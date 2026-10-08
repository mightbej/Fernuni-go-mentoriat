"""Erzeugt index.html aus den HTML-Seiten der Kapitelordner.

Aufruf im Hauptverzeichnis des Repositorys:

    python werkzeuge/startseite.py

Der Linktext jeder Seite kommt aus ihrem <title>.
"""
import glob
import html
import os
import re

REPO = "https://github.com/mightbej/Fernuni-go-mentoriat/tree/main/"

KAPITEL = [
    ("kap3-einfuehrung", "Kapitel 3", "Praktische Einführung in Go"),
    ("kap4-primitive-datentypen", "Kapitel 4", "Primitive Datentypen"),
    ("kap5-anweisungen", "Kapitel 5", "Anweisungen"),
    ("kap6-arrays", "Kapitel 6", "Arrays"),
    ("kap7-funktionen-zeiger", "Kapitel 7", "Funktionen und Zeiger"),
    ("kap8-verbund", "Kapitel 8", "Verbund-Datentyp"),
    ("kap9-dynamische-datenstrukturen", "Kapitel 9", "Dynamische Datenstrukturen"),
    ("exkurse", "Exkurse", "Themen neben dem Kurstext"),
]

GRUPPEN = ["Übersicht", "Erklärseiten", "Quiz", "Exkurse"]

# Seiten, die über den Kurstext hinausgehen. Sie stehen im Kapitel, zu dem sie passen.
EXKURSE = {
    "156_go_k5_zahlenraten_programmieruebung.html",
    "exkurs_von_problem_zu_code.html",
    "157_go_k6_BubbleSort.html",
    "159_go_k6_SelectionSort.html",
}


def titel(pfad):
    text = open(pfad, encoding="utf-8").read()
    treffer = re.search(r"<title>(.*?)</title>", text, re.S)
    name = html.unescape(treffer.group(1).strip()) if treffer else os.path.basename(pfad)
    return re.sub(r"^Go Quiz( K\d)?: ", "", name)


def gruppe(datei):
    if datei in EXKURSE:
        return "Exkurse"
    if "quiz_hub" in datei or "quiz_startseite" in datei:
        return "Übersicht"
    if datei.endswith("_quiz.html"):
        return "Quiz"
    return "Erklärseiten"


def kapitel_block(ordner, nummer, name):
    seiten = sorted(glob.glob(ordner + "/*.html"))
    if not seiten:
        return ""
    teile = ['    <section class="card">', "      <h2>%s: %s</h2>" % (nummer, html.escape(name))]
    for g in GRUPPEN:
        eintraege = sorted((titel(s), os.path.basename(s)) for s in seiten if gruppe(os.path.basename(s)) == g)
        if not eintraege:
            continue
        teile.append("      <h3>%s</h3>" % g)
        teile.append("      <ul>")
        for text, datei in eintraege:
            teile.append('        <li><a href="%s/%s">%s</a></li>' % (ordner, datei, html.escape(text)))
        teile.append("      </ul>")
    if glob.glob(ordner + "/go/*.go"):
        teile.append('      <p class="go"><a href="%s%s/go">Go-Programme zu diesem Abschnitt auf GitHub</a></p>' % (REPO, ordner))
    teile.append("    </section>")
    return "\n".join(teile)


VORLAGE = """<!doctype html>
<html lang="de">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Begleitmaterial zum Mentoriat Imperative Programmierung</title>
  <style>
    :root { --bg:#f7f9fc;--card:#fff;--text:#1f2937;--muted:#4b5563;--line:#e5e7eb;--accent:#0f766e;--highlight:#fef3c7; }
    * { box-sizing:border-box; }
    body { margin:0;font-family:"Segoe UI",Tahoma,Geneva,Verdana,sans-serif;color:var(--text);background:radial-gradient(circle at 18% -12%,#d9fbe8 0%,#f7f9fc 46%);line-height:1.6; }
    .container { max-width:1000px;margin:0 auto;padding:24px; }
    .hero { background:linear-gradient(135deg,#0f766e 0%,#1e40af 100%);color:#fff;border-radius:14px;padding:24px;box-shadow:0 10px 24px rgba(15,118,110,.25);margin-bottom:20px; }
    .hero h1 { margin:0 0 10px;font-size:1.9rem; }
    .hero p { margin:0;opacity:.96;font-size:1.05rem; }
    .card { background:var(--card);border:1px solid var(--line);border-radius:12px;padding:20px;margin-bottom:18px; }
    h2 { margin-top:0;font-size:1.3rem;color:var(--accent);border-bottom:2px solid var(--line);padding-bottom:8px; }
    h3 { font-size:1.05rem;margin:18px 0 8px;color:var(--muted); }
    ul { margin:0;padding-left:22px;columns:2;column-gap:32px; }
    @media (max-width:700px) { ul { columns:1; } }
    li { margin:4px 0;break-inside:avoid; }
    a { color:#1e40af; }
    a:focus-visible { outline:2px solid var(--accent);outline-offset:2px; }
    .hinweis { background:var(--highlight);padding:12px 14px;border-radius:8px;border-left:4px solid #f59e0b;margin-bottom:18px; }
    .go { margin:14px 0 0;font-size:.95rem; }
    footer { color:var(--muted);font-size:.9rem;padding:8px 4px 24px; }
  </style>
</head>
<body>
  <div class="container">
    <header class="hero">
      <h1>Begleitmaterial zum Mentoriat</h1>
      <p>Kurs 01613 „Imperative Programmierung“ mit Go. Erklärseiten und Quizze, die wir in den Sitzungen des Mentoriats verwenden.</p>
    </header>
    <div class="hinweis">
      Diese Seiten sind ein privates Angebot des Mentors und begleiten die Sitzungen. Verbindlich für Kurs und Klausur sind allein der Kurstext und die Moodle-Umgebung der FernUniversität in Hagen.
    </div>
@KAPITEL@
    <footer>Die Seiten sind nach den Kapiteln des Kurstexts geordnet. Jede Seite läuft im Browser, eine Installation ist dafür nicht nötig.</footer>
  </div>
</body>
</html>
"""


def main():
    bloecke = [kapitel_block(*k) for k in KAPITEL]
    with open("index.html", "w", encoding="utf-8", newline="\n") as f:
        f.write(VORLAGE.replace("@KAPITEL@", "\n".join(b for b in bloecke if b)))


if __name__ == "__main__":
    main()
