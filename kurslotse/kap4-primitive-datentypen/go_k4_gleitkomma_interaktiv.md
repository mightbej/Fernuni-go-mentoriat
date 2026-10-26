# Gleitkommazahlen: Mantisse, Exponent und Rundungsfehler in Go

Die Seite erklärt, wie eine Gleitkommazahl aus Vorzeichen, Mantisse und Exponent besteht und warum beim Speichern und beim Rechnen Fehler entstehen. Sie beginnt bei dezimalen Literalen im Quelltext. Danach rechnet sie mit dem Modell des Kurstexts, einer Mantisse aus 8 Bit, zeigt einen Vergleich mit `float64` und schließt mit Grenzen beim Rechnen. Sie gehört zu Abschnitt 4.4 des Kurstexts.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Die Seite schreibt Zahlen im Text und im Modell mit Komma, zum Beispiel 0,67. In Schritt 1, in Schritt 4 und in Go-Code steht der Punkt, zum Beispiel 0.1.
- Das Modell: Die Mantisse hat 8 Bit und beginnt mit 0,1. Sie liegt also zwischen 1/2 und 1. Der Exponent reicht von −12 bis 12. Der Wert ist Mantisse mal 2 hoch Exponent.
- Die acht Bits der Mantisse haben die Wertigkeiten 1/2, 1/4, 1/8, 1/16, 1/32, 1/64, 1/128 und 1/256.
- Das Modell schneidet Stellen ab, die nicht in die Mantisse passen. So rechnen die Beispiele im Kurstext. Go rundet bei `float32` und `float64` zur nächsten darstellbaren Zahl.
- Der Abstand zweier darstellbarer Zahlen ist 2 hoch (Exponent − 8). Bei Exponent 0 ist das 1/256, bei Exponent 7 ist es 0,5.
- Darstellungsfehler: der Unterschied zwischen der eingegebenen und der gespeicherten Zahl. Rundungsfehler: Der Fehler wächst beim Rechnen.
- Die Demos in Schritt 2, 3 und 5 rechnen im Modell und nicht mit Go. Die Demo in Schritt 4 rechnet im Browser mit 64-Bit-Gleitkommazahlen nach IEEE 754, der Darstellung von `float64`.
- Auf der Seite läuft kein Go. Der Student kann keinen eigenen Code eingeben.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.
- Schritt 2 bis 4 beginnen mit einem blauen Kasten „Erst vorhersagen, dann ausprobieren“. Unter jeder Frage steht eine aufklappbare „Auflösung“. Die Antwort ist verdeckt, bis der Student sie aufklappt.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Oben stehen die drei Teile einer Gleitkommazahl: Vorzeichen, Mantisse, Exponent.
- Der Teil „Warum sie Gleitkommazahl heißt“ sagt: Der Exponent verschiebt das Komma der Binärzahl. Exponent 7 macht aus 0,11001000 die Zahl 1100100, also 100.
- Der Teil „Das Modell dieser Lektion“ nennt die 8 Bit der Mantisse und die 52 gespeicherten Nachkommabits bei `float64` und die führende 1, die dazukommt und nicht gespeichert wird. Das Wort „normalisiert“ steht hier noch nicht, es kommt in Schritt 3.
- Der Teil „Zwei Fehler“ erklärt Darstellungsfehler und Rundungsfehler.
- Eine Liste nennt drei frühere Stellen des Kurstexts: Abschnitt 2.2, Abschnitt 3.3.3 und Abschnitt 4.3.5.
- Eine zweite Liste sagt, was die fünf Schritte der Seite zeigen.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den Seiten „Variablen, Typen und Speicher“ und „Zahlenkreis“ folgen. Beide sind eigene HTML-Seiten.
- In diesem Schritt lässt sich nichts ausprobieren. Die Demos stehen in Schritt 1 bis 5.

### Abschnitte des Kurstexts

- 4.4.2
- 4.4.3

## Schritt 1: Gleitkommazahlen im Programm schreiben

Reiter: literale

### Aufbau der Darstellung

