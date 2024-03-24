package service

type memory struct {
	Max     *int
	Current *int
}

type status struct {
	Name   *string
	Memory *memory
	State  *string
	ID     *int
}

func GetPids() []int {
	return findAllPid()
}

func GetPid(id int) status {
	return status{
		Name:   getName(id),
		Memory: getMemory(id),
		State:  getState(id),
		ID:     &id,
	}

}
