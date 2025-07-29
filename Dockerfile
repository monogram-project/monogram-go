# ---------
# Build stage
# ---------
FROM golang:1.23.8-alpine AS builder
WORKDIR /app

# Copy the Go project directory (adjust the context below if needed)
COPY . .

# Install essential tools
RUN apk update && apk add --no-cache wget tar jq

# Install Just (latest version) from GitHub releases
RUN DOWNLOAD_URL=$(wget -qO- https://api.github.com/repos/casey/just/releases/latest | jq -r '.assets[] | select(.name | contains("x86_64-unknown-linux-musl.tar.gz")) | .browser_download_url') && \
    wget -qO- "$DOWNLOAD_URL" | tar xvz && \
    mv just /usr/local/bin/

# Run the Justfile recipe to build the monogram executable
RUN just build-for-docker

# ----------
# Runtime stage
# ----------
FROM alpine:3.18
WORKDIR /app

# Copy the monogram binary from the builder stage
COPY --from=builder /app/monogram /app/monogram

# Ensure the binary is executable
RUN chmod +x /app/monogram

# Create directories for input/output file mounting
RUN mkdir -p /app/input /app/output

# Expose the port that the --test flag uses when running in a container.
EXPOSE 8080

# Set entrypoint to allow arguments to pass through
ENTRYPOINT ["/app/monogram"]
CMD []
