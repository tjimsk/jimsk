FROM golang:1.11 as builder
ENV GOBIN /go/bin/
ENV APPSRC .
ENV APPPATH $GOBIN/jimsk
ENV APPDIR /root/jimsk
# copy go.mod only and get dependencies first
COPY $APPSRC/go.mod $APPDIR/go.mod
WORKDIR $APPDIR
RUN go mod download
# copy the rest of the app directory
COPY $APPSRC $APPDIR
RUN CGO_ENABLED=0 GOOS=linux go build -o $APPPATH -a -installsuffix cgo $APPDIR/main/main.go

FROM alpine:3.4 as image
ENV GOBIN /go/bin/
ENV GOSTATIC /etc/jimsk/static
ENV APPPATH $GOBIN/jimsk
ENV GOPORT :80
COPY static $GOSTATIC
COPY --from=builder $APPPATH $APPPATH
EXPOSE 80
ENTRYPOINT $APPPATH
