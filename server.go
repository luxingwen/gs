package gs

import (
	"net"
	"runtime"
	"sync"
)

type Server struct {
	ln               net.Listener
	wg               *sync.WaitGroup
	exitChan         chan struct{}
	ClientCallBack   func() ClientCallBack
	protocol         Protocol
	clientIDSequence int64
	clientLock       sync.RWMutex
	clients          map[int64]*Client
}

func NewServer(ln net.Listener, f func() ClientCallBack, protocol Protocol) *Server {
	s := &Server{
		ln:             ln,
		wg:             new(sync.WaitGroup),
		exitChan:       make(chan struct{}),
		ClientCallBack: f,
		protocol:       protocol,
		clients:        make(map[int64]*Client),
	}
	return s
}

func (s *Server) Start() {
	s.wg.Add(1)
	defer func() {
		s.ln.Close()
		s.wg.Done()

	}()
	for {
		clientConn, err := s.ln.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				runtime.Gosched()
				continue
			}
			return
		}
		s.wg.Add(1)
		go func() {
			newClient(s, clientConn).Do()
			s.wg.Done()
		}()
	}
}

func (s *Server) Stop() {
	s.ln.Close()
	close(s.exitChan)
	s.wg.Wait()
}

func (s *Server) waitGroupFunc(f func()) {
	s.wg.Add(1)
	go func() {
		f()
		s.wg.Done()
	}()
}

func (s *Server) AddClient(clientId int64, c *Client) {
	s.clientLock.Lock()
	s.clients[clientId] = c
	s.clientLock.Unlock()
}

func (s *Server) RemoveClient(clientId int64) {
	s.clientLock.Lock()
	defer s.clientLock.Unlock()
	if _, ok := s.clients[clientId]; !ok {
		return
	}
	delete(s.clients, clientId)
}
