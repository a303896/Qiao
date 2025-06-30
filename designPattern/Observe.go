package designPattern

import (
	"context"
	"fmt"
	"sync"
)

type Event struct {
	Topic string
	Value interface{}
}

type Observer interface {
	Onchange(ctx context.Context, e *Event) error
}

type EventBus interface {
	Subscribe(topic string, ob Observer)
	Unsubscribe(topic string, ob Observer)
	Publish(ctx context.Context, e *Event)
}

type BaseObserver struct {
	Name string
}

func NewBaseObserver(name string) *BaseObserver {
	return &BaseObserver{name}
}

func (o *BaseObserver) Onchange(ctx context.Context, e *Event) error {
	fmt.Printf("%s observed by %s\n", o.Name, e.Topic)
	return nil
}

type BaseEventBus struct {
	mutex     sync.RWMutex
	observers map[string]map[Observer]bool
}

func NewBaseEventBus() *BaseEventBus {
	return &BaseEventBus{observers: make(map[string]map[Observer]bool)}
}

func (b *BaseEventBus) Subscribe(topic string, ob Observer) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	_, ok := b.observers[topic]
	if !ok {
		b.observers[topic] = make(map[Observer]bool)
	}
	b.observers[topic][ob] = true
}

func (b *BaseEventBus) Unsubscribe(topic string, ob Observer) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	_, ok := b.observers[topic]
	if ok {
		delete(b.observers[topic], ob)
	}
}

// 同步模式
type SyncBaseEventBus struct {
	BaseEventBus
}

func (b *SyncBaseEventBus) Publish(ctx context.Context, e *Event) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	subscribers := b.observers[e.Topic]
	errs := make(map[Observer]error)
	for subscriber := range subscribers {
		if err := subscriber.Onchange(ctx, e); err != nil {
			errs[subscriber] = err
		}
	}
	b.Handler(errs)
}

// 错误处理
func (b *SyncBaseEventBus) Handler(errors map[Observer]error) {
	for ob, err := range errors {
		fmt.Printf("observer %v error: %v\n", ob, err)
	}
}
