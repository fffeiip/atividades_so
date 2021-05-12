package main

import (
	"container/list"
	"fmt"
	"time"
)

type Process struct {
	id                  int
	nome                string
	tempo_cpu           int
	tempo_cpu_decorrido int
	finalizado_em       time.Time
}

func main() {

	queue := list.New()
	var proc Process
	var nome string
	var tempo_cpu int
	var quantum int
	var total_processes int
	var e *list.Element
	var next *list.Element

	fmt.Println("Exemplo de algorítimo Round-Robin")
	fmt.Println("Favor informar o total de processos:")
	fmt.Scanln(&total_processes)

	fmt.Println("Favor informar o tempo de quantum:")
	fmt.Scanln(&quantum)

	for i := 1; i <= total_processes; i++ {
		fmt.Printf("=====\tProcesso %d\t=====\nInforme o nome do processo\n", i)
		fmt.Scanln(&nome)
		fmt.Println("\nTempo de cpu:")
		fmt.Scanln(&tempo_cpu)
		proc.id = i
		proc.nome = nome
		proc.tempo_cpu = tempo_cpu
		queue.PushBack(proc)
	}

	start := time.Now()
	e = queue.Front()
	for {
		proc = e.Value.(Process)
		next = e.Next()
		fmt.Printf("\n====\tProcesso %s\t====", proc.nome)
		time.Sleep(time.Duration(quantum) * time.Millisecond)
		proc.tempo_cpu_decorrido = proc.tempo_cpu_decorrido + quantum
		if proc.tempo_cpu > quantum && proc.tempo_cpu_decorrido < proc.tempo_cpu {
			fmt.Printf("\nProcesso voltou pro final da fila")
			fmt.Printf("\nTempo de cpu nescessário %dms\nTempo de processamento decorrido: %dms", proc.tempo_cpu, proc.tempo_cpu_decorrido)
			queue.PushBack(proc)
		} else {
			proc.finalizado_em = time.Now()
			fmt.Printf("\nProcesso %d finalizado", proc.id)
			fmt.Println("\nTempo Total: ", proc.finalizado_em.Sub(start))
		}
		fmt.Println("\n============================")
		queue.Remove(e)

		if queue.Len() == 0 {
			break
		}
		if next == nil {
			e = queue.Front()
		} else {
			e = next
		}
	}
}
