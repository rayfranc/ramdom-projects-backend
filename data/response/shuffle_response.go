package response


type Projects struct {
	Name string `json:"name"`
}

type PersonTasks struct {
	Name string `json:"name"`
	Projects []Projects `json:"projects"`
}

type ShuffleResponse struct {
	Persons []PersonTasks `json:"persons"`
}