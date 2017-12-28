# Distance de Hamming
TP d'Intelligence Artificielle

## Auteur
Thomas Sauvajon

## Installer Go sur Linux

``` sh
# Dernière version stable Linux 64 bits au moment de l'écriture du Readme : go1.9.2.linux-amd64.tar.gz
# Consulter https://golang.org/dl/ pour versions + récentes
VERSION=go1.9.2
OS=linux
ARCH=amd64
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

## GOPATH
Les sources Go doivent se trouver dans le `$GOPATH` (par défaut `$HOME/go`)

Pour le changer :

``` sh
export GOPATH=/nouvelle/localisation/du/gopath
```

Les projets Go se trouvent tous dans le même dossier, et sont organisés
par organisation/username.

Par exemple pour ce repo:
`$HOME/go/src/github.com/tsauvajon/distance-hamming`.

# Lancement
``` sh
go build
./distance-hamming
```

# Utilisation

Pour utiliser le jeu de test de l'énoncé, copier/coller le contenu de
`Saisie.txt` dans la console.

Pour utiliser un autre jeu de test, suivre les consignes du programme !.

# Approche

## Papier + crayon

J'ai commencé par une approche papier avec l'exemple donné (10 exemples,
4 colonnes) :

Tout d'abord, transformation des exemples en true/false, par exemple
clair = true et foncée = false pour rendre l'algorithme plus facile à concevoir.

**Matrice d'exemples :**

|    |  1  |  2  |  3  |  4  |
|----|-----|-----|-----|-----|
| 1  |  T  |  F  |  F  |  T  |
| 2  |  T  |  T  |  F  |  T  |
| 3  |  F  |  F  |  F  |  T  |
| 4  |  T  |  F  |  T  |  T  |
| 5  |  T  |  F  |  F  |  F  |
| 6  |  T  |  T  |  T  |  F  |
| 7  |  F  |  F  |  F  |  F  |
| 8  |  F  |  T  |  T  |  T  |
| 9  |  F  |  T  |  T  |  F  |
| 10 |  F  |  F  |  T  |  F  |

Ensuite, création d'une matrice de 10 x 10 (symétrique) qui représente la
distance de hamming entre chaque paire d'éléments :

**Distances de hamming :**

|   |  1  |  2  |  3  |  4  |  5  |  6  |  7  |  8  |  9  |  10 |
|---|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|
|1  |  -  |  1  |  1  |  1  |  1  |  3  |  2  |  3  |  4  |  3  |
|2  |  1  |  -  |  2  |  2  |  2  |  2  |  3  |  2  |  3  |  4  |
|3  |  1  |  2  |  -  |  2  |  2  |  4  |  1  |  2  |  3  |  2  |
|4  |  1  |  2  |  2  |  -  |  2  |  2  |  3  |  2  |  3  |  2  |
|5  |  1  |  2  |  2  |  2  |  -  |  2  |  1  |  4  |  3  |  2  |
|6  |  3  |  2  |  4  |  2  |  2  |  -  |  3  |  2  |  1  |  2  |
|7  |  2  |  3  |  1  |  3  |  1  |  3  |  -  |  3  |  2  |  1  |
|8  |  3  |  2  |  2  |  2  |  4  |  2  |  3  |  -  |  1  |  2  |
|9  |  4  |  3  |  3  |  3  |  3  |  1  |  2  |  1  |  -  |  1  |
|10 |  3  |  4  |  2  |  2  |  2  |  2  |  1  |  2  |  1  |  -  |

Pour calculer les valeurs, j'ai calculé les différences manuellement entre
2 lignes d'exemples.

Par exemple, la ligne d'exemple 1 et la ligne d'exemple 6 ont une valeur
commune et trois valeurs difféntes.

J'inscris donc 3 comme valeur pour [1, 6] et [6, 1].

Je répète ces étapes jusqu'à ce que ma matrice soit complète.

## Clustering

J'ai ensuite séparé ces données en 2 clusters :
pour celà, j'ai essayé de regrouper au maximum les éléments ayant une distance
faible entre eux (1), et de mettre dans des clusters différents les éléments
ayant une distance forte entre eux [4].

**Distances de hamming :**

|   |  1  |  2  |  3  |  4  |  5  |  6  |  7  |  8  |  9  |  10 |
|---|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|
|1  |  -  |  1  |  1  |  1  |  1  |  3  |  2  |  3  | [4] |  3  |
|2  | (1) |  -  |  2  |  2  |  2  |  2  |  3  |  2  |  3  | [4] |
|3  | (1) |  2  |  -  |  2  |  2  | [4] | (1) |  2  |  3  |  2  |
|4  | (1) |  2  |  2  |  -  |  2  |  2  |  3  |  2  |  3  |  2  |
|5  | (1) |  2  |  2  |  2  |  -  |  2  | (1) | [4] |  3  |  2  |
|6  |  3  |  2  | [4] |  2  |  2  |  -  |  3  |  2  | (1) |  2  |
|7  |  2  |  3  | (1) |  3  | (1) |  3  |  -  |  3  |  2  | (1) |
|8  |  3  |  2  |  2  |  2  | [4] |  2  |  3  |  -  | (1) |  2  |
|9  | [4] |  3  |  3  |  3  |  3  | (1) |  2  | (1) |  -  | (1) |
|10 |  3  | [4] |  2  |  2  |  2  |  2  | (1) |  2  | (1) |  -  |

J'obtiens donc les 2 clusters suivants :

## Cluster 1

**Matrice d'exemples**

|   |  1  |  2  |  3  |  4  |
|---|-----|-----|-----|-----|
|1  |  T  |  F  |  F  |  T  |
|2  |  T  |  T  |  F  |  T  |
|3  |  F  |  F  |  F  |  T  |
|4  |  T  |  F  |  T  |  T  |
|5  |  T  |  F  |  F  |  F  |
|7  |  F  |  F  |  F  |  F  |

**Distances de hamming**

|   |  1  |  2  |  3  |  4  |  5  |  7  |
|---|-----|-----|-----|-----|-----|-----|
|1  |  -  |  1  |  1  |  1  |  1  |  2  |
|2  |  1  |  -  |  2  |  2  |  2  |  3  |
|3  |  1  |  2  |  -  |  2  |  2  |  1  |
|4  |  1  |  2  |  2  |  -  |  2  |  3  |
|5  |  1  |  2  |  2  |  2  |  -  |  1  |
|7  |  2  |  3  |  1  |  3  |  1  |  -  |

## Cluster 2

**Matrice d'exemples**

|   |  6  |  8  |  9  |  10 |
|---|-----|-----|-----|-----|
|6  |  T  |  T  |  T  |  F  |
|8  |  F  |  T  |  T  |  T  |
|9  |  F  |  T  |  T  |  F  |
|10 |  F  |  F  |  T  |  F  |

**Distances de hamming**

|   |  1  |  2  |  3  |  4  |
|---|-----|-----|-----|-----|
|1  |  -  |  2  |  1  |  2  |
|2  |  2  |  -  |  1  |  2  |
|3  |  1  |  1  |  -  |  1  |
|4  |  2  |  2  |  1  |  -  |

## Conclusions

On constate qu'avec ces 2 clusters, la solution n'est pas satisfaisante :
l'exemple 7 a une distance de 3 avec l'exemple 3 (distance interne),
alors qu'il a une distance de 1 avec l'exemple 10 (distance externe).

Il est impossible d'obtenir une solution satisfaisante avec seulement 2 clusters,
et l'élément 7 donne des solutions assez similaires et mitigées s'il est placé
soit dans le cluster 1, soit dans le cluster 2.

Une solution pourrait être atteinte avec plus de clusters, mais il serait très
long et difficile de les obtenir manuellement. Un algorithme s'impose donc.

# Algorithme

## En langage naturel :

N initial: 2

- 1 On démarre en splittant aléatoirement les exemples en N clusters (au moins 2 éléments par cluster) => si impossible, on retourne les derniers clusters obtenus
- 2 On vérifie si les conditions sont remplies => Oui = fini
  Conditions = toutes les distances internes <= toutes les distances externes
- 3 On sépare les exemples qui ont une trop grande distance entre eux (1 transfert d'un cluster vers un autre)
  Le transfert se fait vers le cluster qui a la distance moyenne la plus faible avec l'élément à transférer
- 4 Si on a déjà eu cette configuration => on augmente le nombre de clusters et retour à l'étape **#1**
- 5 Si un des clusters a 1 seul élément => "transvaser" un élément d'un autre cluster (celui avec la + grosse distance moyenne) et retour à l'étape **#4**
- 6 Retour à l'étape **#2**

## Conditions d'arrêt :

- toutes les distances internes <= les distances externes (#2) -> résultat satisfaisant
- on ne peut pas séparer en N clusters -> résultat partiel (peut être ou ne pas être un résultat "optimal")

## En Go :

**Syntaxe de Go**

Pour information, en Go, `varA := 2` déclare une variable varA (et détermine
le type), `varA = 2` assigne une valeur à la variable existante `varA`.

En Go, les `switch` `break` par défaut. Il faut explicitement indiquer `fallthrough`
pour exécuter l'instruction du `case` suivant.

**Algo**

TODO

## Complexité :

TODO

# Remarques

Lancer 2 fois l'algorithme avec les mêmes entrées ne fournira pas toujours
les mêmes sorties.

Selon les entrées, il n'est pas certain qu'il existe de solution viable :
dans ce cas, l'algorithme retourne une solution qui réduit au maximum les
distances internes, et qui souvent comprendra un nombre élevé de clusters
(nombre de clusters maximum : nombre d'exemples / 2).

# Limitaions

Le système de hashes stockés dans une map a ses limitations : il existe un risque
de collision. Dans ce cas, l'algo sortira prématuremment. Sachant qu'un uint32 va
de 0 à 4294967295, le risque de collisions avec 10 exemples reste assez limité.

Cet algorithme ne fonctionnerait pas avec des valeurs non binaires (plus
de 2 valeurs possibles pour un même critère, faible / moyen / fort par exemple).

# Ouverture

La complexité de cet algorithme pourrait largement être améliorée, et ne
permet pas de traiter de problèmes avec une grosse masse de données.

A la suite du codage de cet algorithme, le langage Go semble bien adapté
à ce type de problèmes, au même titre que le Python, en comparaison à des
langages orientés Objet (Java, C#).

L'algorithme cherche à réduire les distances internes, mais pas à maximiser
les distances externes. L'approche inverse pourrait également fonctionner.

L'algorithme cherche également à réduire les distances internes en se
préoccupant peu du nombre de clusters ; selon le but recherché, il
faudrait mettre + l'accent sur la diminution du nombre de clusters.

Evolution possible : lancer un grand nombre de fois l'algorithme avec les
mêmes entrées, comparer les résultats et garder la meilleure solution
recontrée.

# Collaboration

J'ai discuté sur la phase "Crayon + papier" avec :

- Guillaume BAECHLER
- Loïc Thiaw-Wing-Kaï
- Floriane Ziégelé

J'ai échangé sur nos algorithmes respectifs avec Théophile Cousin, et 
étudié son code sur https://github.com/theocousin/ai_hamming
