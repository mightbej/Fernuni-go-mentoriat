# Bitverschiebung: Schieben, verdoppeln und halbieren in Go

Die Seite zeigt, was die Operatoren `<<` und `>>` mit den Bits einer Ganzzahl tun und was das für den Wert der Zahl bedeutet. Sie rechnet mit `uint8` und `int8`, wie die Beispiele des Kurstexts.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt. Der Wert von a gilt unter „Worum es geht“, in Schritt 1 und in Schritt 2 gemeinsam: Wer ihn dort ändert, sieht ihn in den beiden anderen wieder.

Den Code auf der Seite kann der Student nicht ändern und nicht ausführen. Die Seite schreibt die Programmzeilen selbst und passt sie an, wenn der Student ein Bit umschaltet.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Reihe mit den acht Bits von a. Über jedem Bit steht seine Wertigkeit: 128, 64, 32, 16, 8, 4, 2, 1.
- Unter der Reihe steht die Summe der Wertigkeiten aller Bits, die auf 1 stehen. Das ist der Wert von a als `uint8`.
- Hier wird noch nichts geschoben.
- Darunter stehen zwei Tabellen mit Text: die Operatoren `<<` und `>>`, und welche Bits je Typ beim Schieben nach rechts nachrücken.

### Handgriffe

- Ein Bit von a anklicken und damit umschalten.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.3.7.2.2

## Schritt 1: Ein Bitmuster schieben

Reiter: schieben

### Aufbau der Darstellung

- Zu sehen ist eine Bit-Tabelle mit drei Zeilen untereinander. Über den Zeilen stehen die Wertigkeiten der acht Bits.
- Obere Zeile, beschriftet mit `a << n`: das Bitmuster von a, um n Stellen nach links geschoben.
- Mittlere Zeile, beschriftet mit `a`: das Bitmuster von a selbst. Diese Zeile bleibt beim Schieben stehen. Sie ist der feste Ausgangswert, mit dem der Student die obere und die untere Zeile vergleicht. Der Ausdruck `a << n` berechnet einen neuen Wert und lässt a unverändert.
- Untere Zeile, beschriftet mit `a >> n`: das Bitmuster von a, um n Stellen nach rechts geschoben.
- Verändert der Student n, dann ändern sich die obere und die untere Zeile. Die mittlere Zeile ändert sich nur, wenn er a ändert.
- Nur die Bits der mittleren Zeile lassen sich anklicken.
- Blau gestrichelt umrandet sind die Bits, die nachgerückt sind: in der oberen Zeile die n Bits ganz rechts, in der unteren Zeile die n Bits ganz links.
- Rot durchgestrichen stehen am Rand die Bits, die weggefallen sind: links neben der oberen Zeile und rechts neben der unteren Zeile.
- Rechts neben jeder Zeile steht ihr Wert als Dezimalzahl.
- Unter der Bit-Tabelle stehen eine Erklärung in Worten und die Programmzeilen in Go. Hinter `//` steht dort die Ausgabe.

### Handgriffe

- Ein Bit von a anklicken und damit umschalten.
- Den Typ von a wählen: `uint8` oder `int8`.
- Für a eine Zahl eingeben: 0 bis 255 bei `uint8`, −128 bis 127 bei `int8`.
- n mit den Knöpfen − und + verändern: 0 bis 8.
- Ein Beispiel des Kurstexts wählen: 13 um 1 Stelle, −13 um 1 Stelle, −13 um 4 Stellen.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.3.7.2.2

## Schritt 2: Verdoppeln und halbieren

Reiter: verdoppeln

### Aufbau der Darstellung

- Zu sehen ist eine Bit-Tabelle mit neun Zeilen untereinander.
- Mittlere Zeile, beschriftet mit `a`: das Bitmuster von a selbst. Sie bleibt stehen, und nur ihre Bits lassen sich anklicken.
- Die vier Zeilen darüber zeigen `a << 1` bis `a << 4`. Direkt über a steht `a << 1`, ganz oben `a << 4`. Jede Zeile nach oben verdoppelt den Wert.
- Die vier Zeilen darunter zeigen `a >> 1` bis `a >> 4`. Direkt unter a steht `a >> 1`, ganz unten `a >> 4`. Jede Zeile nach unten halbiert den Wert und rundet ab.
- Rechts neben jeder Zeile steht die Rechnung, zum Beispiel 26 = 13 · 2. Rot steht ein Ergebnis, das übergelaufen ist.
- Blau gestrichelt umrandet sind nachgerückte Bits. Rot durchgestrichen am Rand stehen weggefallene Bits.

### Handgriffe

