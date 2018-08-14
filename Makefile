REPO=github.com/team-a-hacks/tas-blog-api
HOST=https:\/\/tas-blog-api.herokuapp.com

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

# deploy to heroku
deploy:
	git checkout -b deploy
	dep ensure
	sed -i -e "s/\"basePath\": \"http:\/\/localhost:1323\"/\"basePath\": \"$(HOST)\"/g" swagger-ui/docs.go
	sed -i -e "s/\"basePath\": \"http:\/\/localhost:1323\/swagger-ui\"/\"basePath\": \"$(HOST)\/swagger-ui\"/g" swagger-ui/docs.go
	GOOS=linux GOARCH=amd64 go build -o db/seeds/seeds db/seeds/seed.go
	GOOS=linux GOARCH=amd64 go build
	git add -f tas-blog-api db/seeds/seeds
	git add swagger-ui/docs.go
	git commit -m 'deploy to heroku'
	git push -f heroku deploy:master
