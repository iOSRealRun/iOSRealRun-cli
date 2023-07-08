package parseRoute

import (
	"encoding/json"
	"os"
	"strconv"
)

func Split(inp string) (out []map[string]float64) {
	// read file
	content, err := os.ReadFile(inp)
	if err != nil {
		panic(err)
	}
	// add [] to the beginning and end of the content
	content = append([]byte("["), content...)
	content = append(content, []byte("]")...)

	// parse using json
	var tmp []map[string]string
	err = json.Unmarshal(content, &tmp)
	if err != nil {
		panic(err)
	}

	// convert string to float64
	out = make([]map[string]float64, 0)
	for _, i := range tmp {
		tmp2 := make(map[string]float64)
		for k, v := range i {
			tmp2[k], err = strconv.ParseFloat(v, 64)
			if err != nil {
				panic(err)
			}
		}
		out = append(out, tmp2)
	}
	return
}
