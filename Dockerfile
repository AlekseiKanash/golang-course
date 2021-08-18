##
## Build
##

FROM golang:latest AS build
WORKDIR /app

COPY / /app

WORKDIR /app/lesson_08/web/

RUN go build -o _build/srv

##
## Deploy
##

FROM golang:latest AS run

WORKDIR /app

COPY --from=build /app/lesson_08/web/_build/srv /app/
COPY --from=build /app/lesson_08/web/user.html /app/

EXPOSE 80/tcp

CMD ["./srv"]