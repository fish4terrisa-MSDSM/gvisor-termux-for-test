// automatically generated by stateify.

package nvproxy

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (dev *frontendDevice) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.frontendDevice"
}

func (dev *frontendDevice) StateFields() []string {
	return []string{
		"nvp",
		"minor",
	}
}

func (dev *frontendDevice) beforeSave() {}

// +checklocksignore
func (dev *frontendDevice) StateSave(stateSinkObject state.Sink) {
	dev.beforeSave()
	stateSinkObject.Save(0, &dev.nvp)
	stateSinkObject.Save(1, &dev.minor)
}

func (dev *frontendDevice) afterLoad(context.Context) {}

// +checklocksignore
func (dev *frontendDevice) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &dev.nvp)
	stateSourceObject.Load(1, &dev.minor)
}

func (fd *frontendFD) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.frontendFD"
}

func (fd *frontendFD) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
		"dev",
		"containerName",
		"hostFD",
		"memmapFile",
		"queue",
		"clients",
	}
}

// +checklocksignore
func (fd *frontendFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &fd.NoLockFD)
	stateSinkObject.Save(4, &fd.dev)
	stateSinkObject.Save(5, &fd.containerName)
	stateSinkObject.Save(6, &fd.hostFD)
	stateSinkObject.Save(7, &fd.memmapFile)
	stateSinkObject.Save(8, &fd.queue)
	stateSinkObject.Save(9, &fd.clients)
}

// +checklocksignore
func (fd *frontendFD) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &fd.NoLockFD)
	stateSourceObject.Load(4, &fd.dev)
	stateSourceObject.Load(5, &fd.containerName)
	stateSourceObject.Load(6, &fd.hostFD)
	stateSourceObject.Load(7, &fd.memmapFile)
	stateSourceObject.Load(8, &fd.queue)
	stateSourceObject.Load(9, &fd.clients)
	stateSourceObject.AfterLoad(func() { fd.afterLoad(ctx) })
}

func (mf *frontendFDMemmapFile) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.frontendFDMemmapFile"
}

func (mf *frontendFDMemmapFile) StateFields() []string {
	return []string{
		"NoBufferedIOFallback",
		"fd",
	}
}

func (mf *frontendFDMemmapFile) beforeSave() {}

// +checklocksignore
func (mf *frontendFDMemmapFile) StateSave(stateSinkObject state.Sink) {
	mf.beforeSave()
	stateSinkObject.Save(0, &mf.NoBufferedIOFallback)
	stateSinkObject.Save(1, &mf.fd)
}

func (mf *frontendFDMemmapFile) afterLoad(context.Context) {}

// +checklocksignore
func (mf *frontendFDMemmapFile) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &mf.NoBufferedIOFallback)
	stateSourceObject.Load(1, &mf.fd)
}

func (nvp *nvproxy) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.nvproxy"
}

func (nvp *nvproxy) StateFields() []string {
	return []string{
		"version",
		"frontendFDs",
		"clients",
	}
}

// +checklocksignore
func (nvp *nvproxy) StateSave(stateSinkObject state.Sink) {
	nvp.beforeSave()
	stateSinkObject.Save(0, &nvp.version)
	stateSinkObject.Save(1, &nvp.frontendFDs)
	stateSinkObject.Save(2, &nvp.clients)
}

// +checklocksignore
func (nvp *nvproxy) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &nvp.version)
	stateSourceObject.Load(1, &nvp.frontendFDs)
	stateSourceObject.Load(2, &nvp.clients)
	stateSourceObject.AfterLoad(func() { nvp.afterLoad(ctx) })
}

func (o *object) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.object"
}

func (o *object) StateFields() []string {
	return []string{
		"nvp",
		"client",
		"class",
		"handle",
		"impl",
		"deps",
		"rdeps",
		"objectFreeEntry",
	}
}

func (o *object) beforeSave() {}

// +checklocksignore
func (o *object) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.nvp)
	stateSinkObject.Save(1, &o.client)
	stateSinkObject.Save(2, &o.class)
	stateSinkObject.Save(3, &o.handle)
	stateSinkObject.Save(4, &o.impl)
	stateSinkObject.Save(5, &o.deps)
	stateSinkObject.Save(6, &o.rdeps)
	stateSinkObject.Save(7, &o.objectFreeEntry)
}

func (o *object) afterLoad(context.Context) {}

// +checklocksignore
func (o *object) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.nvp)
	stateSourceObject.Load(1, &o.client)
	stateSourceObject.Load(2, &o.class)
	stateSourceObject.Load(3, &o.handle)
	stateSourceObject.Load(4, &o.impl)
	stateSourceObject.Load(5, &o.deps)
	stateSourceObject.Load(6, &o.rdeps)
	stateSourceObject.Load(7, &o.objectFreeEntry)
}

func (c *capturedRmAllocParams) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.capturedRmAllocParams"
}

func (c *capturedRmAllocParams) StateFields() []string {
	return []string{
		"fd",
		"ioctlParams",
		"rightsRequested",
		"allocParams",
	}
}

func (c *capturedRmAllocParams) beforeSave() {}

// +checklocksignore
func (c *capturedRmAllocParams) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.fd)
	stateSinkObject.Save(1, &c.ioctlParams)
	stateSinkObject.Save(2, &c.rightsRequested)
	stateSinkObject.Save(3, &c.allocParams)
}

func (c *capturedRmAllocParams) afterLoad(context.Context) {}

