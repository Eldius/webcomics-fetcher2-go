
refresh:
	go run main.go plugin refresh

list:
	go run main.go plugin list

fetch-oots:
	go run main.go strips fetch oots

list-oots:
	go run main.go strips list oots

lint:
	golangci-lint run && revive -formatter friendly ./...
