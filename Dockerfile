FROM golang:1.17

RUN mkdir /app

WORKDIR /app 

COPY . ./

RUN go mod download

RUN go build -o ./builder

EXPOSE 8000

CMD ./builder