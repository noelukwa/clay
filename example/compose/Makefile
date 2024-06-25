# Makefile

# Specify the output file name
OUTFILE = work.wasm

# Specify the source file
SRC = main.go

# Build command
build:
	@GOOS=js GOARCH=wasm go build -o $(OUTFILE) $(SRC)
	
# Clean command to remove the generated .wasm file
clean:
	@rm -f $(OUTFILE)

.PHONY: build clean
