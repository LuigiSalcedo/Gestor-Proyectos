package repository

import(
	"sftsa/data"
	"errors"
	"strings"
)

// Manejo de errores
var (
	NotValidIDErr = errors.New("Código de registro no valido.")
	NotValidNameErr = errors.New("Nombre de registro no valido.")
	AlreadyRegisterErr = errors.New("El código ya se encuentra registrado.")
	NotValidSalaryErr = errors.New("Salario ingresado no es valido.")
)

// Estrucutra de almacenaje
var (
	employeedsData = map[int32]data.Employeed{}
	jobsData = map[int32]data.Occupation{}
	projectsData = map[int32]data.Project{}
)

// Funciones generales
func validateGeneralData(element *data.Element) error {
	element.Name = strings.Trim(element.Name, " ")
	if element.Name == "" {
		return NotValidNameErr
	}

	if element.Code <= 0 {
		return NotValidIDErr
	}

	return nil
}

// Funciones para el manejo de ocupaciones
func AddOccupation(occupation data.Occupation) error {
	err := validateGeneralData(&occupation.Element)

	if err != nil {
		return err
	}

	_, registered := jobsData[occupation.Code]
	if registered {
		return AlreadyRegisterErr
	}

	jobsData[occupation.Code] = occupation
	return nil
}

func RemoveOccupation(code int32) bool {
	delete(jobsData, code)
	_, deleted := jobsData[code]
	return !deleted
}

func SearchOccupation(code int32) (data.Occupation, bool) {
	data, ok := jobsData[code]
	return data, ok
}

func UpdateOccupation(job data.Occupation) {
	jobsData[job.Code] = job
}

// Funciones para el manejo de empleados
func AddEmployeed(emp data.Employeed) error {
	err := validateGeneralData(&emp.Element)

	if err != nil {
		return err
	}

	_, registered := employeedsData[emp.Code]
	if registered {
		return AlreadyRegisterErr
	}

	employeedsData[emp.Code] = emp
	return nil
}

func RemoveEmployeed(code int32) bool {
	delete(employeedsData, code)
	_, deleted := employeedsData[code]
	return !deleted
}

func SearchEmployeed(code int32) (data.Employeed, bool) {
	data, ok := employeedsData[code]
	return data, ok
}

func UpdateEmployeed(emp data.Employeed) {
	employeedsData[emp.Code] = emp
}

// Funciones de manejo de proyectos

func AddProject(project data.Project) error {
	err := validateGeneralData(&project.Element)

	if err != nil {
		return err
	}

	_, registered := projectsData[project.Code]
	if registered {
		return AlreadyRegisterErr
	}

	projectsData[project.Code] = project
	return nil
}

func RemoveProject(code int32) bool {
	delete(projectsData, code)
	_, deleted := projectsData[code]
	return !deleted
}

func SearchProject(code int32) (data.Project, bool) {
	data, ok := projectsData[code]
	return data, ok
}

func UpdateProject(project data.Project) {
	projectsData[project.Code] = project
}

func AddMember(p *data.Project, emp data.Employeed) bool {
	for _, emps := range p.Members {
		if emps.Code == emp.Code {
			return false
		}
	}
	p.Members = append(p.Members, emp)
	return true
}