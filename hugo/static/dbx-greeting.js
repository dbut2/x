(function () {
  const root = document.querySelector('[data-dbx-greeting]');
  if (!root) return;
  if (matchMedia('(prefers-reduced-motion: reduce)').matches) return;

  let lines;
  try {
    lines = JSON.parse(root.getAttribute('data-dbx-greeting'));
  } catch (e) {
    return;
  }
  if (!Array.isArray(lines) || lines.length < 2) return;

  const num = (name, fallback) => {
    const v = Number(root.getAttribute(name));
    return Number.isFinite(v) && v > 0 ? v : fallback;
  };
  const interval = num('data-dbx-greeting-interval', 3400);
  const fade = num('data-dbx-greeting-fade', 280);

  let i = 0;
  let held = false;
  let swap;

  const advance = () => {
    root.style.opacity = '0';
    clearTimeout(swap);
    swap = setTimeout(() => {
      i = (i + 1) % lines.length;
      root.textContent = lines[i].text;
      root.setAttribute('lang', lines[i].lang);
      root.style.opacity = '1';
    }, fade);
  };

  root.addEventListener('mouseenter', () => { held = true; });
  root.addEventListener('mouseleave', () => { held = false; });
  setInterval(() => { if (!held && !document.hidden) advance(); }, interval);
})();
