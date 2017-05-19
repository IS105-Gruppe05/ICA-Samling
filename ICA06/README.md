# ICA06
ICA06 Gruppe05 
Navnliste: Abdikani Gureye, Brede Knutsen Meli, Eirik Aanestad Fintland, Jan Kevin Henriksen, Mats Skjærvik, Mikael Kimerud, Morten Johan Mygland, Nils Fredrik Iselvmo Bjerk.

- I denne oppgaven har alle bidratt. Ikke alle jobber på hver sin maskin, vi jobber av og til bare på noen av pcene og det er da de som laster opp det som er gjort. Derfor vil det ikke alltid være at alle har pushet opp noe til github.

For at disse skal virke må du bruke chrome browser.

Eksperiment 1.

Vi har sett på hvordan formantene i ordet "hjelp" ser ut. Måten vi gjorde dette på var at alle
på gruppen spilte inn ordet i programmet praat. Vi sammenlignet deretter bildene og så da
spesielt på F1 og F2 linjene. Siden disse linjene er mest presise og nok for å utelukke feil. I
eksperimentet så vi likhet på bildene med tanke på frekvensene for hver enkelt bokstav. Når
det kommer til hvilken metode man burde bruke for å klassifisere lyd til tekst er det flere
mulige valg. Det å gjenkjenne hele ord en en mulighet vi ville foretrukket. En utfordring med
hele ord er forskjellige dialekter. Noen uttaler ordet “hjelp” mens andre sier “hjælp”, dette er
noe som må tas hensyn til. Vi så da at “æ” var en lavere frekvens kontra “e”.

Eksperiment 2.

main.go i servermappa kan kjøres for å starte en lokal server(localhost:8080) som deretter
sender en request til http: //158.37.63.148:8080 som kjører espeak via espeakbox. main
benytter eksterne golang-pakker(https://github.com/meinside/wit.ai-go
https://github.com/nicolaifsf/go-speak) for å kjøre, og de kan installeres med go get -u
URLforGitHubRepoHer hvis GOPATH har blitt satt opp.

Fungerende tekst til tale: http://158.37.63.148:8080/speech?text=hei&voice=no

Eksperiment 3.

Vi har fått til å kjøre wit.ai i et go-program(main_ai.go i wit.ai-go-mappa), men vi har ikke fått
til filopplastning via nettleser som deretter kan leses av wit API for tolkning.

![Bilde1](https://i.gyazo.com/5e479a1aed9f95dc1130a30be5e0ee6c.png)
Fungerende nettside: http://158.37.63.22:8085
