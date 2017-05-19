# is105-ICA01

Navnliste: 
Abdikani Gureye, 
Brede Knutsen Meli, 
Eirik Aanestad Fintland, 
Jan Kevin Henriksen, 
Mats Skjærvik, 
Mikael Kimerud, 
Morten Johan Mygland, 
Nils Fredrik Iselvmo Bjerk.

- I denne oppgaven har alle bidratt. Ikke alle jobber på hver sin maskin, vi jobber av og til bare på noen av pcene og det er da de som laster opp det som er gjort. Derfor vil det ikke alltid være at alle har pushet opp noe til github.

## Oppgave 1.2.1

Titall til totall:     
1 = 0001

2 = 0010

5 = 0101

8 = 1000

16 = 0001 0000

256 = 0001 0000 0000

Verdier på de ulike plassene:                 
2^0=1                                                         
2^1=2                                                         
2^2=4                                                         
2^3=8                                                         
2^4=16       
2^5=32    
2^6=64  
2^7=128  
2^8=256   
2^9=512

Totall til titall:  
100 = 1*2^2 + 0*2^1 + 0*2^0 = 4 + 0 + 0 = 4  
1001 = 1*2^3 + 0*2^2 + 0*2^1 + 1*2^0 = 8 + 0 + 0 +1 = 9  
1100110011 = 512 + 256 + 0x128 + 0x64 + 32 + 16 + 0x8 + 0x4 + 2 + 1 = 819   

^10 sifre, derfor vet man det må hvertfall være 2^9

Når man skal konvertere fra totallsystemet til titallsystemet, så bruker man det faktum at det kun brukes to sifre i totallsystemet, 0 og 1. 0 og 1 representerer da ulike verdier, basert på hvilken posisjon de har. Ut ifra dette bruker man grunntallet 2, og opphøyer dette i sifferets posisjon nummer og multipliserer dette med sifferverdien.  10 i totallsystemet blir da f.eks: (1 * 2^1) + (0 * 2^0) = 2 + 0 = 2 i titallsystemet. 

Fra titallsystemet til totallsystemet bruker man det samme prinsippet med at alle sifrene har posisjoner med forskjellige verdier, bare motsatt vei. Og da må man regne ut slik at de ulike verdiene til sammen blir 10. Over har vi laget en liten tabell som skriver hvilke verdier de ulike posisjonene i totallsystemet har. Ut ifra den kan man se at man trenger 1 på 2^3 plassen (for det blir 8, ”åtterplass”) og 1 i 2^1 plassen (det blir 2, ”toerplass”). De andre plassene vil da få 0. Altså blir 10 i titallsystemet da f.eks: 1010 i totallsystemet. (1 * 2^3) + (0 * 2^2) + (1 * 2^1) + (0 * 2^0) = 8 + 0 + 2 + 0 = 10. 

## Oppgave 1.2.2

3-bit = 8 muligheter (tall, 0-7). (000, 001, 010, 011, 100, 101, 110, 111).

(1) 4 Oddetall (1, 3, 5 og 7)
Vi har en formel som sier noe om hvordan man skal regne informasjonsmengden basert på antall sannsynlige valgmuligheter, og også hvis man får mer informasjon om et valg som begrenser disse valgmulighetene igjen. Formelen er log2(1 / (M/N) ), der N er antall valgmuligheter (8, alle valgmulighetene, 0-7), og M antall informasjon som begrenser valgmuligheter(4, fordi det er fire oddetall) N. Ved bruk av denne formelen vil svaret bli:  log2(1/(4/8)) = log2(2) = 1 bit.

(2) Ikke multiplum av 3 (0, 3, 6). Da gjenstår 1, 2, 4, 5, 7.
Log2(1/(5/8)) = 0.678072 bit. 

(3) Tallet inneholder nøyaktig to enere: 011, 101, 110
Log2(1/(3/8)) =  ~1.4 bit.

(4) Ut i fra opplysningene i oppgave 1,2 og 3, så vet vi at det eneste tallet som passer alle beskrivelsene er  5. Da blir formelen:
Log2(1/(1/8)) = 3 bit. Hun har altså all den informasjonen hun trenger for å vite tallet.

Jo mer informasjon (bit) en person har, jo større sannsynlighet har de for å gjette riktig tall.


