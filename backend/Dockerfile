FROM golang:1.14-alpine as builder

WORKDIR /go/src/app
COPY . .

RUN go get -v

# Build statically linked binary
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a

FROM alpine 

WORKDIR /go/src/app

COPY --from=builder /go/src/app/backend .

# TODO expose ports
# docker build -t bigtwo-container .
# docker run -p 8000:8000 bigtwo-container 
EXPOSE 8000
CMD /go/src/app/backend $PORT
