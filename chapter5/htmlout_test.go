package htmlout

import (
	// "net/http"
	"testing"
)

// func TestHTMLOut(t *testing.T) {
// 	resp, _ := http.Get("https://www.google.com")
// 	Output(resp.Body)
// 	resp.Body.Close()
// }

// func TestTopoSort(t *testing.T) {
// 	fmt.Println(topoSort(prereqs))
// }

func TestPermute(t *testing.T) {
	permute([]int{0, 1})
}