- Literale sind Zahlen direkt im Quelltext. Go verwendet den Dezimalpunkt, zum Beispiel `12.5`.
- `e` oder `E` leitet einen Exponenten zur Basis 10 ein. `28.5e6` ist 28,5 mal 10^6, also 28.500.000.
- Im hellen Kasten stehen die Gleichung, die Ziffern mit blau markiertem Dezimalpunkt und die Erklärung seiner Verschiebung. Die Ziffern bleiben an derselben Stelle; der Punkt wandert bei positiven Exponenten nach rechts, bei negativen nach links. Graue Nullen füllen freie Stellen auf. Ist der Kasten schmaler als die Zeile, etwa auf dem Handy, rollt er nach jeder Änderung so, dass der Punkt in der Mitte steht. Die schwarzen Ziffern und die grauen Nullen behalten ihre Position beim Regeln. Exponent 0 lässt den Punkt stehen.
- Der blaue Fragekasten fragt nach `28.5e-2`; die verdeckte Auflösung ist 0,285.
- `0.25`, `.25` und `25e-2` sind gleichwertig. `867.` bedeutet 867,0. Weitere Beispiele erklären ein optionales Pluszeichen, Unterstriche zwischen Ziffern und führende Nullen.
- Der abschließende Hinweis trennt die Basis 10 der dezimalen Literale von der Basis 2 des Speichermodells. Erst die Zuweisung an eine `float64`-Variable überträgt den Wert in diesen Typ und rundet nötigenfalls. Verschiedene Literale mit demselben Wert haben dieselben Darstellungsgrenzen.
- Hexadezimale Gleitkomma-Literale werden wie im Kurstext nicht behandelt.

### Handgriffe

- „Zahl vor e“ wählen: 28.5, 25 oder 0.25. Freie Eingabe gibt es hier nicht.
- Den Regler „Exponent“ von −6 bis 6 in ganzen Schritten verändern.
- „Zeichen“ zwischen e und E wechseln. Das ändert den Zahlenwert nicht.
- Die Auflösung aufklappen. Auf dieser HTML-Seite wird kein Go ausgeführt.

### Abschnitte des Kurstexts

- 4.4.1

## Schritt 2: Eine Zahl speichern

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

## Schritt 3: Im Modell rechnen

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
- Der Teil „Wo die Gefahr liegt“ nennt die Subtraktion zweier ähnlich großer Zahlen und nennt das Aufheben der vorderen Stellen Auslöschung. Er erklärt die Normalisierung: Die Mantisse wird nach links geschoben, bis sie wieder mit 0,1 beginnt, und von rechts mit Nullen aufgefüllt.
- Das Bild aus dem Kurstext: Eine Centmünze soll gewogen werden, dazu wird ein Auto einmal mit und einmal ohne Münze gewogen. Der Unterschied zeigt nur die Ungenauigkeit der Autowaage.

### Handgriffe

- Zwei Zahlen eintragen, die Rechenart wählen und „rechnen“ wählen.
- Einen der vier Knöpfe mit Beispielen wählen.
- Die Rechenarten sind nur −, + und ·. Eine Division gibt es in der Demo nicht.
- Die Auflösungen der zwei Fragen aufklappen.

### Abschnitte des Kurstexts

- 4.4.3

## Schritt 4: In Go mit float64

Reiter: float64

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit einer Frage: `a` ist 0.1, `b` ist 0.2. Was liefert `a + b == 0.3`?
- Die Tabelle „Die Gleitkomma-Typen in Go“ hat drei Zeilen: `float32` mit 32 Bit, 8 Bit Exponent, 23 gespeicherten Nachkommabits und etwa 7 Dezimalziffern. `float64` mit 64 Bit, 11 Bit Exponent, 52 gespeicherten Nachkommabits und etwa 16 Dezimalziffern. Das Modell der Seite mit 8 Bit Mantisse und etwa 2 Dezimalziffern.
- Die führende 1 normalisierter Zahlen wird nicht gespeichert. Zählt man sie mit, hat float32 24 und float64 53 signifikante Binärstellen, das sind die Stellen der Mantisse ab der führenden 1. Im Modell zählen die 8 Bit die führende 1 bereits mit.
- Im Kurstext wird empfohlen, in der Praxis `float64` zu verwenden.
- Die Demo hat drei Eingabefelder für `a`, `b` und `c` und den Knopf „vergleichen“. Voreingestellt sind 0.1, 0.2 und 0.3.
- Eine Tabelle zeigt die gespeicherten Werte von `a`, `b`, `a + b` und `c` auf 22 Ziffern. Die Zeile `a + b` ist gelb unterlegt.
- Darunter steht ein dunkler Codeblock mit vier Zeilen Go: die Deklaration der drei Variablen als `float64` und drei Ausgaben. Hinter `//` steht jeweils die Ausgabe.
- Mit 0.1, 0.2 und 0.3: `a + b` ergibt 0.30000000000000004. `a + b == c` ergibt `false`. `math.Abs(a+b-c) < 1e-9` ergibt `true`.
- Diese drei Ausgaben sind mit Go 1.26 nachgeprüft. Für andere Zahlen rechnet der Browser, mit derselben Darstellung wie `float64`.
- Ein gelber Kasten sagt: `0.1 + 0.2 == 0.3` mit Literalen liefert in Go `true`, weil Go konstante Ausdrücke beim Übersetzen exakt ausrechnet. Der Fehler entsteht erst mit Variablen vom Typ `float64`.
- Der Teil „Die Regel für den Vergleich“ sagt: Bei berechneten Näherungen prüft man, ob ihr Abstand kleiner ist als eine passende Fehlerschranke ε. `==` ist erlaubt und vergleicht gespeicherte Werte exakt. Für gezielt gewählte, genau darstellbare Eingabewerte kann das sinnvoll sein. Die Fehlerschranke muss zur Größenordnung und zur geforderten Genauigkeit passen. `math.Abs` liefert den Betrag einer Zahl.

