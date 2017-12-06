package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/log"
)

var logger = log.Initialize("test/myRpcClient")

func init() {
	SetupMyRpcIncrementReply(incrementReply)

}

func incrementReply(count uint64) (MyRpcProcedure) {
	logger.Printf("myRpcClient: Recieved Increment Reply: %v\n", count)
	return nil
}

func main () {
	var n uint64
	n = 200
	logger.Printf("myRpcClient: before call\n")

	fd, status := altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		logger.Printf("Ipc failed %v\n", status)
		altEthos.Exit(status)
	}

	call := MyRpcIncrement{n}
	status = altEthos.ClientCall(fd, &call)
	if status != syscall.StatusOk {
		logger.Printf("clientCall failed: %v\n", status)
		altEthos.Exit(status)
	}

	logger.Printf("mtRpcClient: done\n")
}
