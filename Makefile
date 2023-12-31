prepare: 
	rm -rf ./build/*
	mkdir -p ./build/static

build: prepare
	tailwindcss -c ./tailwind.config.js -m -i ./styles/style.css -o ./static/style.css
	cp -r ./static/ ./build/
	tsc
	templ generate
	go build -o build/pld-alarm main.go

run: build
	cd build && ./pld-alarm
