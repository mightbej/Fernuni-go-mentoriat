// Teilt eine Erklärseite in Bereiche. Jeder Abschnitt nennt seinen Bereich im Attribut data-bereich.
// Die Leiste mit den Knöpfen entsteht hier. Ohne JavaScript bleiben alle Abschnitte sichtbar.
(function () {
    const NAMEN = {
        grundlagen: "Grundlagen",
        ausprobieren: "Ausprobieren",
        vertiefung: "Vertiefung",
        faustregeln: "Faustregeln"
    };
    const ALLES = "alles";

    function einrichten() {
        const abschnitte = Array.from(document.querySelectorAll("[data-bereich]"));
        const bereiche = [];
        abschnitte.forEach(a => { if (!bereiche.includes(a.dataset.bereich)) bereiche.push(a.dataset.bereich); });
        if (bereiche.length < 2) return;

        const leiste = document.createElement("nav");
        leiste.className = "bereiche";
        leiste.setAttribute("aria-label", "Bereiche dieser Seite");
        const knoepfe = {};
        bereiche.concat(ALLES).forEach(b => {
            const k = document.createElement("button");
            k.type = "button";
            k.className = "bereich-knopf";
            k.textContent = b === ALLES ? "Alles zeigen" : (NAMEN[b] || b);
            k.addEventListener("click", () => zeige(b, true));
            knoepfe[b] = k;
            leiste.appendChild(k);
        });
        abschnitte[0].parentNode.insertBefore(leiste, abschnitte[0]);

        function zeige(bereich, merken) {
            if (!knoepfe[bereich]) bereich = bereiche[0];
            abschnitte.forEach(a => { a.hidden = bereich !== ALLES && a.dataset.bereich !== bereich; });
            Object.keys(knoepfe).forEach(b => knoepfe[b].setAttribute("aria-pressed", b === bereich ? "true" : "false"));
            if (merken) {
                history.replaceState(null, "", "#" + bereich);
                leiste.scrollIntoView({ block: "start" });
            }
        }

        // Die Adresse mit #ausprobieren öffnet die Seite gleich im Bereich "Ausprobieren".
        zeige(location.hash.slice(1), false);
        window.addEventListener("hashchange", () => zeige(location.hash.slice(1), false));
    }

    if (document.readyState === "loading") {
        document.addEventListener("DOMContentLoaded", einrichten);
    } else {
        einrichten();
    }
})();