- Ein Bit von a anklicken und damit umschalten.
- Ein Beispiel wählen: a = 13 als `uint8`, a = −13 als `int8`.
- Mit den Reitern oben zu einem anderen Schritt wechseln. Den Typ und eine beliebige Zahl für a stellt der Student in Schritt 1 ein.

### Abschnitte des Kurstexts

- 4.3.7.2.2

## Schritt 3: Einzelne Bits setzen mit 1 << i

Reiter: bits-setzen

### Aufbau der Darstellung

- Zu sehen ist eine Reihe mit den acht Bits von `muster`. Darüber stehen zwei Zeilen: die Nummer jedes Bits von 7 bis 0 und seine Wertigkeit von 128 bis 1.
- Unter der Reihe steht der Go-Ausdruck aus `1<<i` und `|`, der genau dieses Bitmuster erzeugt.
- `muster` ist ein eigener Wert. Er hängt nicht mit a aus den Schritten davor zusammen.

### Handgriffe

- Ein Bit von `muster` anklicken und damit umschalten.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.3.7.2.2

## Schritt 4: Das Komma schieben

Reiter: komma

### Aufbau der Darstellung

- Zu sehen sind zwei Bytes nebeneinander: links Byte 1 mit der Mantisse, rechts Byte 2 mit dem Exponenten.
- In der Mantisse steht ein Komma. Es wandert, wenn sich der Exponent ändert. Die Bits der Mantisse bleiben stehen.
- Über jedem Bit der Mantisse steht seine Wertigkeit. Sie ändert sich mit dem Exponenten.
- Das erste Bit der Mantisse ist immer 1 und lässt sich nicht umschalten.
- Der Exponent ist im Zweierkomplement gespeichert, wie ein Wert vom Typ `int8`.
- Unter den zwei Bytes steht der Wert der Gleitkommazahl als Rechnung: Mantisse mal 2 hoch Exponent.
- Das Vorzeichen der Gleitkommazahl lässt die Seite weg.
- Am Ende des Schritts unter „Wie es weitergeht“ stehen Links auf: „Quiz zur Bitverschiebung“.

### Handgriffe

- Ein Bit der Mantisse anklicken, außer dem ersten.
- Ein Bit des Exponenten anklicken.
- Den Exponenten mit den Knöpfen − und + verändern.
- Ein Beispiel des Kurstexts wählen: 2/3 im Modell, 0,67 im Modell.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.4.2
- 4.4.3

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Schieben, 4.3.7.2.2, Seite 75: Schiebe-Operatoren; 4.1, Seite 46, Fußnote 4: Typen der Operanden: Die Operatoren << und >>, die Beispiele mit 13 und −13 und die Regel zu den Typen der Operanden stammen aus dem Kurstext. Das anklickbare Bitmuster mit den nachgerückten und den weggefallenen Bits gibt die Lektion dazu.
  - 2 Verdoppeln und halbieren, 4.3.7.2.2, Seiten 75–77: Tabelle 4.2 und Fußnote 23: Die Ergebnisse für 13 und −13 stehen in Tabelle 4.2. Dass >> und / bei negativen Zahlen verschieden runden, steht in Fußnote 23. Die rot markierten Ergebnisse mit Überlauf gibt die Lektion dazu.
  - 3 Bits setzen, 4.3.7.2.2, Seite 78: Expertenwissen 4.13: Der Ausdruck 1 << i und das Zusammenfügen mit | stammen aus dem Expertenwissen 4.13. Die Demo, die zu einem Bitmuster den Ausdruck zeigt, gibt die Lektion dazu.
  - 4 Komma schieben, 4.4.2, Seite 87: Gleitkomma-Typen; 4.4.3, Seite 88: Normalisieren: Mantisse, Exponent und das Normalisieren stammen aus dem Kurstext. Den Vergleich zwischen dem Schieben von Bits und dem Wandern des Kommas zieht die Lektion.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.3, Seiten 10–11, Rechnen: Der Rechner wandelt Folgen von 0 und 1 in andere Folgen um. Lektion: Exkurs: Speichern und Rechnen im Rechner.
  - Abschnitt 4.3.7.2.1, Seiten 72–74, bit-weise Operatoren: &, | und ^. Lektion: Rechenwerk.
  - Abschnitt 4.3.8, Seiten 78–84, Überlauf: Was mit dem Bit geschieht, das links hinausgeschoben wird. Lektion: Überlauf bei Ganzzahlen in Go.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 4“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.3
- 4.1
- 4.3.7.2.1
- 4.3.7.2.2
- 4.3.8
- 4.4.2
- 4.4.3
