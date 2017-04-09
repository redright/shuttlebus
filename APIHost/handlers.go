package APIHost

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"io"
	"io/ioutil"

	"github.com/redright/shuttlebus/appServices"
)

func Operation(w http.ResponseWriter, r *http.Request) {

	var operation OperationRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &operation); err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	services := appServices.Services{}
	servicesV := reflect.ValueOf(&services).Elem()
	serviceT, _ := servicesV.Type().FieldByName(operation.ServiceName)
	if serviceT.Type == nil {
		w.WriteHeader(404) // unprocessable entity
		return
	}
	method, _ := serviceT.Type.MethodByName(operation.MethodName)
	methodT := method.Type
	parameterCount := methodT.NumIn() - 1

	if parameterCount != len(operation.Parameters) {
		panic("ParameterCountMissmatch")
	}
	refParams := make([]reflect.Value, parameterCount+1)
	refParams[0] = reflect.New(serviceT.Type.Elem())
	for i := 1; i < methodT.NumIn(); i++ {
		paramType := methodT.In(i)
		t := reflect.New(paramType)
		if err := json.Unmarshal([]byte(operation.Parameters[i-1]), t.Interface()); err != nil {
			fmt.Println(err)
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
			return
		}
		refParams[i] = reflect.ValueOf(t.Elem().Interface())
	}
	// serviceInitMethod, _ := serviceT.Type.MethodByName("Init")
	// serviceInitMethod.Func.Call([]reflect.Value{refParams[0]})
	result := method.Func.Call(refParams)

	response := OperationResponse{}
	// if err != nil {
	// 	response.Error = string(err)
	// }

	if len(result) > 0 {
		errorInterface := reflect.TypeOf((*error)(nil)).Elem()
		if len(result) > 0 && !result[0].Type().Implements(errorInterface) {
			response.Result = result[0].Interface()
		} else {
			response.Error = result[0].Interface().(error).Error()
		}

		if len(result)-1 > 0 {
			err := result[len(result)-1].Interface()
			if err != nil {
				response.Error = err.(error).Error()
			}
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
