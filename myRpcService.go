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
var myRpc_count uint64 = 0
var logger = log.Initialize("test/myRpcService")
//list of files, we set set an initial dimension, but slices in go can be easily resized


// Thread synchronization variables
var sem = make(chan uint64, MAX_ELEMENTS) // Indicates there are files to print
var mutex = make(chan uint64, 1) // Mutex to access the queue
var N uint32 = 0 // Number of elements in the queue
var queue = make(chan string, MAX_ELEMENTS) // The queue


func init() {
	SetupMyRpcFileTransfer(fileTransfer)
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
			case 2:
				var fd syscall.Fd
				var fi kernelTypes.Account
				logger.Printf("myRpcService: Recived Account!")
				myDir += "/Account"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 3:
				var fd syscall.Fd
				var fi kernelTypes.Annotation
				logger.Printf("myRpcService: Recived Annotation!")
				myDir += "/Annotation"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 4:
				var fd syscall.Fd
				var fi kernelTypes.Any
				logger.Printf("myRpcService: Recived Any!")
				myDir += "/Any"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 5:
				var fd syscall.Fd
				var fi kernelTypes.Authenticator
				logger.Printf("myRpcService: Recived Authenticator!")
				myDir += "/Authenticator"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 6:
				var fd syscall.Fd
				var fi kernelTypes.Base
				logger.Printf("myRpcService: Recived Base!")
				myDir += "/Base"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 8:
				var fd syscall.Fd
				var fi kernelTypes.CertificateHeader
				logger.Printf("myRpcService: Recived CertificateHeader!")
				myDir += "/CertificateHeader"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 9:
				var fd syscall.Fd
				var fi kernelTypes.DirectoryAttributes
				logger.Printf("myRpcService: Recived DirectoryAttributes!")
				myDir += "/DirectoryAttributes"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 10:
				var fd syscall.Fd
				var fi kernelTypes.DirectoryCertificate
				logger.Printf("myRpcService: Recived DirectoryCertificate!")
				myDir += "/DirectoryCertificate"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 11:
				var fd syscall.Fd
				var fi kernelTypes.DirectoryServiceRecord
				logger.Printf("myRpcService: Recived DirectoryServiceRecord!")
				myDir += "/DirectoryServiceRecord"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 12:
				var fd syscall.Fd
				var fi kernelTypes.DualEncryptionKeyPair
				logger.Printf("myRpcService: Recived DualEncryptionKeyPair!")
				myDir += "/DualEncryptionKeyPair"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 13:
				var fd syscall.Fd
				var fi kernelTypes.EncryptionKeyPair
				logger.Printf("myRpcService: Recived EncryptionKeyPair!")
				myDir += "/EncryptionKeyPair"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 14:
				var fd syscall.Fd
				var fi kernelTypes.EncryptionPrivateKey
				logger.Printf("myRpcService: Recived EncryptionPrivateKey!")
				myDir += "/EncryptionPrivateKey"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 15:
				var fd syscall.Fd
				var fi kernelTypes.EncryptionPublicKey
				logger.Printf("myRpcService: Recived EncryptionPublicKey!")
				myDir += "/EncryptionPublicKey"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 16:
				var fd syscall.Fd
				var fi kernelTypes.EphemeralHostCertificate
				logger.Printf("myRpcService: Recived EphemeralHostCertificate!")
				myDir += "/EphemeralHostCertificate"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 17:
				var fd syscall.Fd
				var fi kernelTypes.EphemeralKeyAndExpiration
				logger.Printf("myRpcService: Recived EphemeralKeyAndExpiration!")
				myDir += "/EphemeralKeyAndExpiration"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 18:
				var fd syscall.Fd
				var fi kernelTypes.EthernetMac
				logger.Printf("myRpcService: Recived EthernetMac!")
				myDir += "/EthernetMac"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 19:
				var fd syscall.Fd
				var fi kernelTypes.EthernetMacString
				logger.Printf("myRpcService: Recived EthernetMacString!")
				myDir += "/EthernetMacString"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 20:
				var fd syscall.Fd
				var fi kernelTypes.ExitStatus
				logger.Printf("myRpcService: Recived ExitStatus!")
				myDir += "/ExitStatus"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 22:
				var fd syscall.Fd
				var fi kernelTypes.FilePermission
				logger.Printf("myRpcService: Recived FilePermission!")
				myDir += "/FilePermission"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 25:
				var fd syscall.Fd
				var fi kernelTypes.GdbProxyCall
				logger.Printf("myRpcService: Recived GdbProxyCall!")
				myDir += "/GdbProxyCall"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 26:
				var fd syscall.Fd
				var fi kernelTypes.GdbProxyReply
				logger.Printf("myRpcService: Recived GdbProxyReply!")
				myDir += "/GdbProxyReply"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 27:
				var fd syscall.Fd
				var fi kernelTypes.GRegisters
				logger.Printf("myRpcService: Recived GRegisters!")
				myDir += "/GRegisters"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 28:
				var fd syscall.Fd
				var fi kernelTypes.GroupType
				logger.Printf("myRpcService: Recived GroupType!")
				myDir += "/GroupType"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 29:
				var fd syscall.Fd
				var fi kernelTypes.HashValue
				logger.Printf("myRpcService: Recived HashValue!")
				myDir += "/HashValue"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 30:
				var fd syscall.Fd
				var fi kernelTypes.HostCertificate
				logger.Printf("myRpcService: Recived HostCertificate!")
				myDir += "/HostCertificate"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 31:
				var fd syscall.Fd
				var fi kernelTypes.HostGroupCertificate
				logger.Printf("myRpcService: Recived HostGroupCertificate!")
				myDir += "/HostGroupCertificate"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 32:
				var fd syscall.Fd
				var fi kernelTypes.HostRecord
				logger.Printf("myRpcService: Recived HostRecord!")
				myDir += "/HostRecord"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 37:
				var fd syscall.Fd
				var fi kernelTypes.IpAddress
				logger.Printf("myRpcService: Recived IpAddress!")
				myDir += "/IpAddress"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 38:
				var fd syscall.Fd
				var fi kernelTypes.IpAddressString
				logger.Printf("myRpcService: Recived IpAddressString!")
				myDir += "/IpAddressString"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 39:
				var fd syscall.Fd
				var fi kernelTypes.LabelType
				logger.Printf("myRpcService: Recived LabelType!")
				myDir += "/LabelType"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 41:
				var fd syscall.Fd
				var fi kernelTypes.Nanoseconds
				logger.Printf("myRpcService: Recived Nanoseconds!")
				myDir += "/Nanoseconds"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 42:
				var fd syscall.Fd
				var fi kernelTypes.Password
				logger.Printf("myRpcService: Recived Password!")
				myDir += "/Password"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 43:
				var fd syscall.Fd
				var fi kernelTypes.Permission
				logger.Printf("myRpcService: Recived Permission!")
				myDir += "/Permission"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 44:
				var fd syscall.Fd
				var fi kernelTypes.ProcessInfo
				logger.Printf("myRpcService: Recived ProcessInfo!")
				myDir += "/ProcessInfo"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 46:
				var fd syscall.Fd
				var fi kernelTypes.ProgramName
				logger.Printf("myRpcService: Recived ProgramName!")
				myDir += "/ProgramName"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 47:
				var fd syscall.Fd
				var fi kernelTypes.Puzzle
				logger.Printf("myRpcService: Recived Puzzle!")
				myDir += "/Puzzle"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 48:
				var fd syscall.Fd
				var fi kernelTypes.PuzzleSecret
				logger.Printf("myRpcService: Recived PuzzleSecret!")
				myDir += "/PuzzleSecret"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 49:
				var fd syscall.Fd
				var fi kernelTypes.PuzzleSolution
				logger.Printf("myRpcService: Recived PuzzleSolution!")
				myDir += "/PuzzleSolution"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 57:
				var fd syscall.Fd
				var fi kernelTypes.Seconds
				logger.Printf("myRpcService: Recived Seconds!")
				myDir += "/Seconds"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 58:
				var fd syscall.Fd
				var fi kernelTypes.Signature
				logger.Printf("myRpcService: Recived Signature!")
				myDir += "/Signature"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 59:
				var fd syscall.Fd
				var fi kernelTypes.SignatureKeyPair
				logger.Printf("myRpcService: Recived SignatureKeyPair!")
				myDir += "/SignatureKeyPair"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 60:
				var fd syscall.Fd
				var fi kernelTypes.SignaturePrivateKey
				logger.Printf("myRpcService: Recived SignaturePrivateKey!")
				myDir += "/SignaturePrivateKey"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 61:
				var fd syscall.Fd
				var fi kernelTypes.SignaturePublicKey
				logger.Printf("myRpcService: Recived SignaturePublicKey!")
				myDir += "/SignaturePublicKey"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 64:
				var fd syscall.Fd
				var fi kernelTypes.Time
				logger.Printf("myRpcService: Recived Time!")
				myDir += "/Time"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 65:
				var fd syscall.Fd
				var fi kernelTypes.Time64
				logger.Printf("myRpcService: Recived Time64!")
				myDir += "/Time64"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 66:
				var fd syscall.Fd
				var fi kernelTypes.TypeNode
				logger.Printf("myRpcService: Recived TypeNode!")
				myDir += "/TypeNode"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 67:
				var fd syscall.Fd
				var fi kernelTypes.Types
				logger.Printf("myRpcService: Recived Types!")
				myDir += "/Types"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 68:
				var fd syscall.Fd
				var fi kernelTypes.UdpPort
				logger.Printf("myRpcService: Recived UdpPort!")
				myDir += "/UdpPort"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 73:
				var fd syscall.Fd
				var fi kernelTypes.UserCertificate
				logger.Printf("myRpcService: Recived UserCertificate!")
				myDir += "/UserCertificate"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 74:
				var fd syscall.Fd
				var fi kernelTypes.UserGroupCertificate
				logger.Printf("myRpcService: Recived UserGroupCertificate!")
				myDir += "/UserGroupCertificate"
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
				status = altEthos.WriteVarRaw(fd, name, buff)
				if status != syscall.StatusOk {
					logger.Printf("myRpcService: Error, couldn't save recived data %v\n", status)
					ret = 4
					mutex <- 1 // sem_post(mutex)
					return &MyRpcFileTransferReply{ret}
				}
				break
			case 75:
				var fd syscall.Fd
				var fi kernelTypes.UserRecord
				logger.Printf("myRpcService: Recived UserRecord!")
				myDir += "/UserRecord"
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
			if (isType(fileInfo.TypeHash, "Account")) {
				var file kernelTypes.Account
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "Annotation")) {
				var file kernelTypes.Annotation
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "Any")) {
				var file Any
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "Authenticator")) {
				var file kernelTypes.Authenticator
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "Base")) {
				var file kernelTypes.Base
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "CertificateHeader")) {
				var file kernelTypes.CertificateHeader
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "DirectoryAttributes")) {
				var file kernelTypes.DirectoryAttributes
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "DirectoryCertificate")) {
				var file kernelTypes.DirectoryCertificate
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "DirectoryServiceRecord")) {
				var file kernelTypes.DirectoryServiceRecord
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "DualEncryptionKeyPair")) {
				var file kernelTypes.DualEncryptionKeyPair
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "EncryptionKeyPair")) {
				var file kernelTypes.EncryptionKeyPair
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "EncryptionPrivateKey")) {
				var file kernelTypes.EncryptionPrivateKey
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "EncryptionPublicKey")) {
				var file kernelTypes.EncryptionPublicKey
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "EphemeralHostCertificate")) {
				var file kernelTypes.EphemeralHostCertificate
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "EphemeralKeyAndExpiration")) {
				var file kernelTypes.EphemeralKeyAndExpiration
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "EthernetMac")) {
				var file kernelTypes.EthernetMac
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "EthernetMacString")) {
				var file kernelTypes.EthernetMacString
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "ExitStatus")) {
				var file kernelTypes.ExitStatus
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "FileInformation")) {
				var file kernelTypes.FileInformation
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "FilePermission")) {
				var file kernelTypes.FilePermission
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "GdbProxyCall")) {
				var file kernelTypes.GdbProxyCall
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "GdbProxyReply")) {
				var file kernelTypes.GdbProxyReply
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "GRegisters")) {
				var file kernelTypes.GRegisters
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "GroupType")) {
				var file kernelTypes.GroupType
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "HashValue")) {
				var file kernelTypes.HashValue
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "HostCertificate")) {
				var file kernelTypes.HostCertificate
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "HostGroupCertificate")) {
				var file kernelTypes.HostGroupCertificate
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "HostRecord")) {
				var file kernelTypes.HostRecord
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "IpAddress")) {
				var file kernelTypes.IpAddress
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "IpAddressString")) {
				var file kernelTypes.IpAddressString
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "LabelType")) {
				var file kernelTypes.LabelType
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "Nanoseconds")) {
				var file kernelTypes.Nanoseconds
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "Password")) {
				var file kernelTypes.Password
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "Permission")) {
				var file kernelTypes.Permission
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "ProcessInfo")) {
				var file kernelTypes.ProcessInfo
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "ProgramName")) {
				var file kernelTypes.ProgramName
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "Puzzle")) {
				var file kernelTypes.Puzzle
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "PuzzleSecret")) {
				var file kernelTypes.PuzzleSecret
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "PuzzleSolution")) {
				var file kernelTypes.PuzzleSolution
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "Seconds")) {
				var file kernelTypes.Seconds
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "Signature")) {
				var file kernelTypes.Signature
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "SignatureKeyPair")) {
				var file kernelTypes.SignatureKeyPair
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "SignaturePrivateKey")) {
				var file kernelTypes.SignaturePrivateKey
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "SignaturePublicKey")) {
				var file kernelTypes.SignaturePublicKey
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "Time")) {
				var file kernelTypes.Time
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "Time64")) {
				var file kernelTypes.Time64
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "TypeNode")) {
				var file kernelTypes.TypeNode
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "Types")) {
				var file kernelTypes.Types
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "UdpPort")) {
				var file kernelTypes.UdpPort
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "UserCertificate")) {
				var file kernelTypes.UserCertificate
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "UserGroupCertificate")) {
				var file kernelTypes.UserGroupCertificate
				status = altEthos.Read(path, &file)
				check(status, "Could not read file")
				// Simulate the print action
				logger.Printf("[PRINTER] %v\n", file)
				altEthos.Beep(altEthos.GetTime()+time)
				logger.Printf("[PRINTER] done\n")
			} else
			if (isType(fileInfo.TypeHash, "UserRecord")) {
				var file kernelTypes.UserRecord
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

