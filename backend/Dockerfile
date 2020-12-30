FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -v

RUN make

# TODO expose ports
# ENTRYPOINT
ENTRYPOINT ["./backend"]