package entitys

import (
	customerros "golang/internal/custom_erros"
	"log"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	// ouvindo eventos de interrupção do sistema
	interrupt chan os.Signal

	// ouvindo se ouve algum erro ao completar a tarefa, caso nil, sucesso.
	complete chan error

	// ouvindo error de expiração de execução
	timeout <-chan time.Time

	// list de tarefas a serem executas
	tasks []func(int)
}

func NewRunner(duration time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(duration),
	}
}

// associa uma tarefa na fila de execução recebendo um id.
func (r *Runner) AddTask(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (runner *Runner) Start() error {
	// 'pluga o canal de erros de execução do sistema para o canal do runner.
	signal.Notify(runner.interrupt, os.Interrupt)

	go func() {
		runner.complete <- runner.run()
	}()

	select {
	case err := <-runner.complete:
		return err
	case <-runner.timeout:
		return customerros.NewTimeoutError()
	}
}

func (runner *Runner) run() error {
	// verifica se houve interrupção do sistema e percorre todas as tarefas executando-as.
	for id, task := range runner.tasks {
		if runner.isInterrupted() {
			return customerros.NewInterruptError()
		}

		task(id)
	}
	return nil
}

func (runner *Runner) isInterrupted() bool {
	select {
	// se chegar algum valor desse canal, eu paro o sistema com o erro que recebi.
	case <-runner.interrupt:
		signal.Stop(runner.interrupt)
		return true
	default:
		return false
	}
}

func (runner *Runner) CreateTaskForRunner() func(int) {
	return func(id int) {
		log.Printf("Processor - T#%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
