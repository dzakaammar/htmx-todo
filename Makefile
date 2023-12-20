generate:
	@templ generate
	@npx tailwindcss -i ./static/input.css -o ./static/dist/main.css
run:
	@air
deps:
	@npm install -D tailwindcss