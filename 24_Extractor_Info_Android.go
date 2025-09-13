// Con el siguiente codigo de GO se extrae la mayor cantidad de datos
// De el Android donde se ejecute
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func crearDirectorio() string {
	timestamp := time.Now().Format("20060102_150405")
	dir := "Info_Android_" + timestamp
	_ = os.MkdirAll(dir, 0755)
	_ = os.Chdir(dir)
	return dir
}

func guardarSalidaADB(comando string, archivo string) {
	cmd := exec.Command("adb", "shell", comando)
	out, err := cmd.CombinedOutput()
	if err != nil {
		out = append(out, []byte("\n[ERROR]: "+err.Error())...)
	}
	os.WriteFile(archivo, out, 0644)
}

func main() {
	dir := crearDirectorio()
	fmt.Println("Guardando información en:", dir)

	// Información general del sistema
	guardarSalidaADB("getprop", "sistema.txt")

	// Información de red
	guardarSalidaADB("ip addr show", "red.txt")
	guardarSalidaADB("netstat", "conexiones.txt")

	// Procesos
	guardarSalidaADB("ps", "procesos.txt")
	guardarSalidaADB("top -n 1", "procesos_detallados.txt")

	// Información de hardware
	guardarSalidaADB("cat /proc/cpuinfo", "cpu.txt")
	guardarSalidaADB("cat /proc/meminfo", "memoria.txt")
	guardarSalidaADB("df -h", "discos.txt")
	guardarSalidaADB("lsblk", "discos_detallado.txt") // puede no existir

	// Programas instalados
	guardarSalidaADB("pm list packages -f", "programas.txt")

	// Servicios (limitado sin root)
	guardarSalidaADB("service list", "servicios.txt")

	// Usuarios
	guardarSalidaADB("whoami", "usuarios.txt")

	// Variables de entorno
	guardarSalidaADB("printenv", "entorno.txt")

	// Información de BIOS (limitada en Android)
	guardarSalidaADB("getprop ro.bootloader && getprop ro.hardware", "bios.txt")

	// Eventos (logs)
	guardarSalidaADB("logcat -d -t 100", "eventos.txt")

	// Firewall (solo si el dispositivo está rooteado)
	guardarSalidaADB("su -c iptables -L", "firewall.txt")

	// Actualizaciones (no estándar)
	guardarSalidaADB("dumpsys package", "actualizaciones.txt")

	fmt.Println("(+) Información guardada en la carpeta:", dir)
}

// go build -o recolector.exe main.go
// recolector.exe
