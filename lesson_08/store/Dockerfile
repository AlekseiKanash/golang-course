##
## Build
##

FROM golang:latest AS build

RUN apt update
RUN apt install -y protobuf-compiler

WORKDIR /app

COPY /build /app

RUN /app/prepare_dependencies.sh
RUN /app/generate_proto.sh

WORKDIR /app/store/src
RUN go build -o ../_build/store

##
## Deploy
##

FROM golang:latest AS run

WORKDIR /app

COPY --from=build /app/store/_build/store /app/

EXPOSE 9000/tcp

CMD ["./store"]