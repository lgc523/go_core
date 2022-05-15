package micro_kernel

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

type State int

const (
	Running State = iota
	Waiting
)

var WrongStataError = errors.New("can not take the operation in the current state")

type CollectorsError struct {
	CollectorsError []error
}

func (ce CollectorsError) Error() string {
	var sts []string
	for _, err := range ce.CollectorsError {
		sts = append(sts, err.Error())
	}
	return strings.Join(sts, ",")
}

type Event struct {
	Source  string
	Content string
}
type EventReceiver interface {
	OnEvent(evt Event)
}
type Collector interface {
	Init(evtReceiver EventReceiver) error
	Start(agtCtx context.Context) error
	Stop() error
	Destroy() error
}

type Agent struct {
	collectors map[string]Collector
	evtBuf     chan Event
	cancel     context.CancelFunc
	ctx        context.Context
	state      State
}

func (agt *Agent) EventProcessGoroutine() {
	var evtSeg [10]Event
	for {
		for i := 0; i < 10; i++ {
			select {
			case evtSeg[i] = <-agt.evtBuf:
			case <-agt.ctx.Done():
				return

			}
		}
		fmt.Println(evtSeg)
	}
}

func NewAgent(sizeEvtBuf int) *Agent {
	agt := Agent{
		collectors: map[string]Collector{},
		evtBuf:     make(chan Event, sizeEvtBuf),
		state:      Waiting,
	}
	return &agt
}

func (agt *Agent) registerCollector(name string, collector Collector) error {
	if agt.state != Waiting {
		return WrongStataError
	}
	agt.collectors[name] = collector
	return collector.Init(agt)
}

func (agt *Agent) OnEvent(evt Event) {
	agt.evtBuf <- evt
}

func (agt *Agent) StartCollectors() error {
	var err error
	var errs CollectorsError
	var mutex sync.Mutex

	for name, collector := range agt.collectors {
		go func(name string, collector Collector, ctx context.Context) {
			defer func() {
				mutex.Unlock()
			}()
			err = collector.Start(ctx)
			mutex.Lock()
			if err != nil {
				errs.CollectorsError = append(errs.CollectorsError, errors.New(name+":"+err.Error()))
			}
		}(name, collector, agt.ctx)
	}
	return errs
}

func (agt *Agent) stopCollector() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Stop(); err != nil {
			errs.CollectorsError = append(errs.CollectorsError, errors.New(name+":"+err.Error()))
		}
	}
	return errs
}

func (agt *Agent) destroyCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Stop(); err != nil {
			errs.CollectorsError = append(errs.CollectorsError, errors.New(name+":"+err.Error()))
		}
	}
	return errs
}

func (agt *Agent) Start() error {
	if agt.state != Waiting {
		return WrongStataError
	}
	agt.state = Running
	agt.ctx, agt.cancel = context.WithCancel(context.Background())
	go agt.EventProcessGoroutine()
	return agt.StartCollectors()
}
func (agt *Agent) Stop() error {
	if agt.state != Running {
		return WrongStataError
	}
	agt.state = Waiting
	agt.cancel()
	return agt.stopCollector()
}
func (agt *Agent) Destroy() error {
	if agt.state != Waiting {
		return WrongStataError
	}
	return agt.destroyCollectors()
}
