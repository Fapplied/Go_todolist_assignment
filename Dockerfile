FROM golang:1.12.0-alpine3.9
RUN mkdir /app 
ADD . /app
WORKDIR /app
ENV DB_URL = "host=snuffleupagus.db.elephantsql.com user=bcpnkdnx password=Z7eR4RPpnjmz0IF3B-NwzhV0S7jE53cW dbname=bcpnkdnx port=5432 sslmode=disable"
RUN go build -o main . 
CMD ["/app/main"]