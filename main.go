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

	var para bool

	arquivo, err := os.Open("huge.log")

	if err != nil {
		log.Fatalf("ERRO: %v", err)
	}

	defer arquivo.Close()

	r := bufio.NewReader(arquivo)

	line, prefix, erro := r.ReadLine()
	if erro != nil {
	}
	log.Println(prefix)

	var listanova []string

	for para == false {

		listanova = append(listanova, string(line))

		line, prefix, erro = r.ReadLine()

		if erro != nil {
			para = true
		}

	}

	log.Println("antes qsort")
	listanova = qsort(listanova)
	log.Println("depois qsort")
	log.Println(len(listanova))

	errArq := escreverTexto(listanova, "huge_orded.log")
	if errArq != nil {
		log.Fatalf("Erro:", errArq)
	} else {
		setmd5()
	}
}

func mainzz() {

	var para bool

	arquivo, err := os.Open("huge.log")

	if err != nil {
		log.Fatalf("ERRO: %v", err)
	}

	defer arquivo.Close()

	r := bufio.NewReader(arquivo)

	line, prefix, erro := r.ReadLine()
	if erro != nil {
	}
	log.Println(prefix)

	var listanova []string

	for para == false {

		listanova = sortByLine(listanova, string(line))
		log.Println(len(listanova))

		line, prefix, erro = r.ReadLine()

		if erro != nil {
			para = true
		}

	}

	errArq := escreverTexto(listanova, "huge_orded.log")
	if errArq != nil {
		log.Fatalf("Erro:", errArq)
	} else {
		setmd5()
	}

	/**TESTE

	// testando
	var listateste []string

	listateste = append(listateste, "ab")
	listateste = append(listateste, "zx")
	listateste = append(listateste, "rt")
	listateste = append(listateste, "iu")
	listateste = append(listateste, "aa")
	listateste = append(listateste, "ll")
	listateste = append(listateste, "6y")
	listateste = append(listateste, "u8")
	listateste = append(listateste, "0p")
	listateste = append(listateste, "uu")
	listateste = append(listateste, "00")
	listateste = append(listateste, "fg")
	listateste = append(listateste, "ra")
	listateste = append(listateste, "zz")
	listateste = append(listateste, "cp")
	listateste = append(listateste, "ep")
	listateste = append(listateste, "fc")
	listateste = append(listateste, "bp")

	var listanova []string

	for index := 0; index < len(listateste); index++ {
		listanova = sortByLine(listanova, listateste[index])
		log.Println(listanova, len(listanova))
	}

	errArq := escreverTexto(listanova, "huge_orded.log")
	if errArq != nil {
		log.Fatalf("Erro:", errArq)
	} else {
		setmd5()
	}
	 FIM TESTE **/

}

func sortByLine(lista []string, linha string) []string {

	if len(lista) == 0 {
		lista = append(lista, linha)
		return lista
	}

	if len(lista) == 1 {
		if lista[0] > linha {
			lista = append(lista, lista[0])
			lista[0] = linha
		} else {
			lista = append(lista, linha)
		}
		return lista
	}

	var linhaAux string
	var subs bool
	var pos int
	var posAux int
	var encontrou bool
	var maior bool
	var menor bool

	pos = len(lista) / 2

	for encontrou == false {
		//log.Println(pos, menor, maior)
		if lista[pos] > linha {
			maior = true
			if menor == false {
				posAux = pos
				pos = pos / 2
			} else {
				encontrou = true
			}
		}

		if lista[pos] < linha {
			menor = true
			if maior == false {
				posAux = pos
				pos = pos + (len(lista) / 4)
				if pos > len(lista)-1 {
					pos = len(lista) - 1
					encontrou = true
				}
				if pos == 1 {
					pos = 0
				}
			} else {
				encontrou = true
			}
		}
		if pos == 0 {
			encontrou = true
		}

	}

	if posAux < pos {
		pos = posAux
	}

	for index := pos; index < len(lista); index++ {
		if subs == false {
			if linha < lista[index] {
				linhaAux = lista[index]
				lista[index] = linha
				linha = linhaAux
				subs = true
			}
		} else {
			linhaAux = lista[index]
			lista[index] = linha
			linha = linhaAux
		}
	}

	lista = append(lista, linha)

	return lista
}

func escreverTexto(linhas []string, caminhoDoArquivo string) error {
	// Cria o arquivo de texto
	arquivo, err := os.Create(caminhoDoArquivo)
	// Caso tenha encontrado algum erro retornar ele
	if err != nil {
		return err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um escritor responsavel por escrever cada linha do slice no arquivo de texto
	escritor := bufio.NewWriter(arquivo)
	for _, linha := range linhas {
		fmt.Fprintln(escritor, linha)
	}

	// Caso a funcao flush retorne um erro ele sera retornado aqui tambem
	return escritor.Flush()
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

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	qsort(a[:left])
	qsort(a[left+1:])

	return a
}
