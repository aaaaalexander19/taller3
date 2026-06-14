package listas

import "testing"

func TestLRUBasico(t *testing.T) {
	cache := NuevoLRU(2)

	cache.Put(1, 100)
	cache.Put(2, 200)

	if val, ok := cache.Get(1); !ok || val != 100 {
		t.Errorf("Se esperaba encontrar la clave 1 con valor 100")
	}
}

func TestLRUMoveToHead(t *testing.T) {
	cache := NuevoLRU(2)
	cache.Put(1, 100)
	cache.Put(2, 200)

	// Al hacer Get(1), el 1 pasa a ser el más reciente, y el 2 se queda atrás
	cache.Get(1)

	// Esto debería expulsar al 2, no al 1
	cache.Put(3, 300)

	if _, ok := cache.Get(2); ok {
		t.Errorf("La clave 2 debió ser desalojada")
	}
	if _, ok := cache.Get(1); !ok {
		t.Errorf("La clave 1 debía mantenerse en la caché")
	}
}

func TestLRUEvict(t *testing.T) {
	cache := NuevoLRU(2)
	cache.Put(1, 100)
	cache.Put(2, 200)
	cache.Put(3, 300) // Supera capacidad, desaloja al 1

	if _, ok := cache.Get(1); ok {
		t.Errorf("La clave 1 debió ser desalojada")
	}
}
