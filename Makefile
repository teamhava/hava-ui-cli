
# Build locally, will make ./hava available to run
# Adds in some build flags to identify the binary in the future
build-local:
	go build -v \
	-ldflags="\
	-X 'github.com/teamhava/hava-ui-cli/version.Version=0.0.0' \
	-X 'github.com/teamhava/hava-ui-cli/version.Prerelease=alpha' \
	-X 'github.com/teamhava/hava-ui-cli/version.Build=local' \
	-X 'github.com/teamhava/hava-ui-cli/version.BuiltBy=$(shell whoami)' \
	-X 'github.com/teamhava/hava-ui-cli/version.Date=$(shell date)'"

# Updated go packages (will touch go.mod and go.sum)
update:
	go get -u
	go mod tidy

format:
	go fmt