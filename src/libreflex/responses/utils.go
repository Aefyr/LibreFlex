package responses

import (
	"encoding/base64"
	"encoding/json"
)

func jsonEncode(obj interface{}) []byte {
	b, _ := json.Marshal(obj)
	return b
}

func b64encode(data []byte) []byte {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(buf, data)
	return buf
}
