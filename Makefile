PROJECT_NAME=gpa

build:
	go build -o ./build/${PROJECT_NAME}

clean:
	rm -rf build/