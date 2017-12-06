package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/log"
)

var logger = log.Initialize("test/myRpcClient")

func init() {
	SetupMyRpcIncrementReply(incrementReply)
	SetupMyRpcChunkReply(chunkReply)
	SetupMyRpcBoxReply(boxReply)
}

func chunkReply(count uint64) (MyRpcProcedure) {
	logger.Printf("myRpcClient: Service wrote %v bytes\n", count)
	return nil
}

func incrementReply(count uint64) (MyRpcProcedure) {
	logger.Printf("myRpcClient: Recieved Increment Reply: %v\n", count)
	return nil
}

func boxReply(count uint64) (MyRpcProcedure) {
	logger.Printf("myRpcClient: number of box printed: %v\n", count)
	return nil
}

func Random(size uint32) (randomString[]byte, status syscall.Status) {
	eventId, status := syscall.Random(size)
	if status != syscall.StatusOk {
		return
	}
	randomString, _, status = syscall.BlockAndRetire(eventId)
	return
}

func main () {
//	var n uint64
//	var size uint64
//	var filename string
//	var buffer []byte
	var boxBuff Box
//	n = 200
//	filename = "/user/nobody/printFile"

	logger.Printf("myRpcClient: before call\n")

//	fileInfo, status := altEthos.GetFileIndormation(filename)
	//TODO: Check the status like below

	fd, status := altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		logger.Printf("Ipc failed %v\n", status)
		altEthos.Exit(status)
	}
/*
	for i:=0; i<10; i++ {
		call := MyRpcIncrement{n}
		status = altEthos.ClientCall(fd, &call)
		if status != syscall.StatusOk {
			logger.Printf("clientCall failed: %v\n", status)
			altEthos.Exit(status)
		}
	}
	*/
// Generate the box
	rands, _ := Random(4)
	ll := Point{int32(rands[0]), int32(rands[2])}
	ur := Point{int32(rands[1]), int32(rands[3])}
	boxBuff = Box{ll, ur}
	call := MyRpcBox{boxBuff}
	status = altEthos.ClientCall(fd, &call)
	if status != syscall.StatusOk {
		logger.Printf("clientCall failed: %v\n", status)
		altEthos.Exit(status)
	}
	logger.Printf("mtRpcClient: done\n")
}
