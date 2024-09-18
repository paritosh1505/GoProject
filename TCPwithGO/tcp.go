package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Client[T any] struct {
	sendPort string
	channel  chan T
}

func NewSender[T any](port string) (*Client[T], error) {
	connSender, error := net.Dial("tcp",port)
	if error!=nil{
		//fmt.Println("Erro while sending the data")
		return nil,fmt.Errorf("error in COnnection %w",error)
	}
	senderinit := &Client[T]{
		sendPort :port,
		channel: make(chan T),
	}
	fmt.Println("Connection value is",connSender)
	go senderinit.writeConnDetail(connSender)
	return senderinit,nil
}
func (c *Client[T]) writeConnDetail(conval net.Conn){ 
defer conval.Close()
for data:= range c.channel{
	gob.NewEncoder(conval).Encode(data)
}

}
type Server[T any] struct{
	acceptPort string
	channel chan T 
} 

func NewServer[T any](portval string) (*Server[T],error){
    acceptSender,err := net.Listen("tcp",portval)
	if err!=nil{
		fmt.Println("Error while receving the connection")
		return nil,fmt.Errorf("error in server Connection")
	}
	serverval:=&Server[T]{
		acceptPort: portval,
		channel: make(chan T),

	}
	go serverval.acceptConn(acceptSender)
	fmt.Println("Server Connection is",acceptSender)
	return serverval,nil

}
 

func (s *Server[T]) acceptConn(ln net.Listener){
	defer ln.Close()
	for{
		conn,err:=ln.Accept()
		if err!=nil{
			fmt.Println("Error in connction.retryng....")
		}
		fmt.Println("Remote address is",conn.RemoteAddr())
		go s.handleConnectionVal(conn)
	}
	
}
func(s *Server[T]) handleConnectionVal(conval net.Conn){
defer conval.Close()
fmt.Println("************")
//bufferval:= make([]byte,1024)
var data T;

for{
	gob.NewDecoder(conval).Decode(&data)
	s.channel<-data
	
}
fmt.Println("Outside")

}