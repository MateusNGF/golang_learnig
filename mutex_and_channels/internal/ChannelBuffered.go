package internal

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/*
	Este exemplo de programa mostra como usar um canal com buffer
  para trabalhar com varias tarefas usando um numero predefinido de gorutines.

  Um Channel Buffered é um canal com capacidade para armazenar um ou mais valores antes que eles
  sejam recebidos. Um canal sem buffer garante que a troca entre duas goroutines seja efetuadas no
  momento em que o envio e a recepção ocorrem. UM CANAL COM BUFFER NÃO OFERECE ESSA GARANTIA.

*/

const (
	numberGorutinesChannelBuffered = 10
	taskLoadChannelBuffered        = 1e4
)

var wgChannelBuffered sync.WaitGroup

// init é chamado antes da execução de todas as chamadas
func init() {

	// gera a semente do gerador de numeros aleatorios.
	rand.Seed(time.Now().Unix())
}

func ChannelBuffered() {
	runtime.GOMAXPROCS(1)
	// cria um canal 'bufferizado'.
	task := make(chan string, taskLoadChannelBuffered)

	wgChannelBuffered.Add(numberGorutinesChannelBuffered)

	for i := 0; i < numberGorutinesChannelBuffered; i++ {
		go Work(task, i)
	}

	// popula o canal com algumas tarefas a serem feitas
	for i := 0; i < taskLoadChannelBuffered; i++ {
		task <- fmt.Sprintf("T-%d", i)
	}

	close(task)
	wgChannelBuffered.Wait()
}

func Work(task chan string, work int) {
	defer wgChannelBuffered.Done()

	for {
		task, ok := <-task
		if !ok {
			// esse trecho verifica se o canal foi fechado ou/e vazio
			fmt.Printf("Worker: %d: Shutting Down\n", work)
			return
		}

		// Apresenta a inicialização do worker
		fmt.Printf("Worker: %d: Started %s\n", work, task)

		// simula um tempo aleatório de execução
		sleep := rand.Intn(100)
		time.Sleep(time.Duration(sleep) * time.Microsecond)

		// finaliza a execução da task
		fmt.Printf("Worker: %d: Completed %s\n", work, task)
	}
}
