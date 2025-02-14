FROM golang:latest


WORKDIR /app

COPY go.mod go.sum ./

# Get Dependancies 
RUN go mod download

COPY . /app

RUN go build -o /fetch-app

EXPOSE 8080

CMD ["/fetch-app"]