# Lab 5 - Docker

## Wykonane czynności

1. **Przygotowanie plików**:
   - Pobranie obrazu Alpine: `alpine-minirootfs-3.21.3-aarch64.tar.gz`
   - Utworzenie plików: `Dockerfile`, `main.go`, `nginx.conf`

2. **Budowanie obrazu**:
   ```bash
   docker build --build-arg VERSION=1.0.0 -t my-go-app .

3. Uruchomienie kontenera:
   ```bash
   docker run -d -p 8080:80 --name my-running-app my-go-app

4. **Sprawdzenie działania kontenera:
   ```bash
   docker ps

5. **Sprawdzenie ustawionych portów:
   ```bash
   docker port my-running-app

6. **Uruchomienie aplikacji w przeglądarce:
   [http://localhost:8080/](http://localhost:8080/)