## Oppgave 1.2.3

Begynte med å opprette et repository fra GitHub og klonet dette ved bruk av
```
~ $ git clone <URL til repository>.
```

Så laget vi filen "Hello.go", og brukte
```
~ $ git add Hello.go
```
og
```
~ $ git status
```
for å se om filen ble lagt til i staging. Deretter
```
~ $ git commit -m "melding"
```
og
```
~ $ git push origin master.
```
for å opplaste den til GitHub.

Hvis flere brukere ønsker å "pushe" endringer av samme linje så vil endringene bli "merged" av den siste brukeren 
som "pusher". Serveren "merger" altså da ikke endringene, men brukerne.

Det kan også oppstå en rekke situasjoner dersom begge endrer på samme linje i en fil:

- 1) Brukeren som ikke enda har "committet" sin endring, vil måtte "stashe" sine endringer eller "merge" endringene når
han "puller" fra server.
- 2) Den mest nylig endringen er den som blir "committet".


## Oppgave 1.2.4

(1) Hvilken fordeler og ulemper har en git-flow-modell med en hovedrepository?

Link til hovedrepository: https://github.com/IS105-Gruppe05/ICA01 
Fordeler: Parallell jobbing fungerer godt, enkelte kan gjøre endringer, så kan andre 
i gruppen sjekke om jobber er gjort bra. Det er også lettere å gå tilbake å gjør endringer
hvis det er noe som ikke fungerer. Samtlige har også mulighet til å kun gjøre endringer i sin egen branch, slik at man unngår å overskrive andres arbeid. Gjorde mye av det samme som i 1.2.3 når vi først hadde fått tilgang til hovedrepositoren,
bare at vi gjorde endringer i egen branch (som ble laget ved git checkout -b <branchnavn>), og sendte enn pull request og merget med master. (git merge <branchnavn>).
Ulemper: Hvis man ikke er flink til å update/merge branchen sin ofte, så kan det bli
mye og uoversiktlig. Også hvis flere gjør endringer samtidig så kan det bli komplikasjoner.


(2) Finn ut hva heter objektfiler for de mest brukte platformer (Unix/Linux, MS Windows, Mac OS X)! Hvorfor, etter deres mening, har disse platformene så forskjellige objektfil-formater?

En objekt-fil er en fil som inneholder objekt kode, som betyr at kompilatoren oppretter en objekt fil for hver kilde-fil før den setter dem sammen til noe kjørbart. I begynnelsen var det vanlig at hver type datamaskin hadde sitt eget unike format. Senere kom Unix, og andre operativ systemer som kan brukes I andre systemer enn det som den var laget for. Dette førte til at formatter som COFF og ELF nå brukes på ulike typer systemer. Systemene har også ulike objective formatter fordi de har ulike måter å lese og kalle på binære tallkoder.

Oversikt over objektfiler for de forskjellige platformene:
Windows: PE, DLL
Mac: Mach-O
Linux: ELF (avhengig av build)


(3) Hvilke forskjeller ser dere i forhold til programmeringsspråket Java?

Forskjellene når vi kjører kode I Golang og Java, er at Golang kompilatoren oversetter kode til objektfil (varierer med operativsystem) direkte, mens Java kompilatoren først oversetter til java bytecode slik at den kan utføres av et operativsystem, og deretter oversettes til det aktuelle objektfil-formatet.

“Golang leser koden en linje av gangen” Når vi f eks, kjører kommandoen “ Go run <programmet vårt.go> til et format som operativsystemet kan utføre. Operativsystemet har så et program som kan lese den kompilerte koden linje for linje.

Forskjeller mellom Golang og Java

-	i Golang har vi ikke nøkkelord som Public, og Private. Funksjoner har stor forbokstav.
-	Kompilerer raskere fordi Java bruker VM for å kjøre sin kode som gjør at det tar lengre tid enn for Go.
-	Ingen semikolon nødvendig for Golang. Du kan bruke dem, men du må ikke, sammenlignet med Java.
-	Når det kommer til variabeldeklarasjon, så krever Java at man spesifiserer type for variabelen, mens I Golang så antar kompilatoren typen hvis den ikke er definert på forhånd, f.eks. vil t := 2 bli antatt som en type av int(som kan være enten 32 eller 64bits-størrelse I følge godoc.)
-	Golang er ikke objekt-orientert, men har visse egenskaper som man kan finne I andre objekt-orienterte spark som Java, C++ osv. F.eks, Structs tilsvarer klasser med felt.



