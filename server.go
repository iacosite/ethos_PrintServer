package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/log"
)

var myRpc_increment_counter uint64 = 0
var logger = log.Initialize("test/myRpcService")

func init() {
	SetupMyRpcIncrement(increment)
}

func increment() (MyRpcProcedure) {
	logger.Printf("myRpcService: called increment\n")
	myRpc_increment_counter++
	return &MyrpcIncrementReply(myRpc_increment_counter)
}

func main () {
	listeningFd, status := altEthos.Advertise("myRpc")
	if status != syscall.StatusOk {
		logger.Printf("Advertising service failes %s\n", status)
		altEthos.Exit(status)
	}

	for {
		-, fd, status := altEthos.Import(listeningfd)
		if status != syscall.StatusOk {
			logger.Printf("Error calling import %v\n", status)
		}

		logger.Printf("myRpcService: new connection accepted\n")

		t:= MyRpc{}
		altEthos.Handle(fd, &t)
	}
}
