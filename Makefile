build:
	tailwindcss -c ./tailwind.config.js -m -i ./styles/style.css -o ./build/static/style.css
	templ generate
	go build -o build/pld-alarm main.go
