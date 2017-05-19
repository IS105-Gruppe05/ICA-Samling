# ICA03

Navnliste: Abdikani Gureye, Brede Knutsen Meli, Eirik Aanestad Fintland, Jan Kevin Henriksen, Mats Skjærvik, Mikael Kimerud, Morten Johan Mygland, Nils Fredrik Iselvmo Bjerk.
- I denne oppgaven har alle bidratt. Ikke alle jobber på hver sin maskin, vi jobber av og til bare på noen av pcene og det er da de som laster opp det som er gjort. Derfor vil det ikke alltid være at alle har pushet opp noe til github.


Repository for IS-105 experiments with encoding of signs and symbols.

## Oppgave 1: 
## a)

```
go run ascii_main.go
```

![Bilde1](https://i.gyazo.com/7a8bcacfacde867cd4e9cf5c352a68fc.png)

Ved å sjekke tabellen på ascii-code.com fant vi ut at de første 31 karakterene i ASCII-tabellen er kontroll-karakterer som ikke kan printes ut. Et eksempel er #24 i lista, hex 18 som er Cancel.
![Bilde2](https://i.gyazo.com/f0177cd8b6d6bcea6349d63eb33bba9b.png)

Her er bilde av de 16 første karaktere i ASCII tabellen. 

## b)

![Bilde3](https://i.gyazo.com/83e515b1437e4428055e6ce9979cd1f9.png)

Her kan vi se at tegn som “ og mellomrom også har sine egne ASCII-verdier, som vi kan bruke for å hardkode en setning. I dette tilfellet er det standard ASCII, som vil si at vi ikke trenger å encode setningen noe mer utover dette. 

## c) 

```
go test
```
Hver karakter har en verdi. I standard ASCII er dette fra 1-127, og utvidet ASCII er over 127. Vi lagde en test som kontrollerer om verdien på karakterene er under 127, siden alle karakterene i funksjonen er i standard ASCII så består testen.
VI har brukt “range” koden slik den kan finne ut av hvor den er i lista og lagrer den lista den har gått igjennom så langt. Hadde vi bare brukt en for-løkke ville koden ikke visst hvor den var  i lista. 

![Bilde4](https://i.gyazo.com/06b39d9d8780ab306edb79cb4832a91d.png)

Vi prøvde å sette inn tegn som ikke var ascii, men da fikk vi bare tilbakemelding om at dette var “Feil”, og at det ikke var riktige ascii tegn. Det er en feilkode som vi selv har kodet inn for å hjelpe oss i å forstå hva som  er feil i koden.



## Oppgave 2:
## a)

```
go run main.go
```

![Bilde5](https://i.gyazo.com/cf301c09caba65e055dadee85481e29f.png)

![Bilde6](https://i.gyazo.com/08db71ed88d2dad1763351fc91faaa6c.png)

Her kan vi se forskjellen på forskjellig encoding i selve terminalen, i det øverste bildet bruker terminalen (Git Bash i dette tilfellet) CP-1252 encoding, mens det nederste bruker UTF-8 encoding. Vi kan se at 1252 tillater tegn som ½, mens UTF-8 printer de fleste utvidete ASCII-karakterene, men ikke ½.
Platformene spiller ikke en stor rolle for vi må å passe på hvilken type terminal vi bruker og hvilken encoding terminalen bruker. 

## b)

![Bilde7](https://i.gyazo.com/67f8c5605785dc466810064826d798ac.png)

Ubuntu: Se bildene over
Vi fant ikke ut av hvordan vi skulle få til €(euro tegnet). Vi prøvde ulike innstillinger inne på terminalen der som gjorde at vi fikk opp € tegnet, men da var det andre tegn som ikke virket i stedet for. I encoding CP1252 fikk vi opp €.

ISO-8859-1 kodesett er mer begrenset i forhold til UTF-8 som brukes mer i Golang og andre programmeringsmiljøer. Vesten bruker UTF-8 encoding for det meste og det er kun gamlere programmer som ikke klarer å håndtere utf-8 kodesett, og disse bruker da generelt ascii eller annet kodesett som deres operativsystem støtter.

## c)
For å kunne utføre testen må du være lokalisert i iso mappen.
```
go test
```

Hver karakter har en verdi. I standard ASCII er dette fra 1-127, og utvidet ASCII er over 127. Vi brukte samme test som vi brukte på oppgave 1c, testen ser om verdien til alle tegnene er over 127. På grunn av at alle karakterene ikke er i extended ASCII så blir testen ikke godkjent.



## Oppgave 3: 


## a)

```
go run formatering.go
```

%s formateringen returner en “uninterpreted” byte av stringen eller slicen i UTF-8 (Go-standard), og de symbolene som ikke er i den første ASCII-tegnsettet returneres som “?”.
Med %q lar Go de tegnene den ikke forstår bare være (“xbd” etc).
%x returner hexadesimal. Stor X gjør at det returneres i store bokstaver. 
%c leser en og en karakter og tolker og skriver ut dette.

Så vi ser på ASCII tabellen at bd er ½. Så for å få dette til å fungere med %s, så må den gjøres til UTF-8. Den er \xC2\xBD, og hvis man setter inn dette i koden blir det slikt:

![Bilde8](https://i.gyazo.com/582ae218bd06878bbda692380c40a1f6.png)


## b) 

```
go run fileutils_main.go
```

![Bilde9](https://i.gyazo.com/bef76ce664b684e48cd7ade04a1a3805.png)
![Bilde10](https://i.gyazo.com/93c89e54abc36f23d23677d14fdb2aad.png)
![Bilde11](https://i.gyazo.com/27fc841c241a144bb78200e8aca7b043.png)
![Bilde12](https://i.gyazo.com/a702a55ed0a0ffe9eee069636b1fb17c.png)

Vi kan se at lang01 er kyrillisk(vi fikk ikke til riktig output i terminal), lang02 er islandsk og lang03 er norsk. På mac fikk vi til å skrive ut ordene, men ikke riktige bytes. Så vi trur det kan ha noe å gjøre med enkodinga.

De første 16 bytes tilsvarer de første 16 karakterene i printen, dvs. FE FD 73 6B 61 72 0A FE FD 73 6B 61 72 61 6E 61 tilsvarer þ ý s k a r þ ý s k a r a n a þ ý.

## c)

```
go run treasure_main.go
```

Treasure stringen blir gjort om til UTF-8, slik at vi får ut tekst som gir mening. 

![Bilde13](https://i.gyazo.com/023de41bf40f235e95fbba72b4f3a6ce.png)

## Oppgave 4:
## a) 
Koden tar parameter med js / jp
```
go run unicode_main.go js
```
```
go run unicode_main.go jp
```

I dette tilfellet bruker vi heksadesimal og unicode.
Heksadesimaler for: “nord og sør”.
 Norsk = "\x93\x6e\x6f\x72\x64\x20\x6f\x67\x20\x73\xf8\x72"
 Islansk = "\u0022\x6e\x6f\x72\u00f0\x75\x72\x20\x6f\x20\x73\x75\u00f0\x75\x72\u0022"
 Japansk = "\u0022\xe5\x8c\x97\xE3\x81\xA8\xE5\x8D\x97\u0022"

Det går også an å representere karakterer via desimalverdien deres, som f.eks. “ som blir 34 i desimalverdi. Hvis vi vil printe ut 34 som desimal kan vi først definere 34 som en variabel: x:= 34
og deretter kan vi skrive: fmt.Printf(“%c”, x) som vil printe ut “.

![Bilde14](http://i.imgur.com/dwX7FDe.png)

## b)
```
go run server.go
```
Bildet viser tiden da det ble tatt.

![Bilde15](https://github.com/IS105-Gruppe05/ICA03/blob/master/Bilder/Oppgave4/Bilde15.png)

