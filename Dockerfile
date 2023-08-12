FROM golang:latest

WORKDIR /api

CMD [ "tail", "-f", "/dev/null" ]
