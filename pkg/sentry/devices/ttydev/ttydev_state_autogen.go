// automatically generated by stateify.

package ttydev

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (t *ttyDevice) StateTypeName() string {
	return "pkg/sentry/devices/ttydev.ttyDevice"
}

func (t *ttyDevice) StateFields() []string {
	return []string{}
}

func (t *ttyDevice) beforeSave() {}

// +checklocksignore
func (t *ttyDevice) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
}

func (t *ttyDevice) afterLoad(context.Context) {}

// +checklocksignore
func (t *ttyDevice) StateLoad(ctx context.Context, stateSourceObject state.Source) {
}

func init() {
	state.Register((*ttyDevice)(nil))
}
