"use strict";
(function () {
    const init = () => {
        const player = document.getElementById("player");
        const play = player === null || player === void 0 ? void 0 : player.querySelector("#play");
        const audio = player === null || player === void 0 ? void 0 : player.querySelector("audio");
        const progress = player === null || player === void 0 ? void 0 : player.querySelector("#progress_bar");
        if (play && audio && progress) {
            setTimeout(() => {
                audio.play();
            }, 2000);
            audio.addEventListener("timeupdate", () => {
                const percent = Math.round(audio.currentTime / audio.duration * 100);
                progress.style.width = percent + "%";
            });
            console.log("Play now!!");
        }
    };
    init();
})();
