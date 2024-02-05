FROM golang:1.21 AS build

WORKDIR /root/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o dist/ -ldflags "-s -w"

FROM gcr.io/distroless/base-debian12

WORKDIR /root/
COPY --from=build /root/app/dist .

EXPOSE 8080

CMD ["./club-portal"]
