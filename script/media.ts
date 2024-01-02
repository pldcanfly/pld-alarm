(function() {

	const init = () => {
		const player = document.getElementById("player");
		const play = player?.querySelector<HTMLElement>("#play")
		const audio = player?.querySelector<HTMLAudioElement>("audio");
		const progress = player?.querySelector<HTMLDivElement>("#progress_bar")
		if (play && audio && progress) {
			setTimeout(() => {
				audio.play();
			}, 2000)
			audio.addEventListener("timeupdate", () => {
				const percent = Math.round(audio.currentTime / audio.duration * 100);
				progress.style.width = percent + "%";
			});
			console.log("Play now!!");
		}
	};


	init();
})();
