package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/log"
	"ethos/kernelTypes"
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
	SetupMyRpcFileTransfer(fileTransfer)
	SetupMyRpcTestVars(test)
}

func test(tmp Par) (MyRpcProcedure) {
	logger.Printf("myRpcService: called test %v\n", tmp)
	return &MyRpcIncrementReply{myRpc_increment_counter}
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


func fileTransfer (param Param) (MyRpcProcedure) {
	var ret uint64
	name := param.name
	t := param.t
	buff := param.buff
	logger.Printf("I recived something")
	me := syscall.GetUser()
	myDir := "/user/" + me 
	switch t {
		case 21:
			var fd syscall.Fd
			var fi kernelTypes.FileInformation
			logger.Printf("myRpcService: Recived FileInformation!")
			myDir += "/FileInformation"
			fd, status := altEthos.DirectoryOpen(myDir)
			if status != syscall.StatusOk {
				status = altEthos.DirectoryCreate(myDir, &fi, "boh")
				if status != syscall.StatusOk {
					logger.Fatalf("myRpcService: Couldn't create directory %v: %v\n", myDir, status)
					ret = 2
					return &MyRpcFileTransferReply{ret}
				}
			}
			fd, status = altEthos.DirectoryOpen(myDir)
			check(status, "Couldn't open directory " + myDir)
			// TODO check if the filename is already present
			status = altEthos.WriteVarRaw(fd, name, buff)
			if status != syscall.StatusOk {
				logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
				ret = 3
				return &MyRpcFileTransferReply{ret}
			}
			break
		default:
			ret = 1
			logger.Printf("myRpcService: FileType not found!")
			return &MyRpcFileTransferReply{ret}
	}
	ret = 0
	return &MyRpcFileTransferReply{uint64(0)}
}

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
	check(status, "Advertising service failes")

	for {
		_, fd, status := altEthos.Import(listeningFd)
		check(status, "Error calling import")
		logger.Printf("myRpcService: new connection accepted\n")

		//loop for client sending more RPCs
		for {
			t:= MyRpc{}
			altEthos.Handle(fd, &t)
		}

	}
}

func check(status syscall.Status, s string) {
	if status != syscall.StatusOk {
		logger.Fatalf("myRpcService: " + s + " - %v\n",  status)
	}
}
