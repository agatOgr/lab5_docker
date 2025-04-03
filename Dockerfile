# Etap 1: Kompilacja aplikacji w Go
FROM golang:alpine AS builder

WORKDIR /app
COPY main.go .

# Inicjalizacja modułu Go (jeśli potrzebna)
RUN go mod init lab5 || true
RUN go mod tidy || true

# Kompilacja aplikacji z dynamicznym nadaniem wersji
ARG VERSION="dev"
RUN go build -ldflags="-X main.version=$VERSION" -o app .

# Etap 2: Tworzenie minimalnego obrazu
FROM scratch

# Dodanie systemu plików Alpine (mini rootfs)
ADD alpine-minirootfs-3.21.3-aarch64.tar.gz /

# Skopiowanie skompilowanej aplikacji
COPY --from=builder /app/app /app

# Ustawienie domyślnego portu
EXPOSE 8080

# Uruchomienie aplikacji
CMD ["/app"]


