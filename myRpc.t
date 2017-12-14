MyRpc interface {
	FileTransfer(buffer Param) (count uint64)
}

Param struct {
	buff []uint8
	t uint32
	name string
}
