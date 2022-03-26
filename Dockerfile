# using alpine for small size
FROM golang:1.17-alpine

WORKDIR /app

COPY . .

# get dependencies
RUN go mod tidy

# expose port the application is running on
EXPOSE 8081

# run application
CMD ["go", "run", "."]