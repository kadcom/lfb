.PHONY: all clean

all: simple simple-linux simple-linux.zip

static/styles.css: input.css 
	@echo "Compiling CSS"
	@yarn run build 

simple-linux: static/styles.css
	@echo "Building simple-linux"
	@GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o simple-linux

simple: static/styles.css
	@echo "Building native image"
	@go build -ldflags="-w -s" -o simple

simple-linux.zip: simple-linux static/styles.css
	@echo "Creating zip file"
	@zip simple-linux.zip simple-linux static/styles.css

clean:
	@echo "Cleaning up"
	@rm -f simple simple-linux static/styles.css simple-linux.zip

