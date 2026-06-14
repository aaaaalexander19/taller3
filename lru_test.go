package listas

// Nodo representa un elemento en la lista doblemente enlazada
type Nodo struct {
	clave, valor int
	prev, next   *Nodo
}

// LRU representa la caché Least Recently Used
type LRU struct {
	cap  int
	mapa map[int]*Nodo
	head *Nodo
	tail *Nodo
}

// NuevoLRU inicializa una nueva caché con una capacidad dada
func NuevoLRU(capacidad int) *LRU {
	return &LRU{
		cap:  capacidad,
		mapa: make(map[int]*Nodo),
	}
}

// Get busca un elemento en la caché
func (c *LRU) Get(clave int) (int, bool) {
	if nodo, existe := c.mapa[clave]; existe {
		c.moveToHead(nodo)
		return nodo.valor, true
	}
	return 0, false
}

// Put inserta o actualiza un elemento en la caché
func (c *LRU) Put(clave, valor int) {
	if nodo, existe := c.mapa[clave]; existe {
		nodo.valor = valor
		c.moveToHead(nodo)
		return
	}

	nuevoNodo := &Nodo{clave: clave, valor: valor}
	c.mapa[clave] = nuevoNodo
	c.addToHead(nuevoNodo)

	if len(c.mapa) > c.cap {
		c.evict()
	}
}

// --- Métodos Auxiliares para Manipulación de Punteros O(1) ---

func (c *LRU) addToHead(nodo *Nodo) {
	nodo.next = c.head
	nodo.prev = nil
	if c.head != nil {
		c.head.prev = nodo
	}
	c.head = nodo
	if c.tail == nil {
		c.tail = nodo
	}
}

func (c *LRU) remove(nodo *Nodo) {
	if nodo.prev != nil {
		nodo.prev.next = nodo.next
	} else {
		c.head = nodo.next
	}
	if nodo.next != nil {
		nodo.next.prev = nodo.prev
	} else {
		c.tail = nodo.prev
	}
}

func (c *LRU) moveToHead(nodo *Nodo) {
	c.remove(nodo)
	c.addToHead(nodo)
}

func (c *LRU) evict() {
	if c.tail == nil {
		return
	}
	eliminado := c.tail
	c.remove(eliminado)
	delete(c.mapa, eliminado.clave)
}
