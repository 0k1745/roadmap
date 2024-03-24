package service

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Cache struct {
	expiration time.Time
	value      string
}

var cache = make(map[int]*Cache)

const pattern = "[0-9]+"

var regexProcessFiles = regexp.MustCompile(pattern)

const patternName = "Name:[\t ]*(.*)"

var regexProcessName = regexp.MustCompile(patternName)

const patternMemoryMax = "VmPeak:[\t ]*([0-9]*) kB"

var regexProcessMemoryMax = regexp.MustCompile(patternMemoryMax)

const patternMemoryCurrent = "VmHWM:[\t ]*([0-9]*) kB"

var regexProcessMemoryCurrent = regexp.MustCompile(patternMemoryCurrent)

const patternState = "State:[\t ]*(.*)"

var regexProcessState = regexp.MustCompile(patternState)

func findAllPid() []int {
	dir, _ := os.ReadDir("/proc/")
	var pids = make([]int, len(dir))
	nbElements := 0
	for index, entry := range dir {
		if regexProcessFiles.MatchString(entry.Name()) {
			pid, _ := strconv.Atoi(entry.Name())
			pids[index] = pid
			nbElements++
		}
	}

	return pids[:nbElements]
}

func getName(id int) *string {
	content, err := getStatusProcessFile(id)
	if err != nil {
		return nil
	}
	return &regexProcessName.FindStringSubmatch(content)[1]
}

func getStatusProcessFile(id int) (string, error) {
	var content string
	if cache[id] == nil || time.Now().After(cache[id].expiration) {
		filePath := "/proc/" + strconv.Itoa(id) + "/status"
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			// Check if the error indicates that the file doesn't exist
			if os.IsNotExist(err) {
				fmt.Println("File does not exist:", filePath)
			} else {
				// Other error occurred
				fmt.Println("Error reading file:", err)
			}
			return "", err
		}
		content = string(fileContent)
		cacheValue := Cache{
			expiration: time.Now().Add(time.Second * 20),
			value:      content,
		}
		cache[id] = &cacheValue
		log.Print("File cache refreshed")
	} else {
		content = cache[id].value
	}
	return content, nil
}

func getMemory(id int) *memory {
	content, err := getStatusProcessFile(id)
	if err != nil {
		return nil
	}
	vmPeak := regexProcessMemoryMax.FindStringSubmatch(content)
	// In kB
	maxMemory, _ := strconv.Atoi(vmPeak[1])
	vmHWM := regexProcessMemoryCurrent.FindStringSubmatch(content)
	// In kB
	currentMemory, _ := strconv.Atoi(vmHWM[1])
	return &memory{
		Max:     &maxMemory,
		Current: &currentMemory,
	}
}

func getState(id int) *string {
	content, err := getStatusProcessFile(id)
	if err != nil {
		return nil
	}
	state := parseState(content)
	return &state
}

func parseState(content string) string {
	processState := regexProcessState.FindStringSubmatch(content)[1]
	//https://github.com/torvalds/linux/blob/bfa8f18691ed2e978e4dd51190569c434f93e268/fs/proc/array.c#L130
	switch processState {
	case "R (running)":
		return "running"
	case "S (sleeping)":
		return "sleeping"
	case "D (disk sleep)":
		return "disk sleep"
	case "T (stopped)":
		return "stopped"
	case "t (tracing stop)":
		return "tracing stop"
	case "X (dead)":
		return "dead"
	case "Z (zombie)":
		return "zombie"
	case "P (parked)":
		return "parked"
	case "I (idle)":
		return "idle"
	default:
		return "no status"
	}
}
