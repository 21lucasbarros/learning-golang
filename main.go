package main

import "fmt"

func main() {
	/*
		    *	var nome string = "Lucas"
			*   var idade int = 23
			*   var altura float64 = 1.75
			*   var estudante bool = true
	*/

	// o operador := declara uma variavel e infere seu tipo
	nome := "Lucas"
	idade := 23
	altura := 1.80
	estudante := true

	fmt.Printf("%s %d %.2f %t\n", nome, idade, altura, estudante)

	/*
	 * Símbolo    Significado                  Exemplo
	 * %s         String                       Lucas
	 * %d         Inteiro (decimal)            23
	 * %f         Número decimal               1.800000
	 * %.2f       Decimal com 2 casas          1.80
	 * %t         Boolean                      true
	 * %v         Valor em formato padrão      Lucas
	 * %T         Tipo da variável             string
	 */
}
