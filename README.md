# FernUni Go-Mentoriat

Dieses Repository enthält Go-Demonstrationsprogramme und interaktive HTML-Lernmaterialien für das Modul 01603 „Imperative Programmierung“ der FernUniversität in Hagen.

## Inhalt

| Mentoriat                    | Inhalt                                                                  | Status          |
| ---------------------------- | ----------------------------------------------------------------------- | --------------- |
| 01–07                        | Materialien werden später ergänzt                                       | in Vorbereitung |
| [08](mentoriat-08/README.md) | Verbunddatentypen, dynamische Datenstrukturen und praktische Live-Demos | vorhanden       |

## Repository-Struktur

```text
Fernuni-go-mentoriat/
├── go.mod
├── go.sum
├── README.md
└── mentoriat-08/
    ├── s201-232-verbunddatentypen/
    ├── s233-250-dynamische-datenstrukturen/
    ├── zusatzdemos/
    ├── prompts/
    └── README.md
```

Jedes Mentoriat erhält einen eigenen Ordner und ein eigenes README mit Startbefehlen, Themenübersicht und Hinweisen zur HTML-Anzeige.

## Nutzung

Die Bearbeitung und Ausführung erfolgt vorzugsweise in GitHub Codespaces:

1. Auf GitHub **Code → Codespaces** öffnen.
2. Einen vorhandenen Codespace starten oder einen neuen erstellen.
3. Go-Programme im Terminal gezielt mit `go run <Dateipfad>` ausführen.
4. HTML-Seiten über einen Webserver im Codespace anzeigen.
5. Änderungen committen und anschließend mit GitHub synchronisieren.

Die konkreten Befehle und Inhalte stehen im README des jeweiligen Mentoriats.

## Technische Grundlage

* Go-Moduldefinition im Repository-Hauptverzeichnis
* browserbasierte Entwicklungsumgebung über GitHub Codespaces
* eigenständige HTML-Dateien mit integriertem CSS und JavaScript
* GitHub als zentrale, versionierte Ablage

## Hinweis

Die HTML-Seiten veranschaulichen die Go-Konzepte mit JavaScript. Sie führen keinen Go-Code direkt im Browser aus. Die eigentlichen Go-Programme werden im Codespace-Terminal ausgeführt.
