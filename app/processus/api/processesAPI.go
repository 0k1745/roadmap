package api

import (
	"encoding/json"
	"log"
	"net/http"
	"processus/service"
	"strconv"
	"strings"
)

type listOfResources struct {
	Resources []location `json:"resources"`
}

type location struct {
	Location *string `json:"location"`
	Type     string  `json:"type"`
}

type ram struct {
	Max     *int `json:"max"`
	current *int `json:"current"`
}

type ProcessDetails struct {
	Name *string `json:"name"`
	Ram  ram     `json:"ram" `
}

func handlerProcess(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id, _ := strconv.Atoi(segments[len(segments)-1])
	pidDetail := service.GetPid(id)
	convertToJsonResponse(w, pidDetail)

}

func handlerListProcesses(w http.ResponseWriter, r *http.Request) {
	pids := service.GetPids()
	var processes = make([]location, len(pids))
	for index, pid := range pids {
		locationToProcess := "/process/" + strconv.Itoa(pid)
		processes[index] = location{
			Location: &locationToProcess,
			Type:     "process",
		}
	}
	processesList := listOfResources{Resources: processes}
	convertToJsonResponse(w, processesList)
}

func convertToJsonResponse(w http.ResponseWriter, object any) {
	jsonData, err := json.Marshal(object)
	if err != nil {
		log.Print("Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}

func GetProcessesConfigurationsPaths() map[string]func(http.ResponseWriter, *http.Request) {
	var myMap map[string]func(w http.ResponseWriter, r *http.Request)
	myMap = make(map[string]func(w http.ResponseWriter, r *http.Request))
	myMap["/process/"] = handlerProcess
	myMap["/process"] = handlerListProcesses
	return myMap
}
