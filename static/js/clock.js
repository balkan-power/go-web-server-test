function updateClock() {
    // setting it inside home.go only sets it once, therefore
    // not a live clock. no choice but to use j*v*scr*pt
    const now = new Date();
    const pad = n => String(n).padStart(2, "0");

    // ai slop
    document.getElementById("clock").textContent =
    "Today is: " +
    `${pad(now.getDate())}.` +
    `${pad(now.getMonth() + 1)}.` +
    `${now.getFullYear()} ` +
    `${pad(now.getHours())}:` +
    `${pad(now.getMinutes())}:` +
    `${pad(now.getSeconds())}`;
}

updateClock();
setInterval(updateClock, 1000); // update once every second (i hate js)