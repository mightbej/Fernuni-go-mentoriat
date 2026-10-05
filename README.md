# Begleitmaterial zum Mentoriat „Imperative Programmierung“

Erklärseiten, Quizze und Go-Programme, die ich in meinem Mentoriat zum Kurs 01613 der FernUniversität in Hagen verwende.

**Zu den Seiten:** https://mightbej.github.io/Fernuni-go-mentoriat/

## Einordnung

Dieses Repository ist ein privates Angebot des Mentors und begleitet die Sitzungen. Verbindlich für Kurs und Klausur sind allein der Kurstext und die Moodle-Umgebung der FernUniversität. Kurstext, Einsendeaufgaben und Klausuren liegen hier nicht.

## Aufbau

Das Material ist nach den Kapiteln des Kurstexts geordnet.

| Ordner | Kapitel des Kurstexts |
| --- | --- |
| `kap3-einfuehrung` | 3 Praktische Einführung in Go |
| `kap4-primitive-datentypen` | 4 Primitive Datentypen |
| `kap6-arrays` | 6 Arrays |
| `kap7-funktionen-zeiger` | 7 Funktionen und Zeiger |
| `kap8-verbund` | 8 Verbund-Datentyp |
| `kap9-dynamische-datenstrukturen` | 9 Dynamische Datenstrukturen |
| `exkurse` | Themen neben dem Kurstext |

Die HTML-Seiten liegen direkt im Kapitelordner. Go-Programme liegen dort im Unterordner `go`. Die Kapitel 5 bis 7 werden im Lauf des Semesters ergänzt.

## HTML-Seiten

Jede Seite ist eine einzelne Datei mit eigenem CSS und JavaScript und läuft im Browser. Die Seiten veranschaulichen Go mit JavaScript, sie führen keinen Go-Code aus.

## Go-Programme ausführen

Jede Datei enthält ein eigenes `main`-Programm. Deshalb startet man immer eine einzelne Datei, im Hauptverzeichnis des Repositorys:

```bash
go run kap8-verbund/go/S207_struct_demo.go
```

Ohne lokale Go-Installation geht das in GitHub Codespaces: auf GitHub **Code → Codespaces** wählen und den Befehl dort im Terminal eingeben.

Die Programme im Ordner `exkurse/go` holen Daten von fremden Diensten und brauchen eine Internetverbindung.

## Quellen

Zwei Programme bauen auf Listings des Kurstexts auf und nennen die Quelle im Dateikopf: `kap8-verbund/go/S230_binaere_suche_demo.go` und `kap9-dynamische-datenstrukturen/go/S234_bubblesort_demo.go`.
