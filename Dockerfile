FROM ubuntu:18.04

LABEL name="Martynov Anton"

ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

USER root
ENV DEBIAN_FRONTEND 'noninteractive'

RUN apt-get update -y
RUN apt-get install -y --no-install-recommends apt-utils

RUN apt-get install -y git wget

RUN wget https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.12.5.linux-amd64.tar.gz

ENV GOROOT /usr/local/go
ENV GOPATH /opt/go
ENV PATH $GOROOT/bin:$GOPATH/bin:/usr/local/go/bin:$PATH

RUN apt-get update && apt-get install -y wget gnupg &&     wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -

#RUN echo "deb http://apt.postgresql.org/pub/repos/apt bionic-pgdg main" > /etc/apt/sources.list.d/PostgreSQL.list

#RUN apt-get update && apt-get install -y postgresql-11

WORKDIR /Ambitine

COPY . .

EXPOSE 9090

RUN go get -u

USER postgres

RUN /etc/init.d/postgresql start &&\
    psql -c "CREATE USER ambitine WITH SUPERUSER PASSWORD '1488';" &&\
    createdb -O ambitine ambitine &&\
    psql -c "GRANT ALL ON DATABASE ambitine TO ambitine;" &&\
    psql -d forum -c "CREATE EXTENSION IF NOT EXISTS citext;" &&\
    /etc/init.d/postgresql stop

USER root

#VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

CMD go run main.go