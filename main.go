package main

import (
	"fmt"
)

var (
	nbExemples, nbColonnes, intValue int
	matrice                          [][]bool
	distancesHamming                 [][]int
)

func main() {
	fmt.Println("Nombre d'exemples : ")
	fmt.Scanln(&nbExemples)

	fmt.Println("Nombre de colonnes : ")
	fmt.Scanln(&nbColonnes)

	fmt.Printf("Matrice de %d x %d\n\n", nbExemples, nbColonnes)

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

	fmt.Print("   |  ")

	for i := 0; i < nbExemples; i++ {
		fmt.Printf("%d  |  ", i+1)
	}

	fmt.Println()

	// Matrice nbExemples x nbExemples contenatn les distances de hamming
	// Calcul et affichage de toutes les distances de Hamming
	for i := 0; i < nbExemples; i++ {
		fmt.Printf("%d  |", i+1)

		for j := 0; j < nbExemples; j++ {
			if i == j {
				fmt.Print("  -  |")
				continue
			}

			distance := distanceHamming(i, j)

			fmt.Printf("  %d  |", distance)
		}

		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
}

func distanceHamming(a, b int) int {
	// Si on a déjà calculé cette valeur, on la retourne
	if distancesHamming[a][b] > 0 {
		// -1 car on ne veut pas stocker 0, étant la valeur par défaut.
		// On ajoute donc +1 au moment de stocker la valeur et on la soustrait en la récupérant
		return distancesHamming[a][b] - 1
	}

	// Si la valeur n'a pas encore été calculée, on la calcule
	count := 0

	for i, a := range matrice[a] {
		if a != matrice[b][i] {
			count++
		}
	}

	// On stocke la valeur calculée pour éviter de la calculer 2 fois
	distancesHamming[a][b] = count + 1
	distancesHamming[b][a] = count + 1

	return count
}
