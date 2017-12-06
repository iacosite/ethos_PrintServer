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
	SetupMyRpcBox(boxhandle)
	SetupMyRpcChunk(chunk)
}

func increment(quantity uint64) (MyRpcProcedure) {
	logger.Printf("myRpcService: called increment\n")
	myRpc_increment_counter = myRpc_increment_counter + quantity
	return &MyRpcIncrementReply{myRpc_increment_counter}
}

func boxhandle(buff Box) (MyRpcProcedure) {
	var count uint64
	count = 1
	logger.Printf("myRpcService: Recived box %v\n", buff)
	return &MyRpcBoxReply{count}
}
func chunk(chunk []byte) (MyRpcProcedure) {
	var count uint64
	count = 1
	logger.Printf("myRpcService: called chunk\n")
	return &MyRpcBoxReply{count}
}
func main () {
	listeningFd, status := altEthos.Advertise("myRpc")
	if status != syscall.StatusOk {
		logger.Printf("Advertising service failes %s\n", status)
		altEthos.Exit(status)
	}

	for {
		_, fd, status := altEthos.Import(listeningFd)
		if status != syscall.StatusOk {
			logger.Printf("Error calling import %v\n", status)
		}

		logger.Printf("myRpcService: new connection accepted\n")

		t:= MyRpc{}
	//	for i:=0; i<10; i++{
		altEthos.Handle(fd, &t)
//		}
	}
}
