FROM golang:latest

WORKDIR /go/src/github.com/Vico1993/CryptoTweets/src

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

# RUN GOPROXY="direct" go get -v -u go.mongodb.org/mongo-driver/mongo
# RUN GOPROXY="direct" go get -d -v ./...
RUN dep init
RUN dep ensure
RUN go install -v ./...

## THE LIFE SAVER
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

CMD "/wait" && "src"