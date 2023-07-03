package main

import (
	"sftsa/data"
	"sftsa/data/repository"
	"fmt"
)

func MostrarError(err error){
	fmt.Printf("\nUps, parace que ocurrió un error: %v\n\n", err)
}

func RegistrarCargo() {
	var codigo int32
	var nombre string
	var salario uint32

	fmt.Println("\n___________________ REGISTRANDO CARGO ___________________\n")
	fmt.Print("Código del cargo: ")
	fmt.Scan(&codigo)
	fmt.Print("Nombre del cargo: ")
	fmt.Scan(&nombre)
	fmt.Print("Salario del cargo: $")
	fmt.Scan(&salario)

	cargo := data.NewOccupation(codigo, nombre, salario)

	if err := repository.AddOccupation(cargo); err != nil {
		MostrarError(err)
		return
	}
	fmt.Println("\nCargo registrado correctamente\n")
}

func RegistrarEmpleado() {
	var codigo int32
	var nombre string
	var edad uint8
	var codigoCargo int32

	fmt.Println("\n___________________ REGISTRANDO EMPLEADO ___________________\n")

	fmt.Print("Ingrese el código del cargo del empleado: ")
	fmt.Scan(&codigoCargo)

	cargo, existe := repository.SearchOccupation(codigoCargo)

	if !existe {
		fmt.Println("\nParece que el no existe ningún cargo asociado a este código . . .\n")
		return
	}

	fmt.Printf("\nDatos del cargo a asingar:\n %v\n\n", cargo)
	fmt.Print("Código del empleado: ")
	fmt.Scan(&codigo)
	fmt.Print("Nombre del empleado: ")
	fmt.Scan(&nombre)
	fmt.Print("Edad: ")
	fmt.Scan(&edad)

	empleado := data.NewEmployeed(codigo, nombre, edad, cargo)

	if err := repository.AddEmployeed(empleado); err != nil {
		MostrarError(err)
		return
	}

	fmt.Println("\nEmpleado registrado correctamente\n")
}

func RegistrarProyecto() {
	var codigo int32
	var nombre string

	fmt.Println("\n___________________ REGISTRANDO PROYECTO ___________________\n")

	fmt.Print("Código del proyecto: ")
	fmt.Scan(&codigo)
	fmt.Print("Nombre del proyecto: ")
	fmt.Scan(&nombre)

	proyecto := data.NewProject(codigo, nombre)

	if err := repository.AddProject(proyecto); err != nil {
		MostrarError(err)
		return
	}

	fmt.Println("\nProyecto registrado correctamente\n")
}

func AsignarEmpleado() {
	var codigoEmpleado int32
	var codigoProyecto int32

	fmt.Println("\n___________________ ASIGNANDO EMPLEADO AL PROYECTO ___________________\n")

	fmt.Print("Ingrese código del empleado: ")
	fmt.Scan(&codigoEmpleado)

	empleado, ok := repository.SearchEmployeed(codigoEmpleado)

	if !ok {
		fmt.Println("\nParece que no existe ningún empleado registrado con este código . . .\n")
		return
	}

	fmt.Printf("Datos del empleado: \n%v\n\n", empleado)

	fmt.Print("Ingrese código del proyecto: ")
	fmt.Scan(&codigoProyecto)

	proyecto, ok := repository.SearchProject(codigoProyecto)

	if !ok {
		fmt.Println("\nParece que no existe ningún proyecto registrado con este código . . .\n")
		return 
	}

	ok = repository.AddMember(&proyecto, empleado)

	if !ok {
		fmt.Println("\nParace que este empleado ya estaba asignado a este proyecto . . .\n")
		return 
	}

	repository.UpdateProject(proyecto)

	fmt.Println("\nEmpleado asignado correctamente al proyecto . . .\n")
}

func ConsultarCargo() {
	var codigo int32

	fmt.Println("\n___________________ CONSULTANDO CARGO ___________________\n")

	fmt.Print("Ingrese código del cargo a consultar: ")
	fmt.Scan(&codigo)

	data, ok := repository.SearchOccupation(codigo)

	if !ok {
		fmt.Println("\nParece que no existe ningún cargo con este código . . .\n")
		return 
	}

	fmt.Printf("\n\n%v\n\n", data)
}

func ConsultarEmpleado() {
	var codigo int32

	fmt.Println("\n___________________ CONSULTANDO EMPLEADO ___________________\n")

	fmt.Print("Ingrese código del empleado a consultar: ")
	fmt.Scan(&codigo)

	data, ok := repository.SearchEmployeed(codigo)

	if !ok {
		fmt.Println("\nParece que no existe ningún empleado con este código . . .\n")
		return 
	}

	fmt.Printf("\n\n%v\n\n", data)
}

func ConsultarProyecto() {
	var codigo int32

	fmt.Println("\n___________________ CONSULTANDO PROYECTO ___________________\n")

	fmt.Print("Ingrese código del proyecto a consultar: ")
	fmt.Scan(&codigo)

	data, ok := repository.SearchProject(codigo)

	if !ok {
		fmt.Println("\nParece que no existe ningún proyecto con este código . . .\n")
		return 
	}

	fmt.Printf("\n\n%v\n\n", data)	
}

func main() {
	
	op := 0
	 
	for op != 8 {
		fmt.Println("Gestor de proyectos de SFT S.A")
		fmt.Println()
		fmt.Println("[1]. Registrar cargo")
		fmt.Println("[2]. Registrar empleados")
		fmt.Println("[3]. Registrar proyecto")
		fmt.Println("[4]. Asignar empleado a proyecto")
		fmt.Println("[5]. Consultar cargo")
		fmt.Println("[6]. Consultar empleado")
		fmt.Println("[7]. Consultar proyecto")
		fmt.Println()
		fmt.Println("[8]. Salir")
		fmt.Println()
		fmt.Print("Seleccionar una opción: ")
		fmt.Scan(&op)
		switch op {
			case 1:
				RegistrarCargo()
			case 2:
				RegistrarEmpleado()
			case 3:
				RegistrarProyecto()
			case 4:
				AsignarEmpleado()
			case 5:
				ConsultarCargo()
			case 6:
				ConsultarEmpleado()
			case 7:
				ConsultarProyecto()
			case 8:
				break
			default:
				fmt.Println("\nOpción seleccionada no valida\n")
		}
	}
}