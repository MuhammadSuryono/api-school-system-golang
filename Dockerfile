# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./ 
RUN go mod download
COPY . ./
RUN go build -o /api-school-system
EXPOSE 8080
CMD ["/api-school-system"]


##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /api-school-system /api-school-system

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/api-school-system"]
