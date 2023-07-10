package parseRoute

import (
	"encoding/json"
	"strconv"

	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/config"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/logger"
)

var myLogger = logger.NewMyLogger("log.log", config.Config.LogLevel)

func ParseRoute(inp []byte) (out []map[string]float64) {
	// add [] to the beginning and end of the content, don't modify the original content
	myLogger.Infoln("Start parsing route ...")
	myLogger.Debugln("inp:", string(inp))
	content := make([]byte, len(inp))
	copy(content, inp)
	content = append([]byte("["), content...)
	content = append(content, []byte("]")...)

	// parse using json
	var tmp []map[string]string
	err := json.Unmarshal(content, &tmp)
	if err != nil {
		myLogger.Errorln("Parse route failed")
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
	myLogger.Infoln("Parse route successfully")
	return
}
