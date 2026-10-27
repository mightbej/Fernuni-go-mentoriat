# Variablen, Typen und Speicher: was eine Variable im Programm ist

Die Seite erklärt, was eine Variable in einem Programm ist: ein Stück Speicher mit einem Namen und einem Typ. Sie stellt die Variable der Mathematik neben die Variable im Programm, klärt die Begriffe Typ, Operator, Konstante und Literal, zeigt zusammengesetzte Typen und zeigt die Variable als Speicherplatz aus acht Speicherelementen. Aus Abschnitt 2.4 des Kurstexts nimmt sie die Konversion in Schritt 2 und die Deklaration in Schritt 4 auf.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Die Seite gehört zu Kapitel 2 des Kurstexts. Dort kommt noch kein Go vor. Die Programmzeilen der Seite gehören zu keiner bestimmten Programmiersprache.
- Der Pfeil `←` bedeutet: Die Variable links bekommt den Wert, der rechts steht. Go schreibt dafür `=`.
- Die Seite führt kein Programm aus. Der Student kann die Programmzeilen nicht ändern und nicht starten.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und einem Bild des Rechners. Eine Demo gibt es hier nicht.
- Eine Liste ordnet drei Abschnitte des Kurstexts je einem Werk des Rechners zu: Abschnitt 2.1 dem Steuerwerk, Abschnitt 2.2 dem Speicherwerk, Abschnitt 2.3 dem Rechenwerk. Beim Steuerwerk steht ein Link auf die Lektion „Imperative Programmierung“. Dort ändern Verzweigung und Schleife die Reihenfolge der Anweisungen.
- Das Bild zeigt vier Kästen. Oben stehen Steuerwerk und Rechenwerk nebeneinander in einem gestrichelten Rahmen mit der Aufschrift „Prozessor“. Darunter stehen Speicherwerk und „Ein- und Ausgabe“ nebeneinander.
- Der Kasten „Speicherwerk“ ist gelb hervorgehoben, weil die Seite zu Abschnitt 2.2 gehört. Die Kästen lassen sich nicht anklicken.
- Die Aufteilung in vier Werke geht auf John von Neumann zurück. Der Kurstext nennt den Namen von Neumann nicht.
- Unter dem Bild steht ein Link auf den Exkurs „Speichern und Rechnen im Rechner“. Dort arbeiten die vier Werke an einer Zuweisung zusammen.
- Darunter steht, was die vier Schritte der Seite zeigen.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Dem Link zum Exkurs oder dem Link zur Lektion „Imperative Programmierung“ folgen.

### Abschnitte des Kurstexts

- 2.1
- 2.2
- 2.3

## Schritt 1: Variable in Mathematik und Informatik

Reiter: variable

### Aufbau der Darstellung

- Zu sehen sind zwei Kästen nebeneinander: links „In der Mathematik“, rechts „Im Programm“. Auf einem schmalen Bildschirm stehen sie untereinander.
- Links oben steht die Gleichung `a · x^2 + b · x + c = 0`. Darunter erklärt eine Liste: a, b und c sind Parameter, x ist die Unbekannte.
- Links stehen drei Knöpfe „Beispiel 1“, „Beispiel 2“ und „Beispiel 3“. Der gewählte Knopf ist grün.
- Unter den Knöpfen steht ein Feld mit drei Zeilen: die Parameter, die Gleichung mit eingesetzten Parametern und die Werte für x als Menge.
- Beispiel 1: a = 1, b = −3, c = 2. Die Menge ist {1, 2}. Beispiel 2: a = 1, b = −2, c = 1. Die Menge ist {1}. Beispiel 3: a = 1, b = 0, c = 1. Die Menge ist leer.
- Rechts oben steht in einem dunklen Kasten ein Programm mit drei Anweisungen: `x ← 3`, `x ← 7`, `x ← x + 1`.
- Hell hinterlegt mit einem gelben Strich am linken Rand ist die Anweisung, die zuletzt ausgeführt wurde. Am Anfang ist keine markiert.
- Unter den Knöpfen rechts steht der Name `x` und daneben ein Kasten mit dem Wert von x. Am Anfang steht dort grau „kein Wert“ in einem gestrichelten Rahmen. Danach steht dort der Wert auf gelbem Grund.
- Die Zeile „Werte von x bisher“ nennt alle Werte, die x im Lauf schon hatte.
- Die Werte nach den drei Anweisungen sind 3, dann 7, dann 8.
- Unter jedem der zwei Kästen steht in einem grünen Kasten ein erklärender Satz.
- Am Anfang hat x keinen Wert, die Variable ist undefiniert. In Go kommt das nicht vor, dort hat jede Variable von Anfang an einen Wert.
- Die dritte Anweisung `x ← x + 1` rechnet rechts mit dem bisherigen Wert 7 und beschreibt x dann mit 8. Als Gleichung der Mathematik gelesen hätte x = x + 1 keine Lösung.

