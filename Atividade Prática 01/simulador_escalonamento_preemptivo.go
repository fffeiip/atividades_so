package main

import (
	"container/list"
	"fmt"
	"time"
)

type Process struct {
	id                  int
	nome                string
	prioridade          int
	tempo_cpu           int
	tempo_cpu_decorrido int
	finalizado_em       time.Time
}

func main() {

	queue := list.New()
	processes_ready := list.New()
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
		if proc.tempo_cpu > quantum && proc.tempo_cpu_decorrido < quantum {
			fmt.Printf("Processo %d voltou pro final da fila\n", proc.id)
			proc.tempo_cpu_decorrido = proc.tempo_cpu_decorrido + quantum
			fmt.Println("Tempo restante para finalizar: ", proc.tempo_cpu-proc.tempo_cpu_decorrido)
			queue.PushBack(proc)
		} else {
			fmt.Printf("Processo %d finalizado\n", proc.id)
			proc.tempo_cpu_decorrido = proc.tempo_cpu_decorrido + proc.tempo_cpu
			processes_ready.PushBack(proc)
		}
		queue.Remove(e)
		if next == nil && queue.Len() == 0 {
			break
		}
		e = next
	}
	for element := processes_ready.Front(); element != nil; element = element.Next() {
		proc = element.Value.(Process)
		finalizado := proc.finalizado_em.Sub(start)
		fmt.Printf("Processo %d - %s.\n", proc.id, proc.nome)
		fmt.Println("Tempo decorrido: ", finalizado)
	}

}
