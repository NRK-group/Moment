package Test

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"

	"backend/pkg/auth"
)

func TestWriteImage(t *testing.T) {
	t.Run("Upload valid file", func(t *testing.T) {
		file, err := os.ReadFile("./TestImages/golang1.png")
		if err != nil {
			t.Errorf("Failed to open file: %s", "./TestImages/golang1.png")
		}

		data := url.Values{}
		// res, _ := json.Marshal(data)
		data.Add("file", string(file))
		fmt.Println()
		fmt.Println()
		fmt.Println("FILE ===== ", string(file))
		fmt.Println()
		fmt.Println()

		req := &http.Request{
			Method: "POST",
			Form: data,
		}

		// req.Form
		// req.Header["Accept-Encoding"] =
		fmt.Println()
		fmt.Println()
		fmt.Println("FILE VALUE ---- ", req.Header)
		fmt.Println()
		fmt.Println()
		req.Header.Set("content-type", "multipart/form-data")

		valid, resDir := auth.WriteImage("./TestImages", req)
		if !valid {
			t.Errorf("Problem writing function: %v", resDir)
		}
		fmt.Println(resDir)
	})
}
