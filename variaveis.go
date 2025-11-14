package main

import "fmt"

func main() {
	// Declaração de variáveis com var
	var nome string
	var idade int
	var altura float64
	var maiorDeIdade bool

	// Atribuição de valores
	nome = "João"
	idade = 25
	altura = 1.75
	maiorDeIdade = true

	fmt.Println("=== Declaração com var ===")
	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Idade: %d\n", idade)
	fmt.Printf("Altura: %.2f\n", altura)
	fmt.Printf("Maior de idade: %t\n", maiorDeIdade)

	// Declaração com inferência de tipo (:=)
	nome2 := "Maria"
	idade2 := 18
	altura2 := 1.65
	maiorDeIdade2 := false

	fmt.Println("\n=== Declaração com inferência (:=) ===")
	fmt.Printf("Nome: %s\n", nome2)
	fmt.Printf("Idade: %d\n", idade2)
	fmt.Printf("Altura: %.2f\n", altura2)
	fmt.Printf("Maior de idade: %t\n", maiorDeIdade2)

	// Declaração múltipla
	var (
		numero1 int     = 10
		numero2 int     = 20
		pi      float64 = 3.14
	)

	fmt.Println("\n=== Declaração múltipla ===")
	fmt.Printf("Número 1: %d\n", numero1)
	fmt.Printf("Número 2: %d\n", numero2)
	fmt.Printf("PI: %.2f\n", pi)

	// Declaração com inferência múltipla
	x, y, z := 5, 10, 15
	fmt.Println("\n=== Declaração múltipla com inferência ===")
	fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)

	// Demonstração de tipos básicos
	fmt.Println("\n=== Tipos Básicos ===")
	var texto string = "Olá, Go!"
	var numero int = 42
	var decimal float64 = 3.14159
	var verdadeiro bool = true

	fmt.Printf("String: %s (tipo: %T)\n", texto, texto)
	fmt.Printf("Int: %d (tipo: %T)\n", numero, numero)
	fmt.Printf("Float64: %.5f (tipo: %T)\n", decimal, decimal)
	fmt.Printf("Bool: %t (tipo: %T)\n", verdadeiro, verdadeiro)
}


