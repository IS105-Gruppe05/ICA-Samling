# ICA07

Navnliste: Abdikani Gureye, Brede Knutsen Meli, Eirik Aanestad Fintland, Jan Kevin Henriksen, Mats Skjærvik, Mikael Kimerud, Morten Johan Mygland, Nils Fredrik Iselvmo Bjerk.

- I denne oppgaven har alle bidratt. Ikke alle jobber på hver sin maskin, vi jobber av og til bare på noen av pcene og det er da de som laster opp det som er gjort. Derfor vil det ikke alltid være at alle har pushet opp noe til github.

Koden i dette repositoriet er modifisert og/eller inspirert fra flere kilder:


http://stackoverflow.com/a/26032240

https://austburn.me/posts/creating_a_secure_server_in_golang

http://www.minaandrawos.com/2016/05/14/udp-vs-tcp-in-golang/

http://www.sohrabvakharia.in/practical-3-diffie-hellman-key-exchange-algorithm/

https://gist.github.com/manishtpatel/8222606


## 1 a og b)

UDP klient og server kjøres lokalt. Det er nødvendig å opprette to terminaler for å få kjørt prosessen. En til server og en til klient. Outputen i server vil da være meldingen fra klient "Møte Fr 5.5 14:45 Flåklypa".

```
go run udpserver.go
```
```
go run udpclient.go
```

UDP klient og server kjøres lokalt. Det er nødvendig å opprette to terminaler for å få kjørt prosessen. En til server og en til klient. Outputen i server vil da være meldingen fra klient "Møte Fr 5.5 14:45 Flåklypa".

```
go run udpserver.go
```
```
go run udpclient.go
```

