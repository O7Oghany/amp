APP?=amp
GO111MODULES=auto
REGISTRY?=abdelmmu/images
COMMIT_SHA=$(shell git rev-parse --short HEAD)
AMPSDK_PATH=$(GOPATH/src/github.com/abdelmmu/${APP})
.PHONY: setup
##setup: will make new mod for you
setup:  
	@go mod init \
		&& go mod tidy


.PHONY: format
##format: will run go vet as well as golint
format:
	@echo "formating phase"
	go vet ./... 
	golint -set_exit_status $(go list ./...)

.PHONY: build
##build: runs go get for the whole project and then buid out the binary based on the APPName
build: clean
	@echo "starting to get the Prjoect "
	go get -t -v ./...
	go build -o ${APP} main.go

.PHONY: test
##test: will run go test for the whole project
test:
	@echo "starting Test Phase " 
	go test -cover -v -race ./...   

.PHONY: test_report
##test_report: generat html report for test cover
test_report:
	go test -cover ./... -coverprofile=${APP}.out -covermode count
	go tool cover -html=${APP}.out -o ${APP}_test_report.html

.PHONY: test_bench
##test_bench: generat benchmark test
test_bench:
	@echo "starting Test Phase "
	go test -bench ./... -benchtime 10s


.PHONY: run
##run: will run go run --race
run:
	@echo "running phase"
	go run -race main.go

.PHONY: clean
##clean: will remove the binary of the app and run go clean as well
clean:
	@echo "Cleaning"
	go clean
	rm -rf ${APP} ${APP}.out ${APP}_test_report.html

#helper rule for deployment
check-environment:
ifndef ENV
	$(error ENV not set, allowed values - `staging` or `production`)
endif


.PHONY: docker-build
##docker-build: will build tha image and tage it with the ${APP}:${COMMIT_SHA}	
docker-build: build
	docker build -t ${APP}:${COMMIT_SHA} .

.PHONY: docker-push
##docker-push: will push the image to ${REGISTRY}/${ENV}/${APP}:${COMMIT_SHA} 
docker-push: check-environment docker-build
	docker push ${REGISTRY}/${ENV}/${APP}:${COMMIT_SHA}

.PHONY: help
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


