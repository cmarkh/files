package files

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/cmarkh/errs"
	"github.com/jszwec/csvutil"
)

// ReadCSV reads a csv
func ReadCSV(filename string) (records [][]string, err error) {
	csvfile, err := os.Open(filename)
	if err != nil {
		err = errs.WrapErr(err, "Couldn't open the csv file")
		return
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)

	return r.ReadAll()
}

// CSVToStruct unmarshalls the given CSV to the given slice of structs
func CSVToStruct(filename string, strct interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %s", r)
		}
	}()

	csvfile, err := os.Open(filename)
	if err != nil {
		return errs.WrapErr(err, "Couldn't open the csv file")
	}

	csvReader := csv.NewReader(csvfile)

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return errs.WrapErr(err)
	}

	err = dec.Decode(&strct)
	if err != nil {
		return errs.WrapErr(err)
	}

	return
}
