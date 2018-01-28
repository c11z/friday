# Run friday cli with local environment variables
run: cmd/friday.go
	env $$(cat local.env | xargs) go run $<
