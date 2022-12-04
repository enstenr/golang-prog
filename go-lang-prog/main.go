package main
import "fmt"
type void1 struct{}
var member void1
func main() {
    set := make(map[string]void1)
    set["apple"] = member
    set["orange"] = member
    set["mango"] = member
    delete(set, "apple")
    for k := range set {
        fmt.Println(k)
    }
}