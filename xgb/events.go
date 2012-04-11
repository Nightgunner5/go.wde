package xgb

import (
	"fmt"
	"io"
	"code.google.com/p/jamslam-x-go-binding/xgb"
	"sync"
)

var eventChans = map[xgb.Id] chan interface{}{}
var eventLock sync.Mutex

func handleEvents(conn *xgb.Conn) {
	for {
		e, err := conn.WaitForEvent()
		fmt.Printf("%T, %s\n", e, err)

		if err == io.EOF {
			break
		}

		// var id xgb.Id

		// eventLock.Lock()
		// ch, ok := eventChans[id]
		// if ok {
		// 	ch <- e
		// }
		// eventLock.Unlock()
	}

	eventLock.Lock()
	defer eventLock.Unlock()
	for _, ch := range eventChans {
		close(ch)
	}
}

func registerId(ch chan interface{}, id xgb.Id) {
	eventLock.Lock()
	defer eventLock.Unlock()
	eventChans[id] = ch
}

func (w *Window) EventChan() (events <-chan interface{}) {
	ch := make(chan interface{})
	registerId(ch, w.id)
	events = ch

	return
}