# ICA02

Navnliste: Abdikani Gureye, Brede Knutsen Meli, Eirik Aanestad Fintland, Jan Kevin Henriksen, Mats Skjærvik, Mikael Kimerud, Morten Johan Mygland, Nils Fredrik Iselvmo Bjerk.

- I denne oppgaven har alle bidratt. Ikke alle jobber på hver sin maskin, vi jobber av og til bare på noen av pcene og det er da de som laster opp det som er gjort. Derfor vil det ikke alltid være at alle har pushet opp noe til github.


## Oppgave 1:

Specs ligger inne i docs (linken under).
https://docs.google.com/spreadsheets/d/1TUlQXl-SlkGzoc0ZrM6NQQLayWV9iOdv5WyQhnWOXKE/edit#gid=0 

Oppstart: POST →  BIOS/EFI
Prosesser: CPU, primærminne, dvs. ram og sekundærminne, dvs. HDD/SDD.

Maskinen starter ved å gjennomføre en power-on self test (POST), og deretter går videre til oppstart programmer i firmware på hovedkortet. Basic Input/Output System (BIOS) er en fellesbetegnelse for slike programmer. Det er vanlig å inkludere Extended Firmware Interface (EFI) etter 2010.

En prosess må ha tilgang til prosessor, primær- og sekundær minne(for å kunne overleve uten strøm) for å kjøre. Hver plattform har sin egen arkitektur for administrasjon av prosesser. OSX har /sbin/launchd prosess med ID 1, Linux /sbin/init (kan variere for ulike distros) med ID 1. OSX og Linux (fordi de er baserte på unix) har begge et hierarkisk system for prosesser. NT Windows har ntoskrnl.exe, og bruker ikke et direkte hierarkisk system på samme måte som OSX og Linux. 

I Windows kan man bruke taskmgr for å ha oversikt over hvilke prosesser som kjører og hvilke som ligger klar til bruk (idle). Når en prosess befinner seg “idle” kjøres en software med lavt aktivitet nivå, slik at prosessen er klar til å startes når den trengs. Man kan lukke ikke-kritiske prosesser via taskmgr. I Linux og OSX kan man bruke top, ps og  for å ha oversikt for prosesser. 

Prosesser:
Både linux serveren, windows og mac varierer prosessene som kjøres. Når du starter et program i windows vil nødvendige prosesser bli tatt i bruk. Under oppstart av et program vil flere prosesser bli brukt, for så å gå i idle tilstand igjen. Det varierer dog hvor mange prosesser hvert system benytter for programmer. Når det gjelder linux serveren virker den å være mer stabil enn windows, men varierer fremdeles litt.


## Oppgave 2:

