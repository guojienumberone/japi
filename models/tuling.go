package models

import (
	 "net/http"
	 "encoding/json"
	 "io/ioutil"
	 "fmt"
)

type TuLing struct {
	Code int
	Text string
}

func (tl *TuLing) Translation(chat string) {
	key := "fcfea588759708b454baa23fe8f17b80"
	response,_:=http.Get("http://www.tuling123.com/openapi/api?key="+key+"&info="+chat)
	defer response.Body.Close()
	
	if response.StatusCode == 200 {
		body,_:=ioutil.ReadAll(response.Body)
		
		err:=json.Unmarshal([]byte(body), &tl)
		if err==nil {
			fmt.Println(tl)
		} else {
			fmt.Println(err)
		}
	}
}

