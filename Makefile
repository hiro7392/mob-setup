# 変数定義
BINARY_NAME=myproject
GO_FILES=$(shell find . -type f -name '*.go')

# デフォルトターゲット
.PHONY: all
all: build

# ビルドターゲット
.PHONY: build
build:
	@echo "Building the binary..."
	go build -o $(BINARY_NAME)

# 実行ターゲット
.PHONY: run
run:
	@echo "Running the project..."
	go run .

# テストターゲット
.PHONY: test
test:
	@echo "Running tests..."
	go test ./...

# フォーマットターゲット
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

# リントターゲット (golangci-lint使用)
.PHONY: lint
lint:
	@echo "Linting code..."
	golangci-lint run

# クリーンターゲット
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

# モジュールの依存関係を整理
.PHONY: tidy
tidy:
	@echo "Tidying up dependencies..."
	go mod tidy
