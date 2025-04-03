# lab5_docker

# Wykonane czynności


Pobranie oraz przeniesienie pliku alpine-minirootfs-3.21.3-aarch64.tar.gz do katalogu projektu

Utowrzenie pliku Dockerfile

Utworzenie aplikacji main.go

Utworzenie pliku nginx.conf

Zbudowanie obrazu Docker
docker build --build-arg VERSION=1.0.0 -t my-go-app .

Uruchomienie kontenera 
docker run -d -p 8080:80 --name my-running-app my-go-app

Sprawdzenie działania kontenera
docker ps

Sprawdzenie ustawionych portów
docker port my-running-app

Uruchomienie aplikacji w przeglądarce 
http://localhost:8080/
