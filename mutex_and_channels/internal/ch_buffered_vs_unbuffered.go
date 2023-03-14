package internal

import (
	"fmt"
	"time"
)

func ChannelBufferedVSChannelUnbuffered() {
	// showExemploUnbuffered()
	showExemploBuffered()
}

func showExemploUnbuffered() {
	channelUnBuffered := make(chan string)
	fmt.Printf("cap(channelBuffered_Vs): %v\n", cap(channelUnBuffered))

	// ² se tentarmos habilitar o lado Escritor antes dará error. ³
	// channelUnBuffered <- "A_"

	go func() {
		// ¹ aqui 'inicializa' o channel o lado 'Ouvinte', portanto habilitando o lado 'Escritor' ²
		for x := range channelUnBuffered {
			fmt.Printf("x: %v\n", x)
		}
	}()

	/*
		³ Isso define o comportamento do canal não bufferizado, ele garante que para cada escritor tenha um leitor.
	*/

	channelUnBuffered <- "B"
	close(channelUnBuffered)
	time.Sleep(2 * time.Second)
}

func showExemploBuffered() {
	channelBuffered := make(chan string, 2)
	fmt.Printf("cap(channelBuffered_Vs): %v\n", cap(channelBuffered))

	/*
			Quando definimos uma capacidade pre determinada, conseguimos escrever/inserir valores
		  sem precisar que tenhamos alguem lendo/recebendo os mesmos. Contudo, você só conseguirá inserir
		  a quantidade da capacidade, DEPOIS desse valor cairá na regra de canal não bufferizado.
		  	Analogamente, pense como um cano, que tenha um determinado raio e um determinado comprimeto,
		  com isso você consegue calcular seu volume, se tamparmos um dos lados só consigos encher, sem perdas,
		  somente até o volume desse cano. No entanto, para não termos perdas, precisamos dar VASÃO CONSTANTE.
	*/

	channelBuffered <- "COOMEÇO"
	channelBuffered <- "MEIO"

	go func() {
		for x := range channelBuffered {
			fmt.Println("x:", x)
		}
	}()

	channelBuffered <- "FIM"
	close(channelBuffered)
	time.Sleep(2 * time.Second)
}
