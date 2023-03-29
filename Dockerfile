FROM golang:1.20 AS build

WORKDIR /root/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o dist/ -ldflags "-s -w"

FROM gcr.io/distroless/base

WORKDIR /root/
COPY --from=build /root/app/dist .

EXPOSE 8080

CMD ["./club-portal"]
