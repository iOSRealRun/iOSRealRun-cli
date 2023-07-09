package parseRoute

import (
	"encoding/json"
	"strconv"
)

func Split(inp []byte) (out []map[string]float64) {
	// add [] to the beginning and end of the content, don't modify the original content
	content := make([]byte, len(inp))
	copy(content, inp)
	content = append([]byte("["), content...)
	content = append(content, []byte("]")...)

	// parse using json
	var tmp []map[string]string
	err := json.Unmarshal(content, &tmp)
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
