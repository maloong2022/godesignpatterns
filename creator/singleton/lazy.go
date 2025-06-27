package singleton

import (
	"fmt"
	"sync"
)

// 单例模式 -> 饿汉模式

// 初始化一个全局的单例变量
var (
	lazy *Lazy
	once sync.Once
)

// 定义单例模式类型
type Lazy struct{}

// 定义 Lazy 类型 SayHi 方法
func (lz *Lazy) SayHi() {
	fmt.Println("Hi")
}

// GetLazy 不加锁，可能会存在并发问题
// 初始化并获取全局单例实例。
// 懒汉方法在获取初始化实例时，可能需要传入参数，这样不是很优雅
// 所以实际开发中，我更喜欢用饿汉模式
// 使用once.Do可以确保 lazy 实例全局只被创建一次，once.Do 函数
// 还可以确保当同时有多个创建动作时，只有一个创建动作在被执行。
// 使用这种方法来获取单例实例，既提高了代码效率，又保证了并发安全。
func GetLazy() *Lazy {
	once.Do(func() {
		lazy = &Lazy{}
	})
	return lazy
}
