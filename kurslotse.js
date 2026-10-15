// Dialog mit dem Kurslotsen: eine Leiste rechts neben der Seite, in der der Student Fragen zum sichtbaren Schritt stellt.
//
// Der Kurslotse ist ein Programm auf dem Rechner des Studenten. Es reicht die Frage an ein Sprachmodell weiter.
// Läuft es nicht, erscheint die Leiste nicht, und die Seite sieht aus wie ohne diese Datei.
//
// Die Leiste selbst steht in der Datei lotsenleiste.js neben dieser Datei. Sie ist für alle Lotsen dieselbe. In dieser Datei
// steht, was nur die HTML-Seiten des Mentoriats betrifft: was die Seite mit einer Frage mitschickt und wie sie den
// Kurslotsen auf dem Rechner des Studenten findet.
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
  // Liefert ein Kurslotse die Seite selbst aus, genügt die eigene Adresse, unter welchem Port er auch läuft.
  const LOKAL = location.protocol === 'http:' && (location.hostname === 'localhost' || location.hostname === '127.0.0.1');
  const LOTSE = LOKAL ? '' : 'http://127.0.0.1:' + PORT;
  // Das Hauptverzeichnis der HTML-Seiten: der Ordner, in dem diese Datei liegt. Bei GitHub Pages und beim Kurslotsen
  // auf dem eigenen Rechner ist das eine andere Adresse.
  const WURZEL = new URL('.', document.currentScript.src).href;
  // Diese Teile des sichtbaren Schritts gehen als Text mit, wenn die Seite nichts anderes nennt.
  const TEXT = '.card > p:not(.quelle), .legende, .hinweis, .erklaerung, .ergebnis, table.text';

  const el = id => document.getElementById(id);
  const esc = t => String(t).replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
  let seite = null;          // die Angaben der HTML-Seite
  let leiste = null;         // die Leiste aus lotsenleiste.js, sobald die Datei geladen ist
  let standZuletzt = null;   // der Stand der Demo bei der letzten Frage, als ein Text
  const beschreibung = { seite: '', schritte: {} };

  // Solange der Kurslotse nicht verbunden ist, steht unten rechts der Knopf „Kurslotsen verbinden“, darüber eine Meldung.
  const GESTALTUNG = `
    .lotse-zeile { position:fixed;right:18px;bottom:18px;z-index:19;display:flex;flex-direction:column;align-items:flex-end;gap:8px;max-width:320px;font-family:"Segoe UI",Tahoma,Geneva,Verdana,sans-serif; }
    .lotse-zeile[hidden] { display:none; }
    .lotse-zeile .lotse-auf { position:static;background:#fff;color:#0f766e;box-shadow:0 6px 16px rgba(15,23,42,.12); }
    .lotse-zeile .lotse-auf:hover { background:#ecfdf5; }
    .lotse-status { background:#fff;border:1px solid #e5e7eb;border-radius:8px;padding:8px 12px;font-size:.88rem;line-height:1.5;color:#4b5563;box-shadow:0 6px 16px rgba(15,23,42,.12); }
    .lotse-status:empty { display:none; }`;

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

  // Was mit jeder Frage an den Kurslotsen geht. Hat der Student die Demo seit der letzten Frage verändert, gehen die
  // früheren Fragen und Antworten nicht mehr mit. Sonst antwortet das Sprachmodell zum Stand, von dem sie handeln.
  // Auf dem Bildschirm bleiben sie stehen.
  const datei = adresse => new URL(adresse, location.href).pathname.split('/').pop();
  function mitschicken() {
    const jetzt = seite.stand().join('\n');
    if (standZuletzt !== null && jetzt !== standZuletzt) leiste.neuerStand();
    standZuletzt = jetzt;
    return { kontext: kontext(), zeilen: zeilen(), abschnitte: abschnitte(), thema: (seite.thema || seite.titel) + ', ' + seite.schritt(), seite: datei(location.href) };
  }

  // Ein Link unter einer Antwort zeigt auf GitHub Pages. Er führt in das Hauptverzeichnis, aus dem diese Seite stammt:
  // Liefert der Kurslotse die Seiten selbst aus, bleibt der Student auf dem eigenen Rechner.
  function link(adresse) {
    const pfad = adresse.split('/Fernuni-go-mentoriat/')[1];
    return pfad ? new URL(pfad, WURZEL).href : new URL(adresse).href;
  }

  // Die Zeile über dem Eingabefeld: welchen Stand die nächste Frage mitschickt
  function zeigeStand() {
    if (!seite || !leiste) return;
    const kurz = seite.kurz ? seite.kurz() : '';
    leiste.setzeLage('Die nächste Frage schickt mit: <b>' + esc(seite.schritt().split(':')[0]) + '</b>' + (kurz ? ', ' + esc(kurz) : ''));
  }

  // ---------- Den Kurslotsen finden ----------

  // Liefert der Kurslotse die Seite selbst aus, sucht die Seite ihn gleich. Von GitHub Pages aus sucht sie ihn erst,
  // wenn der Student einmal „Kurslotsen verbinden“ gewählt hat: Der Browser kann beim Zugriff auf den eigenen Rechner
  // nachfragen, und diese Rückfrage soll niemand sehen, der den Kurslotsen nicht eingerichtet hat.
  async function sucheLotsen(aufWunsch) {
    try {
      const d = await (await fetch(LOTSE + '/api/bereit', { signal: AbortSignal.timeout(4000) })).json();
      if (!d.schluessel) {
        el('lotseStatus').innerHTML = 'Der Kurslotse läuft, hat aber keinen Schlüssel für KI:connect. Der Schlüssel gehört in die Umgebungsvariable KICONNECT_KEY. ' +
          '<a href="' + (LOTSE || location.origin) + '/einrichtung.html" target="_blank" rel="noopener">Anleitung öffnen</a>';
        return;
      }
      el('lotseZeile').hidden = true;
      leiste.setzeZusatz('Modell: ' + d.modell);
      try { localStorage.setItem('lotseVerbunden', '1'); } catch (fehler) {}
      leiste.zeige(true, aufWunsch ? true : undefined);
    } catch (fehler) {
      if (aufWunsch) el('lotseStatus').textContent = 'Der Kurslotse antwortet auf diesem Rechner nicht. Er muss laufen, und der Browser muss den Zugriff auf den eigenen Rechner erlauben.';
    }
  }

  // ---------- Anmeldung der Seite ----------

  function einrichten(angaben) {
    seite = angaben;
    if (window.Leiste) { baue(); return; }
    const skript = document.createElement('script');
    skript.src = WURZEL + 'lotsenleiste.js';
    skript.onload = baue;
    document.head.appendChild(skript);
  }

  function baue() {
    leiste = Leiste.einrichten({
      name: 'Kurslotse',
      adresse: LOTSE + '/api/frage',
      wo: 'auf der HTML-Seite',
      sonst: 'Zu diesem Schritt',
      hinweis: '<p>Hier stellst du Fragen zu dem Schritt, der links zu sehen ist. Die Frage geht an das Sprachmodell von KI:connect. Die Seite schickt den Stand der Demo mit: Wer links etwas umstellt und danach fragt, bekommt eine Antwort zum neuen Stand. Wer links eine Stelle markiert, schickt sie mit der nächsten Frage mit. Wer oben den Schritt wechselt, beginnt einen neuen Dialog. „Kopieren“ legt den Dialog vorher in die Zwischenablage.</p>',
      platzhalter: 'Zum Beispiel: ' + (seite.beispiel || 'Was sehe ich hier?'),
      merken: 'lotseOffen',
      markieren: () => true,
      zeilen: seite.zeilen,
      link: link,
      mitschicken: mitschicken,
      marke: () => seite.schritt(),
      kopfzeile: () => {
        const kurz = seite.kurz ? seite.kurz() : '';
        return [(seite.thema || seite.titel) + ', ' + seite.schritt()].concat(kurz ? ['Stand der Demo: ' + kurz] : []);
      },
      beimLeeren: () => { standZuletzt = null; }
    });

    const stil = document.createElement('style');
    stil.textContent = GESTALTUNG;
    document.head.appendChild(stil);
    const zeile = document.createElement('div');
    zeile.className = 'lotse-zeile';
    zeile.id = 'lotseZeile';
    zeile.innerHTML = '<div class="lotse-status" id="lotseStatus" aria-live="polite"></div><button type="button" class="lotse-auf" id="lotseVerbinden">Kurslotsen verbinden</button>';
    document.body.appendChild(zeile);
    el('lotseVerbinden').addEventListener('click', () => { el('lotseStatus').textContent = ''; sucheLotsen(true); });

    zeigeStand();
    ladeBeschreibung();
    let verbunden = null;
    try { verbunden = localStorage.getItem('lotseVerbunden'); } catch (fehler) {}
    if (LOTSE === '' || verbunden === '1') sucheLotsen(false);
  }

  window.Kurslotse = {
    einrichten: einrichten,
    schrittGewechselt: () => { if (leiste) leiste.leere(); zeigeStand(); },
    standGeaendert: zeigeStand
  };
})();
