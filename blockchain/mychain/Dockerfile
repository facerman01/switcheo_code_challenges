FROM golang:1.23

# Install Ignite CLI
RUN curl https://get.ignite.com/cli! | bash

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

EXPOSE 26657 9090 1317
RUN yes | ignite chain build
CMD yes | ignite chain serve