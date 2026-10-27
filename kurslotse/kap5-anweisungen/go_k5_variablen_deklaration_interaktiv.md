# Variablendeklaration in Go

Die Seite erklärt, wie Go eine Variable deklariert: ausführlich mit `var` und kurz mit `:=`. Sie zeigt den Unterschied zwischen Deklaration und Zuweisung, lokale und globale Variablen an Listing 5.1 und die Meldung des Compilers bei einer ungenutzten Variable. Sie gehört zu Abschnitt 5.1 des Kurstexts.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ausgaben und Meldungen des Compilers stammen aus Läufen mit Go 1.26 und stehen fest auf der Seite.
- Auf der Seite läuft kein Go. Der Student kann keinen eigenen Code eingeben und kein Programm starten. Was sich einstellen lässt, steht unter „Handgriffe“.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.
- Eine Deklaration erzeugt eine Variable, gibt ihr einen Typ und einen Anfangswert und verbindet sie mit einem neuen Bezeichner. Der Anfangswert heißt Initialwert.
- `:=` ist der Deklarationsoperator. `=` ist der Zuweisungsoperator. Eine Zuweisung setzt eine Variable voraus, die schon deklariert ist.
- Ein Programm einer Demo steht in einem dunklen Kasten. Hat es Zeilennummern, stehen sie am linken Rand.
- Eine Meldung des Compilers steht in roter Schrift unter dem Programm.
- Die Reihenfolge der Schritte folgt dem Kurstext: 5.1.1, 5.1.2, 5.1.3, 5.1.4, 5.1.5.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und einem Codeblock. Eine Demo gibt es hier nicht.
- Oben steht, was eine Deklaration tut: erzeugen, belegen, benennen.
- Der Codeblock zeigt die zwei Schreibweisen: `var foo int = 42` und `foo := 42`. In einem Programm steht nur eine der zwei Zeilen. Stünden beide untereinander, meldete der Compiler einen Fehler, weil `foo` dann zweimal deklariert wäre.
- Eine Liste nennt drei frühere Stellen des Kurstexts: Abschnitt 2.2 mit Variable, Wert, Typ und Speicherplatz, Abschnitt 3.3.3 mit `var tempC float64` aus Listing 3.2, Kapitel 4 mit den Typen.
- Eine zweite Liste sagt, was die fünf Schritte der Seite zeigen.
- Ein gelber Kasten sagt, dass die Programme der Seite mit Go 1.26 ausgeführt sind.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Dem Link zur Seite „Variablen, Typen und Speicher“ folgen. Sie ist eine eigene HTML-Seite.
- Der Codeblock lässt sich nicht ändern und nicht ausführen.
- In diesem Schritt lässt sich nichts ausprobieren. Die Demos stehen in Schritt 1 bis 5.

### Abschnitte des Kurstexts

- 5.1
- 2.2
- 3.3.3

## Schritt 1: Ausführliche Deklaration

Reiter: var

### Aufbau der Darstellung

- Oben steht die Deklaration `var foo int = 42` mit ihren vier Teilen: Schlüsselwort `var`, Bezeichner `foo`, Typ `int`, Initialwert `= 42`.
- Darunter steht: Der Typ oder der Wert darf fehlen, aber nicht beides.
- Fehlt der Typ, ermittelt Go ihn aus dem Initialwert. Ein Integer-Literal wie `42` hat den Standardtyp `int`, ein Floating-Point-Literal wie `9.95` den Standardtyp `float64`.
- Fehlt der Wert, bekommt die Variable den Null-Wert ihres Typs. Eine Tabelle nennt die Null-Werte: `0` für alle numerischen Typen, `false` für `bool`, `""` für `string`.
- Die Demo heißt „Ausprobieren: eine Deklaration zusammenstellen“. Sie hat zwei Reihen von Knöpfen. Der gewählte Knopf ist grün.
- Obere Reihe, der Typ: `int`, `float64`, `string`, `bool`. Zu jedem Typ gehören ein fester Name und ein fester Wert: `foo` mit 42, `preis` mit 9.95, `name` mit "Peter", `bestanden` mit `true`.
- Untere Reihe, die Form: „mit Typ und Wert“, „ohne Typ“, „ohne Wert“, „ohne Typ und ohne Wert“.
- Unter den Knöpfen steht die Deklaration in einem dunklen Kasten, zum Beispiel `var preis float64` für `float64` und „ohne Wert“.
- Darunter stehen drei Angaben: „Ergebnis“ oder „Meldung des Compilers“, „Typ der Variable“ und „Wert nach der Deklaration“. Darunter erklärt ein grüner Kasten das Ergebnis.
- Die Form „ohne Typ und ohne Wert“ ergibt zum Beispiel `var foo`. Go übersetzt das nicht. Die Meldung: `syntax error: unexpected newline, expected type`. Bei „Typ der Variable“ und „Wert nach der Deklaration“ steht dann „keine Variable“.
- Der Null-Wert von `float64` steht auf der Seite als `0`. So gibt Go ihn aus.
- Unter der Demo stehen drei Teile: „Wozu der Null-Wert gut ist“ mit dem Vergleich zu C, „Mehrere Variablen in einer Deklaration“ mit drei Zeilen und der Schreibweise `var ( … )`, und ein Satz mit Link zu den Konstanten.
- In `var name, note, bestanden = "Peter", 1.7, true` bekommt `name` den Typ `string`, `note` den Typ `float64` und `bestanden` den Typ `bool`.
- In `var a, b, c int` bekommen alle drei Variablen den Wert 0.

