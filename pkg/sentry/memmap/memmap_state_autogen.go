// automatically generated by stateify.

package memmap

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (fr *FileRange) StateTypeName() string {
	return "pkg/sentry/memmap.FileRange"
}

func (fr *FileRange) StateFields() []string {
	return []string{
		"Start",
		"End",
	}
}

func (fr *FileRange) beforeSave() {}

// +checklocksignore
func (fr *FileRange) StateSave(stateSinkObject state.Sink) {
	fr.beforeSave()
	stateSinkObject.Save(0, &fr.Start)
	stateSinkObject.Save(1, &fr.End)
}

func (fr *FileRange) afterLoad(context.Context) {}

// +checklocksignore
func (fr *FileRange) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fr.Start)
	stateSourceObject.Load(1, &fr.End)
}

func (mr *MappableRange) StateTypeName() string {
	return "pkg/sentry/memmap.MappableRange"
}

func (mr *MappableRange) StateFields() []string {
	return []string{
		"Start",
		"End",
	}
}

func (mr *MappableRange) beforeSave() {}

// +checklocksignore
func (mr *MappableRange) StateSave(stateSinkObject state.Sink) {
	mr.beforeSave()
	stateSinkObject.Save(0, &mr.Start)
	stateSinkObject.Save(1, &mr.End)
}

func (mr *MappableRange) afterLoad(context.Context) {}

// +checklocksignore
func (mr *MappableRange) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &mr.Start)
	stateSourceObject.Load(1, &mr.End)
}

func (r *MappingOfRange) StateTypeName() string {
	return "pkg/sentry/memmap.MappingOfRange"
}

func (r *MappingOfRange) StateFields() []string {
	return []string{
		"MappingSpace",
		"AddrRange",
		"Writable",
	}
}

func (r *MappingOfRange) beforeSave() {}

// +checklocksignore
func (r *MappingOfRange) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.MappingSpace)
	stateSinkObject.Save(1, &r.AddrRange)
	stateSinkObject.Save(2, &r.Writable)
}

func (r *MappingOfRange) afterLoad(context.Context) {}

// +checklocksignore
func (r *MappingOfRange) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.MappingSpace)
	stateSourceObject.Load(1, &r.AddrRange)
	stateSourceObject.Load(2, &r.Writable)
}

func init() {
	state.Register((*FileRange)(nil))
	state.Register((*MappableRange)(nil))
	state.Register((*MappingOfRange)(nil))
}
