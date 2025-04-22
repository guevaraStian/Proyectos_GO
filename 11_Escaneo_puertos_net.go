package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	Ruta_victima := "scanme.nmap.org" // Puedes cambiarlo por una IP o dominio
	Puerto_inicial := 1
	Puerto_final := 56065

	fmt.Printf("Empieza el escaneo de los puertos %s (%d-%d)...\n", Ruta_victima, Puerto_inicial, Puerto_final)

	for port := Puerto_inicial; port <= Puerto_final; port++ {
		address := fmt.Sprintf("%s:%d", Ruta_victima, port)

		// Intentar conectarse con timeout de 500ms
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		if err != nil {
			continue // puerto cerrado o inaccesible
		}
		conn.Close()
		fmt.Printf("[+] Puerto %d abierto\n", port)
	}
}
