package Subscript

import (
	"fmt"
	"sync"
)

func main()  {

}

type Subscript struct {
	mutx  sync.Mutex 						//读写锁替换
	MessageChans map[string][]chan string   
	Message map[string]string 				//不存消息的话可以删除
}

func NewSubscript() *Subscript {
	return &Subscript{
		mutx: sync.Mutex{},
		MessageChans: make(map[string][]chan string),
		Message: make(map[string]string),
	}
}

func (s *Subscript) Subscription(receive chan string, topic string) {
	s.mutx.Lock()
	defer s.mutx.Unlock()
	if _,ok := s.Message[topic];!ok {
		s.Message[topic] = ""
	}
	if v,ok := s.MessageChans[topic];ok {
		v = append(v, receive)
		s.MessageChans[topic] = v
	} else {
		chans := make([]chan string,0)
		chans = append(chans,receive)
		s.MessageChans[topic] = chans
	}
}

func (s *Subscript) NewMessage(topic string,message string) {
	s.mutx.Lock()
	defer s.mutx.Unlock()
	if _,ok := s.Message[topic]; ok {
		receives := s.MessageChans[topic]
		for _,rec := range receives {
			rec <- message //用go func去防阻塞
		}
	} else {
		fmt.Println("none sub")
	}
}

//func ()  {
//
//} 取消订阅
