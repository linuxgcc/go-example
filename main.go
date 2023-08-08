package main
import (
	"fmt"
)
type context struct {
}
type MapStr map[string]interface{}

type message struct {
	context context
	event   MapStr
	events  []MapStr
}
type syncPublisher struct {
}
type syncClient func(message) bool

type eventPublisher interface {
	PublishEvent(ctx *context, event MapStr) bool
	PublishEvents(ctx *context, events []MapStr) bool
}

func (c syncClient) PublishEvent(ctx *context, event MapStr) bool {
	return c(message{context: *ctx, event: event})
}
func (c syncClient) PublishEvents(ctx *context, events []MapStr) bool {
	return c(message{context: *ctx, events: events})
}

func (p *syncPublisher) client() eventPublisher {
	return syncClient(p.forward)
}
func (p *syncPublisher) forward(m message) bool {
	return true
}

func main() {
	p := syncPublisher{}
	cc := p.client()
	fmt.Println("hello word!")
}
