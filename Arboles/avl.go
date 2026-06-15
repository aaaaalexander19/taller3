package main

import (
	"math"
)

type Pelicula struct {
	ID      string
	Titulo  string
	Generos string
}

type NodoAVL struct {
	Clave float64
	Datos []Pelicula
	Alt   int
	Izq   *NodoAVL
	Der   *NodoAVL
}

type ArbolAVL struct {
	Raiz       *NodoAVL
	NumNodos   int
	Rotaciones int
}

func Altura(n *NodoAVL) int {
	if n == nil {
		return 0
	}
	return n.Alt
}

func factorBalance(n *NodoAVL) int {
	if n == nil {
		return 0
	}
	return Altura(n.Izq) - Altura(n.Der)
}

func (arbol *ArbolAVL) rotarDer(y *NodoAVL) *NodoAVL {
	arbol.Rotaciones++
	x := y.Izq
	T2 := x.Der

	x.Der = y
	y.Izq = T2

	y.Alt = int(math.Max(float64(Altura(y.Izq)), float64(Altura(y.Der)))) + 1
	x.Alt = int(math.Max(float64(Altura(x.Izq)), float64(Altura(x.Der)))) + 1

	return x
}

func (arbol *ArbolAVL) rotarIzq(x *NodoAVL) *NodoAVL {
	arbol.Rotaciones++
	y := x.Der
	T2 := y.Izq

	y.Izq = x
	x.Der = T2

	x.Alt = int(math.Max(float64(Altura(x.Izq)), float64(Altura(x.Der)))) + 1
	y.Alt = int(math.Max(float64(Altura(y.Izq)), float64(Altura(y.Der)))) + 1

	return y
}

func (arbol *ArbolAVL) Insertar(clave float64, dato Pelicula) {
	arbol.Raiz = arbol.insertarRecursivo(arbol.Raiz, clave, dato)
}

func (arbol *ArbolAVL) insertarRecursivo(nodo *NodoAVL, clave float64, dato Pelicula) *NodoAVL {
	if nodo == nil {
		arbol.NumNodos++ // Nuevo nodo único (nuevo año)
		return &NodoAVL{Clave: clave, Datos: []Pelicula{dato}, Alt: 1}
	}

	if clave < nodo.Clave {
		nodo.Izq = arbol.insertarRecursivo(nodo.Izq, clave, dato)
	} else if clave > nodo.Clave {
		nodo.Der = arbol.insertarRecursivo(nodo.Der, clave, dato)
	} else {
		// La clave (el año) ya existe, simplemente agregamos la película a la lista.
		nodo.Datos = append(nodo.Datos, dato)
		return nodo
	}

	nodo.Alt = int(math.Max(float64(Altura(nodo.Izq)), float64(Altura(nodo.Der)))) + 1
	balance := factorBalance(nodo)

	if balance > 1 && clave < nodo.Izq.Clave {
		return arbol.rotarDer(nodo)
	}
	if balance < -1 && clave > nodo.Der.Clave {
		return arbol.rotarIzq(nodo)
	}
	if balance > 1 && clave > nodo.Izq.Clave {
		nodo.Izq = arbol.rotarIzq(nodo.Izq)
		return arbol.rotarDer(nodo)
	}
	if balance < -1 && clave < nodo.Der.Clave {
		nodo.Der = arbol.rotarDer(nodo.Der)
		return arbol.rotarIzq(nodo)
	}

	return nodo
}

func (arbol *ArbolAVL) ConsultaRango(a, b float64) []Pelicula {
	return arbol.consultaRecursiva(arbol.Raiz, a, b)
}

func (arbol *ArbolAVL) consultaRecursiva(nodo *NodoAVL, a, b float64) []Pelicula {
	var resultados []Pelicula
	if nodo == nil {
		return resultados
	}

	if a < nodo.Clave {
		resultados = append(resultados, arbol.consultaRecursiva(nodo.Izq, a, b)...)
	}
	if a <= nodo.Clave && nodo.Clave <= b {
		resultados = append(resultados, nodo.Datos...)
	}
	if b > nodo.Clave {
		resultados = append(resultados, arbol.consultaRecursiva(nodo.Der, a, b)...)
	}

	return resultados
}
