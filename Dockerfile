FROM golang:latest

WORKDIR /Ambitine
COPY . .

#EXPOSE 9090

RUN go mod init github.com/Pickausernaame/Ambitine_backend

RUN go get -u

CMD go run main.go