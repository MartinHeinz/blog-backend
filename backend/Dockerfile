FROM golang as builder

COPY src /go/src
RUN go get -d ./src/app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./src/app

FROM alpine:latest

RUN apk --no-cache add ca-certificates
COPY --from=builder /go/app .

CMD ./app
EXPOSE 8080

# docker run -p 8080:8080 <image_id>