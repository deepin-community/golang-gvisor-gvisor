// automatically generated by stateify.

package inet

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (a *abstractEndpoint) StateTypeName() string {
	return "pkg/sentry/inet.abstractEndpoint"
}

func (a *abstractEndpoint) StateFields() []string {
	return []string{
		"ep",
		"socket",
		"name",
		"ns",
	}
}

func (a *abstractEndpoint) beforeSave() {}

// +checklocksignore
func (a *abstractEndpoint) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.ep)
	stateSinkObject.Save(1, &a.socket)
	stateSinkObject.Save(2, &a.name)
	stateSinkObject.Save(3, &a.ns)
}

func (a *abstractEndpoint) afterLoad(context.Context) {}

// +checklocksignore
func (a *abstractEndpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.ep)
	stateSourceObject.Load(1, &a.socket)
	stateSourceObject.Load(2, &a.name)
	stateSourceObject.Load(3, &a.ns)
}

func (a *AbstractSocketNamespace) StateTypeName() string {
	return "pkg/sentry/inet.AbstractSocketNamespace"
}

func (a *AbstractSocketNamespace) StateFields() []string {
	return []string{
		"endpoints",
	}
}

func (a *AbstractSocketNamespace) beforeSave() {}

// +checklocksignore
func (a *AbstractSocketNamespace) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.endpoints)
}

func (a *AbstractSocketNamespace) afterLoad(context.Context) {}

// +checklocksignore
func (a *AbstractSocketNamespace) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.endpoints)
}

func (t *TCPBufferSize) StateTypeName() string {
	return "pkg/sentry/inet.TCPBufferSize"
}

func (t *TCPBufferSize) StateFields() []string {
	return []string{
		"Min",
		"Default",
		"Max",
	}
}

func (t *TCPBufferSize) beforeSave() {}

// +checklocksignore
func (t *TCPBufferSize) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.Min)
	stateSinkObject.Save(1, &t.Default)
	stateSinkObject.Save(2, &t.Max)
}

func (t *TCPBufferSize) afterLoad(context.Context) {}

// +checklocksignore
func (t *TCPBufferSize) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.Min)
	stateSourceObject.Load(1, &t.Default)
	stateSourceObject.Load(2, &t.Max)
}

func (n *Namespace) StateTypeName() string {
	return "pkg/sentry/inet.Namespace"
}

func (n *Namespace) StateFields() []string {
	return []string{
		"inode",
		"creator",
		"isRoot",
		"userNS",
		"abstractSockets",
	}
}

func (n *Namespace) beforeSave() {}

// +checklocksignore
func (n *Namespace) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.inode)
	stateSinkObject.Save(1, &n.creator)
	stateSinkObject.Save(2, &n.isRoot)
	stateSinkObject.Save(3, &n.userNS)
	stateSinkObject.Save(4, &n.abstractSockets)
}

// +checklocksignore
func (n *Namespace) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.inode)
	stateSourceObject.LoadWait(1, &n.creator)
	stateSourceObject.Load(2, &n.isRoot)
	stateSourceObject.Load(3, &n.userNS)
	stateSourceObject.Load(4, &n.abstractSockets)
	stateSourceObject.AfterLoad(func() { n.afterLoad(ctx) })
}

func (r *namespaceRefs) StateTypeName() string {
	return "pkg/sentry/inet.namespaceRefs"
}

func (r *namespaceRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *namespaceRefs) beforeSave() {}

// +checklocksignore
func (r *namespaceRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *namespaceRefs) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(func() { r.afterLoad(ctx) })
}

func init() {
	state.Register((*abstractEndpoint)(nil))
	state.Register((*AbstractSocketNamespace)(nil))
	state.Register((*TCPBufferSize)(nil))
	state.Register((*Namespace)(nil))
	state.Register((*namespaceRefs)(nil))
}
