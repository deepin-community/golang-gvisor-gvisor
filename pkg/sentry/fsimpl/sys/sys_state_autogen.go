// automatically generated by stateify.

package sys

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (r *dirRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.dirRefs"
}

func (r *dirRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *dirRefs) beforeSave() {}

// +checklocksignore
func (r *dirRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *dirRefs) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(func() { r.afterLoad(ctx) })
}

func (i *kcovInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.kcovInode"
}

func (i *kcovInode) StateFields() []string {
	return []string{
		"InodeAttrs",
		"InodeNoopRefCount",
		"InodeNotAnonymous",
		"InodeNotDirectory",
		"InodeNotSymlink",
		"InodeWatches",
		"implStatFS",
	}
}

func (i *kcovInode) beforeSave() {}

// +checklocksignore
func (i *kcovInode) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.InodeAttrs)
	stateSinkObject.Save(1, &i.InodeNoopRefCount)
	stateSinkObject.Save(2, &i.InodeNotAnonymous)
	stateSinkObject.Save(3, &i.InodeNotDirectory)
	stateSinkObject.Save(4, &i.InodeNotSymlink)
	stateSinkObject.Save(5, &i.InodeWatches)
	stateSinkObject.Save(6, &i.implStatFS)
}

func (i *kcovInode) afterLoad(context.Context) {}

// +checklocksignore
func (i *kcovInode) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.InodeAttrs)
	stateSourceObject.Load(1, &i.InodeNoopRefCount)
	stateSourceObject.Load(2, &i.InodeNotAnonymous)
	stateSourceObject.Load(3, &i.InodeNotDirectory)
	stateSourceObject.Load(4, &i.InodeNotSymlink)
	stateSourceObject.Load(5, &i.InodeWatches)
	stateSourceObject.Load(6, &i.implStatFS)
}

func (fd *kcovFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.kcovFD"
}

func (fd *kcovFD) StateFields() []string {
	return []string{
		"FileDescriptionDefaultImpl",
		"NoLockFD",
		"vfsfd",
		"inode",
		"kcov",
	}
}

func (fd *kcovFD) beforeSave() {}

// +checklocksignore
func (fd *kcovFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(1, &fd.NoLockFD)
	stateSinkObject.Save(2, &fd.vfsfd)
	stateSinkObject.Save(3, &fd.inode)
	stateSinkObject.Save(4, &fd.kcov)
}

func (fd *kcovFD) afterLoad(context.Context) {}

// +checklocksignore
func (fd *kcovFD) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(1, &fd.NoLockFD)
	stateSourceObject.Load(2, &fd.vfsfd)
	stateSourceObject.Load(3, &fd.inode)
	stateSourceObject.Load(4, &fd.kcov)
}

func (fsType *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.FilesystemType"
}

func (fsType *FilesystemType) StateFields() []string {
	return []string{}
}

func (fsType *FilesystemType) beforeSave() {}

// +checklocksignore
func (fsType *FilesystemType) StateSave(stateSinkObject state.Sink) {
	fsType.beforeSave()
}

func (fsType *FilesystemType) afterLoad(context.Context) {}

// +checklocksignore
func (fsType *FilesystemType) StateLoad(ctx context.Context, stateSourceObject state.Source) {
}

func (i *InternalData) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.InternalData"
}

func (i *InternalData) StateFields() []string {
	return []string{
		"ProductName",
		"EnableTPUProxyPaths",
		"TestSysfsPathPrefix",
	}
}

func (i *InternalData) beforeSave() {}

// +checklocksignore
func (i *InternalData) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.ProductName)
	stateSinkObject.Save(1, &i.EnableTPUProxyPaths)
	stateSinkObject.Save(2, &i.TestSysfsPathPrefix)
}

func (i *InternalData) afterLoad(context.Context) {}

// +checklocksignore
func (i *InternalData) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.ProductName)
	stateSourceObject.Load(1, &i.EnableTPUProxyPaths)
	stateSourceObject.Load(2, &i.TestSysfsPathPrefix)
}

func (fs *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.filesystem"
}

func (fs *filesystem) StateFields() []string {
	return []string{
		"Filesystem",
		"devMinor",
	}
}

func (fs *filesystem) beforeSave() {}

// +checklocksignore
func (fs *filesystem) StateSave(stateSinkObject state.Sink) {
	fs.beforeSave()
	stateSinkObject.Save(0, &fs.Filesystem)
	stateSinkObject.Save(1, &fs.devMinor)
}

func (fs *filesystem) afterLoad(context.Context) {}

// +checklocksignore
func (fs *filesystem) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fs.Filesystem)
	stateSourceObject.Load(1, &fs.devMinor)
}

func (d *dir) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.dir"
}

