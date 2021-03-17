package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
	brand string
}

//Implements the interface witho a pointer receiver
//使用指针类型时，可以改变值，对外是生效的。
func (nokiaPhone *NokiaPhone) call() {
	nokiaPhone.brand = "nokia"
	fmt.Println("I am " + nokiaPhone.brand + ", I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone

	//phone = new(NokiaPhone)
	phone = &NokiaPhone{"444"}
	phone.call()
	var phone2 = &NokiaPhone{"444"}
	phone2.call()
	fmt.Println(phone2.brand)

	phone = new(IPhone)
	phone.call()

}
