package main
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
func main() {
	pembaca := bufio.NewReader(os.Stdin)
	masukan, _ := pembaca.ReadString('\n')
	masukan = strings.TrimSpace(masukan)
	bagian := strings.Split(masukan, " ")

	var data []int
	for _, bagian := range bagian {
		angka, _ := strconv.Atoi(bagian)
		if angka == -5313 {
		break
	} else if angka == 0 {
		median := hitungMedian(data)
		fmt.Println(median)
	} else {
		data = append(data, angka)
		}
	}
}
func hitungMedian(data []int) int {
	sort.Ints(data)
	panjang := len(data)
	if panjang == 0 {
		return 0
	} else if panjang%2 == 0 {
		return (data[panjang/2-1] + data[panjang/2]) / 2
	} else {
		return data[panjang/2]
	}
}
