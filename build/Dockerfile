FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod download
 
RUN go build -o ./out/luckydraw cmd/main.go

EXPOSE 9090
 
CMD ["/cmd"]