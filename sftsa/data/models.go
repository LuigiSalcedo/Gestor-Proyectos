package data

import (
	"strings"
	"fmt"
)

type Money uint32

type Element struct {
	Code int32
	Name string
}

type Occupation struct {
	Element
	Salary Money
}

type Employeed struct {
	Element
	Age uint8
	Job Occupation
}

type Project struct {
	Element
	Finished bool
	Members []Employeed
}

func (m Money) String() string {
	return fmt.Sprintf("$%d", m)
}

func (e Element) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Código: %d\n", e.Code))
	sb.WriteString(fmt.Sprintf("Nombre: %s\n", e.Name))
	return sb.String()
}

func (o Occupation) String() string {
	var sb strings.Builder
	sb.WriteString("--------------- Datos Cargo ---------------- \n")
	sb.WriteString(o.Element.String())
	sb.WriteString(fmt.Sprintf("Salario: %v", o.Salary))
	sb.WriteString("\n-------------------------------------------- ")
	return sb.String()
}

func (e Employeed) String() string {
	var sb strings.Builder
	sb.WriteString("*********************** Datos Personales ***************************** \n")
	sb.WriteString(e.Element.String())
	sb.WriteString(fmt.Sprintf("Edad: %d\n", e.Age))
	sb.WriteString(fmt.Sprintf("%v", e.Job))
	sb.WriteString("\n*********************************************************************")
	return sb.String()
}

func (p Project) String() string {
	var sb strings.Builder
	var finished string
	if p.Finished {
		finished = "Sí"
	} else {
		finished = "No"
	}
	sb.WriteString(p.Element.String())
	sb.WriteString("Finalizado: ")
	sb.WriteString(finished)
	sb.WriteString("\n===================== Datos de los empleados =============================\n")
	for _, emp := range p.Members {
		sb.WriteString(emp.String())
		sb.WriteString("\n")
	}
	sb.WriteString("================================================================================")
	return sb.String()
}
