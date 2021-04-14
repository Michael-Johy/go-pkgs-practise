package main

import (
	"log"
	"unsafe"
)

type Person struct {
	name string
	age  int
}

/**
	Summary
    1.任何类型的指针值都可以转换成unsafe.Pointer
	2.unsafe.Pointer <=> uintptr
*/
func main() {

	p := &Person{}

	namePointer := (*string)(unsafe.Pointer(p))
	*namePointer = "yangTao"

	agePointer := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(namePointer)) + unsafe.Offsetof(p.age)))
	*agePointer = 11

	log.Println("name : ", p.name)
	log.Println("age : ", p.age)

}
