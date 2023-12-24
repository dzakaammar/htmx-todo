generate:
	@templ generate
	@npx tailwindcss -i ./static/input.css -o ./static/dist/main.css
run:
	@air
deps:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/cosmtrek/air@latest
	@npm install -D tailwindcss
slide:
	@marp --html slides/slides.md