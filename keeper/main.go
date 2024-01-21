package main

import (
	"limit/conf"
	"limit/contract"
	"limit/subgraph"
	"limit/cron"
	"limit/executor"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf.Init()
	contract.Init()
	subgraph.Init()
	cron.Init()
	go executor.Run()

	//如果监听到系统信号 SIGQUIT 就退出程序，否则一直阻塞
	exitChan := make(chan int)
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		log.Print("Received signal SIGQUIT")
		exitChan <- 1
	}()
	signal.Notify(signalChan, syscall.SIGQUIT)
	<-exitChan
}
