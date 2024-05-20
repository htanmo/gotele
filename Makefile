default: main.go
	@go build -o bot 

clean: bot
	@rm bot
