# Gleitkommazahlen: Mantisse, Exponent und Rundungsfehler in Go

Die Seite erklärt, wie eine Gleitkommazahl aus Vorzeichen, Mantisse und Exponent besteht und warum beim Speichern und beim Rechnen Fehler entstehen. Sie rechnet mit dem Modell des Kurstexts, einer Mantisse aus 8 Bit, und zeigt am Ende einen Vergleich mit `float64`. Sie gehört zu Abschnitt 4.4 des Kurstexts.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Die Seite schreibt Zahlen im Text und im Modell mit Komma, zum Beispiel 0,67. In Schritt 3 und in Go-Code steht der Punkt, zum Beispiel 0.1.
- Das Modell: Die Mantisse hat 8 Bit und beginnt mit 0,1. Sie liegt also zwischen 1/2 und 1. Der Exponent reicht von −12 bis 12. Der Wert ist Mantisse mal 2 hoch Exponent.
- Die acht Bits der Mantisse haben die Wertigkeiten 1/2, 1/4, 1/8, 1/16, 1/32, 1/64, 1/128 und 1/256.
- Das Modell schneidet Stellen ab, die nicht in die Mantisse passen. So rechnen die Beispiele im Kurstext. Go rundet bei `float32` und `float64` zur nächsten darstellbaren Zahl.
- Der Abstand zweier darstellbarer Zahlen ist 2 hoch (Exponent − 8). Bei Exponent 0 ist das 1/256, bei Exponent 7 ist es 0,5.
- Darstellungsfehler: der Unterschied zwischen der eingegebenen und der gespeicherten Zahl. Rundungsfehler: Der Fehler wächst beim Rechnen.
- Die Demos in Schritt 1 und 2 rechnen im Modell und nicht mit Go. Die Demo in Schritt 3 rechnet im Browser mit 64-Bit-Gleitkommazahlen nach IEEE 754, der Darstellung von `float64`.
- Auf der Seite läuft kein Go. Der Student kann keinen eigenen Code eingeben.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.
- Schritt 1 bis 3 beginnen mit einem blauen Kasten „Erst vorhersagen, dann ausprobieren“. Unter jeder Frage steht eine aufklappbare „Auflösung“. Die Antwort ist verdeckt, bis der Student sie aufklappt.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Oben stehen die drei Teile einer Gleitkommazahl: Vorzeichen, Mantisse, Exponent.
- Der Teil „Warum sie Gleitkommazahl heißt“ sagt: Der Exponent verschiebt das Komma der Binärzahl. Exponent 7 macht aus 0,11001000 die Zahl 1100100, also 100.
- Der Teil „Das Modell dieser Seite“ nennt die 8 Bit der Mantisse und die 52 Bit bei `float64`.
- Der Teil „Zwei Fehler“ erklärt Darstellungsfehler und Rundungsfehler.
- Eine Liste nennt drei frühere Stellen des Kurstexts: Abschnitt 2.2, Abschnitt 3.3.3 und Abschnitt 4.3.5.
- Eine zweite Liste sagt, was die drei Schritte der Seite zeigen.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den Seiten „Variablen, Typen und Speicher“ und „Zahlenkreis“ folgen. Beide sind eigene HTML-Seiten.
- In diesem Schritt lässt sich nichts ausprobieren. Die Demos stehen in Schritt 1 bis 3.

### Abschnitte des Kurstexts

- 4.4.2
- 4.4.3

## Schritt 1: Eine Zahl speichern

