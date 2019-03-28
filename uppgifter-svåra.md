## 1. I rekursion.go finns en simpel rekursiv struktur. En funktion som kallar på sig själv.

En rekursiv funktioninnehåller:

 - ett basfall, något som stoppar den. I vårt exempel ```if len(indentering) == 10```
 - ett kall till den rekursiva funktionen (ett kall till sig själv)

Utgående från detta, skriv en rekursiv funktion som räknar ut fibonaccis talföljd, som slutar efter x gånger (du kan ge x in som ett argument).

```
Fibonaccis talföljd är nästa tal summan av de två förra talen:

0
1  = 0
1  = 1+0
2  = 1+1
3  = 2+1
5  = 3+2
8  = 5+3
13 = 8+5
21 = 13+8
34 = 21+13
etc.

```

## 2. Skriv om förra uppgiften ```iterativt```: alltså som en vanlig for-loop

## 3. En ```struct``` är en samling av fält i ett värde. Lite som ett objekt i python.

T.ex.

```
type Person struct{
	Namn string
	Ålder int
}
```

Definierar en struct som går att skapa på följande sätt:

```
var p = Person{"Kalle",12}
```

Vi kan efter det komma åt personens namn:

```
fmt.Printf("%v är %v år gammal",p.Namn,p.Ålder)
```

### Skriv en struct som heter ```Vektor``` som innehåller två tal, kalla dem x och y

Testa din struct genom att skapa en Vektor och printa ut x och y

Du kan testa exemplet med Person först, sedan göra uppgiften.

## 4. Skriv en funktion som adderar två Vektorer

Vektorer adderas genom att addera varje del skiljt. x:n ska adderas och y:n ska adderas. Resultatet är en ny Vektor.

## 5. Utgå från träd.go:

Skriv ett program som skriver ut trädet, skriv ut varje node med ett antal mellanslag (" ") eller tabs ("\t") som visar att noden ligger på ett visst djup.

     Alltså, t.ex.:
     ```
     handel och data:
      datanom:
       data18
       data18st
       dat17
       dat17st
       dat16
      merkonom:
       mera18
       mer17
       mer16
      ```

Tips: Du kan använda rekursion.go för inspiration.

## 6 Använd dlv för att printa en stack trace

### Använd debuggern delve för att printa en stack-trace när du kollar på studentgruppen dat17

#### Instruktioner

Installera programmet delve:

```
go get -u github.com/go-delve/delve/cmd/dlv
```

Om din GOPATH inte är del av din PATH kommer du inte att kunna köra programmet.

Om du inte har en GOPATH måste du skapa den.

Kolla om du har GOPATH med env|grep GOPATH

Om du inte får ett svar måste du skriva in din GOPATH i din .bashrc:

```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Du får också använda en annan GOPATH än $HOME/go, men då måste du flytta din go-katalog till den stigen.

Det kan vara att du måste vara i en katalog med enbart en .go-fil
Starta debuggern med 
```
dlv debug
```

Sätt en breakpoint med 

```
b $fil:$radnummer
```

Starta programmet med:

```
continue
```

Du får en stack-trace med 

```
stack
```
