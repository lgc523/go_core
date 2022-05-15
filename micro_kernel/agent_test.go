package micro_kernel

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type DemoCollector struct {
	evtReceiver EventReceiver
	agtCtx      context.Context
	stopChan    chan struct{}
	name        string
	content     string
}

func NewCollect(name string, content string) *DemoCollector {
	return &DemoCollector{
		stopChan: make(chan struct{}),
		name:     name,
		content:  content,
	}
}

func (c *DemoCollector) Init(evtReceiver EventReceiver) error {
	fmt.Println("initialize collect", c.name)
	c.evtReceiver = evtReceiver
	return nil
}

func (c *DemoCollector) Start(agtCtx context.Context) error {
	fmt.Println("start collect", c.name)
	for {
		select {
		case <-agtCtx.Done():
			c.stopChan <- struct{}{}
			break
		default:
			time.Sleep(time.Millisecond * 50)
			c.evtReceiver.OnEvent(
				Event{c.name, c.content})
		}
	}
}

func (c *DemoCollector) Stop() error {
	fmt.Println("stop collect", c.name)
	select {
	case <-c.stopChan:
		return nil
	case <-time.After(time.Second):
		return errors.New("failed to stop for timeout")
	}
}

func (c *DemoCollector) Destroy() error {
	fmt.Println(c.name, "released resources.")
	return nil
}

func TestAgent(t *testing.T) {
	agt := NewAgent(523)
	c1 := NewCollect("c1", "1")
	c2 := NewCollect("c2", "2")
	c3 := NewCollect("c3", "3")
	agt.registerCollector("c1", c1)
	agt.registerCollector("c2", c2)
	agt.registerCollector("c3", c3)
	agt.Start()
	fmt.Println(agt.state)
	time.Sleep(time.Second)
	agt.Stop()
	agt.Destroy()
}
