package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
)

func main() {

	var stop bool

	// Abre o arquivo
	file, err := os.Open("huge.log")

	if err != nil {
		log.Fatalf("ERRO: %v", err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	// leitura linha a linha e coloca na lista newlist
	line, prefix, erro := r.ReadLine()
	if erro != nil {
	}
	log.Println(prefix)

	var newlist []string

	for stop == false {

		newlist = append(newlist, string(line))

		line, prefix, erro = r.ReadLine()

		if erro != nil {
			stop = true
		}

	}

	// logs e chama função qsort
	log.Println("antes qsort")
	newlist = qsort(newlist)
	log.Println("depois qsort")
	log.Println(len(newlist))

	// escreve o arquivo ordenado
	errFile := writeFile(newlist, "huge_orded.log")
	if errFile != nil {
		log.Fatalf("Erro:", errFile)
	} else {
		// aplica o md5sum
		setmd5()
	}
}

func writeFile(lines []string, fileName string) error {
	// Cria o arquivo de texto
	file, err := os.Create(fileName)
	// Caso tenha encontrado algum erro retornar ele
	if err != nil {
		return err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer file.Close()

	// Cria um escritor responsavel por escrever cada linha do slice no arquivo de texto
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	// Caso a funcao flush retorne um erro ele sera retornado aqui tambem
	return writer.Flush()
}

func setmd5() {
	f, err := os.Open("huge_orded.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
}

func qsort(a []string) []string {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Escolhe um pivo
	pivotIndex := rand.Int() % len(a)

	// Move o pivo para a direita
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Coloca a pilha de ementos menor a esquerda do pivo
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Coloca o pivo apos o ultimo elemento menor
	a[left], a[right] = a[right], a[left]

	// Vai fundo!!
	qsort(a[:left])
	qsort(a[left+1:])

	return a
}
