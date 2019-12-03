# PRR-Labo2

> Tiago Povoa et Burgener François

## Lancement du projet

Nous avons 3 manières de lancer notre projet. Via deux script, windows et linux, ou alors via ligne de commande

### Windows

Pour lancer le script il faut aller dans le dossier labo2 ``PRR-Labo2/labo2`` et de lancer le script ``startWindows.bat``

```
$ ./startWindows.bat <nombre de processus>
...
Appuyez sur une touche pour continuer...
```

L'argument est le nombre de processus que l'on souhaite lancer. Ensuite appuyer sur une touche pour qu'un processus se lancer. Une fois que le processus a terminé de exécuter appuyer de-nouveau sur une touche pour lancer le prochain processus. Fait cela pour tout les processus.

### Linux

Pour utiliser le script sur linux, il vous faut avoir le terminal gnome-terminal. Si vous ne l'avais pas, vous pourrait lancer chacun des processus via ligne de commande.

Pour lancer le script il faut aller dans le dossier labo2 ``PRR-Labo2/labo2`` et de lancer le script ``startLinux.sh``

```
$ ./startLinux.sh <nombre de processus>
```

L'argument est le nombre de processus que l'on souhaite lancer.

### Ligne de commande

Pour lancer en ligne de commande il faudra tout d'abords aller dans le dossier ``PRR-Labo2/labo2`` et exécuter la ligne suivante dans différent terminal

```
go run main.go -proc <id du processus> -N <nombre de processus>
```

Les id des processus commencent à **0**

**Example**

```
go run main.go -proc 0 -N 3
go run main.go -proc 1 -N 3
go run main.go -proc 2 -N 3
```

## ...



