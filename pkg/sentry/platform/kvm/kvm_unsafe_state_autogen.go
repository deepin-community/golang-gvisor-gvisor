// automatically generated by stateify.

//go:build go1.18 && go1.18
// +build go1.18,go1.18

package kvm

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (p *machineAtomicPtr) StateTypeName() string {
	return "pkg/sentry/platform/kvm.machineAtomicPtr"
}

func (p *machineAtomicPtr) StateFields() []string {
	return []string{
		"ptr",
	}
}

func (p *machineAtomicPtr) beforeSave() {}

// +checklocksignore
func (p *machineAtomicPtr) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	var ptrValue *machine
	ptrValue = p.savePtr()
	stateSinkObject.SaveValue(0, ptrValue)
}

func (p *machineAtomicPtr) afterLoad(context.Context) {}

// +checklocksignore
func (p *machineAtomicPtr) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new(*machine), func(y any) { p.loadPtr(ctx, y.(*machine)) })
}

func init() {
	state.Register((*machineAtomicPtr)(nil))
}
