package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/log"
	"ethos/kernelTypes"
)

var logger = log.Initialize("test/myRpcClient")

func init() {
	SetupMyRpcIncrementReply(incrementReply)
	SetupMyRpcChunkReply(chunkReply)
	SetupMyRpcBoxReply(boxReply)
	SetupMyRpcUint32sReply(uint32Reply)
//	SetupMyRpcFileInformationReply(fileInformationReply)
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

func uint32Reply(count uint64) (MyRpcProcedure) {
	logger.Printf("myRpcClient: number of Uint32 printed: %v\n", count)
	return nil
}

//func fileInformatoinReply(count uint64) (MyRpcProcedure) {
//	logger.Printf("myRpcClient: number of FileInformation printed: %v\n", count)
//	return nil
//}

func Random(size uint32) (randomString[]byte, status syscall.Status) {
	eventId, status := syscall.Random(size)
	if status != syscall.StatusOk {
		return
	}
	randomString, _, status = syscall.BlockAndRetire(eventId)
	return
}

func main () {
	var boxBuff Box
	filename := "/user/test/"
	createfile()

	logger.Printf("myRpcClient: before call\n")

	fileInfo, status := altEthos.GetFileInformation(filename)
	if status != syscall.StatusOk {
		logger.Fatalf("myRpcClient: GetFileInformation failed %v\n", status)
		altEthos.Exit(status)
	} else {
		logger.Printf("myRpcClient: File info: %v", fileInfo)
		logger.Printf("Label: %v", fileInfo.Label)
	}
	typeHash := fileInfo.TypeHash



	tmpHash, status := altEthos.TypeNameToHash("kernelTypes", "FileInformation")
	if status != syscall.StatusOk {
		logger.Printf("myRpcClient: TypeNameToHash failed %v\n", status)
		altEthos.Exit(status)
	}
	if (compareHash(tmpHash, typeHash)) {
		logger.Printf("type recognized")
//		call := MyRpcFileInformation{buff}
//		status = altEthos.ClientCall(fd, &call)
//		if status != syscall.StatusOk {
//			logger.Printf("clientCall failed: %v\n", status)
//			altEthos.Exit(status)
//		}
	} else {
		logger.Printf("wrong type mate")
	}


// da qui c'e solo roba inutile, l'ho tenuta per avere un programma che funziona un minimo e trasmette un file
//Generate the box
	fd, status := altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		logger.Printf("Ipc failed %v\n", status)
		altEthos.Exit(status)
	}
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
	logger.Printf("myRpcClient: done\n")
}

// Create a fake file. Needed in order to be sent to the server. 
func createfile() {
	// Get info about some random file (just use directory, since they are typed. GetFileInformation fails if address is a file)
	filename := "/types/spec/kernelTypes/"
	fileInfo, status := altEthos.GetFileInformation(filename)
	if status != syscall.StatusOk {
		logger.Fatalf("myRpcClient: GetFileInformation failed %v\n", status)
		altEthos.Exit(status)
	}
	// Store the file in a fixed directory
	path := "/user/test/printFile"
	status = altEthos.DirectoryCreate("/user/test", &fileInfo, "boh")
	if status != syscall.StatusOk {
		logger.Fatalf("myRpcClient: Couldn't create directory %v: %v\n", path, status)
	}
	status = altEthos.Write(path, &fileInfo)
	if status != syscall.StatusOk {
		logger.Fatalf("myRpcClient: Couldn't write file %v\n", status)
	} else {
		logger.Printf("myRpcClient: File written in %v:\n%v\n", path, fileInfo)
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
