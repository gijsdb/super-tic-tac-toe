package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gijsdb/super-tic-tac-toe/internal/api/models"
	"github.com/inconshreveable/log15"
)

func GenerateResponseError(message string, rerr error, debugStack string) (res []byte, error interface{}) {
	res, err := json.Marshal(models.ResponseError{
		Err: fmt.Sprintf("%s: %s", message, rerr.Error()),
		//	Msg:        message,
		Stacktrace: debugStack,
	})
	//log15.Error("response error", "err", err, "res", string(res))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}

func StringToInt(str string) int {
	out, err := strconv.Atoi(str)
	if err != nil {
		log15.Crit("Error converting string to int in StringToInt()::helpers.go", "err", err)
		return -1
	}
	return out
}
