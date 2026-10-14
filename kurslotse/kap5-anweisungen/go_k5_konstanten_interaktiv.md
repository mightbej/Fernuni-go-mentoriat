# Konstanten in Go: ein fester Wert mit Namen

Die Seite erklärt die Konstanten in Go: die Deklaration mit `const`, die vier Zwecke einer Konstante und die nicht-typisierten Konstanten. Sie gehört zu Abschnitt 5.2 des Kurstexts.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Jeder Code auf der Seite ist Go. Ausgaben und Meldungen des Compilers stammen aus Läufen mit Go 1.26 und stehen fest auf der Seite.
- Auf der Seite läuft kein Go. Der Student kann keinen eigenen Code eingeben und kein Programm starten. Was sich einstellen lässt, steht unter „Handgriffe“.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.
- Ein Literal ist ein Wert, der wörtlich im Programm steht, zum Beispiel `12`. Eine benannte Konstante ist ein Name für einen festen Wert, zum Beispiel `monateImJahr`.

## Worum es geht

Reiter: start

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Liste nennt drei Stellen, an denen Konstanten im Kurstext vor Abschnitt 5.2 vorkommen: Abschnitt 2.2 mit dem Begriff und der Zahl Pi, Abschnitt 3.3.3 mit `const nullpunkt float64 = -273.15`, Abschnitt 4.1 mit dem konstanten Ausdruck.
- Ein konstanter Ausdruck besteht nur aus Literalen und Konstanten. Der Compiler rechnet ihn beim Übersetzen aus.
- Eine zweite Liste sagt, was die vier Schritte der Seite zeigen.
- Ein gelber Kasten sagt, dass die Programme der Seite mit Go 1.26 ausgeführt sind.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Dem Link zur Seite „Variablen, Typen und Speicher“ folgen.

### Abschnitte des Kurstexts

- 2.2
- 3.3.3
- 4.1
- 5.2

## Schritt 1: Deklaration

Reiter: deklaration

### Aufbau der Darstellung

- Oben steht, worin sich die Deklaration einer Konstante von der ausführlichen Deklaration einer Variable unterscheidet: Sie beginnt mit `const`, und der Wert muss dastehen.
- Darunter steht die Demo „Vier Fassungen zum Ausprobieren“ mit vier Knöpfen: „mit Typ“, „ohne Typ“, „ohne Wert“, „Wert später ändern“. Der gewählte Knopf ist grün.
- Unter den Knöpfen steht in einem dunklen Kasten das Programmstück der gewählten Fassung. Darunter steht entweder „Ausgabe“ mit der Ausgabe oder „Meldung des Compilers“ mit der Meldung in roter Schrift. Darunter erklärt ein grüner Kasten die Fassung.
- „mit Typ“: `const maxAnzahl int = 500`, die Ausgabe ist 500.
- „ohne Typ“: `const maxAnzahl = 500`, die Ausgabe ist 500.
- „ohne Wert“: `const maxAnzahl int`. Der Compiler meldet `missing init expr for maxAnzahl`. Auf Deutsch: Der Ausdruck für den Anfangswert fehlt.
- „Wert später ändern“: Nach der Deklaration steht `maxAnzahl = 600`. Der Compiler meldet `cannot assign to maxAnzahl (neither addressable nor a map index expression)`. Auf Deutsch: `maxAnzahl` lässt sich nichts zuweisen.
- Darunter steht der Teil „Mehrere Konstanten in einer Klammer“ mit einem Auszug aus dem Paket `math`: `E`, `Pi`, `Phi`, `Ln2`, `Log2E`, `Ln10`, `Log10E`. Der Auszug lässt sich nicht anklicken.
- `Log2E = 1 / Ln2` ist erlaubt, weil rechts ein konstanter Ausdruck steht.
- Am Ende steht eine Tabelle, die Konstante und Variable in vier Eigenschaften vergleicht: Schlüsselwort, Wert in der Deklaration, Zuweisung nach der Deklaration, Typ in der Deklaration.
- Eine Variable ohne Wert bekommt den Null-Wert ihres Typs, bei `int` ist das 0.

### Handgriffe

- Eine der vier Fassungen wählen. Andere Fassungen gibt es nicht.
- Dem Link zur Seite „Variablendeklaration“ folgen.
- Die Programmzeilen lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 5.2
- 5.1

## Schritt 2: Wozu eine Konstante

