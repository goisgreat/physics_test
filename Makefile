compile:
	@echo "Installing $(notdir $(shell pwd))"
	go install .
	@echo "Done! Errors will be shown above."
execute:
	@echo ~/go/bin/$(notdir $(shell pwd))
	@~/go/bin/$(notdir $(shell pwd))
clean:
	rm ~/go/bin/$(notdir $(shell pwd))
.SILENT:
run: compile execute