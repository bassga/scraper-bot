# 出力バイナリ名
BINARY_NAME=scraper-bot

# デフォルトはbuild
all: build

# buildターゲット
build:
	go build -o $(BINARY_NAME) ./cmd/scraper

# runターゲット
run: build
	./$(BINARY_NAME)

# cleanターゲット
clean:
	rm -f $(BINARY_NAME)