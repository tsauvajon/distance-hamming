// Ce fichier contient tout ce qui permet de calculer et comparer des distances de hamming
package main

import "fmt"
import "math"

// DistancesHamming est un alias pour [][]int
type DistancesHamming [][]int

// Calcul de toutes les distances de hamming pour un cluster
// Retourne un tableau de distances de hamming, et la distance max
// a l'intérieur de ce cluster
func calculeDistancesHamming(cluster Cluster) (DistancesHamming, int) {
	// On crée un tableau de la taille du cluster x la taille du cluster
	distancesDeHamming := make(DistancesHamming, len(cluster))

	max := 0

	// On fait un tour complet d'initialisation,
	// sinon distancesDeHamming[j][i] ne sera
	// par toujours initialisé à temps
	for i := range cluster {
		distancesDeHamming[i] = make([]int, len(cluster))
	}

	for i := range cluster {
		for j := i + 1; j < len(cluster); j++ {
			distance := distanceHamming(cluster[i], cluster[j], distancesDeHamming) - 1

			if distance > max {
				max = distance
			}

			jID := cluster[j].id
			iID := cluster[i].id

			distancesDeHamming[iID][jID] = distance + 1
			distancesDeHamming[jID][iID] = distance + 1
		}
	}

	return distancesDeHamming, max
}

func distanceHamming(a, b Exemple, distancesDejaCalculees DistancesHamming) int {
	// Si on a déjà calculé cette valeur, on la retourne
	if distancesDejaCalculees[a.id][b.id] > 0 {
		// -1 car on ne veut pas stocker 0, étant la valeur par défaut.
		// Ca fausserait les comparaisons.
		// On ajoute donc +1 au moment de stocker la valeur et on la soustrait en la récupérant
		// TODO : utiliser un type nullable, ou une structure
		return distancesDejaCalculees[a.id][b.id]
	}

	// Si la valeur n'a pas encore été calculée, on la calcule
	count := 0

	// Pour chaque valeur différente, on incrémente le compteur
	for i, aVal := range a.valeurs {
		if aVal != b.valeurs[i] {
			count++
		}
	}

	// Count + 1 pour éviter les comparaisons avec 0
	return count + 1
}

// On sauvegarde les résultats pour éviter de calculer 2 fois la même chose
type compareResult struct {
	calculated         bool
	min, max           int
	minIndex, maxIndex int
	moy                float32
}

var saveCompare map[Hash]compareResult

// Compare un exemple avec un cluster (interne ou externe)
// et retourne min (distance minimum entre cet exemple et les éléments du cluster),
// max (pareil, maximum) et moy (distance moyenne entre cet élément
// et le cluster cible)
func compareAvecCluster(exemple Exemple, cluster Cluster, distances DistancesHamming) (min, max, minIndex, maxIndex int, moy float32) {
	concat := fmt.Sprintf("cp%#v%#v%#v", exemple, cluster, distances)
	h := hash(concat)
	res := saveCompare[h]

	if res.calculated {
		return res.min, res.max, res.minIndex, res.maxIndex, res.moy
	}

	total := 0
	count := 0
	min = math.MaxInt32
	max = 0

	for i, ex := range cluster {
		if ex.id == exemple.id {
			continue
		}

		count++
		distance := distances[ex.id][exemple.id]
		total += distance

		if distance > max {
			max = distance
			maxIndex = i
		}

		if distance < min {
			min = distance
			minIndex = i
		}
	}

	moy = float32(total) / float32(count)

	saveCompare[h] = compareResult{
		calculated: true,
		min:        min,
		max:        max,
		minIndex:   minIndex,
		maxIndex:   maxIndex,
		moy:        moy,
	}

	return min, max, minIndex, maxIndex, moy
}

// Récupère la distance interne mini, maxi et la moyenne des distances internes d'un cluster
func distancesInternes(cluster Cluster, distances DistancesHamming) (min, max, minIndex, maxIndex int, moy float32) {
	concat := fmt.Sprintf("di%#v%#v", cluster, distances)
	h := hash(concat)
	res := saveCompare[h]

	if res.calculated {
		return res.min, res.max, res.minIndex, res.maxIndex, res.moy
	}

	total := float32(0)
	min = math.MaxInt32
	minMoy := float32(0)
	maxMoy := float32(0)

	for _, exemple := range cluster {
		mi, ma, minI, maxI, mo := compareAvecCluster(exemple, cluster, distances)

		total += mo

		// on sauvegarde le mini avec la distance moyenne mini
		if mi < min || (mi == min && minMoy < mo) {
			min = mi
			minIndex = minI
			minMoy = mo
		}

		// et le maxi avec la distance moyenne maxi
		if ma > max || (ma == max && maxMoy > mo) {
			max = ma
			maxIndex = maxI
			maxMoy = mo
		}
	}

	moy = total / float32(len(cluster))

	saveCompare[h] = compareResult{
		calculated: true,
		min:        min,
		max:        max,
		minIndex:   minIndex,
		maxIndex:   maxIndex,
		moy:        moy,
	}

	return min, max, minIndex, maxIndex, moy
}

func maxDistancesInternes(clusters []Cluster, distances DistancesHamming) (max int, maxIndex int) {
	maxMoy := float32(0)
	for i, cluster := range clusters {
		if _, m, _, _, mo := distancesInternes(cluster, distances); m > max || (m == max && mo > maxMoy) {
			max = m
			maxIndex = i
			maxMoy = mo
		}
	}

	return max, maxIndex
}

func distancesExternes(cluster Cluster, clusters []Cluster, distances DistancesHamming) (min, max, minIndex, maxIndex int, moy float32) {
	concat := fmt.Sprintf("de%#v%#v%#v", cluster, clusters, distances)
	h := hash(concat)
	res := saveCompare[h]

	if res.calculated {
		return res.min, res.max, res.minIndex, res.maxIndex, res.moy
	}

	min = math.MaxInt32
	total := float32(0)
	count := 0
	minMoy := float32(0)
	maxMoy := float32(0)

	for _, exemple := range cluster {
		for i, cls := range clusters {
			// Si c'est le même cluster
			if exemple.id == cls[0].id {
				continue
			}

			mi, ma, _, _, mo := compareAvecCluster(exemple, cls, distances)

			total += mo

			count++

			// on sauvegarde le mini avec la distance moyenne mini
			if mi < min || (mi == min && minMoy < mo) {
				min = mi
				minIndex = i
				minMoy = mo
			}

			// et le maxi avec la distance moyenne maxi
			if ma > max || (ma == max && maxMoy > mo) {
				max = ma
				maxIndex = i
				maxMoy = mo
			}
		}
	}

	moy = total / float32(count)

	saveCompare[h] = compareResult{
		calculated: true,
		min:        min,
		max:        max,
		minIndex:   minIndex,
		maxIndex:   maxIndex,
		moy:        moy,
	}

	return min, max, minIndex, maxIndex, moy
}

func minDistancesExternes(clusters []Cluster, distances DistancesHamming) (min, minIndex int) {
	for _, cluster := range clusters {
		if m, _, minI, _, _ := distancesExternes(cluster, clusters, distances); m > min {
			min = m
			minIndex = minI
		}
	}

	return min, minIndex
}
