// Con el siguiente codigo de GO se extrae la mayor cantidad de datos
// De el Windows donde se ejecute
// Se crea la imagen de la pantalla solicitada screenshot
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func crearDirectorio() string {
	hostname, _ := os.Hostname()
	timestamp := time.Now().Format("20060102_150405")
	dirName := fmt.Sprintf("Info_Windows_%s_%s", hostname, timestamp)
	os.MkdirAll(dirName, os.ModePerm)
	os.Chdir(dirName)
	return dirName
}

func guardarSalida(comando string, args []string, archivo string) {
	cmd := exec.Command(comando, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		output = append(output, []byte("\n[ERROR]: "+err.Error())...)
	}
	os.WriteFile(archivo, output, 0644)
}

func main() {
	dir := crearDirectorio()
	fmt.Println("(+) Guardando información en:", dir)

	// Sistema
	guardarSalida("systeminfo", nil, "sistema.txt")

	// Red
	guardarSalida("ipconfig", []string{"/all"}, "red.txt")
	guardarSalida("netstat", []string{"-anob"}, "conexiones.txt")

	// Procesos
	guardarSalida("tasklist", nil, "procesos.txt")
	guardarSalida("powershell", []string{"Get-Process | Sort-Object CPU -Descending"}, "procesos_detallados.txt")

	// Hardware
	guardarSalida("wmic", []string{"cpu", "get", "name,NumberOfCores,NumberOfLogicalProcessors"}, "cpu.txt")
	guardarSalida("wmic", []string{"memorychip", "get", "capacity,manufacturer,speed"}, "memoria.txt")
	guardarSalida("wmic", []string{"diskdrive", "get", "model,size,serialnumber"}, "discos.txt")
	guardarSalida("powershell", []string{"Get-PhysicalDisk | Format-Table"}, "discos_detallado.txt")

	// Programas instalados
	guardarSalida("wmic", []string{"product", "get", "name,version"}, "programas.txt")
	guardarSalida("powershell", []string{"Get-WmiObject -Class Win32_Product | Select-Object Name, Version"}, "programas_detallado.txt")

	// Servicios
	guardarSalida("sc", []string{"query"}, "servicios.txt")
	guardarSalida("powershell", []string{"Get-Service | Where-Object { $_.Status -eq 'Running' }"}, "servicios_activos.txt")

	// Usuarios
	guardarSalida("net", []string{"user"}, "usuarios.txt")
	guardarSalida("powershell", []string{"Get-LocalUser"}, "usuarios_detallado.txt")

	// Entorno
	guardarSalida("cmd", []string{"/C", "set"}, "entorno.txt")

	// BIOS
	guardarSalida("wmic", []string{"bios", "get", "manufacturer,smbiosbiosversion"}, "bios.txt")

	// Eventos recientes
	guardarSalida("powershell", []string{"Get-EventLog -LogName System -Newest 20"}, "eventos.txt")

	// Firewall
	guardarSalida("netsh", []string{"advfirewall", "show", "allprofiles"}, "firewall.txt")

	// Actualizaciones
	guardarSalida("powershell", []string{"Get-HotFix"}, "actualizaciones.txt")

	fmt.Println("(+) Información del sistema guardada en la carpeta", dir)
}

// go build -o recolector.exe main.go
// recolector.exe
