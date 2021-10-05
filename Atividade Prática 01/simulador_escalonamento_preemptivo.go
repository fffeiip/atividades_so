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
	last_iteration      time.Time
	wait_time           []time.Duration
	finalizado_em       time.Time
}

func main() {

	queue := list.New()
	finished_processes := list.New()
	var proc Process
	var wait_time time.Duration
	var sum int
	var nome string
	var tempo_cpu int
	var quantum int
	var total_processes int

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

	for e := queue.Front(); e != nil; e = e.Next() {
		proc = e.Value.(Process)
		fmt.Printf("\n====\tProcesso %s\t====", proc.nome)

		if proc.last_iteration.IsZero() {
			wait_time = time.Now().Sub(start)
		} else {
			wait_time = time.Now().Sub(proc.last_iteration)
		}
		proc.wait_time = append(proc.wait_time, wait_time)
		time.Sleep(time.Duration(quantum) * time.Millisecond)
		proc.tempo_cpu_decorrido = proc.tempo_cpu_decorrido + quantum
		if proc.tempo_cpu > quantum && proc.tempo_cpu_decorrido < proc.tempo_cpu {
			fmt.Printf("\nProcesso voltou pro final da fila")
			fmt.Printf("\nTempo de cpu nescessário %dms\nTempo de processamento decorrido: %dms", proc.tempo_cpu, proc.tempo_cpu_decorrido)
			proc.last_iteration = time.Now()
			queue.PushBack(proc)
		} else {
			proc.finalizado_em = time.Now()
			finished_processes.PushBack(proc)
			fmt.Printf("\nProcesso %d finalizado", proc.id)
			fmt.Printf("\nTempo de cpu necessário %dms", proc.tempo_cpu)
			fmt.Println("\nTurnaround: ", proc.finalizado_em.Sub(start))
		}
		fmt.Println("\n============================")
	}

	fmt.Println("\n************Relatorio Final************")

	for e := finished_processes.Front(); e != nil; e = e.Next() {
		proc = e.Value.(Process)
		sum = 0
		fmt.Printf("\n======		Processo %s		======", proc.nome)
		fmt.Println("\nTempo de turnaround: ", proc.finalizado_em.Sub(start))
		for i := 0; i < len(proc.wait_time); i++ {
			sum += int(proc.wait_time[i])
		}
		fmt.Println("\nTempo médio de espera: ", time.Duration(sum/len(proc.wait_time)))
		fmt.Println("\n===================================")
	}
}
