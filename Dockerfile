FROM golang:1.11.1-stretch

WORKDIR /go/src/app

COPY . .
RUN mv go-ethereum /go/src/.

RUN bash build.sh
RUN make sol build install

CMD ["run_backend"]
