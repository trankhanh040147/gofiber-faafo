server:
	go run cmd/main.go
start-dm:
	CompileDaemon -command="./go-mongo-tut"
reown:
	sudo chown -R $(USER) .
chmod:
	sudo chmod -R 777 .
start-dB:
	sudo docker start mymongo

.PHONY: server reown start-dm start-db
