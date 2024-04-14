server:
	go run main.go
start-dm:
	CompileDaemon -command="./go-mongo-tut"
reown:
	sudo chown -R $(USER) .
chmod:
	sudo chmod -R 777 .

.PHONY: server reown start-dm
