build:
	@go build -o chinaup-to-date

run: build
	@ chinaup-to-date

test: 
	@go test -v ./...