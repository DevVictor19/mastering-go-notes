package main

import "fmt"

func main() {
	ref1 := make([]int, 0, 3) // length: 0, cap: 3
	ref2 := ref1

	ref2 = append(ref2, 1, 2, 3) // não excedeu a capacidade (array interno mantem o mesmo)

	fmt.Printf("ref1 %p\n", ref1)
	fmt.Printf("ref2 %p\n", ref2)

	fmt.Println("ref1 slice", ref1) // ref1 ainda guarda uma slice len:0 e cap:3 (aparece vazio)
	fmt.Println("ref2 slice", ref2) // ref2 agora guarda uma slice len:3 e cap:3 (criada no append)

	// o array interno é o mesmo
	fmt.Println("ref1 internal", ref1[:3])
	fmt.Println("ref2 internal", ref2[:3])

	ref2 = append(ref2, 4) // fazendo o array interno mudar

	fmt.Printf("ref1 %p\n", ref1)
	fmt.Printf("ref2 %p\n", ref2) // referencia diferente

	// arrays internos diferentes
	fmt.Println("ref1 internal", ref1[:3]) // tentar acessar :4 vai dar panic
	fmt.Println("ref2 internal", ref2[:4])
}
