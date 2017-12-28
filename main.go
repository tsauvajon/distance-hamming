package main

import (
	"fmt"
)

func main() {
	var (
		nbExemples int // Nombre de lignes d'exemple
		nbColonnes int // Nombre de colonnes pour chaque ligne d'exemples
		nbClusters int // Nombre de clusters à obtenir
	)

	// ## Saisies utilisateur

	fmt.Println("Nombre d'exemples : ")
	fmt.Scanln(&nbExemples)

	fmt.Println("Nombre de colonnes : ")
	fmt.Scanln(&nbColonnes)

	fmt.Println("Nombre de clusters à créer : ")
	fmt.Scanln(&nbClusters)

	fmt.Printf("Matrice de %d x %d, à séparer en %d clusters\n\n", nbExemples, nbColonnes, nbClusters)

	// On demande à l'utilisateur de remplir la matrice initiale (les exemples)
	matrice := saisieCluster(nbExemples, nbColonnes)

	fmt.Println("Matrice d'exemples :")

	afficherCluster(matrice)

	// ## Calcul des distances de Hamming de la matrice d'exemples

	fmt.Println("Distances de hamming :")

	distancesDeHamming, _ := calculeDistancesHamming(matrice)

	afficheDistancesHamming(distancesDeHamming)

	// ## Début de l'algorithme

	// On sépare la matrice en 2 clusters aléatoires
	// clusters := randomSplit(matrice, 2)
}
