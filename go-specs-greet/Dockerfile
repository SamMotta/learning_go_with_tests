FROM golang:1.22.5

WORKDIR /app

ARG bin_to_build

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr cmd/${bin_to_build}/main.go

EXPOSE 8080
CMD [ "./svr" ]