Reiter: wozu

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und vier Überschriften. Eine Demo gibt es hier nicht.
- Zweck 1, der Name ist verständlicher als die Zahl: Beispiel `math.Pi`. Zwei Programmzeilen berechnen denselben Umfang, einmal mit `math.Pi`, einmal mit `3.141592653589793`.
- Zweck 2, ein häufig gebrauchter Wert wird nicht jedes Mal neu berechnet: Beispiel `math.Log10E`, deklariert als `1 / Ln10`. Der Compiler rechnet die Division einmal beim Übersetzen aus.
- Zweck 3, der Wert lässt sich an einer Stelle ändern: Beispiel `const maxAnzahl = 500`.
- Zweck 4, der Name sagt, was die Zahl bedeutet: Beispiel `const monateImJahr = 12`. Die Zahl 12 kann im selben Programm auch die Stunden auf der Uhr bedeuten.
- Die Reihenfolge der vier Zwecke ist die des Kurstexts.

### Handgriffe

- Mit „Weiter zu Schritt 3“ zur Aufgabe mit der Zahl 12 wechseln. In diesem Schritt lässt sich nichts einstellen.

### Abschnitte des Kurstexts

- 5.2

## Schritt 3: Die Zahl 12 suchen

Reiter: suchen

### Aufbau der Darstellung

- Die Aufgabe: Ein Jahr soll 13 Monate haben. Der Student ändert das Programm so, dass es mit 13 Monaten rechnet. Die Uhr behält ihre 12 Stunden.
- Zu sehen ist ein Go-Programm in einem dunklen Kasten, mit Zeilennummern am linken Rand.
- Das Programm rechnet drei Werte aus: das Jahresgehalt aus dem Monatsgehalt 3000, den Monatsbeitrag aus dem Jahresbeitrag 1560 und die Stunde auf einem Zifferblatt mit 12 Stunden für 15 Uhr.
- Das Programm gibt es in zwei Fassungen. Zu sehen ist immer nur eine. Der Knopf der gewählten Fassung ist grün.
- Fassung mit Literal: 15 Zeilen. Die Zahl 12 steht an fünf Stellen. In den Zeilen 7, 8 und 14 bedeutet sie die Monate im Jahr. In den Zeilen 10 und 11 bedeutet sie die Stunden auf der Uhr.
- Fassung mit Konstante: 16 Zeilen. Zeile 6 deklariert `const monateImJahr = 12`. Der Name `monateImJahr` steht außerdem in den Zeilen 8, 9 und 15. Die Zahl 12 steht in Zeile 6 und in den Zeilen 11 und 12, dort bedeutet sie die Stunden auf der Uhr.
- Jede Zahl 12 im Programm ist ein heller Knopf. Eine Zahl, die auf 13 geändert ist, hat einen roten Rand.
- Eine Suche markiert ihre Treffer gelb. Gesucht wird nach `12` oder nach `monateImJahr`.
- Unter dem Programm stehen zwei Zeilen: „Treffer der Suche“ mit der Anzahl und „Ausgabe des Programms“ mit dem, was das Programm mit den eingestellten Zahlen ausgibt.
- Die Ausgabe ohne Änderung: `36000 130 3` und `Monate im Jahr: 12`. Die Ausgabe der gelösten Aufgabe: `39000 120 3` und `Monate im Jahr: 13`.
- Darunter sagt ein grüner Kasten, was die Suche gefunden hat und wie weit die Aufgabe gelöst ist.
- Die Aufgabe ist gelöst, wenn alle Stellen mit der Bedeutung „Monate“ auf 13 stehen und die zwei Stellen der Uhr auf 12.
- In der Fassung mit Literal sind dafür drei Zahlen zu ändern, in der Fassung mit Konstante eine, die in der Deklaration.
- Die Suche nach `monateImJahr` findet in der Fassung mit Literal keine Stelle, in der Fassung mit Konstante vier: die Deklaration und drei Verwendungen.
- Die zwei Fassungen behalten ihre Änderungen getrennt.
- Unter der Demo stehen zwei Listen: „Was die Aufgabe zeigt“ und vier Faustregeln.

### Handgriffe

- Die Fassung wählen: „Fassung mit Literal 12“ oder „Fassung mit Konstante monateImJahr“.
- „nach 12 suchen“ oder „nach monateImJahr suchen“ wählen. Nach anderen Wörtern lässt sich nicht suchen. „Suche aufheben“ nimmt die Markierung weg.
- Eine Zahl 12 im Programm anklicken: Sie wird zur 13. Ein zweiter Klick macht sie wieder zur 12. Andere Werte als 12 und 13 gibt es nicht.
- „von vorn“: In beiden Fassungen stehen wieder alle Zahlen auf 12, die Suche ist aufgehoben.
- Die übrigen Teile des Programms lassen sich nicht ändern, auch nicht 3000, 1560 und 15.

