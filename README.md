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

Les projets Go se trouvent tous dans le même dossier, et sont organisés par organisation/username
Par exemple pour ce repo:
$HOME/go/src/github.com/tsauvajon/distance-hamming

# Lancement
``` sh
go build
./distance-hamming
```

# Utilisation

Pour utiliser le jeu de test de l'énoncé, copier/coller le contenu de `Saisie.txt` dans la console
Pour utiliser un autre jeu de test, suivre les consignes du programme !

# Approche

## Papier + crayon

J'ai commencé par une approche papier avec l'exemple donné (10 exemples, 4 colonnes) :
Tout d'abord, transformation des exemples en true/false, par exemple clair = true et foncée = false pour rendre l'algorithme plus facile à concevoir.

**Matrice d'exemples :**

   |  1  |  2  |  3  |  4  |
1  |  T  |  F  |  F  |  T  |
2  |  T  |  T  |  F  |  T  |
3  |  F  |  F  |  F  |  T  |
4  |  T  |  F  |  T  |  T  |
5  |  T  |  F  |  F  |  F  |
6  |  T  |  T  |  T  |  F  |
7  |  F  |  F  |  F  |  F  |
8  |  F  |  T  |  T  |  T  |
9  |  F  |  T  |  T  |  F  |
10 |  F  |  F  |  T  |  F  |

Ensuite, création d'une matrice de 10 x 10 (symétrique) qui représente la distance de hamming entre chaque paire d'éléments :

**Distances de hamming :**

   |  1  |  2  |  3  |  4  |  5  |  6  |  7  |  8  |  9  |  10 |
1  |  -  |  1  |  1  |  1  |  1  |  3  |  2  |  3  |  4  |  3  |
2  |  1  |  -  |  2  |  2  |  2  |  2  |  3  |  2  |  3  |  4  |
3  |  1  |  2  |  -  |  2  |  2  |  4  |  1  |  2  |  3  |  2  |
4  |  1  |  2  |  2  |  -  |  2  |  2  |  3  |  2  |  3  |  2  |
5  |  1  |  2  |  2  |  2  |  -  |  2  |  1  |  4  |  3  |  2  |
6  |  3  |  2  |  4  |  2  |  2  |  -  |  3  |  2  |  1  |  2  |
7  |  2  |  3  |  1  |  3  |  1  |  3  |  -  |  3  |  2  |  1  |
8  |  3  |  2  |  2  |  2  |  4  |  2  |  3  |  -  |  1  |  2  |
9  |  4  |  3  |  3  |  3  |  3  |  1  |  2  |  1  |  -  |  1  |
10 |  3  |  4  |  2  |  2  |  2  |  2  |  1  |  2  |  1  |  -  |

Pour calculer les valeurs, j'ai calculé les différences manuellement entre 2 lignes d'exemples.

Par exemple, la ligne d'exemple 1 et la ligne d'exemple 6 ont une valeur commune et trois valeurs difféntes.

J'inscris donc 3 comme valeur pour [1, 6] et [6, 1]

## Clustering

J'ai ensuite séparé ces données en 2 clusters :
pour celà, j'ai essayé de regrouper au maximum les éléments ayant une distance faible entre eux (1), et
de mettre dans des clusters différents les éléments ayant une distance forte entre eux [4]

**Distances de hamming :**

   |  1  |  2  |  3  |  4  |  5  |  6  |  7  |  8  |  9  |  10 |
1  |  -  |  1  |  1  |  1  |  1  |  3  |  2  |  3  | [4] |  3  |
2  | (1) |  -  |  2  |  2  |  2  |  2  |  3  |  2  |  3  | [4] |
3  | (1) |  2  |  -  |  2  |  2  | [4] | (1) |  2  |  3  |  2  |
4  | (1) |  2  |  2  |  -  |  2  |  2  |  3  |  2  |  3  |  2  |
5  | (1) |  2  |  2  |  2  |  -  |  2  | (1) | [4] |  3  |  2  |
6  |  3  |  2  | [4] |  2  |  2  |  -  |  3  |  2  | (1) |  2  |
7  |  2  |  3  | (1) |  3  | (1) |  3  |  -  |  3  |  2  | (1) |
8  |  3  |  2  |  2  |  2  | [4] |  2  |  3  |  -  | (1) |  2  |
9  | [4] |  3  |  3  |  3  |  3  | (1) |  2  | (1) |  -  | (1) |
10 |  3  | [4] |  2  |  2  |  2  |  2  | (1) |  2  | (1) |  -  |

J'obtiens donc les 2 clusters suivants :

## Cluster 1

**Matrice d'exemples**

   |  1  |  2  |  3  |  4  |
1  |  T  |  F  |  F  |  T  |
2  |  T  |  T  |  F  |  T  |
3  |  F  |  F  |  F  |  T  |
4  |  T  |  F  |  T  |  T  |
5  |  T  |  F  |  F  |  F  |
7  |  F  |  F  |  F  |  F  |

**Distances de hamming**

   |  1  |  2  |  3  |  4  |  5  |  7  |
1  |  -  |  1  |  1  |  1  |  1  |  2  |
2  |  1  |  -  |  2  |  2  |  2  |  3  |
3  |  1  |  2  |  -  |  2  |  2  |  1  |
4  |  1  |  2  |  2  |  -  |  2  |  3  |
5  |  1  |  2  |  2  |  2  |  -  |  1  |
7  |  2  |  3  |  1  |  3  |  1  |  -  |

## Cluster 2

**Matrice d'exemples**

   |  6  |  8  |  9  |  10 |
6  |  T  |  T  |  T  |  F  |
8  |  F  |  T  |  T  |  T  |
9  |  F  |  T  |  T  |  F  |
10 |  F  |  F  |  T  |  F  |

**Distances de hamming**

## Conclusions

On constate qu'avec ces 2 clusters, la solution n'est pas satisfaisante :
l'exemple 7 a une distance de 3 avec l'exemple 3 (distance interne),
alors qu'il a une distance de 1 avec l'exemple 10 (distance externe).

Il est impossible d'obtenir une solution satisfaisante avec seulement 2 clusters,
et l'élément 7 donne des solutions assez similaires et mitigées s'il est placé
soit dans le cluster 1, soit dans le cluster 2.

Une solution pourrait être atteinte avec plus de clusters, mais il serait très long
et difficile de les obtenir manuellement. Un algorithme s'impose donc.

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

TODO

## Complexité :

TODO

# Remarques

Lancer 2 fois l'algorithme avec les mêmes entrées ne fournira pas toujours les mêmes sorties

Pour information, en Go, `varA := 2` déclare une variable varA (et détermine le type), `varA = 2` assigne une valeur à la variable existante `varA`

# Ouverture

Cet algorithme ne fonctionnerait pas avec des valeurs non binaires (plus de 2 valeurs possibles pour un même critère, faible / moyen / fort par exemple).