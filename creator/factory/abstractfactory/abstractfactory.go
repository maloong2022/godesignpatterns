package abstractfactory

import "fmt"

// DoorFactory是抽象工厂接口，定义了创建门和门把手的方法
type DoorFactory interface {
	CreateDoor() Door
	CreateDoorHandler() DoorHandle
}

// Door 是门接口
type Door interface {
	Open()
	Close()
}

// DoorHandle 是门把手接口
type DoorHandle interface {
	Press()
}

// WoodenDoorFactory 是一个具体的木门工厂，实现了 DoorFactory接口
type WoodenDoorFactory struct{}

func (f *WoodenDoorFactory) CreateDoor() Door {
	return &WoodenDoor{}
}

func (f *WoodenDoorFactory) CreateDoorHandler() DoorHandle {
	return &WoodenDoorHandle{}
}

// WoodenDoor 是木门实现
type WoodenDoor struct{}

func (d *WoodenDoor) Open() {
	fmt.Println("Wooden door is opened")
}

func (d *WoodenDoor) Close() {
	fmt.Println("Wooden door is closed")
}

// WoodenDoorHandle 是木门把手实现
type WoodenDoorHandle struct{}

func (h *WoodenDoorHandle) Press() {
	fmt.Println("Wooden handle is pressed")
}
