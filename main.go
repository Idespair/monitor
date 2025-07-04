//Executes the program

// Declares the package name
package main

//Imports packages required for it to work

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Idespair/monitor/hardware"
)

type server struct {
	subscriberMessageBuffer int
	mux                     http.ServeMux
}

func NewServer() *server {
	s := &server{
		subscriberMessageBuffer: 10,
	}

	s.mux.Handle("/", http.FileServer(http.Dir("./htmx")))

	return s
}

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

	srv := NewServer()
	err := http.ListenAndServe(":8080", &srv.mux)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