(4) I mappen Log finner dere en main.go fil som importerer en funksjon fra log.go (finnes i mappen log). Den beregner logartimen av base 2 til et tall som dere bestemmer. F.eks. hvis dere vil regne ut logaritmen av base 2 til 10:
```
~ $ go run main.go 10
```
![Alt Bilde](https://i.gyazo.com/247d897c28e584464b71ee76be7ae7a8.png)

Hvilke viktige poeng illustrerer denne øvelsen når det gjelder bruk av et programmeringsmiljø på en platform?

Programmering på en platform har mange fordeler, der i blant versjonskontroll, og effektivitet. Vi fant også ut at mappestruktur er veldig viktig for å kunne importere en selvlaget pakke, i dette tilfellet “package log”, som ble importert i main.go via “import ./log”. En forutsetning for at dette skal kunne gå er at filen som det importeres fra, må være i en undermappe i den mappen som main filen ligger i.

Versjonskontroll trenger ikke nødvendigvis å være plattformuavhengig enda, siden de fleste systemer er basert på unix, så er det vel enklere å kunne tilby Gitbash som en løsning for windows baserte brukere. På spørsmål om det bør være plattformavhengig, så vil det nok I fremtiden komme ulike operativsystemer som ikke bare tilhører unix og da kan det kanskje være mer relevant og utvikle versjonskontroll som kan være plattformuavhengig.


Et  eksempel på effektivitet knyttet til versjonskontroll  kan være at du har to personer som jobber med debugging/ utvikling av en ny versjon.
Du når et punkt der du ønsker å gi ut programvaren du har utviklet, men du må debugge/ teste. I mellomtiden jobber din kollega på neste versjon av programvaren. Så I dette eksempelet så har du eldre kode og kode under utvikling separert som er en effektiv måte man kan jobbe på.

GOPATH bruker vi for å vise til steder vi kan se etter Go kode. Den brukes også til å hente, bygge og installere pakker utenfor det som er standard I Go. 


(5) I mappen logcli finnes en logcli.go fil hvor også en kan velge et tall, og regne dette tallets logartime av base 2.
Men i denne så ligger all koden i samme, slik at det er lettere å builde. Vi har valgt å ikke legge med noen kjørbar fil, siden det er lettere for hver enkel å bare builde selv, slik at de for et program som kan kjøres på sitt eget operativsystem.

```
~ $ go build logcli.go 
```
Da har man buildet et kjørbart program, og kan kjøres ved denne kommandoen, hvis man igjen vil bruke 10 som eksempel:

```
~ $ ./logcli 10
```

Er det hensiktsmessig å legge inn denne filen i git repository? Begrunn svaret!

Nei! Den eneste hensikten med å ha et utførbart program i Repositoryen er for å dele den med resten av gruppen. MEN, siden alle har tilgang til .go filene vil alle kunne builde sine egne programmer, så mye av hensikten faller egentlig bort.
i tillegg så vil man måtte “builde” forskjellige filer for forskjellige operativsystemer.

![Alt Bilde](https://i.gyazo.com/17151739e0314b4df73f68ce781979ff.png)

(6) I mappen logbcli finnes en logbcli.go fil som er veldig like logcli.go filen, bare at i denne versjonen kan man også velge base, hvor basen er det første argumentet. Så hvis man vil regne logaritmen av base 10 til 5, må man først builde logbcli.go filen, og deretter kjøre programmet:

```
~ $ ./logbcli 10 5
```

Eller hvis man vil regne logaritmen av base 2 til 3:

```
~ $ ./logbcli 2 3
```
![Alt Bilde](https://i.gyazo.com/f162c6d25edf1c2ff4e835c66a2c9d07.png)

Hvordan skiller pakken log​, som dere har implementert, seg fra andre pakker i go, som, for eksempel, fmt​? 


Log er en lokal pakke som vi har skrevet selv, og vi har full kontroll over funksjonaliteten til denne pakken. fmt er en del av standardbiblioteket for golang, og vi kan kun bruke funksjonene som definert i fmt dersom vi importerer fmt. 
