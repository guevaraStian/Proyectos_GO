// Se descargan las librerias que se usaran en el proyecto
// go get github.com/google/gopacket
// go get github.com/google/gopacket/pcap

package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Lista las interfaces disponibles
	interfaces, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Interfaces disponibles:")
	for i, device := range interfaces {
		fmt.Printf("[%d] %s - %s\n", i, device.Name, device.Description)
	}

	// Selecciona la primera interfaz (puedes cambiarlo por otra)
	if len(interfaces) == 0 {
		log.Fatal("No se encontraron interfaces de red.")
	}
	device := interfaces[0].Name

	// Abre la interfaz
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Empieza a capturar paquetes
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println("Escuchando paquetes en la interfaz:", device)

	for packet := range packetSource.Packets() {
		fmt.Println("Paquete capturado:", packet)
	}
}
