package main

import "fmt"

type Persona struct{	
	nombre string
	estatura float64
}

func main() {
	luis := Persona{"Luis David",1.80}
	luis.correr()
}

func (person *Persona)correr(){
	fmt.Printf("%s corriendo\n",person.nombre)
}