### Handgriffe

- Den Typ wählen: nur mit den vier Knöpfen `int`, `float64`, `string` und `bool`. Andere Typen gibt es in der Demo nicht.
- Die Form wählen: nur mit den vier Knöpfen.
- Name und Wert lassen sich nicht ändern. Ein Eingabefeld gibt es nicht.
- Die Codeblöcke unter der Demo sind fest und haben keine Knöpfe.
- Dem Link zur Seite „Konstanten“ folgen. Sie ist eine eigene HTML-Seite.

### Abschnitte des Kurstexts

- 5.1.1

## Schritt 2: Kurze Deklaration

Reiter: kurz

### Aufbau der Darstellung

- Oben stehen `var foo = 42` und `foo := 42` untereinander. Beide bedeuten das Gleiche. Die zweite Zeile ist die kurze Variablendeklaration.
- Für die kurze Deklaration gilt: Der Initialwert ist Pflicht. Ein Typ lässt sich nicht angeben, die Variable bekommt den Standardtyp ihres Initialwerts. Mehrere Variablen gehen auch.
- Der Teil „:= und = nicht verwechseln“ stellt `x := 8` als Deklaration und `x = 10` als Zuweisung gegenüber. In Pascal ist `:=` der Zuweisungsoperator.
- Die Demo heißt „Ausprobieren: was nach x := 8 erlaubt ist“. Sie zeigt ein Programm mit zwei Zeilen und Zeilennummern. Zeile 1 ist immer `x := 8`. Zeile 2 wählt der Student mit einem von sechs Knöpfen. Zeile 2 ist gelb markiert.
- Hinter den zwei Zeilen folgt im ganzen Programm eine Ausgabe aller Variablen. Sie ist auf der Seite nicht abgedruckt. Ohne sie meldete Go die Variablen als ungenutzt.
- Unter dem Programm stehen zwei Angaben: „Ergebnis“ oder „Meldung des Compilers“ und „Variablen nach Zeile 2“. Darunter erklärt ein grüner Kasten Zeile 2.
- `x = 10`: Zuweisung. Go übersetzt. Danach `x = 10`.
- `x := 10`: zweite Deklaration von `x`. Meldung: `no new variables on left side of :=`.
- `var x int = 10`: zweite Deklaration von `x` in ausführlicher Form. Meldung: `x redeclared in this block`.
- `y = 10`: Zuweisung an eine Variable, die nicht deklariert ist. Meldung: `undefined: y`.
- `y := 10`: Deklaration einer zweiten Variable. Go übersetzt. Danach `x = 8, y = 10`.
- `x, y := 10, 20`: Mischform. `x` bekommt 10 zugewiesen, `y` wird deklariert. Go übersetzt. Danach `x = 10, y = 20`.
- Bei einer Meldung läuft das Programm nicht. Bei „Variablen nach Zeile 2“ steht dann, was nach Zeile 1 gälte: `x = 8`.
- Unter der Demo steht der Teil „Die Mischform aus Deklaration und Zuweisung“: Links von `:=` darf ein Teil der Bezeichner schon deklariert sein, mindestens einer muss neu sein. Im Kurstext ist das Expertenwissen 5.3, dort mit den Variablen `s` und `t`. Wozu die Mischform gut ist, wird in Abschnitt 7.1.2 erklärt.

### Handgriffe

- Zeile 2 wählen: nur mit den sechs Knöpfen `x = 10`, `x := 10`, `var x int = 10`, `y = 10`, `y := 10` und `x, y := 10, 20`. Ein Eingabefeld gibt es nicht, andere Zeilen gehen nicht.
- Zeile 1 lässt sich nicht ändern.
- Einen Lauf Anweisung für Anweisung gibt es in diesem Schritt nicht.

