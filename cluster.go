// Ce fichier contient tout ce qui permet de créer,
// modifier des clusters
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

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

// Sépare aléatoirement une matrice d'exemples en n clusters
func randomSplit(matrice Cluster, n int) ([]Cluster, error) {
	// Si on ne peut pas mettre au moins 2 éléments dans chaque cluster
	// => nique la police
	if len(matrice)/2 < n {
		return nil, errors.New("Impossible de séparer cette matrice : il faut au moins 2 éléments par cluster")
	}

	clusters := make([]Cluster, n)

	for {
		// Initialisation des clusters vides
		for i := range clusters {
			clusters[i] = make(Cluster, 0)
		}

		// On ajoute chaque exemple à un cluster aléatoire
		seed := rand.NewSource(time.Now().UnixNano())
		randomizer := rand.New(seed)
		for _, exemple := range matrice {
			rand := randomizer.Intn(n)
			clusters[rand] = append(clusters[rand], exemple)
		}

		// Si on a pas un split acceptable (au moins 2 exemples par cluster)
		// on réitère
		if !ontDeuxExemples(clusters) {
			continue
		}

		// Sinon tout est ok renvoie les clusters obtenus
		return clusters, nil
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

// On remarquera le magnifique franglais
func areConditionsSatisfaites(clusters []Cluster, distances DistancesHamming) bool {
	maxToutesDistancesInternes := 0
	// minToutesDistancesExternes := math.MaxInt32

	// Calcul de la plus grande de toutes les distances internes
	for _, cluster := range clusters {
		// calcul max distance interne
		// compare avec chaque distance externe
		// si une externe > interne return false
		_, dist, _ := maxDistanceInterne(cluster, distances)

		if dist > maxToutesDistancesInternes {
			maxToutesDistancesInternes = dist
		}
	}

	// Calcul de la plus petite de toutes les dinstances externes
	return true
}
