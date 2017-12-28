package main

import (
	"fmt"
)

var (
	nbExemples, nbColonnes, nbClusters int // paramètres de l'application
	intValue                           int // saisie d'entier dans la console
	matrice                            [][]bool
	distancesHamming                   [][]int
)

func main() {
	fmt.Println("Nombre d'exemples : ")
	fmt.Scanln(&nbExemples)

	fmt.Println("Nombre de colonnes : ")
	fmt.Scanln(&nbColonnes)

	fmt.Printf("Matrice de %d x %d\n\n", nbExemples, nbColonnes)

	fmt.Println("Nombre de clusters à créer : ")
	fmt.Scanln(&nbClusters)

	// Création des matrices : première dimension
	matrice = make([][]bool, nbExemples)
	distancesHamming = make([][]int, nbExemples)

	// Création des matrices : 2e dimension
	for i := 0; i < nbExemples; i++ {
		matrice[i] = make([]bool, nbColonnes)
		distancesHamming[i] = make([]int, nbExemples)
	}

	// Pour chaque valeur, renseigner la classe (1 ou 2)
	fmt.Println("Remplissage de la matrice: [ligne d'example; colonne]")

	for i := 0; i < nbExemples; i++ {
		for j := 0; j < nbColonnes; j++ {
			// +1 pour avoir avoir un format "humain" démarrant par 1 plutôt que 0
			fmt.Printf("[%d; %d] : ", i+1, j+1)
			fmt.Scan(&intValue)
			fmt.Println()

			matrice[i][j] = intValue == 1
		}
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("Matrice d'exemples :")

	fmt.Print("   |  ")

	for i := 0; i < nbColonnes; i++ {
		fmt.Printf("%d  |  ", i+1)
	}

	fmt.Println()

	for i := 0; i < nbExemples; i++ {
		fmt.Printf("%d  |", i+1)

		for j := 0; j < nbColonnes; j++ {
			var text = "F"
			if matrice[i][j] {
				text = "T"
			}

			fmt.Printf("  %s  |", text)
		}

		fmt.Println()
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("Distances de hamming :")

	distancesDeHamming, _ := calculeDistancesHamming(matrice)
	afficheDistancesHamming(distancesDeHamming)
}

// Calcul de toutes les distances de hamming pour un cluster
// Retourne un tableau de distances de hamming, et la distance max
// a l'intérieur de ce cluster
func calculeDistancesHamming(cluster [][]bool) ([][]int, int) {
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

	for i := 0; i < nbExemples; i++ {
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

func distanceHamming(a, b int, distancesDejaCalculees [][]int, cluster [][]bool) int {
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
