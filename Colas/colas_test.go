package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// --- Pruebas para la estructura Cola ---

func TestCola_Enqueue(t *testing.T) {
	c := &Cola{}
	c.Enqueue(10)
	if c.size != 1 {
		t.Errorf("Enqueue() falló, tamaño esperado 1, pero fue %d", c.size)
	}
	if c.head.ts != 10 {
		t.Errorf("Enqueue() falló, valor de head esperado 10, pero fue %d", c.head.ts)
	}
	if c.tail.ts != 10 {
		t.Errorf("Enqueue() falló, valor de tail esperado 10, pero fue %d", c.tail.ts)
	}

	c.Enqueue(20)
	if c.size != 2 {
		t.Errorf("Enqueue() falló, tamaño esperado 2, pero fue %d", c.size)
	}
	if c.tail.ts != 20 {
		t.Errorf("Enqueue() falló, valor de tail esperado 20, pero fue %d", c.tail.ts)
	}
}

func TestCola_Dequeue(t *testing.T) {
	c := &Cola{}
	c.Enqueue(10)
	c.Enqueue(20)

	val, ok := c.Dequeue()
	if !ok || val != 10 {
		t.Errorf("Dequeue() falló, se esperaba valor 10 y ok=true, pero fue %d y %v", val, ok)
	}
	if c.size != 1 {
		t.Errorf("Dequeue() falló, tamaño esperado 1, pero fue %d", c.size)
	}
	if c.head.ts != 20 {
		t.Errorf("Dequeue() falló, head esperado con valor 20, pero fue %d", c.head.ts)
	}

	val, ok = c.Dequeue()
	if !ok || val != 20 {
		t.Errorf("Dequeue() falló, se esperaba valor 20 y ok=true, pero fue %d y %v", val, ok)
	}
	if c.size != 0 {
		t.Errorf("Dequeue() falló, tamaño esperado 0, pero fue %d", c.size)
	}
	if c.head != nil || c.tail != nil {
		t.Errorf("Dequeue() falló, head y tail deberían ser nil en una cola vacía")
	}
}

func TestCola_Vacia(t *testing.T) {
	c := &Cola{}

	// Probar Front en cola vacía
	_, ok := c.Front()
	if ok {
		t.Error("Front() en cola vacía debería retornar ok=false")
	}

	// Probar Dequeue en cola vacía
	_, ok = c.Dequeue()
	if ok {
		t.Error("Dequeue() en cola vacía debería retornar ok=false")
	}

	// Probar Len en cola vacía
	if c.Len() != 0 {
		t.Errorf("Len() en cola vacía debería ser 0, pero fue %d", c.Len())
	}
}

func TestCola_Len(t *testing.T) {
	c := &Cola{}
	if c.Len() != 0 {
		t.Errorf("Len() inicial debería ser 0, pero fue %d", c.Len())
	}
	c.Enqueue(1)
	if c.Len() != 1 {
		t.Errorf("Len() después de 1 Enqueue debería ser 1, pero fue %d", c.Len())
	}
	c.Enqueue(2)
	if c.Len() != 2 {
		t.Errorf("Len() después de 2 Enqueue debería ser 2, pero fue %d", c.Len())
	}
	c.Dequeue()
	if c.Len() != 1 {
		t.Errorf("Len() después de 1 Dequeue debería ser 1, pero fue %d", c.Len())
	}
	c.Dequeue()
	if c.Len() != 0 {
		t.Errorf("Len() después de vaciar la cola debería ser 0, pero fue %d", c.Len())
	}
}

func TestParsearLinea(t *testing.T) {
	// Caso normal
	linea := `192.168.1.1 - - [08/Jun/2026:10:00:00 -0500] "GET / HTTP/1.1" 200 1234`
	ip, ts, err := ParsearLinea(linea)
	if err != nil {
		t.Fatalf("ParsearLinea() falló con error inesperado: %v", err)
	}
	if ip != "192.168.1.1" {
		t.Errorf("IP esperada '192.168.1.1', pero se obtuvo '%s'", ip)
	}

	// Para hacer el test determinístico, calculamos el timestamp esperado de la misma forma que la función lo hace.
	layout := "[02/Jan/2006:15:04:05"
	expectedTime, _ := time.Parse(layout, "[08/Jun/2026:10:00:00")
	expectedTs := expectedTime.Unix()

	if ts != expectedTs {
		t.Errorf("Timestamp esperado %d, pero se obtuvo %d", expectedTs, ts)
	}

	// Caso de error: línea mal formada
	_, _, err = ParsearLinea("linea corta")
	if err == nil {
		t.Error("ParsearLinea() debería haber fallado con una línea mal formada, pero no lo hizo")
	}

	// Caso de error: fecha inválida
	lineaFechaInvalida := `192.168.1.1 - - [99/Invalid/9999:10:00:00 -0500] "GET /"`
	_, _, err = ParsearLinea(lineaFechaInvalida)
	if err == nil {
		t.Error("ParsearLinea() debería haber fallado con una fecha inválida, pero no lo hizo")
	}
}

