/*
   func: XSS Protect for web params
   By: Chih-Han LIu
   Date:2019-04-24
   Lastupdate:
*/
package main // SherryXSS

import (
   "fmt"
   "html"
   "reflect"
)

type XSS struct {
}

// Transfer html params to string, ex: html space to &nbsp;
// Note: Need transfer struct to interface....!!!?
func (xss *XSS) HtmlEscape(vals interface{}) (interface{}, error) {
   elems := reflect.TypeOf(vals)
   if elems.Kind() != reflect.Struct {
      return nil, fmt.Errorf("Type of param must be struct, but got %v", elems.Kind())
   }
   elemsValue := reflect.ValueOf(vals)
   newStruct := reflect.New(elems)
   for i := 0; i < elems.NumField(); i++ {
      elem := elems.Field(i)
      val := elemsValue.Field(i).Interface()
      switch elem.Type.Name() {
          case "string":
             s := html.EscapeString(val.(string))  // 型態強制轉化 2 string
             newStruct.Elem().Field(i).SetString(s)
         default:
             t := elem.Type
             s := val.(t)
fmt.Printf("%v = %v\n", s, t)
/*
            val := reflect.ValueOf(vals).Field(i) // .Interface().(x)  // 取得欄位內容
            x := reflect.ValueOf(val)   // <int Value>
            fmt.Printf("%v\n", x)
            // newStruct.Elem().Field(i).Set(x)
*/
      }
   }
   return newStruct.Interface(), nil
}

type XX struct {
        IntField int
        StringField string
}

func main() {
   myStruct := XX {
      IntField: 42,
      StringField: "<script>alert('foo');</script>",
    }

   xs := XSS{}
   s, err := xs.HtmlEscape(myStruct)

   if err != nil {
      fmt.Printf("Error:%v\n", err)
   } else {
      fmt.Printf("Return struct is :%v\n", s)
   }
}
