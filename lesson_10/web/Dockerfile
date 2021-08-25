##
## Build
##

FROM golang:latest AS build

WORKDIR /app

COPY /build /app

RUN /app/prepare_dependencies.sh

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