### Handgriffe

- Links eines der drei Beispiele wählen. Eigene Parameter lassen sich nicht eingeben.
- Rechts „nächste Anweisung“ wählen: Das Programm führt eine Anweisung aus. Nach der dritten Anweisung ist der Knopf gesperrt.
- „eine Anweisung zurück“: Das Programm geht eine Anweisung zurück.
- „von vorn“: Das Programm beginnt neu, x hat wieder keinen Wert.
- Die drei Anweisungen lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 2.2
- 2.3

## Schritt 2: Typ, Operator, Konstante, Literal

Reiter: begriffe

### Aufbau der Darstellung

- Zu sehen ist oben ein dunkler Kasten mit zwei Programmzeilen. Sie berechnen die Fläche eines Kreises.
- Zeile 1: `PI ist die Konstante 3,14159265359`. Zeile 2: `flaeche ← PI · r · r`.
- Anklickbar sind die hellen Teile: `PI`, `3,14159265359`, `flaeche`, der Punkt `·` und `r`. Die Wörter „ist die Konstante“ und der Pfeil `←` sind nicht anklickbar.
- Der gewählte Teil ist gelb markiert. Kommt er in den zwei Zeilen mehrmals vor, sind alle Stellen gelb.
- Unter dem Kasten steht in einem grünen Kasten die Auskunft zum gewählten Teil.
- `PI` ist eine Konstante: eine Variable, deren Wert nie verändert wird.
- `3,14159265359` ist ein Literal: ein Wert, der wörtlich im Programm steht. Go schreibt einen Punkt an der Stelle des Kommas.
- `flaeche` und `r` sind Variablen mit dem Typ Gleitkommazahl.
- Der Punkt `·` ist ein Operator, hier die Multiplikation. Go schreibt dafür `*`.
- Darunter steht eine Tabelle mit zwei Typen. Ganze Zahl: Die Addition macht aus 17 und 4 die Zahl 21. Zeichenkette: Die Konkatenation macht aus „Haus“ und „tür“ die Zeichenkette „Haustür“.
- Unter der Tabelle steht der Teil „Zwei Typen in einer Rechnung: die Konversion“. Er besteht nur aus Text, eine Demo hat er nicht.
- Das Beispiel dort: Die Variable `n` hat den Typ ganze Zahl und den Wert 17. Die Variable `y` hat den Typ Gleitkommazahl und den Wert 0,5. Das Programm soll `n + y` ausrechnen.
- Die Addition verknüpft nur zwei Werte desselben Typs. Der Wert von `n` wird deshalb erst in die Gleitkommazahl 17,0 überführt, danach ergibt die Addition 17,5. Die Variable `n` behält ihren Typ und ihren Wert.
- Das Beispiel gehört zu keiner bestimmten Programmiersprache. Ob die Konversion implizit oder explizit geschieht, lässt es offen.
- Konversion heißt: Ein Wert wird in einen anderen Typ überführt.
- Implizit heißt: Der Rechner nimmt die Konversion von selbst vor, sie steht nicht im Programm. Explizit heißt: Der Programmierer schreibt die Konversion ins Programm.
- Eine Sprache, die explizite Konversionen verlangt, heißt streng typisiert. Go ist streng typisiert.
- In Go heißt die Konversion Typkonvertierung und wird als Typname mit runden Klammern geschrieben, zum Beispiel `float64(n)`. Den Typ einer Variable passt Go nie von selbst an. In Go wäre `n + y` deshalb ein Fehler, richtig ist `float64(n) + y`.
- Die Variablen `n` und `y` kommen in den zwei Programmzeilen oben nicht vor. Der Teil zur Konversion ist reiner Lesetext. Ausprobieren lässt sich die Konversion erst auf der Seite „Typkonvertierung“, zu der der Link am Ende des Schritts führt.
- Am Ende steht eine Liste mit den Stellen, an denen Literale, Konstanten, die Konkatenation und die Konversion in Go behandelt werden.

