# push-swap

Implémentation de **push-swap** et **checker** en Go, utilisant deux piles (`a` et `b`) et un ensemble d'instructions (`pa`, `pb`, `sa`, `sb`, `ss`, `ra`, `rb`, `rr`, `rra`, `rrb`, `rrr`).

## Structure du projet

- `cmd/push-swap` : point d'entrée du programme `push-swap` qui génère la suite d'instructions pour trier la pile `a`.
- `cmd/checker` : point d'entrée du programme `checker` qui lit les instructions sur `stdin`, les exécute et affiche `OK` ou `KO`.
- `internal/pushswap` : parsing des arguments, gestion des piles et des opérations, algorithme de tri et tests unitaires.

## Compilation

Depuis la racine du projet :

```powershell
go build -o push-swap.exe ./cmd/push-swap
go build -o checker.exe ./cmd/checker 
```
## utilisation

```powershell
.\push-swap.exe "2 1 3 6 5 8"
.\push-swap.exe "2 1 3 6 5 8" | .\checker.exe "2 1 3 6 5 8"
```
## Tests
 
```powershell
go test -v ./...
```