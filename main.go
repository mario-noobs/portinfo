// main.go

package main

import (
	"fmt"
	"log"
	"github.com/mario-noobs/portinfo/port"
)

func main() {
	// Get processes
	processes, err := portinfo.GetProcesses()
	if err != nil {
		log.Fatal(err)
	}
	
	// Convert the processes to JSON
	jsonData, err := portinfo.ToJSON(processes)
	if err != nil {
		log.Fatal(err)
	}

	// Print the JSON data
	fmt.Println(jsonData)
}