### Abschnitte des Kurstexts

- 5.1.2

## Schritt 3: Lokal und global

Reiter: lokal

### Aufbau der Darstellung

- Oben stehen die zwei Begriffe. Lokale Variable: in einer Funktion deklariert, ihr Bezeichner gilt nur in dem Block, in dem die Deklaration steht. Globale Variable: außerhalb aller Funktionen deklariert, sie gehört zum Paket, ihr Bezeichner gilt in allen Funktionen des Pakets. Sie heißt auch package-level Variable.
- Für eine globale Variable ist nur die ausführliche Deklaration mit `var` erlaubt.
- Die Demo zeigt Listing 5.1 des Kurstexts, `lokal_vs_global.go`, mit 22 Zeilen und Zeilennummern.
- Zeile 5: `var klima string`, die globale Variable. Zeilen 7 bis 13: die Funktion `f`. Zeilen 15 bis 22: die Funktion `main`.
- In `f` deklariert Zeile 8 die lokale Variable `wetter` mit "Regen". Zeile 9 gibt `klima` und `wetter` aus. Zeile 10 setzt `klima` auf "warm". Zeile 11 hängt "schauer" an `wetter`. Zeile 12 gibt beide aus.
- In `main` deklariert Zeile 16 die lokale Variable `wetter` mit "Sonne". Zeile 17 setzt `klima` auf "kalt". Zeile 18 gibt beide aus. Zeile 19 ruft `f` auf. Zeile 20 hängt "nschein" an `wetter`. Zeile 21 gibt beide aus.
- Der Student geht den Lauf Anweisung für Anweisung durch. Der Lauf hat 14 Anweisungen und hält der Reihe nach bei den Zeilen 5, 16, 17, 18, 19, 8, 9, 10, 11, 12, 13, 20, 21 und 22.
- Der Lauf beginnt bei Zeile 5, weil Go die globale Deklaration verarbeitet, bevor `main` beginnt. `klima` hat dann den Null-Wert, den leeren String.
- Die gelb markierte Programmzeile ist die Zeile, die gerade ausgeführt wurde. Ein grüner Strich links an einer Programmzeile bedeutet: Die Zeile ist in diesem Lauf schon gelaufen.
- Über dem Programm stehen drei Kästen nebeneinander. Der rote Kasten „Paket main: globale Variable“ zeigt `klima`. Der erste blaue Kasten „Funktion main: lokale Variable“ zeigt die Variable `wetter` von `main`. Der zweite blaue Kasten „Funktion f: lokale Variable“ zeigt die Variable `wetter` von `f`.
- Eine lokale Variable, die es in diesem Augenblick nicht gibt, steht als „gibt es noch nicht“ oder „gibt es nicht mehr“ im Kasten. Die Variable `wetter` von `f` gibt es nur von Zeile 8 bis Zeile 13.
- Ein gelb umrandeter Kasten gehört zu der Funktion, die gerade läuft. Eine gelb unterlegte Variable hat sich mit der letzten Anweisung geändert.
- Unter den Kästen erklärt ein grüner Kasten die markierte Zeile. Unter dem Programm steht ein schwarzer Kasten „Ausgabe bisher“.
- Die ganze Ausgabe des Programms hat vier Zeilen: „Ausgabe main: kalt, Sonne“, „Ausgabe f:    kalt, Regen“, „Ausgabe f:    warm, Regenschauer“, „Ausgabe main: warm, Sonnenschein“.
- Was der Lauf zeigt: Die globale Variable `klima` teilen sich `main` und `f`. Die zwei Variablen `wetter` sind zwei verschiedene Variablen mit dem gleichen Namen und beeinflussen einander nicht.
- Unter der Demo steht der Teil „Wo die Deklaration stehen muss“. Eine lokale Variable muss vor ihrer ersten Verwendung deklariert sein, sonst meldet der Compiler `undefined: a`. Eine globale Variable darf im Quellcode auch hinter ihrer ersten Verwendung deklariert sein. Zu `klima := "kalt"` außerhalb einer Funktion meldet der Compiler `syntax error: non-declaration statement outside function body`.
- Funktionen und ihr Aufruf werden im Kurstext erst in Abschnitt 3.4 kurz und in Kapitel 7 ganz erklärt. Hier genügt: Der Aufruf `f()` führt den Rumpf von `f` aus und kehrt danach zurück.

### Handgriffe

