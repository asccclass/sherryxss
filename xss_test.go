package SherryXSS

import (
   "testing"
)


func TestXSS() {
   var want string = `struct { IntField int; StringField string }{IntField:42, StringField:"&lt;script&gt;alert(&#39;foo&#39;);&lt;/script&gt;"}`

   xs := XSS{IntField:42, Intfield: "<script>alert("foo");</script>"}}
   s, err := xs.HtmlEscape(myStruct)

   if err != nil {
      fmt.Printf("Error:%v\n", err)
   } else if xs["IntField"] != want["IntField"] || xs["StringField"] != want["StringField"]) {
      fmt.Printf("Return struct is :%v\n", s)
   }
}
_
