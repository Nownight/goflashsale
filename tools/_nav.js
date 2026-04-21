// Shared top-nav renderer for xtp tool pages.
// Usage: <div id="nav"></div>  + <script src="_nav.js" data-active="shop"></script>
(function() {
  const scriptEl = document.currentScript;
  const active = scriptEl?.dataset.active || "";
  const links = [
    { key: "portal",    label: "工作台",      href: "portal.html" },
    { key: "shop",      label: "商店",        href: "shop.html" },
    { key: "orders",    label: "订单",        href: "orders.html" },
    { key: "workspace", label: "生产力引擎",  href: "workspace.html" },
  ];

  function loggedInName() {
    try { return (JSON.parse(localStorage.getItem("garden.user")) || {}).name || null; }
    catch { return null; }
  }

  function render() {
    const mount = document.getElementById("nav");
    if (!mount) return;
    const user = loggedInName();
    mount.innerHTML = `
      <nav class="glass-nav">
        <div class="glass-nav-left">
          <a href="portal.html" class="glass-nav-brand">
            <span class="glass-nav-brand-mark">⌁</span>
            <span>xtp Workspace</span>
          </a>
          ${links.map(l => `
            <a href="${l.href}" class="nav-link ${l.key === active ? "active" : ""}">${l.label}</a>
          `).join("")}
        </div>
        <div class="glass-nav-right">
          ${user ? `<span style="color:var(--text-mute); font-size:12px;">Signed in · <b style="color:var(--text);">${user}</b></span>` : ""}
          <a class="back-to-garden" href="../小菜园.html">← 返回菜园</a>
        </div>
      </nav>
    `;
  }
  render();
  window.addEventListener("storage", render);
})();
