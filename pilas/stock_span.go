package pilas

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
)

type Stack[T any] struct {
	elementos []T
}

func (s *Stack[T]) Push(v T) {
	s.elementos = append(s.elementos, v)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var Zero T
		return Zero, errors.New("La pila esta vacia")
	}
	aux := s.elementos[len(s.elementos)-1]
	s.elementos = s.elementos[:len(s.elementos)-1]
	return aux, nil
}
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var Zero T
		return Zero, errors.New("La pila esta vacia")
	}
	aux := s.elementos[len(s.elementos)-1]
	return aux, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elementos) == 0
}

func CalcularStockSpan(precios []float64) []int {
	spans := make([]int, len(precios))
	Stack := &Stack[int]{}
	for i := 0; i < len(precios); i++ {
		for !Stack.IsEmpty() {
			indiceDelTope, _ := Stack.Peek()
			if precios[indiceDelTope] <= precios[i] {
				Stack.Pop()
			} else {
				break
			}
		}
		if Stack.IsEmpty() {
			spans[i] = i + 1
		} else {
			indiceDelTope, _ := Stack.Peek()
			spans[i] = i - indiceDelTope
		}
		Stack.Push(i)
	}
	return spans
}

type Registro struct {
	Fecha  string
	Precio float64
}

func LeerPrecios(ruta string) ([]Registro, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	lector := csv.NewReader(archivo)
	lineas, err := lector.ReadAll()
	if err != nil {
		return nil, err
	}
	var listaRegistros []Registro
	for i := 1; i < len(lineas); i++ {
		fecha := lineas[i][0]
		precioCierre, err := strconv.ParseFloat(lineas[i][4], 64)
		if err == nil {
			reg := Registro{
				Fecha:  fecha,
				Precio: precioCierre,
			}
			listaRegistros = append(listaRegistros, reg)
		}
	}
	return listaRegistros, nil
}
