// méthodes utilitaires
package main

import (
	"fmt"
	"hash/fnv"
)

// Génère un hash uint32 à partir de n'importe quel objet
func hash(o interface{}) uint32 {
	return hashString(fmt.Sprintf("%#v", o))
}

// Génère un hash uint32 à partir d'un string
// Peut générer des collisions
func hashString(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
