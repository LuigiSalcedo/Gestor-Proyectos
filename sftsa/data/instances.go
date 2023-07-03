package data

func NewOccupation(code int32, name string, salary uint32) Occupation {
	job := Occupation{}
	job.Code = code
	job.Name = name
	job.Salary = Money(salary)
	return job
}

func NewEmployeed(code int32, name string, edad uint8, job Occupation) Employeed {
	emp := Employeed{}
	emp.Code = code
	emp.Name = name
	emp.Age = edad
	emp.Job = job
	return emp
}

func NewProject(code int32, name string) Project {
	project := Project{}
	project.Code = code
	project.Name = name
	project.Members = make([]Employeed, 0, 10)
	return project
}