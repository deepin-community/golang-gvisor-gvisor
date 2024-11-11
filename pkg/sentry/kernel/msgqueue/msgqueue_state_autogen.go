// automatically generated by stateify.

package msgqueue

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (l *msgList) StateTypeName() string {
	return "pkg/sentry/kernel/msgqueue.msgList"
}

func (l *msgList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *msgList) beforeSave() {}

// +checklocksignore
func (l *msgList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *msgList) afterLoad(context.Context) {}

// +checklocksignore
func (l *msgList) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *msgEntry) StateTypeName() string {
	return "pkg/sentry/kernel/msgqueue.msgEntry"
}

func (e *msgEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *msgEntry) beforeSave() {}

// +checklocksignore
func (e *msgEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *msgEntry) afterLoad(context.Context) {}

// +checklocksignore
func (e *msgEntry) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (r *Registry) StateTypeName() string {
	return "pkg/sentry/kernel/msgqueue.Registry"
}

func (r *Registry) StateFields() []string {
	return []string{
		"reg",
	}
}

func (r *Registry) beforeSave() {}

// +checklocksignore
func (r *Registry) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.reg)
}

func (r *Registry) afterLoad(context.Context) {}

// +checklocksignore
func (r *Registry) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.reg)
}

func (q *Queue) StateTypeName() string {
	return "pkg/sentry/kernel/msgqueue.Queue"
}

func (q *Queue) StateFields() []string {
	return []string{
		"registry",
		"dead",
		"obj",
		"senders",
		"receivers",
		"messages",
		"sendTime",
		"receiveTime",
		"changeTime",
		"byteCount",
		"messageCount",
		"maxBytes",
		"sendPID",
		"receivePID",
	}
}

func (q *Queue) beforeSave() {}

// +checklocksignore
func (q *Queue) StateSave(stateSinkObject state.Sink) {
	q.beforeSave()
	stateSinkObject.Save(0, &q.registry)
	stateSinkObject.Save(1, &q.dead)
	stateSinkObject.Save(2, &q.obj)
	stateSinkObject.Save(3, &q.senders)
	stateSinkObject.Save(4, &q.receivers)
	stateSinkObject.Save(5, &q.messages)
	stateSinkObject.Save(6, &q.sendTime)
	stateSinkObject.Save(7, &q.receiveTime)
	stateSinkObject.Save(8, &q.changeTime)
	stateSinkObject.Save(9, &q.byteCount)
	stateSinkObject.Save(10, &q.messageCount)
	stateSinkObject.Save(11, &q.maxBytes)
	stateSinkObject.Save(12, &q.sendPID)
	stateSinkObject.Save(13, &q.receivePID)
}

func (q *Queue) afterLoad(context.Context) {}

// +checklocksignore
func (q *Queue) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &q.registry)
	stateSourceObject.Load(1, &q.dead)
	stateSourceObject.Load(2, &q.obj)
	stateSourceObject.Load(3, &q.senders)
	stateSourceObject.Load(4, &q.receivers)
	stateSourceObject.Load(5, &q.messages)
	stateSourceObject.Load(6, &q.sendTime)
	stateSourceObject.Load(7, &q.receiveTime)
	stateSourceObject.Load(8, &q.changeTime)
	stateSourceObject.Load(9, &q.byteCount)
	stateSourceObject.Load(10, &q.messageCount)
	stateSourceObject.Load(11, &q.maxBytes)
	stateSourceObject.Load(12, &q.sendPID)
	stateSourceObject.Load(13, &q.receivePID)
}

func (m *Message) StateTypeName() string {
	return "pkg/sentry/kernel/msgqueue.Message"
}

func (m *Message) StateFields() []string {
	return []string{
		"msgEntry",
		"Type",
		"Text",
		"Size",
	}
}

func (m *Message) beforeSave() {}

// +checklocksignore
func (m *Message) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.msgEntry)
	stateSinkObject.Save(1, &m.Type)
	stateSinkObject.Save(2, &m.Text)
	stateSinkObject.Save(3, &m.Size)
}

func (m *Message) afterLoad(context.Context) {}

// +checklocksignore
func (m *Message) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.msgEntry)
	stateSourceObject.Load(1, &m.Type)
	stateSourceObject.Load(2, &m.Text)
	stateSourceObject.Load(3, &m.Size)
}

func init() {
	state.Register((*msgList)(nil))
	state.Register((*msgEntry)(nil))
	state.Register((*Registry)(nil))
	state.Register((*Queue)(nil))
	state.Register((*Message)(nil))
}