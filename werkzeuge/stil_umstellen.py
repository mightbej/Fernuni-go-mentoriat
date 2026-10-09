"""Stellt Seiten auf die gemeinsame stil.css um.

Aufruf im Repository:  python umstellen.py [--schreiben] <Glob> [<Glob> ...]

Jede Seite wird aus HEAD neu aufgebaut. Außerhalb des <style>-Blocks ändert sich nur
die <link>-Zeile und der Farbtausch. Das Skript prüft das selbst und schreibt sonst nicht.
"""
import glob
import os
import re
import subprocess
import sys

sys.stdout.reconfigure(encoding="utf-8")

LINK = '<link rel="stylesheet" href="../stil.css">'
HELL = "radial-gradient(circle at 18% -12%,#d9fbe8 0%,#f7f9fc 46%)"

FARBEN = [("#667eea", "#0f766e"), ("#764ba2", "#1e40af"), ("#8b5cf6", "#0f766e"), ("#6366f1", "#1e40af"),
          ("#7c3aed", "#115e59"), ("#ede9fe", "#f0fdfa"), ("#ddd6fe", "#99f6e4")]
RGBA = [(r"rgba\(\s*102\s*,\s*126\s*,\s*234\s*,", "rgba(15, 118, 110,"),
        (r"rgba\(\s*139\s*,\s*92\s*,\s*246\s*,", "rgba(15, 118, 110,")]

DUNKEL = ("Blockchain_Analogie", "Zusatz_Demos", "S219_verbund_parameter_demo_veranschaulichung",
          "S220_hierarchische_verbunde_demo_veranschaulichung", "S219_verbund_parameter_demo_quiz",
          "S220_hierarchische_verbunde_demo_quiz", "169-201_go_k7_funktionen_zeiger_rekursion_quiz",
          "S207_struct_quiz", "S212_verbund_felder_demo_quiz", "S214_zeiger_auf_verbund_demo_quiz",
          "S216_verbund_vergleich_demo_quiz")

# Nacharbeit je Seite im <style>-Block: Schrift, die auf dem hellen Grund nicht mehr lesbar ist.
ORANGE = [r"button", r"button:hover", r"\.quiz-question h4", r"\.quiz-score", r"\.quiz-option:hover",
          r"\.quiz-option\.selected", r"\.input-group select:focus", r"\.input-group input:focus",
          r"\.conversion-arrow", r"\.example-card", r"\.highlight-box"]


def nacharbeit(name, css):
    if name == "S157-168_go_k6_arrays_quiz.html":
        css = css.replace(".progress-label { color: white;", ".progress-label { color: #4b5563;")
    if name == "156_go_k5_zahlenraten_programmieruebung.html":
        css = re.sub(r"(\.cross-link \{\n)(\s*)(background: #e8eaf6;)", r"\1\2color: #1f2937;\n\2\3", css)
    if name.startswith("go_k4_typkonvertierung_"):
        for sel in ORANGE:
            css = re.sub(r"(?m)(^\s*" + sel + r"\s*\{[^}]*\})",
                         lambda m: m.group(1).replace("#f59e0b", "#0f766e").replace("#d97706", "#115e59"), css)
    return css


FREMD =("go_k4_quiz_hub.html", "go_k4_rechenwerk_interaktiv.html", "go_k4_bitverschiebung")
BEHAELTER = r"(?:\.section|\.card|\.content|\.container|\.content-area|\.step-content|\.level-content|main)"


def farbtausch(t):
    for a, n in FARBEN:
        t = re.sub(re.escape(a), n, t, flags=re.I)
    for a, n in RGBA:
        t = re.sub(a, n, t)
    return t


def familie(eig):
    """Welche Eigenschaften einer Seitenregel deckt eine Eigenschaft der stil.css ab?"""
    if eig == "border":
        return r"border(-(top|right|bottom|left))?(-(color|width|style))?"
    if eig in ("border-left", "border-bottom", "border-top", "border-right"):
        return eig + r"(-(color|width|style))?"
    if eig in ("padding", "margin"):
        return eig + r"(-(top|right|bottom|left))?"
    if eig == "background":
        return r"background(-(color|image))?"
    if eig == "overflow":
        return r"overflow(-[xy])?"
    return re.escape(eig)


def regeln(css):
    """Liefert (anfang, ende, kopf, rumpf_anfang, rumpf_ende, tiefe_at) für jede Stilregel, auch in @media."""
    erg = []

    def lauf(a, e):
        i = a
        start = a
        while i < e:
            if css.startswith("/*", i):
                j = css.find("*/", i)
                i = e if j < 0 else j + 2
                start = i
                continue
            if css[i] == "{":
                tiefe, j = 1, i + 1
                while j < e and tiefe:
                    if css[j] == "{":
                        tiefe += 1
                    elif css[j] == "}":
                        tiefe -= 1
                    j += 1
                kopf = css[start:i]
                if kopf.strip().startswith("@"):
                    if re.match(r"\s*@(media|supports)", kopf):
                        lauf(i + 1, j - 1)
                else:
                    erg.append((start, j, kopf, i + 1, j - 1))
                i = j
                start = i
                continue
            i += 1

    lauf(0, len(css))
    return erg


