FROM golang:latest

RUN mkdir app/

COPY . app/

WORKDIR app/api/

RUN go build -o ./main

EXPOSE 8080

CMD ["./main"]

