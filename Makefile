# Makefile
.EXPORT_ALL_VARIABLES:	

EMAIL_HOST=
EMAiL_USERNAME=
EMAIL_PASSWORD=
EMAIL_PORT=

#GO111MODULE=on
#GOPROXY=direct
#GOSUMDB=off
GOPRIVATE=github.com/jeffotoni/gosendmail

build:
	@echo "########## Compilando nossa API ... "
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o gosendmail main.go
	@echo "buid completo..."
	@echo "\033[0;33m################ Enviando para o server #####################\033[0m"

update:
	@echo "########## Compilando nossa API ... "
	@rm -f go.*
	go mod init github.com/jeffotoni/gobrz/gosendmail
	go mod tidy
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o gosendmail main.go
	@echo "buid update completo..."
	@echo "fim"

tests:
	go test github.com/jeffotoni/gosendmail -v
	