TARGET = gitman
TARGET_UPDATE = gitman_update
SRC = src/gitman.go
SRC_UPDATE = src-update/gitman_update.go

build: $(TARGET)

$(TARGET): $(SRC)
	go build -o $(TARGET) $(SRC)
	go build -o $(TARGET_UPDATE) $(SRC_UPDATE)

install: build
	sudo cp $(TARGET) /usr/local/bin/$(TARGET)
	@if [ -f "$(TARGET_UPDATE)" ]; then sudo cp $(TARGET_UPDATE) /usr/local/bin/$(TARGET_UPDATE); fi

clean:
	rm -f $(TARGET) $(TARGET_UPDATE)

.PHONY: build install clean