![Alt Text](https://github.com/IS105-Gruppe05/ICA02/blob/master/bilder/Oppgave%202/Skjermbilde%202017-05-12%20kl.%2013.35.36.png?raw=trueg)

![Alt Text](https://github.com/IS105-Gruppe05/ICA02/blob/master/bilder/Oppgave%202/Skjermbilde%202017-05-12%20kl.%2013.36.28.png)

Viser her at vi kan kjøre hello.exe på windows.
```
Kjør med "go run Hello.go"
```

![Alt Text](https://github.com/IS105-Gruppe05/ICA02/blob/master/bilder/Oppgave%202/cmd%20windows.png)



## Oppgave 3:
I oppgave 3 er det laget to mapper for test. Sum_fail hvor alle testene feiler, og sum hvor alle testene godkjennes.
For å kjøre en test må du være lokalisert i mappen sum_fail / sum.
```
go test
```
For å kjøre "kalkulatoren" hvor vi legger sammen to tall av int 64.
```
go run main_sum.go
```


Vi lagde Sum funksjoner for de forskjellige numeriske typene i golang (int8, int32, int64, float64 og uint32). Deretter lagde vi tilsvarende tester (sumType_test.go) for de forskjellige typene og lagde en rekkevidde av tall for hver type, som vi skulle teste.

Vi har laget tester som produserer owerflow feil, samt tester som går gjennom. For int8 lagde vi en test som gikk utover rekkevidden. {1267, 1, 127} (kan leses som tall + tall = tall). UInt32 betyr unassigned integer, som vil si at typen kun aksepterer positive verdier. For å teste og produsere en feil for uint32 lagde vi en 128, -1, 127. Dette gir en overflow feil da uint ikke aksepterer negative tall.

De andre typene (int32, int64 og float64) har mye større rekkevidde og oppfører seg mer eller mindre likt. Forskjellen er at float er for desimaltall, hvor 64 er veldig stor presisjon og 32 tar mindre plass enn 64. I følge golang dokumentasjonen er det ikke mulig å mikse numeriske typer i golang. 

I main_sum.go lagde vi et program som aksepterer 2 int64 parametere og regner ut summen via terminal, syntaks go run main_sum.go tall1 tall2 vil gi sum av tall1 og tall2.

En mulig løsning for å gjøre pakken “sum” trygg og brukervennlig vil være å f.eks vise til at man kun kan bruke positive tall som en feilmelding hvis de prøver å skrive negative tall for typen uint32. Samt produsere logiske feilmeldinger til bruker.


## Oppgave 4:

For å kjøre benchmark må du være lokalisert i mappen oppgave4.
```
go test -bench=.
```

Big O, er en måte å måle hvor effektiv en algoritme er og hvor godt den skalerer I forhold til størrelsen på innholdet. Det er vanlig å se bort fra best case, og heller fokusere på worst case noe vi også gjør I denne oppgaven.

Bubble sort algoritmen som vi modifiserte har en O(n^2) algoritme fordi vi går gjennom lista og sammenligner hvert element med det ved siden av, og bytter elementene hvis de ikke står I riktig rekkefølge. Etter første iterasjonen, begynner vi så fra begynnelsen men til forskjell fra den første iterasjonen, så vet vi nå at det siste elementet I lista er korrekt.

Quicksort algoritmene i likhet med den originale bubble sort algoritmen du la inn i sorting_go bruker også en O(n2) algoritme.
Bubble_sort_mod2 benytter seg av en boolean “swap” funksjon for å konstatere at elementene i lista er blitt byttet. Mod2 bytter elementene når den finner ut hvilken som er størst. Både bubble_sort og bubble_sort_modified lager en midlertidig liste hvor et element lagres før de byttes. Mod2 vil derfor være mest effektiv av disse funksjonene.

Bilde over benchmark fra bubblesort funksjonene.
![Alt Text](https://github.com/IS105-Gruppe05/ICA02/blob/master/bilder/Oppgave%204/18450035_444169509276108_1144329361_n.png)


## Oppgave 5:
For å kjøre alle filene må det benyttes tre terminaler. Lokalt på maskinen benytte "top / ps" for å se prosessene. På Ubuntu serveren er det brukt "htop".
```
go run boring_main.go
```
```
main_boring_goroutine.go
```


Vi fant ut at boring_main.go programmet opptok 5 tråder på Mac og 4 på Ubuntu. Programmet var en “running process”, dvs. at programmet kjørte i realtime.
Du stopper prosessen med å trykke: ctrl + C. 

Med kommandoen “tasklist” i cmd på windows kan vi se hvor mange tråder og minne boring.exe og main_boring_goroutine.exe bruker. Vi kan se at programmet benytter 3 tråder i windows. Samme resultat oppstår også i PowerShell med kommandoen get-process.

Bilder av tråder når vi kjører både boring_main og main_boring_goroutine:

mac:
![Alt Text](https://github.com/IS105-Gruppe05/ICA02/blob/master/bilder/Oppgave%205/Skjermbilde%202017-05-15%20kl.%2011.34.14.png)

Windows:

![Alt Text](https://github.com/IS105-Gruppe05/ICA02/blob/master/bilder/Oppgave%205/18492732_120332000804713234_240020223_n.png)

Ubuntu server:
![Alt Text](https://raw.github.com/IS105-Gruppe05/ICA02/master/bilder/Oppgave%205/18516655_120332000775565656_1898529906_o.png)
