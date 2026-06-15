# Taller 3 - Algoritmos y Estructuras de Datos
**Universidad ESAN** Este repositorio contiene la implementación en Go de estructuras de datos lineales y no lineales aplicadas a datasets reales, correspondiente al Taller 3.

##  Integrantes del Equipo
* **Farid Alexander Poma Huaman** - Ejercicio 4.1: Pilas (Stack)
* **Engels Alessandro Quispe Hernandez** - Ejercicio 4.2: Colas (Queue)
* **Julio Lorenzo Santana Mendoza** - Ejercicio 4.3: Listas Enlazadas (LRU Cache)
* **Fabrizio Gabriel Siliman Roman** - Ejercicio 4.4: Árboles Binarios (AVL)

---

##  Estructura del Repositorio
El proyecto sigue la estructura exigida en la especificación del taller:
- `/pilas` - Implementación genérica de Pila y cálculo de Stock Span.
- `/colas` - Implementación de Cola FIFO y Rate Limiter. *(Pendiente de integración)*
- `/listas` - Implementación de Lista Doblemente Enlazada y Caché LRU. *(Pendiente de integración)*
- `/arboles` - Implementación de Árbol AVL y consultas por rango. *(Pendiente de integración)*
- `/diagramas` - Diagramas de funciones de todas las estructuras (PNG/PDF).
- `informe_performance.pdf` - Documento unificado con mediciones de tiempo/memoria y justificación Big-O.

---

##  Origen de los Datasets (No incluidos en el repositorio)
Por políticas del proyecto, los datasets masivos no se suben a Git. A continuación, los enlaces para descargar los datos originales utilizados en cada ejercicio:

1. **Ejercicio 4.1 - Pilas (Stock Span):** [Huge Stock Market Dataset - Kaggle](https://www.kaggle.com/datasets/borismarjanovic/price-volume-data-for-all-us-stocks-etfs) *(Archivo utilizado: amzn.us.txt)*
2. **Ejercicio 4.2 - Colas (Rate Limiter):** [Web Server Access Logs - Kaggle](https://www.kaggle.com/datasets/eliasdabbas/web-server-access-logs)
3. **Ejercicio 4.3 - Listas (Caché LRU):** [MovieLens Datasets - GroupLens](https://grouplens.org/datasets/movielens/)
4. **Ejercicio 4.4 - Árboles (Índice AVL):** [MovieLens Datasets - GroupLens](https://grouplens.org/datasets/movielens/)

---

##  Instrucciones de Ejecución y Pruebas
Para compilar, ejecutar y probar este proyecto localmente, siga los pasos a continuación garantizando que se cumplan las especificaciones del taller:

### 1. Pre-requisitos e Inicialización del Proyecto
* Verifique tener instalado **Go 1.21** o superior.
* Si está configurando el proyecto por primera vez o desde una carpeta vacía, inicialice el módulo con el nombre oficial del proyecto y sincronice las dependencias nativas:
  ```bash
  go mod init taller3
  go mod tidy
2. Preparación de los Datos
Descargue los datasets listados en la sección anterior.

Ubique el archivo de texto correspondiente dentro del directorio local respectivo (por ejemplo, coloque amzn.us.txt en la raíz del proyecto o en la carpeta /pilas/ según se encuentre parametrizado el argumento de ruta en el archivo principal).

3. Ejecución del Programa Principal
El archivo main.go actúa como la compuerta central encargada de instanciar y ejecutar secuencialmente cada uno de los submódulos integrados por el grupo.
* ```bash
  go run main.go
4. Pruebas Unitarias y Cobertura (Testing)Las implementaciones de software incluyen suites de pruebas exhaustivas parametrizadas para validar casos normales, límites (estructuras de datos vacías o con un solo elemento) y aserciones de manejo de errores. Para certificar el estándar de cobertura exigido ($\geq 70\%$):
* ```bash
  # Ejecutar todas las pruebas unitarias del repositorio simultáneamente
  go test -v ./...
  # Analizar la tasa de cobertura de código por módulo
  go test -cover ./...
  # Ejecutar de forma aislada las pruebas del módulo de Pilas (Ejercicio 4.1)
  go test -v ./pilas
  go test -cover ./pilas
5. Evaluación de Benchmarks (Performance)
Para validar de forma empírica la inercia del tiempo de respuesta y las asignaciones del recolector de basura en el Heap:
* ```bash
  go test -bench . -benchmem ./...
## Enlaces a Videos Explicativos (YouTube)
Cada integrante expone de manera individual la sustentación técnica de su código y la correcta interpretación de los resultados:

* Ejercicio 4.1 (Pilas): [Enlace al video explicativo de Stock Span]

* Ejercicio 4.2 (Colas): [Enlace al video explicativo de Rate Limiter]

* Ejercicio 4.3 (Listas): [Enlace al video explicativo de Caché LRU]

* Ejercicio 4.4 (Árboles): [Enlace al video explicativo de Índice AVL]

## Declaración de Ética y Código de Honor
En concordancia estricta con las pautas de Integridad Académica establecidas por la institución:

* Originalidad: Toda la lógica de control, los algoritmos de las estructuras dinámicas de datos y las pruebas automatizadas han sido programados de forma autónoma por los integrantes de este equipo de trabajo.

* Uso Responsable de Inteligencia Artificial: Se emplearon modelos de lenguaje predictivo exclusivamente como herramientas de soporte pedagógico para la definición de plantillas estructurales y la documentación base de datos de prueba en la suite de testing. La revisión lógica paso a paso, el manejo seguro de tipos de datos en Go y el análisis de la complejidad asintótica fueron auditados íntegramente por el equipo operativo.
