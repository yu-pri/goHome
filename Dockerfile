FROM golang as build

WORKDIR /build
COPY go.* ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=arm GOARM=7 go build -v -o ./goHome
RUN chmod -R 755 goHome


FROM debian

COPY --from=build /build/goHome goHome
RUN apt-get update && apt-get install -y file
RUN file /goHome > file.txt
EXPOSE 3000
CMD ["/goHome"]
