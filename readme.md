🐹 TP1 – Implémentation d’un cache manuel en Go

Ce premier travail pratique consiste à implémenter un système de cache en mémoire manuellement, sans dépendance externe.
L’objectif est de comprendre le fonctionnement d’un cache avant d’utiliser des bibliothèques comme GoCache.

🎯 Objectifs pédagogiques

Comprendre la gestion de la mémoire et la synchronisation en Go

Implémenter un cache à durée de vie limitée (TTL)

Utiliser les mutex (sync.RWMutex) pour assurer la sécurité des accès concurrents

Tester la logique d’expiration et de récupération des données

🧰 Technologies utilisées

Langage : Go 1.22+

Packages standard utilisés :

sync → gestion de la concurrence

time → gestion du TTL

fmt → affichage console

🧩 Fonctionnement

Le cache stocke des paires clé/valeur avec une date d’expiration (TTL).
Lorsqu’un élément expire, il est automatiquement ignoré lors des lectures.

🐹 TP2  Go — Utilisation de GoCache

Ce projet est un travail pratique (TP) réalisé en Go (Golang).
L’objectif est d’implémenter un système de mise en cache performant en utilisant la bibliothèque gocache
.

🚀 Objectifs du TP 2

Découvrir la gestion du cache en Go

Manipuler des structures de données en mémoire

Améliorer les performances d’accès aux données

Mettre en pratique la gestion de la durée de vie (TTL) des objets en cache

🧰 Technologies utilisées

Langage : Go 1.25+

Librairie : patrickmn/go-cache

Outils : go mod, go run, go test