def lies_stil(pfad):
    css = open(pfad, encoding="utf-8").read()
    karte = {}
    for a, e, kopf, ra, re_ in regeln(css):
        eigs = re.findall(r"(?:^|;)\s*([\w-]+)\s*:", re.sub(r"/\*.*?\*/", "", css[ra:re_], flags=re.S))
        for sel in kopf.split(","):
            karte.setdefault(" ".join(sel.split()), set()).update(eigs)
    return karte


def deckung(sel, karte):
    sel = " ".join(sel.split())
    if sel in karte:
        return set(karte[sel])
    m = re.fullmatch(BEHAELTER + r"\s+(h2|h3)", sel)
    if m:
        return set(karte[m.group(1)]) | {"font-size", "color", "border-bottom", "padding-bottom"}
    return None


def baue_stil(css, karte, weisse_karte):
    aus = css
    geloescht = 0
    for a, e, kopf, ra, re_ in sorted(regeln(css), reverse=True):
        sels = [s for s in kopf.split(",")]
        deck = [deckung(s, karte) for s in sels]
        einzel = " ".join(kopf.split())
        if all(d is not None for d in deck):
            eigs = set.intersection(*deck)
        elif einzel == ".content" and weisse_karte:
            eigs = {"padding"}
        else:
            continue
        muster = "|".join(familie(x) for x in sorted(eigs))
        rumpf = aus[ra:re_]
        neu, n = re.subn(r"[ \t]*(?:" + muster + r")\s*:[^;{}]*;?[ \t]*(?:/\*[^*]*\*/)?[ \t]*\r?\n?", "", rumpf)
        if not n:
            continue
        geloescht += n
        if re.sub(r"/\*.*?\*/", "", neu, flags=re.S).strip() == "":
            # ganze Regel entfernen, samt der Leerzeile davor
            anfang = a
            while anfang > 0 and aus[anfang - 1] in " \t":
                anfang -= 1
            ende = e
            m = re.match(r"[ \t]*\r?\n(?:[ \t]*\r?\n)?", aus[ende:])
            if m:
                ende += m.end()
            aus = aus[:anfang] + aus[ende:] if aus[anfang - 1:anfang] in ("\n", "") else aus[:a] + aus[e:]
        else:
            aus = aus[:ra] + neu + aus[re_:]
    return aus, geloescht


def ohne_stil(t):
    return re.sub(r"(?s)<style[^>]*>.*?</style>", "<style></style>", t)


def art(pfad, alt):
    name = os.path.basename(pfad)
    if any(x in name for x in FREMD):
        return "fremd"
    if any(x in name for x in DUNKEL) or pfad.replace("\\", "/").startswith("kap9"):
        return "dunkel"
    stil = "".join(re.findall(r"(?s)<style[^>]*>.*?</style>", alt))
    if ".content-area" in stil or re.search(r"\.container\s*\{[^}]*height:\s*100vh", stil):
        return "fenster"
    return "voll"


def stelle_um(pfad, karte, schreiben):
    alt = subprocess.run(["git", "show", "HEAD:" + pfad.replace("\\", "/")], capture_output=True).stdout.decode("utf-8")
    alt = alt.replace("\r\n", "\n")
    a = art(pfad, alt)
    if a in ("fremd", "dunkel"):
        return a, 0
    m = re.search(r"(?s)([ \t]*)<style[^>]*>(.*?)</style>", alt)
    if not m or alt.count("<style") != 1:
        return "kein einzelner <style>-Block", 0
    css = m.group(2)
    if a == "fenster":
        # nur Farben: heller Hintergrund, kein Link
        css2, n = re.subn(r"(body\s*\{[^}]*?background(?:-color)?\s*:\s*)linear-gradient\([^;]*\)\s*;", r"\g<1>" + HELL + ";", css, count=1)
        neu = alt[:m.start(2)] + nacharbeit(os.path.basename(pfad), css2) + alt[m.end(2):]
        neu = farbtausch(neu)
        soll = farbtausch(ohne_stil(alt))
    else:
        weiss = bool(re.search(r"\.container\s*\{[^}]*box-shadow", css))
        css2, n = baue_stil(css, karte, weiss)
        css2 = nacharbeit(os.path.basename(pfad), css2)
        neu = alt[:m.start(1)] + m.group(1) + LINK + "\n" + alt[m.start(1):m.start(2)] + css2 + alt[m.end(2):]
        neu = farbtausch(neu)
        soll = farbtausch(ohne_stil(alt)).replace(m.group(1) + "<style></style>", m.group(1) + LINK + "\n" + m.group(1) + "<style></style>", 1)
    if ohne_stil(neu) != soll:
        return "PRÜFUNG FEHLGESCHLAGEN", 0
    if schreiben:
        crlf = b"\r\n" in open(pfad, "rb").read()
        open(pfad, "w", encoding="utf-8", newline="").write(neu.replace("\n", "\r\n") if crlf else neu)
    return a, n


def main():
    schreiben = "--schreiben" in sys.argv
    muster = [x for x in sys.argv[1:] if not x.startswith("--")]
    karte = lies_stil("stil.css")
    for mu in muster:
        for pfad in sorted(glob.glob(mu)):
            a, n = stelle_um(pfad, karte, schreiben)
            print("%-9s %4d  %s" % (a if len(a) < 10 else "FEHLER", n, pfad if len(a) < 10 else pfad + "  -> " + a))


main()
