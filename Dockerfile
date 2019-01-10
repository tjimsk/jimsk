FROM golang:1.11 as builder
ENV GOBIN /go/bin
ENV APP_SRC .
ENV APP_DIR /root/jimsk
ENV APP_BIN $GOBIN/jimsk

# copy go.mod only and get dependencies first
COPY $APP_SRC/go.mod $APP_DIR/go.mod
WORKDIR $APP_DIR
RUN go mod download
# copy the rest of the app directory
COPY $APP_SRC $APP_DIR
RUN CGO_ENABLED=0 GOOS=linux go build -o $APP_BIN -a -installsuffix cgo $APP_DIR/main

FROM alpine:3.4 as image
ENV GOBIN /go/bin
ENV APP_STATIC /etc/jimsk/static
ENV APP_BIN $GOBIN/jimsk
ENV APP_PORT :80

COPY static $APP_STATIC
COPY --from=builder $APP_BIN $APP_BIN
EXPOSE 80
ENTRYPOINT $APP_BIN