- „Nächste Anweisung“ führt die nächste Anweisung aus. „Bis zum Ende“ führt den Rest aus. „Von vorn“ beginnt wieder bei Zeile 5. Einen Schritt zurück gibt es nicht.
- Die Programmzeilen und die Werte lassen sich nicht ändern. Das Programm hat keine Eingabe.

### Abschnitte des Kurstexts

- 5.1.3

## Schritt 4: Welche Form wann

Reiter: wahl

### Aufbau der Darstellung

- Oben steht: Ob eine Variable ausführlich oder kurz deklariert ist, ändert an der Variable nichts. Es ist eine Frage des Stils.
- Die Empfehlungen dieses Schritts stehen im Kurstext in Abschnitt 5.1.4. Dort sind sie aus einem Lehrbuch zitiert, das im Kurstext als [DK15] geführt ist. Die Seite gibt sie auf Deutsch wieder.
- Die Demo heißt „Vier Deklarationen, eine Variable“. Vier Knöpfe zeigen vier Zeilen: `s := ""`, `var s string`, `var s = ""` und `var s string = ""`. Jede erzeugt eine Variable `s` vom Typ `string` mit dem leeren String als Wert. Der gewählte Knopf ist grün.
- Unter den Knöpfen steht die gewählte Zeile in einem dunklen Kasten. Darunter stehen zwei Angaben: „In einer Funktion“ und „Außerhalb einer Funktion“. Darunter erklärt ein grüner Kasten, was die Form dem Leser sagt.
- `s := ""` ist außerhalb einer Funktion nicht erlaubt. Dort steht in roter Schrift die Meldung `syntax error: non-declaration statement outside function body`. Die drei Formen mit `var` sind an beiden Stellen erlaubt.
- `s := ""`: die kürzeste Form. Sie sagt: Auf den Anfangswert kommt es an.
- `var s string`: verlässt sich auf den Null-Wert. Sie sagt: Der Anfangswert spielt keine Rolle.
- `var s = ""`: selten, üblich nur bei mehreren Variablen in einer Deklaration.
- `var s string = ""`: nennt den Typ ausdrücklich. Nötig, wenn der Typ der Variable ein anderer sein soll als der Typ des Initialwerts.
- Darunter steht der Teil „Die Empfehlungen“ mit vier Listenpunkten. Das Beispiel `var klein int8 = 42` zeigt eine Variable mit einem anderen Typ als dem Standardtyp. Mit `klein := 42` bekäme sie den Typ `int`.
- Der Teil „Wann der Typ dastehen sollte“ zeigt drei Zeilen: `a := 219`, `b := 100`, `c := a / b`. `c` hat den Typ `int` und den Wert 2, weil `a / b` bei ganzen Zahlen der Ganzzahlquotient ist. Im Kurstext ist das Hinweis 5.4.
- Zu `var c float64 = a / b` meldet der Compiler `cannot use a / b (value of type int) as float64 value in variable declaration`.
- `x := math.Exp2(8)` ergibt eine Variable vom Typ `float64` mit dem Wert 256. Den Typ sieht man der Zeile nicht an.

### Handgriffe

- Eine der vier Deklarationen wählen: nur mit den vier Knöpfen. Andere Zeilen gehen nicht.
- Die Codeblöcke unter der Demo sind fest und haben keine Knöpfe. Kein Code dieses Schritts lässt sich ändern oder ausführen.
- Dem Link zur Seite „Typkonvertierung“ folgen. Sie ist eine eigene HTML-Seite.

### Abschnitte des Kurstexts

- 5.1.4
- 5.1.2

## Schritt 5: Ungenutzte Variablen

Reiter: ungenutzt

### Aufbau der Darstellung

