package listas

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"sort"
	"strconv"
)

// RegistroRating estructura los datos individuales de cada línea de u.data
type RegistroRating struct {
	MovieID   int
	Timestamp int64
}

// CargarSecuencia lee el archivo de datos, lo ordena cronológicamente y extrae los IDs.
func CargarSecuencia(ruta string) ([]int, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	// Parseo con bufio y encoding/csv sin librerías externas (Requisito del taller)
	lector := csv.NewReader(bufio.NewReader(archivo))
	lector.Comma = '\t' // Configuración obligatoria para las tabulaciones de u.data

	var registros []RegistroRating

	for {
		linea, err := lector.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Estructura de u.data: userID (0), itemID (1), rating (2), timestamp (3)
		movieID, errID := strconv.Atoi(linea[1])
		timestamp, errTS := strconv.ParseInt(linea[3], 10, 64)

		// Si alguna línea está corrupta o mal formateada, se ignora de forma segura
		if errID != nil || errTS != nil {
			continue
		}

		registros = append(registros, RegistroRating{
			MovieID:   movieID,
			Timestamp: timestamp,
		})
	}

	// Ordenamiento cronológico exigido por la rúbrica usando el paquete nativo 'sort'
	sort.Slice(registros, func(i, j int) bool {
		return registros[i].Timestamp < registros[j].Timestamp
	})

	// Construcción de la secuencia final de accesos para alimentar la caché LRU
	secuencia := make([]int, len(registros))
	for i, reg := range registros {
		secuencia[i] = reg.MovieID
	}

	return secuencia, nil
}
