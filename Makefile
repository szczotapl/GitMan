TARGET = gitman
TARGET_UPDATE = gitman_update
SRC = src/gitman.go
SRC_UPDATE = src-update/gitman_update.go

build: $(TARGET)

$(TARGET): $(SRC)
	go build -o $(TARGET) $(SRC)
	go build -o $(TARGET_UPDATE) $(SRC_UPDATE)

install: $(TARGET)
	sudo cp $(TARGET) /usr/local/bin/$(TARGET)
	sudo cp $(TARGET_UPDATE) /usr/local/bin/$(TARGET_UPDATE)

clean:
	rm -f $(TARGET)
	rm -f $(TARGET_UPDATE)

.PHONY: build install clean
