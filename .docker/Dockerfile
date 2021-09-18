FROM golang:1.17-buster
RUN apt update
RUN apt install -y git make bash
RUN mkdir -p /src
COPY . /src
WORKDIR /src
RUN make clean build

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
COPY bin/ /src
CMD ["./app"]  

