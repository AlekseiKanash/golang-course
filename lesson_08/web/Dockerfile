##
## Build
##

FROM golang:latest AS build
WORKDIR /app

COPY / /app

WORKDIR /app/src

RUN go build -o _build/srv

##
## Deploy
##

FROM golang:latest AS run

WORKDIR /app

COPY --from=build /app/src/_build/srv /app/
COPY --from=build /app/src/user.html /app/

EXPOSE 80/tcp

CMD ["./srv"]