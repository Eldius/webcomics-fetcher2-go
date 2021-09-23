.EXPORT_ALL_VARIABLES:
WEBCOMICS_FETCHER_PLUGIN_FOLDER=../webcomics-fetcher2-plugins/plugins
WEBCOMICS_FETCHER_STRIPS_FOLDER=../webcomics-fetcher2-plugins/strips


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
