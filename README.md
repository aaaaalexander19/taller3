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
  ```bash
  go run main.go
