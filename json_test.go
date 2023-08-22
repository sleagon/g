package g

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJMap(t *testing.T) {

	is := assert.New(t)

	// 测试json对于各种数据解析的能力

	{
		val := `{"a":1,"b":2,"c":3,"d":4,"e":5,"f":6,"g":7,"h":8}`

		data := make(map[string]interface{})

		err := json.Unmarshal([]byte(val), Ptr(JMap(data)))

		is.Nil(err)
		is.Equal(8, len(data))
		is.Equal(int64(1), data["a"])
		is.Equal(int64(2), data["b"])
		is.Equal(int64(3), data["c"])
	}

	// 测试对于其他类型的解析能力
	{
		val := `{"name": "golang", "age": 10, "score": 99.9, "is_ok": true, "is_not_ok": false, "nil": null}`
		data := make(map[string]interface{})
		err := json.Unmarshal([]byte(val), Ptr(JMap(data)))
		is.Nil(err)
		is.Equal(6, len(data))
		is.Equal("golang", data["name"])
		is.Equal(int64(10), data["age"])
		is.Equal(99.9, data["score"])
		is.Equal(true, data["is_ok"])
		is.Equal(false, data["is_not_ok"])
		is.Equal(nil, data["nil"])
	}

	// 测试bigint的精度能力
	{
		val := `{"a": 9223372036854775807, "b": 1.2e8}`
		data := make(map[string]interface{})
		err := json.Unmarshal([]byte(val), Ptr(JMap(data)))
		is.Nil(err)
		is.Equal(2, len(data))
		is.Equal(int64(9223372036854775807), data["a"])
		is.Equal(1.2e8, data["b"])
	}
}
