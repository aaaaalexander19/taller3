package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Nodo struct {
	ts   int64
	next *Nodo
}

// Cola ahora usa punteros al primer y último nodo, garantizando O(1)
type Cola struct {
	head *Nodo // Frente de la cola (por donde salen)
	tail *Nodo // Final de la cola (por donde entran)
	size int   // Mantenemos el tamaño para que Len() sea O(1)
}

func (c *Cola) Enqueue(ts int64) {
	nuevoNodo := &Nodo{ts: ts, next: nil}

	if c.size == 0 {
		// Si está vacía, el nuevo nodo es tanto la cabeza como la cola
		c.head = nuevoNodo
		c.tail = nuevoNodo
	} else {
		// Lo agregamos al final y actualizamos la cola
		c.tail.next = nuevoNodo
		c.tail = nuevoNodo
	}
	c.size++
}

func (c *Cola) Dequeue() (int64, bool) {
	if c.size == 0 {
		return 0, false
	}

	val := c.head.ts
	// Movemos la cabeza al siguiente nodo.
	// El nodo viejo queda sin referencias y el "Garbage Collector" de Go lo elimina.
	c.head = c.head.next
	c.size--

	if c.size == 0 {
		c.tail = nil // Seguridad extra si la cola quedó vacía
	}

	return val, true
}

func (c *Cola) Front() (int64, bool) {
	if c.size == 0 {
		return 0, false
	}
	return c.head.ts, true
}

func (c *Cola) Len() int {
	return c.size
}

// Lógica intacta: El Rate Limiter no necesita saber CÓMO funciona la cola por dentro
func PermitirPeticion(colas map[string]*Cola, ip string, ts int64, M int, T int64) bool {
	colaIP, existe := colas[ip]
	if !existe {
		colaIP = &Cola{}
		colas[ip] = colaIP
	}

	tiempoLimite := ts - T

	frente, ok := colaIP.Front()
	for ok && frente < tiempoLimite {
		colaIP.Dequeue()
		frente, ok = colaIP.Front()
	}

	if colaIP.Len() < M {
		colaIP.Enqueue(ts)
		return true
	}

	return false
}

func ParsearLinea(linea string) (ip string, ts int64, err error) {
	partes := strings.Split(linea, " ")
	if len(partes) < 4 {
		return "", 0, errors.New("línea mal formada")
	}

	ip = partes[0]
	fechaString := partes[3]

	layout := "[02/Jan/2006:15:04:05"
	fechaGo, errParseo := time.Parse(layout, fechaString)
	if errParseo != nil {
		return "", 0, errParseo
	}

	ts = fechaGo.Unix()
	return ip, ts, nil
}

func main() {
	rutaLog := "dataColas.log"
	M := 10        // Máximo 10 peticiones
	T := int64(60) // Ventana de 60 segundos

	colas := make(map[string]*Cola)
	rechazosPorIP := make(map[string]int)
	totalRechazos := 0

	archivo, err := os.Open(rutaLog)
	if err != nil {
		fmt.Printf("Error al abrir el archivo %s: %v\n", rutaLog, err)
		return
	}
	defer archivo.Close()

	peticionesProcesadas := 0
	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		linea := scanner.Text()
		ip, ts, err := ParsearLinea(linea)
		if err != nil {
			continue
		}

		peticionesProcesadas++
		esMuestra := peticionesProcesadas <= 20

		if !PermitirPeticion(colas, ip, ts, M, T) {
			if esMuestra {
				fmt.Printf("RECHAZADA: IP=%s\n", ip)
			}
			rechazosPorIP[ip]++
			totalRechazos++
		} else {
			if esMuestra {
				fmt.Printf("ACEPTADA:  IP=%s\n", ip)
			}
		}
	}

	fmt.Println("\n... procesamiento completado ...")

	fmt.Println("\n--- RESUMEN DE RECHAZOS ---")
	fmt.Printf("Total de peticiones rechazadas: %d\n", totalRechazos)

	type ipCount struct {
		IP    string
		Count int
	}

	var listaRechazos []ipCount
	for ip, count := range rechazosPorIP {
		listaRechazos = append(listaRechazos, ipCount{IP: ip, Count: count})
	}

	n := len(listaRechazos)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if listaRechazos[j].Count < listaRechazos[j+1].Count {
				listaRechazos[j], listaRechazos[j+1] = listaRechazos[j+1], listaRechazos[j]
			}
		}
	}

	fmt.Println("\nTop 5 IPs con más rechazos:")
	for i := 0; i < 5 && i < len(listaRechazos); i++ {
		fmt.Printf("%d. IP: %-15s | Rechazos: %d\n", i+1, listaRechazos[i].IP, listaRechazos[i].Count)
	}
}