- Oben steht: Go übersetzt ein Programm nicht, wenn eine lokale Variable deklariert ist und nie gelesen wird. Eine Zuweisung an die Variable zählt nicht als Benutzung.
- Die Demo heißt „Ausprobieren: vier Programme“. Vier Knöpfe zeigen vier Programme. Der gewählte Knopf ist grün.
- Unter den Knöpfen steht das gewählte Programm mit Zeilennummern. Die Programme sind Ausschnitte ohne `package` und `import`.
- Gelb markiert sind die Zeilen, in denen sich das gewählte Programm vom ersten Programm unterscheidet. Hier bedeutet Gelb also keinen Lauf. Einen Lauf Anweisung für Anweisung gibt es in diesem Schritt nicht.
- Unter dem Programm stehen zwei Angaben: „Ergebnis“ oder „Meldung des Compilers“ und „Ausgabe“. Darunter erklärt ein grüner Kasten das Programm.
- Programm „b wird nie gelesen“, sechs Zeilen: `var a, b int`, `a = 5`, `b = a * 100`, `fmt.Println(a)`. Meldung: `declared and not used: b`. Keine Zeile ist markiert.
- Programm „b wird ausgegeben“, sieben Zeilen: Zeile 6 ist `fmt.Println("ausführliche Ausgabe: ", b)`. Go übersetzt. Die Ausgabe hat zwei Zeilen: 5 und „ausführliche Ausgabe:  500“.
- Programm „Ausgabe auskommentiert, mit _ = b“, acht Zeilen: Zeile 6 ist `_ = b`, Zeile 7 ist die auskommentierte Ausgabe. Go übersetzt. Die Ausgabe ist 5.
- Programm „b ist global“, acht Zeilen: Zeile 1 ist `var b int` außerhalb der Funktion, Zeile 4 ist `var a int`. Go übersetzt. Die Ausgabe ist 5. Dieses Programm steht nicht im Kurstext. Im Kurstext steht nur, dass globale Variablen ungenutzt bleiben dürfen und dass man das nicht ausnutzen soll.
- `a` ist in allen Programmen in Ordnung, weil `b = a * 100` und `fmt.Println(a)` den Wert von `a` lesen.
- Unter der Demo stehen drei Teile: „Warum Go so streng ist“, „Der leere Bezeichner _“ mit `_ = b` und `_, _ = a, b`, und „Globale Variablen“.
- Ein gelber Kasten sagt: Im Kurstext steht die Meldung als `b declared but not used`, Go 1.26 schreibt `declared and not used: b`. Gemeint ist das Gleiche.
- Am Ende des Schritts unter „Nächste Schritte“ steht ein Link zum Selbsttest. Welche Fragen der Selbsttest stellt, steht auf dieser Seite nicht.

### Handgriffe

- Eines der vier Programme wählen: nur mit den vier Knöpfen.
- Die Programmzeilen lassen sich nicht ändern. Ein Eingabefeld gibt es nicht.
- Dem Link zum Selbsttest folgen. Er ist eine eigene HTML-Seite.

### Abschnitte des Kurstexts

- 5.1.5

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Ausführliche Deklaration, 5.1.1, Seiten 114–116: Sprachvergleich 5.1: Die Form mit var, der Null-Wert und der Sprachvergleich stammen aus dem Kurstext. Die Deklarationen mit preis sind eigene Beispiele.
  - 2 Kurze Deklaration, 5.1.2, Seiten 116–118: Hinweis 5.2 und Expertenwissen 5.3: Die kurze Deklaration mit := und die Regel zur Mischform stammen aus dem Kurstext. Die Zeilen der Demo sind eigene Beispiele.
  - 3 Lokal und global, 5.1.3, Seiten 119–120: Listing 5.1 und Hinweis 5.5: Listing 5.1 mit der globalen Variable klima stammt aus dem Kurstext. Die zwei Meldungen des Compilers gibt die Lektion dazu.
  - 4 Welche Form wann, 5.1.4, Seite 121; 5.1.2, Seite 118: Hinweis 5.4: Die Empfehlung, wann welche Form passt, stammt aus dem Kurstext. Das Beispiel mit int8 und die Meldung des Compilers gibt die Lektion dazu.
  - 5 Ungenutzte Variablen, 5.1.5, Seiten 122–123: Die Meldung zur ungenutzten Variable und der leere Bezeichner als Ausweg stammen aus dem Kurstext. Das Programm mit der globalen Variable b ist ein eigenes Beispiel.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.4, Seiten 11–12, Deklaration: Eine Deklaration stellt Speicherplatz bereit, der über den Namen der Variable angesprochen wird. Lektion: Variablen, Typen und Speicher, Schritt 4.
  - Abschnitt 3.3.3, Seiten 29–30, die erste Deklaration: var tempC float64 in Listing 3.2, mit dem Null-Wert 0. Eine Lektion zu diesem Abschnitt gibt es nicht.
  - Abschnitt 3.3.4, Seiten 30–32, Bezeichner: Welche Namen eine Variable haben darf. Lektion: Bezeichner und Schlüsselwörter in Go.
  - Abschnitt 3.4, Seiten 37–41, Gültigkeitsbereich: Lektion: Fehlersuche: drei Go-Programme mit mehreren Fehlern, Schritt 1.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 5“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.4
- 3.3.3
- 3.3.4
- 3.4
- 5.1.1
- 5.1.2
- 5.1.3
- 5.1.4
- 5.1.5
