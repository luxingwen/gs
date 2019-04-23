package sproto

type SpProtocol struct {
	Type     int32
	Name     string
	Request  reflect.Type
	Response reflect.Type
}

func (p *SpProtocol) HasRequest() bool {
	return p.Request != nil
}

func (p *SpProtocol) HasResponse() bool {
	return p.Response != nil
}
