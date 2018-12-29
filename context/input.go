package context


import(
  "net/http"
  "strconv"
  // "fmt"
)

type InputParams map[string]interface{}

type Input struct {
  Params InputParams
  Request *http.Request
}

func (i *Input) GetString(key string, defaultValue string) string{
	value, ok := i.Params[key]

  if ok == false{
    return defaultValue
  }

  if ret, ok := value.(string); ok{
    return ret
  }

  ret, ok := value.([]string)
	if ok{
		return ret[0]
	}

	return defaultValue
}

func (i *Input) GetUint(key string, defaultValue uint64) uint64{
  value, ok := i.Params[key]
  var ret uint64

  if ok == false{
    return defaultValue
  }
  // integer from json unmarshal will turn to float64 default
  if float_64, ok := value.(float64); ok{
    return uint64(float_64)
  }

  if ret, ok = value.(uint64); ok{
    return ret
  }

  if str, ok := value.(string);ok{
    ret, err := strconv.ParseUint(str, 10, 64)
    if err == nil{
      return ret
    }
  }
  return defaultValue
}

func (i *Input) GetStrings(key string, defaultValue ...string) []string{
	value, ok := i.Params[key]
	if ok == false{
		return defaultValue
	}

  ret, ok := value.([]string)
  if ok{
    return ret
  }

	return defaultValue
}