### Handgriffe

- Einen hellen Teil der zwei Programmzeilen anklicken.
- Den Links zu den Seiten „Konstanten“ und „Typkonvertierung“ folgen. „Typkonvertierung“ ist eine eigene HTML-Seite zu Kapitel 4 und kein Reiter dieser Seite. Auf dieser Seite lässt sich kein Go-Code ausführen.
- Zur Konversion gibt es genau einen Handgriff: dem Link zur Seite „Typkonvertierung“ folgen.
- Die Programmzeilen lassen sich nicht ändern. Einen Wert für `r` kann der Student nicht eingeben, die Seite rechnet die Fläche nicht aus.

### Abschnitte des Kurstexts

- 2.2
- 2.4
- 4.3.1
- 4.6.2.2
- 4.7
- 5.2

## Schritt 3: Zusammengesetzte Typen

Reiter: typen

### Aufbau der Darstellung

- Zu sehen sind zwei Darstellungen untereinander: oben ein Vektor, darunter ein Datenverbund.
- Der Vektor ist eine Reihe aus fünf Kästen ohne Lücke. In den Kästen stehen von links nach rechts die Messwerte 12, 15, 11, 18 und 14. Jeder ist eine ganze Zahl.
- Der Datenverbund ist ein Rahmen mit der Aufschrift „Person“. Im Rahmen stehen untereinander die Felder Vorname mit „Erika“, Nachname mit „Mustermann“, dann ein blauer Rahmen mit der Aufschrift „Adresse“, dann Telefonnummer mit „0351 5550123“.
- Im blauen Rahmen „Adresse“ stehen vier Felder: Straße mit „Heidestraße“, Hausnummer mit 17, Postleitzahl mit „01067“ und Ort mit „Dresden“.
- Die Adresse ist selbst ein Datenverbund. Ein Datenverbund kann also einen anderen enthalten.
- Die Hausnummer ist eine ganze Zahl. Alle anderen Felder sind Zeichenketten. Die Postleitzahl ist eine Zeichenkette, weil sie mit einer 0 beginnen kann. Die Telefonnummer ist eine Zeichenkette, weil niemand mit ihr rechnet.
- Das gewählte Element oder Feld ist gelb markiert. Es ist immer nur eines gewählt, im Vektor oder im Datenverbund.
- Unter dem Datenverbund steht in einem grünen Kasten die Auskunft: der Typ des gewählten Elements oder Felds und wozu es gehört.
- Die Person und ihre Angaben sind erfunden.
- Am Ende steht eine Liste: Der Vektor heißt in Go Array, der Datenverbund heißt in Go `struct`.

### Handgriffe

- Eines der fünf Elemente des Vektors anklicken.
- Eines der sieben Felder des Datenverbunds anklicken. Die Rahmen „Person“ und „Adresse“ selbst lassen sich nicht anklicken.
- Den Links zu den Seiten „Arrays“, „struct in Go“ und „Hierarchische Verbünde“ folgen.
- Die Werte lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 2.2
- 6
- 8.2
- 8.3

## Schritt 4: Eine Variable ist Speicherplatz

Reiter: speicher

### Aufbau der Darstellung

