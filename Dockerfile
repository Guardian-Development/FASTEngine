# simple image to build and test the library
FROM golang:1.13.5-alpine AS build-env

# copy library source
WORKDIR /go/src/github.com/Guardian-Development/FASTEngine
COPY . .

# run the tests with coverage 
RUN go test ./... -cover -v