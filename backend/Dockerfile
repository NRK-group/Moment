FROM golang:1.19.3-alpine3.16 AS build
WORKDIR /moment-backend
COPY . .
RUN apk add --update gcc g++
RUN go build -o moment-backend .

# # # Run stage
FROM alpine:3.17
WORKDIR /moment-backend
COPY --from=build /moment-backend .
EXPOSE 5070
CMD ["./moment-backend"]