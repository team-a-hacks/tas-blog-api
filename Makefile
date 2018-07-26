REPO=github.com/team-a-hacks/tas-blog-api

## generate swagger
gen-swagger:
	# swaggerファイル生成
	swagger -apiPackage="$(REPO)" -mainApiFile="$(REPO)/main.go" -output=./swagger-ui
	sed -i -e "s/package main/package swaggerui/g" ./swagger-ui/docs.go
	sed -i -e "s/\"basePath\": \"http:\/\/localhost:1323\/swagger-ui\"/\"basePath\": \"http:\/\/localhost:1323\"/g" swagger-ui/docs.go
	sed -i -e "1,/\"basePath\": \"http:\/\/localhost:1323\"/s/\"basePath\": \"http:\/\/localhost:1323\"/\"basePath\": \"http:\/\/localhost:1323\/swagger-ui\"/" swagger-ui/docs.go

# build and run
build:
	go build
	./tas-blog-api
