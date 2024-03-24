package utils

import "main/data/response"

type MinLenghtPersons struct{
	Name string
	RealIndex int
	Projects []response.Projects
}

func ErrorPanic(err error){
	if err !=nil{
		panic(err)
	  }
}

func Filter(ss []response.PersonTasks, test func(response.PersonTasks) bool) (ret []MinLenghtPersons) {
    for i, s := range ss {
        if test(s) {
            ret = append(ret, MinLenghtPersons{
				Name:s.Name,
				RealIndex: i,
				Projects: s.Projects,
			} )
        }
    }
    return
}