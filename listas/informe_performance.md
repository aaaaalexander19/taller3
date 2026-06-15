# Informe de Performance - Taller 3

**Alumno:** Julio Lorenzo Santana Mendoza

**Código:** 25100130
## 1. Especificaciones del Entorno de Medición

| Componente | Detalle |
| :--- | :--- |
| **Procesador (CPU)** | AMD Ryzen 7 5700U with Radeon Graphics (1.80 GHz) |
| **Memoria RAM** | 12.0 GB (11.3 GB utilizable) |
| **Sistema Operativo** | Windows 11 Home Single Language|
| **Versión de Go** | go1.26.1 windows/amd64 |

---

## 2. Ejercicio 4.3: Caché LRU con Lista Doblemente Enlazada

### A. Resultados de la Simulación (Dataset MovieLens 100K)
Se procesó la secuencia completa de accesos del archivo `u.data`, ordenada cronológicamente por su marca de tiempo (*timestamp*). A continuación, se presentan los resultados de **Hit Ratio** (Aciertos / Accesos totales) variando la capacidad máxima de la caché según lo solicitado:

| Capacidad de Caché | Hits (Aciertos) | Misses (Fallos) | Hit Ratio (%) |
| :---: | :---: | :---: | :---: |
| **50** | 4485 | 95515 | 4.49% |
| **100** | 10864 | 89136 | 10.86% |
| **500** | 64146 | 35854 | 64.15% |
| **1000** | 92572 | 7428 | 92.57% |

### B. Análisis de la Tendencia Observada
Al analizar los datos y la gráfica de rendimiento, podemos notar un comportamiento en forma de asíntota ascendente: **a medida que aumenta la capacidad de la caché, el Hit Ratio se eleva notablemente**. 

Esto se justifica mediante el principio de **localidad temporal** de los datos bursátiles y de consumo en datasets reales. En el contexto de MovieLens, los usuarios suelen interactuar con películas populares o realizar evaluaciones en ráfagas de tiempo cercanas. Una caché con mayor capacidad (`cap`) retiene estos elementos calientes en memoria durante más tiempo antes de que sean expulsados por el extremo de la cola (*Tail*), disminuyendo drásticamente los accesos costosos a almacenamiento (Misses).

### C. Complejidad Teórica vs. Empírica
La rúbrica del taller exige garantizar operaciones de lectura y escritura en tiempo constante:
* **Complejidad Teórica:** O(1) tanto para `Get` como para `Put`.
* **Justificación de Diseño:** 1. El acceso directo al elemento se realiza mediante un mapa (`map[int]*Nodo`), lo que permite buscar o validar la existencia de cualquier clave en tiempo constante O(1).
  2. La reordenación de la prioridad se delega en una lista doblemente enlazada implementada manualmente. Mover un nodo existente a la cabeza (*Head*) o expulsar el nodo menos usado (*Tail*) se logra modificando únicamente los punteros adyacentes (`prev` y `next`), una operación que toma tiempo estrictamente constante O(1) sin importar el tamaño total de los datos en memoria.

### D. Evidencia de Benchmarks Automatizados
*Resultados obtenidos tras ejecutar las pruebas de rendimiento con el comando nativo de Go `go test -bench=. -benchmem ./listas`*

```text
goos: windows
goarch: amd64
pkg: taller3/listas
cpu: AMD Ryzen 7 5700U with Radeon Graphics
=== RUN   BenchmarkLRUPut
BenchmarkLRUPut
BenchmarkLRUPut-16
11607492               109.6 ns/op            32 B/op          1 allocs/op
PASS
ok      taller3/listas  1.771s
