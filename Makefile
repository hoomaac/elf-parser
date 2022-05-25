


static:
	@echo "Build statically"
	@go build -ldflags="-linkmode external -extldflags -static"

dynamic:
	@echo "Build dynamically"
	@go build -gcflags

static_memcheck:
	@echo "Build statically with memcheck"
	@go build -ldflags="-linkmode external -extldflags -static" -gcflags '-m -l'

dynamic_memcheck:
	@echo "Build dynamically with memcheck"
	@go build -gcflags '-m -l'

test:
	@go test -v

benchmem:
	@echo "Memory Bemchmark"
	@go test -bench . -benchmem
