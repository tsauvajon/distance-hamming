// Ce fichier contient tout ce qui permet de calculer et afficher des distances de hamming
package main

import "fmt"

// DistancesHamming est un alias pour [][]int
type DistancesHamming [][]int

// Calcul de toutes les distances de hamming pour un cluster
// Retourne un tableau de distances de hamming, et la distance max
// a l'intérieur de ce cluster
func calculeDistancesHamming(cluster Cluster) ([][]int, int) {
	// On crée un tableau de la taille du cluster x la taille du cluster
	distancesDeHamming := make([][]int, len(cluster))

	max := 0

	// On fait un tour complet d'initialisation,
	// sinon distancesDeHamming[j][i] ne sera
	// par toujours initialisé à temps
	for i := 0; i < len(cluster); i++ {
		distancesDeHamming[i] = make([]int, len(cluster))
	}

	for i := 0; i < len(cluster); i++ {
		for j := i + 1; j < len(cluster); j++ {
			distance := distanceHamming(i, j, distancesDeHamming, cluster) - 1

			if distance > max {
				max = distance
			}

			distancesDeHamming[i][j] = distance + 1
			distancesDeHamming[j][i] = distance + 1
		}
	}

	return distancesDeHamming, max
}

// Parcourt la matrice de distances de hamming et l'affiche
// pour un rendu lisible
func afficheDistancesHamming(distancesDeHamming [][]int) {
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

func distanceHamming(a, b int, distancesDejaCalculees [][]int, cluster Cluster) int {
	// Si on a déjà calculé cette valeur, on la retourne
	if distancesDejaCalculees[a][b] > 0 {
		// -1 car on ne veut pas stocker 0, étant la valeur par défaut.
		// Ca fausserait les comparaisons.
		// On ajoute donc +1 au moment de stocker la valeur et on la soustrait en la récupérant
		// TODO : utiliser un type nullable, ou une structure
		return distancesDejaCalculees[a][b]
	}

	// Si la valeur n'a pas encore été calculée, on la calcule
	count := 0

	for i, a := range cluster[a] {
		if a != cluster[b][i] {
			count++
		}
	}

	return count + 1
}
