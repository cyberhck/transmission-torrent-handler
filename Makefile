build-deploy-clean: build deploy clean
clean:
	rm -Rf server
build:
	GOOS=linux GOARCH=arm GOARM=5 go build -o server ./server.go
deploy:
	scp ./server pi@nas.local:/home/pi/server
	scp internal/templates/* pi@nas.local:/home/pi/internal/templates
