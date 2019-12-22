FROM golang:latest

WORKDIR /Gorush
COPY . .

EXPOSE 8088

CMD ./gorush ${PWD}/config.yml:/config.yml
