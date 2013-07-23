build:
	@go install aldaviva.com/mailsender

run: build
	@bin/mailsender "Spore Reloaded"