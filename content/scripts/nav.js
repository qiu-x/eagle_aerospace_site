(function () {
  const toggle = document.querySelector('.nav-toggle');
  const panel  = document.getElementById('navPanel');
  if (!toggle || !panel) return;

  function setOpen(open) {
    panel.classList.toggle('is-open', open);
    toggle.setAttribute('aria-expanded', String(open));
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
