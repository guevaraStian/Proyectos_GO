// Captura de pantalla del navegador hecho con el lenguaje de programacion GO
// go get -u github.com/chromedp/chromedp
// Se crea la imagen de la pantalla solicitada screenshot
package main

import (
	"fmt"
	"math"
	"net"
	"sort"
	"time"
)

// Puerto y tiempo de respuesta
type PortInfo struct {
	Port int
	Time float64 // tiempo en milisegundos
}

// Escaneo de puertos
func scanPort(host string, port int) PortInfo {
	start := time.Now()
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	elapsed := time.Since(start).Seconds() * 1000 // en milisegundos

	if err != nil {
		return PortInfo{Port: port, Time: elapsed}
	}
	conn.Close()
	return PortInfo{Port: port, Time: elapsed}
}

// K-means clustering básico (1D)
func kmeans(data []float64, k int, maxIter int) []int {
	n := len(data)
	labels := make([]int, n)
	centroids := make([]float64, k)

	// Inicialización simple: primeros k valores como centroides
	copy(centroids, data[:k])

	for iter := 0; iter < maxIter; iter++ {
		// Asignación de clusters
		for i := 0; i < n; i++ {
			minDist := math.MaxFloat64
			for j := 0; j < k; j++ {
				dist := math.Abs(data[i] - centroids[j])
				if dist < minDist {
					minDist = dist
					labels[i] = j
				}
			}
		}

		// Recalcular centroides
		count := make([]int, k)
		sum := make([]float64, k)
		for i := 0; i < n; i++ {
			sum[labels[i]] += data[i]
			count[labels[i]]++
		}
		for j := 0; j < k; j++ {
			if count[j] > 0 {
				centroids[j] = sum[j] / float64(count[j])
			}
		}
	}

	return labels
}

func main() {
	host := "scanme.nmap.org" // sitio público de prueba
	ports := []int{20, 43, 445}
	var results []PortInfo
	var times []float64

	fmt.Println("Escaneando puertos...")
	for _, port := range ports {
		info := scanPort(host, port)
		results = append(results, info)
		times = append(times, info.Time)
	}

	labels := kmeans(times, 2, 10) // Agrupar en 2 clusters

	// Mostrar resultados
	fmt.Println("\nResultados:")
	for i, info := range results {
		fmt.Printf("Puerto %d - Tiempo %.2f ms - Cluster %d\n", info.Port, info.Time, labels[i])
	}

	// Opcional: ordenado por tiempo
	sort.Slice(results, func(i, j int) bool {
		return results[i].Time < results[j].Time
	})
}
