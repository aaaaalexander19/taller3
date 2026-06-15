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

3. Ejecución con Datos Reales (Pruebas de Integración)
Para evaluar los algoritmos procesando los archivos masivos reales y visualizar los resultados tabulados en la consola, cada integrante ha diseñado pruebas de integración específicas:
Ejercicio 4.1 - Pilas (Stock Span):
* ```bash
  go test -v -run TestIntegracionDatasetReal ./pilas
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

* Ejercicio 4.1 (Pilas): https://youtu.be/37XYDGvqGE4

* Ejercicio 4.2 (Colas): [Enlace al video explicativo de Rate Limiter]

* Ejercicio 4.3 (Listas): https://youtu.be/NHgsPvbMgPc

* Ejercicio 4.4 (Árboles): [Enlace al video explicativo de Índice AVL]

## Declaración de uso de Inteligencia Artificial

Para la realización de este taller, se utilizó asistencia de Inteligencia Artificial (Gemini) en el siguiente aspecto específico:

Testing (colas_test.go): Se empleó asistencia de IA para la generación estructurada de las pruebas automatizadas (unitarias y benchmarks), específicamente para las funciones main(), ParsearLinea() y PermitirPeticion(), con el objetivo de establecer una suite de testing robusta que superara el 70% de cobertura (coverage) requerido en la rúbrica.

Lógica principal (colas.go): El uso de IA se limitó estrictamente a la asistencia en la verificación de funciones y depuración (debugging) para solucionar errores lógicos puntuales, manteniendo la autoría del equipo sobre la estructura de datos.

Validación: Todo el código sugerido o modificado por la IA fue leído, analizado, refactorizado y comprendido por los autores. La validación empírica se realizó ejecutando go test -v, go test -cover y comprobando manualmente que el resultado de main() sobre el archivo dataColas.log (recorte del dataset original) fuera el esperado y respetara  la complejidad algorítmica solicitada (O(1)).

## Declaración de Ética y Código de Honor
En concordancia estricta con las pautas de Integridad Académica establecidas por la institución:

* Originalidad: Toda la lógica de control y los algoritmos fundamentales de las estructuras dinámicas de datos han sido pensados, diseñados y programados por los integrantes de este equipo de trabajo

* Uso Responsable de Inteligencia Artificial:  Se emplearon modelos de lenguaje predictivo exclusivamente como herramientas de soporte pedagógico (generación de la base de la suite de pruebas y asistencia en depuración). La revisión lógica paso a paso, el manejo seguro de tipos de datos en Go, la adaptación al contexto del problema y el análisis de la complejidad asintótica fueron auditados íntegramente por el equipo operativo, quienes asumen total autoría y capacidad de sustentar cada línea del proyecto