// +checklocksignore
func (c *capturedRmAllocParams) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.fd)
	stateSourceObject.Load(1, &c.ioctlParams)
	stateSourceObject.Load(2, &c.rightsRequested)
	stateSourceObject.Load(3, &c.allocParams)
}

func (o *rmAllocObject) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.rmAllocObject"
}

func (o *rmAllocObject) StateFields() []string {
	return []string{
		"object",
		"params",
	}
}

func (o *rmAllocObject) beforeSave() {}

// +checklocksignore
func (o *rmAllocObject) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.object)
	stateSinkObject.Save(1, &o.params)
}

func (o *rmAllocObject) afterLoad(context.Context) {}

// +checklocksignore
func (o *rmAllocObject) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.object)
	stateSourceObject.Load(1, &o.params)
}

func (o *rootClient) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.rootClient"
}

func (o *rootClient) StateFields() []string {
	return []string{
		"object",
		"resources",
		"params",
	}
}

func (o *rootClient) beforeSave() {}

// +checklocksignore
func (o *rootClient) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.object)
	stateSinkObject.Save(1, &o.resources)
	stateSinkObject.Save(2, &o.params)
}

func (o *rootClient) afterLoad(context.Context) {}

// +checklocksignore
func (o *rootClient) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.object)
	stateSourceObject.Load(1, &o.resources)
	stateSourceObject.Load(2, &o.params)
}

func (l *objectFreeList) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.objectFreeList"
}

func (l *objectFreeList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *objectFreeList) beforeSave() {}

// +checklocksignore
func (l *objectFreeList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *objectFreeList) afterLoad(context.Context) {}

// +checklocksignore
func (l *objectFreeList) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *objectFreeEntry) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.objectFreeEntry"
}

func (e *objectFreeEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *objectFreeEntry) beforeSave() {}

// +checklocksignore
func (e *objectFreeEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *objectFreeEntry) afterLoad(context.Context) {}

// +checklocksignore
func (e *objectFreeEntry) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (dev *uvmDevice) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.uvmDevice"
}

func (dev *uvmDevice) StateFields() []string {
	return []string{
		"nvp",
	}
}

func (dev *uvmDevice) beforeSave() {}

// +checklocksignore
func (dev *uvmDevice) StateSave(stateSinkObject state.Sink) {
	dev.beforeSave()
	stateSinkObject.Save(0, &dev.nvp)
}

func (dev *uvmDevice) afterLoad(context.Context) {}

// +checklocksignore
func (dev *uvmDevice) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &dev.nvp)
}

func (fd *uvmFD) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.uvmFD"
}

func (fd *uvmFD) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
		"dev",
		"containerName",
		"hostFD",
		"memmapFile",
		"queue",
	}
}

// +checklocksignore
func (fd *uvmFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &fd.NoLockFD)
	stateSinkObject.Save(4, &fd.dev)
	stateSinkObject.Save(5, &fd.containerName)
	stateSinkObject.Save(6, &fd.hostFD)
	stateSinkObject.Save(7, &fd.memmapFile)
	stateSinkObject.Save(8, &fd.queue)
}

// +checklocksignore
func (fd *uvmFD) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &fd.NoLockFD)
	stateSourceObject.Load(4, &fd.dev)
	stateSourceObject.Load(5, &fd.containerName)
	stateSourceObject.Load(6, &fd.hostFD)
	stateSourceObject.Load(7, &fd.memmapFile)
	stateSourceObject.Load(8, &fd.queue)
	stateSourceObject.AfterLoad(func() { fd.afterLoad(ctx) })
}

func (mf *uvmFDMemmapFile) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.uvmFDMemmapFile"
}

func (mf *uvmFDMemmapFile) StateFields() []string {
	return []string{
		"fd",
	}
}

func (mf *uvmFDMemmapFile) beforeSave() {}

// +checklocksignore
func (mf *uvmFDMemmapFile) StateSave(stateSinkObject state.Sink) {
	mf.beforeSave()
	stateSinkObject.Save(0, &mf.fd)
}

func (mf *uvmFDMemmapFile) afterLoad(context.Context) {}

// +checklocksignore
func (mf *uvmFDMemmapFile) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &mf.fd)
}

func (v *DriverVersion) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.DriverVersion"
}

func (v *DriverVersion) StateFields() []string {
	return []string{
		"major",
		"minor",
		"patch",
	}
}

func (v *DriverVersion) beforeSave() {}

// +checklocksignore
func (v *DriverVersion) StateSave(stateSinkObject state.Sink) {
	v.beforeSave()
	stateSinkObject.Save(0, &v.major)
	stateSinkObject.Save(1, &v.minor)
	stateSinkObject.Save(2, &v.patch)
}

func (v *DriverVersion) afterLoad(context.Context) {}

// +checklocksignore
func (v *DriverVersion) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &v.major)
	stateSourceObject.Load(1, &v.minor)
	stateSourceObject.Load(2, &v.patch)
}

func init() {
	state.Register((*frontendDevice)(nil))
	state.Register((*frontendFD)(nil))
	state.Register((*frontendFDMemmapFile)(nil))
	state.Register((*nvproxy)(nil))
	state.Register((*object)(nil))
	state.Register((*capturedRmAllocParams)(nil))
	state.Register((*rmAllocObject)(nil))
	state.Register((*rootClient)(nil))
	state.Register((*objectFreeList)(nil))
	state.Register((*objectFreeEntry)(nil))
	state.Register((*uvmDevice)(nil))
	state.Register((*uvmFD)(nil))
	state.Register((*uvmFDMemmapFile)(nil))
	state.Register((*DriverVersion)(nil))
}
