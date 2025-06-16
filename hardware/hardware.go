// Responsible for reading the hardware information

// Declares the package name
package hardware

//Imports packages required for it to work
import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// Function to get system information
func GetSystemSection() (string, error) {

	//Retrives the Operational System
	runTimeOS := runtime.GOOS

	//Retrives memory statistics
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}

	//Retrives information from the host computer
	hostStat, err := host.Info()
	if err != nil {
		return "", err
	}

	//Formats the output to send it as a return
	output := fmt.Sprintf("Hostname: %s\nTotal Memory: %d\nUsed Memory: %d\nOS: %s",
		hostStat.Hostname, vmStat.Total, vmStat.Used, runTimeOS)

	//Returns the output and any possible error
	return output, nil
}
