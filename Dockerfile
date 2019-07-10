FROM golang:1.12

LABEL maintainer="Kaue Bonfim <kaueoliveir95@hotmail.com>"

RUN go get -u -v github.com/KaueBonfim/desafio-golang-translate
WORKDIR $GOPATH/src/github.com/KaueBonfim/desafio-golang-translate

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go build -v app.go

# This container exposes 
EXPOSE 8080
EXPOSE 5080
RUN chmod 777 ./app

ENTRYPOINT [ "entry.sh" ]