package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {

	dir, err := os.Getwd()
	checkErr(err)

	f, err := os.Open(dir + "/movies.csv")
	checkErr(err)
	defer f.Close()

	o, err := os.Create(dir + "/output.txt")
	defer o.Close()

	reader := bufio.NewReader(f)

	var line string
	var lineCount int

	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			checkErr(err)
		}

		if lineCount == 0 {
			header := fmt.Sprintf("Movie Title|Movie Description|Movie Date\n")
			o.WriteString(header)
			lineCount++
			continue
		} else {
			lineCount++
		}

		r := csv.NewReader(strings.NewReader(string(line)))
		r.Comma = ','
		records, err := r.ReadAll()
		checkErr(err)
		for _, column := range records {
			mydate, err := time.Parse("1/02/06", column[2])
			checkErr(err)
			data := fmt.Sprintf("%s|%s|%s\n", column[0], column[1], mydate.Format("01 02 2006"))
			o.WriteString(data)

		}

	}
	o.Sync()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
