FROM golang:alpine

WORKDIR $GOPATH/src/vinda-api

COPY . .

ENV GIN_MODE release
ENV GO111MODULE on
ENV GOPROXY https://goproxy.io

#RUN go mod tidy
#RUN gcc -v && go env && go build -o app .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

EXPOSE 3000

ENTRYPOINT ["./app"]

