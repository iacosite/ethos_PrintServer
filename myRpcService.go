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
}

func test(a []uint8) (MyRpcProcedure) {
	logger.Printf("myRpcService: called increment\n")
	return &MyRpcIncrementReply{myRpc_increment_counter}
}

func increment(quantity uint64) (MyRpcProcedure) {
	logger.Printf("myRpcService: called increment\n")
	myRpc_increment_counter = myRpc_increment_counter + quantity
	return &MyRpcIncrementReply{myRpc_increment_counter}
}

func boxhandle(buff Box) (MyRpcProcedure) {
	myRpc_count += 1
<<<<<<< HEAD
// tried to implement a sleep with Beep, but it doesn't sleep the amount of time we want

=======
>>>>>>> 3a68f25a59fb543bfd94501110c06f931c64e833

	sem <- 1//it's like a signal(sem+1), it blocks if sem==MAX, that is the buffer(sem) is full
	logger.Printf("myRpcService: Box received %v and sent to the printer\n", buff)

<<<<<<< HEAD
//Assuming that we have 6 seconds of execution with sudo ethosRun -t
//It blocks at the third block if the buffer(sem) length is 2 and the printer takes 6 seconds to process 1 Box
//It doesn't block if the buffer(sem) length is 2 and the printer takes 1 secondto process 1 Box


//implement our printer as a concurrent thread with a goroutine
	go func() {

	//implement a sleep with a loop
	//the sleep stands for the time the printer takes to process 1 Box
		time1 := syscall.GetTime()

		logger.Printf("start wait\n")
		var i uint64 = 1
	//it waits around 1 sec with 3*10^9, so for instance 6*(3*10^9) it's around 6 sec
		for i <= 3000000000*6 {
			i = i+1
		}
		time2 := syscall.GetTime()
		logger.Printf("time waited in ms %v\n", (time2-time1)/1000000)

		<-sem //it's like a wait(sem-1), it blocks if sem==0(in our case it never blocks, it simply means that we have finished to process our Box and so we decrement the buffer(sem) by 1)
	}()




//tried to implement a sleep with time.Sleep, but you can't use time.Second or time.Millisecond to put inside like for instance time.Sleep(time.Second * 3)


=======
>>>>>>> 3a68f25a59fb543bfd94501110c06f931c64e833
	return &MyRpcBoxReply{myRpc_count}
}


func fileTransfer (buff []uint8, t uint32, name string) (MyRpcProcedure) {
	var ret uint64
	logger.Printf("I recived something")
	me := syscall.GetUser()
	myDir := "/user/" + me + "/printService/"
	switch t {
		case 1:
			var fd syscall.Fd
			var fi kernelTypes.FileInformation
			logger.Printf("myRpcService: Recived FileInformation!")
			myDir += "FileInformation"
			fd, status := altEthos.DirectoryOpen(myDir)
			if status != syscall.StatusOk {
				status = altEthos.DirectoryCreate(myDir, &fi, "boh")
				if status != syscall.StatusOk {
					logger.Fatalf("myRpcService: Couldn't create directory %v: %v\n", myDir, status)
					ret = 2
					return &MyRpcFileTransferReply{ret}
				}
			}
			// TODO check if the filename is already present
			status = altEthos.WriteVarRaw(fd, myDir + name, buff)
			if status != syscall.StatusOk {
				logger.Printf("myRpcService: Error, couldn't get var from data %v\n", status)
				ret = 3
				return &MyRpcFileTransferReply{ret}
			}
			break
		default:
			ret = 1
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
