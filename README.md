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


## pinOut
Outil ecrit en go permettant de passer l'état d'un port gpio en sortie à l'état souhaité pour une durée déterminée.

Peut être utile, entre autres, pour activer un relais de manière récurrente (par exemple via crontab
) 

 * Build

Pour builder pour une raspbian, 

```
cd pinOut
env GOOS=linux GOARCH=arm GOARM=5 go build -o pinOu
```
 * Paramètres acceptés :
```
./pinOut -h

Usage of ./pinOut:
  -d duration
    	duration
  -n int
    	BCM Gpio port number
  -s int
    	GPIO state:  - 0 : low  - 1 : high

```
 * Usage :
 
 Par exemple l'arrosage de la serre se fait via l'éléctrovanne branché sur le relais commandé par le port GPIO BCM 17.
 Pour que le robinet soit ouvert l'état du port 17 doit être à 0
 
 ``
 sudo ./pinOutLow -n 17 -d 60s -s 0
 ``
 ouvre le robinet permettant d'arroser la serre pendant 1 minute
 