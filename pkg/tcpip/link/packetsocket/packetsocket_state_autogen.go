// automatically generated by stateify.

package packetsocket

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (e *endpoint) StateTypeName() string {
	return "pkg/tcpip/link/packetsocket.endpoint"
}

func (e *endpoint) StateFields() []string {
	return []string{
		"Endpoint",
	}
}

func (e *endpoint) beforeSave() {}

// +checklocksignore
func (e *endpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.Endpoint)
}

func (e *endpoint) afterLoad(context.Context) {}

// +checklocksignore
func (e *endpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.Endpoint)
}

func init() {
	state.Register((*endpoint)(nil))
}