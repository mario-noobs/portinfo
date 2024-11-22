// process.go

package portinfo

import (
	"os/exec"
	"strings"
	"encoding/json"
)

// Process struct represents the information of a process.
type Process struct {
	Port    string `json:"port"`
	PID     string `json:"pid"`
	User    string `json:"user"`
	Command string `json:"command"`
}

// GetProcesses fetches processes running on listening ports and returns them as a slice of Process structs.
func GetProcesses() ([]Process, error) {
	// Execute the `lsof` command to get listening processes.
	out, err := exec.Command("lsof", "-i", "-P", "-n", "-sTCP:LISTEN").Output()
	if err != nil {
		return nil, err
	}

	// Convert output to string and split by lines.
	strStdout := string(out)
	lines := strings.Split(strStdout, "\n")

	var processes []Process

	// Iterate through each line of the output.
	for i, line := range lines {
		// Skip empty lines and the header.
		if len(line) == 0 || i == 0 {
			continue
		}

		// Split the line into fields.
		fields := strings.Fields(line)
		if len(fields) < 9 {
			continue
		}

		// Extract relevant information.
		pid := fields[1]
		user := fields[2]
		port := strings.Split(fields[8], ":")[1]
		command := fields[0]

		// Append the process to the result list.
		processes = append(processes, Process{
			Port:   port,
			PID:    pid,
			User:   user,
			Command: command,
		})
	}

	return processes, nil
}

// ToJSON serializes the list of processes into a single-line JSON string.
func ToJSON(processes []Process) (string, error) {
    jsonData, err := json.Marshal(processes)
    if err != nil {
        return "", err
    }
    return string(jsonData), nil
}
