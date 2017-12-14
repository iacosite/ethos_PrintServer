package main

import (
	"reflect"
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/log"
	"ethos/kernelTypes"
)

var logger = log.Initialize("test/myRpcClient")

func init() {
	SetupMyRpcIncrementReply(incrementReply)
	SetupMyRpcBoxReply(boxReply)
	SetupMyRpcFileTransferReply(fileTransferReply)
}


func incrementReply(count uint64) (MyRpcProcedure) {
	logger.Printf("myRpcClient: Recieved Increment Reply: %v\n", count)
	return nil
}

func boxReply(count uint64) (MyRpcProcedure) {
	logger.Printf("myRpcClient: number of box printed: %v\n", count)
	return nil
}


func fileTransferReply(count uint64) (MyRpcProcedure) {
	logger.Printf("myRpcClient: Recieved %v from server after sending the file\n", count)
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
	createfile()
	// Name of the file we will need to print
	// Get the directory in order to understand the type of file
	filename := "/user/test/"
	name := "test"
	var t uint32
	var bytes []uint8
	var fd syscall.Fd
	logger.Printf("myRpcClient: before call\n")
	// Get file information
	fileInfo, status := altEthos.GetFileInformation(filename)
	check(status, "GetFileInformation failed")
	// Beginning of the if-else waterfall
	if (isType(fileInfo.TypeHash, "FileInformation")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(filename)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, "printFile")
		check(status, "ReadVarWar failed")
		// 1 is the identifier relative to FileInformation
		t = 1
	} else {
		// Another type
		logger.Printf("wrong type mate")
	}
	logger.Printf("Bytes[%v]: %v\n", reflect.TypeOf(bytes), bytes)
	logger.Printf("Type[%v]: %v\n", reflect.TypeOf(t), t)
	logger.Printf("Name[%v]: %v\n", reflect.TypeOf(name), name)
	call := MyRpcFileTransfer{bytes, t, name}
	status = altEthos.ClientCall(fd, &call)
	check(status, "Failed to send data")

	logger.Printf("myRpcClient: done\n")
}

// Create a fake file. Needed in order to be sent to the server. 
func createfile() {
	// Get info about some random file
	filename := "/types/spec/kernelTypes/"
	fileInfo, status := altEthos.GetFileInformation(filename)
	check(status, "GetFileInformation failed")
	// Store the file in a fixed directory
	path := "/user/test/printFile"
	status = altEthos.DirectoryCreate("/user/test", &fileInfo, "boh")
	check(status, "Couldn't create directory " + path)
	status = altEthos.Write(path, &fileInfo)
	check(status, "Couldn't write file")
}
func sendBoxes() {
	var boxBuff Box
	fd, status := altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		logger.Printf("Ipc failed %v\n", status)
		altEthos.Exit(status)
	}
//generate and send 3 random boxes
	for i := 0; i<=3; i++{
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
	}

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
	check(status, "TypeNameToHash failed")
	return compareHash(h, tmpHash)
}

func check(status syscall.Status, s string) {
	if status != syscall.StatusOk {
		logger.Fatalf("myRpcClient: " + s + " - %v\n",  status)
	}
}
