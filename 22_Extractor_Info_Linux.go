// Con el siguiente codigo de GO se extrae la mayor cantidad de datos
// De el Linux donde se ejecute

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
	dirName := fmt.Sprintf("Info_Linux_%s_%s", hostname, timestamp)
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

	// Información general
	guardarSalida("bash", []string{"-c", "uname -a && lsb_release -a"}, "sistema.txt")

	// Red
	guardarSalida("bash", []string{"-c", "ip a && ip r && cat /etc/resolv.conf"}, "red.txt")
	guardarSalida("ss", []string{"-tunap"}, "conexiones.txt")

	// Procesos
	guardarSalida("ps", []string{"aux"}, "procesos.txt")
	guardarSalida("bash", []string{"-c", "ps aux --sort=-%cpu"}, "procesos_detallados.txt")

	// Hardware
	guardarSalida("lscpu", nil, "cpu.txt")
	guardarSalida("free", []string{"-h"}, "memoria.txt")
	guardarSalida("bash", []string{"-c", "lsblk && df -h"}, "discos.txt")
	guardarSalida("lsblk", []string{"-o", "NAME,SIZE,MODEL,SERIAL"}, "discos_detallado.txt")

	// Programas instalados
	guardarSalida("dpkg", []string{"-l"}, "programas.txt") // cambia a rpm -qa si es RHEL/Fedora

	// Servicios
	guardarSalida("systemctl", []string{"list-units", "--type=service"}, "servicios.txt")
	guardarSalida("systemctl", []string{"list-units", "--type=service", "--state=running"}, "servicios_activos.txt")

	// Usuarios
	guardarSalida("cat", []string{"/etc/passwd"}, "usuarios.txt")

	// Variables de entorno
	guardarSalida("printenv", nil, "entorno.txt")

	// BIOS
	guardarSalida("sudo", []string{"dmidecode", "-t", "bios"}, "bios.txt")

	// Eventos recientes
	guardarSalida("journalctl", []string{"-n", "20"}, "eventos.txt")

	// Firewall
	guardarSalida("sudo", []string{"ufw", "status"}, "firewall.txt")

	// Actualizaciones
	guardarSalida("apt", []string{"list", "--upgradable"}, "actualizaciones.txt")

	fmt.Println("(+) Información del sistema guardada en la carpeta", dir)
}

// go build -o recolector.exe main.go
// recolector.exe
