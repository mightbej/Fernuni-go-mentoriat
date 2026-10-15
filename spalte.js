// Die Spalte der Lotsen: rechts neben der Seite stellt der Student Fragen und liest die Antworten des Sprachmodells.
//
// Diese Datei ist für alle Lotsen dieselbe. Sie liegt im Modul lotsen, Ordner lotsen-gemeinsam, und nur dort wird sie
// geändert. Skriptlotse, Klausurenlotse und Kurslotse liefern sie unter /spalte.js aus. Im Repository der HTML-Seiten
// liegt eine Kopie neben kurslotse.js, die das Skript werkzeuge/spalte_kopieren.py im Ordner mentoriat-dev schreibt.
//
// Eine Seite bindet die Datei ein und richtet die Spalte ein:
//
//   const spalte = Spalte.einrichten({
//     name: 'Klausurenlotse',                      // steht im Kopf, vor jeder Antwort und auf dem Knopf zum Öffnen
//     adresse: '/api/frage',                       // dorthin geht jede Frage
//     hinweis: '<p>Hier stellst du Fragen …</p>',  // steht über einem leeren Gespräch
//     platzhalter: 'Zum Beispiel: …',              // steht im leeren Eingabefeld
//     wo: 'in der Aufgabe',                        // „Markierte Stelle in der Aufgabe: …“
//     sonst: 'Zu dieser Aufgabe',                  // „Zu dieser Aufgabe im Kurstext:“, wenn zur Frage selbst keine Stelle gefunden ist
//     markieren: knoten => !!knoten.closest('main'), // welche Markierungen auf der Seite mit der nächsten Frage mitgehen
//     mitschicken: frage => ({ id: aufgabe.id })   // was das Programm außer Frage und Gespräch braucht
//   });
//
// Dazu nach Bedarf:
//   knopf            Beschriftung des Knopfs, sonst „Fragen“
//   weitere          weitere Knöpfe: [{ art, knopf, text, brauchtMarkierung }]
//   zeilen           () => Programmzeilen der Seite, wenn das Programm nicht in pre steht
//   link             adresse => Adresse, unter der ein Link auf eine HTML-Seite des Mentoriats öffnet
//   kopfzeile        () => Zeilen, die beim Kopieren über dem Gespräch stehen
//   marke            () => kleine Zeile über der Frage des Studenten, zum Beispiel der sichtbare Schritt
//   beimLeeren       () => wird gerufen, wenn das Gespräch geleert ist
//   merken           Name, unter dem sich der Browser merkt, ob die Spalte offen ist
//
// Was die Seite mit der Spalte tun kann:
//   spalte.zeige(an)               Die Spalte gibt es oder gibt es nicht, zum Beispiel erst, wenn eine Aufgabe zu sehen ist.
//   spalte.oeffne(offen)           Die Spalte auf- oder zuklappen.
//   spalte.sperre(html, platzhalter)  Die Spalte nimmt keine Frage an und sagt, warum. Mit sperre(null) nimmt sie wieder Fragen an.
//   spalte.leere()                 Das Gespräch leeren.
//   spalte.markiere(text, wo)      Eine markierte Stelle melden, wenn die Seite sie selbst findet, zum Beispiel im PDF.
//                                  wo ist die Umgebung, wenn die Seite sie kennt: { art: 'code', text, titel, zeigen },
//                                  zum Beispiel im Editorlotsen die ganzen Programmzeilen. Ohne Text verschwindet die Stelle.
//   spalte.frage({ art, text, extra })    Eine Frage stellen, als hätte der Student sie getippt.
//   spalte.einzeln({ art, text, extra })  Eine Frage außerhalb des Gesprächs. Liefert die Antwort als HTML.
//   spalte.neuerStand()            Frühere Fragen und Antworten gehen ab jetzt nicht mehr an das Sprachmodell.
//   spalte.setzeLage(html)         Die Zeile über dem Eingabefeld: worauf sich die nächste Frage bezieht.
//   spalte.setzeHinweis(html), spalte.setzePlatzhalter(text), spalte.setzeZusatz(text)
//
// Für die übrigen Anfragen der Seite an ihren Lotsen gibt es Spalte.hole(pfad, daten), zum Beispiel
// Spalte.hole('/api/bereit'). Mit daten geht die Anfrage als POST mit JSON.
//
// In einem Fenster ab 1100 Pixel Breite zieht der Student die Spalte am linken Rand breiter oder schmaler, zwischen
// 320 Pixeln und 60 Prozent des Fensters. Der Browser merkt sich die Breite unter lotseBreite, je Adresse getrennt:
// Auf GitHub Pages gilt sie für alle HTML-Seiten, jeder Lotse auf dem eigenen Rechner hat seine eigene.
//
// Mit jeder Frage schickt die Spalte: art, markiert, umgebung, markiertCode, getippt und verlauf. Dazu kommt, was
// mitschicken liefert. Die Antwort des Programms: antwort, quellen, zurFrage und nach Bedarf hinweis.
(function () {
  const esc = t => String(t).replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/"/g, '&quot;');
  const glatt = t => t.replace(/\s+/g, ' ').trim();

  const GESTALTUNG = `
    .lotse, .lotse * { box-sizing:border-box; }
    .lotse { position:fixed;top:0;right:0;bottom:0;width:420px;max-width:100%;margin:0;padding:0;display:flex;flex-direction:column;background:#fff;border:0;border-left:1px solid #e5e7eb;border-radius:0;box-shadow:-6px 0 18px rgba(15,23,42,.08);z-index:20;font-family:"Segoe UI",Tahoma,Geneva,Verdana,sans-serif;font-size:16px;color:#1f2937;line-height:1.6;max-height:none; }
    .lotse[hidden],.lotse-auf[hidden],.lotse [hidden] { display:none; }
    .lotse button { font-family:inherit;font-size:.9rem;font-weight:600;line-height:1.4;cursor:pointer;border:1px solid #e5e7eb;border-radius:8px;background:#fff;color:#1f2937;padding:6px 11px;margin:0;width:auto; }
    .lotse button:disabled { opacity:.5;cursor:default; }
    .lotse-kopf { display:flex;align-items:center;gap:8px;padding:12px 16px;background:linear-gradient(135deg,#0f766e 0%,#1e40af 100%);color:#fff; }
    .lotse-kopf strong { font-size:1.1rem;color:#fff; }
    .lotse-kopf span { flex:1;font-size:.82rem;opacity:.9;white-space:nowrap;overflow:hidden;text-overflow:ellipsis; }
    .lotse .lotse-kopf button { background:transparent;color:#fff;border-color:rgba(255,255,255,.5);font-size:.85rem; }
    .lotse .lotse-kopf button.zu { padding:4px 11px;font-size:1.1rem;line-height:1.2; }
    .lotse .lotse-kopf button:hover { background:rgba(255,255,255,.18); }
    .lotse-verlauf { flex:1;overflow-y:auto;display:flex;flex-direction:column;gap:10px;padding:14px 16px;min-height:0; }
    .lotse-hinweis, .lotse-gesperrt { color:#4b5563;font-size:.92rem;margin:0; }
    .lotse-hinweis p, .lotse-gesperrt p { margin:0 0 8px; }
    .lotse-hinweis ul { margin:0;padding-left:20px; } .lotse-hinweis li { margin:4px 0; }
    .lotse.gesperrt .beitrag, .lotse.gesperrt .lotse-hinweis { display:none; }
    .lotse .beitrag { padding:10px 14px;border-radius:10px;max-width:94%;overflow-wrap:anywhere;font-size:.95rem;margin:0; }
    .lotse .beitrag.student { align-self:flex-end;background:#dbeafe;border:1px solid #bfdbfe;white-space:pre-wrap; }
    .lotse .beitrag.modell { align-self:flex-start;background:#f0fdfa;border:1px solid #99f6e4; }
    .lotse .beitrag.fehler { align-self:flex-start;background:#fef2f2;border:1px solid #fecaca;color:#991b1b; }
    .lotse .beitrag.wartet { align-self:flex-start;color:#4b5563;font-style:italic; }
    .lotse .beitrag .stand, .lotse .beitrag .wer { display:block;font-size:.8rem;color:#4b5563;margin-bottom:4px;font-family:inherit;font-weight:400;font-style:normal; }
    .lotse .beitrag pre { margin:8px 0;background:#1e293b;color:#e2e8f0;border-radius:8px;padding:10px;font-family:Consolas,monospace;font-size:.85rem;line-height:1.5;white-space:pre-wrap;overflow-x:auto; }
    .lotse .beitrag code { background:#f3f4f6;color:#1f2937;padding:2px 6px;border-radius:6px;font-family:Consolas,"Courier New",monospace;font-size:.9rem; }
    .lotse .beitrag ul,.lotse .beitrag ol { margin:6px 0;padding-left:20px; }
    .lotse .beitrag li { margin:3px 0; }
    .lotse .beitrag p { margin:0 0 8px; } .lotse .beitrag p:last-child { margin:0; }
    .lotse .beitrag .quellen { margin-top:10px;padding-top:6px;border-top:1px solid #99f6e4;font-size:.82rem;color:#4b5563; }
    .lotse .beitrag .quellen ul ul { margin:2px 0 6px; }
    .lotse .beitrag .quellen a { color:#0f766e;font-weight:600; }
    .lotse .beitrag .zitat { border-left:3px solid #1e40af;padding-left:8px;margin-bottom:6px;color:#4b5563;font-style:italic;max-height:110px;overflow:auto; }
    .lotse .umgebung { display:block;margin-top:4px;font-size:.8rem;font-style:normal;color:#4b5563;white-space:pre-wrap; }
    .lotse .umgebung.code { font-family:Consolas,"Courier New",monospace; }
    .lotse-fuss { border-top:1px solid #e5e7eb;padding:10px 16px 14px;background:#f8fafc; }
    .lotse-lage { font-size:.82rem;color:#4b5563;margin-bottom:8px; }
    .lotse-lage:empty, .lotse-meldung:empty { display:none; }
    .lotse-lage b { color:#1f2937;font-weight:600; }
    .lotse-markiert { position:relative;font-size:.85rem;color:#4b5563;border-left:3px solid #0f766e;background:#fff;padding:6px 28px 6px 8px;margin-bottom:8px;max-height:90px;overflow:auto;white-space:pre-wrap; }
    .lotse .lotse-markiert button { position:absolute;top:0;right:0;padding:2px 7px;border:0;background:transparent;color:#4b5563;font-weight:400; }
    .lotse .lotse-eingabe textarea { display:block;width:100%;min-height:72px;margin:0;border:1px solid #e5e7eb;border-radius:8px;padding:10px;background:#fff;color:#1f2937;font-family:inherit;font-size:.95rem;line-height:1.5;resize:vertical; }
    .lotse .lotse-eingabe textarea:focus { outline:2px solid #0f766e;outline-offset:2px; }
    .lotse .lotse-eingabe textarea:disabled { opacity:.55; }
    .lotse-knoepfe { display:flex;gap:8px;flex-wrap:wrap;margin-top:8px; }
    .lotse .lotse-knoepfe button { border-color:#0f766e;color:#0f766e;padding:8px 14px;font-size:.95rem; }
    .lotse .lotse-knoepfe button.haupt { background:#0f766e;color:#fff; }
    .lotse-meldung { font-size:.8rem;color:#92400e;margin-top:6px; }
    .lotse-meldung a, .lotse-gesperrt a { color:#0f766e;font-weight:600; }
    .lotse-auf { position:fixed;right:18px;bottom:18px;z-index:19;width:auto;margin:0;background:#0f766e;color:#fff;border:1px solid #0f766e;border-radius:8px;font-weight:600;font-size:.95rem;line-height:1.4;font-family:"Segoe UI",Tahoma,Geneva,Verdana,sans-serif;cursor:pointer;box-shadow:0 6px 16px rgba(15,118,110,.35);padding:11px 16px; }
    .lotse-auf:hover { background:#115e59; }
    /* Der Griff am linken Rand: Ziehen ändert die Breite der Spalte, ein Doppelklick setzt sie zurück. */
    .lotse-griff { display:none;position:absolute;top:0;bottom:0;left:-5px;width:10px;cursor:col-resize;z-index:1;touch-action:none; }
    .lotse-griff::after { content:"";position:absolute;top:50%;left:3px;width:4px;height:44px;margin-top:-22px;border-radius:2px;background:#cbd5e1; }
    .lotse-griff:hover::after, .lotse-griff:focus::after, body.lotse-ziehen .lotse-griff::after { background:#0f766e; }
    .lotse-griff:focus { outline:none; }
    body.lotse-ziehen { cursor:col-resize;user-select:none; }
    body.lotse-ziehen iframe { pointer-events:none; }
    /* In einem breiten Fenster rückt die Seite nach links und macht der Spalte Platz. Der Student kann die Spalte dort
       breiter oder schmaler ziehen. In einem schmalen Fenster liegt die Spalte über der Seite und hat ihre feste Breite. */
    @media (min-width:1100px) {
      .lotse { width:var(--lotse-breite,420px); }
      .lotse-griff { display:block; }
      body.lotse-offen { padding-right:var(--lotse-breite,420px); }
    }`;

  // ---------- Die Breite der Spalte ----------

  // Die Breite gilt für alle Seiten unter derselben Adresse. Der Browser merkt sie sich für den nächsten Besuch.
  const BREITE = 420, SCHMALSTE = 320, MERKEN_BREITE = 'lotseBreite';
  let breite = BREITE;
  // Die Breite des Fensters ohne die Bildlaufleiste: Bis zu ihr reicht die Spalte.
  const fenster = () => document.documentElement.clientWidth || window.innerWidth;
  // Die Spalte nimmt höchstens 60 Prozent des Fensters ein, damit von der Seite genug zu sehen bleibt.
  const begrenze = b => Math.round(Math.max(SCHMALSTE, Math.min(b, Math.max(SCHMALSTE, fenster() * 0.6))));
  function setzeBreite(b, merke) {
    breite = b;
    document.documentElement.style.setProperty('--lotse-breite', begrenze(b) + 'px');
    if (merke) { try { localStorage.setItem(MERKEN_BREITE, String(Math.round(b))); } catch (fehler) {} }
  }
  function ladeBreite() {
    let gemerkt = null;
    try { gemerkt = Number(localStorage.getItem(MERKEN_BREITE)); } catch (fehler) {}
    setzeBreite(gemerkt > 0 ? gemerkt : BREITE, false);
    // Wird das Fenster schmaler, wird die Spalte notfalls mit schmaler. Wird es wieder breiter, kommt die gewählte Breite zurück.
    window.addEventListener('resize', () => setzeBreite(breite, false));
  }
  function richteGriffEin(griff) {
    let zieht = false;
    griff.addEventListener('pointerdown', ev => {
      if (ev.button !== 0) return;
      ev.preventDefault();
      zieht = true;
      griff.setPointerCapture(ev.pointerId);
      document.body.classList.add('lotse-ziehen');
    });
    griff.addEventListener('pointermove', ev => {
      if (zieht) setzeBreite(begrenze(fenster() - ev.clientX), false);
    });
    const ende = () => {
      if (!zieht) return;
      zieht = false;
      document.body.classList.remove('lotse-ziehen');
      setzeBreite(breite, true);
    };
    griff.addEventListener('pointerup', ende);
    griff.addEventListener('pointercancel', ende);
    griff.addEventListener('dblclick', () => setzeBreite(BREITE, true));
    // Mit der Tastatur: Pfeil nach links macht die Spalte breiter, Pfeil nach rechts schmaler, Pos1 setzt sie zurück.
    griff.addEventListener('keydown', ev => {
      const schritt = ev.key === 'ArrowLeft' ? 20 : ev.key === 'ArrowRight' ? -20 : 0;
      if (schritt) { ev.preventDefault(); setzeBreite(begrenze(begrenze(breite) + schritt), true); }
      if (ev.key === 'Home') { ev.preventDefault(); setzeBreite(BREITE, true); }
    });
  }

  // ---------- Darstellung einer Antwort ----------

  // Text ohne Codeblock: Zeilen mit „- “ oder „1. “ werden zu einer Liste, eine Tabelle bleibt, wie sie ist, der Rest wird zu Absätzen.
  function absaetze(text) {
    // `Code` kommt zuerst beiseite: Ein Stern darin ist in Go ein Zeiger, zum Beispiel `*auto`, und keine Auszeichnung.
    const fein = z => {
      const code = [];
      return esc(z).replace(/`([^`\n]+)`/g, (ganz, t) => '\u0000' + (code.push(t) - 1) + '\u0000')
        .replace(/\*\*([^*\n]+)\*\*/g, '<strong>$1</strong>').replace(/\*([^*\n]+)\*/g, '<em>$1</em>')
        .replace(/\u0000(\d+)\u0000/g, (ganz, i) => '<code>' + code[Number(i)] + '</code>');
    };
    let html = '', liste = '', tabelle = [];
    const schliesse = () => {
      if (liste) { html += '</' + liste + '>'; liste = ''; }
      if (tabelle.length) { html += '<pre>' + esc(tabelle.join('\n')) + '</pre>'; tabelle = []; }
    };
    text.split('\n').forEach(zeile => {
      const z = zeile.trim(), punkt = z.match(/^[-*•]\s+(.*)$/), nummer = z.match(/^\d+[.)]\s+(.*)$/);
      if (!z) { schliesse(); return; }
      if (/^\|.*\|$/.test(z)) { if (liste) schliesse(); tabelle.push(z); return; }
      if (tabelle.length) schliesse();
      const art = punkt ? 'ul' : nummer ? 'ol' : '';
      if (art !== liste) { schliesse(); if (art) { html += '<' + art + '>'; liste = art; } }
      html += art ? '<li>' + fein((punkt || nummer)[1]) + '</li>' : '<p>' + fein(z.replace(/^#+\s*/, '')) + '</p>';
    });
    schliesse();
    return html;
  }
  // Antwort des Sprachmodells: Codeblöcke, Listen, Tabellen, `Code` und **fett** darstellen, alles andere als Text
  function alsHTML(t) {
    return String(t).split('```').map((teil, i) => i % 2
      ? '<pre>' + esc(teil.replace(/^[a-z]*\n/, '').replace(/\n$/, '')) + '</pre>'
      : absaetze(teil)).join('');
  }

  // ---------- Die Umgebung einer markierten Stelle ----------

  // Eine Markierung wie &b sagt allein nichts, die Seite weiß aber, wo sie steht: in einem Programm die ganzen Zeilen,
  // in einer Tabelle die Tabellenzeile mit den Spaltenköpfen, im Text der Satz. art sagt dem Programm, ob die Stelle
  // Code ist: Mit Code sucht es nicht im ganzen Kurstext.
  function umgebungVon(auswahl, zeilen) {
    const bereich = auswahl.getRangeAt(0);
    const element = knoten => knoten.nodeType === 1 ? knoten : knoten.parentElement;
    const start = element(bereich.startContainer);
    const laenge = bereich.toString().length;
    // so viele Zeichen des Blocks stehen vor der Markierung
    const davor = block => { const r = document.createRange(); r.selectNodeContents(block); r.setEnd(bereich.startContainer, bereich.startOffset); return r.toString().length; };
    const kasten = start.closest('.card, .karte, section, article');
    const kopf = kasten && kasten.querySelector('h2, h3, h4');
    const fertig = (art, text) => {
      text = text.slice(0, 600);
      return { art: art, text: text, titel: kopf ? glatt(kopf.textContent) : '', zeigen: glatt(text) !== glatt(auswahl.toString()) };
    };

    const pre = start.closest('pre');
    if (pre) {
      // Ein Programm, dessen Zeilen einzeln in der Seite stehen, zum Beispiel mit Zeilennummern
      const erste = start.closest('.z');
      if (erste && pre.contains(erste)) {
        const letzte = element(bereich.endContainer).closest('.z') || erste;
        const alle = Array.from(pre.querySelectorAll('.z'));
        const teil = alle.slice(alle.indexOf(erste), alle.indexOf(letzte) + 1).map(z => z.textContent.trimEnd());
        return fertig('code', teil.length === 1 ? teil[0].trim() : teil.join('\n'));
      }
      const ganz = pre.textContent, a = davor(pre), b = a + laenge;
      const von = ganz.lastIndexOf('\n', a - 1) + 1;
      let bis = ganz.indexOf('\n', Math.max(b - 1, a));
      if (bis < 0) bis = ganz.length;
      const teil = ganz.slice(von, bis).split('\n').map(z => z.trimEnd());
      return fertig('code', teil.length === 1 ? teil[0].trim() : teil.join('\n'));
    }
    const zelle = start.closest('td, th');
    if (zelle) {
      const zeile = zelle.parentElement, oben = zeile.closest('table').querySelector('tr');
      const koepfe = oben === zeile ? [] : Array.from(oben.children).map(k => k.tagName === 'TH' ? glatt(k.textContent) : '');
      return fertig('tabelle', Array.from(zeile.children).map((z, i) => (koepfe[i] ? koepfe[i] + ': ' : '') + glatt(z.textContent)).join(' | '));
    }
    const block = start.closest('p, li, dd, dt, h1, h2, h3, h4, figcaption, div');
    if (!block) return null;
    // Ein Programm, das nicht in pre steht: Die Seite kennt ihre Programmzeilen selbst.
    const anfang = glatt(auswahl.toString().split('\n')[0]);
    if (block.tagName === 'DIV' && zeilen && anfang) {
      const zeile = zeilen().find(z => glatt(z).includes(anfang));
      if (zeile) return fertig('code', zeile.trim());
    }
    const ganz = block.textContent, a = davor(block), b = a + laenge;
    if (b > ganz.length) return null;                        // die Markierung reicht über den Block hinaus
    let von = 0, bis = ganz.length, m;
    const ende = /[.!?]\s/g;
    while ((m = ende.exec(ganz))) {
      const p = m.index + 1;
      if (p <= a) von = p; else if (p >= b) { bis = p; break; }
    }
    return fertig('text', glatt(ganz.slice(von, bis)));
  }
  const WO = { code: 'aus dem Programm', text: 'aus dem Satz', tabelle: 'aus der Tabellenzeile' };
  // Die Umgebung als Zeile unter der markierten Stelle, über dem Eingabefeld und in der Frage
  const umgebungHTML = u => u && u.zeigen ? '<span class="umgebung' + (u.art === 'code' ? ' code' : '') + '">' + (u.art === 'code' ? '' : WO[u.art] + ': ') + esc(u.text) + '</span>' : '';

  // ---------- Eine Spalte einrichten ----------

  let gestaltet = false;

  function einrichten(angaben) {
    if (!gestaltet) {
      const stil = document.createElement('style');
      stil.textContent = GESTALTUNG;
      document.head.appendChild(stil);
      gestaltet = true;
      ladeBreite();
    }
    const name = angaben.name;
    const wurzel = document.createElement('aside');
    wurzel.className = 'lotse';
    wurzel.hidden = true;
    wurzel.setAttribute('aria-label', 'Fragen an den ' + name + 'n');
    wurzel.innerHTML = `
      <div class="lotse-griff" role="separator" aria-orientation="vertical" aria-label="Breite der Spalte" tabindex="0" title="Ziehen ändert die Breite der Spalte. Doppelklick setzt sie zurück."></div>
      <div class="lotse-kopf">
        <strong>${esc(name)}</strong>
        <span class="lotse-zusatz"></span>
        <button type="button" class="lotse-kopieren" title="Das Gespräch in die Zwischenablage kopieren">Kopieren</button>
        <button type="button" class="lotse-leeren" title="Das Gespräch leeren und neu beginnen">Leeren</button>
        <button type="button" class="zu lotse-zu" aria-label="Spalte schließen" title="Spalte schließen">×</button>
      </div>
      <div class="lotse-verlauf" aria-live="polite">
        <div class="lotse-hinweis"></div>
        <div class="lotse-gesperrt" hidden></div>
      </div>
      <div class="lotse-fuss">
        <div class="lotse-lage"></div>
        <div class="lotse-markiert" hidden><span class="lotse-markiert-text"></span><span class="lotse-umgebung"></span><button type="button" class="lotse-verwerfen" title="Markierte Stelle nicht mitschicken" aria-label="Markierte Stelle nicht mitschicken">✕</button></div>
        <div class="lotse-eingabe"><textarea aria-label="Frage an das Sprachmodell"></textarea></div>
        <div class="lotse-knoepfe"><button type="button" class="haupt lotse-senden"></button></div>
        <div class="lotse-meldung" aria-live="polite"></div>
      </div>`;
    const auf = document.createElement('button');
    auf.type = 'button';
    auf.className = 'lotse-auf';
    auf.hidden = true;
    auf.textContent = name + 'n fragen';
    document.body.append(wurzel, auf);

    const teil = klasse => wurzel.querySelector('.' + klasse);
    const feld = wurzel.querySelector('textarea'), verlaufFeld = teil('lotse-verlauf'), senden = teil('lotse-senden');
    senden.textContent = angaben.knopf || 'Fragen';
    teil('lotse-hinweis').innerHTML = angaben.hinweis || '';
    feld.placeholder = angaben.platzhalter || '';
    (angaben.weitere || []).forEach(w => {
      const knopf = document.createElement('button');
      knopf.type = 'button';
      knopf.textContent = w.knopf;
      knopf.addEventListener('click', () => stelle({ art: w.art, text: w.text, brauchtMarkierung: w.brauchtMarkierung, mitMarkierung: true }));
      teil('lotse-knoepfe').appendChild(knopf);
    });
    const knoepfe = () => Array.from(teil('lotse-knoepfe').querySelectorAll('button'));

    let verlauf = [];          // das Gespräch für das Sprachmodell
    let abHier = 0;            // ab dieser Nachricht geht das Gespräch an das Sprachmodell
    let lauf = 0;              // zählt die Gespräche: eine Antwort auf ein geleertes Gespräch verfällt
    let langsamUhr;            // meldet nach 15 Sekunden, dass KI:connect langsam antwortet
    let markiert = '';         // die markierte Stelle
    let umgebung = null;       // wo sie steht: { art, text, titel, zeigen }
    let gesperrt = false, wartet = false, da = false, platzhalter = angaben.platzhalter || '';
    const merken = angaben.merken || 'lotseOffen-' + name;

    // ----- kleine Helfer -----

    let meldeUhr;
    function melde(html) {
      teil('lotse-meldung').innerHTML = html;
      clearTimeout(meldeUhr);
      if (html) meldeUhr = setTimeout(() => { teil('lotse-meldung').innerHTML = ''; }, 8000);
    }
    function schalte() {
      feld.disabled = gesperrt;
      knoepfe().forEach(k => { k.disabled = gesperrt || wartet; });
    }
    function beitrag(klasse, html) {
      const d = document.createElement('div');
      d.className = 'beitrag ' + klasse;
      d.innerHTML = html;
      teil('lotse-hinweis').hidden = true;
      verlaufFeld.appendChild(d);
      d.scrollIntoView({ block: 'end' });
      return d;
    }
    function zeigeMarkierung(text, wo) {
      markiert = text || '';
      umgebung = markiert ? wo || null : null;
      teil('lotse-markiert-text').textContent = markiert ? '„' + markiert + '“' : '';
      teil('lotse-umgebung').innerHTML = umgebungHTML(umgebung);
      teil('lotse-markiert').hidden = !markiert;
    }
    // Nimmt die markierte Stelle weg, über dem Eingabefeld und auf der Seite. Bliebe die Markierung auf der Seite
    // stehen, käme sie mit der nächsten Bewegung des Zeigers wieder.
    function verwirf() {
      zeigeMarkierung('');
      if (angaben.markieren) { try { document.getSelection().removeAllRanges(); } catch (fehler) {} }
    }
    // Die markierte Stelle am Anfang der Frage, so wie sie an das Sprachmodell geht
    function alsText(zitat, u) {
      if (!zitat) return '';
      let t = 'Markierte Stelle ' + (angaben.wo || 'auf der Seite') + ': „' + zitat + '“\n';
      if (u && u.zeigen) {
        const ort = u.art === 'code' ? (u.text.includes('\n') ? 'in diesen Programmzeilen' : 'in dieser Programmzeile') : u.art === 'tabelle' ? 'in dieser Tabellenzeile' : 'in diesem Satz';
        t += 'Sie steht ' + ort + (u.titel ? ' unter der Überschrift „' + u.titel + '“' : '') + ':\n' + u.text + '\n';
      }
      return t + '\n';
    }

    // ----- Quellen unter einer Antwort -----

    // Die HTML-Seiten des Mentoriats zu einem Abschnitt des Kurstexts: die ihn erklären, sonst die ihn nennen. Die Liste
    // schreibt das Programm aus der Datei verweise.json. Ein Link öffnet einen neuen Tab, damit das Gespräch stehen bleibt.
    function htmlSeiten(q) {
      const liste = (q.htmlSeiten || []).filter(v => v.adresse);
      if (!liste.length) return '';
      return '<ul>' + liste.map(v => '<li>' + (v.art === 'erklaert' ? 'Im Mentoriat erklärt' : 'Im Mentoriat genannt') + ': <a href="' +
        esc(angaben.link ? angaben.link(v.adresse) : v.adresse) + '" target="_blank" rel="noopener">' + esc(v.titel) + '</a></li>').join('') + '</ul>';
    }
    // Die Abschnitte des Kurstexts unter einer Antwort. Die Liste stammt vom Programm, nicht vom Sprachmodell.
    function quellenHTML(d) {
      if (!d.quellen) return '';
      if (!d.quellen.length) {
        return '<div class="quellen">' + esc(d.hinweis || 'Im Kurstext hat das Programm dazu keine Stelle gefunden. Die Antwort stützt sich also nicht auf den Kurstext. Frag im Forum oder im Mentoriat nach, wenn du sicher sein willst.') + '</div>';
      }
      return '<div class="quellen">' + (d.zurFrage ? 'Dazu' : esc(angaben.sonst || 'Dazu')) + ' im Kurstext:<ul>' + d.quellen.map(q =>
        '<li>Abschnitt ' + esc(q.abschnitt + ' ' + q.titel) + (q.seiten ? ', ' + (/[–-]/.test(q.seiten) ? 'Seiten ' : 'Seite ') + esc(q.seiten) : '') + htmlSeiten(q) + '</li>').join('') + '</ul></div>';
    }
    const antwortHTML = (d, mitQuellen) => '<span class="wer">' + esc(name) + ' · Sprachmodell, kann sich irren</span>' + alsHTML(d.antwort) + (mitQuellen ? quellenHTML(d) : '');

    // ----- Fragen -----

    function schicke(daten) {
      return hole(typeof angaben.adresse === 'function' ? angaben.adresse() : angaben.adresse, daten);
    }

    // Stellt eine Frage im Gespräch. auftrag.text fehlt, wenn der Student die Frage getippt hat.
    async function stelle(auftrag) {
      if (gesperrt) return;
      if (wartet) { melde(esc(name) + ' antwortet noch auf die letzte Frage.'); return; }
      const mitMarkierung = auftrag.text === undefined || auftrag.mitMarkierung;
      const zitat = mitMarkierung ? markiert : '', wo = mitMarkierung ? umgebung : null;
      const getippt = auftrag.text === undefined ? feld.value.trim() : '';
      const text = auftrag.text || getippt || (zitat ? 'Erkläre mir diese Stelle.' : '');
      if (auftrag.brauchtMarkierung && !zitat) { melde('Markiere erst eine Stelle.'); return; }
      if (!text) { melde(angaben.markieren || angaben.weitere ? 'Markiere eine Stelle oder stelle eine Frage.' : 'Tippe erst eine Frage.'); return; }
      melde('');
      const meiner = lauf, art = auftrag.art || 'frage';
      wartet = true;
      schalte();
      const frage = { art: art, markiert: zitat, umgebung: wo ? wo.text : '', markiertCode: !!(wo && wo.art === 'code'), getippt: getippt };
      let warte = null;
      try {
        const mehr = angaben.mitschicken ? await angaben.mitschicken(frage) : {};
        verlauf.push({ role: 'user', content: alsText(zitat, wo) + text });
        const marke = angaben.marke ? angaben.marke() : '';
        beitrag('student', (marke ? '<span class="stand">' + esc(marke) + '</span>' : '') + (zitat ? '<div class="zitat">„' + esc(zitat) + '“' + umgebungHTML(wo) + '</div>' : '') + esc(text));
        if (mitMarkierung) { feld.value = ''; verwirf(); }
        warte = beitrag('wartet', esc(name) + ' sucht im Kurstext und antwortet …');
        // Nach 15 Sekunden ohne Antwort sieht das Warten wie ein Absturz aus. Meist ist KI:connect nur langsam.
        const langsam = warte;
        langsamUhr = setTimeout(() => {
          if (langsam.className === 'beitrag wartet') langsam.innerHTML += '<br>KI:connect antwortet gerade langsam. Du kannst weiterarbeiten, die Antwort erscheint hier.';
        }, 15000);
        const d = await schicke(Object.assign(frage, { verlauf: verlauf.slice(abHier) }, auftrag.extra, mehr));
        if (meiner !== lauf) return;
        verlauf.push({ role: 'assistant', content: d.antwort });
        warte.className = 'beitrag modell';
        warte.innerHTML = antwortHTML(d, true);
        warte.scrollIntoView({ block: 'nearest' });
      } catch (fehler) {
        if (meiner !== lauf) return;
        if (warte) {
          verlauf.pop();                             // die Frage zählt nicht, der Student stellt sie neu
          warte.className = 'beitrag fehler';
          warte.textContent = 'Die Frage ist nicht angekommen: ' + fehler.message;
          // Frage und markierte Stelle stehen wieder da, wenn der Student inzwischen nichts Neues begonnen hat.
          if (mitMarkierung && !feld.value) feld.value = getippt;
          if (mitMarkierung && zitat && !markiert) zeigeMarkierung(zitat, wo);
        } else {
          melde(esc(fehler.message));
        }
      } finally {
        clearTimeout(langsamUhr);
        if (meiner === lauf) { wartet = false; schalte(); if (!gesperrt && mitMarkierung) feld.focus(); }
      }
    }

    // Stellt eine Frage außerhalb des Gesprächs, zum Beispiel für ein Urteil über eine Antwort des Studenten.
    async function einzeln(auftrag) {
      const frage = { art: auftrag.art, markiert: '', umgebung: '', markiertCode: false, getippt: '' };
      const mehr = angaben.mitschicken ? await angaben.mitschicken(frage) : {};
      const d = await schicke(Object.assign(frage, { verlauf: [{ role: 'user', content: auftrag.text }] }, auftrag.extra, mehr));
      return antwortHTML(d, auftrag.mitQuellen !== false);
    }

    function leere() {
      lauf++;
      verlauf = [];
      abHier = 0;
      wartet = false;
      verlaufFeld.querySelectorAll('.beitrag').forEach(b => b.remove());
      teil('lotse-hinweis').hidden = false;
      feld.value = '';
      zeigeMarkierung('');
      melde('');
      schalte();
      if (angaben.beimLeeren) angaben.beimLeeren();
    }

    // Das Gespräch als Text: oben die Kopfzeilen der Seite, darunter Fragen und Antworten
    async function kopiere() {
      const knopf = teil('lotse-kopieren');
      const sage = wort => { knopf.textContent = wort; setTimeout(() => { knopf.textContent = 'Kopieren'; }, 1600); };
      if (!verlauf.length) { sage('Noch leer'); return; }
      const kopf = angaben.kopfzeile ? angaben.kopfzeile() : [];
      const text = kopf.concat(kopf.length ? [''] : [], verlauf.map(n => (n.role === 'user' ? 'Du: ' : name + ': ') + n.content + '\n')).join('\n');
      try {
        await navigator.clipboard.writeText(text);
        sage('Kopiert');
      } catch (fehler) {
        sage('Geht hier nicht');
      }
    }

    // ----- Auf und zu -----

    // Die Spalte auf- und zuklappen. Die Wahl merkt sich der Browser für den nächsten Besuch.
    function oeffne(offen, merke) {
      wurzel.hidden = !offen;
      auf.hidden = offen;
      document.body.classList.toggle('lotse-offen', offen);
      if (merke) { try { localStorage.setItem(merken, offen ? '1' : '0'); } catch (fehler) {} }
      if (offen && merke && !gesperrt) feld.focus();
    }
    // Die Spalte gibt es oder gibt es nicht. Gibt es sie, steht sie so da, wie der Student sie zuletzt gelassen hat.
    function zeige(an, offen) {
      da = an;
      if (!an) { wurzel.hidden = true; auf.hidden = true; document.body.classList.remove('lotse-offen'); return; }
      let gemerkt = null;
      try { gemerkt = localStorage.getItem(merken); } catch (fehler) {}
      oeffne(offen !== undefined ? offen : gemerkt === null ? window.innerWidth >= 1100 : gemerkt === '1', false);
    }
    function sperre(html, ersatz) {
      gesperrt = !!html;
      wurzel.classList.toggle('gesperrt', gesperrt);
      teil('lotse-gesperrt').innerHTML = html || '';
      teil('lotse-gesperrt').hidden = !gesperrt;
      feld.placeholder = gesperrt ? ersatz || '' : platzhalter;
      if (gesperrt) zeigeMarkierung('');
      schalte();
    }

    // ----- Was der Student auf der Seite markiert -----

    // Die Stelle erscheint über dem Eingabefeld und geht mit der nächsten Frage mit. Sie bleibt dort stehen, auch wenn
    // die Markierung auf der Seite beim Klick ins Eingabefeld verschwindet.
    let markierUhr;
    if (angaben.markieren) {
      document.addEventListener('selectionchange', () => {
        clearTimeout(markierUhr);
        markierUhr = setTimeout(() => {
          const auswahl = document.getSelection();
          const text = auswahl.toString().replace(/[ \t]+\n/g, '\n').trim().slice(0, 1500);
          const knoten = auswahl.anchorNode && (auswahl.anchorNode.nodeType === 1 ? auswahl.anchorNode : auswahl.anchorNode.parentElement);
          if (!text || !knoten || knoten.closest('.lotse') || wurzel.hidden || gesperrt || !angaben.markieren(knoten)) return;
          let wo = null;
          try { wo = umgebungVon(auswahl, angaben.zeilen); } catch (fehler) {}
          zeigeMarkierung(text, wo);
        }, 250);
      });
    }

    senden.addEventListener('click', () => stelle({}));
    feld.addEventListener('keydown', ev => { if (ev.key === 'Enter' && !ev.shiftKey) { ev.preventDefault(); stelle({}); } });
    teil('lotse-kopieren').addEventListener('click', kopiere);
    teil('lotse-leeren').addEventListener('click', () => { leere(); if (!gesperrt) feld.focus(); });
    teil('lotse-verwerfen').addEventListener('click', verwirf);
    teil('lotse-zu').addEventListener('click', () => oeffne(false, true));
    richteGriffEin(teil('lotse-griff'));
    auf.addEventListener('click', () => oeffne(true, true));
    schalte();

    return {
      zeige: zeige,
      oeffne: offen => { if (da) oeffne(offen, false); },
      sperre: sperre,
      leere: leere,
      markiere: (text, wo) => { if (!gesperrt) zeigeMarkierung(String(text || '').slice(0, 1500), text ? wo || null : null); },
      frage: auftrag => stelle(auftrag),
      einzeln: einzeln,
      neuerStand: () => { abHier = verlauf.length; },
      setzeLage: html => { teil('lotse-lage').innerHTML = html || ''; },
      setzeHinweis: html => { teil('lotse-hinweis').innerHTML = html || ''; },
      setzePlatzhalter: text => { platzhalter = text || ''; if (!gesperrt) feld.placeholder = platzhalter; },
      setzeZusatz: text => { teil('lotse-zusatz').textContent = text || ''; },
      melde: melde,
    };
  }

  // hole schickt eine Anfrage an den Lotsen und liefert seine Antwort. Mit daten geht sie als POST mit JSON. Meldet der
  // Lotse einen Fehler, wirft hole ihn mit dem Text des Lotsen, sonst mit der Nummer des Fehlers.
  async function hole(pfad, daten) {
    const r = await fetch(pfad, daten ? { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify(daten) } : undefined);
    const d = await r.json().catch(() => null);
    if (!r.ok || d === null) throw new Error((d && d.fehler) || 'Fehler ' + r.status);
    return d;
  }

  window.Spalte = { einrichten: einrichten, alsHTML: alsHTML, esc: esc, hole: hole };
})();
