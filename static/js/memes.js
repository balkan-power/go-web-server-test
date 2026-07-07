const memes = [
    "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
    "https://www.youtube.com/watch?v=TNUXzTP79e0",
    "https://www.youtube.com/watch?v=3c2wdlxLy7Q",
    "https://www.youtube.com/watch?v=XcyzLyZeqf4?t=43",
    "https://www.youtube.com/watch?v=j308rH1j_bg",
    "https://www.youtube.com/watch?v=vbIbIurceNU",
    "https://www.youtube.com/watch?v=EAzskL3gNxc",
    "https://www.youtube.com/watch?v=dQa8lydtFmE",
    "https://www.youtube.com/watch?v=s0Gpd2ooB9w",
    "https://www.youtube.com/watch?v=wa5inGuht_o",
    "https://www.youtube.com/watch?v=bPzVV_5sQtc",
    "https://www.youtube.com/watch?v=sUSN7fqVBio",
];

function randomMeme(e) {
    e.preventDefault();
    const url = memes[Math.floor(Math.random() * memes.length)];
    window.location.href = url;
}