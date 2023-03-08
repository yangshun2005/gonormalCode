package downloadfile

// # ioutil befer golang 1.16 had been drop
import (
	"io/ioutil"
	"net/http"
)

func DownloadFiled(uri string, filename string) {

	// Get the data
	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(filename, data, 0755)

}
