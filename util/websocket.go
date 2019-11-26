package util

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type Connection struct {
	SocketId       string
	wsConn         *websocket.Conn
	inChan         chan []byte
	outJsonChan    chan interface{}
	outMessageChan chan []byte
	closeChan      chan byte
	mutex          sync.Mutex
	isClosed       bool
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		SocketId:       wsConn.RemoteAddr().String(),
		wsConn:         wsConn,
		inChan:         make(chan []byte, 100),
		outJsonChan:    make(chan interface{}, 100),
		outMessageChan: make(chan []byte, 100),
		closeChan:      make(chan byte, 1),
	}

	//启动协程去读取消息
	go conn.readLoop()
	//启动协程去发送消息
	go conn.writeLoop()
	return
}

func (c *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-c.inChan:
	case <-c.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

func (c *Connection) WriteMessage(data []byte) (err error) {
	select {
	case c.outMessageChan <- data:
	case <-c.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

func (c *Connection) WriteJson(data interface{}) (err error) {
	select {
	case c.outJsonChan <- data:
	case <-c.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

func (c *Connection) Close() {
	//线程安全的Close
	c.wsConn.Close()

	//这一行代码只需要执行一次
	c.mutex.Lock()
	if !c.isClosed {
		close(c.closeChan)
		c.isClosed = true
	}
	c.mutex.Unlock()
}

func (c *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		_, data, err = c.wsConn.ReadMessage()
		if err != nil {
			goto ERROR
		}
		select {
		case c.inChan <- data:
		case <-c.closeChan:
			//todo:关闭
			goto ERROR
		}
	}

ERROR:
	//todo:关闭连接操作
	c.Close()
}

func (c *Connection) writeLoop() {
	var (
		message []byte
		json    interface{}
		err     error
	)
	for {
		select {
		case message = <-c.outMessageChan:
			err = c.wsConn.WriteMessage(websocket.TextMessage, message)
		case json = <-c.outJsonChan:
			err = c.wsConn.WriteJSON(json)
		case <-c.closeChan:
			goto ERROR
		}

		if err != nil {
			goto ERROR
		}
	}
ERROR:
	//todo:关闭连接操作
	c.Close()
}
