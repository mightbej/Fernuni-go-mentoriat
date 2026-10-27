# Fernlotse: Begleitmaterial zum Mentoriat „Imperative Programmierung“

Lektionen, Quizze und Go-Programme, die ich in meinem Mentoriat zum Modul 63811 „Einführung in die imperative Programmierung“ der FernUniversität in Hagen verwende.

**Zu den Seiten:** https://mightbej.github.io/Fernuni-go-mentoriat/

## Einordnung

Der Fernlotse ist ein privates Angebot des Mentors und begleitet die Sitzungen. Verbindlich für Kurs und Klausur sind allein der Kurstext und die Moodle-Umgebung der FernUniversität. Kurstext, Einsendeaufgaben und Klausuren liegen hier nicht.

## Aufbau

Das Material ist nach den Kapiteln des Kurstexts geordnet.

| Ordner | Kapitel des Kurstexts |
| --- | --- |
| `kap1-einleitung` | 1 Einleitung |
| `kap2-grundlagen` | 2 Grundlagen imperativer Programmierung |
| `kap3-einfuehrung` | 3 Praktische Einführung in Go |
| `kap4-primitive-datentypen` | 4 Primitive Datentypen |
| `kap5-anweisungen` | 5 Anweisungen |
| `kap6-arrays` | 6 Arrays |
| `kap7-funktionen-zeiger` | 7 Funktionen und Zeiger |
| `kap8-verbund` | 8 Verbund-Datentyp |
| `kap9-dynamische-datenstrukturen` | 9 Dynamische Datenstrukturen |
| `exkurse` | Themen neben dem Kurstext |

Die HTML-Seiten liegen direkt im Kapitelordner. Gehören zu einem Kapitel Go-Programme, liegen sie im Unterordner `go`.

Daneben gibt es drei Ordner:

- **`lotsen`:** Beschreibung und Anleitung zu jedem Lotsen. Die Programme selbst gibt es nur über Moodle.
- **`kurslotse`:** zu jeder Lektion eine Beschreibung für das Sprachmodell. Zum Lernen ist dieser Ordner nicht gedacht.
- **`daten`:** zwei Listen, die die Lotsen aus dem Netz holen. `verweise.json` ordnet den Abschnitten des Kurstexts die Lektionen zu. `programme.json` nennt die Go-Programme für den Editorlotsen.

## HTML-Seiten

Die HTML-Seiten laufen im Browser. Sie veranschaulichen Go mit JavaScript. Es gibt zwei Arten:

- **Lektion:** Eine Lektion führt in Schritten durch ein Thema. Oben stehen Reiter, sichtbar ist immer ein Schritt. Der letzte Reiter heißt „Bezug zum Kurstext“.
- **Quiz:** Ein Quiz stellt Fragen zu einer Lektion.

Vier Dateien im Hauptverzeichnis gehören zu allen Lektionen:

- **`stil.css`:** das Aussehen.
- **`kurslotse.js`:** der Dialog mit dem Kurslotsen.
- **`lotsenleiste.js`:** die Leiste rechts, in der der Dialog steht.
- **`ausprobieren.js`:** der Reiter „Selbst ausprobieren“.

## Die Lotsen

Jede Lektion hat rechts eine Leiste, in der man dem Kurslotsen Fragen zum sichtbaren Schritt stellt. Der Kurslotse ist ein Programm mit einem Sprachmodell. Es läuft auf dem eigenen Rechner und liegt nicht in diesem Repository. Ohne den Kurslotsen arbeiten die Lektionen wie alle anderen HTML-Seiten.

Im Reiter „Selbst ausprobieren“ ändert man ein Go-Programm und führt es aus. Das geht nur, wenn der Editorlotse die Lektion zeigt. Auf GitHub Pages ist der Reiter grau und sagt, wie man die Lektion im Editorlotsen öffnet.

## Go-Programme ausführen

Jede Datei enthält ein eigenes `main`-Programm. Deshalb startet man immer eine einzelne Datei, im Hauptverzeichnis des Repositorys:

```bash
go run kap8-verbund/go/S207_struct_demo.go
```

Ohne lokale Go-Installation geht das im Editorlotsen oder in GitHub Codespaces: auf GitHub **Code → Codespaces** wählen und den Befehl dort im Terminal eingeben.

Drei Programme im Ordner `exkurse/go` holen Daten von fremden Diensten und brauchen eine Internetverbindung: `aktien_portfolio.go`, `wetter_vergleich.go` und `reiseangebote.go`.

## Quellen

Zwei Programme bauen auf Listings des Kurstexts auf und nennen die Quelle im Dateikopf: `kap8-verbund/go/S230_binaere_suche_demo.go` und `kap9-dynamische-datenstrukturen/go/S234_bubblesort_demo.go`.

Einige Lektionen drucken ein Listing des Kurstexts ab. Jede nennt über dem Programm das Listing und den Abschnitt des Kurstexts.
