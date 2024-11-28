FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /loyalty_system

EXPOSE 8080

# Run
CMD [ "/loyalty_system" ]
