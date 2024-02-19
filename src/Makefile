TARGET = gitman
SRC = gitman.go

build: $(TARGET)

$(TARGET): $(SRC)
	go build -o $(TARGET) $(SRC)

install: $(TARGET)
	sudo cp $(TARGET) /usr/local/bin/$(TARGET)

clean:
	rm -f $(TARGET)

.PHONY: build install clean
