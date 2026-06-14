package pilas

import (
	"os"
	"reflect"
	"testing"
)

func TestPila(t *testing.T) {
	t.Run("Pila Vacía - Casos de Error", func(t *testing.T) {
		pila := &Stack[int]{}
		if !pila.IsEmpty() {
			t.Errorf("Se esperaba que la pila estuviera vacía")
		}
		_, err := pila.Pop()
		if err == nil {
			t.Errorf("Se esperaba un error al hacer Pop en una pila vacía")
		}
		_, err = pila.Peek()
		if err == nil {
			t.Errorf("Se esperaba un error al hacer Peek en una pila vacía")
		}
	})

	t.Run("Operaciones Normales (Push, Pop, Peek)", func(t *testing.T) {
		pila := &Stack[int]{}

		pila.Push(10)
		pila.Push(20)
		valor, err := pila.Peek()
		if err != nil || valor != 20 {
			t.Errorf("Peek falló: se esperaba 20, se obtuvo %v", valor)
		}
		valor, err = pila.Pop()
		if err != nil || valor != 20 {
			t.Errorf("Pop falló: se esperaba 20, se obtuvo %v", valor)
		}
		if pila.IsEmpty() {
			t.Errorf("La pila no debería estar vacía")
		}
	})
}
func TestCalcularStockSpan(t *testing.T) {
	pruebas := []struct {
		nombre   string
		precios  []float64
		esperado []int
	}{
		{
			nombre:   "Caso Normal - Fluctuación de precios",
			precios:  []float64{100, 80, 60, 70, 60, 75, 85},
			esperado: []int{1, 1, 1, 2, 1, 4, 6},
		},
		{
			nombre:   "Caso Límite - Estructura Vacía",
			precios:  []float64{},
			esperado: []int{},
		},
		{
			nombre:   "Caso Límite - Un solo elemento",
			precios:  []float64{150.5},
			esperado: []int{1},
		},
		{
			nombre:   "Tendencia Bajista Pura (Todos los span deben ser 1)",
			precios:  []float64{100, 90, 80, 70, 60},
			esperado: []int{1, 1, 1, 1, 1},
		},
		{
			nombre:   "Tendencia Alcista Pura",
			precios:  []float64{10, 20, 30, 40, 50},
			esperado: []int{1, 2, 3, 4, 5},
		},
	}

	for _, prueba := range pruebas {
		t.Run(prueba.nombre, func(t *testing.T) {
			resultado := CalcularStockSpan(prueba.precios)
			if !reflect.DeepEqual(resultado, prueba.esperado) {
				t.Errorf("Prueba '%s' falló. Esperado: %v, Obtenido: %v", prueba.nombre, prueba.esperado, resultado)
			}
		})
	}
}
func generarPrecios(n int) []float64 {
	precios := make([]float64, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			precios[i] = float64(100 - (i % 50))
		} else {
			precios[i] = float64(80 + (i % 20))
		}
	}
	return precios
}

func BenchmarkCalcularStockSpan1000(b *testing.B) {
	datos := generarPrecios(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalcularStockSpan(datos)
	}
}

func BenchmarkCalcularStockSpan10000(b *testing.B) {
	datos := generarPrecios(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalcularStockSpan(datos)
	}
}

func BenchmarkCalcularStockSpan100000(b *testing.B) {
	datos := generarPrecios(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalcularStockSpan(datos)
	}
}
func TestLeerPrecios(t *testing.T) {
	contenidoCSV := "Date,Open,High,Low,Close,Volume,OpenInt\n1997-05-16,1.97,1.98,1.71,1.73,14700000,0\n1997-05-19,1.76,1.76,1.62,1.71,6106800,0\n"
	rutaTemporal := "test_amzn.txt"

	err := os.WriteFile(rutaTemporal, []byte(contenidoCSV), 0644)
	if err != nil {
		t.Fatalf("No se pudo crear el archivo temporal de prueba")
	}
	defer os.Remove(rutaTemporal)
	registros, err := LeerPrecios(rutaTemporal)
	if err != nil {
		t.Errorf("No se esperaba error al leer, se obtuvo: %v", err)
	}
	if len(registros) != 2 {
		t.Errorf("Se esperaban 2 registros, se obtuvieron %d", len(registros))
	}
	if len(registros) > 0 && registros[0].Precio != 1.73 {
		t.Errorf("Precio incorrecto, esperado 1.73, obtenido %f", registros[0].Precio)
	}
	_, err = LeerPrecios("archivo_falso_que_no_existe.txt")
	if err == nil {
		t.Errorf("Se esperaba un error al intentar leer un archivo que no existe")
	}
}