func (d *dir) StateFields() []string {
	return []string{
		"dirRefs",
		"InodeAlwaysValid",
		"InodeAttrs",
		"InodeDirectoryNoNewChildren",
		"InodeNotAnonymous",
		"InodeNotSymlink",
		"InodeTemporary",
		"InodeWatches",
		"OrderedChildren",
		"locks",
	}
}

func (d *dir) beforeSave() {}

// +checklocksignore
func (d *dir) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.dirRefs)
	stateSinkObject.Save(1, &d.InodeAlwaysValid)
	stateSinkObject.Save(2, &d.InodeAttrs)
	stateSinkObject.Save(3, &d.InodeDirectoryNoNewChildren)
	stateSinkObject.Save(4, &d.InodeNotAnonymous)
	stateSinkObject.Save(5, &d.InodeNotSymlink)
	stateSinkObject.Save(6, &d.InodeTemporary)
	stateSinkObject.Save(7, &d.InodeWatches)
	stateSinkObject.Save(8, &d.OrderedChildren)
	stateSinkObject.Save(9, &d.locks)
}

func (d *dir) afterLoad(context.Context) {}

// +checklocksignore
func (d *dir) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.dirRefs)
	stateSourceObject.Load(1, &d.InodeAlwaysValid)
	stateSourceObject.Load(2, &d.InodeAttrs)
	stateSourceObject.Load(3, &d.InodeDirectoryNoNewChildren)
	stateSourceObject.Load(4, &d.InodeNotAnonymous)
	stateSourceObject.Load(5, &d.InodeNotSymlink)
	stateSourceObject.Load(6, &d.InodeTemporary)
	stateSourceObject.Load(7, &d.InodeWatches)
	stateSourceObject.Load(8, &d.OrderedChildren)
	stateSourceObject.Load(9, &d.locks)
}

func (d *cgroupDir) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.cgroupDir"
}

func (d *cgroupDir) StateFields() []string {
	return []string{
		"dir",
	}
}

func (d *cgroupDir) beforeSave() {}

// +checklocksignore
func (d *cgroupDir) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.dir)
}

func (d *cgroupDir) afterLoad(context.Context) {}

// +checklocksignore
func (d *cgroupDir) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.dir)
}

func (c *cpuFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.cpuFile"
}

func (c *cpuFile) StateFields() []string {
	return []string{
		"implStatFS",
		"DynamicBytesFile",
		"maxCores",
	}
}

func (c *cpuFile) beforeSave() {}

// +checklocksignore
func (c *cpuFile) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.implStatFS)
	stateSinkObject.Save(1, &c.DynamicBytesFile)
	stateSinkObject.Save(2, &c.maxCores)
}

func (c *cpuFile) afterLoad(context.Context) {}

// +checklocksignore
func (c *cpuFile) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.implStatFS)
	stateSourceObject.Load(1, &c.DynamicBytesFile)
	stateSourceObject.Load(2, &c.maxCores)
}

func (i *implStatFS) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.implStatFS"
}

func (i *implStatFS) StateFields() []string {
	return []string{}
}

func (i *implStatFS) beforeSave() {}

// +checklocksignore
func (i *implStatFS) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
}

func (i *implStatFS) afterLoad(context.Context) {}

// +checklocksignore
func (i *implStatFS) StateLoad(ctx context.Context, stateSourceObject state.Source) {
}

func (s *staticFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.staticFile"
}

func (s *staticFile) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"StaticData",
	}
}

func (s *staticFile) beforeSave() {}

// +checklocksignore
func (s *staticFile) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.DynamicBytesFile)
	stateSinkObject.Save(1, &s.StaticData)
}

func (s *staticFile) afterLoad(context.Context) {}

// +checklocksignore
func (s *staticFile) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.DynamicBytesFile)
	stateSourceObject.Load(1, &s.StaticData)
}

func (hf *hostFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.hostFile"
}

func (hf *hostFile) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"hostPath",
	}
}

func (hf *hostFile) beforeSave() {}

// +checklocksignore
func (hf *hostFile) StateSave(stateSinkObject state.Sink) {
	hf.beforeSave()
	stateSinkObject.Save(0, &hf.DynamicBytesFile)
	stateSinkObject.Save(1, &hf.hostPath)
}

func (hf *hostFile) afterLoad(context.Context) {}

// +checklocksignore
func (hf *hostFile) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &hf.DynamicBytesFile)
	stateSourceObject.Load(1, &hf.hostPath)
}

func init() {
	state.Register((*dirRefs)(nil))
	state.Register((*kcovInode)(nil))
	state.Register((*kcovFD)(nil))
	state.Register((*FilesystemType)(nil))
	state.Register((*InternalData)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*dir)(nil))
	state.Register((*cgroupDir)(nil))
	state.Register((*cpuFile)(nil))
	state.Register((*implStatFS)(nil))
	state.Register((*staticFile)(nil))
	state.Register((*hostFile)(nil))
}
