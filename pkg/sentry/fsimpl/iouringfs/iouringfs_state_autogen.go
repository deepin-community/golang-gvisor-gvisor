// automatically generated by stateify.

package iouringfs

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (fd *FileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/iouringfs.FileDescription"
}

func (fd *FileDescription) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
		"rbmf",
		"sqemf",
		"running",
		"ioRings",
		"remap",
	}
}

// +checklocksignore
func (fd *FileDescription) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &fd.NoLockFD)
	stateSinkObject.Save(4, &fd.rbmf)
	stateSinkObject.Save(5, &fd.sqemf)
	stateSinkObject.Save(6, &fd.running)
	stateSinkObject.Save(7, &fd.ioRings)
	stateSinkObject.Save(8, &fd.remap)
}

// +checklocksignore
func (fd *FileDescription) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &fd.NoLockFD)
	stateSourceObject.Load(4, &fd.rbmf)
	stateSourceObject.Load(5, &fd.sqemf)
	stateSourceObject.Load(6, &fd.running)
	stateSourceObject.Load(7, &fd.ioRings)
	stateSourceObject.Load(8, &fd.remap)
	stateSourceObject.AfterLoad(func() { fd.afterLoad(ctx) })
}

func (sqemf *sqEntriesFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/iouringfs.sqEntriesFile"
}

func (sqemf *sqEntriesFile) StateFields() []string {
	return []string{
		"fr",
	}
}

func (sqemf *sqEntriesFile) beforeSave() {}

// +checklocksignore
func (sqemf *sqEntriesFile) StateSave(stateSinkObject state.Sink) {
	sqemf.beforeSave()
	stateSinkObject.Save(0, &sqemf.fr)
}

func (sqemf *sqEntriesFile) afterLoad(context.Context) {}

// +checklocksignore
func (sqemf *sqEntriesFile) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &sqemf.fr)
}

func (rbmf *ringsBufferFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/iouringfs.ringsBufferFile"
}

func (rbmf *ringsBufferFile) StateFields() []string {
	return []string{
		"fr",
	}
}

func (rbmf *ringsBufferFile) beforeSave() {}

// +checklocksignore
func (rbmf *ringsBufferFile) StateSave(stateSinkObject state.Sink) {
	rbmf.beforeSave()
	stateSinkObject.Save(0, &rbmf.fr)
}

func (rbmf *ringsBufferFile) afterLoad(context.Context) {}

// +checklocksignore
func (rbmf *ringsBufferFile) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &rbmf.fr)
}

func init() {
	state.Register((*FileDescription)(nil))
	state.Register((*sqEntriesFile)(nil))
	state.Register((*ringsBufferFile)(nil))
}
