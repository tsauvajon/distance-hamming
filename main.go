package main

import (
	"fmt"
)

func main() {
	var (
		nbExemples    int // Nombre de lignes d'exemple
		nbColonnes    int // Nombre de colonnes pour chaque ligne d'exemples
		nbClustersMax int // Nombre de clusters maximum à obtenir
	)

	// ## Saisies utilisateur

	fmt.Println("Nombre d'exemples : ")
	fmt.Scanln(&nbExemples)

	fmt.Println("Nombre de colonnes : ")
	fmt.Scanln(&nbColonnes)

	fmt.Println("Nombre de clusters max à créer : ")
	fmt.Scanln(&nbClustersMax)

	fmt.Printf("Matrice de %d x %d, à séparer en %d clusters max.\n\n", nbExemples, nbColonnes, nbClustersMax)

	// On demande à l'utilisateur de remplir la matrice initiale (les exemples)
	matrice := saisieCluster(nbExemples, nbColonnes)

	fmt.Println("Matrice d'exemples :")

	afficheCluster(matrice)

	// ## Calcul des distances de Hamming de la matrice d'exemples

	fmt.Println("Distances de hamming :")

	distancesDeHamming, _ := calculeDistancesHamming(matrice)

	afficheDistancesHamming(distancesDeHamming)

	// ## Début de l'algorithme

	var solutionRetenue []Cluster

	nbClusters := 2

	for nbClusters <= nbClustersMax {
		// On sépare la matrice en 2 clusters aléatoires
		clusters, err := randomSplit(matrice, nbClusters)

		// Si une erreur est retournée c'est qu'on ne peut pas
		// split en N clusters (pas assez d'éléments).
		// Dans ce cas soit on a déjà une solution et on l'affiche,
		// soit on n'a pas de solution, la seule solution possible
		// est donc la matrice de départ.
		if err != nil {
			if solutionRetenue == nil {
				fmt.Printf("Impossible de séparer en %d clusters : pas assez d'éléments.\n", nbClusters)
				fmt.Println("1 seul cluster = celui de départ :")
				fmt.Println()

				afficheCluster(matrice)

				return
			}

			fmt.Println("Impossible de trouver un solution satisfaisante.")
			fmt.Printf("Arrêté car on ne peut pas avoir plus de %d clusters avec cette matrice\n", nbClusters)
			fmt.Println("Meilleure solution trouvée :")
			fmt.Println()

			afficheClusters(solutionRetenue)

			return
		}

		// La hashMap va contenir les hashes des clusters.
		// Cela permet de vérifier, avec une complexité très faible,
		// si cette disposition de clusters a déjà été parcourue
		hashMap := make(map[uint32]bool)

		for {
			h := hash(clusters)

			fmt.Println("hash : ", h)

			if hashMap[h] {
				fmt.Println("Cette disposition de clusters a déjà été parcourue.")
				nbClusters++
				fmt.Println("Augmentation du nombre de clusters: ", nbClusters)
				fmt.Println()

				solutionRetenue = clusters

				break
			}

			// On enregistre la disposition actuelle
			hashMap[h] = true

			// Si tout est ok on tient notre solution, on l'affiche
			if areConditionsSatisfaites(clusters, distancesDeHamming) {
				fmt.Println("Solution satisfaisante !!")
				fmt.Println()

				afficheClusters(clusters)

				return
			}

			// 1) Transférer un élement d'un cluster à un autre
			// 2) S'il n'y a plus qu'un élément dans le cluster,
			// transférer un élement dans ce cluster
			// repeat 2)
		}
	}

	fmt.Println("Impossible de trouver une solution satisfaisante.")
	fmt.Println("Arrêté car le nombre max de clusters a été atteint")
	fmt.Println("Meilleure solution trouvée :")
	fmt.Println()

	afficheClusters(solutionRetenue)
}
