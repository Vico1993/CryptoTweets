FROM golang:latest

WORKDIR /go/src/github.com/Vico1993/CryptoTweets/src

COPY . .

RUN GOPROXY="direct" go get -d -v ./...
RUN go install -v ./...

## THE LIFE SAVER
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

CMD "/wait" && "src"