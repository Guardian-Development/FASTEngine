# simple image to build and test the library
FROM golang:1.13.5-alpine AS build-env

# copy library source
COPY . .

# run the tests
RUN go test -v