### Abschnitte des Kurstexts

- 5.2

## Schritt 4: Nicht-typisierte Konstanten

Reiter: untyped

### Aufbau der Darstellung

- Der Schritt gehört zum Expertenwissen 5.6 des Kurstexts.
- Oben steht als Text: In Go heißt jeder Ausdruck mit festem Wert Konstante, auch ein Literal. Ein Integer-Literal hat den Typ untyped integer, ein Floating-Point-Literal den Typ untyped float.
- Ein Wert vom Typ untyped integer nimmt den Typ an, den die Stelle verlangt, an der er steht. Verlangt die Stelle keinen Typ, wird er zum Standardtyp `int`.
- Darunter steht die Demo: ein dunkler Kasten mit fünf nummerierten Programmzeilen aus dem Kurstext und drei Knöpfe „a + 88“, „b + 88“ und „%T mit 88“. Der gewählte Knopf ist grün.
- Die Zeile, zu der der gewählte Knopf gehört, ist heller hinterlegt und hat links einen gelben Strich: Zeile 3 bei „a + 88“, Zeile 4 bei „b + 88“, Zeile 5 bei „%T mit 88“.
- Unter dem Kasten steht „Typ von 88 an dieser Stelle“: `uint8` in Zeile 3, `int64` in Zeile 4, `int` in Zeile 5. Darunter erklärt ein grüner Kasten den Grund.
- Das Verb `%T` gibt den Typ des Arguments aus.
- Darunter steht der Teil „Eine benannte Konstante ohne Typ“. `const ohneTyp = 500` hat den Typ untyped integer, `const mitTyp int = 500` hat den Typ `int`.
- Ein Codeblock zeigt die Folge: `var x float64 = ohneTyp` ist erlaubt. `var y float64 = mitTyp` übersetzt der Compiler nicht, seine Meldung steht in roter Schrift darunter. Dieser Codeblock ist fest und hat keine Knöpfe.
- Am Ende steht der Teil „Der Compiler rechnet genauer als float64“: Der Compiler rechnet mit nicht-typisierten Konstanten auf mindestens 256 Bit genau. `math.Pi` ist mit mehr als 60 Dezimalziffern deklariert, ein `float64` hält etwa 16.

### Handgriffe

- Einen der drei Knöpfe „a + 88“, „b + 88“ oder „%T mit 88“ wählen. Andere Stellen und andere Literale gibt es in der Demo nicht.
- Die Programmzeilen lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 5.2
- 4.7
- 4.1

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Deklaration, 5.2, Seite 124: Die Deklaration mit const und der Auszug aus dem Paket math stammen aus dem Kurstext. Die Meldungen des Compilers gibt die Lektion dazu.
  - 2 Wozu eine Konstante, 5.2, Seiten 124–125: Die Gründe für eine Konstante stammen aus dem Kurstext. Die Demo gibt die Lektion dazu.
  - 3 Die Zahl 12 suchen, 5.2, Seite 125: Die Aufgabe, ein Jahr mit 13 Monaten anzunehmen, stammt aus dem Kurstext. Das Programm mit der Zahl 12 an fünf Stellen ist ein eigenes Beispiel.
  - 4 Nicht-typisierte Konstanten, 5.2, Seiten 125–126: Expertenwissen 5.6: Die fünf Programmzeilen stammen aus dem Expertenwissen 5.6. Die Auswahl der Zeilen mit dem Typ, den Go jeweils annimmt, gibt die Lektion dazu.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.2, Seiten 8–10, Konstante und Literal: Eine Konstante ist ein Name für einen Wert, der sich nie ändert. Ein Literal ist der Wert selbst. Lektion: Variablen, Typen und Speicher, Schritt 2.
  - Abschnitt 3.3.3, Seiten 29–30, die erste Konstante: const nullpunkt float64 = -273.15 in Listing 3.2. Eine Lektion zu diesem Abschnitt gibt es nicht.
  - Abschnitt 4.7, Seiten 109–111, implizite Typanpassung: Lektion: Typkonvertierung.
- Unter „Wie es weitergeht“ stehen zwei Links: zur Seite „Typkonvertierung“ und zum Quiz zu den Konstanten.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 4“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.
- Den Links zur Seite „Typkonvertierung“ und zum Quiz folgen. Beide sind eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 2.2
- 3.3.3
- 4.7
- 5.2
