package foo

import "fmt"

/*
#include "c_component/universe.c"
*/
import "C"

func Foo() {
	res := C.bar()
	fmt.Println("Answer to everything:", res)
}

func MyBar() {
	fmt.Println("The cake is a lie")
}
