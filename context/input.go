package context


import(
  "net/http"
)

type Input struct {
  Params map[string][]string
  Request *http.Request
}

func (i *Input) GetString(key string, defaultValue string) string{
	value, ok := i.Params[key]
	if ok{
		return value[0]
	}

	return defaultValue
}

func (i *Input) GetStrings(key string, defaultValue ...string) []string{
	value, ok := i.Params[key]
	if ok{
		return value
	}

	return defaultValue
}