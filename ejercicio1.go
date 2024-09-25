package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func imprimir() {
	fmt.Println("Texto desde la funcion imprimir")
}

func hola_string(s string) string {
	return "hola " + s
}

func sumar(a int, b int) int {
	return a + b
}

func comparar_bool() {
	var a bool
	b := true
	a = false
	if a == b {
		fmt.Println("A y B son igules")
	} else {
		fmt.Println("A y B no son igules")
	}
}

func arreglo() {
	var aprendices [5]string

	aprendices[0] = "Leonardo"
	aprendices[1] = "Juan Sebastian"
	aprendices[2] = " Frank"
	aprendices[3] = "Juan Jose"
	aprendices[4] = " Jaider"

	fmt.Println(aprendices[0])
}

func tipo_datos() {
	var texto string = "Fabian"
	var numero int = 1000
	var decimal float64 = 0.00001

	fmt.Println(reflect.TypeOf(texto))
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println(reflect.TypeOf(decimal))
}

func converstringtoboolean() {
	var palabra string = "true"

	boolean, err := strconv.ParseBool(palabra)
	fmt.Println(boolean, reflect.TypeOf(boolean))
	fmt.Println("Este es el error", err)
}
func converbooltostring() {
	var palabra_bool bool = true

	strbool := strconv.FormatBool(palabra_bool)
	fmt.Println(strbool, reflect.TypeOf(strbool))
}

func main() {
	//fmt.Println("texto desde la funcion main")
	//imprimir()
	//fmt.Println(hola_string("Fabian"))
	//fmt.Println(sumar(3, 5))
	//comparar_bool()
	//arreglo()
	//tipo_datos()
	converstringtoboolean()
	//converbooltostring()
}
