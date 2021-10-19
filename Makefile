build-deploy-clean: build deploy clean
clean:
	rm -Rf server
build:
	GOOS=linux GOARCH=arm GOARM=5 go build -o server ./server.go
deploy:
	ssh pi@nas.local 'mkdir -p transmission-torrent-handler/internal/templates'
	ssh pi@nas.local 'mkdir -p transmission-torrent-handler/certs'
	scp internal/templates/* pi@nas.local:/home/pi/transmission-torrent-handler/internal/templates
	scp ./certs/* pi@nas.local:/home/pi/transmission-torrent-handler/certs
	scp ./transmission-torrent-handler.service pi@nas.local:/home/pi/transmission-torrent-handler/transmission-torrent-handler.service
	ssh pi@nas.local 'sudo systemctl stop transmission-torrent-handler.service'
	scp ./server pi@nas.local:/home/pi/transmission-torrent-handler/server
	ssh pi@nas.local 'sudo mv /home/pi/transmission-torrent-handler/transmission-torrent-handler.service /etc/systemd/system/transmission-torrent-handler.service'
	ssh pi@nas.local 'sudo systemctl start transmission-torrent-handler.service'
