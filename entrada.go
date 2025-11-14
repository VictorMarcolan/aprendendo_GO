package main

import "fmt"

func main() {
	var nome string
	var idade int
	var peso, altura float64

	fmt.Println("=== Entrada de Dados ===")

	// Lendo nome e idade
	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&idade)

	// Verificar maioridade
	fmt.Println("\n=== Verificação de Maioridade ===")
	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Idade: %d anos\n", idade)

	// Operadores relacionais para verificar maioridade
	maiorDeIdade := idade >= 18
	if maiorDeIdade {
		fmt.Println("Status: Maior de idade ✓")
	} else {
		fmt.Println("Status: Menor de idade")
	}

	// Calcular IMC
	fmt.Println("\n=== Cálculo de IMC ===")
	fmt.Print("Digite seu peso (kg): ")
	fmt.Scanln(&peso)

	fmt.Print("Digite sua altura (m): ")
	fmt.Scanln(&altura)

	// Cálculo do IMC: peso / (altura * altura)
	imc := peso / (altura * altura)

	fmt.Printf("\nNome: %s\n", nome)
	fmt.Printf("Idade: %d anos\n", idade)
	fmt.Printf("Peso: %.2f kg\n", peso)
	fmt.Printf("Altura: %.2f m\n", altura)
	fmt.Printf("IMC: %.2f\n", imc)

	// Classificação do IMC usando operadores relacionais e lógicos
	fmt.Print("Classificação: ")
	if imc < 18.5 {
		fmt.Println("Abaixo do peso")
	} else if imc >= 18.5 && imc < 25 {
		fmt.Println("Peso normal")
	} else if imc >= 25 && imc < 30 {
		fmt.Println("Sobrepeso")
	} else {
		fmt.Println("Obesidade")
	}

	// Demonstração adicional com fmt.Scan (lendo múltiplos valores)
	fmt.Println("\n=== Exemplo com fmt.Scan (múltiplos valores) ===")
	var a, b int
	fmt.Print("Digite dois números separados por espaço: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Você digitou: %d e %d\n", a, b)
	fmt.Printf("Soma: %d\n", a+b)
}


