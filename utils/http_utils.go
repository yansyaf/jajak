package utils

import "log"

func ThrowPanic(e error) {
	if e != nil {
		log.Panicf("throwing panic %v", e)
		panic(e)
	}
}

func CommonPanicHandler() {
	if r := recover(); r != nil {
		log.Fatalf("Unhandled Panic: %v", r)
	}
}
