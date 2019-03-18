# Webb servers i go

Utgå från webbserver.go

## 2.1. Skriv en webbserver som skriver ut vad tiden är
	
tips: paketet "time" har en funktion Now()

## 2.2. Skriv en webbserver som skriver tillbaka stigen i URLen (det som kommer efter domänen och porten)
	
tips: stigen kommer du åt via http.Request.URL.Path (variabeln r i vårt exempel)

## 2.3. Skriv en webbserver som servar filen index.html

tips: paketet ioutil innehåller en funktion ReadFile läser en fil från filsystemet och returnerar en sträng och ett error (om något gått fel)

exempel:
	```
	//Läser filen fil.txt in i variabeln text
	text,err := ioutil.ReadFile("fil.txt")
	```
	Minns att alla variabler du deklarerar måste användas. Om du inte vill använda errorn kan du skriva _ istället

## 2.4. Skriv en ny endpoint till webbservern från uppgift 3. Enpointen ska heta "/hund" och ska serva filen hund.jpeg

tips: Browsern vet inte vilken filtyp en fil har om det inte står i http-svaret.
Webbservrar använder http-headers för att berätta om för browsern vilken filtyp en fil har 
	
http-headers används för mycket annat också.

http-headern för att bestämma vilken filtyp en fil har heter ```Content-Type```

Content-Typen för en jpeg-fil är ```image/jpeg```

I go kan du ändra på en http-header genom http.ResponseWriterns Header().Set() funktion.

Exempel:

	```
	//använder w som vår ResponseWriter hette i exempelfilen webbserver.go
	//Ändrar filtypen till webm
	w.Header().Set("Content-Type","video/webm")
	```

## 2.5. Skriv en ny webbserver som servar filer från den katalog du kör den från. 

Om du hämtar localhost:8080/a.txt ska du:
	
 - skriva innehållet av filen a.txt om filen finns
 - skriva ett felmeddelande om filen inte finns (använd errorn från ioutil.ReadFile)

Du behöver inte oroa dig om filtyper

tips: använd http.Request.URL.Path igen. Du måste skära av första symbolen ("/") från filstigen.

tips: använd ioutil.ReadFile igen

## 2.6. Finns det säkerhetsluckor i din filserver? Går det att läsa filer som man inte borde kunna läsa? Kan du göra något för att minska på säkerhetsproblemen?
