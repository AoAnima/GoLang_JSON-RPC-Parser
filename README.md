# GoLang_JSON-RPC-Parser
Мини парсер JSON-RPC для Вызова удалённых процедур

Добавляем новые функции к структуре 
~~~~
func (str *JsqonQlReq) ИмяФункции(arg  interface{}) (interface{}, error){
  // код функции который нужно выполнить
	err := errors.New("какаято ошибка")

	return result, err
}
~~~~
в json должен приходить объект в формате json-rpc
~~~~
 data = []byte({"method": "ИмяФункции", "params": [{"name":"vale"},...{"name":"vale"}], "id": 1})
 
 result := ParseJsonQl(data)

result будет в формате 
result = {
	Result interface{} //"резудтат выполнения функции"
	Error  interface{} `json:"error"` // ваша ошибка или nil 
	Id string // id который пришёл в изначальном json объекте 
}

~~~~
