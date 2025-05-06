FROM golang:1.22-bullseye

# Node.js (dla mockly-cli)
RUN apt-get update && apt-get install -y curl && \
    curl -fsSL https://deb.nodesource.com/setup_22.x | bash - && \
    apt-get install -y nodejs

# Dodanie użytkownika appuser
RUN useradd -m -u 1001 appuser
USER appuser

# Instalacja mockly-cli
RUN npm install -g mockly-cli

WORKDIR /app
COPY --chown=appuser:appuser go.mod go.sum ./
RUN go mod download
COPY --chown=appuser:appuser . .

EXPOSE 8080
ENTRYPOINT ["go", "run", "./cmd/main.go"]
