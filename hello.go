package main

import (
	"fmt"
	"net/http"
	"os"
)

func showIntro() {
	name := "My"
	version := 1.0
	fmt.Println("Olá", name)
	fmt.Println("This program is currently on version ", version)
}

func showMenu() {
	fmt.Println("1 - Monitorar")
	fmt.Println("2 - Logs")
	fmt.Println("0 - Sair")
}

func getCommand() int {
	var something int
	fmt.Scan(&something)

	return something
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	sites := []string{"https://www.izirh.io", "https://www.alura.com.br", "https://www.google.com"}

	for _, site := range sites {
		trySite(site)
	}

	fmt.Println("")
}

func trySite(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site ", site, "está funcionando corretamente")
	} else {
		fmt.Println("Deu problema. Status:", response.StatusCode)
	}
}

func main() {

	showIntro()

	for {
		showMenu()
		something := getCommand()

		switch something {
		case 1:
			startMonitoring()

		case 2:
			fmt.Println("Exibindo logs...")

		case 0:
			fmt.Println("Bye")
			os.Exit(0)

		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
	}
}
