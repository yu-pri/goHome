FROM arm32v7/golang as build

WORKDIR /build
COPY go.* ./
RUN go mod download

COPY . .
RUN go build -v -o ./goHome

FROM scratch

COPY --from=build /build/goHome goHome

EXPOSE 3000
ENTRYPOINT goHome
