FROM golang:latest


WORKDIR /app

COPY go.mod go.sum ./

# Get Dependancies 
RUN go mod download

COPY . .

RUN go build 

EXPOSE 8080

CMD ["app/fetch"]