FROM golang:1.17-buster as build
RUN apt update
RUN apt install -y git make bash
RUN mkdir -p /src
COPY . /src
WORKDIR /src
RUN make clean build

FROM alpine:latest
WORKDIR /src
RUN apk --no-cache add ca-certificates
COPY --from=build /src/bin /src
CMD ["/src/app"]

