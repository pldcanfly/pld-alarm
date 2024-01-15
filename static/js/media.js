"use strict";
(function () {
    const init = () => {
        const ws = new WebSocket("ws://localhost:8080/ws/media");
        ws.onopen = () => {
            console.log("connected");
        };
        ws.onmessage = handleMessage;
        const player = document.getElementById("controls");
        const play = player === null || player === void 0 ? void 0 : player.querySelector("#play");
        const audio = player === null || player === void 0 ? void 0 : player.querySelector("audio");
        const progress = player === null || player === void 0 ? void 0 : player.querySelector("#progress_bar");
        if (play && audio && progress) {
            play.addEventListener("click", () => {
                console.log("Play now!!");
                audio.play();
                ws.send(JSON.stringify({
                    sender: "controls",
                    action: "playing",
                    data: {
                        currentTime: audio.currentTime,
                        duration: audio.duration
                    },
                }));
            });
            audio.addEventListener("timeupdate", () => {
                const percent = Math.round(audio.currentTime / audio.duration * 100);
                progress.style.width = percent + "%";
                ws.send(JSON.stringify({
                    sender: "controls",
                    action: "timeupdate",
                    data: {
                        currentTime: audio.currentTime,
                        duration: audio.duration
                    },
                }));
            });
        }
    };
    const handleMessage = (e) => {
        console.log(e.data);
    };
    init();
})();
