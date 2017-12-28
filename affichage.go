// Affichage des différentes parties de l'algo
package main

import "fmt"

// Parcourt la matrice de distances de hamming et l'affiche
// pour un rendu lisible
func afficheDistancesHamming(distancesDeHamming DistancesHamming) {
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

// Affiche le contenu d'un cluster dans la console
func afficheCluster(cls Cluster) {
	fmt.Print("   |  ")

	for i := range cls[0].valeurs {
		fmt.Printf("%d  |  ", i+1)
	}

	fmt.Println()

	for _, exemple := range cls {
		fmt.Printf("%d  |", exemple.id+1)

		for _, valeur := range exemple.valeurs {
			text := "F"
			if valeur {
				text = "T"
			}

			fmt.Printf("  %s  |", text)
		}

		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
}

func afficheClusters(clusters []Cluster) {
	for i, cluster := range clusters {
		fmt.Println()
		fmt.Printf("CLUSTER %d : \n", i+1)
		fmt.Println()
		afficheCluster(cluster)
		fmt.Println("__________")
		fmt.Println()
	}
}
