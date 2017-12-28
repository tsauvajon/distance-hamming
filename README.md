# Distance de Hamming
TP d'Intelligence Artificielle

## Auteur
Thomas Sauvajon

## Lancement
``` sh
go build
./distance-hamming
```

Pour utiliser le jeu de test de l'énoncé, copier le contenu de Saisie.txt dans la console
Pour utiliser un autre jeu de test, suivre les consignes du programme !

## Algorithme

N initial: 2

- 1 On démarre en splittant aléatoirement les exemples en N clusters (au moins 2 éléments par cluster) => si impossible, on retourne les derniers clusters obtenus
- 2 On vérifie si les conditions sont remplies => Oui = fini
  Conditions = toutes les distances internes <= toutes les distances externes
- 3 On sépare les exemples qui ont une trop grande distance entre eux (1 transfert d'un cluster vers un autre)
  Le transfert se fait vers le cluster qui a la distance moyenne la plus faible avec l'élément à transférer
- 4 Si on a déjà eu cette configuration => on augmente le nombre de clusters et retour à l'étape #1
- 5 Si un des clusters a 1 seul élément => "transvaser" un élément d'un autre cluster (celui avec la + grosse distance moyenne) et retour à l'étape #4
- 6 Retour à l'étape #2

Conditions d'arrêt :

- toutes les distances internes <= les distances externes (#2) -> résultat satisfaisant
- on ne peut pas séparer en N clusters -> résultat partiel (peut être ou ne pas être un résultat "optimal")

## Remarques

Lancer 2 fois l'algorithme avec les mêmes entrées ne fournira pas toujours les mêmes sorties
