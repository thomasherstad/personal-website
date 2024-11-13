function toggleMenu() {
    console.log("toggling menu")
    const menu = document.querySelector(".menu-links");
    const icon = document.querySelector(".hamburger-icon");
    menu.classList.toggle("open");
    icon.classList.toggle("open");
}

console.log("script.js has loaded")