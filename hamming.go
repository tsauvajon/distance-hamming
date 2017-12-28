// Ce fichier contient tout ce qui permet de calculer et afficher des distances de hamming
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

// Parcourt la matrice de distances de hamming et l'affiche
// pour un rendu lisible
func afficheDistancesHamming(distancesDeHamming DistancesHamming) {
	fmt.Print("   |  ")

	for i := range distancesDeHamming {
		fmt.Printf("%d  |  ", i+1)
	}

	fmt.Println()

	// Matrice nbExemples x nbExemples contenatn les distances de hamming
	// Calcul et affichage de toutes les distances de Hamming
	for i, row := range distancesDeHamming {
		fmt.Printf("%d  |", i+1)

		for j, dist := range row {
			if i == j {
				fmt.Print("  -  |")
				continue
			}

			fmt.Printf("  %d  |", dist-1)
		}

		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
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
func moyenneDistancesHamming(exemple Exemple) (distance int) {
	distance = 0
	// TODO
	return distance
}

// Retourne la moyenne de toutes les distances internes d'un cluster
func moyennesDistancesHamming(cluster Cluster) (distances []int, maxIndex int) {
	distances = make([]int, len(cluster))
	max := 0
	maxIndex = 0

	for i, exemple := range cluster {
		distances[i] = moyenneDistancesHamming(exemple)

		if distances[i] > max {
			maxIndex = i
		}
	}

	return distances, maxIndex
}

// Trouve l'élement d'un cluster qui a la distance de hamming
// interne maximum (parmi tous les élements de ce cluster)
// Retourne le tuple index, max
func maxDistanceHamming(cluster Cluster) (index, max int) {
	index = 0
	max = 0
	// TODO
	// Si un seul élement a une distance interne max (par exemple 4
	// dans l'énoncé) on le retourne
	// Sinon on retourne l'élement avec la moyenne maximum
	return index, max
}