### Handgriffe

- Drei Zahlen für `a`, `b` und `c` eintragen und „vergleichen“ wählen. Der Betrag jeder Zahl darf höchstens 10^15 sein.
- Die Seite prüft immer `a + b` gegen `c`. Eine andere Rechnung lässt sich nicht einstellen.
- Die Auflösung der Frage aufklappen.

### Abschnitte des Kurstexts

- 4.4.2
- 4.4.4

## Schritt 5: Grenzen beim Rechnen

Reiter: grenzen

### Aufbau der Darstellung

- Der Einstieg erklärt: Mehr Mantissenstellen verkleinern Fehler, beseitigen aber weder periodische Binärbrüche noch die Gefahr bei der Subtraktion ähnlicher Zahlen.
- Die neue Demo verwendet dieselben 8 Mantissenbits und dieselben Exponentengrenzen −12 bis 12 wie die bisherigen Modelldemos. Startwert ist 1 = 0,10000000 mal 2^1.
- Drei Knöpfe stehen über der Tabelle: „Verdoppeln“, „Halbieren“, „Zurück auf 1“. Die Tabelle wächst bei jedem Rechenschritt um eine Zeile. Spalten: Rechnung, Exponent, Wert im Modell. Die Mantisse bleibt 0,10000000.
- Beim elften Verdoppeln wird 2048 erreicht, beim zwölften bräuchte 4096 den Exponenten 13. Die Demo hält mit Überlauf an. Der größte Modellwert ist 4080, den diese Folge nicht trifft.
- Beim dreizehnten Halbieren wird der kleinste positive Modellwert 0,0001220703125 mit Exponent −12 erreicht. Beim vierzehnten bräuchte 0,00006103515625 den Exponenten −13. Die Demo meldet Unterlauf und setzt auf 0. Diese Behandlung des Unterlaufs ist eine Festlegung dieser Demo, keine allgemeine Regel für Go. Das bisherige Speichern und Rechnen weist Ergebnisse außerhalb des Modellbereichs weiterhin zurück.
- Ein gelber Kasten unter der Tabelle nennt den aktuellen Stand und zeigt bei darstellbaren Werten, ob die Fehlerschranke 0,001 schon erreicht ist.
- Nach Überlauf oder Unterlauf sind beide Rechenknöpfe gesperrt; „Zurück auf 1“ bleibt bedienbar. Beim Zurücksetzen wird die Tabelle geleert und der Startwert eingetragen.
- Der Text trennt Go vom Modell: `float64` hat zusätzlich subnormale Zahlen unterhalb der normalisierten Werte; dort sinkt die Genauigkeit. Noch kleinere Ergebnisse runden auf 0. Positiver Überlauf kann +Inf liefern; es gibt keinen umlaufenden Exponenten wie im Zahlenkreis.
- Zwei feste Codeblöcke zeigen mit Go geprüfte Beispiele. `gross * 2` bei `gross = 1e308` ergibt +Inf. `klein / 1e100` bei `klein = 1e-300` ergibt 0. Ein nicht darstellbarer konstanter Wert bei der Zuweisung `var x float64 = 1e400` wird schon beim Übersetzen abgelehnt.
- Die eigene Halbierungsfolge erreicht 0,001 nach 10 Schritten mit 0,0009765625. Die Schranke 0,00001 liegt unter dem kleinsten positiven Modellwert. Ein Nullvergleich würde bis zum Unterlauf weiterrechnen. Die entstehende 0 beweist keine mathematisch exakte Rechnung.
- Im zweiten Go-Beispiel ist h = 1e-20 positiv, aber `1+h == 1` ist true. Die kleine Änderung geht bei der Addition verloren. Das ist Rundung bei 1, kein Unterlauf von h. Der Fragekasten deckt diesen Unterschied auf.
- Ratschläge: eine passende darstellbare Fehlerschranke wählen, unveränderte gespeicherte Ergebnisse erkennen, Wiederholungen begrenzen. Die Lektion führt keine Schleifensyntax ein.
- Abschnitt 4.4.5 enthält eine Steigungsberechnung mit ε. Die vereinfachte Halbierungsfolge und die festen Go-Beispiele sind eigene Ergänzungen; subnormale Zahlen und +Inf präzisieren das Verhalten von Go anhand der Sprachspezifikation.
- Am Ende von Schritt 5 steht „Wie es weitergeht“ mit drei Links: das Quiz zu den Gleitkomma-Typen, die Lektion „for-Anweisung mit Bedingung“, Schritt 5, mit einer Schleife, die wegen eines Rundungsfehlers nie endet, und die Lektion „Bitverschiebung“, Schritt 4. Welche Fragen das Quiz stellt, steht in dieser Lektion nicht.

