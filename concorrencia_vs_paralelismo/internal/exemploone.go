package internal

import (
	"fmt"
	"runtime"
	"sync"
)

var wg1 sync.WaitGroup

func ExemploOne() {

	/**
	 ¹Limitando o numero de processos conseguimos ver que o runtime não vai fazer o swap pois a primeira
	goroutine não esta, digamos, sendo um processo bloqueante, ela esta constantemente respondendo.
	 No entando, se adicionarmos um Sleep depois da exibição do caracter veremos o runtime gerenciando e fazendo swap
	entre as gorutines. Esse gerenciamente(swap) é o que chamamos de Concorrencia. ²No entanto, se adicionarmos mais
	processadores fisicos, veremos outro conceito bem interessante que chamamos de Paralelismo, que nada mais é
	a capacidade de execução de muitas tarefas ao mesmo tempo. Note que Paralelismo é a execução de um ou mais tarefas
	em processos diferentes, e Conconrrencia é o gerenciamento das tarefas em um unico processo.

	**/
	// runtime.GOMAXPROCS(1) //¹
	runtime.GOMAXPROCS(2) //²

	wg1.Add(2)

	go abcDario(true)
	go abcDario(false)

	fmt.Println("Wait finish.")
	wg1.Wait()
	fmt.Println("Program finish.")
}

func abcDario(uppercase bool) {
	defer wg1.Done()
	start := 'A'

	if uppercase {
		start = 'a'
	}

	for i := 0; i < 3; i++ {
		for char := start; char < start+26; char++ {
			fmt.Printf("%c", char)
		}
	}
	fmt.Println('\n')
}
