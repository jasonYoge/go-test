package main

import (
	"fmt"
	"runtime"
)

//func RunMyProgram() {
//	trace.Log(context.TODO(), "myID", "123")
//	fmt.Println("Hello Trace")
//}

func stackExample() {
	stackSlice := make([]byte, 512)
	s := runtime.Stack(stackSlice, false)
	fmt.Printf("\n%s", stackSlice[0:s])
}

func First() {
	Second()
}

func Second() {
	Third()
}

func Third() {
	for c := 0; c < 5; c++ {
		fmt.Println(runtime.Caller(c))
	}
}

func main() {
	//f, _ := os.Create("trace.out")
	//defer f.Close()
	//
	//trace.Start(f)
	//defer trace.Stop()
	//
	//RunMyProgram()
	stackExample()
	First()
}
