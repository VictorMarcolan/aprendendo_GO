package main

import (
    "fmt"
)

func main() {
    var nome string

    fmt.Print("Digite seu nome: ")
    fmt.Scanln(&nome)

    fmt.Printf("olá, %s!\n", nome)
}
