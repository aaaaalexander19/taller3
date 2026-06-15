package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func extraerAnio(titulo string) (float64, error) {
	re := regexp.MustCompile(`\((\d{4})\)\s*$`)
	coincidencias := re.FindStringSubmatch(titulo)
	if len(coincidencias) > 1 {
		anio, err := strconv.ParseFloat(coincidencias[1], 64)
		if err == nil {
			return anio, nil
		}
	}
	return 0, fmt.Errorf("año no encontrado")
}

func main() {
	arbol := &ArbolAVL{}

	rutaArchivo := "movies.csv"
	archivo, err := os.Open(rutaArchivo)
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %v\n", err)
		return
	}
	defer archivo.Close()

	lector := csv.NewReader(bufio.NewReader(archivo))

	_, err = lector.Read()
	if err != nil {
		fmt.Printf("Error al leer cabeceras: %v\n", err)
		return
	}

	fmt.Println("Procesando dataset e insertando en el Árbol AVL...")

	registrosProcesados := 0
	for {
		linea, err := lector.Read()
		if err != nil {
			break
		}

		id := linea[0]
		titulo := linea[1]
		generos := linea[2]

		anio, err := extraerAnio(titulo)
		if err == nil {
			peli := Pelicula{ID: id, Titulo: titulo, Generos: generos}
			arbol.Insertar(anio, peli)
			registrosProcesados++
		}
	}

	fmt.Printf("\n--- MÉTRICAS DEL ÁRBOL AVL ---\n")
	fmt.Printf("Total de películas procesadas: %d\n", registrosProcesados)
	fmt.Printf("Número de Nodos (Años únicos): %d\n", arbol.NumNodos)
	fmt.Printf("Altura del árbol: %d\n", Altura(arbol.Raiz))
	fmt.Printf("Total de Rotaciones realizadas: %d\n", arbol.Rotaciones)

	anioInicio := 1990.0
	anioFin := 1992.0

	fmt.Printf("\n--- CONSULTA POR RANGO: %.0f a %.0f ---\n", anioInicio, anioFin)
	resultados := arbol.ConsultaRango(anioInicio, anioFin)

	fmt.Printf("Se encontraron %d películas en este rango.\n", len(resultados))
	fmt.Println("Mostrando las primeras 5:")

	limite := 5
	if len(resultados) < 5 {
		limite = len(resultados)
	}
	for i := 0; i < limite; i++ {
		fmt.Printf("- [%.0f] %s\n", anioInicio, resultados[i].Titulo) // Nota: para exactitud del año de cada película, se puede agregar a la struct Pelicula.
	}
}
