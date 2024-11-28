FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /loyalty_system

EXPOSE 8080

# Run
CMD [ "/loyalty_system" ]
