# Uppgifter som går att göra på playground

Samma uppgifter som i README minus de som behöver argument (behöver en terminal)

Utgå från filen städer.go för följande uppgifter
-----------------------------------------------

Exempel finns på https://tour.golang.org/moretypes/16

  1. Skriv ut namnen på alla städer. En stad per rad.

  2. Skriv ut namnen på alla städer, när programmet skriver ut Åbo, skriv också ut något annat (valfri text)
  
  if-sats: https://tour.golang.org/flowcontrol/5
  
  3. Skriv ut namnen på städerna i alfabetisk ordning

   tips: import "sort"
   

Utgå från filen träd.go för följande uppgifter
---------------------------------------------

Läs:

 - https://tour.golang.org/moretypes/2
 - https://tour.golang.org/moretypes/3
 - https://tour.golang.org/moretypes/5
 - https://tour.golang.org/moretypes/9

för att förstå hur structs och listor (slices) fungerar

Du behöver en rekursiv funktion för att lösa uppgiften. Du kan försöka googla och söka efter en lösning.

  4. Skriv ett program som skriver ut trädet, skriv ut varje node med ett antal mellanslag (" ") eller tabs ("\t")
     som visar att noden ligger på ett visst djup. Du kan också använda - eller *  
     
     Alltså, t.ex.:
     ```
     handel och data:
       - datanom:
         - data18
         - data18st
         - dat17
         - dat17st
         - dat16
       - merkonom:
         - mera18
         - mer17
         - mer16
      ```

  5. Som 4, men skriv ut till höger om namnet hur många studeranden det finns i noden
