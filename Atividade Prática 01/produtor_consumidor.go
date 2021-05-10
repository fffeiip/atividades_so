package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var contador_arroz = 10
var contador_feijao = 10
var mu sync.Mutex
var wg sync.WaitGroup

func ConsomeArroz() {
	mu.Lock()
	x := contador_arroz
	if x < 1 {
		fmt.Printf("Estoque de arroz: %d. Não foi possível consumir mais.", x)
	} else {
		runtime.Gosched()
		x--
		contador_arroz = x
		fmt.Printf("Consumiu uma unidade de arroz.\nNovo estoque de arroz: %d", x)
	}
	mu.Unlock()
	fmt.Println("\nRoutine do arroz dormindo por 3s")
	fmt.Println("\n\n============================================")
	time.Sleep(3 * time.Second)
}

func ConsomeFeijao() {
	mu.Lock()
	y := contador_feijao
	if y < 1 {
		fmt.Printf("Estoque de feijão: %d. Não foi possível consumir mais.", y)
	} else {
		runtime.Gosched()
		y--
		contador_feijao = y
		fmt.Printf("Consumiu uma unidade de feijão.\nNovo estoque de feijão: %d.", y)
	}
	mu.Unlock()
	fmt.Println("\nRoutine do feijão dormindo por 3s")
	fmt.Println("\n\n============================================")
	time.Sleep(3 * time.Second)
}

func ProduzFeijao() {
	mu.Lock()
	y := contador_feijao
	runtime.Gosched()
	fmt.Println("\n\n============\tProdutor\t=============")
	if y <= 0 {
		fmt.Println("\nFeijão em estoque: 0.\nNão será mais necessária a produção.")
	} else {
		y++
		contador_feijao = y
		fmt.Printf("\nProduziu mais uma unidade de feijão.\nNovo estoque de feijão: %d", y)
	}
	fmt.Println("\n\n=============================================")
	mu.Unlock()
	time.Sleep(3 * time.Second)
}

func ProduzArroz() {
	mu.Lock()
	x := contador_arroz
	runtime.Gosched()
	fmt.Println("\n\n============\tProdutor\t=============")

	if x > 0 {
		x++
		contador_arroz = x
		fmt.Printf("\nProduziu mais uma unidade de arroz.\nNovo estoque de arroz: %d", x)
	} else {
		fmt.Println("\nArroz em estoque: 0.\nNão será mais necessária a produção.")
	}
	fmt.Println("\n\n=============================================")

	mu.Unlock()
	time.Sleep(3 * time.Second)
}

func Consumidor(j int) {
	var wgConsumidor sync.WaitGroup
	wgConsumidor.Add(2)

	go func() {
		for {
			fmt.Printf("\n\n============\tConsumidor %d\t============\n", j)
			fmt.Println("Inicio routine de arroz")
			ConsomeArroz()
			mu.Lock()
			x := contador_arroz
			if x < 1 {
				mu.Unlock()
				break
			}
			mu.Unlock()
		}
		wgConsumidor.Done()
	}()

	go func() {
		for {
			fmt.Printf("\n\n============\tConsumidor %d\t============\n", j)
			fmt.Println("Inicio routine de feijão")
			ConsomeFeijao()
			mu.Lock()
			y := contador_feijao
			if y < 1 {
				mu.Unlock()
				break
			}
			mu.Unlock()
		}
		wgConsumidor.Done()
	}()

	wgConsumidor.Wait()
	wg.Done()
}

func Produtor() {
	var wgProdutor sync.WaitGroup
	wgProdutor.Add(2)
	go func() {
		for {
			ProduzArroz()
			mu.Lock()
			x := contador_arroz
			if x < 1 {
				mu.Unlock()
				break
			}
			mu.Unlock()
		}
		wgProdutor.Done()
	}()

	go func() {
		for {
			ProduzFeijao()
			mu.Lock()
			y := contador_feijao
			if y < 1 {
				mu.Unlock()
				break
			}
			mu.Unlock()
		}
		wgProdutor.Done()
	}()
	wgProdutor.Wait()
	wg.Done()
}

func main() {
	fmt.Println("Estoque inicial:\nArroz: ", contador_arroz, "\tFeijão", contador_feijao)
	totaldegoroutines := 6
	wg.Add(totaldegoroutines)
	go Produtor()
	go Consumidor(1)
	go Consumidor(2)
	go Consumidor(3)
	go Consumidor(4)
	go Consumidor(5)
	wg.Wait()
	fmt.Println("Estoque Final:\nArroz: ", contador_arroz, "\tFeijão", contador_feijao)
}
