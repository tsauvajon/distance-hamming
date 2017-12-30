package main

import (
	"fmt"
)

func main() {
	var (
		nbExemples    int // Nombre de lignes d'exemple
		nbColonnes    int // Nombre de colonnes pour chaque ligne d'exemples
		nbClustersMin int // Nombre de clusters minimum à obtenir
		nbClustersMax int // Nombre de clusters maximum à obtenir
	)

	// ## Saisies utilisateur

	fmt.Println("Nombre d'exemples : ")
	fmt.Scanln(&nbExemples)

	fmt.Println("Nombre de colonnes : ")
	fmt.Scanln(&nbColonnes)

	fmt.Println("Nombre de clusters min à créer : ")
	fmt.Scanln(&nbClustersMin)

	fmt.Println("Nombre de clusters max à créer : ")
	fmt.Scanln(&nbClustersMax)

	fmt.Printf("Matrice de %d x %d, à séparer entre %d et  %d clusters.\n\n", nbExemples, nbColonnes, nbClustersMin, nbClustersMax)

	// On demande à l'utilisateur de remplir la matrice initiale (les exemples)
	matrice := saisieCluster(nbExemples, nbColonnes)

	fmt.Println("Matrice d'exemples :")

	afficheCluster(matrice)

	// ## Calcul des distances de Hamming de la matrice d'exemples

	fmt.Println("Distances de hamming :")

	distancesDeHamming, _ := calculeDistancesHamming(matrice)

	afficheDistancesHamming(distancesDeHamming)

	// ## Début de l'algorithme

	saveCompare = make(map[Hash]compareResult)

	solutionRetenue := make([]Cluster, 1)

	solutionRetenue[0] = matrice

	nbClusters := nbClustersMin

	for nbClusters <= nbClustersMax {
		// On sépare la matrice en 2 clusters aléatoires
		clusters, err := randomSplit(matrice, nbClusters)

		// Si une erreur est retournée c'est qu'on ne peut pas
		// split en N clusters (pas assez d'éléments).
		// On termine donc l'algo
		if err != nil {
			break
		}

		// La hashMap va contenir les hashes des clusters.
		// Cela permet de vérifier, avec une complexité plus
		// faible que de faire un for dans un for dans un for,
		// si cette disposition de clusters a déjà été parcourue
		hashMap := make(map[Hash]bool)

		for {
			h := hash(clusters)

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

			// A améliorer
			hashMap2 := make(map[Hash]bool)
			exit := false

			for {
				clusterIndex, elemIndex := trouverElementADeplacer(clusters, distancesDeHamming)

				elemADeplacer := clusters[clusterIndex][elemIndex]

				versIndex := trouverVersOuDeplacer(elemADeplacer, clusterIndex, clusters, distancesDeHamming)

				fmt.Printf("Déplacer l'exemple clusters[%d][%d] (id: %d) vers le cluster clusters[%d]\n", clusterIndex, elemIndex, elemADeplacer.id, versIndex)

				clusters = transfereElement(clusterIndex, elemIndex, versIndex, clusters)

				// ## ANTI BOUCLE INFINIE DEGUEULASSE
				h := hash(clusters)

				if hashMap2[h] {
					fmt.Println("Cette disposition de clusters a déjà été parcourue.")
					nbClusters++
					fmt.Println("Augmentation du nombre de clusters: ", nbClusters)
					fmt.Println()

					solutionRetenue = clusters

					exit = true

					break
				}

				// On enregistre la disposition actuelle
				hashMap2[h] = true
				// ## FIN ANTI BOUCLE INFINIE DEGUEULASSE

				// Continue jusqu'à ce que tous les clusters aient au moins
				// 2 exemples
				if ontTousDeuxExemples(clusters) {
					break
				}
			}

			if exit {
				break
			}
		}
	}

	fmt.Println("Impossible de trouver une solution satisfaisante.")
	fmt.Println("Meilleure solution trouvée :")
	fmt.Println()

	afficheClusters(solutionRetenue)
}
