package main

import (
	"strings"
	"path/filepath"
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/log"
	"ethos/kernelTypes"
)
var MAX_ELEMENTS uint32 = 2
var myRpc_increment_counter uint64 = 0
var myRpc_count uint64 = 0
var logger = log.Initialize("test/myRpcService")
//list of files, we set set an initial dimension, but slices in go can be easily resized


// Thread synchronization variables
var sem = make(chan uint64, MAX_ELEMENTS) // Indicates there are files to print
var mutex = make(chan uint64, 1) // Mutex to access the queue
var N uint32 = 0 // Number of elements in the queue
var queue = make(chan string, MAX_ELEMENTS) // The queue


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
	logger.Printf("I recived something")

	<-mutex // sem_wait(mutex)
	if N >= MAX_ELEMENTS {
		// 1 Indicates the buffer is full
		ret = 1
	} else {
		name := param.name
		t := param.t
		buff := param.buff
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
						ret = 3
						mutex <- 1 // sem_post(mutex)
						return &MyRpcFileTransferReply{ret}
					}
				}
				fd, status = altEthos.DirectoryOpen(myDir)
				check(status, "Couldn't open directory " + myDir)
				// TODO check if the filename is already present
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			default:
				ret = 2
				logger.Printf("myRpcService: FileType not found!")
				mutex <- 1 // sem_post(mutex)
				return &MyRpcFileTransferReply{ret}
		}
		// 0 indicates the file is now in print queue
		ret = 0
		queue <- myDir + "/" + name
		N++
		sem <- 1 // sem_post(sem))
	}
	mutex <- 1 // sem_post(mutex)
	return &MyRpcFileTransferReply{ret}
}

func main () {
// Initialize semaphores
	mutex <- 1
	//go printer()

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
			printer()
		}

	}
}
func printer() {

		var dname, path string
		var time syscall.Time64
		time = 2*1000000000 // 6 seconds(time is in nanoseconds)

//		for {
			<- sem //it's like a wait(sem-1), it blocks if sem==0(in our case it never blocks, it simply means that we have finished to process our Box and so we decrement the buffer(sem) by 1)
			<- mutex // sem_wait(mutex)
			N --
			path = <- queue
			dname, _ = parseName(path)
			mutex <- 1 // sem_post(mutex)
			// Read the file and print it 
			fileInfo, status := altEthos.GetFileInformation(dname)
			check(status, "GetFileInformation failed")
			logger.Printf("[PRINTER] Printing file %v\n", path)
			// Beginning of the if-else waterfall
			if (isType(fileInfo.TypeHash, "FileInformation")) {
				var file kernelTypes.FileInformation
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else {
				logger.Printf("[PRINTER] File type not found\n")
			}
//		}
}
func check(status syscall.Status, s string) {
	if status != syscall.StatusOk {
		logger.Fatalf("myRpcService: " + s + " - %v\n",  status)
	}
}

func parseName(str string) (directory string, filename string) {
	filename = filepath.Base(str)
	i := strings.LastIndex(str, filename)
	directory = str[0:i]
	return
}

func compareHash(h1 kernelTypes.HashValue, h2 kernelTypes.HashValue) bool {
	isEqual := true
	for i, e := range h1 {
		if h2[i] != e {
			isEqual = false
		}
	}
	return isEqual
}

func isType(h kernelTypes.HashValue, t string) bool {
	tmpHash, status := altEthos.TypeNameToHash("kernelTypes", t)
	check(status, "TypeNameToHash failed on " + t)
	return compareHash(h, tmpHash)
}