Reiter: speichern

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit drei Fragen zu 0,5, zu 0,1 und zum Exponenten.
- Die Demo hat oben ein Eingabefeld „Zahl“, den Knopf „speichern“ und fünf Knöpfe mit Beispielen: 2/3, 0,67, 0,1, 0,5 und 100,3.
- Darunter steht ein heller Kasten mit dem Bitmuster. Von links nach rechts: der Knopf für das Vorzeichen, die Mantisse als „0,“ und acht große Bits, der Exponent mit den Knöpfen − und +.
- Über jedem Bit der Mantisse steht seine Wertigkeit, von 1/2 bis 1/256. Ein Bit mit dem Wert 1 ist grün unterlegt. Ein Klick auf ein Bit schaltet es um. Das erste Bit ist immer 1 und lässt sich nicht umschalten.
- Unter den Bits stehen drei Zeilen: die Mantisse als Summe ihrer Wertigkeiten, der Wert als Mantisse mal Zweierpotenz und dieselbe Zahl als Binärzahl mit verschobenem Komma.
- Darunter steht ein Zahlenstrahl. Der grüne Strich in der Mitte ist die gespeicherte Zahl. Die schwarzen Striche links und rechts sind ihre darstellbaren Nachbarn. Ein rotes Dreieck über dem Strahl zeigt die eingegebene Zahl. Eine blaue Linie unten zeigt den Abstand zur nächstgrößeren Zahl.
- Unter dem Zahlenstrahl stehen vier Angaben: „Eingegeben“, „Gespeichert“, „Darstellungsfehler“ und „Abstand zur nächstgrößeren Zahl“.
- 0,67 wird als Mantisse 0,10101011 mit Exponent 0 gespeichert, das ist 0,66796875. Der Darstellungsfehler ist etwa 0,002, das sind etwa 0,3 Prozent.
- 2/3 wird als 0,10101010 mit Exponent 0 gespeichert, das ist 0,6640625.
- 0,5 wird genau gespeichert: Mantisse 0,10000000, Exponent 0, Darstellungsfehler 0.
- 0,1 wird als 0,099609375 gespeichert: Mantisse 0,11001100, Exponent −3. Ein Zehntel ist im Binärsystem ein unendlicher periodischer Bruch.
- 100,3 wird als 100 gespeichert: Exponent 7, der Abstand zweier Zahlen ist dort 0,5.
- Hat der Student ein Bit oder den Exponenten von Hand verändert, gibt es keine eingegebene Zahl mehr. Bei „Eingegeben“ und „Darstellungsfehler“ steht dann ein Strich.
- Eine Zahl, die einen Exponenten außerhalb von −12 bis 12 bräuchte, nimmt das Modell nicht an. Die Seite meldet dann Überlauf oder Unterlauf in roter Schrift.

### Handgriffe

- Eine Zahl in das Feld eintragen und „speichern“ wählen. Erlaubt sind Dezimalzahlen mit Komma oder Punkt und Brüche wie 2/3.
- Einen der fünf Knöpfe 2/3, 0,67, 0,1, 0,5 und 100,3 wählen.
- Ein Bit der Mantisse anklicken, um es umzuschalten. Das erste Bit geht nicht.
- Das Vorzeichen anklicken, um es umzuschalten.
- Den Exponenten mit − und + um 1 verändern, von −12 bis 12.
- Die Auflösungen der drei Fragen aufklappen.

### Abschnitte des Kurstexts

- 4.4.2
- 4.4.3

## Schritt 2: Im Modell rechnen

Reiter: rechnen

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit zwei Fragen zu 0,67 − 2/3 und zu 100 + 0,2.
- Die Demo hat zwei Eingabefelder für zwei Zahlen, dazwischen eine Auswahl der Rechenart −, + oder ·, und den Knopf „rechnen“. Darunter stehen vier Knöpfe mit Beispielen: 0,67 − 2/3, 0,67 + 2/3, 3 · 2/3 und 100 + 0,2.
- Das Modell speichert erst beide Zahlen mit 8 Bit Mantisse, rechnet dann mit den gespeicherten Werten und speichert das Ergebnis wieder mit 8 Bit Mantisse.
- Die Tabelle hat fünf Zeilen: „Erste Zahl, gespeichert“, „Zweite Zahl, gespeichert“, „Ergebnis im Modell“, „Wahres Ergebnis“ und „Abweichung“. Die Spalten zeigen Mantisse und Exponent und den Wert. Das Ergebnis im Modell ist gelb unterlegt. Eine Abweichung ab 1 Prozent ist rötlich unterlegt.
- Unter der Tabelle steht bei manchen Rechnungen ein gelber Kasten mit einem Hinweis.
- 0,67 − 2/3: Das Modell rechnet 0,66796875 − 0,6640625 = 0,00390625. Das wahre Ergebnis ist 0,00333… Die Abweichung beträgt etwa 17 Prozent. Im Kurstext steht dieselbe Rechnung mit 17,2 Prozent.
- 0,67 + 2/3: Die Abweichung bleibt klein.
- 3 · 2/3: Das Modell liefert nicht genau 2.
- 100 + 0,2: Das Modell liefert 100. Der Summand 0,2 ist kleiner als der Abstand 0,5 zweier darstellbarer Zahlen bei 100 und verschwindet.
- Der Teil „Wo die Gefahr liegt“ nennt die Subtraktion zweier ähnlich großer Zahlen und erklärt die Normalisierung: Die Mantisse wird nach links geschoben, bis sie wieder mit 0,1 beginnt, und von rechts mit Nullen aufgefüllt.
- Das Bild aus dem Kurstext: eine Centmünze wiegen, indem man ein Auto mit und ohne Münze wiegt.

### Handgriffe

- Zwei Zahlen eintragen, die Rechenart wählen und „rechnen“ wählen.
- Einen der vier Knöpfe mit Beispielen wählen.
- Die Rechenarten sind nur −, + und ·. Eine Division gibt es in der Demo nicht.
- Die Auflösungen der zwei Fragen aufklappen.

### Abschnitte des Kurstexts

- 4.4.3

