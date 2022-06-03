package auxiliar

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func Escrver() {
	fmt.Println("Escrever publico")
	escrever2()
	erro := checkmail.ValidateFormat("phaaphaa.123phaa@gmail.com")
	fmt.Println(erro)
}
