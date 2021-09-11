# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./api-school-system .
EXPOSE 8080
CMD [ "./api-school-system" ]