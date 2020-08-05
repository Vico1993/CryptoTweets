FROM golang:latest

WORKDIR /go/src/github.com/Vico1993/CryptoTweets/src/fetcher

COPY . .

RUN ls -la

RUN go get -d -v ./...
RUN go install -v ./...

## THE LIFE SAVER
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait

CMD "/wait" && "fetcher"