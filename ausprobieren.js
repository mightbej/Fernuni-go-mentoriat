// Selbst ausprobieren: ein Kasten auf einer Lektion, in dem der Student ein Go-Programm ändert und ausführt.
//
// Ausführen kann nur ein Programm auf dem Rechner des Studenten. Das ist der Editorlotse. Thomas' Beschluss vom
// 2026-10-15: Die Lektion läuft dafür im Editorlotsen. Er liefert sie aus, sie sieht aus wie auf GitHub Pages, und der
// Kasten schickt das Programm an denselben Editorlotsen. Der hat Go immer und liefert auch den Editor CodeMirror unter
// /editor.js. Auf GitHub Pages liegt deshalb keine fremde Bibliothek, und keine Seite von GitHub Pages lässt den Rechner
// des Studenten ein Programm ausführen. Welche Lektionen der Editorlotse ausliefert, steht in daten/programme.json.
//
// Der Kasten hat zwei Zustände:
//
//   im Editorlotsen   Der Editor steht da, „Ausführen“ übersetzt und startet das Programm. Ein Programm mit fmt.Scan
//                     bekommt seine Eingabe in der Zeile unter der Ausgabe. Ein Link führt zurück zum Editor des
//                     Editorlotsen, im selben Tab.
//   anderswo          Auf GitHub Pages oder bei einem Lotsen ohne Go wählt der Student eine Fassung und sieht das
//                     Programm und was Go dazu gemeldet oder ausgegeben hat, festgehalten mit Go 1.26.1. Ein Satz sagt,
//                     wie der Student die Lektion im Editorlotsen öffnet: Editorlotse starten, im Menü „Programm wählen“
//                     das Kapitel und darin die Lektion wählen.
//                     Ein Link dorthin fehlt mit Absicht, Thomas' Beschluss vom 2026-10-15: Mit ihm hätte der Student
//                     zwei Tabs, die Seite des Editorlotsen und die Lektion. Der Reiter ist grau.
//
// Eine Lektion bindet die Datei ein und richtet den Kasten ein:
//
//   const probe = Ausprobieren.einrichten({
//     ziel: document.getElementById('werkstatt'),                 // in dieses Element kommt der Kasten
//     reiter: document.querySelector('.tab[data-stufe="5"]'),     // der Reiter, der grau wird, solange nichts läuft
//     fassungen: [
//       { name: 'Listing 3.1', code: '…', ergebnis: { ausgabe: 'Hello, World!\n' } },
//       { name: 'ohne import', code: '…', ergebnis: { meldungen: ['.\\main.go:4:2: undefined: fmt'] } }
//     ],
//     beiAenderung: () => Kurslotse.standGeaendert()                // optional
//   });
//
// probe.zeilen() liefert das Programm im Kasten, probe.stand() den Stand in Sätzen, beides für die Leiste.
(function () {
  'use strict';
  // Liefert ein Lotse die Seite selbst aus, fragt der Kasten ihn. Von GitHub Pages aus fragt er keinen.
  const LOKAL = location.protocol === 'http:' && (location.hostname === 'localhost' || location.hostname === '127.0.0.1');
  const ANLEITUNG = 'https://mightbej.github.io/Fernuni-go-mentoriat/lotsen/editorlotse.html';
  // Das Kapitel der Lektion, wie es im Menü des Editorlotsen steht: aus dem Ordner, zum Beispiel kap3-einfuehrung.
  const kapitel = () => { const k = location.pathname.match(/\/kap(\d+)-/); return k ? 'Kapitel ' + k[1] : 'Exkurse'; };
  const esc = t => String(t).replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');

  const GESTALTUNG = `
    .probe { margin:12px 0; }
    .probe .fassungen { display:flex;flex-wrap:wrap;gap:8px;margin:10px 0; }
    .probe .lage { border-left:4px solid #1e40af;background:#eff6ff;border-radius:8px;padding:10px 14px;margin:10px 0;font-size:.95rem; }
    .probe .lage p { margin:0 0 8px; } .probe .lage p:last-child { margin-bottom:0; }
    .probe .lage.fehler { border-left-color:#b91c1c;background:#fef2f2; }
    .probe .balken { height:10px;background:#e5e7eb;border-radius:5px;overflow:hidden;margin:8px 0; }
    .probe .balken span { display:block;height:100%;width:0;background:#0f766e;transition:width .3s; }
    .probe pre.quelltext { margin:10px 0;background:#1e293b;color:#e2e8f0;border-radius:8px;padding:10px 14px;font-family:Consolas,"Courier New",monospace;font-size:.95rem;line-height:1.6;overflow-x:auto; }
    .probe .editor { height:var(--probe-editor, 300px);border:1px solid #cbd5e1;border-radius:8px 8px 0 0;overflow:hidden;margin:10px 0 0; }
    .probe .griff { position:relative;height:12px;margin:0 0 10px;background:#e2e8f0;border:1px solid #cbd5e1;border-top:0;border-radius:0 0 8px 8px;cursor:row-resize;touch-action:none; }
    .probe .griff[hidden] { display:none; }
    .probe .griff::after { content:"";position:absolute;left:50%;top:4px;width:44px;height:4px;margin-left:-22px;border-radius:2px;background:#94a3b8; }
    .probe .griff:hover::after, .probe .griff:focus::after, body.probe-ziehen .probe .griff::after { background:#0f766e; }
    .probe .griff:focus { outline:none; }
    body.probe-ziehen { cursor:row-resize;user-select:none; }
    .probe .knoepfe { display:flex;flex-wrap:wrap;gap:8px;align-items:center;margin:10px 0; }
    .probe .knoepfe .stand { color:#4b5563;font-size:.9rem; }
    .probe .ausgabe { background:#111827;color:#f9fafb;border-radius:8px;padding:10px 14px;font-family:Consolas,"Courier New",monospace;white-space:pre-wrap;min-height:2.4em;max-height:320px;overflow-y:auto;overflow-wrap:anywhere; }
    .probe .ausgabe .meldung { color:#fca5a5; }
    .probe .ausgabe .ok { color:#86efac; }
    .probe .ausgabe .eingabe { color:#7dd3fc; }
    .probe .ausgabe .still { color:#9ca3af; }
    .probe form.eingabezeile { display:flex;gap:8px;margin:8px 0; }
    .probe form.eingabezeile[hidden] { display:none; }
    .probe form.eingabezeile input { flex:1;min-width:0;font:15px Consolas,"Courier New",monospace;padding:6px 8px;border-radius:6px;border:1px solid #cbd5e1; }
    .tab.ausgegraut { opacity:.55; }`;

  // Eine Anfrage an den Lotsen, der die Seite ausliefert. Mit daten geht sie als POST mit JSON.
  async function hole(pfad, daten) {
    const r = await fetch(pfad, daten !== undefined ? { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify(daten) } : { signal: AbortSignal.timeout(4000) });
    const d = await r.json().catch(() => null);
    if (!r.ok || d === null) throw new Error((d && d.fehler) || 'Fehler ' + r.status);
    return d;
  }
  let editorGeladen = null;   // das Laden von editor.js, einmal je Seite
  function ladeEditor() {
    if (window.Editor) return Promise.resolve();
    if (!editorGeladen) editorGeladen = new Promise((fertig, fehler) => {
      const s = document.createElement('script');
      s.src = '/editor.js';
      s.onload = fertig;
      s.onerror = () => { editorGeladen = null; fehler(new Error('Der Editor ließ sich nicht vom Editorlotsen laden.')); };
      document.head.appendChild(s);
    });
    return editorGeladen;
  }

  // Die Höhe des Editors: Der Student zieht den Griff unter dem Editor. Sie gilt für jeden Kasten auf jeder Lektion,
  // der Browser merkt sie sich. Ein Doppelklick oder Pos1 setzt sie auf 300 Pixel zurück.
  const HOEHE = 300, NIEDRIGSTE = 150, MERKEN_HOEHE = 'probeEditorHoehe';
  let hoehe = HOEHE;
  const begrenze = h => Math.round(Math.max(NIEDRIGSTE, Math.min(h, Math.max(NIEDRIGSTE, window.innerHeight * 0.85))));
  function setzeHoehe(h, merke) {
    hoehe = h;
    document.documentElement.style.setProperty('--probe-editor', begrenze(h) + 'px');
    if (merke) { try { localStorage.setItem(MERKEN_HOEHE, String(Math.round(h))); } catch (fehler) {} }
  }
  try { const h = Number(localStorage.getItem(MERKEN_HOEHE)); if (h > 0) setzeHoehe(h, false); } catch (fehler) {}
  window.addEventListener('resize', () => { if (hoehe !== HOEHE) setzeHoehe(hoehe, false); });
  function richteGriffEin(griff) {
    let start = null;   // Mausstelle und Höhe beim Beginn des Ziehens
    griff.addEventListener('pointerdown', ev => {
      if (ev.button !== 0) return;
      ev.preventDefault();
      start = { y: ev.clientY, h: begrenze(hoehe) };
      griff.setPointerCapture(ev.pointerId);
      document.body.classList.add('probe-ziehen');
    });
    griff.addEventListener('pointermove', ev => { if (start) setzeHoehe(begrenze(start.h + ev.clientY - start.y), false); });
    const ende = () => {
      if (!start) return;
      start = null;
      document.body.classList.remove('probe-ziehen');
      setzeHoehe(hoehe, true);
    };
    griff.addEventListener('pointerup', ende);
    griff.addEventListener('pointercancel', ende);
    griff.addEventListener('dblclick', () => setzeHoehe(HOEHE, true));
    // Mit der Tastatur: Pfeil nach unten macht den Editor höher, Pfeil nach oben niedriger.
    griff.addEventListener('keydown', ev => {
      const schritt = ev.key === 'ArrowDown' ? 20 : ev.key === 'ArrowUp' ? -20 : 0;
      if (schritt) { ev.preventDefault(); setzeHoehe(begrenze(begrenze(hoehe) + schritt), true); }
      if (ev.key === 'Home') { ev.preventDefault(); setzeHoehe(HOEHE, true); }
    });
  }

  function einrichten(angaben) {
    if (!document.getElementById('probeStil')) {
      const stil = document.createElement('style');
      stil.id = 'probeStil';
      stil.textContent = GESTALTUNG;
      document.head.appendChild(stil);
    }
    const fassungen = angaben.fassungen;
    const melde = angaben.beiAenderung || (() => {});
    const kasten = angaben.ziel;
    kasten.classList.add('probe');
    kasten.innerHTML = `
      <div class="fassungen" role="group" aria-label="Fassung des Programms"></div>
      <div class="lage" aria-live="polite"></div>
      <pre class="quelltext"></pre>
      <div class="editor" hidden></div>
      <div class="griff" hidden role="separator" aria-orientation="horizontal" aria-label="Höhe des Editors" tabindex="0" title="Ziehen ändert die Höhe des Editors. Doppelklick setzt sie zurück."></div>
      <div class="knoepfe">
        <button type="button" class="haupt ausfuehren" hidden>Ausführen</button>
        <button type="button" class="stopp" hidden>Stopp</button>
        <button type="button" class="zuruecksetzen" hidden>Fassung zurücksetzen</button>
        <span class="stand"></span>
      </div>
      <p><strong class="ueber">Was Go dazu meldet oder ausgibt</strong></p>
      <div class="ausgabe" aria-live="polite"></div>
      <form class="eingabezeile" hidden>
        <input autocomplete="off" spellcheck="false" aria-label="Eingabe für das Programm" placeholder="Wartet das Programm auf eine Eingabe? Tippe sie hier ein und drück die Eingabetaste.">
        <button type="submit">Eingeben</button>
      </form>`;
    const $ = s => kasten.querySelector(s);
    richteGriffEin($('.griff'));

    let zustand = 'ohne';     // 'ohne': nicht im Editorlotsen, 'bereit': im Editorlotsen mit Go
    let gewaehlt = 0;         // Nummer der gewählten Fassung
    let ed = null;            // der Editor, sobald er da ist
    let lauf = 0, quelle = null;
    let protokoll = '';       // was der letzte Lauf ausgegeben oder gemeldet hat, für die Leiste
    let gelaufen = false;     // ob der Student die Fassung im Kasten ausgeführt hat

    const code = () => ed ? ed.text() : fassungen[gewaehlt].code;

    // ---------- Fassungen ----------
    fassungen.forEach((f, i) => {
      const k = document.createElement('button');
      k.type = 'button';
      k.textContent = f.name;
      k.addEventListener('click', () => waehle(i));
      $('.fassungen').appendChild(k);
    });
    function waehle(i) {
      gewaehlt = i;
      kasten.querySelectorAll('.fassungen button').forEach((k, j) => { k.classList.toggle('gewaehlt', j === i); k.setAttribute('aria-pressed', String(j === i)); });
      stoppe();
      gelaufen = false;
      if (ed) {
        ed.setzeText(fassungen[i].code);
        ed.markiere([]);
        leereAusgabe();
        schreibe('Wähle „Ausführen“. Go übersetzt das Programm auf deinem Rechner und führt es aus.', 'still');
        protokoll = '';
      } else {
        $('.quelltext').textContent = fassungen[i].code;
        zeigeErgebnis(fassungen[i].ergebnis);
      }
      melde();
    }
    // Was Go 1.26.1 zu einer Fassung gemeldet oder ausgegeben hat, festgehalten auf der Seite.
    function zeigeErgebnis(e) {
      leereAusgabe();
      if (e.meldungen) {
        schreibe(e.meldungen.join('\n') + '\n', 'meldung');
        protokoll = 'Go meldet beim Übersetzen:\n' + e.meldungen.join('\n');
      } else {
        schreibe(e.ausgabe);
        protokoll = 'Ausgabe des Programms:\n' + e.ausgabe;
      }
    }

    // ---------- Ausgabe ----------
    function leereAusgabe() { $('.ausgabe').innerHTML = ''; }
    function schreibe(text, klasse) {
      const a = $('.ausgabe');
      const s = document.createElement('span');
      if (klasse) s.className = klasse;
      s.textContent = text;
      a.appendChild(s);
      a.scrollTop = a.scrollHeight;
    }

    // ---------- Zustand ----------
    function setzeLage(html, fehler) {
      $('.lage').innerHTML = html;
      $('.lage').classList.toggle('fehler', !!fehler);
      $('.lage').hidden = !html;
    }
    function zeigeZustand() {
      if (angaben.reiter) {
        angaben.reiter.classList.toggle('ausgegraut', zustand !== 'bereit');
        angaben.reiter.title = zustand === 'bereit' ? '' : 'Ausführen geht im Editorlotsen. Hier zeigt der Schritt, was Go meldet oder ausgibt.';
      }
      const bereit = zustand === 'bereit' && ed;
      $('.quelltext').hidden = !!bereit;
      $('.editor').hidden = !bereit;
      $('.griff').hidden = !bereit;
      $('.ausfuehren').hidden = !bereit;
      $('.zuruecksetzen').hidden = !bereit;
      $('.ueber').textContent = bereit ? 'Ausgabe' : 'Was Go dazu meldet oder ausgibt';
      if (zustand === 'ohne') {
        setzeLage('<p>Hier steht das Programm, und darunter, was Go dazu meldet oder ausgibt. Ändern und ausführen kannst du es im Editorlotsen auf deinem Rechner:</p>' +
          '<ol><li>Starte den Editorlotsen mit einem Doppelklick auf <code>editorlotse.exe</code>.</li>' +
          '<li>Wähle dort im Menü „Programm wählen“ den Eintrag „' + esc(kapitel()) + '“ und darin „' + esc(document.title) + '“.</li></ol>' +
          '<p>Den Editorlotsen und seine Anleitung findest du unter <a href="' + ANLEITUNG + '" target="_blank" rel="noopener">Editorlotse</a>.</p>');
      } else if (!ed) {
        setzeLage('<p>Der Editor wird geladen …</p>');
      } else {
        setzeLage('<p>Eigene Programme schreibst und speicherst du im <a href="/">Editor des Editorlotsen</a>.</p>');
      }
    }

    // Fragt den Lotsen, der die Seite ausliefert, ob er Programme ausführt. Das kann nur der Editorlotse, und nur er
    // meldet unter /api/bereit, ob er Go hat.
    async function pruefe() {
      let d;
      try { d = await hole('/api/bereit'); } catch (fehler) { return; }
      if (d.go === false) {
        setzeLage('<p>Der Editorlotse hat noch kein Go. Öffne seine <a href="/">Startseite</a>, dort richtet er Go ein. Lade danach diese Seite neu.</p>', true);
        return;
      }
      if (d.go !== true) return;   // ein Lotse, der nichts ausführt, zum Beispiel der Kurslotse
      zustand = 'bereit';
      zeigeZustand();
      try {
        await ladeEditor();
      } catch (fehler) {
        setzeLage('<p>' + esc(fehler.message) + '</p>', true);
        return;
      }
      ed = Editor.erzeuge($('.editor'), { text: fassungen[gewaehlt].code, beiAenderung: () => { gelaufen = false; melde(); } });
      leereAusgabe();
      schreibe('Wähle „Ausführen“. Go übersetzt das Programm auf deinem Rechner und führt es aus.', 'still');
      protokoll = '';
      zeigeZustand();
      melde();
    }

    // ---------- Ausführen ----------
    function stoppe() {
      if (quelle) { quelle.close(); quelle = null; }
      $('.stopp').hidden = true;
      $('.eingabezeile').hidden = true;
      if (zustand === 'bereit' && ed) $('.ausfuehren').hidden = false;
    }
    async function fuehreAus() {
      stoppe();
      leereAusgabe();
      ed.markiere([]);
      protokoll = '';
      gelaufen = true;
      $('.ausfuehren').disabled = true;
      $('.stand').textContent = 'Go übersetzt das Programm …';
      let d;
      try {
        d = await hole('/api/lauf', { code: ed.text() });
      } catch (fehler) {
        $('.ausfuehren').disabled = false;
        $('.stand').textContent = '';
        schreibe(fehler.message + '\n', 'meldung');
        protokoll = fehler.message;
        melde();
        return;
      }
      $('.ausfuehren').disabled = false;
      if (!d.lauf) {
        // Die Meldungen in der Form, in der Go sie im Terminal zeigt, wie auf der Seite festgehalten.
        const zeilen = (d.meldungen || []).map(m => '.\\main.go:' + m.zeile + (m.spalte ? ':' + m.spalte : '') + ': ' + m.text);
        ed.markiere(d.meldungen || []);
        schreibe((zeilen.length ? zeilen.join('\n') : d.text) + '\n', 'meldung');
        protokoll = 'Go meldet beim Übersetzen:\n' + (zeilen.length ? zeilen.join('\n') : d.text);
        $('.stand').textContent = (d.meldungen || []).some(m => m.zeile > 0) ? 'Nicht übersetzt. Die Stelle ist im Editor rot unterstrichen.' : 'Nicht übersetzt.';
        melde();
        return;
      }
      lauf = d.lauf;
      $('.stand').textContent = 'Das Programm läuft.' + (d.sekunden > 3 ? ' Das erste Übersetzen hat ' + d.sekunden.toFixed(0) + ' Sekunden gedauert, jedes weitere geht schneller.' : '');
      $('.ausfuehren').hidden = true;
      $('.stopp').hidden = false;
      $('.eingabezeile').hidden = false;
      protokoll = 'Ausgabe des Programms:\n';
      quelle = new EventSource('/api/lauf/ausgabe?lauf=' + lauf);
      quelle.addEventListener('aus', e => { const t = JSON.parse(e.data); protokoll += t; schreibe(t); });
      quelle.addEventListener('ende', e => {
        const ende = JSON.parse(e.data);
        stoppe();
        if (ende.grund) {
          schreibe('\n' + ende.grund + '\n', 'meldung');
          protokoll += '\n' + ende.grund;
          $('.stand').textContent = 'Beendet.';
        } else if (ende.code === 0) {
          $('.stand').textContent = 'Das Programm ist fertig.';
        } else {
          schreibe('\nDas Programm ist mit dem Exit-Code ' + ende.code + ' abgebrochen.\n', 'meldung');
          protokoll += '\nDas Programm ist mit dem Exit-Code ' + ende.code + ' abgebrochen.';
          $('.stand').textContent = 'Abgebrochen.';
        }
        melde();
      });
      quelle.onerror = () => {
        if (!quelle) return;
        stoppe();
        $('.stand').textContent = 'Die Verbindung zum Editorlotsen ist abgebrochen. Läuft das schwarze Fenster noch?';
      };
    }
    $('.ausfuehren').addEventListener('click', fuehreAus);
    $('.stopp').addEventListener('click', () => hole('/api/lauf/stopp', {}).catch(() => {}));
    $('.zuruecksetzen').addEventListener('click', () => waehle(gewaehlt));
    $('.eingabezeile').addEventListener('submit', async e => {
      e.preventDefault();
      const feld = $('.eingabezeile input');
      const zeile = feld.value;
      feld.value = '';
      schreibe(zeile + '\n', 'eingabe');
      protokoll += zeile + '\n';
      try { await hole('/api/lauf/eingabe', { lauf, zeile }); } catch (fehler) { schreibe(fehler.message + '\n', 'meldung'); }
    });

    // ---------- Start ----------
    waehle(0);
    zeigeZustand();
    if (LOKAL) pruefe();

    return {
      zeilen: () => code().split('\n'),
      stand: () => {
        const f = fassungen[gewaehlt];
        const saetze = ['Gewählt ist die Fassung „' + f.name + '“.'];
        if (zustand !== 'bereit') {
          saetze.push('Die Seite läuft nicht im Editorlotsen. Der Student kann hier nichts ändern und nichts ausführen. Die Seite zeigt das Programm und darunter, was Go 1.26.1 dazu gemeldet oder ausgegeben hat.');
        } else {
          saetze.push('Die Seite läuft im Editorlotsen. Der Student kann das Programm im Editor ändern und mit „Ausführen“ ausführen.');
          saetze.push(ed && ed.text() !== f.code ? 'Der Student hat das Programm im Editor geändert. Die Programmzeilen unten zeigen seine Fassung.' : 'Das Programm im Editor ist die Fassung, wie die Seite sie lädt.');
          if (!gelaufen) saetze.push('Diese Fassung ist seit der letzten Änderung nicht ausgeführt.');
        }
        if (protokoll) saetze.push(protokoll.slice(0, 3000));
        return saetze;
      },
      kurz: () => fassungen[gewaehlt].name + (zustand === 'bereit' && ed && ed.text() !== fassungen[gewaehlt].code ? ', geändert' : '')
    };
  }

  window.Ausprobieren = { einrichten: einrichten };
})();
