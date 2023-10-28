FROM golang:1.21.3

WORKDIR /app
RUN apt-get update && apt-get install -y git

# Clone app sources
RUN git clone https://github.com/hendri-marcolia/go-crawler.git .

# Build application
RUN go mod download
RUN go build -o fetch

# Entrypoint for execution
ENTRYPOINT ["/app/fetch"]
