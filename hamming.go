// Ce fichier contient tout ce qui permet de calculer et comparer des distances de hamming
package main

import "fmt"

// DistancesHamming est un alias pour [][]int
type DistancesHamming [][]int

// Calcul de toutes les distances de hamming pour un cluster
// Retourne un tableau de distances de hamming, et la distance max
// a l'intérieur de ce cluster
func calculeDistancesHamming(cluster Cluster) (DistancesHamming, int) {
	// On crée un tableau de la taille du cluster x la taille du cluster
	distancesDeHamming := make(DistancesHamming, len(cluster))

	max := 0

	// On fait un tour complet d'initialisation,
	// sinon distancesDeHamming[j][i] ne sera
	// par toujours initialisé à temps
	for i := range cluster {
		distancesDeHamming[i] = make([]int, len(cluster))
	}

	for i := range cluster {
		for j := i + 1; j < len(cluster); j++ {
			distance := distanceHamming(cluster[i], cluster[j], distancesDeHamming) - 1

			if distance > max {
				max = distance
			}

			jID := cluster[j].id
			iID := cluster[i].id

			distancesDeHamming[iID][jID] = distance + 1
			distancesDeHamming[jID][iID] = distance + 1
		}
	}

	return distancesDeHamming, max
}

func distanceHamming(a, b Exemple, distancesDejaCalculees DistancesHamming) int {
	// Si on a déjà calculé cette valeur, on la retourne
	if distancesDejaCalculees[a.id][b.id] > 0 {
		// -1 car on ne veut pas stocker 0, étant la valeur par défaut.
		// Ca fausserait les comparaisons.
		// On ajoute donc +1 au moment de stocker la valeur et on la soustrait en la récupérant
		// TODO : utiliser un type nullable, ou une structure
		return distancesDejaCalculees[a.id][b.id]
	}

	// Si la valeur n'a pas encore été calculée, on la calcule
	count := 0

	// Pour chaque valeur différente, on incrémente le compteur
	for i, aVal := range a.valeurs {
		if aVal != b.valeurs[i] {
			count++
		}
	}

	// Count + 1 pour éviter les comparaisons avec 0
	return count + 1
}

// Retourne la moyenne des distances internes d'un exemple donné en paramètre
func moyenneDistancesInternes(index int, cluster Cluster, distances DistancesHamming) (moyenne float32) {
	total := 0

	a := cluster[index]

	for j, b := range cluster {
		if index == j {
			continue
		}

		total += distances[a.id][b.id]
	}

	moyenne = float32(total) / float32(len(cluster))

	return moyenne
}

// Retourne la moyenne de toutes les distances internes d'un cluster
// Mais est-ce seulement utile ?
func moyennesDistancesInternes(cluster Cluster, distances DistancesHamming) (moyennes []float32, maxIndex int) {
	var max float32
	maxIndex = 0
	moyennes = make([]float32, len(cluster))

	for i := range cluster {
		moyennes[i] = moyenneDistancesInternes(i, cluster, distances)

		if moyennes[i] > max {
			maxIndex = i
			max = moyennes[i]
		}
	}

	return moyennes, maxIndex
}

// Pour éviter de faire 50 fois les mêmes calculs on sauvegarde dans une map
type mdiResult struct {
	calculated bool
	index, max int
	maxMoy     float32
}

var sauvegardeMaxDistancesInternes map[uint32]mdiResult

// Trouve les éléments qui ont la + grande distance interne,
// et parmi ceux là celui qui a la + grande moyenne de distances internes
// TODO : méthode qui compare UN EXEMPLE avec UN CLUSTER (interne ou externe)
func maxDistanceInterne(cluster Cluster, distances DistancesHamming) (index, max int, maxMoy float32) {
	concat := fmt.Sprintf("%#v%#v", cluster, distances)
	h := hash(concat)

	// Si on a déjà calculé ce résultat on le renvoie direct
	if save := sauvegardeMaxDistancesInternes[h]; save.calculated {
		return save.index, save.max, save.maxMoy
	}

	// utilisation d'un map plutôt qu'un array pour pouvoir
	// facilement vérifier si elle contient un élément
	var indexes map[int]bool
	max = 0
	maxMoy = 0

	// On cherche la distance max, et tous les exemples qui "participent"
	// à cette distance max
	for i, exemple := range cluster {
		for j := i + 1; j < len(cluster); j++ {
			a := exemple.id
			b := cluster[j].id

			distance := distances[a][b]

			switch {
			case distance > max:
				// on vide le tableau d'indexes
				indexes = make(map[int]bool, 0)
				max = distance

				fallthrough
			case distance == max:
				// on ajoute les indexes des 2 distances
				indexes[i] = true
				indexes[j] = true
			}
		}
	}

	// On a tous nos index à inspecter.
	// On récupère l'index qui a la distance interne moyenne max
	// parmi ces index là
	for i := range indexes {
		moy := moyenneDistancesInternes(i, cluster, distances)

		if moy > maxMoy {
			maxMoy = moy
			index = i
		}
	}

	sauvegardeMaxDistancesInternes[h] = mdiResult{
		calculated: true,
		index:      index,
		max:        max,
		maxMoy:     maxMoy,
	}

	return index, max, maxMoy
}

// Converts int map keys to an int array
func mapToArray(m map[int]interface{}) (out []int) {
	out = make([]int, len(m))

	for key := range m {
		out = append(out, key)
	}

	return out
}
