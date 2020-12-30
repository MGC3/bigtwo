FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -v

RUN make

# TODO expose ports
# docker build -t bigtwo-container .
# docker run -p 8000:8000 bigtwo-container 
EXPOSE 8000
ENTRYPOINT ["./backend"]