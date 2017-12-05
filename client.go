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
	logger.Printf("myRpcClient: before call\n")

	fd, status := altEthos.lpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		logger.Printf("lpc failed %v\n", status)
		altEthos(status)
	}

	logger.Printf("mtRpcClient: done\n")
}
