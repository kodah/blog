build:
	go build -o dist/blogctl .

watch:
	air -c '.air-main.toml'