- Zu sehen ist eine Reihe aus acht Kästen. Jeder Kasten ist ein binäres Speicherelement. Zusammen sind sie der Speicherplatz einer Variable.
- Über jedem Kasten steht seine Wertigkeit, von links nach rechts: 128, 64, 32, 16, 8, 4, 2, 1.
- Am Anfang steht in jedem Kasten ein graues Fragezeichen in einem gestrichelten Rahmen. Das heißt: Das Speicherelement ist noch nicht beschrieben und hat keinen definierten Wert.
- Ein beschriebenes Speicherelement zeigt 0 oder 1. Ein Kasten mit 1 ist gelb, ein Kasten mit 0 ist weiß.
- Unter der Reihe steht ein Feld mit drei Zeilen: das Bitmuster, die Folge als ganze Zahl gelesen und die Folge als Zeichen gelesen.
- Solange ein Speicherelement nicht beschrieben ist, steht in den Zeilen „als ganze Zahl gelesen“ und „als Zeichen gelesen“ das Wort „undefiniert“.
- Die ganze Zahl ist die Summe der Wertigkeiten aller Kästen mit 1. Sie liegt zwischen 0 und 255. Negative Zahlen zeigt die Seite nicht.
- Das Zeichen ist das Zeichen mit dieser Nummer in der Zeichentabelle ASCII. Bei 65 ist es A, bei 97 ist es a. Bei den Zahlen 0 bis 31 und 127 bis 255 steht dort „kein druckbares Zeichen“, bei 32 „das Leerzeichen“.
- Unter dem Feld steht in einem grünen Kasten ein Satz, der den letzten Handgriff erklärt.
- Die Zuweisung eines Werts beschreibt alle acht Speicherelemente auf einmal: 65 ergibt 01000001, 97 ergibt 01100001, 17 ergibt 00010001.
- Dieselbe Folge aus 0 und 1 bedeutet je nach Typ etwas anderes: 01000001 ist als ganze Zahl 65 und als Zeichen A.
- Unter der Demo steht der Teil „Die Deklaration“. Er besteht nur aus Text.
- Eine Deklaration legt für eine Variable fest: den Namen, den es nur einmal gibt, den Typ, den Speicherplatz und oft einen Anfangswert.
- Aus dem Typ ergibt sich, wie viele Speicherelemente die Variable braucht. Der Rechner stellt sie bereit. Das Fachwort dafür ist „allokieren“.
- Fehlt der Anfangswert, stehen in den Speicherelementen noch die Werte aus ihrer vorherigen Nutzung. Die Variable heißt dann undefiniert. Die Fragezeichen der Demo stehen für diesen Zustand.
- In Go bekommt jede Variable bei der Deklaration einen Wert. Eine undefinierte Variable gibt es in Go nicht.
- Die Demo selbst zeigt keine Deklaration. Name und Typ der Variable mit den acht Speicherelementen nennt die Seite nicht.
- Am Ende des Schritts steht „Nächste Schritte“:
  - Nächste Lektion: „Größter gemeinsamer Teiler: vom Struktogramm zum Go-Programm“.

### Handgriffe

- „Wert 65 zuweisen“, „Wert 97 zuweisen“ oder „Wert 17 zuweisen“ wählen. Knöpfe für andere Werte gibt es nicht.
- Einen der acht Kästen anklicken. Ein Kasten mit Fragezeichen wird mit 1 beschrieben. Ein Kasten mit 1 wechselt auf 0, ein Kasten mit 0 wechselt auf 1.
- Jede ganze Zahl von 0 bis 255 lässt sich so einstellen: erst einen der drei Werte zuweisen, dann einzelne Kästen anklicken. Beispiel: 200 ist 128 + 64 + 8, das Bitmuster ist 11001000.
- „von vorn“: Alle acht Kästen zeigen wieder das Fragezeichen.
- Den Links zur Seite „Variablendeklaration“, zum Exkurs „Speichern und Rechnen im Rechner“, zum Speichermodell und zum Zahlenkreis folgen.
- Mehr als acht Speicherelemente zeigt die Seite nicht.

### Abschnitte des Kurstexts

