//Executes the program

// Declares the package name
package main

//Imports packages required for it to work

import (
	"fmt"
	"time"

	"github.com/Idespair/monitor/hardware"
)

// Starts the main function
func main() {
	fmt.Println("Starting system monitor..")

	go func() {
		for {
			systemSection, err := hardware.GetSystemSection()
			if err != nil {
				fmt.Println(err)
			}

			cpuSection, err := hardware.GetCpuSection()
			if err != nil {
				fmt.Println(err)
			}

			diskSection, err := hardware.GetDiskSection()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(systemSection)
			fmt.Println(cpuSection)
			fmt.Println(diskSection)

			time.Sleep(1 * time.Minute)
		}
	}()

	time.Sleep(5 * time.Minute)
}
