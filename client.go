package gs

import (
	"errors"
	"net"
	"sync"
	"sync/atomic"
)

var (
	ErrConnClosing = errors.New("use of closed network connection")
)

type Client struct {
	id         int64
	server     *Server
	conn       net.Conn
	closeOnece sync.Once
	closeFlag  int32
	closeChan  chan struct{}
	recv       chan interface{}
	secv       chan interface{}
	callBack   ClientCallBack
}

type ClientCallBack interface {
	OnConnect(*Client) bool
	OnMessage(*Client, interface{})
	OnClose(*Client)
}

func newClient(server *Server, conn net.Conn) *Client {
	return &Client{
		server:    server,
		conn:      conn,
		closeChan: make(chan struct{}),
		recv:      make(chan interface{}, 20),
		secv:      make(chan interface{}, 20),
		callBack:  server.ClientCallBack(),
	}
}

func (c *Client) Close() {
	c.closeOnece.Do(func() {
		atomic.StoreInt32(&c.closeFlag, 1)
		close(c.closeChan)
		close(c.recv)
		close(c.secv)
		c.server.RemoveClient(c.id)
		c.callBack.OnClose(c)
		c.conn.Close()
	})
}

func (c *Client) ClientId() int64 {
	return c.id
}

func (c *Client) IsClose() bool {
	return atomic.LoadInt32(&c.closeFlag) == 1
}

func (c *Client) WriteMsg(p interface{}) (err error) {
	if c.IsClose() {
		return ErrConnClosing
	}
	c.secv <- p
	return nil
}

func (c *Client) Do() {
	if !c.callBack.OnConnect(c) {
		return
	}
	atomic.AddInt64(&c.server.clientIDSequence, 1)
	c.id = c.server.clientIDSequence
	c.server.AddClient(c.id, c)
	c.server.waitGroupFunc(c.readLoop)
	c.server.waitGroupFunc(c.writeLoop)
	c.server.waitGroupFunc(c.handleLoop)
}

func (c *Client) readLoop() {
	defer func() {
		recover()
		c.Close()
	}()

	for {
		select {
		case <-c.server.exitChan:
			return
		case <-c.closeChan:
			return
		default:
		}
		b, err := c.server.protocol.Read(c.conn)
		if err != nil {
			return
		}
		r, err := c.server.protocol.Unmarshal(b)
		if err != nil {
			return
		}
		c.recv <- r
	}
}

func (c *Client) writeLoop() {
	defer func() {
		recover()
		c.Close()
	}()
	for {
		select {
		case <-c.server.exitChan:
			return
		case <-c.closeChan:
			return
		case p := <-c.secv:
			if c.IsClose() {
				return
			}
			b, err := c.server.protocol.Marshal(p)
			if err != nil {
				break
			}
			if err := c.server.protocol.Write(c.conn, b...); err != nil {
				return
			}
		}
	}
}

func (c *Client) handleLoop() {
	defer func() {
		recover()
		c.Close()
	}()
	for {
		select {
		case <-c.server.exitChan:
			return
		case <-c.closeChan:
			return
		case p := <-c.recv:
			c.callBack.OnMessage(c, p)
		}
	}
}
