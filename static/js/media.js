"use strict";
(function () {
    const init = () => {
        const player = document.getElementById("player");
        if (player) {
            const play = player.querySelector("#play");
            play === null || play === void 0 ? void 0 : play.addEventListener('click', () => {
                console.log("Play now!!");
                ;
            });
        }
    };
    init();
})();
