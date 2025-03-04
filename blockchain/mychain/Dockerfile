FROM golang:1.23
RUN curl https://get.ignite.com | bash

WORKDIR /app
COPY go.mod go.sum ./
RUN go.mod download
COPY . .

EXPOSE 26657 9090 1317
CMD ["ignite", "chain", "serve"]