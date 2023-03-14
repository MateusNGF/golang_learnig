package main

import (
	"golang/internal/entitys"
	"log"
	"os"
	"time"
)

func main() {
	RunRunner()
}

func RunRunner() {

	// definimos o tempo de execução maximo do programa.
	timeout := time.Duration(7 * time.Second)

	log.Println("==> Starting work")

	// criamos um runner e configuramos com o timeout
	runner := entitys.NewRunner(timeout)

	var tasks []func(int)

	// para facilitar os testes, fiz um loop para adicionar tarefas.
	for i := 0; i < 10; i++ {
		tasks = append(tasks, runner.CreateTaskForRunner())
	}

	// defino as tarefas para o runner executar.
	runner.AddTask(tasks...)

	// inicio a execução do programa e fico observando erros.
	if err := runner.Start(); err != nil {
		log.Printf("Terminating due to " + err.Error())
		os.Exit(1)
	}

	log.Println("Process end. <3")
}
