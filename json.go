package files

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/cmarkh/errs"
)

func GetJSON[V any](url string) (data V, err error) {
	resp, err := http.Get(url)
	if err != nil {
		err = errs.WrapErr(err)
		return
	}
	defer resp.Body.Close()

	//PrintResp(resp)

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		err = errs.WrapErr(err)
		return
	}

	return
}

func GetJSONFile[V any](path string) (data V, err error) {
	file, err := os.Open(path)
	if err != nil {
		err = errs.WrapErr(err)
		return
	}
	defer file.Close()

	//tools.PrintResp(resp)

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		err = errs.WrapErr(err)
		return
	}

	return
}
