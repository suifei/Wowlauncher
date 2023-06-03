build:
	go build . -o bin/Wowlauncher
	
all:
	go install github.com/mitchellh/gox@latest
	rm -rf ./bin
	mkdir bin && cd bin
	gox ../
	cd ..

push:
	git add .
	git commit -m "update"
	git push

pull:
	git pull