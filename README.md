# serrepiGo
Outils de gestion de la serre


## pinOutHigh
Outil ecrit en go permettant de passer l'état d'un port gpio en sortie à l'état haut pour une durée déterminée.

Peut être utile, entre autres, pour activer un relais de manière récurrente (par exemple via crontab
)
 * Build

Pour builder pour une raspbian, après installation de go et de https://github.com/stianeikeland/go-rpio

```
env GOOS=linux GOARCH=arm GOARM=5 go build -o pinOutHigh
```
 * Paramètres acceptés :
```
./pinOutHigh -h

Usage of ./pinOutHigh:
  -d duration
    	duration
  -n int
    	BCM Gpio port number
```
