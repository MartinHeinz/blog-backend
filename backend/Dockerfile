FROM golang as builder

COPY src /go/src
RUN go get -d ./src/app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./src/app

FROM alpine:latest

RUN apk --no-cache add ca-certificates
COPY --from=builder /go/app .

CMD ./app
EXPOSE 1234

# docker run -p 1234:1234 <image_id>s