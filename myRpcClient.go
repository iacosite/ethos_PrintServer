package main

import (
	"strings"
	"path/filepath"
	"reflect"
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/log"
	"ethos/kernelTypes"
)

var logger = log.Initialize("test/myRpcClient")

func init() {
	SetupMyRpcFileTransferReply(fileTransferReply)
}

func fileTransferReply(count uint64) (MyRpcProcedure) {
	logger.Printf("myRpcClient: Recieved %v from server after sending the file\n", count)
	switch count {
	case 1:
		logger.Printf("Apparently the print queue is full...\n")
		break
	case 2:
		logger.Printf("File type not found by server\n")
		break
	case 3:
		logger.Printf("Server could not save file\n")
		break
	case 4:
		logger.Printf("Error in data transfer\n")
		break
	default:
		logger.Printf("Return value not recognized \n")
		break
	}
	return nil
}

func main () {
	// filename will be the only parameter recived by the program
	filename := "/user/test/printFile"

	var t uint32
	var bytes []uint8
	var fd syscall.Fd
	sock, status := altEthos.IpcRepeat("myRpc", "", nil)
	check(status, "Ipc failed")
	createfile()
	// Name of the file we will need to print
	// Get the directory in order to understand the type of file
	dname, fname := parseName(filename) 
	logger.Printf("myRpcClient: before call\n")
	// Get file information
	fileInfo, status := altEthos.GetFileInformation(dname)
	check(status, "GetFileInformation failed")
	// Beginning of the if-else waterfall
	if (isType(fileInfo.TypeHash, "Account")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 2 is the identifier relative to Account
		t = 2
	} else
	if (isType(fileInfo.TypeHash, "Annotation")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 3 is the identifier relative to Annotation
		t = 3
	} else
	if (isType(fileInfo.TypeHash, "Any")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 4 is the identifier relative to Any
		t = 4
	} else
	if (isType(fileInfo.TypeHash, "Authenticator")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 5 is the identifier relative to Authenticator
		t = 5
	} else
	if (isType(fileInfo.TypeHash, "Base")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 6 is the identifier relative to Base
		t = 6
	} else
	if (isType(fileInfo.TypeHash, "bool")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 7 is the identifier relative to bool
		t = 7
	} else
	if (isType(fileInfo.TypeHash, "CertificateHeader")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 8 is the identifier relative to CertificateHeader
		t = 8
	} else
	if (isType(fileInfo.TypeHash, "DirectoryAttributes")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 9 is the identifier relative to DirectoryAttributes
		t = 9
	} else
	if (isType(fileInfo.TypeHash, "DirectoryCertificate")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 10 is the identifier relative to DirectoryCertificate
		t = 10
	} else
	if (isType(fileInfo.TypeHash, "DirectoryServiceRecord")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 11 is the identifier relative to DirectoryServiceRecord
		t = 11
	} else
	if (isType(fileInfo.TypeHash, "DualEncryptionKeyPair")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 12 is the identifier relative to DualEncryptionKeyPair
		t = 12
	} else
	if (isType(fileInfo.TypeHash, "EncryptionKeyPair")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 13 is the identifier relative to EncryptionKeyPair
		t = 13
	} else
	if (isType(fileInfo.TypeHash, "EncryptionPrivateKey")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 14 is the identifier relative to EncryptionPrivateKey
		t = 14
	} else
	if (isType(fileInfo.TypeHash, "EncryptionPublicKey")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 15 is the identifier relative to EncryptionPublicKey
		t = 15
	} else
	if (isType(fileInfo.TypeHash, "EphemeralHostCertificate")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 16 is the identifier relative to EphemeralHostCertificate
		t = 16
	} else
	if (isType(fileInfo.TypeHash, "EphemeralKeyAndExpiration")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 17 is the identifier relative to EphemeralKeyAndExpiration
		t = 17
	} else
	if (isType(fileInfo.TypeHash, "EthernetMac")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 18 is the identifier relative to EthernetMac
		t = 18
	} else
	if (isType(fileInfo.TypeHash, "EthernetMacString")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 19 is the identifier relative to EthernetMacString
		t = 19
	} else
	if (isType(fileInfo.TypeHash, "ExitStatus")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 20 is the identifier relative to ExitStatus
		t = 20
	} else
	if (isType(fileInfo.TypeHash, "FileInformation")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 21 is the identifier relative to FileInformation
		t = 21
	} else
	if (isType(fileInfo.TypeHash, "FilePermission")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 22 is the identifier relative to FilePermission
		t = 22
	} else
	if (isType(fileInfo.TypeHash, "float32")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 23 is the identifier relative to float32
		t = 23
	} else
	if (isType(fileInfo.TypeHash, "float64")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 24 is the identifier relative to float64
		t = 24
	} else
	if (isType(fileInfo.TypeHash, "GdbProxyCall")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 25 is the identifier relative to GdbProxyCall
		t = 25
	} else
	if (isType(fileInfo.TypeHash, "GdbProxyReply")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 26 is the identifier relative to GdbProxyReply
		t = 26
	} else
	if (isType(fileInfo.TypeHash, "GRegisters")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 27 is the identifier relative to GRegisters
		t = 27
	} else
	if (isType(fileInfo.TypeHash, "GroupType")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 28 is the identifier relative to GroupType
		t = 28
	} else
	if (isType(fileInfo.TypeHash, "HashValue")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 29 is the identifier relative to HashValue
		t = 29
	} else
	if (isType(fileInfo.TypeHash, "HostCertificate")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 30 is the identifier relative to HostCertificate
		t = 30
	} else
	if (isType(fileInfo.TypeHash, "HostGroupCertificate")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 31 is the identifier relative to HostGroupCertificate
		t = 31
	} else
	if (isType(fileInfo.TypeHash, "HostRecord")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 32 is the identifier relative to HostRecord
		t = 32
	} else
	if (isType(fileInfo.TypeHash, "int16")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 33 is the identifier relative to int16
		t = 33
	} else
	if (isType(fileInfo.TypeHash, "int32")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 34 is the identifier relative to int32
		t = 34
	} else
	if (isType(fileInfo.TypeHash, "int64")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 35 is the identifier relative to int64
		t = 35
	} else
	if (isType(fileInfo.TypeHash, "int8")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 36 is the identifier relative to int8
		t = 36
	} else
	if (isType(fileInfo.TypeHash, "IpAddress")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 37 is the identifier relative to IpAddress
		t = 37
	} else
	if (isType(fileInfo.TypeHash, "IpAddressString")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 38 is the identifier relative to IpAddressString
		t = 38
	} else
	if (isType(fileInfo.TypeHash, "LabelType")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 39 is the identifier relative to LabelType
		t = 39
	} else
	if (isType(fileInfo.TypeHash, "MinimaltdRpc")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 40 is the identifier relative to MinimaltdRpc
		t = 40
	} else
	if (isType(fileInfo.TypeHash, "Nanoseconds")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 41 is the identifier relative to Nanoseconds
		t = 41
	} else
	if (isType(fileInfo.TypeHash, "Password")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 42 is the identifier relative to Password
		t = 42
	} else
	if (isType(fileInfo.TypeHash, "Permission")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 43 is the identifier relative to Permission
		t = 43
	} else
	if (isType(fileInfo.TypeHash, "ProcessInfo")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 44 is the identifier relative to ProcessInfo
		t = 44
	} else
	if (isType(fileInfo.TypeHash, "ProcStatus")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 45 is the identifier relative to ProcStatus
		t = 45
	} else
	if (isType(fileInfo.TypeHash, "ProgramName")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 46 is the identifier relative to ProgramName
		t = 46
	} else
	if (isType(fileInfo.TypeHash, "Puzzle")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 47 is the identifier relative to Puzzle
		t = 47
	} else
	if (isType(fileInfo.TypeHash, "PuzzleSecret")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 48 is the identifier relative to PuzzleSecret
		t = 48
	} else
	if (isType(fileInfo.TypeHash, "PuzzleSolution")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 49 is the identifier relative to PuzzleSolution
		t = 49
	} else
	if (isType(fileInfo.TypeHash, "RelabelPermission")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 50 is the identifier relative to RelabelPermission
		t = 50
	} else
	if (isType(fileInfo.TypeHash, "RpcCrypto")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 51 is the identifier relative to RpcCrypto
		t = 51
	} else
	if (isType(fileInfo.TypeHash, "RpcDirectoryCache")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 52 is the identifier relative to RpcDirectoryCache
		t = 52
	} else
	if (isType(fileInfo.TypeHash, "RpcEfs")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 53 is the identifier relative to RpcEfs
		t = 53
	} else
	if (isType(fileInfo.TypeHash, "RpcMinimalt")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 54 is the identifier relative to RpcMinimalt
		t = 54
	} else
	if (isType(fileInfo.TypeHash, "RpcShadowDaemon")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 55 is the identifier relative to RpcShadowDaemon
		t = 55
	} else
	if (isType(fileInfo.TypeHash, "RpcUserSpaceDebug")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 56 is the identifier relative to RpcUserSpaceDebug
		t = 56
	} else
	if (isType(fileInfo.TypeHash, "Seconds")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 57 is the identifier relative to Seconds
		t = 57
	} else
	if (isType(fileInfo.TypeHash, "Signature")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 58 is the identifier relative to Signature
		t = 58
	} else
	if (isType(fileInfo.TypeHash, "SignatureKeyPair")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 59 is the identifier relative to SignatureKeyPair
		t = 59
	} else
	if (isType(fileInfo.TypeHash, "SignaturePrivateKey")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 60 is the identifier relative to SignaturePrivateKey
		t = 60
	} else
	if (isType(fileInfo.TypeHash, "SignaturePublicKey")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 61 is the identifier relative to SignaturePublicKey
		t = 61
	} else
	if (isType(fileInfo.TypeHash, "Status")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 62 is the identifier relative to Status
		t = 62
	} else
	if (isType(fileInfo.TypeHash, "Time")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 64 is the identifier relative to Time
		t = 64
	} else
	if (isType(fileInfo.TypeHash, "Time64")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 65 is the identifier relative to Time64
		t = 65
	} else
	if (isType(fileInfo.TypeHash, "TypeNode")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 66 is the identifier relative to TypeNode
		t = 66
	} else
	if (isType(fileInfo.TypeHash, "Types")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 67 is the identifier relative to Types
		t = 67
	} else
	if (isType(fileInfo.TypeHash, "UdpPort")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 68 is the identifier relative to UdpPort
		t = 68
	} else
	if (isType(fileInfo.TypeHash, "UserCertificate")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 73 is the identifier relative to UserCertificate
		t = 73
	} else
	if (isType(fileInfo.TypeHash, "UserGroupCertificate")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 74 is the identifier relative to UserGroupCertificate
		t = 74
	} else
	if (isType(fileInfo.TypeHash, "UserRecord")) {
		logger.Printf("type recognized")
		fd, status = altEthos.DirectoryOpen(dname)
		check(status, "DirectoryOpen failed")
		// Read the file in raw form (only bytes)
		bytes, status = altEthos.ReadVarRaw(fd, fname)
		check(status, "ReadVarRaw failed")
		// 75 is the identifier relative to UserRecord
		t = 75
	} else {
		// Another type
		logger.Printf("wrong type mate")
	}
	logger.Printf("Bytes[%v]: %v\n", reflect.TypeOf(bytes), bytes)
	logger.Printf("Type[%v]: %v\n", reflect.TypeOf(t), t)
	logger.Printf("Name[%v]: %v\n", reflect.TypeOf(fname), fname)
	param := Param{bytes, t, fname}
	call := MyRpcFileTransfer{param}
	status = altEthos.ClientCall(sock, &call)
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

func check(status syscall.Status, s string) {
	if status != syscall.StatusOk {
		logger.Fatalf("myRpcClient: " + s + " - %v\n",  status)
	}
}

func parseName(str string) (directory string, filename string) {
	filename = filepath.Base(str)
	i := strings.LastIndex(str, filename)
	directory = str[0:i]
	return
}
