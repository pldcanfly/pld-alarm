(function() {
	const init = () => {
		const player = document.getElementById("player");
		if (player) {
			const play = player.querySelector<HTMLElement>("#play")
			play?.addEventListener('click', () => {
				console.log("Play now!!");;
			});
		}
	};


	init();
})();
