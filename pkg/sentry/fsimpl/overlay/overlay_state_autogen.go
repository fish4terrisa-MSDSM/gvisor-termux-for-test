// automatically generated by stateify.

package overlay

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (fd *directoryFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/overlay.directoryFD"
}

func (fd *directoryFD) StateFields() []string {
	return []string{
		"fileDescription",
		"DirectoryFileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"off",
		"dirents",
	}
}

func (fd *directoryFD) beforeSave() {}

// +checklocksignore
func (fd *directoryFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.fileDescription)
	stateSinkObject.Save(1, &fd.DirectoryFileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &fd.off)
	stateSinkObject.Save(4, &fd.dirents)
}

func (fd *directoryFD) afterLoad(context.Context) {}

// +checklocksignore
func (fd *directoryFD) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.fileDescription)
	stateSourceObject.Load(1, &fd.DirectoryFileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &fd.off)
	stateSourceObject.Load(4, &fd.dirents)
}

func (fstype *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/overlay.FilesystemType"
}

func (fstype *FilesystemType) StateFields() []string {
	return []string{}
}

func (fstype *FilesystemType) beforeSave() {}

// +checklocksignore
func (fstype *FilesystemType) StateSave(stateSinkObject state.Sink) {
	fstype.beforeSave()
}

func (fstype *FilesystemType) afterLoad(context.Context) {}

// +checklocksignore
func (fstype *FilesystemType) StateLoad(ctx context.Context, stateSourceObject state.Source) {
}

func (f *FilesystemOptions) StateTypeName() string {
	return "pkg/sentry/fsimpl/overlay.FilesystemOptions"
}

func (f *FilesystemOptions) StateFields() []string {
	return []string{
		"UpperRoot",
		"LowerRoots",
	}
}

func (f *FilesystemOptions) beforeSave() {}

// +checklocksignore
func (f *FilesystemOptions) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.UpperRoot)
	stateSinkObject.Save(1, &f.LowerRoots)
}

func (f *FilesystemOptions) afterLoad(context.Context) {}

// +checklocksignore
func (f *FilesystemOptions) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.UpperRoot)
	stateSourceObject.Load(1, &f.LowerRoots)
}

func (fs *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/overlay.filesystem"
}

func (fs *filesystem) StateFields() []string {
	return []string{
		"vfsfs",
		"opts",
		"creds",
		"dirDevMinor",
		"lowerDevMinors",
		"dirInoCache",
		"lastDirIno",
		"maxFilenameLen",
	}
}

func (fs *filesystem) beforeSave() {}

// +checklocksignore
func (fs *filesystem) StateSave(stateSinkObject state.Sink) {
	fs.beforeSave()
	stateSinkObject.Save(0, &fs.vfsfs)
	stateSinkObject.Save(1, &fs.opts)
	stateSinkObject.Save(2, &fs.creds)
	stateSinkObject.Save(3, &fs.dirDevMinor)
	stateSinkObject.Save(4, &fs.lowerDevMinors)
	stateSinkObject.Save(5, &fs.dirInoCache)
	stateSinkObject.Save(6, &fs.lastDirIno)
	stateSinkObject.Save(7, &fs.maxFilenameLen)
}

func (fs *filesystem) afterLoad(context.Context) {}

// +checklocksignore
func (fs *filesystem) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fs.vfsfs)
	stateSourceObject.Load(1, &fs.opts)
	stateSourceObject.Load(2, &fs.creds)
	stateSourceObject.Load(3, &fs.dirDevMinor)
	stateSourceObject.Load(4, &fs.lowerDevMinors)
	stateSourceObject.Load(5, &fs.dirInoCache)
	stateSourceObject.Load(6, &fs.lastDirIno)
	stateSourceObject.Load(7, &fs.maxFilenameLen)
}

func (l *layerDevNumber) StateTypeName() string {
	return "pkg/sentry/fsimpl/overlay.layerDevNumber"
}

func (l *layerDevNumber) StateFields() []string {
	return []string{
		"major",
		"minor",
	}
}

func (l *layerDevNumber) beforeSave() {}

// +checklocksignore
func (l *layerDevNumber) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.major)
	stateSinkObject.Save(1, &l.minor)
}

func (l *layerDevNumber) afterLoad(context.Context) {}

// +checklocksignore
func (l *layerDevNumber) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.major)
	stateSourceObject.Load(1, &l.minor)
}

func (l *layerDevNoAndIno) StateTypeName() string {
	return "pkg/sentry/fsimpl/overlay.layerDevNoAndIno"
}

func (l *layerDevNoAndIno) StateFields() []string {
	return []string{
		"layerDevNumber",
		"ino",
	}
}

func (l *layerDevNoAndIno) beforeSave() {}

// +checklocksignore
func (l *layerDevNoAndIno) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.layerDevNumber)
	stateSinkObject.Save(1, &l.ino)
}

func (l *layerDevNoAndIno) afterLoad(context.Context) {}

// +checklocksignore
func (l *layerDevNoAndIno) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.layerDevNumber)
	stateSourceObject.Load(1, &l.ino)
}

