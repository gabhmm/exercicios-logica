package main

import "fmt"

//exercicio de interface e polimorfismo com structs


type pessoa struct {
	nome    string
	idade   int
	salario float64
}

type medico struct {
	pessoa
	cirurgias int
}

type enfermeiro struct {
	pessoa
	injeções int
}

type profissao interface {
	cumprimento()
}

func (m medico) cumprimento() {
	fmt.Printf("Eu trabalho como médico(a)!! Meu nome é %v, tenho %v anos, recebo %v e já fiz %v cirurgias\n", m.nome, m.idade, m.salario, m.cirurgias)
}

func (e enfermeiro) cumprimento() {
	fmt.Printf("Eu trabalho como enfermeiro(a)!! Meu nome é %v, tenho %v anos, recebo %v e já apliquei %v injeções", e.nome, e.idade, e.salario, e.injeções)
}

func gente(p profissao) {
	p.cumprimento()
}

func main() {
	m := medico{
		pessoa: pessoa{
			nome:    "Fernando",
			idade:   33,
			salario: 10500,
		},
		cirurgias: 23,
	}
	e := enfermeiro{
		pessoa: pessoa{
			nome:    "Marcia",
			idade:   26,
			salario: 7000,
		},
		injeções: 450,
	}
	gente(m)
	gente(e)
}
