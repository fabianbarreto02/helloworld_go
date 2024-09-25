package main

import "fmt"

func equivalenciaEnPies(altura float32) float32 {
	altura = altura * 3.28
	return altura
}

var altura float32 = 1.70

func main() {
	fmt.Println("La altura es:", altura, "mts")
	fmt.Println("La altura es:", equivalenciaEnPies(altura), " pies")
	fmt.Println("La nueva altura es:", altura, "mts")
}
