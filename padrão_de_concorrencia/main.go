package main

import (
	customerros "golang/internal/custom_erros"
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
	timeout := time.Duration(3 * time.Second)

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
		switch err.(type) {
		case *customerros.InterruptError:
			log.Fatalln("Process has interrupted.")
			os.Exit(1)
		case *customerros.TimeoutError:
			log.Fatalln("Process timeout of execution.")
			os.Exit(2)
		}
	}

	log.Println("Process end. <3")
}
