package ch41

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

const (
	Waiting = iota
	Running
)

var WrongStateError = errors.New("can not take the operation int the current state")

type CollectorsError struct {
	CollectorsError []error
}

func (ce CollectorsError) Error() string {
	var messages []string
	for _, err := range ce.CollectorsError {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, ";")
}

type Event struct {
	Source  string
	Content string
}

type EventReceiver interface {
	OnEvent(event Event)
}

type Collector interface {
	Init(evtReceiver EventReceiver) error
	Start(context context.Context) error
	Stop() error
	Destory() error
}

type Agent struct {
	collectors map[string]Collector
	eventBuf   chan Event
	cancel     context.CancelFunc
	ctx        context.Context
	state      int
}

func (agt *Agent) OnEvent(event Event) {
	agt.eventBuf <- event
}

func (agt *Agent) EventProcessGroutine() {
	var eventSeg [10]Event
	for {
		for i := 0; i < 10; i++ {
			select {
			case eventSeg[i] = <-agt.eventBuf:
			case <-agt.ctx.Done():
				return
			}
		}
		fmt.Println(eventSeg)
	}
}

func (agt *Agent) RegisterCollector(name string, collector Collector) error {
	if agt.state != Waiting {
		return WrongStateError
	}
	agt.collectors[name] = collector
	return collector.Init(agt)
}

func (agt *Agent) startCollectors() error {
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
				errs.CollectorsError = append(errs.CollectorsError, errors.New(name+": "+err.Error()))
			}
		}(name, collector, agt.ctx)
	}
	if len(errs.CollectorsError) == 0 {
		return nil
	}
	return errs
}

func (agt *Agent) stopCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Destory(); err != nil {
			errs.CollectorsError = append(errs.CollectorsError, errors.New(name+":"+err.Error()))
		}
	}
	if len(errs.CollectorsError) == 0 {
		return nil
	}
	return errs
}

func (agt *Agent) destoryCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Destory(); err != nil {
			errs.CollectorsError = append(errs.CollectorsError, errors.New(name+":"+err.Error()))
		}
	}
	if len(errs.CollectorsError) == 0 {
		return nil
	}
	return errs
}

func (agt *Agent) Start() error {
	if agt.state != Waiting {
		return WrongStateError
	}
	agt.state = Running
	agt.ctx, agt.cancel = context.WithCancel(context.Background())
	go agt.EventProcessGroutine()
	return agt.startCollectors()
}

func (agt *Agent) Stop() error {
	if agt.state != Running {
		return WrongStateError
	}
	agt.state = Waiting
	agt.cancel()
	return agt.stopCollectors()
}

func (agt *Agent) Destory() error {
	if agt.state != Waiting {
		return WrongStateError
	}
	return agt.destoryCollectors()
}

func NewAgent(sizeEventBuf int) *Agent {
	agt := Agent{
		collectors: map[string]Collector{},
		eventBuf:   make(chan Event, sizeEventBuf),
		state:      Waiting,
	}
	return &agt
}
