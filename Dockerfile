FROM golang:latest

WORKDIR /Ambitine
COPY . .

#EXPOSE 9090

RUN go get -u

CMD go run main.go