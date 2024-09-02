package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

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
	sites := getSitesFromFilte()

	for i := 0; i < monitoring; i++ {
		for _, site := range sites {
			trySite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")
}

func trySite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site ", site, "está funcionando corretamente")
		logStart(site, true)
	} else {
		fmt.Println("Deu problema. Status:", response.StatusCode)
		logStart(site, false)
	}
}

func getSitesFromFilte() []string {
	sites := []string{}

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if err == io.EOF {
			break
		} else {
			sites = append(sites, line)
		}
	}

	file.Close()

	return sites
}

func logStart(site string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func getLogs() {
	file, err := os.ReadFile("logs.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}

func main() {

	showIntro()

	logStart("Log start", false)

	for {
		showMenu()
		something := getCommand()

		switch something {
		case 1:
			startMonitoring()

		case 2:
			fmt.Println("Exibindo logs...")
			getLogs()
		case 0:
			fmt.Println("Bye")
			os.Exit(0)

		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
	}
}
