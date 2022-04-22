package main

import (
	"fmt"
	"time"
)

func main(){
	go recevier()
	for {
		fmt.Println("send a message..")
		sender()
		<-time.NewTicker(time.Second*1).C
	}
}
