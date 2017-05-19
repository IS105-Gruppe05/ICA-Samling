# ICA05

Navnliste: Abdikani Gureye, Brede Knutsen Meli, Eirik Aanestad Fintland, Jan Kevin Henriksen, Mats Skjærvik, Mikael Kimerud, Morten Johan Mygland, Nils Fredrik Iselvmo Bjerk.

- I denne oppgaven har alle bidratt. Ikke alle jobber på hver sin maskin, vi jobber av og til bare på noen av pcene og det er da de som laster opp det som er gjort. Derfor vil det ikke alltid være at alle har pushet opp noe til github. 

Kildekode for serveren finnes i mappen Oppgave 1, server.go

Brukerscenario:
Per skal ha påskeferie om noen få dager, og ha skal på tur oppi fjellet. Dermed er det viktig å vite hva slags vær det vil bli oppe i fjellet i påske. Så han bestemmer seg for å finne det ut med å åpne nettsiden til en av hans favoritte værtjeneste [guttegruppa-agder.go](http://158.37.63.91:8001). Han åpner nettsiden og skriver inn hvor han skal være i påske. Etter å ha trykket på knappen “søk” kommer det en ny side opp med all informasjon han trenger. Informasjon om hvor varmt/kaldt, hvor sterk/svak vind og hvilket vær det kommer til å være oppe fjellet i påske. Så nå var Per godt fornøyd med informasjons han hadde fått fra [guttegruppa-agder.go](http://158.37.63.91:8001) og kunne glede seg til påske. 


Brukerscenario 2: 
Per skal på påskeferie oppe i fjellet hvis det blir fint vær i påske. Så han bestemmer seg for å finne det ut med å åpne sin favoritt værtjeneste [guttegruppa-agder.go](http://158.37.63.91:8001) for å finne ut hvordan været blir. Han åpner siden og skriver inn søkestedet sitt i søkefeltet og trykker søk. På siden kommer det opp hvor varmt/kaldt det blir, hvor sterk/svak vind det blir og hvilket vær det blir i påske. Da ser han fort at i påske blir det kaldt, men det kommer til å blåse mye og melder bare regn, så da frister det lite med å reise på påskeferie. Så han tar den riktige valget med å bli hjemme å ha en gaming påske. Da slår han opp siden [guttegruppa-agder.go](http://158.37.63.91:8001), der han logger på med brukeren sin. Derfra går han inn på vennelista, for å finner ut hvor mange av vennene hans er online slik han kan snakke med de og spille med dem.

Skjermbilde fra nettside:

![Bilde1](https://i.gyazo.com/13e3d35744d85fa1bd11de2388c33f19.png)


![Bilde2](https://i.gyazo.com/20ef26c2f2fd8ff413031871741f95d1.png)

Under json mappen finnes en main.go funksjon som bruker goroutines. Den henter inn værinformasjon fra 5 forskjellige nettsider.
