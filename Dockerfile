FROM golang:1.21.4

RUN apt update && apt upgrade -y \
    && apt install -y git\
    make openssh-client

WORKDIR /go/src/app

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

COPY . .
RUN go mod tidy \
    && go mod verify


EXPOSE 3069

ENTRYPOINT [ "air" ]