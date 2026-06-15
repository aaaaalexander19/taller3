package main

import (
	"math"
	"testing"
)

func esAVLBalanceado(t *testing.T, nodo *NodoAVL) bool {
	if nodo == nil {
		return true
	}
	balance := factorBalance(nodo)
	if math.Abs(float64(balance)) > 1 {
		t.Errorf("Árbol desbalanceado en el nodo con clave %f. Balance: %d", nodo.Clave, balance)
		return false
	}
	return esAVLBalanceado(t, nodo.Izq) && esAVLBalanceado(t, nodo.Der)
}

func TestCasosLimite(t *testing.T) {
	arbol := &ArbolAVL{}

	resultados := arbol.ConsultaRango(1990, 2000)
	if len(resultados) != 0 {
		t.Errorf("Se esperaba 0 resultados en árbol vacío, se obtuvieron %d", len(resultados))
	}

	arbol.Insertar(1995, Pelicula{Titulo: "Toy Story"})
	if arbol.NumNodos != 1 {
		t.Errorf("Se esperaba 1 nodo, se obtuvieron %d", arbol.NumNodos)
	}
	if Altura(arbol.Raiz) != 1 {
		t.Errorf("Se esperaba altura 1, se obtuvo %d", Altura(arbol.Raiz))
	}
}

func TestInsercionesOrdenadas(t *testing.T) {
	arbol := &ArbolAVL{}

	anios := []float64{1990, 1991, 1992, 1993, 1994, 1995, 1996}
	for _, anio := range anios {
		arbol.Insertar(anio, Pelicula{Titulo: "Peli de prueba"})
	}

	if !esAVLBalanceado(t, arbol.Raiz) {
		t.Errorf("El árbol no se balanceó correctamente tras inserciones ordenadas")
	}

	if arbol.Rotaciones == 0 {
		t.Errorf("Se esperaban rotaciones al insertar datos ordenados, pero no hubo ninguna")
	}
}

func TestConsultaRango(t *testing.T) {
	arbol := &ArbolAVL{}

	arbol.Insertar(1990, Pelicula{Titulo: "Goodfellas"})
	arbol.Insertar(1990, Pelicula{Titulo: "Home Alone"})
	arbol.Insertar(1994, Pelicula{Titulo: "Pulp Fiction"})
	arbol.Insertar(1999, Pelicula{Titulo: "The Matrix"})
	arbol.Insertar(2001, Pelicula{Titulo: "Shrek"})

	resultados := arbol.ConsultaRango(1990, 1995)
	if len(resultados) != 3 {
		t.Errorf("Se esperaban 3 películas en el rango, se obtuvieron %d", len(resultados))
	}

	resultadosFuera := arbol.ConsultaRango(2010, 2020)
	if len(resultadosFuera) != 0 {
		t.Errorf("Se esperaban 0 películas, se obtuvieron %d", len(resultadosFuera))
	}
}
func TestExtraerAnio(t *testing.T) {
	anio, err := extraerAnio("Toy Story (1995)")
	if err != nil || anio != 1995 {
		t.Errorf("Se esperaba 1995 sin error, se obtuvo: %v, error: %v", anio, err)
	}

	_, err2 := extraerAnio("Pelicula Rara Sin Fecha")
	if err2 == nil {
		t.Errorf("Se esperaba un error al no encontrar el año, pero no ocurrió")
	}

	anio3, _ := extraerAnio("The Matrix (1999)  ")
	if anio3 != 1999 {
		t.Errorf("Se esperaba 1999, se obtuvo: %v", anio3)
	}
}
func TestMainIntegracion(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("El programa principal falló con un panic: %v", r)
		}
	}()

	main()
}

func benchmarkConsultaRango(b *testing.B, n int) {
	arbol := &ArbolAVL{}
	for j := 0; j < n; j++ {
		arbol.Insertar(float64(j), Pelicula{Titulo: "Benchmark"})
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		arbol.ConsultaRango(float64(n/2), float64(n/2)+10)
	}
}

func BenchmarkConsulta1000(b *testing.B)    { benchmarkConsultaRango(b, 1000) }
func BenchmarkConsulta10000(b *testing.B)   { benchmarkConsultaRango(b, 10000) }
func BenchmarkConsulta100000(b *testing.B)  { benchmarkConsultaRango(b, 100000) }
func BenchmarkConsulta1000000(b *testing.B) { benchmarkConsultaRango(b, 1000000) }
