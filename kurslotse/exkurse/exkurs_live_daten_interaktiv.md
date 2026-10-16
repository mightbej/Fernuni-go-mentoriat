# Ausblick: Verbünde in Programmen mit Live-Daten

Die Seite ist eine Lektion mit Editor am Ende von Kapitel 8. Sie zeigt drei Go-Programme des Mentors, die Daten aus dem Internet holen und in Verbünden ablegen: ein Aktien-Portfolio, einen Wetter-Vergleich und Reiseangebote. Die Lektion unterrichtet nichts Neues. Sie zeigt dem Studenten, was er mit dem Stoff von Kapitel 5 bis 8 schon kann, und lädt ihn ein, die Programme auszuführen und umzubauen.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Die drei Programme stehen nicht im Kurstext. Sie stammen vom Mentor und liegen im Repository des Mentoriats unter `exkurse/go`: `aktien_portfolio.go`, `wetter_vergleich.go` und `reiseangebote.go`.
- Die Programme arbeiten mit Arrays fester Länge und einer Zählvariable, nicht mit Slices. Sie verwenden keine Maps.
- Neu gegenüber dem Kurs sind nur die Pakete `net/http`, `encoding/json`, `encoding/xml` und `regexp`, dazu `time` für Wartezeiten und Datumsangaben. Sie stecken in wenigen Funktionen am Ende jedes Programms. Die Seite erklärt sie nicht.
- Im Editorlotsen steht in jedem Programm-Schritt ein Editor mit dem ganzen Programm. Der Student kann es ändern, mit „Ausführen“ starten und mit „Fassung zurücksetzen“ wiederherstellen. Anderswo, zum Beispiel auf GitHub Pages, zeigt der Schritt das Programm und seine Ausgabe vom 15.10.2026. Im Editorlotsen steht die Lektion im Menü „Programm wählen“ unter „Exkurse“.
- Die Programme brauchen eine Verbindung zum Internet. Die Ausgabe ist bei jedem Lauf anders.
- Eine Frage „Erst vorhersagen“ gibt es auf dieser Seite nicht. Jeder Programm-Schritt endet mit „Ideen zum Ändern“.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität. Die drei Programme gehören nicht zum Stoff des Kurstexts.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Ein Editor steht hier nicht.
- Der Text sagt: Mit Kapitel 8 hat der Student alles beisammen, um Programme mit echten Daten zu schreiben.
- Eine Liste nennt die drei Programme mit je einem Satz.
- Unter „So arbeitest du mit den Programmen“ steht, dass die Programme im Editorlotsen laufen.
- Ein gelber Kasten „Was in den Programmen neu ist“ nennt `net/http`, `encoding/json` und `encoding/xml` und `regexp` mit je einem Satz.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.3

## Schritt 1: Aktien-Portfolio

Reiter: aktien

### Aufbau der Darstellung

- Oben erklärt der Text Aktie, Kurs, Portfolio und Position. Das Programm holt die Kurse von fünf Aktien bei Yahoo Finance: AAPL, MSFT, GOOGL, AMZN und TSLA. Dann legt es ein Portfolio mit drei Positionen an, AAPL, MSFT und TSLA, und berechnet Gewinn und Verlust. Kaufpreise und Kaufdaten sind im Programm fest eingetragen und ausgedacht.
- „Was du darin schon kennst“ nennt:
  - die Verbünde `Portfolio`, `PortfolioPosition` und `Aktie` ineinander, mit dem Ausdruck `p.Positionen[0].Aktie.AktuellerPreis`, Abschnitt 8.3
  - das Array `Positionen [10]PortfolioPosition` mit dem Feld `Anzahl`, Kapitel 6
  - `holeAktie(symbol string) (Aktie, error)`, Abschnitt 7.1.2
  - `fuegeHinzu(p *Portfolio, …)` und `berechneSummen(p *Portfolio)` mit einem Zeiger, Abschnitt 8.2.7
  - `besteUndSchlechteste`, die Suche nach Maximum und Minimum
- „Was neu ist“ sagt: `holeAktie` fragt Yahoo Finance mit `net/http` und liest die Antwort mit `encoding/json` in den Verbund `yahooAntwort`. Yahoo Finance ist keine offizielle Schnittstelle. Scheitert eine Aktie, rechnet das Programm mit den übrigen weiter.
- Darunter steht der Editor mit `aktien_portfolio.go` und die Ausgabe.
- „Ideen zum Ändern“: andere Kürzel in `symbole`, eine vierte Position mit `fuegeHinzu`, die Positionen mit Bubblesort nach dem Gewinn sortieren.

### Handgriffe

