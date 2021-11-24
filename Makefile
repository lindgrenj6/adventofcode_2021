all: test

test:
	@if [ -z "$(day)" ]; then \
		go test ./...; \
	else \
		go test ./$(day); \
	fi

newday:
	@echo "Creating folder for day $(day)"
	@cp -R template/ day$(day)
