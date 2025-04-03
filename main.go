package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

var version = "unknown" // Domyślna wersja, nadpisywana w czasie kompilacji

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		ip := getIPAddress()

		fmt.Fprintf(w, "IP: %s\nHostname: %s\nVersion: %s\n", ip, hostname, version)
	})

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

// Pobieranie adresu IP
func getIPAddress() string {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "unknown"
}
