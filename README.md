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
3. **Ejercicio 4.3 - Listas (Caché LRU):** [MovieLens Datasets - GroupLens](https://grouplens.org/datasets/movielens/) *(Archivo utilizado: ratings.csv)*
4. **Ejercicio 4.4 - Árboles (Índice AVL):** [MovieLens Datasets - GroupLens](https://grouplens.org/datasets/movielens/) *(Archivos utilizados: movies.csv y ratings.csv)*

---

##  Instrucciones de Ejecución
Asegúrate de tener instalado **Go 1.21+**. 

Para ejecutar las pruebas unitarias y verificar la cobertura del código de cualquier estructura, navega a su carpeta y ejecuta:
```bash
go test -v ./pilas      # Reemplazar "pilas" por la carpeta correspondiente
go test -cover ./pilas  # Para ver el porcentaje de cobertura