### Handgriffe

- Wiederholt „Verdoppeln“ oder „Halbieren“ wählen, auch im Wechsel, solange keine Grenze erreicht ist.
- Mit „Zurück auf 1“ einen neuen Versuch beginnen.
- Den Reiter wechseln und zurückkehren: Der Stand bleibt erhalten.
- Die Auflösung aufklappen. Die festen Go-Codeblöcke sind nicht änderbar und nicht ausführbar.

### Abschnitte des Kurstexts

- 4.4.5

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Gleitkommazahlen im Programm schreiben, 4.4.1, Seiten 85–86: dezimale Literale, e/E, verkürzte Schreibweisen, Unterstriche, führende Nullen; die veränderbare Anzeige ist eine eigene Demo.
  - 2 Eine Zahl speichern, 4.4.2, Seiten 86–87: Gleitkomma-Typen; 4.4.3, Seiten 87–88: Rundungsfehler: Das Modell mit 8 Bit Mantisse und die Zahlen 2/3 und 0,67 stammen aus dem Kurstext. Die Demo, die eine eingegebene Zahl im Modell speichert, gibt die Lektion dazu.
  - 3 Im Modell rechnen, 4.4.3, Seiten 87–90: Rundungsfehler: Die Rechnungen 0,67 − 2/3 und 0,67 + 2/3 stammen aus dem Kurstext. Die Rechnung 100 + 0,2 ist ein eigenes Beispiel.
  - 4 In Go mit float64, 4.4.2, Seiten 86–87: Gleitkomma-Typen; 4.4.4, Seiten 90–91: Vergleich: Der Typ float64 und der Vergleich von Gleitkommazahlen stammen aus dem Kurstext. Das Beispiel mit 0.1, 0.2 und 0.3 und der Hinweis zu konstanten Ausdrücken stehen nicht im Kurstext. Dass die führende 1 nicht mitgespeichert wird und float32 und float64 deshalb 24 und 53 signifikante Stellen haben, steht im Kurstext nur als Vorschlag in Fußnote 32.
- Schritt 5 bezieht sich auf 4.4.5, Seiten 91–92: Grenzen der Genauigkeit und des Wertebereichs, Überlauf, Unterlauf und Abbruchgrenzen. Die konkrete Demo und die Go-Beispiele sind eigene Ergänzungen. Die Präzisierungen zu Go verweisen auf die offizielle Sprachspezifikation.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.3, Seiten 10–11, FLOAT: Eine Gleitkommazahl ist eine Näherung einer reellen Zahl. Ihre Operationen arbeiten anders als die für ganze Zahlen, auch wenn sie gleich heißen. Lektion: Exkurs: Speichern und Rechnen im Rechner.
  - Abschnitt 3.3.3, Seiten 29–30, float64: Die Variable tempC hat den Typ float64, eine Näherung der reellen Zahlen mit 64 Bit. Eine Lektion zu diesem Abschnitt gibt es nicht.
  - Abschnitt 4.7, Seiten 109–111, Typkonvertierung: Lektion: Typkonvertierung.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 5“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.
- Den Links zu den Lektionen „Exkurs: Speichern und Rechnen im Rechner“ und „Typkonvertierung“ folgen, dazu den Links auf die Go-Sprachspezifikation und auf den Quelltext von `math.Ldexp`. Die Links zum Quiz, zu „for-Anweisung mit Bedingung“ und zu „Bitverschiebung“ stehen am Ende von Schritt 5.

### Abschnitte des Kurstexts

- 2.3
- 3.3.3
- 4.4.1
- 4.4.2
- 4.4.3
- 4.4.4
- 4.4.5
- 4.7
