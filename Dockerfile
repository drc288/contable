FROM golang:1.18 as builder

# SET WORKING DIRECTYORY
WORKDIR /app
COPY . .

RUN go mod download
RUN go mod tidy
RUN go mod verify
RUN go build -o app .
RUN ls
RUN chmod 755 ./app

RUN chmod -R 777 configuration

EXPOSE 8080

ENTRYPOINT ["./app"]