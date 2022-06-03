package auxiliar

import (
	"fmt"

	"github.com/badoux/checkmail"
)

// funação pubilca
func Escrver() {
	fmt.Println("Escrever publico")
	escrever2()
	// função de outro repositorio
	erro := checkmail.ValidateFormat("phaaphaa.123phaa@gmail.com")
	fmt.Println(erro)
}