## Schritt 3: In Go mit float64

Reiter: float64

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit einer Frage: `a` ist 0.1, `b` ist 0.2. Was liefert `a + b == 0.3`?
- Die Tabelle „Die Gleitkomma-Typen in Go“ hat drei Zeilen: `float32` mit 32 Bit, 8 Bit Exponent, 23 Bit Mantisse und etwa 7 Dezimalziffern. `float64` mit 64 Bit, 11 Bit Exponent, 52 Bit Mantisse und etwa 16 Dezimalziffern. Das Modell der Seite mit 8 Bit Mantisse und etwa 2 Dezimalziffern.
- Im Kurstext wird empfohlen, in der Praxis `float64` zu verwenden.
- Die Demo hat drei Eingabefelder für `a`, `b` und `c` und den Knopf „vergleichen“. Voreingestellt sind 0.1, 0.2 und 0.3.
- Eine Tabelle zeigt die gespeicherten Werte von `a`, `b`, `a + b` und `c` auf 22 Ziffern. Die Zeile `a + b` ist gelb unterlegt.
- Darunter steht ein dunkler Codeblock mit vier Zeilen Go: die Deklaration der drei Variablen als `float64` und drei Ausgaben. Hinter `//` steht jeweils die Ausgabe.
- Mit 0.1, 0.2 und 0.3: `a + b` ergibt 0.30000000000000004. `a + b == c` ergibt `false`. `math.Abs(a+b-c) < 1e-9` ergibt `true`.
- Diese drei Ausgaben sind mit Go 1.26 nachgeprüft. Für andere Zahlen rechnet der Browser, mit derselben Darstellung wie `float64`.
- Ein gelber Kasten sagt: `0.1 + 0.2 == 0.3` mit Literalen liefert in Go `true`, weil Go konstante Ausdrücke beim Übersetzen exakt ausrechnet. Der Fehler entsteht erst mit Variablen vom Typ `float64`.
- Der Teil „Die Regel für den Vergleich“ sagt: Gleitkommazahlen prüft man nicht mit `==`. Man prüft, ob ihr Abstand kleiner ist als ein kleines ε. `math.Abs` liefert den Betrag einer Zahl.

### Handgriffe

- Drei Zahlen für `a`, `b` und `c` eintragen und „vergleichen“ wählen. Der Betrag jeder Zahl darf höchstens 10^15 sein.
- Die Seite prüft immer `a + b` gegen `c`. Eine andere Rechnung lässt sich nicht einstellen.
- Die Auflösung der Frage aufklappen.

### Abschnitte des Kurstexts

- 4.4.2
- 4.4.4

## Schritt 4: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Eine Zahl speichern, 4.4.2, Seiten 86–87: Gleitkomma-Typen; 4.4.3, Seiten 87–88: Rundungsfehler: Das Modell mit 8 Bit Mantisse und die Zahlen 2/3 und 0,67 stammen aus dem Kurstext. Die Demo, die eine eingegebene Zahl im Modell speichert, gibt die Lektion dazu.
  - 2 Im Modell rechnen, 4.4.3, Seiten 87–90: Rundungsfehler: Die Rechnungen 0,67 − 2/3 und 0,67 + 2/3 stammen aus dem Kurstext. Die Rechnung 100 + 0,2 ist ein eigenes Beispiel.
  - 3 In Go mit float64, 4.4.2, Seiten 86–87: Gleitkomma-Typen; 4.4.4, Seiten 90–91: Vergleich: Der Typ float64 und der Vergleich von Gleitkommazahlen stammen aus dem Kurstext. Das Beispiel mit 0.1, 0.2 und 0.3 und der Hinweis zu konstanten Ausdrücken stehen nicht im Kurstext.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.3, Seiten 10–11, FLOAT: Eine Gleitkommazahl ist eine Näherung einer reellen Zahl. Ihre Operationen arbeiten anders als die für ganze Zahlen, auch wenn sie gleich heißen. Lektion: Exkurs: Speichern und Rechnen im Rechner.
  - Abschnitt 3.3.3, Seiten 29–30, float64: Die Variable tempC hat den Typ float64, eine Näherung der reellen Zahlen mit 64 Bit. Eine Lektion zu diesem Abschnitt gibt es nicht.
  - Abschnitt 4.7, Seiten 109–111, Typkonvertierung: Lektion: Typkonvertierung.
- Unter „Wie es weitergeht“ steht ein Link zum Quiz. Welche Fragen das Quiz stellt, steht auf dieser Seite nicht.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 3“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.
- Den Links zum Quiz, zur Seite „for-Anweisung mit Bedingung“ und zur Seite „Bitverschiebung“ folgen. Alle sind eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 2.3
- 3.3.3
- 4.4.2
- 4.4.3
- 4.4.4
- 4.7
