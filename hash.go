// méthodes utilitaires
package main

import (
	"fmt"
	"hash/fnv"
)

// Hash est un alias de uint32
type Hash uint32

// Génère un hash uint32 à partir de n'importe quel objet
func hash(o interface{}) Hash {
	return hashString(fmt.Sprintf("%#v", o))
}

// Génère un hash uint32 à partir d'un string
// Peut générer des collisions
func hashString(s string) Hash {
	h := fnv.New32a()
	h.Write([]byte(s))
	return Hash(h.Sum32())
}
