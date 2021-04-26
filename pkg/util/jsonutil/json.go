// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package jsonutil

import (
	"encoding/json"
	"gin-demo/pkg/logger"
	simplejson "github.com/bitly/go-simplejson"
	jsoniter "github.com/json-iterator/go"
)

type Map map[string]interface{}

var jsonx = jsoniter.ConfigCompatibleWithStandardLibrary

func Unmarshal(json string, val interface{}) error {
	return jsonx.UnmarshalFromString(json, val)
}

func Encode(o interface{}) ([]byte, error) {
	return json.Marshal(o)
}

func Decode(y []byte, o interface{}) error {
	return json.Unmarshal(y, o)
}

func ToString(o interface{}) string {
	b, err := Encode(o)
	if err != nil {
		logger.Errorf("Failed to encode [%+v], error: %+v", o, err)
		return ""
	}
	return string(b)
}

// FIXME: need improve performance
func ToJson(o interface{}) Json {
	var j Json
	j = &fakeJson{simplejson.New()}
	b, err := Encode(o)
	if err != nil {
		logger.Errorf("Failed to encode [%+v] to []byte, error: %+v", o, err)
		return j
	}
	j, err = NewJson(b)
	if err != nil {
		logger.Errorf("Failed to decode [%+v] to Json, error: %+v", o, err)
	}
	return j
}

type fakeJson struct {
	*simplejson.Json
}

func NewJson(y []byte) (Json, error) {
	j, err := simplejson.NewJson(y)
	return &fakeJson{j}, err
}

func (j *fakeJson) Get(key string) Json {
	return &fakeJson{j.Json.Get(key)}
}

func (j *fakeJson) GetPath(branch ...string) Json {
	return &fakeJson{j.Json.GetPath(branch...)}
}

func (j *fakeJson) CheckGet(key string) (Json, bool) {
	result, ok := j.Json.CheckGet(key)
	return &fakeJson{result}, ok
}
