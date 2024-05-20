FROM golang:1.22-bullseye AS build

WORKDIR /movies_apigateway

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

#RUN go build -v -o /usr/local/bin/app 

RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o app

FROM scratch

WORKDIR /

COPY --from=build /movies_apigateway/app /bin/app

EXPOSE 3000

CMD [ "app"]
