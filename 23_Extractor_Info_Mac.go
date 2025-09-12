// Con el siguiente codigo de GO se extrae la mayor cantidad de datos
// De el MAC donde se ejecute

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
	dir := fmt.Sprintf("Info_Mac_%s_%s", hostname, timestamp)
	os.MkdirAll(dir, 0755)
	os.Chdir(dir)
	return dir
}

func guardarSalida(comando string, args []string, archivo string) {
	cmd := exec.Command(comando, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		out = append(out, []byte("\n[ERROR]: "+err.Error())...)
	}
	os.WriteFile(archivo, out, 0644)
}

func main() {
	dir := crearDirectorio()
	fmt.Println("📂 Guardando información en:", dir)

	guardarSalida("system_profiler", []string{"SPSoftwareDataType"}, "sistema.txt")
	guardarSalida("ifconfig", nil, "red.txt")
	guardarSalida("netstat", []string{"-anv"}, "conexiones.txt")
	guardarSalida("ps", []string{"aux"}, "procesos.txt")
	guardarSalida("top", []string{"-l", "1", "-o", "cpu"}, "procesos_detallados.txt")
	guardarSalida("sysctl", []string{"-n", "machdep.cpu.brand_string"}, "cpu.txt")
	guardarSalida("vm_stat", nil, "memoria.txt")
	guardarSalida("df", []string{"-h"}, "discos.txt")
	guardarSalida("diskutil", []string{"list"}, "discos_detallado.txt")
	guardarSalida("ls", []string{"/Applications"}, "programas.txt")
	guardarSalida("launchctl", []string{"list"}, "servicios.txt")
	guardarSalida("dscl", []string{".", "list", "/Users"}, "usuarios.txt")
	guardarSalida("printenv", nil, "entorno.txt")
	guardarSalida("system_profiler", []string{"SPHardwareDataType"}, "bios.txt")
	guardarSalida("log", []string{"show", "--predicate", "eventType == logEvent", "--last", "1h"}, "eventos.txt")
	guardarSalida("defaults", []string{"read", "/Library/Preferences/com.apple.alf", "globalstate"}, "firewall.txt")
	guardarSalida("softwareupdate", []string{"--list"}, "actualizaciones.txt")

	fmt.Println("✅ Información del sistema guardada en:", dir)
}

// go build -o recolector.exe main.go
// recolector.exe
