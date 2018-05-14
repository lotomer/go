package common

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// SignalHandle 系统信号钩子函数
type SignalHandle func()

// ProgramSignalHandle 程序信号处理函数
func ProgramSignalHandle(exitHandle SignalHandle, user1Handle SignalHandle, user2Handle SignalHandle) {
	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	sigusr1 := syscall.Signal(0xa)
	sigusr2 := syscall.Signal(0xc)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, sigusr1, sigusr2)
	//signal.Notify(c)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				if exitHandle != nil {
					exitHandle()
				} else {
					fmt.Println("退出", s)
				}
			case sigusr1:
				if user1Handle != nil {
					user1Handle()
				} else {
					fmt.Println("usr1", s)
				}
			case sigusr2:
				if user2Handle != nil {
					user2Handle()
				} else {
					fmt.Println("usr2", s)
				}
			default:
				fmt.Println("other", s)
			}
		}
	}()
}
