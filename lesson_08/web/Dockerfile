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

WORKDIR /app/web/src
RUN go build -o ../../_build/srv

##
## Deploy
##

FROM golang:latest AS run

WORKDIR /app

COPY --from=build /app/_build/srv /app/
COPY --from=build /app/web/src/user.html /app/

EXPOSE 80/tcp

CMD ["./srv"]