package main

import "fmt"

func main() {

	//declarando variavel
	var variavel1 string = "Variável 1"
	//declarando variavel simplificado
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// declarando mais que uma variavel
	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)

	fmt.Println(variavel3, variavel4)
	// declarando mais que uma variavel somplificado
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)
	// declarando const não pode ser alterada
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// invertendo valores ds variaveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
