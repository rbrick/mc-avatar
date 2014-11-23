package main

import (
  "log"
  "fmt"
  "net/http"
  "github.com/rbrick/mc-avatar"
)

func main() {
   //
   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Path[len("/"):]
        if name == "" {
          fmt.Fprintln(w, "name must not be empty!")
        } else if len(name) > 16 {
          fmt.Fprintln(w, "name to long!")
        } else {
           skin, err := mc.GetSkin(name)

           if err != nil {
              log.Fatal(err)
           }

          img := skin.GetFace(64)

           if img == nil {
             fmt.Fprintln(w, "image not found!")
           }

          http.ServeFile(w,r, name + "_face.png")
        }
   })

   http.ListenAndServe(":80", nil)
}
