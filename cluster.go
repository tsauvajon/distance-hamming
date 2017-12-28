// Ce fichier contient tout ce qui permet de créer,
// modifier, afficher des clusters
package main

import "fmt"

// Cluster est un alias pour []Exemple
type Cluster []Exemple

// Exemple est une ligne d'exemples
type Exemple struct {
	// Contient les valeurs, mises sous forme de booléen,
	// de la ligne d'exemples
	valeurs []bool
	// On sauvegarde un "id" d'exemple, pour permettre de calculer
	// 1 seule fois chaque distance de hamming.
	// Quand on transfère un élément d'un cluster à un autre,
	// son index dans le cluster va changer, et il faut donc
	// un moyen de l'identifier
	id int
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
		matrice[i].valeurs = make([]bool, nbColonnes)
	}

	// Pour chaque valeur, renseigner la classe (1 ou 2)
	fmt.Println("Remplissage de la matrice: [ligne d'example; colonne]")

	for i := 0; i < nbExemples; i++ {
		// on enregistre un "id" de la ligne d'exemple pour pouvoir
		// la retrouver précisemment quand elle sera dans d'autres
		// clusters
		matrice[i].id = i

		for j := 0; j < nbColonnes; j++ {
			// +1 pour avoir avoir un format "humain" démarrant par 1 plutôt que 0
			fmt.Printf("[%d; %d] : ", i+1, j+1)
			fmt.Scan(&intValue)
			fmt.Println()

			// On enregistre la valeur saisie par l'utilisateur
			// 1 = true
			// toutes les autres valeurs = false
			matrice[i].valeurs[j] = intValue == 1
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

	nbColonnes := len(cluster[0].valeurs)

	fmt.Print("   |  ")

	for i := 0; i < nbColonnes; i++ {
		fmt.Printf("%d  |  ", i+1)
	}

	fmt.Println()

	for i := 0; i < nbExemples; i++ {
		fmt.Printf("%d  |", i+1)

		for j := 0; j < nbColonnes; j++ {
			var text = "F"
			if cluster[i].valeurs[j] {
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

func transfereElement(de, vers Cluster, index int) (Cluster, Cluster) {
	// TODO
	return de, vers
}