- 2.2
- 2.4
- 5.1

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - Worum es geht, 2.1, Seiten 5–8; 2.2, Seiten 8–10; 2.3, Seiten 10–11: Die Themen der Abschnitte 2.1 bis 2.3 stammen aus dem Kurstext: die Reihenfolge der Anweisungen, die Variable im Speicher und das Rechnen. Die Zuordnung der drei Abschnitte zu den Werken des Rechners nach John von Neumann gibt die Lektion dazu.
  - 1 Variable in Mathematik und Informatik, 2.2, Seiten 8–9; 2.3, Seiten 10–11: Zuweisung: Die Gleichung a · x² + b · x + c = 0, die Parameter und die Unbekannte stammen aus dem Kurstext. Dort steht auch, dass eine Variable im Programm erst einen Wert haben kann und später einen anderen. Die drei Beispiele zur Gleichung und das Programm mit drei Anweisungen sind eigene Beispiele.
  - 2 Typ, Operator, Konstante, Literal, 2.2, Seite 9; 2.4, Seiten 11–12: Typ, Operator, Konstante und Literal, die Zahl Pi mit dem Wert 3,14159265359 und die Konkatenation stammen aus Abschnitt 2.2. Die Konversion und die streng typisierte Sprache stammen aus Abschnitt 2.4. Die Programmzeilen zur Fläche eines Kreises, die Tabelle der zwei Typen und die Rechnung n + y mit 17 und 0,5 sind eigene Beispiele.
  - 3 Zusammengesetzte Typen, 2.2, Seite 9: Der Vektor und der Datenverbund mit Vorname, Nachname, Adresse und Telefonnummer stammen aus dem Kurstext. Dort ist auch die Adresse ein zusammengesetzter Typ. Die fünf Messwerte und die Werte der Person sind eigene Beispiele.
  - 4 Eine Variable ist Speicherplatz, 2.2, Seiten 9–10; 2.4, Seite 12: Die binären Speicherelemente, das Beschreiben und das Lesen je nach Typ stammen aus Abschnitt 2.2. Die Deklaration mit Name, Typ, Speicherplatz und Anfangswert stammt aus Abschnitt 2.4. Die Variable aus acht Speicherelementen, die Werte 65, 97 und 17 und das Lesen als Zeichen sind eigene Beispiele.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 3.3.3, Seiten 29–30, die erste Variable in Go: var tempC float64 in Listing 3.2, mit Typ, Null-Wert und der Konstante nullpunkt. Eine Lektion zu diesem Abschnitt gibt es nicht.
  - Abschnitt 3.3.5, Seiten 33–35, Zuweisung: Go schreibt die Zuweisung mit =. Lektion: Zuweisungen in Go, Schritt 1.
  - Abschnitt 4.3.1, Seiten 59–61, Literale: Wie Go ganze Zahlen als Literale schreibt. Lektion: Zahlensysteme in Go, Schritt 2.
  - Abschnitt 4.3.4, Seite 63, Ganzzahl-Typen: Wie viele Bit ein Typ von Go belegt. Lektion: Speichermodell.
  - Abschnitt 4.6.2.2, Seiten 107–108, Konkatenation: Go hängt zwei Strings mit + aneinander. Lektion: Strings in Go, Schritt 4.
  - Abschnitt 4.7, Seiten 109–111, Konversion: In Go heißt sie Typkonvertierung. Lektion: Typkonvertierung.
  - Abschnitt 5.1, Seiten 113–123, Deklaration: Lektion: Variablendeklaration in Go.
  - Abschnitt 5.2, Seiten 124–126, Konstanten: Lektion: Konstanten in Go.
  - Kapitel 6, Seiten 151–162, Vektor: In Go heißt er Array. Lektion: Arrays in Go.
  - Kapitel 8, Seiten 195–214, Datenverbund: In Go heißt er Verbund, der Typ ist struct. Der Verbund im Verbund wird in Abschnitt 8.3 behandelt. Lektionen: Verbund in Go und Hierarchische Verbünde.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 4“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.1
- 2.2
- 2.3
- 2.4
- 3.3.3
- 3.3.5
- 4.3.1
- 4.3.4
- 4.6.2.2
- 4.7
- 5.1
- 5.2
- 6
- 8.2
- 8.3
