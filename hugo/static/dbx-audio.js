(function () {
  const fmt = (s) => {
    if (!Number.isFinite(s)) return '--:--';
    const m = Math.floor(s / 60);
    const r = Math.floor(s % 60);
    return String(m).padStart(2, '0') + ':' + String(r).padStart(2, '0');
  };

  const seek = (root, audio, bar, ev) => {
    const rect = bar.getBoundingClientRect();
    const x = Math.max(0, Math.min(1, (ev.clientX - rect.left) / rect.width));
    if (Number.isFinite(audio.duration)) audio.currentTime = x * audio.duration;
  };

  const wire = (root) => {
    const audio    = root.querySelector('[data-dbx-audio-el]');
    const btn      = root.querySelector('[data-dbx-audio-btn]');
    const glyph    = root.querySelector('[data-dbx-audio-glyph]');
    const cur      = root.querySelector('[data-dbx-audio-cur]');
    const dur      = root.querySelector('[data-dbx-audio-dur]');
    const bar      = root.querySelector('[data-dbx-audio-bar]');
    const fill     = root.querySelector('[data-dbx-audio-fill]');
    const buf      = root.querySelector('[data-dbx-audio-buf]');
    const playhead = root.querySelector('[data-dbx-audio-playhead]');

    if (!audio || !btn || !bar) { root.classList.add('is-bare'); return; }

    const setPct = (p) => {
      const v = Math.max(0, Math.min(1, p)) * 100 + '%';
      fill.style.width = v;
      playhead.style.left = v;
    };
    const setBuf = (p) => { buf.style.width = Math.max(0, Math.min(1, p)) * 100 + '%'; };

    const onMeta = () => { dur.textContent = fmt(audio.duration); };
    const onTime = () => {
      cur.textContent = fmt(audio.currentTime);
      if (Number.isFinite(audio.duration) && audio.duration > 0) {
        setPct(audio.currentTime / audio.duration);
      }
    };
    const onProgress = () => {
      if (audio.buffered.length && Number.isFinite(audio.duration) && audio.duration > 0) {
        setBuf(audio.buffered.end(audio.buffered.length - 1) / audio.duration);
      }
    };
    const setGlyph = (state) => {
      glyph.classList.toggle('dbx-audio__glyph--play',  state === 'play');
      glyph.classList.toggle('dbx-audio__glyph--pause', state === 'pause');
    };
    const onPlay  = () => { setGlyph('pause'); btn.setAttribute('aria-label', 'pause'); };
    const onPause = () => { setGlyph('play');  btn.setAttribute('aria-label', 'play');  };

    audio.addEventListener('loadedmetadata', onMeta);
    audio.addEventListener('durationchange', onMeta);
    audio.addEventListener('timeupdate', onTime);
    audio.addEventListener('progress', onProgress);
    audio.addEventListener('play', onPlay);
    audio.addEventListener('pause', onPause);
    audio.addEventListener('ended', onPause);
    audio.addEventListener('error', () => root.classList.add('is-bare'));

    btn.addEventListener('click', () => {
      if (audio.paused) audio.play().catch(() => root.classList.add('is-bare'));
      else audio.pause();
    });

    bar.addEventListener('click', (ev) => seek(root, audio, bar, ev));

    if (audio.readyState >= 1) onMeta();
  };

  const init = () => document.querySelectorAll('[data-dbx-audio]').forEach(wire);
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
  } else {
    init();
  }
})();
