package context


import(
  "net/http"
)

type Input struct {
  Params map[string][]string
  Request *http.Request
}