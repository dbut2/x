(function () {
  let lb, lbImg, lbCap, lbIdx, lbPrev, lbNext, lbClose;
  let state = { photos: [], i: 0 };

  const ensureLightbox = () => {
    if (lb) return;
    lb = document.createElement('div');
    lb.className = 'dbx-lightbox';
    lb.innerHTML =
      '<div class="dbx-lightbox__frame" data-lb-frame>' +
      '  <div class="dbx-lightbox__img" data-lb-img></div>' +
      '  <div class="dbx-lightbox__meta">' +
      '    <span data-lb-cap>untitled</span>' +
      '    <span class="dbx-lightbox__nav">' +
      '      <button type="button" data-lb-prev aria-label="previous">← prev</button>' +
      '      <span class="dbx-lightbox__idx" data-lb-idx>1 / 1</span>' +
      '      <button type="button" data-lb-next aria-label="next">next →</button>' +
      '      <button type="button" data-lb-close aria-label="close">close ✕</button>' +
      '    </span>' +
      '  </div>' +
      '</div>';
    document.body.appendChild(lb);
    lbImg   = lb.querySelector('[data-lb-img]');
    lbCap   = lb.querySelector('[data-lb-cap]');
    lbIdx   = lb.querySelector('[data-lb-idx]');
    lbPrev  = lb.querySelector('[data-lb-prev]');
    lbNext  = lb.querySelector('[data-lb-next]');
    lbClose = lb.querySelector('[data-lb-close]');

    const frame = lb.querySelector('[data-lb-frame]');
    lb.addEventListener('click', close);
    frame.addEventListener('click', (e) => e.stopPropagation());
    lbPrev.addEventListener('click', () => show(state.i - 1));
    lbNext.addEventListener('click', () => show(state.i + 1));
    lbClose.addEventListener('click', close);
    document.addEventListener('keydown', (e) => {
      if (!lb.classList.contains('is-open')) return;
      if (e.key === 'Escape')      close();
      if (e.key === 'ArrowLeft')   show(state.i - 1);
      if (e.key === 'ArrowRight')  show(state.i + 1);
    });
  };

  const show = (i) => {
    const n = state.photos.length;
    if (!n) return;
    state.i = ((i % n) + n) % n;
    const p = state.photos[state.i];
    lbImg.innerHTML = '';
    const img = document.createElement('img');
    img.src = p.src;
    img.alt = p.alt || '';
    lbImg.appendChild(img);
    lbCap.textContent = p.cap || p.alt || 'untitled';
    lbIdx.textContent = (state.i + 1) + ' / ' + n;
  };

  const open = (gallery, startIdx) => {
    ensureLightbox();
    const figs = gallery.querySelectorAll('figure');
    state.photos = Array.from(figs).map((f) => {
      const im = f.querySelector('img');
      const cap = f.querySelector('figcaption');
      return {
        src: im ? im.currentSrc || im.src : '',
        alt: im ? im.alt : '',
        cap: cap ? cap.textContent.trim() : '',
      };
    });
    show(startIdx);
    lb.classList.add('is-open');
  };

  const close = () => { if (lb) lb.classList.remove('is-open'); };

  const init = () => {
    document.querySelectorAll('[data-dbx-gallery]').forEach((g) => {
      const figs = g.querySelectorAll('figure');
      figs.forEach((f, i) => {
        f.addEventListener('click', () => open(g, i));
      });
    });
  };

  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
  } else {
    init();
  }
})();