func TestPermitirPeticion(t *testing.T) {
	colas := make(map[string]*Cola)
	ip := "1.2.3.4"
	M := 2         // 2 peticiones
	T := int64(10) // en 10 segundos

	// 1. Aceptar primera petición
	ts1 := time.Now().Unix()
	if !PermitirPeticion(colas, ip, ts1, M, T) {
		t.Error("Petición 1 debería ser aceptada")
	}

	// 2. Aceptar segunda petición dentro de la ventana
	ts2 := ts1 + 1
	if !PermitirPeticion(colas, ip, ts2, M, T) {
		t.Error("Petición 2 debería ser aceptada")
	}

	// 3. Rechazar tercera petición (límite alcanzado)
	ts3 := ts2 + 1
	if PermitirPeticion(colas, ip, ts3, M, T) {
		t.Error("Petición 3 debería ser rechazada")
	}

	// 4. Aceptar petición después de que la ventana de tiempo haya pasado
	// Esto prueba que la cola descarta marcas vencidas.
	ts4 := ts1 + T + 1 // ts1 ahora está fuera de la ventana [ts4 - T, ts4]
	if !PermitirPeticion(colas, ip, ts4, M, T) {
		t.Error("Petición 4 debería ser aceptada porque las viejas expiraron")
	}
	// Después de la limpieza y el nuevo enqueue, la cola debería tener 2 elementos (ts2, ts4)
	if colas[ip].Len() != 2 {
		t.Errorf("Longitud de cola esperada 2 tras expiración y nueva petición, pero fue %d", colas[ip].Len())
	}
}

// --- Benchmarks ---

func BenchmarkCola_Enqueue(b *testing.B) {
	c := &Cola{}
	for i := 0; i < b.N; i++ {
		c.Enqueue(int64(i))
	}
}

func BenchmarkCola_Dequeue(b *testing.B) {
	c := &Cola{}
	// Pre-llenar la cola para que Dequeue tenga algo que hacer
	for i := 0; i < b.N; i++ {
		c.Enqueue(int64(i))
	}
	b.ResetTimer() // No medir el llenado
	for i := 0; i < b.N; i++ {
		c.Dequeue()
	}
}

func BenchmarkCola_Enqueue_Sizes(b *testing.B) {
	sizes := []int{1000, 10000, 100000, 1000000}
	for _, n := range sizes {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c := &Cola{}
				for j := 0; j < n; j++ {
					c.Enqueue(int64(j))
				}
			}
		})
	}
}

func BenchmarkCola_Dequeue_Sizes(b *testing.B) {
	sizes := []int{1000, 10000, 100000, 1000000}
	for _, n := range sizes {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				c := &Cola{}
				for j := 0; j < n; j++ {
					c.Enqueue(int64(j))
				}
				b.StartTimer()
				for j := 0; j < n; j++ {
					c.Dequeue()
				}
			}
		})
	}
}

func BenchmarkPermitirPeticion(b *testing.B) {
	colas := make(map[string]*Cola)
	ip := "1.2.3.4"
	M := 100
	T := int64(60)
	ts := time.Now().Unix()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Simulamos el paso del tiempo para que no todas las peticiones se rechacen
		PermitirPeticion(colas, ip, ts+int64(i), M, T)
	}
}

func TestMainFunc(t *testing.T) {
	// Crear un archivo temporal "dataColas.log" para la prueba de main
	// de modo que la función main() lo lea sin errores y podamos probar su lógica
	contenidoLog := `192.168.1.1 - - [08/Jun/2026:10:00:00 -0500] "GET / HTTP/1.1" 200 1234
192.168.1.1 - - [08/Jun/2026:10:00:01 -0500] "GET / HTTP/1.1" 200 1234
192.168.1.2 - - [08/Jun/2026:10:00:02 -0500] "GET / HTTP/1.1" 200 1234
192.168.1.1 - - [08/Jun/2026:10:00:03 -0500] "GET / HTTP/1.1" 200 1234
192.168.1.1 - - [08/Jun/2026:10:00:04 -0500] "GET / HTTP/1.1" 200 1234
192.168.1.1 - - [08/Jun/2026:10:00:05 -0500] "GET / HTTP/1.1" 200 1234
192.168.1.1 - - [08/Jun/2026:10:00:06 -0500] "GET / HTTP/1.1" 200 1234
192.168.1.1 - - [08/Jun/2026:10:00:07 -0500] "GET / HTTP/1.1" 200 1234
192.168.1.1 - - [08/Jun/2026:10:00:08 -0500] "GET / HTTP/1.1" 200 1234
192.168.1.1 - - [08/Jun/2026:10:00:09 -0500] "GET / HTTP/1.1" 200 1234
192.168.1.1 - - [08/Jun/2026:10:00:10 -0500] "GET / HTTP/1.1" 200 1234
192.168.1.1 - - [08/Jun/2026:10:00:11 -0500] "GET / HTTP/1.1" 200 1234
linea mal formada que debe ser ignorada
`
	err := os.WriteFile("dataColas.log", []byte(contenidoLog), 0644)
	if err != nil {
		t.Fatalf("No se pudo crear el archivo de prueba: %v", err)
	}
	defer os.Remove("dataColas.log")

	// Capturamos stdout para evitar ensuciar la salida de go test
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Llamamos a main()
	main()

	// Restauramos stdout
	w.Close()
	os.Stdout = oldStdout

	// Solo confirmamos que no hay panic
	var buf [1024]byte
	_, err = r.Read(buf[:])
	if err != nil && err.Error() != "EOF" {
		t.Errorf("Error al leer salida de main: %v", err)
	}
}
