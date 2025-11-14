package main

import "fmt"

func main() {
	var num1, num2 float64
	var operacao string

	fmt.Println("=== Calculadora Simples ===")
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Print("Digite a operação (+, -, *, /, %): ")
	fmt.Scanln(&operacao)

	// Operadores aritméticos
	var resultado float64
	switch operacao {
	case "+":
		resultado = num1 + num2
		fmt.Printf("\n%.2f + %.2f = %.2f\n", num1, num2, resultado)
	case "-":
		resultado = num1 - num2
		fmt.Printf("\n%.2f - %.2f = %.2f\n", num1, num2, resultado)
	case "*":
		resultado = num1 * num2
		fmt.Printf("\n%.2f * %.2f = %.2f\n", num1, num2, resultado)
	case "/":
		if num2 != 0 {
			resultado = num1 / num2
			fmt.Printf("\n%.2f / %.2f = %.2f\n", num1, num2, resultado)
		} else {
			fmt.Println("\nErro: Divisão por zero!")
		}
	case "%":
		// Para módulo, precisamos converter para int
		resultadoInt := int(num1) % int(num2)
		fmt.Printf("\n%d %% %d = %d\n", int(num1), int(num2), resultadoInt)
	default:
		fmt.Println("\nOperação inválida!")
	}

	// Demonstração de operadores relacionais
	fmt.Println("\n=== Operadores Relacionais ===")
	fmt.Printf("%.2f == %.2f: %t\n", num1, num2, num1 == num2)
	fmt.Printf("%.2f != %.2f: %t\n", num1, num2, num1 != num2)
	fmt.Printf("%.2f > %.2f: %t\n", num1, num2, num1 > num2)
	fmt.Printf("%.2f < %.2f: %t\n", num1, num2, num1 < num2)
	fmt.Printf("%.2f >= %.2f: %t\n", num1, num2, num1 >= num2)
	fmt.Printf("%.2f <= %.2f: %t\n", num1, num2, num1 <= num2)

	// Demonstração de operadores lógicos
	fmt.Println("\n=== Operadores Lógicos ===")
	condicao1 := num1 > 0
	condicao2 := num2 > 0

	fmt.Printf("num1 > 0: %t\n", condicao1)
	fmt.Printf("num2 > 0: %t\n", condicao2)
	fmt.Printf("num1 > 0 && num2 > 0: %t (AND)\n", condicao1 && condicao2)
	fmt.Printf("num1 > 0 || num2 > 0: %t (OR)\n", condicao1 || condicao2)
	fmt.Printf("!(num1 > 0): %t (NOT)\n", !condicao1)
}


