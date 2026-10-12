(() => {
  const $ = (s, r = document) => r.querySelector(s);
  const $$ = (s, r = document) => [...r.querySelectorAll(s)];
  const esc = (s) => s.replace(/[&<>"]/g, (c) => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;' }[c]));
  const here = location.pathname;
  const root = document.documentElement;

  // course page: open the topic named in the URL hash, or the first one
  const groups = $$('details.tgroup');
  if (groups.length) {
    const target = location.hash && $(location.hash);
    ((target && $('details.tgroup', target)) || groups[0]).setAttribute('open', '');
  }

  /* ── sidebar ──────────────────────────────────────────────────────────── */
  const sidebar = $('.sidebar');
  if (sidebar) {
    const active = $$('.tree-link', sidebar).find((a) => a.getAttribute('href') === here);
    if (active) {
      active.classList.add('is-active');
      for (let d = active.closest('details'); d; d = d.parentElement.closest('details')) d.open = true;
    } else {
      $$('details', sidebar).filter((d) => here.startsWith(d.dataset.path)).forEach((d) => (d.open = true));
    }
    const saved = sessionStorage.getItem('lp:side');
    if (saved) sidebar.scrollTop = +saved;
    if (active) {
      const a = active.getBoundingClientRect(), s = sidebar.getBoundingClientRect();
      if (a.top < s.top + 40 || a.bottom > s.bottom - 40) active.scrollIntoView({ block: 'center' });
    }
    addEventListener('pagehide', () => sessionStorage.setItem('lp:side', sidebar.scrollTop));
    const menu = $('.menu-btn');
    const setNav = (open) => {
      document.body.classList.toggle('nav-open', open);
      menu?.setAttribute('aria-expanded', open);
      menu?.setAttribute('aria-label', open ? 'Close navigation' : 'Open navigation');
    };
    menu?.addEventListener('click', () => setNav(!document.body.classList.contains('nav-open')));
    $$('[data-nav-open]').forEach((b) => b.addEventListener('click', () => setNav(true)));
    $('.scrim')?.addEventListener('click', () => setNav(false));
    sidebar.addEventListener('click', (e) => { if (e.target.closest('a')) setNav(false); });
    addEventListener('keydown', (e) => { if (e.key === 'Escape') setNav(false); });
    matchMedia('(min-width: 901px)').addEventListener('change', (e) => { if (e.matches) setNav(false); });
  }

  /* ── header state, reading progress, toc scrollspy ───────────────────── */
  const onScroll = () => {
    const max = root.scrollHeight - innerHeight;
    root.style.setProperty('--read', max > 0 ? Math.min(100, (scrollY / max) * 100) : 0);
    root.toggleAttribute('data-scrolled', scrollY > 8);
  };
  addEventListener('scroll', onScroll, { passive: true });
  onScroll();

  // the same heading can be linked from the side TOC and from the context-bar TOC
  const tocLinks = new Map();
  $$('.toc a, .ctx-toc nav a').forEach((a) => {
    const id = decodeURIComponent(a.hash.slice(1));
    tocLinks.set(id, [...(tocLinks.get(id) || []), a]);
  });
  if (tocLinks.size) {
    const heads = $$('.prose h2[id], .prose h3[id]').filter((h) => tocLinks.has(h.id));
    const toc = $('.toc');
    let active;
    const spy = () => {
      const line = parseFloat(getComputedStyle(root).scrollPaddingTop) + 12;
      let cur = heads[0];
      for (const h of heads) if (h.getBoundingClientRect().top < line) cur = h;
      if (cur === active) return;
      active = cur;
      tocLinks.forEach((links) => links.forEach((a) => a.classList.remove('is-active')));
      if (!cur) return;
      tocLinks.get(cur.id).forEach((a) => a.classList.add('is-active'));
      // keep the active entry in view in a long side TOC
      const a = $('a.is-active', toc);
      if (a && toc.scrollHeight > toc.clientHeight) {
        const r = a.getBoundingClientRect(), t = toc.getBoundingClientRect();
        if (r.top < t.top + 60 || r.bottom > t.bottom - 60) toc.scrollTop += r.top - t.top - toc.clientHeight / 2;
      }
    };
    addEventListener('scroll', spy, { passive: true });
    spy();
  }
  const ctxToc = $('.ctx-toc');
  if (ctxToc) {
    ctxToc.addEventListener('click', (e) => { if (e.target.closest('a')) ctxToc.open = false; });
    addEventListener('click', (e) => { if (ctxToc.open && !ctxToc.contains(e.target)) ctxToc.open = false; });
    addEventListener('keydown', (e) => { if (e.key === 'Escape') ctxToc.open = false; });
  }

  /* ── code blocks: copy (the header is rendered at build time) ─────────── */
  $$('.code .copy-btn').forEach((btn) => {
    const label = $('span', btn);
    btn.addEventListener('click', async () => {
      try { await navigator.clipboard.writeText($('pre code', btn.closest('.code')).textContent.replace(/\n$/, '')); } catch { return; }
      label.textContent = 'Copied';
      btn.classList.add('is-copied');
      setTimeout(() => { label.textContent = 'Copy'; btn.classList.remove('is-copied'); }, 1400);
    });
  });

  /* ── diagrams ─────────────────────────────────────────────────────────────
     Mermaid sources are drawn in Excalidraw's style: mermaid's hand-drawn look
     (rough.js, the library Excalidraw itself draws with), Excalidraw's font and
     its colour palette. Lessons only tag nodes with a role (`class A,B compute`);
     the colours for each role live here so every diagram stays consistent. */
  const ROLES = {
    compute: ['#a5d8ff', '#1971c2'], memory: ['#b2f2bb', '#2f9e44'], queue: ['#ffec99', '#e67700'],
    io: ['#d0bfff', '#6741d9'], warn: ['#ffc9c9', '#e03131'], neutral: ['#e9ecef', '#495057'],
    ink: '#1e1e1e', paper: '#ffffff', group: '#f6f6f8',
  };
  const roleDefs = (p) => Object.keys(p).filter((k) => Array.isArray(p[k]))
    .map((k) => `  classDef ${k} fill:${p[k][0]},stroke:${p[k][1]},color:${p.ink},stroke-width:1.5px`).join('\n');
  const usesRoles = /^\s*(flowchart|graph|stateDiagram)/;
  // ":kubernetes:" in a node label becomes that tool's mark (static/logos, listed in LP_LOGOS)
  const LOGOS = new Set((window.LP_LOGOS || '').split(' '));
  const usedLogos = new Set();
  const withLogos = (src) => src.replace(/:([a-z0-9][a-z0-9-]*):/g, (m, n) => {
    if (!LOGOS.has(n)) return m;
    usedLogos.add(n);
    return `<i class='lg lg-${n}'></i>`;
  });

  const diagrams = $$('pre.mermaid');
  diagrams.forEach((d) => (d.dataset.src = d.textContent));
  let mermaid;
  const loadMermaid = () => new Promise((ok, fail) => {
    if (window.mermaid) return ok(window.mermaid);
    const el = document.createElement('script');
    el.src = window.LP_MERMAID;
    el.onload = () => ok(window.mermaid);
    el.onerror = fail;
    document.head.append(el);
  });
  async function renderDiagrams() {
    if (!diagrams.length) return;
    try {
      const hand = getComputedStyle(root).getPropertyValue('--font-hand').trim();
      [mermaid] = await Promise.all([mermaid ?? loadMermaid(), document.fonts.load('16px Excalifont').catch(() => {})]);
      const p = ROLES;
      mermaid.initialize({
        startOnLoad: false,
        securityLevel: 'loose',
        look: 'handDrawn',
        handDrawnSeed: 11, // fixed seed: the same sketch on every load
        theme: 'base',
        fontFamily: hand,
        themeVariables: {
          fontFamily: hand, fontSize: '16px',
          background: p.paper, mainBkg: p.neutral[0],
          primaryColor: p.neutral[0], primaryBorderColor: p.neutral[1], primaryTextColor: p.ink,
          secondaryColor: p.compute[0], tertiaryColor: p.group,
          lineColor: p.ink, textColor: p.ink, titleColor: p.ink,
          clusterBkg: p.group, clusterBorder: p.neutral[1], edgeLabelBackground: p.paper,
          noteBkgColor: p.queue[0], noteBorderColor: p.queue[1], noteTextColor: p.ink,
          actorBkg: p.compute[0], actorBorder: p.compute[1], actorTextColor: p.ink, actorLineColor: p.neutral[1],
          signalColor: p.ink, signalTextColor: p.ink, labelBoxBkgColor: p.io[0], labelBoxBorderColor: p.io[1], labelTextColor: p.ink,
          activationBkgColor: p.memory[0], activationBorderColor: p.memory[1],
          pie1: p.compute[0], pie2: p.memory[0], pie3: p.queue[0], pie4: p.io[0], pie5: p.warn[0], pie6: p.neutral[0],
          pieStrokeColor: p.ink, pieOuterStrokeColor: p.ink, pieTitleTextColor: p.ink, pieSectionTextColor: p.ink, pieLegendTextColor: p.ink,
        },
        flowchart: { curve: 'basis', padding: 16, nodeSpacing: 46, rankSpacing: 52, useMaxWidth: true },
        sequence: { useMaxWidth: true, mirrorActors: false },
      });
      diagrams.forEach((d) => {
        let src = d.dataset.src.replace(/^\s*classDef\s+(compute|memory|queue|io|warn|neutral)\b.*$/gm, '');
        if (usesRoles.test(src)) src = withLogos(src) + '\n' + roleDefs(p);
        d.removeAttribute('data-processed');
        d.textContent = src;
      });
      if (usedLogos.size) {
        const css = document.createElement('style');
        css.textContent = [...usedLogos].map((n) => `.lg-${n}{background-image:url("${window.LP_LOGO_DIR}${n}.svg")}`).join('');
        document.head.append(css);
      }
      await mermaid.run({ nodes: diagrams });
      // A very wide diagram scrolls sideways instead of shrinking its text to nothing.
      diagrams.forEach((d) => {
        const svg = $('svg', d);
        const w = svg?.viewBox.baseVal.width || 0;
        if (svg && w > d.parentElement.clientWidth * 1.3) svg.style.minWidth = `${Math.round(w * 0.78)}px`;
      });
    } catch (err) {
      console.error('mermaid failed', err);
    }
  }
  renderDiagrams.done = renderDiagrams();

  /* Expanded view: the diagram at full size. Drag pans it, scroll or pinch zooms it. */
  const zoom = $('dialog.zoom');
  if (zoom && diagrams.length) {
    const stage = $('.zoom-stage', zoom), scene = $('.zoom-scene', zoom);
    let view = { s: 1, x: 0, y: 0 }, size = null;
    const clamp = (v, lo, hi) => Math.min(hi, Math.max(lo, v));
    const apply = () => { scene.style.transform = `translate(${view.x}px, ${view.y}px) scale(${view.s})`; };
    // lay the drawing out at the size it will be seen, so text stays sharp
    const fitScene = () => {
      if (!size) return;
      const fit = Math.min((stage.clientWidth - 64) / size[0], (stage.clientHeight - 64) / size[1], 2.4);
      scene.style.width = `${size[0] * fit}px`;
      scene.style.height = `${size[1] * fit}px`;
    };
    const reset = () => { fitScene(); view = { s: 1, x: 0, y: 0 }; apply(); };
    new ResizeObserver(() => { if (zoom.open) reset(); }).observe(stage);
    const open = (svg) => {
      zoom.showModal();
      size = [svg.viewBox.baseVal.width, svg.viewBox.baseVal.height];
      const copy = svg.cloneNode(true);
      ['style', 'width', 'height'].forEach((a) => copy.removeAttribute(a));
      scene.replaceChildren(copy);
      reset();
    };
    $$('.diagram-zoom').forEach((btn) => btn.addEventListener('click', () => {
      const svg = $('svg', btn.parentElement);
      if (svg) open(svg);
    }));
    $('.zoom-reset', zoom).addEventListener('click', reset);

    // one pointer pans; two pointers pinch-zoom
    const pts = new Map();
    let pinch = 0;
    stage.addEventListener('pointerdown', (e) => {
      stage.setPointerCapture(e.pointerId);
      pts.set(e.pointerId, e);
      stage.classList.add('is-dragging');
      if (pts.size === 2) { const [a, b] = [...pts.values()]; pinch = Math.hypot(a.clientX - b.clientX, a.clientY - b.clientY); }
    });
    stage.addEventListener('pointermove', (e) => {
      const last = pts.get(e.pointerId);
      if (!last) return;
      pts.set(e.pointerId, e);
      if (pts.size === 2) {
        const [a, b] = [...pts.values()];
        const d = Math.hypot(a.clientX - b.clientX, a.clientY - b.clientY);
        if (pinch) view.s = clamp(view.s * d / pinch, 0.3, 6);
        pinch = d;
      } else {
        view.x += e.clientX - last.clientX;
        view.y += e.clientY - last.clientY;
      }
      apply();
    });
    const drop = (e) => { pts.delete(e.pointerId); pinch = 0; if (!pts.size) stage.classList.remove('is-dragging'); };
    stage.addEventListener('pointerup', drop);
    stage.addEventListener('pointercancel', drop);
    stage.addEventListener('wheel', (e) => {
      e.preventDefault();
      view.s = clamp(view.s * Math.exp(-e.deltaY * 0.0015), 0.3, 6);
      apply();
    }, { passive: false });
    zoom.addEventListener('keydown', (e) => {
      if (e.key === '+' || e.key === '=') view.s = clamp(view.s * 1.15, 0.3, 6);
      else if (e.key === '-') view.s = clamp(view.s / 1.15, 0.3, 6);
      else if (e.key === '0') view = { s: 1, x: 0, y: 0 };
      else return;
      e.preventDefault();
      apply();
    });
    $('.zoom-close', zoom).addEventListener('click', () => zoom.close());
    zoom.addEventListener('close', () => scene.replaceChildren());
    // /lesson/?expand=2 opens the page with its second diagram expanded, so a view can be linked
    const want = +new URLSearchParams(location.search).get('expand');
    if (want) renderDiagrams.done.then(() => { const svg = $('svg', diagrams[want - 1]); if (svg) open(svg); });
  }

  /* ── search: a command palette over pages and their headings ──────────── */
  const dlg = $('dialog.search');
  $$('[data-kbd]').forEach((k) => (k.textContent = /Mac|iPhone|iPad/.test(navigator.platform) ? '⌘K' : 'Ctrl K'));
  const input = $('input', dlg);
  const results = $('.search-results', dlg);
  const counter = $('[data-search-count]', dlg);
  const GROUPS = [['Path', 'Courses', '◆', 4], ['Topic', 'Topics', '▤', 4], ['Lesson', 'Lessons', '¶', 7], ['Section', 'Sections', '#', 5], ['Reference', 'Reference', '§', 3]];
  let index, sel = 0;

  async function openSearch() {
    if (dlg.open) return;
    dlg.showModal();
    input.select();
    index ??= await fetch(window.LP_INDEX).then((r) => r.json()).then((docs) => docs.flatMap((d) => [
      { ...d, _t: d.t.toLowerCase(), _m: `${d.c} ${d.l}`.toLowerCase(), _x: d.x.toLowerCase() },
      // every heading is its own result, so a concept can be found without knowing its page
      ...(d.h || []).map((h) => ({ u: `${d.u}#${h.i}`, t: h.t, c: `${d.c ? d.c + ' / ' : ''}${d.t}`, l: d.l, k: 'Section', x: '', _t: h.t.toLowerCase(), _m: d.t.toLowerCase(), _x: '' })),
    ]));
    run();
  }

  const row = (d, mark, snip) => `<li role="option"><a href="${d.u}"><span class="r-ico" aria-hidden="true">${GROUPS.find((g) => g[0] === d.k)?.[2] || '¶'}</span><span class="r-main"><span class="r-title">${mark(d.t)}</span><span class="r-path">${snip ? mark(snip) : esc(d.c)}</span></span><span class="r-tag">${esc(d.l)}</span></a></li>`;
  const select = (i, scroll) => {
    const items = $$('li[role=option]', results);
    if (!items.length) return;
    sel = (i + items.length) % items.length;
    items.forEach((li, n) => { li.classList.toggle('is-sel', n === sel); li.setAttribute('aria-selected', n === sel); });
    if (scroll) items[sel].scrollIntoView({ block: 'nearest' });
  };

  function run() {
    if (!index) return;
    const terms = input.value.toLowerCase().split(/\s+/).filter(Boolean);
    if (!terms.length) {
      // nothing typed yet: list the courses
      results.innerHTML = `<li class="group">Courses</li>${index.filter((d) => d.k === 'Path').map((d) => row(d, esc)).join('')}`;
      counter.textContent = '';
      select(0);
      return;
    }
    const hits = [];
    for (const d of index) {
      let score = 0, inTitle = true;
      for (const t of terms) {
        const title = d._t.includes(t);
        const s = (d._t.startsWith(t) ? 6 : 0) + (title ? 10 : 0) + (d._m.includes(t) ? 3 : 0) + (d._x.includes(t) ? 1 : 0);
        if (!s) { score = 0; break; }
        if (!title) inTitle = false;
        score += s;
      }
      if (score) hits.push([d.k === 'Section' ? score - 4 : score, d, inTitle]);
    }
    hits.sort((a, b) => b[0] - a[0] || a[1].t.length - b[1].t.length);
    const re = new RegExp(`(${terms.map((t) => t.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')).join('|')})`, 'gi');
    const mark = (s) => esc(s).replace(re, '<mark>$1</mark>');
    let html = '';
    // groups appear in order of their best hit
    const groups = GROUPS.map(([k, name, , max]) => [name, hits.filter((h) => h[1].k === k).slice(0, max)]).filter((g) => g[1].length)
      .sort((a, b) => b[1][0][0] - a[1][0][0]);
    for (const [name, list] of groups) {
      html += `<li class="group">${name}</li>` + list.map(([, d, inTitle]) => {
        const at = inTitle ? -1 : d._x.indexOf(terms[0]);
        return row(d, mark, at >= 0 ? `…${d.x.slice(Math.max(0, at - 30), at + 90)}` : '');
      }).join('');
    }
    results.innerHTML = html || `<li class="empty">No matches for “${esc(input.value)}”.</li>`;
    counter.textContent = hits.length ? `${hits.length} result${hits.length > 1 ? 's' : ''}` : '';
    results.scrollTop = 0;
    select(0);
  }

  input.addEventListener('input', run);
  input.addEventListener('keydown', (e) => {
    if (e.key === 'ArrowDown' || e.key === 'ArrowUp') {
      e.preventDefault();
      select(sel + (e.key === 'ArrowDown' ? 1 : -1), true);
    } else if (e.key === 'Enter') {
      const a = $('li.is-sel a', results);
      if (a) location.href = a.href;
    }
  });
  results.addEventListener('mousemove', (e) => {
    const li = e.target.closest('li[role=option]');
    if (li && !li.classList.contains('is-sel')) select($$('li[role=option]', results).indexOf(li));
  });
  dlg.addEventListener('click', (e) => { if (e.target === dlg) dlg.close(); });
  $('.search-btn')?.addEventListener('click', openSearch);
  // /?q=kv+cache opens search with the query filled in, so searches can be linked
  const q = new URLSearchParams(location.search).get('q');
  if (q) { input.value = q; openSearch(); }
  addEventListener('keydown', (e) => {
    const typing = /^(INPUT|TEXTAREA|SELECT)$/.test(document.activeElement?.tagName);
    if ((e.key === '/' && !typing) || (e.key.toLowerCase() === 'k' && (e.metaKey || e.ctrlKey))) {
      e.preventDefault();
      if (dlg.open && e.key !== '/') dlg.close(); else openSearch();
    }
  });
})();
