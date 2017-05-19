# is105-ica04




Navnliste: Abdikani Gureye, Brede Knutsen Meli, Eirik Aanestad Fintland, Jan Kevin Henriksen, Mats Skjærvik, Mikael Kimerud, Morten Johan Mygland, Nils Fredrik Iselvmo Bjerk.

- I denne oppgaven har alle bidratt. Ikke alle jobber på hver sin maskin, vi jobber av og til bare på noen av pcene og det er da de som laster opp det som er gjort. Derfor vil det ikke alltid være at alle har pushet opp noe til github.

## Oppgave 1
## a)

![Alt Text](https://raw.github.com/IS105-Gruppe05/ICA04/master/Bilder/Oppgave%201a.png)

(Bildet til venstre)
54 65 73 74 65 72 20 6C 69 6E 6A 65 73 6B 69 66 74 2E 0A 4F 67 20 65 6E 20 74 69 6C 20 2E 2E 2E 0A 4F 67 20 65 6E 20 74 69 6C 20 2E 2E 2E 0A 4F 67 20 65 6E 20 74 69 6C 20 2E 2E 2E 0A 
[T e s t e r   l i n j e s k i f t . 
 O g   e n   t i l   . . . 
 O g   e n   t i l   . . . 
 O g   e n   t i l   . . . 
 
( Bildet til høyre)
54 65 73 74 65 72 20 6C 69 6E 6A 65 73 6B 69 66 74 2E 0D 0A 4F 67 20 65 6E 20 74 69 6C 20 2E 2E 2E 0D 0A 4F 67 20 65 6E 20 74 69 6C 20 2E 2E 2E #D 0A 4F 67 20 65 6E 20 74 69 6C 20 2E 2E 2E 0D 0A

[T e s t e  r   l i n j e s k i f t . 

 O g   e n   t i l   . . . 
 
 O g   e n   t i l   . . . 
 
 O g   e n   t i l   . . . 
 
 
Operativsystem har ingenting å si på hvordan en fil blir vist på skjermen. Det er applikasjonene. Hvis man lager noe i notepad, kan de endre line shift etc. Windows systemer kan legge på character/carriage return.
Hvis man skriver ut byte for byte kan man se hva som egentlig skjuler seg i teksten.


Som vi kan se i text1.txt er byten 0D lagt til 4 ganger, noe som forklarer hvorfor den ene filen er 4 bytes større enn den andre.. 0D er carriage return, 0A er linjeskift. Carriage return resetter enheten til å begynne på starten av ny linje. 


## b)
Koden kjøres med en innlagt fil (asdf.txt) i programmet.
```
go run main.go
```

Linjeskift m/ punktum, mac: Vi ser tegnet 0A  fordi tekstfiler opprettet i unix og nyere mac versjoner kun bruker line feed, mens de i de første mac versjonene kun brukte  carriage return (0D) i tillegg til 0A.

Bilde for linjeskift i filen på mac: 
![Alt Text](https://github.com/IS105-Gruppe05/ICA04/blob/master/Bilder/1.1b.png)


Bilde for linjeskift windows: Dette er samme fil som i illustrasjonen over. Siden filen er opprettet i windows sin notepad, så vil tegnene 0D og 0A dukke opp når en trykker på enter tasten. 0A står for line feed, mens 0D står for carriage return. Vi har brukt en funksjon for å telle antall “\x0A” i filen, som vi får ut som tekst. Dette viser oss at det er 4 linjeskift i teksten.
![Alt Text](https://github.com/IS105-Gruppe05/ICA04/blob/master/Bilder/1.2b.png)



## Oppgave 2
## a)
For å kjøre filinfo filen benyttes flag. Filen tar parameter av ei fil, i vårt tilfelle asdf.txt.
```
filinfo.go -f asdf.txt
```
Her kan vi detaljert data om filen og se hvilke rettigheter som tilhører.
![Alt Text](https://github.com/IS105-Gruppe05/ICA04/blob/master/Bilder/2a.png)

## b)
Programmet kjøres i skyinstansen (ubuntu serveren).
```
go run filinfo -f /dev/stdin
```
```
go run filinfo -f dev/ram0
```


![Alt Text](https://github.com/IS105-Gruppe05/ICA04/blob/master/Bilder/2b.png)

/dev/stdin
Dette er en device file/unix device file som har Unix permissions bits: Dcrw--w---- som tilsvarer Owner: read write, Group: write, Others: ingen tilgang. D står for “door”, c står for “device character”.
/dev/ram0
Dette er en devicefile men ikke en Unix character device file. Filen har en Unix permission bits: Drw-rw---- som tilsvarer Owner: read write, Group: read write, Others: ingen tilgang. D står for “door”.

## c)
Kommando for å skjekke fileinfo i cmd, windows.
```
go run fileinfo.go -f mappenavn
```

Bilde av mappe info på windows:
![Alt Text](https://github.com/IS105-Gruppe05/ICA04/blob/master/Bilder/2.2c.png)

Ut i fra dette skjermbildet kan vi se at selve mappen ikke har en størrelse i følge filinfo, selv om mappen inneholdt en fil på 58B. Dette vil sannsynligvis si at win10 ikke teller selve mappen som en fil, men heller innholdet av mappen. Vi testet dette videre ved å lage en tom mappe på win10 og OSX, og kunne se at den tomme mappen vises som 0 bytes på win10, men 68 bytes på OSX. Dette er fordi OSX bygger på unix-prinsippet at alt er en fil, mens Windows ikke gjør dette.

## Oppgave 3
## a)

Det finnes en rekke ulike metoder for å jobbe med filer i GO programmer. Ved hjelp av pakken “os” kan vi lese filer og redigere filer. Det nødvendig å sjekke for feil, når det er gjort så kan man minske en fils innhold  i minnet. Det er vanlig å ville styre hvordan og hvilke deler av en fil som skal leses, og da kan man enten lese stream av bytes med buffer eller uten buffer. Man kan også bruke Seek funksjon for å finne frem til et spesifikt punkt i filen og lese derfra, og vi kan velge å lese et gitt antall bytes.

Vi kan bruke “Bufio” pakken til å lage en buffer skriver som gjør at vi kan jobbe med en buffer i minnet før vi skriver det over til disk. dette er praktisk hvis man ikke ønsker å bruke mye tid på diskens IO. Det er også nyttig hvis man vil skrive en byte om gangen og deretter lagre det i minne før man dumper alt til filen på èn gang, fremfor å skrive IO for hver byte. Vanlig buffer størrelse er på 4096 bytes, og minimum er på 16 bytes. 

## b)
Vi har laget tre metoder for å lese fra fil. Vi har muligheten til å lese et gitt antall bytes, Finne alle runes og linjeskift i en fil og lese med hjelp av buffer. Metodene krever parameter i form av tekst filer.

Metoden vil skrive ut et gitt antall bytes av en fil.
```
go run freqnbytes_main.go pg100.txt
```
Metoden finner alle runes og linjeskift i en fil.
```
go run finnbokstav_main.go middels.txt
```
Metoden leser med hjelp av buffer.
```
go run freqbuffer_main.go
```

## c)
Vi har laget benchmark tester for lesing av filer. Vi tester lesing med buffer, et gitt antall bytes og totalt antall runer i filen. For å kunne bechmarke må du være lokalisert i riktig mappe som filene (den filen du vil teste).

Kommando for å kjøre benchmark tester.
```
go test -bench=.
```

![Alt Text](https://github.com/IS105-Gruppe05/ICA04/blob/master/Bilder/3c.png)

Som vi ser på bildene vil det være raskere å lese filer med hjelp av buffer. Om vi ser på funksjonen finnBokstav ser vi at på store filer bruker den veldig lang tid. Dette fordi denne leser gjennom hele filen og ikke bare en viss kapasitet eller antall bytes. Dette resulterer i at funksjonen bare benchmarker stor fil 20 ganger.



## Oppgave 4
## a)

Antall fakulteter: 6
Antall studenter(2014): 
Helse og idrett:			1829
Humaniora og pedagogikk:		1525
Kunstfag:				420
Teknologi og realfag:			2166
Lærerutdanning:			1506
Økonomi og samfunnsvitenskap:	3093
totalt = 				10´539

Tallene under er rundet av til nærmeste hele prosent
Helse og idrett:			1829 = 17%
Humaniora og pedagogikk:		1525 = 14%
Kunstfag:				420 = 4%
Teknologi og realfag:			2166 = 21%
Lærerutdanning:			1506 = 14%
Økonomi og samfunnsvitenskap:    	3093 = 30%
total prosent: 				100%

På bakgrunn av tallene ovenfor kan man for eksempel komme frem til at det er 17% sannsynlighet at en tilfeldig student tilhører Helse og idrett.


## b)
*Når du lærer (får informasjon) om at en tilfeldig valgt student hører til et spesifikt fakultet, for hvilket fakultet får du MINST informasjon?*

Kunstfaget, er det fakultetet man får minst informasjon fra (4%). Eller får du minst informasjon fra det største fakultetet, altså økonomi??

## c)
 Binært tre for huffmankode nedenfor med kodelengde for hvert fakultet.

![Alt Text](https://github.com/IS105-Gruppe05/ICA04/blob/master/Bilder/4.1c.png)

![Alt Text](https://github.com/IS105-Gruppe05/ICA04/blob/master/Bilder/4.2c.png)

Helse og idrett               A:001
Humaniora og pedagogikk       B:000
Kunstfag 			                 	C:111
Teknologi og realfag		       	D:10
Lærerutdanning 	        		    E:110
Økonomi og samfunnsvitenskap 	F:01

## d)
*Hva er gjennomsnittslengden for en melding som inneholder fakultets koder for 100 tilfeldig valgte studenter? Beregn og forklar.*

Hvis vi følger sannsynligheten, blir antallet studenter, eller bokstaver:
17A, 14B, 4C, 21D, 14E, 30F (fant ut hvor stor antall % fra hvert fakultet tidligere).
AKA: 17x”001”, 14x”000”, 4x”111”, 21x”10”, 14x”110”, 30x”01”(Antalletxbits, så A=17x3 bits)
AKA: 51+42+12+42+42+60
AKA: 249 (bits)

## e)

Vi greide ikke å lage kode som kan både dekode og enkode huffman, men via litt "hjemmesnekring" så har vi fått til et go-program som dekoder en hardkoda string av binære tall. Vi valgte å bruke scanner for å lese hver rune, sjekker om det er 0 eller 1, deretter til neste og sjekker om den er 0 eller 1, osv.

```
go run dekoder.go
```
