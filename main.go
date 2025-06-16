package main

import (
	"fmt"
	"time"

	"github.com/Idespair/monitor/hardware"
)

func main() {
	fmt.Println("Starting system monitor..")

	go func() {
		for {
			systemSection, err := hardware.GetSystemSection()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(systemSection)

			time.Sleep(3 * time.Second)
		}
	}()

	time.Sleep(5 * time.Minute)
}
