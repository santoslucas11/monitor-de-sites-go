package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 4
const delay = 5

func main() {
	introducao()
	registraLogs("Iniciando monitoramento", true)

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			monitoramento()
		case 2:
			imprimeLog()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}

}

func devolveNome() (string, int) {
	nome := "Lucas"
	idade := 26

	return nome, idade
}

func introducao() {
	nome := "Lucas"
	versao := 1.1

	fmt.Println("Bem vindo,", nome)
	fmt.Println("Estamos na versão", versao)
}

func leComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Opção", comandoLido, "selecionada")
	fmt.Println("")

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func monitoramento() {
	fmt.Println("Iniciando monitoramento")

	sites := leArquivoTxt()

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println(i, site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Oops! Ocorreu um erro na chamada do site: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site retornando 200")
		registraLogs(site, true)
		fmt.Println(resp)
	} else {
		fmt.Println("Site fora do ar")
		registraLogs(site, false)
	}
}

func leArquivoTxt() []string {
	var sites []string

	arquivo, err := os.Open("targets.txt")
	if err != nil {
		fmt.Println("Oops! Ocorreu um erro na abertura do arquivo: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLogs(site string, status bool) {

	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	if status == false {
		arquivo.WriteString("[ERROR] " + time.Now().Format("02/01/2006 15:04:05") + " - " + site + " Status: offline " + "\n")

	} else {
		arquivo.WriteString("[ON AIR] " + time.Now().Format("02/01/2006 15:04:05") + " - " + site + " Status online " + "\n")
	}
	arquivo.Close()
}

func imprimeLog() {
	arquivo, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}
