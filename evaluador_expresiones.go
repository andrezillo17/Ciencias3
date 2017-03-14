/*
	Gustavo Andres Naranjo - 20132020015
	David Morales - 20122020069
*/

package main

import(
	io "fmt"
	str "strings"
	conv "strconv"
)



type Nodo struct{//Se crean los nuevos tipo de datos propios del programa
	Valor int
	Nombre string
}

type Stack struct{
	nodos []*Nodo
	contador int
}


//-----------------Información del Nodo------------------
func (nodo *Nodo) String() string{
	return nodo.Nombre
}
//-------------------------------------------------------

//------------------Funciones del Stack------------------
func NuevoStack() *Stack {
	return &Stack{}
}

func (pila *Stack) Push(nodo *Nodo){
	pila.nodos = append(pila.nodos[:pila.contador], nodo)
	pila.contador++
}

func (pila *Stack) Pop() *Nodo{
	if pila.contador == 0{
		return nil
	}
	pila.contador--
	return pila.nodos[pila.contador]
}
//-------------------------------------------------------

func ResolverPila(pila *Stack) int{
	pilaAux := NuevoStack()
	rsta := 0
	for i:=0; i<len(pila.nodos); i++{		
		termino := pila.Pop()
		aux, err := conv.Atoi(termino.Nombre)
		if err != nil{
			switch termino.Nombre{
				case "+":					
					rsta = pilaAux.Pop().Valor + pilaAux.Pop().Valor
					pilaAux.Push(&Nodo{rsta,""})
				case "-":
					rsta = pilaAux.Pop().Valor - pilaAux.Pop().Valor
					pilaAux.Push(&Nodo{rsta,""})
				case "/":
					denominador := pilaAux.Pop().Valor
					if denominador != 0 {
						rsta = pilaAux.Pop().Valor / denominador						
					}else{
						rsta = pilaAux.Pop().Valor / 1
					}
					pilaAux.Push(&Nodo{rsta,""})
				case "*":
					rsta = pilaAux.Pop().Valor * pilaAux.Pop().Valor
					pilaAux.Push(&Nodo{rsta,""})
			}
		}else{
			pilaAux.Push(&Nodo{aux,""})
		}
	}
	return rsta
}

func main(){
	
	arbol1 := ":= + 5 3 x"
	arbol2 := ":= + - 5 2 4 y"
	arbol3 := ":= / x y z"
	pila1 := NuevoStack()
	pila2 := NuevoStack()
	pila3 := NuevoStack()
	array1 := str.Split(arbol1, " ")
	array2 := str.Split(arbol2, " ")
	array3 := str.Split(arbol3, " ")
	for i:=0; i<len(array1); i++{
		pila1.Push(&Nodo{i,array1[i]})
	}
	for i:=0; i<len(array2); i++{
		pila2.Push(&Nodo{i,array2[i]})
	}
	for i:=0; i<len(array3); i++{
		pila3.Push(&Nodo{i,array3[i]})
	}
	x := ResolverPila(pila1)
	y := ResolverPila(pila2)
	arr := [2] int {x,y}
	for i:=0; i<=len(arr); i++{
		pila3.Pop()
	}
	for i:=0; i<len(arr); i++{
		pila3.Push(&Nodo{i,conv.Itoa(arr[i])})
	}
	z  := ResolverPila(pila3)
	io.Println("X=",arr[0]," Y=",arr[1]," Z=",z)

}
