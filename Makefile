SERVICE_NAME:=go_template
SERVICE_NAMESPACE:=go_template
CREATED_BY:=go_template
build:
	go test -mod vendor ./...
	go build -mod vendor ./...
	docker build --build-arg service_name=${SERVICE_NAME} ${SERVICE_NAMESPACE}/${SERVICE_NAME}:latest .
build-ci:
	GO111MODULE=on go build -mod vendor -o ${SERVICE_NAME}

test:
	go test -mod vendor ./...
create_migration:
	migrate create -ext sql -dir migrations -seq $(name)
