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

	// TODO
	// On cherche à séparer la matrice en n clusters

	// ## 1 On démarre en splittant aléatoirement les exemples en N clusters (au moins 2 éléments par cluster)
	// ## 2 On vérifie si les conditions sont remplie => Oui = fini
	// Conditions = toutes les distances internes <= toutes les distances externes
	// ## 3 On sépare les exemples qui ont une trop grande distance entre eux (1 transfert d'un cluster vers un autre)
	// Le transfert se fait vers le cluster qui a la distance moyenne la plus faible avec l'élément à transférer
	// ## 4 Si on a déjà eu cette configuration => fini (solution non complète)
	// ## 5 Si un des clusters a 1 seul élément => "transvaser" un élément d'un autre cluster (celui avec la + grosse distance moyenne) et retour à l'étape ## 4
	// ## 6 Retour à l'étape ## 2
}
