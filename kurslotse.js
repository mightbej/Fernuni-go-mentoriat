// Dialog mit dem Kurslotsen: eine Spalte rechts neben der Seite, in der der Student Fragen zum sichtbaren Schritt stellt.
//
// Der Kurslotse ist ein Programm auf dem Rechner des Studenten. Es reicht die Frage an ein Sprachmodell weiter.
// Läuft es nicht, erscheint die Spalte nicht, und die Seite sieht aus wie ohne diese Datei.
//
// Eine Seite bindet die Datei ein und meldet sich an:
//
//   Kurslotse.einrichten({
//     titel: 'Bitverschiebung: Schieben, verdoppeln und halbieren in Go',
//     bezug: 'Kurstext v3.0.6, Abschnitt 4.3.7.2.2, Seiten 75–78',
//     beschreibung: '../kurslotse/kap4-primitive-datentypen/go_k4_bitverschiebung_interaktiv.md',
//     reiter: () => 'schieben',                       // Anker des sichtbaren Schritts, wie in der Beschreibung hinter „Reiter:“
//     schritt: () => 'Schritt 1: Ein Bitmuster schieben',
//     sichtbar: () => document.getElementById('stufe-1'),
//     stand: () => ['a hat den Wert 13.'],            // der Stand der Demo in Sätzen
//     kurz: () => 'a = 13 als uint8, n = 1',          // dasselbe in einer Zeile, steht über dem Eingabefeld
//     beispiel: 'Warum rückt hier links eine Eins nach?'
//   });
//
// Dazu nach Bedarf: thema (kurzer Name für die Suche), text (Auswahl der Elemente, deren Text mitgeht),
// zeilen und zeilenTitel (die Programmzeilen und ihre Überschrift, wenn das Programm nicht in pre steht).
//
// Wechselt der Student den Schritt, ruft die Seite Kurslotse.schrittGewechselt() auf: Der Dialog leert sich.
// Ändert er einen Wert der Demo, ruft sie Kurslotse.standGeaendert() auf.
(function () {
  const PORT = '8614';
  // Liefert der Kurslotse die Seite selbst aus, genügt die eigene Adresse.
  const LOTSE = location.protocol === 'http:' && location.port === PORT ? '' : 'http://127.0.0.1:' + PORT;
  // Diese Teile des sichtbaren Schritts gehen als Text mit, wenn die Seite nichts anderes nennt.
  const TEXT = '.card > p:not(.quelle), .legende, .hinweis, .erklaerung, .ergebnis, table.text';

  const el = id => document.getElementById(id);
  const esc = t => String(t).replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
  let seite = null;
  let verlauf = [];
  let lauf = 0;                                    // zählt die Dialoge: eine Antwort auf einen geleerten Dialog verfällt
  const beschreibung = { seite: '', schritte: {} };

  const GESTALTUNG = `
    .lotse { position:fixed;top:0;right:0;bottom:0;width:420px;max-width:100%;display:flex;flex-direction:column;background:#fff;border-left:1px solid #e5e7eb;box-shadow:-6px 0 18px rgba(15,23,42,.08);z-index:20;font-family:"Segoe UI",Tahoma,Geneva,Verdana,sans-serif;color:#1f2937;line-height:1.6; }
    .lotse[hidden],.lotse-auf[hidden],.lotse-zeile[hidden] { display:none; }
    .lotse button { border:1px solid #e5e7eb;border-radius:8px;font-weight:600;cursor:pointer;font-family:inherit; }
    .lotse-kopf { display:flex;align-items:center;gap:8px;padding:12px 16px;background:linear-gradient(135deg,#0f766e 0%,#1e40af 100%);color:#fff; }
    .lotse-kopf strong { font-size:1.1rem; }
    .lotse-kopf span { flex:1;font-size:.82rem;opacity:.9;white-space:nowrap;overflow:hidden;text-overflow:ellipsis; }
    .lotse-kopf button { background:transparent;color:#fff;border-color:rgba(255,255,255,.5);padding:6px 11px;font-size:.85rem; }
    .lotse-kopf button.zu { padding:4px 11px;font-size:1.1rem;line-height:1.2; }
    .lotse-kopf button:hover { background:rgba(255,255,255,.18); }
    .lotse-verlauf { flex:1;overflow-y:auto;display:flex;flex-direction:column;gap:10px;padding:14px 16px; }
    .lotse-hinweis { color:#4b5563;font-size:.92rem;margin:0; }
    .lotse .beitrag { padding:10px 14px;border-radius:10px;max-width:94%;overflow-wrap:anywhere;font-size:.95rem; }
    .lotse .beitrag.student { align-self:flex-end;background:#dbeafe;border:1px solid #bfdbfe;white-space:pre-wrap; }
    .lotse .beitrag.modell { align-self:flex-start;background:#f0fdfa;border:1px solid #99f6e4; }
    .lotse .beitrag.fehler { align-self:flex-start;background:#fef2f2;border:1px solid #fecaca;color:#991b1b; }
    .lotse .beitrag.wartet { align-self:flex-start;color:#4b5563;font-style:italic; }
    .lotse .beitrag .stand { display:block;font-size:.8rem;color:#4b5563;margin-bottom:4px;font-family:inherit;font-weight:400; }
    .lotse .beitrag pre { margin:8px 0;background:#1e293b;color:#e2e8f0;border-radius:8px;padding:10px;font-family:Consolas,monospace;font-size:.85rem;line-height:1.5;white-space:pre-wrap; }
    .lotse .beitrag code { background:#f3f4f6;padding:2px 6px;border-radius:6px;font-family:Consolas,"Courier New",monospace;font-size:.9rem; }
    .lotse .beitrag ul,.lotse .beitrag ol { margin:6px 0;padding-left:20px; }
    .lotse .beitrag li { margin:3px 0; }
    .lotse .beitrag p { margin:0 0 8px; } .lotse .beitrag p:last-child { margin:0; }
    .lotse .beitrag .quellen { margin-top:10px;padding-top:6px;border-top:1px solid #99f6e4;font-size:.82rem;color:#4b5563; }
    .lotse-fuss { border-top:1px solid #e5e7eb;padding:10px 16px 14px;background:#f8fafc; }
    .lotse-stand { font-size:.82rem;color:#4b5563;margin-bottom:8px; }
    .lotse-stand b { color:#1f2937;font-weight:600; }
    .lotse-eingabe { display:flex;gap:8px;align-items:flex-end; }
    .lotse-eingabe textarea { flex:1;min-height:72px;border:1px solid #e5e7eb;border-radius:8px;padding:10px;font-family:inherit;font-size:.95rem;resize:vertical; }
    .lotse-eingabe textarea:focus { outline:2px solid #0f766e;outline-offset:2px; }
    .lotse-eingabe button { background:#0f766e;color:#fff;border-color:#0f766e;padding:9px 14px;font-size:.95rem; }
    .lotse-eingabe button:disabled { opacity:.5;cursor:default; }
    .lotse-auf { position:fixed;right:18px;bottom:18px;z-index:19;background:#0f766e;color:#fff;border:1px solid #0f766e;border-radius:8px;font-weight:600;font-size:.95rem;font-family:"Segoe UI",Tahoma,Geneva,Verdana,sans-serif;cursor:pointer;box-shadow:0 6px 16px rgba(15,118,110,.35);padding:11px 16px; }
    .lotse-auf:hover { background:#115e59; }
    /* Solange der Kurslotse nicht verbunden ist, steht an derselben Stelle der Knopf „Kurslotsen verbinden“, darüber eine Meldung. */
    .lotse-zeile { position:fixed;right:18px;bottom:18px;z-index:19;display:flex;flex-direction:column;align-items:flex-end;gap:8px;max-width:320px;font-family:"Segoe UI",Tahoma,Geneva,Verdana,sans-serif; }
    .lotse-zeile .lotse-auf { position:static;background:#fff;color:#0f766e;box-shadow:0 6px 16px rgba(15,23,42,.12); }
    .lotse-zeile .lotse-auf:hover { background:#ecfdf5; }
    .lotse-status { background:#fff;border:1px solid #e5e7eb;border-radius:8px;padding:8px 12px;font-size:.88rem;line-height:1.5;color:#4b5563;box-shadow:0 6px 16px rgba(15,23,42,.12); }
    .lotse-status:empty { display:none; }
    /* In einem breiten Fenster rückt die Seite nach links und macht der Spalte Platz. In einem schmalen liegt die Spalte über der Seite. */
    @media (min-width:1100px) { body.lotse-offen { padding-right:420px; } }`;

  const SPALTE = `
    <aside class="lotse" id="lotse" aria-label="Fragen an den Kurslotsen" hidden>
      <div class="lotse-kopf">
        <strong>Kurslotse</strong>
        <span id="lotseModell"></span>
        <button type="button" id="lotseKopieren" title="Den Dialog in die Zwischenablage kopieren">Kopieren</button>
        <button type="button" id="lotseLeeren" title="Den Dialog leeren und neu beginnen">Leeren</button>
        <button type="button" class="zu" id="lotseZu" aria-label="Dialog einklappen" title="Dialog einklappen">×</button>
      </div>
      <div class="lotse-verlauf" id="lotseVerlauf" aria-live="polite">
        <p class="lotse-hinweis" id="lotseHinweis">Hier stellst du Fragen zu dem Schritt, der links zu sehen ist. Die Frage geht an das Sprachmodell von KI:connect. Die Seite schickt den Stand der Demo mit: Wer links etwas umstellt und danach fragt, bekommt eine Antwort zum neuen Stand. Wer oben den Schritt wechselt, beginnt einen neuen Dialog. „Kopieren“ legt den Dialog vorher in die Zwischenablage.</p>
      </div>
      <div class="lotse-fuss">
        <div class="lotse-stand" id="lotseStand"></div>
        <div class="lotse-eingabe">
          <textarea id="lotseText" aria-label="Frage an das Sprachmodell"></textarea>
          <button type="button" id="lotseSenden">Fragen</button>
        </div>
      </div>
    </aside>
    <button type="button" class="lotse-auf" id="lotseAuf" hidden>Kurslotsen fragen</button>`;

  // ---------- Was die Seite mitschickt ----------

  // Die Beschreibung der Seite: eine Datei im Ordner kurslotse des Repositorys, je Reiter ein Abschnitt mit
  // Aufbau der Darstellung, Handgriffen und Abschnitten des Kurstexts. Fehlt die Datei, antwortet der Kurslotse
  // nur mit dem sichtbaren Text und dem Stand der Demo.
  async function ladeBeschreibung() {
    if (!seite.beschreibung) return;
    try {
      const r = await fetch(seite.beschreibung);
      if (!r.ok) return;
      const teile = (await r.text()).replace(/\r/g, '').split(/^## /m);
      beschreibung.seite = teile[0].replace(/^# .*$/m, '').trim();
      teile.slice(1).forEach(t => {
        const reiter = t.match(/^Reiter:\s*(\S+)\s*$/m);
        if (reiter) beschreibung.schritte[reiter[1]] = t.replace(/^.*\n/, '').replace(/^Reiter:.*$/m, '').trim();
      });
    } catch (fehler) {}
  }
  const schrittText = () => beschreibung.schritte[seite.reiter()] || '';
  // Die Abschnitte des Kurstexts, auf die sich der sichtbare Schritt stützt: die Nummern unter der Überschrift dazu
  const abschnitte = () => ((schrittText().split(/^### Abschnitte des Kurstexts\s*$/m)[1] || '').split(/^### /m)[0].match(/\d+(\.\d+)+/g)) || [];
  // Die Programmzeilen des sichtbaren Schritts, so wie sie auf der Seite stehen.
  // Eine Seite, deren Programm nicht in pre steht, liefert sie selbst mit zeilen.
  const zeilen = () => seite.zeilen ? seite.zeilen() : Array.from(seite.sichtbar().querySelectorAll('pre'))
    .flatMap(p => p.textContent.split('\n')).map(t => t.trimEnd()).filter(Boolean);

  function kontext() {
    // Kopie des sichtbaren Schritts neben dem Bildschirm, damit aus der Hochzahl in 2³ der Text 2^3 wird
    const kopie = seite.sichtbar().cloneNode(true);
    kopie.removeAttribute('id');
    kopie.style.cssText = 'position:absolute;left:-10000px;width:900px';
    kopie.querySelectorAll('sup').forEach(s => s.replaceWith('^' + s.textContent));
    document.body.appendChild(kopie);
    const text = Array.from(kopie.querySelectorAll(seite.text || TEXT)).map(e => e.innerText.trim()).filter(Boolean).join('\n\n');
    kopie.remove();
    const z = zeilen();
    return ['HTML-Seite: ' + seite.titel, 'Bezug: ' + seite.bezug, 'Sichtbarer Schritt: ' + seite.schritt()].concat(
      beschreibung.seite ? ['', 'Was die HTML-Seite leistet:', beschreibung.seite] : [],
      schrittText() ? ['', 'Beschreibung des sichtbaren Schritts:', schrittText().replace(/^### /gm, '')] : [],
      ['', 'Stand der Demo:'], seite.stand(),
      z.length ? ['', seite.zeilenTitel || 'Programmzeilen des sichtbaren Schritts, in Go. Hinter // steht die Ausgabe:'].concat(z) : [],
      ['', 'Text des sichtbaren Schritts:', text.slice(0, 6000)]).join('\n');
  }

  // ---------- Darstellung ----------

  // Text ohne Codeblock: Zeilen mit „- “ oder „1. “ werden zu einer Liste, die übrigen zu Absätzen
  function absaetze(text) {
    // `Code` kommt zuerst beiseite: Ein Stern darin ist in Go ein Zeiger, zum Beispiel `*auto`, und keine Auszeichnung.
    const fein = z => {
      const code = [];
      return esc(z).replace(/`([^`\n]+)`/g, (ganz, t) => '\u0000' + (code.push(t) - 1) + '\u0000')
        .replace(/\*\*([^*\n]+)\*\*/g, '<strong>$1</strong>').replace(/\*([^*\n]+)\*/g, '<em>$1</em>')
        .replace(/\u0000(\d+)\u0000/g, (ganz, i) => '<code>' + code[Number(i)] + '</code>');
    };
    let html = '', liste = '';
    const schliesse = () => { if (liste) { html += '</' + liste + '>'; liste = ''; } };
    text.split('\n').forEach(zeile => {
      const z = zeile.trim(), punkt = z.match(/^[-*•]\s+(.*)$/), nummer = z.match(/^\d+[.)]\s+(.*)$/);
      if (!z) { schliesse(); return; }
      const art = punkt ? 'ul' : nummer ? 'ol' : '';
      if (art !== liste) { schliesse(); if (art) { html += '<' + art + '>'; liste = art; } }
      html += art ? '<li>' + fein((punkt || nummer)[1]) + '</li>' : '<p>' + fein(z.replace(/^#+\s*/, '')) + '</p>';
    });
    schliesse();
    return html;
  }
  // Antwort des Sprachmodells: Codeblöcke, Listen, `Code` und **fett** darstellen, alles andere als Text
  function alsHTML(t) {
    return t.split('```').map((teil, i) => i % 2
      ? '<pre>' + esc(teil.replace(/^[a-z]*\n/, '').replace(/\n$/, '')) + '</pre>'
      : absaetze(teil)).join('');
  }
  function beitrag(klasse, html) {
    const d = document.createElement('div');
    d.className = 'beitrag ' + klasse;
    d.innerHTML = html;
    el('lotseHinweis').hidden = true;
    el('lotseVerlauf').appendChild(d);
    d.scrollIntoView({ block: 'end' });
    return d;
  }
  // Die Zeile über dem Eingabefeld: welchen Stand die nächste Frage mitschickt
  function zeigeStand() {
    if (!seite || !el('lotseStand')) return;
    const kurz = seite.kurz ? seite.kurz() : '';
    el('lotseStand').innerHTML = 'Die nächste Frage schickt mit: <b>' + esc(seite.schritt().split(':')[0]) + '</b>' + (kurz ? ', ' + esc(kurz) : '');
  }
  // Die Spalte auf- und zuklappen. Die Wahl merkt sich der Browser für den nächsten Besuch.
  function klappe(offen, merken) {
    el('lotse').hidden = !offen;
    el('lotseAuf').hidden = offen;
    document.body.classList.toggle('lotse-offen', offen);
    if (merken) { try { localStorage.setItem('lotseOffen', offen ? '1' : '0'); } catch (fehler) {} }
    if (offen && merken) el('lotseText').focus();
  }

  // ---------- Dialog ----------

  async function frage() {
    const feld = el('lotseText'), text = feld.value.trim();
    if (!text || el('lotseSenden').disabled) return;
    const k = kontext(), meiner = lauf;
    verlauf.push({ role: 'user', content: text });
    beitrag('student', '<span class="stand">' + esc(seite.schritt()) + '</span>' + esc(text));
    feld.value = '';
    el('lotseSenden').disabled = true;
    const wartet = beitrag('wartet', 'Das Sprachmodell antwortet …');
    try {
      const r = await fetch(LOTSE + '/api/frage', { method: 'POST', headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ kontext: k, zeilen: zeilen(), abschnitte: abschnitte(), thema: (seite.thema || seite.titel) + ', ' + seite.schritt(), verlauf: verlauf }) });
      const d = await r.json();
      if (meiner !== lauf) return;
      if (!r.ok) throw new Error(d.fehler || 'Status ' + r.status);
      verlauf.push({ role: 'assistant', content: d.antwort });
      wartet.className = 'beitrag modell';
      wartet.innerHTML = alsHTML(d.antwort) + (d.quellen && d.quellen.length
        ? '<div class="quellen">' + (d.zurFrage ? 'Dazu im Kurstext:' : 'Zu diesem Schritt im Kurstext:') + '<ul>' + d.quellen.map(q => '<li>Abschnitt ' + esc(q.abschnitt + ' ' + q.titel) + ', ' + (/[–-]/.test(q.seiten) ? 'Seiten ' : 'Seite ') + esc(q.seiten) + '</li>').join('') + '</ul></div>' : '');
      wartet.scrollIntoView({ block: 'nearest' });
    } catch (fehler) {
      if (meiner !== lauf) return;
      verlauf.pop();                               // die Frage zählt nicht, der Student stellt sie neu
      wartet.className = 'beitrag fehler';
      wartet.textContent = 'Die Frage ist nicht angekommen: ' + fehler.message;
      feld.value = text;
    }
    el('lotseSenden').disabled = false;
    feld.focus();
  }
  function leere() {
    if (!seite || !el('lotseVerlauf')) return;
    lauf++;
    verlauf = [];
    el('lotseVerlauf').querySelectorAll('.beitrag').forEach(b => b.remove());
    el('lotseHinweis').hidden = false;
    el('lotseSenden').disabled = false;
  }
  // Der Dialog als Text: oben Seite, Schritt und Stand der Demo, darunter Fragen und Antworten
  async function kopiere() {
    const knopf = el('lotseKopieren');
    const melde = wort => { knopf.textContent = wort; setTimeout(() => { knopf.textContent = 'Kopieren'; }, 1600); };
    if (!verlauf.length) { melde('Noch leer'); return; }
    const kurz = seite.kurz ? seite.kurz() : '';
    const text = [(seite.thema || seite.titel) + ', ' + seite.schritt()].concat(kurz ? ['Stand der Demo: ' + kurz] : [], [''],
      verlauf.map(n => (n.role === 'user' ? 'Frage: ' : 'Kurslotse: ') + n.content + '\n')).join('\n');
    try {
      await navigator.clipboard.writeText(text);
      melde('Kopiert');
    } catch (fehler) {
      melde('Geht hier nicht');
    }
  }

  // Liefert der Kurslotse die Seite selbst aus, sucht die Seite ihn gleich. Von GitHub Pages aus sucht sie ihn erst,
  // wenn der Student einmal „Kurslotsen verbinden“ gewählt hat: Der Browser kann beim Zugriff auf den eigenen Rechner
  // nachfragen, und diese Rückfrage soll niemand sehen, der den Kurslotsen nicht eingerichtet hat.
  async function sucheLotsen(aufWunsch) {
    try {
      const d = await (await fetch(LOTSE + '/api/bereit', { signal: AbortSignal.timeout(4000) })).json();
      if (!d.schluessel) {
        el('lotseStatus').innerHTML = 'Der Kurslotse läuft, hat aber keinen Schlüssel für KI:connect. Der Schlüssel gehört in die Umgebungsvariable KICONNECT_KEY. ' +
          '<a href="http://localhost:' + PORT + '/" target="_blank" rel="noopener">Anleitung öffnen</a>';
        return;
      }
      el('lotseZeile').hidden = true;
      el('lotseModell').textContent = 'Modell: ' + d.modell;
      let gemerkt = null;
      try { gemerkt = localStorage.getItem('lotseOffen'); localStorage.setItem('lotseVerbunden', '1'); } catch (fehler) {}
      klappe(aufWunsch || (gemerkt === null ? window.innerWidth >= 1100 : gemerkt === '1'), false);
    } catch (fehler) {
      if (aufWunsch) el('lotseStatus').textContent = 'Der Kurslotse antwortet auf diesem Rechner nicht. Er muss laufen, und der Browser muss den Zugriff auf den eigenen Rechner erlauben.';
    }
  }

  // ---------- Anmeldung der Seite ----------

  function einrichten(angaben) {
    seite = angaben;
    const stil = document.createElement('style');
    stil.textContent = GESTALTUNG;
    document.head.appendChild(stil);
    document.body.insertAdjacentHTML('beforeend', SPALTE);
    const zeile = document.createElement('div');
    zeile.className = 'lotse-zeile';
    zeile.id = 'lotseZeile';
    zeile.innerHTML = '<div class="lotse-status" id="lotseStatus" aria-live="polite"></div><button type="button" class="lotse-auf" id="lotseVerbinden">Kurslotsen verbinden</button>';
    document.body.appendChild(zeile);
    el('lotseText').placeholder = 'Zum Beispiel: ' + (seite.beispiel || 'Was sehe ich hier?');

    el('lotseSenden').addEventListener('click', frage);
    el('lotseText').addEventListener('keydown', ev => { if (ev.key === 'Enter' && !ev.shiftKey) { ev.preventDefault(); frage(); } });
    el('lotseKopieren').addEventListener('click', kopiere);
    el('lotseLeeren').addEventListener('click', () => { leere(); el('lotseText').focus(); });
    el('lotseZu').addEventListener('click', () => klappe(false, true));
    el('lotseAuf').addEventListener('click', () => klappe(true, true));
    el('lotseVerbinden').addEventListener('click', () => { el('lotseStatus').textContent = ''; sucheLotsen(true); });

    zeigeStand();
    ladeBeschreibung();
    let verbunden = null;
    try { verbunden = localStorage.getItem('lotseVerbunden'); } catch (fehler) {}
    if (LOTSE === '' || verbunden === '1') sucheLotsen(false);
  }

  window.Kurslotse = {
    einrichten: einrichten,
    schrittGewechselt: () => { leere(); zeigeStand(); },
    standGeaendert: zeigeStand
  };
})();
