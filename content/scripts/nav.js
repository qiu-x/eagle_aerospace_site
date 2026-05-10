(function () {
  /** @type {HTMLInputElement | null} */
  const toggle = document.querySelector('.nav-toggle');
  /** @type {HTMLElement | null} */
  const panel  = document.getElementById('navPanel');
  if (!toggle || !panel) return;

  const labelOpen  = 'Otwórz menu';
  const labelClose = 'Zamknij menu';

  /** @param {boolean | undefined} open */
  function setOpen(open) {
    if (!panel || !toggle) return;
    panel.classList.toggle('is-open', open);
    toggle.setAttribute('aria-expanded', String(open));
    toggle.setAttribute('aria-label', open ? labelClose : labelOpen);
    document.body.classList.toggle('nav-open', open);
    if (open) {
      const first = panel.querySelector('a');
      if (first) first.focus();
    } else {
      toggle.focus();
    }
  }

  toggle.addEventListener('click', () => {
    setOpen(!panel.classList.contains('is-open'));
  });

  panel.addEventListener('click', e => {
    if (e.target instanceof HTMLAnchorElement) setOpen(false);
  });

  document.addEventListener('keydown', e => {
    if (e.key === 'Escape' && panel.classList.contains('is-open')) {
      setOpen(false);
    }
  });
})();
