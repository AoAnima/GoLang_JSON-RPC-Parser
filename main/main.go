package main

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
)

type JsqonQlResp struct {
	Result interface{}
	Error  interface{} `json:"error"`
	Id string
}
type JsqonQlReq struct {
	Method string `json:"method"`
	Params []map[string]interface{}
	Id string
}


func (str *JsqonQlReq) Get_data(arg  interface{}) (interface{}, error){


	err := errors.New("какаято ошибка")

	return arg, err
}

func parseJsonQl(JsonQl []byte) JsqonQlResp{
	var jsonData JsqonQlReq
	json.Unmarshal(JsonQl, &jsonData)
	arg := make([]reflect.Value,1)
	arg[0] = reflect.ValueOf(jsonData.Params)
	var t JsqonQlReq

	resultCall := reflect.ValueOf(&t).MethodByName(jsonData.Method).Call(arg)

	response := JsqonQlResp{}
	for key, val := range resultCall{

		if	val.Type().String() == "error"{
			response.Error = resultCall[key].Interface()
		} else {
			response.Result = resultCall[key].Interface()
		}
	}
	response.Id = jsonData.Id
	log.Printf("92 %+v\n", response)

	return response
}
