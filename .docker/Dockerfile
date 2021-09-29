FROM golang:1.17-buster as build
RUN apt update
RUN apt install -y git make bash
RUN mkdir -p /service
COPY . /service
WORKDIR /service
RUN make clean build

FROM alpine:latest
WORKDIR /service
RUN apk --no-cache add ca-certificates
COPY --from=build /service/build/bin/app /service/bin/app
COPY --from=build /service/src/static /service/src/static
CMD ["/service/bin/app"]

