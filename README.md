# serrepiGo
Outils de gestion de la serre avec un raspberry pi sous Raspbian

## dependances
Installation de go, par exemple
``
snap install --classic go
``

Based on great work found on :

https://github.com/stianeikeland/go-rpio

Thanks !!


## pinOutLow
Outil ecrit en go permettant de passer l'état d'un port gpio en sortie à l'état bas pour une durée déterminée.

Peut être utile, entre autres, pour activer un relais de manière récurrente (par exemple via crontab
) 

 * Build

Pour builder pour une raspbian, 

```
env GOOS=linux GOARCH=arm GOARM=5 go build -o pinOutHigh
```
 * Paramètres acceptés :
```
./pinOutLow -h

Usage of ./pinOutLow:
  -d duration
    	duration
  -n int
    	BCM Gpio port number
```
 * Usage :
 
 Par exemple l'arrosage de la serre se fait via l'éléctrovanne branché sur le relais commandé par le port GPIO BCM 17
 
 ``
 sudo ./pinOutLow -n 17 -d 60s
 ``
 ouvre le robinet permettant d'arroser la serre pendant 1 minute
 
 
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
 * Usage :
  non averé pour la gestion de la serre
