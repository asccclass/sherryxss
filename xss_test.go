package SherryXSS

import (
   "testing"
   "fmt"
   // "reflect"
)

func TestXSS(t *testing.T) {
   want := XX{IntField:42, StringField:"&lt;script&gt;alert(&#39;foo&#39;);&lt;/script&gt;"}
   test := XX{IntField:42, StringField: "<script>alert(\"foo\");</script>"}

   xs := XSS{}
   s, err := xs.HtmlEscape(test)

   if err != nil {
      fmt.Printf("Error: %v \n", err)
   } 

   vals := s.(XX)
   
   if vals != want {
      t.Error("Something error...")
   }

   fmt.Printf("Wnt:%v\nGot:%v", want, vals)
/*

   if vals.Field(0).Interface().(int) != want.IntField {
      fmt.Printf("Return struct is :%v\n", s)
   }
 //|| val(1).Interfacce().Elem().(string) != want.StringField {
*/
}
