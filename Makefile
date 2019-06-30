GO=go
GOBUILD=$(GO) build
GOGET=$(GO) get -u -v
GOTEST=$(GO) test -v
BINNAME=commupace
GOARCH=amd64
DB_DIR='./back/db/'

all:
	@tree ${GOPATH} |grep gin-gonic >> /dev/null && tree ${GOPATH} |grep go-sql-driver >> /dev/null || make deps
	@$(GOBUILD) -o $(BINNAME)
	@./$(BINNAME) &

deps:
	@$(GOGET) github.com/go-sql-driver/mysql
	@$(GOGET) github.com/gin-gonic/gin
	@echo "Done"

build:
	@$(GOBUILD) -o $(BINNAME)

run:
	@$(GOBUILD) -o $(BINNAME)
	@./$(BINNAME)
