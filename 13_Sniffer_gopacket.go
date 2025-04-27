// En este codigo se muestra como se puede crear
// un sniffer con la libreria gopacket de en lenguaje de programacion Go
// Primero instalamos la libreria que usaremos
// go get github.com/google/gopacket
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Se muestra en pantalla la lista de interfaces de red disponibles
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Interfaces disponibles:")
	for i, device := range devices {
		fmt.Printf("[%d] %s - %s\n", i, device.Name, device.Description)
	}

	// Elige una interfaz de red (puedes cambiarlo por el nombre directamente)
	device := devices[0].Name // <- Aquí puedes cambiar a "eth0", "wlan0", etc.

	// Abrir la interfaz para captura
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	fmt.Println("Capturando paquetes en:", device)

	// Crear un paquete fuente
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Mostrar en pantalla los paquetes que pasen por la interface de red
	for packet := range packetSource.Packets() {
		fmt.Println("\n----------------------")
		fmt.Println("El tiempo es :", time.Now())
		fmt.Println("La descripcion del paque es:", packet)
	}
}
