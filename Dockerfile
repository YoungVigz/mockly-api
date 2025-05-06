# Base image z Node.js i Go
FROM node:22.14.0-bullseye

# Instalacja zależności systemowych
RUN apt-get update && apt-get install -y wget gcc git

# Instalacja Go 1.24.1
RUN wget -qO- https://dl.google.com/go/go1.24.1.linux-amd64.tar.gz | tar -xz -C /usr/local

# Konfiguracja środowiska Go
ENV GOPATH /home/appuser/go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin

# Utworzenie użytkownika nie-root i katalogów z odpowiednimi uprawnieniami
RUN useradd -m -u 1001 appuser && \
    mkdir -p $GOPATH && \
    chown -R appuser:appuser $GOPATH

# Przełącz na użytkownika appuser
USER appuser

# Konfiguracja npm dla użytkownika
RUN npm config set prefix '/home/appuser/.npm-global'
ENV PATH $PATH:/home/appuser/.npm-global/bin

# Instalacja mockly-cli
RUN npm install -g mockly-cli

# Środowisko pracy
WORKDIR /app

# Kopiowanie plików Go mod i pobieranie zależności
COPY --chown=appuser:appuser go.mod go.sum ./
RUN go mod download

# Kopiowanie źródła
COPY --chown=appuser:appuser . .

EXPOSE 8080

# Kompilacja i uruchomienie
ENTRYPOINT ["go", "run", "./cmd/main.go"]