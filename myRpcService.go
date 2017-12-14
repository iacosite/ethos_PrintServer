package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/log"
//	"time"
//	"ethos/kernelTypes"
)

var myRpc_increment_counter uint64 = 0
var myRpc_count uint64 = 0
var logger = log.Initialize("test/myRpcService")

//list of files, we set set an initial dimension, but slices in go can be easily resized


//create a buffered channel
var sem = make(chan uint64)

func init() {
	SetupMyRpcIncrement(increment)
	SetupMyRpcBox(boxhandle)
	SetupMyRpcUint32s(uint32handle)
	SetupMyRpcChunk(chunk)
//	SetupMyRpcFileInformaiton(fileInformation)
}

func increment(quantity uint64) (MyRpcProcedure) {
	logger.Printf("myRpcService: called increment\n")
	myRpc_increment_counter = myRpc_increment_counter + quantity
	return &MyRpcIncrementReply{myRpc_increment_counter}
}

func boxhandle(buff Box) (MyRpcProcedure) {
	myRpc_count += 1

	sem <- 1//it's like a signal(sem+1), it blocks if sem==MAX, that is the buffer(sem) is full
	logger.Printf("myRpcService: Box received %v and sent to the printer\n", buff)

	return &MyRpcBoxReply{myRpc_count}
}

func uint32handle(buff uint32) (MyRpcProcedure) {
	var count uint64
	count = 1
	logger.Printf("myRpcService: Recived Uint32 %v\n", buff)
	return &MyRpcUint32sReply{count}
}

func chunk(chunk []byte) (MyRpcProcedure) {
	var count uint64
	count = 1
	logger.Printf("myRpcService: called chunk\n")
	return &MyRpcBoxReply{count}
}

//func fileInformation (buff kernelTypes.FileInformation) (MyRpcProcedure) {
//	var count uint64
//	count = 1
//	logger.Printf("myRpcService: recived fileInformation %v \n", buff)
//	return &MyRpcBoxReply{count}
//}

func main () {

//implement our printer as a concurrent thread with a goroutine
	go func() {
		<-sem //it's like a wait(sem-1), it blocks if sem==0(in our case it never blocks, it simply means that we have finished to process our Box and so we decrement the buffer(sem) by 1)
	//the sleep stands for the time the printer takes to process 1 Box
		var time syscall.Time64

		time = 6*1000000000 // 1 seconds(time is in nanoseconds)

	//sleep until current time + time
		altEthos.Beep(altEthos.GetTime()+time)

	}()



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

		//loop for client sending more RPCs
		for {
			t:= MyRpc{}
			altEthos.Handle(fd, &t)
		}

	}
}