- Im Editorlotsen: das Programm ändern, „Ausführen“, „Stopp“, „Fassung zurücksetzen“.
- Anderswo lässt sich das Programm nur lesen.

### Abschnitte des Kurstexts

- 6
- 7.1.2
- 8.2.7
- 8.3

## Schritt 2: Wetter-Vergleich

Reiter: wetter

### Aufbau der Darstellung

- Oben erklärt der Text Open-Meteo, einen freien Wetterdienst. Das Programm holt das Wetter für sechs Städte: Berlin, München, Hamburg, Hagen, Paris und London. Es sucht die wärmste, die kälteste und die windigste Stadt, rechnet den Durchschnitt und zeigt die Woche für Berlin.
- „Was du darin schon kennst“ nennt:
  - `WetterBericht` mit `Stadt` mit `Koordinaten`, der Ausdruck `bericht.Stadt.Koordinaten.Breitengrad`, Abschnitt 8.3
  - das Array als Feld `Vorhersage [7]TagesVorhersage`
  - das Array-Literal `[6]Stadt{…}`, Abschnitt 8.2.1
  - eine Schleife, die Maximum, Minimum und Summe zugleich bestimmt
  - `switch` in `wetterText`, Abschnitt 5.5.2
- „Was neu ist“ sagt: `holeBericht` fragt Open-Meteo mit `net/http` und liest die Antwort mit `encoding/json`. Die Vorhersage kommt als Arrays mit sieben Werten.
- Darunter steht der Editor mit `wetter_vergleich.go` und die Ausgabe.
- „Ideen zum Ändern“: eine eigene Stadt in `staedte`, die Woche für Hagen, die Stadt mit der höchsten Luftfeuchtigkeit suchen.

### Handgriffe

- Im Editorlotsen: das Programm ändern, „Ausführen“, „Stopp“, „Fassung zurücksetzen“.
- Anderswo lässt sich das Programm nur lesen.

### Abschnitte des Kurstexts

- 5.5.2
- 6
- 8.2.1
- 8.3

## Schritt 3: Reiseangebote

Reiter: reisen

### Aufbau der Darstellung

- Oben erklärt der Text, was ein Feed ist. Das Programm liest die Feeds von mydealz, Gruppe Reisen, und Urlaubspiraten, sucht im Titel jedes Angebots Zielort, Preis und Nächte, holt für jeden Zielort das Wetter bei Open-Meteo und sortiert nach dem Preis je Nacht. Es kennt 32 Zielorte im Array `orte`.
- „Was du darin schon kennst“ nennt:
  - `Wetter *WetterInfo` als Zeiger, `nil`, wenn es kein Wetter gibt, und die Prüfung `if d.Wetter != nil`, Abschnitt 8.2.4 und 7.4
  - `deals[i].Wetter = &wetter[k]`: Alle Angebote mit dem gleichen Zielort zeigen auf dasselbe Element des Arrays `wetter`
  - das Array `deals [maxDeals]ReiseDeal` mit `anzahl`, wie im Telefonbuch
  - Bubblesort in `sortiere`, Abschnitt 8.4.2
  - `werteAus(d *ReiseDeal)` mit einem Zeiger, Abschnitt 8.2.7
- „Was neu ist“ sagt: `leseFeed` liest den Feed mit `encoding/xml` Element für Element, `werteAus` sucht mit `regexp`. Das Programm rät: „Flüge: Paris – Westkanada“ gilt als Angebot nach Paris. Am 15.10.2026 nannten 7 von 70 Angeboten Zielort und Preis.
- Darunter steht der Editor mit `reiseangebote.go` und die Ausgabe. Ein Fragezeichen in der Spalte „Nächte“ heißt: Das Angebot nennt keine Zahl der Nächte. Der Preis je Nacht ist dann der ganze Preis.
- „Ideen zum Ändern“: ein eigenes Ziel in `orte`, nur Angebote unter 100 € je Nacht, nur Ziele wärmer als 20 °C.

### Handgriffe

- Im Editorlotsen: das Programm ändern, „Ausführen“, „Stopp“, „Fassung zurücksetzen“.
- Anderswo lässt sich das Programm nur lesen.

### Abschnitte des Kurstexts

- 7.4
- 8.2.4
- 8.2.7
- 8.4.2

## Schritt 4: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Eine Tabelle nennt je Programm die Abschnitte im Kurstext mit Seiten, was aus dem Kurstext stammt und was nicht im Kurstext steht.
- Darunter steht „Wo das Thema im Kurstext noch vorkommt“ mit den Abschnitten 2.2, 8.2 und 8.3, je mit Lektion.
- Unter „Wie es weitergeht“ steht der Link auf den Exkurs zu Tether.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.3
