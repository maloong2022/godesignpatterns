package abstractfactory

func ExampleDoorFactory() {
	// 创建一个木门工厂
	woodenFactory := &WoodenDoorFactory{}

	// 使用木门工厂创建门和门把手
	door := woodenFactory.CreateDoor()
	doorHandle := woodenFactory.CreateDoorHandler()

	// 使用创建的门和门把手
	door.Open()
	doorHandle.Press()
	// Output:
	// Wooden door is opened
	// Wooden handle is pressed
}
