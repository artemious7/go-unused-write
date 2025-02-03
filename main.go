package main

type MyType struct {
	value string
}

func (receiver MyType) Modify() {
	receiver.value = "new value"
}

func main() {
	instance := MyType{value: "old value"}
	instance.Modify()
}
