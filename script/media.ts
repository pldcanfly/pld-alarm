(function() {

	const init = () => {
		const ws = new WebSocket("ws://localhost:8080/ws");

		ws.onopen = () => {
			console.log("connected");
		}
		ws.onmessage = (e) => {
			console.log("WS:", e.data);
		}

		const player = document.getElementById("controls");
		const play = player?.querySelector<HTMLElement>("#play")
		const audio = player?.querySelector<HTMLAudioElement>("audio");
		const progress = player?.querySelector<HTMLDivElement>("#progress_bar")
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
				}))
				audio.addEventListener("timeupdate", () => {
					const percent = Math.round(audio.currentTime / audio.duration * 100);
					progress.style.width = percent + "%";
					ws.send(JSON.stringify({
						sender: "controls",
						action: "update",
						data: {
							currentTime: audio.currentTime,
							duration: audio.duration
						},
					}))
				});
			});
		}
	};


	init();
})();
