FROM golang:1.16.5-buster

RUN apt-get update && apt-get install -y libgtk-3-dev libappindicator3-dev

ADD entrypoint.sh /entrypoint.sh

RUN chmod 755 /entrypoint.sh

ENTRYPOINT /entrypoint.sh