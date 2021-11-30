package windows

import "errors"

const(
	IPC_CREATE  = 00001000
	IPC_EXCL    = 00002000
	IPC_NOWAIT  = 00004000
	IPC_DIPC    = 00010000
	IPC_OWN     = 00020000
	IPC_PRIVATE = 0
	IPC_RMID    = 0
	IPC_SET     = 1
	IPC_STAT    = 2
	IPC_INFO    = 3
	IPC_OLD     = 0
	IPC_64      = 0x0100
)

type Segment interface {
	Id() int
	Size() int
	Attach() (uintptr, error)
	Detach() error
}
func Create(size int, flags int, mode int) (Segment, error) {
	return OpenSegment(size, IPC_PRIVATE, flags, mode)
}
var (
	ErrUnspportedTarget = errors.New("unspported target")
)

func OpenSegment(size int, key int, flags int, mode int) (Segment, error) {
	return nil, ErrUnspportedTarget
}

func Open(key int) (Segment, error) {
	return OpenSegment(0, key, 0, 0666)
}

