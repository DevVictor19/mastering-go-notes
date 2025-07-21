package main

import "fmt"

func main() {
	a := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println("a", a)

	s := a[0:1] // outra slice que aponta para o mesmo array interno
	fmt.Println("s", s)
	fmt.Println("s cap", cap(s))

	s[0] = "1"
	s = append(s, "2")
	s = append(s, "3")
	s = append(s, "4")
	s = append(s, "5")
	// fazendo S mudar a capacidade e quebrando o vinculo com o array principal de a

	// aqui foi criado outro array para S
	s[0] = "No modifications"
	fmt.Println("a", a)

	// agora a e s possuem arrays diferentes por baixo dos panos (referencia diferente)
	fmt.Println("s", s)
	change(s)
	fmt.Println("a", a)
}

func change(s []string) {
	s[0] = "Change_function"
}
