// Ce fichier contient tout ce qui permet de créer,
// modifier, afficher des clusters
package main

import "fmt"

// Cluster est un alias pour [][]bool
type Cluster [][]bool

// Exemple est une ligne d'exemples
type Exemple struct {
	valeurs []bool
	id      int
}

// Demande à l'utilisateur de saisir le contenu d'un cluster
func saisieCluster(nbExemples, nbColonnes int) Cluster {
	var (
		matrice  Cluster // Matrice d'exemples à remplir
		intValue int     // Stocke la saisie utilisateur
	)

	// Création des matrices : première dimension
	matrice = make(Cluster, nbExemples)

	// Création des matrices : 2e dimension
	for i := 0; i < nbExemples; i++ {
		matrice[i] = make([]bool, nbColonnes)
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

	return matrice
}

// Affiche le contenu d'un cluster dans la console
func afficherCluster(cluster Cluster) {
	nbExemples := len(cluster)

	if nbExemples == 0 {
		return
	}

	nbColonnes := len(cluster[0])

	fmt.Print("   |  ")

	for i := 0; i < nbColonnes; i++ {
		fmt.Printf("%d  |  ", i+1)
	}

	fmt.Println()

	for i := 0; i < nbExemples; i++ {
		fmt.Printf("%d  |", i+1)

		for j := 0; j < nbColonnes; j++ {
			var text = "F"
			if cluster[i][j] {
				text = "T"
			}

			fmt.Printf("  %s  |", text)
		}

		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
}

// Sépare aléatoirement une matrice d'exemples en n clusters
func randomSplit(matrice Cluster, n int) []Cluster {
	clusters := make([]Cluster, n)

	for {
		for i := 0; i < len(matrice); i++ {

		}

		// Si on a un split acceptable (au moins 2 exemples par cluster)
		// on renvoie les clusters obtenus
		if ontDeuxExemples(clusters) {
			return clusters
		}
	}
}

// Retourne true si les clusters ont tous au moins 2 exemples
func ontDeuxExemples(clusters []Cluster) bool {
	for _, cluster := range clusters {
		if len(cluster) < 2 {
			return false
		}
	}

	return true
}

// Retourne la moyenne, pour chaque exemple, de ses distances internes
func moyenneDistancesHamming(cluster Cluster) (distances []int) {
	distances = make([]int, len(cluster))
	// TODO
	return distances
}

// Trouve l'élement d'un cluster qui a la distance de hamming
// interne maximum (parmi tous les élements de ce cluster)
// Retourne le tuple index, max
func maxMoyDistanceHamming(cluster Cluster) (index, max int) {
	index = 0
	max = 0
	// TODO
	return index, max
}

func transfereElement(de, vers Cluster, index int) (Cluster, Cluster) {
	// TODO
	return de, vers
}
