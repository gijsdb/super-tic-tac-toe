package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/gijsdb/super-tic-tac-toe/internal/api/models"
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