func (d *dentry) StateTypeName() string {
	return "pkg/sentry/fsimpl/overlay.dentry"
}

func (d *dentry) StateFields() []string {
	return []string{
		"vfsd",
		"refs",
		"fs",
		"mode",
		"uid",
		"gid",
		"copiedUp",
		"parent",
		"name",
		"children",
		"dirents",
		"upperVD",
		"lowerVDs",
		"inlineLowerVDs",
		"devMajor",
		"devMinor",
		"ino",
		"lowerMappings",
		"wrappedMappable",
		"isMappable",
		"locks",
		"watches",
	}
}

func (d *dentry) beforeSave() {}

// +checklocksignore
func (d *dentry) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	var parentValue *dentry
	parentValue = d.saveParent()
	stateSinkObject.SaveValue(7, parentValue)
	stateSinkObject.Save(0, &d.vfsd)
	stateSinkObject.Save(1, &d.refs)
	stateSinkObject.Save(2, &d.fs)
	stateSinkObject.Save(3, &d.mode)
	stateSinkObject.Save(4, &d.uid)
	stateSinkObject.Save(5, &d.gid)
	stateSinkObject.Save(6, &d.copiedUp)
	stateSinkObject.Save(8, &d.name)
	stateSinkObject.Save(9, &d.children)
	stateSinkObject.Save(10, &d.dirents)
	stateSinkObject.Save(11, &d.upperVD)
	stateSinkObject.Save(12, &d.lowerVDs)
	stateSinkObject.Save(13, &d.inlineLowerVDs)
	stateSinkObject.Save(14, &d.devMajor)
	stateSinkObject.Save(15, &d.devMinor)
	stateSinkObject.Save(16, &d.ino)
	stateSinkObject.Save(17, &d.lowerMappings)
	stateSinkObject.Save(18, &d.wrappedMappable)
	stateSinkObject.Save(19, &d.isMappable)
	stateSinkObject.Save(20, &d.locks)
	stateSinkObject.Save(21, &d.watches)
}

// +checklocksignore
func (d *dentry) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.vfsd)
	stateSourceObject.Load(1, &d.refs)
	stateSourceObject.Load(2, &d.fs)
	stateSourceObject.Load(3, &d.mode)
	stateSourceObject.Load(4, &d.uid)
	stateSourceObject.Load(5, &d.gid)
	stateSourceObject.Load(6, &d.copiedUp)
	stateSourceObject.Load(8, &d.name)
	stateSourceObject.Load(9, &d.children)
	stateSourceObject.Load(10, &d.dirents)
	stateSourceObject.Load(11, &d.upperVD)
	stateSourceObject.Load(12, &d.lowerVDs)
	stateSourceObject.Load(13, &d.inlineLowerVDs)
	stateSourceObject.Load(14, &d.devMajor)
	stateSourceObject.Load(15, &d.devMinor)
	stateSourceObject.Load(16, &d.ino)
	stateSourceObject.Load(17, &d.lowerMappings)
	stateSourceObject.Load(18, &d.wrappedMappable)
	stateSourceObject.Load(19, &d.isMappable)
	stateSourceObject.Load(20, &d.locks)
	stateSourceObject.Load(21, &d.watches)
	stateSourceObject.LoadValue(7, new(*dentry), func(y any) { d.loadParent(ctx, y.(*dentry)) })
	stateSourceObject.AfterLoad(func() { d.afterLoad(ctx) })
}

func (fd *fileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/overlay.fileDescription"
}

func (fd *fileDescription) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"LockFD",
	}
}

func (fd *fileDescription) beforeSave() {}

// +checklocksignore
func (fd *fileDescription) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.LockFD)
}

func (fd *fileDescription) afterLoad(context.Context) {}

// +checklocksignore
func (fd *fileDescription) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.LockFD)
}

func (fd *regularFileFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/overlay.regularFileFD"
}

func (fd *regularFileFD) StateFields() []string {
	return []string{
		"fileDescription",
		"copiedUp",
		"cachedFD",
		"cachedFlags",
	}
}

func (fd *regularFileFD) beforeSave() {}

// +checklocksignore
func (fd *regularFileFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.fileDescription)
	stateSinkObject.Save(1, &fd.copiedUp)
	stateSinkObject.Save(2, &fd.cachedFD)
	stateSinkObject.Save(3, &fd.cachedFlags)
}

func (fd *regularFileFD) afterLoad(context.Context) {}

// +checklocksignore
func (fd *regularFileFD) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.fileDescription)
	stateSourceObject.Load(1, &fd.copiedUp)
	stateSourceObject.Load(2, &fd.cachedFD)
	stateSourceObject.Load(3, &fd.cachedFlags)
}

func init() {
	state.Register((*directoryFD)(nil))
	state.Register((*FilesystemType)(nil))
	state.Register((*FilesystemOptions)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*layerDevNumber)(nil))
	state.Register((*layerDevNoAndIno)(nil))
	state.Register((*dentry)(nil))
	state.Register((*fileDescription)(nil))
	state.Register((*regularFileFD)(nil))
}