!![Alt Text](https://raw.githubusercontent.com/IS105-Gruppe05/ICA07/master/Bilder/Bilde1.png)

## c)


![Alt Text](https://github.com/IS105-Gruppe05/ICA07/blob/master/Bilder/Bilde2.png)

![Alt Text](https://github.com/IS105-Gruppe05/ICA07/blob/master/Bilder/Bilde3.png?raw=true)










## i 1)
Med wireshark kan vi se hvor mange bytes som sendes over udp med meldingen vår "Møte Fr 5.5 14:45 Flåklypa". Vi fanget opp totalt antall bytes71, hvor headeren opptar 20 bytes. Videre i ser vi at total lengde er på 57.

![Alt Text](https://github.com/IS105-Gruppe05/ICA07/blob/master/Bilder/Bilde4.png?raw=true)

## i 2)
Mange ulike tjenester begrenser de største UDP pakkene til rundt 1472(Trekke ifra UDP og IP header på henholdsvis 8 og 20bytes.). Det er mulig å gå over ca 1472 bytes men da risikerer man packet loss og andre problemer har større sannsynlighet for å oppstå. Når vi referrer til det største en UDP pakke kan være, så tar vi i betrakning at pakkene brytes opp i biter pga Ethernet 2s MTU(Maximum transmission unit) begrensning.. Dette er noe vi også gjør for TCP pakker, som kommer lenger ned.
## ii: Over NIC.
## ii 1)

![Alt text](https://i.gyazo.com/432f797006229363553698bbb662e67f.png)

<<<<<<< HEAD
Her kan vi se prosenten på protokollen på UDP og TCP pakker.
=======
Her kan vi se prosenten på protokollen på UDP og TCP pakker. 
>>>>>>> dc10d6e2b053d5ce086b6dccc3bfaaa8c11906ec

## ii 2)

I wireshark kan vi filtrere etter protocoll for å finne de pakkene vi er interisert i. Det er mulig å kombinere filterkrav med hjelp av opperatører som ```&&``` (and) og ```||``` (or).

```ip.proto == "TCP"```
```ip.proto == "UDP"```


## ii 3)
















## Oppgave 2 (TCP)
## a)

TCP server og klient kjører lokalt og trenger derfor to terminalvinduer for å utføre prosessen. Man kan skrive valgfrie meldinger som input i klient som sendes til server.

```
go run tcp_server.go
```
```
go run tcp_client.go
```

TCP server og klient kjører lokalt og trenger derfor to terminalvinduer for å utføre prosessen. Man kan skrive valgfrie meldinger som input i klient som sendes til server.

```
go run tcp_server.go
```
```
go run tcp_client.go
```

![Alt Text](https://github.com/IS105-Gruppe05/ICA07/blob/master/Bilder/Bilde5.png?raw=true)

## b)

## i)
De største forskjellene mellom UDP og TCP er at UDP ikke trenger tilkobling, det vil si at når data eller meldinger sendes, vet man ikke om den kommer frem eller blir tapt på veien mellom sender og mottaker. TCP er tilkoblings-orientert, som vil si at når en melding blir sendt så vil meldingen komme frem så lenge tilkoblingen ikke blir brutt. Hvis tilkoblingen skulle vise seg å bli brutt så vil serveren fremdeles spørre etter den tapte delen.

Når det gjelder rekkefølge av meldinger, så vil man med TCP ikke måtte bekymre seg over at data kan forekomme i feil rekkefølge. med UDP så vet man ikke hvilken melding som ble sendt ut først, hvis man for eksempel sender ut to meldinger.

Om et program bruker TCP eller UDP er opp til utviklerne, og kommer an på hva programmet trenger. Ønsker du feil-korrigering og en sterk tilkobling så er TCP det bedre valget. UDP brukes der hastighet er viktig, og nettverket/ operativsystemet ikke trenger å gjøre mye for å oversette data som kommer fra pakkene.


## ii)
En TCP pakke kan være på rundt 1500 bytes når man sender pakker over nettverk (Ethernet2 MTU). Blir den større enn det, kan det by på flere problemer som tap av data, pakke forkortelse osv. Det er som regel lurt å være under den øverste grensen på 1500, siden man må regne med data for Ip/TCP header ( 20 bytes hver ca)  osv som kan gjøre at du kommer over toppen. 1400 bytes før man legger til bytes for IP og TCP header, gjør at man kommer under grensen der problemer lettere kan oppstå. Hvis man er over 1500 bytes kan en løsning være å dele opp i to pakker i stedet.

NB: Nå skrev jeg pakker som jo eksisterer for TCP, men at man ikke kan se de. TCP håndterer data i form av segmenter.

## iii)
Fragmentering oppstår når lagringsplass ikke brukes effektivt, og reduserer da både kapasitet og opptreden. Hvordan fragmentering oppstår, avhenger av hvordan datamaskinen du bruker lagrer filer. Hvis en fil som behøver å lagres er for stor til å passe de ledige plassene i driven, som for eksempel en stor filmfil på mange GB. Når dette skjer, så må den delen av operativsystemet som kontrollerer hvordan filer lagres ta en avgjørelse om hvor filen skal lagres, og noen ganger er den beste avgjørelsen og dele filen opp i mange mindre deler. Man kan løse fragmenteringsproblemer ved å sette til side plass om det skulle trengs. (Pre-Allocation).

## iv)
TCP: Sending og mottak av email, SMS service, musikkstreaming.

UDP:  - Voice over IP. Eksempel: Du bruker Skype/Viber og har en videokonferanse med kollegaer. Du ender opp med å miste pakker, men kan fortsatt høre vedkommende selv om det stopper/hakker litt i samtalen. Det samme eksemplet gjelder også for video spill online og streaming av filmer.


## Oppgave 3:

## a)

Både UDP og TCP klientene sender en kryptert melding over til sine representative servere. Opperasjonen foregår lokalt og det må derfor benyttes to terminaler.

```
go run udpserver.go
```
```
go run udpclient.go
```

```
go run tcpserver.go
```
```
go run tcpclient.go
```

## b)

<<<<<<< HEAD
Vi fikk ikke til å implementere Diffie-Hellman nøkkelutveksling for klient-server i Golang, og greide heller ikke å forstå hvordan verktøy som NaCL fungerte, så vi valgte å se etter implementasjoner i Java og prøve å reversere det til Golang-kode. Dette gikk heller ikke fordi vi ikke fant noe tilsvarende java.math.BigInteger i Golang. Vi prøvde å implementere primtall-sortering i Golang, men valgte å stoppe når vi kom til primitive  røtter av primtall når vi innså at det ville blitt vår egen implementering av nøkkelutveksling, som er utenfor våre egenskaper på dette tidspunktet.
Vi har vedlagt inputtest.go som sorterer primtall, og javakode i diffie-hellman.java og en jar-fil som viser en lokal form av Diffie-Hellman.

diffie-hellman.jar kjøres via

```
java -jar diffie-hellman.jar
```

og inputtest.go kan kjøres via

```
go run inputtest.go
```
=======
Vi fikk ikke til å implementere Diffie-Hellman nøkkelutveksling for klient-server i Golang, og greide heller ikke å forstå hvordan verktøy som NaCL fungerte, så vi valgte å se etter implementasjoner i Java og prøve å reversere det til Golang-kode. Dette gikk heller ikke fordi vi ikke fant noe tilsvarende java.math.BigInteger i Golang. Vi prøvde å implementere primtall-sortering i Golang, men valgte å stoppe når vi kom til primitive  røtter av primtall når vi innså at det ville blitt vår egen implementering av nøkkelutveksling, som er utenfor våre egenskaper på dette tidspunktet. 
Vi har vedlagt inputtest.go som sorterer primtall, og javakode i diffie-hellman.java og en jar-fil som viser en lokal form av Diffie-Hellman. 
>>>>>>> dc10d6e2b053d5ce086b6dccc3bfaaa8c11906ec
