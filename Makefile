run: build
	if pgrep "pld-alarm"; then pkill "pld-alarm"; fi
	cd build && ./pld-alarm

build:
	tailwindcss -i ./styles/style.css -o ./build/static/style.css
	templ generate
	go build -o build/pld-alarm main.go
