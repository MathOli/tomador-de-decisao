package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	//pergunta
	intro()

	// criando um slice de strings
	frases := criandoLista()

	//gerando um numero aleatorio
	numAl := gerandoNumero(frases)

	//imprimindo a frase
	imprimirFrase(frases, numAl)

}

func intro() {
	var pergunta string
	fmt.Println("Olá, seja bem vindo ao Tomador de Decisão!")
	fmt.Print("Faça sua pergunta: ")
	fmt.Scan(&pergunta)
}

func criandoLista() []string {
	var frases []string

	arquivo, err := os.Open("frases.txt")

	if err != nil {
		fmt.Println("Erro ao abrir arquivo")
		return []string{"erro"}
	}

	leitor := bufio.NewReader(arquivo)

	for {
		frase, err := leitor.ReadString('\n')
		frase = strings.TrimSpace(frase)

		if err == io.EOF {
			break
		}

		frases = append(frases, frase)
	}

	arquivo.Close()

	return frases
}

func gerandoNumero(frases []string) int {
	rand.Seed(time.Now().UnixNano())
	numAl := rand.Intn(len(frases)-0) + 0
	fmt.Println("tamanho: ", len(frases))
	return numAl
}

func imprimirFrase(frases []string, numAl int) {
	resposta := fmt.Sprintf("resposta: %s", frases[numAl])
	fmt.Println(resposta)
}
