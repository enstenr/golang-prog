package main 
import (
	"bufio"
	"strconv"
	"os"
	"fmt"
	"strings"
)
type Memory struct {
    MemTotal     int
    MemFree      int
    MemAvailable int
}
func main(){file, err := os.Open("/proc/meminfo")
if err != nil {
	panic(err)
}
defer file.Close()
bufio.NewScanner(file)
scanner := bufio.NewScanner(file)
res := Memory{}
for scanner.Scan() {
	
	key, value := parseLine(scanner.Text())
	switch key {
	case "MemTotal":
		res.MemTotal = value
	case "MemFree":
		res.MemFree = value
	case "MemAvailable":
		res.MemAvailable = value
	}
}
}
func parseLine(raw string) (key string, value int) {
    fmt.Println(raw)
    text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
    keyValue := strings.Split(text, ":")
    return keyValue[0], toInt(keyValue[1])
}

func toInt(raw string) int {
    if raw == "" {
        return 0
    }
    res, err := strconv.Atoi(raw)
    if err != nil {
        panic(err)
    }
    return res
}
