package main

import (
  "log"
  "net/http"
  "github.com/go-martini/martini"
  "strconv"
  "fmt"
)

func main() {
    m := martini.Classic()

    m.Get("/:player", func(w http.ResponseWriter,r *http.Request,params martini.Params) {
      name := params["player"]
      if name == "" {
        fmt.Fprintln(w,"name must not be empty!")
        } else if len(name) > 16 {
          fmt.Fprintln(w,"name to long!")
          } else {
            skin, err := GetSkin(name)

            if err != nil {
              log.Fatal(err)
            }

            img := skin.GetFace(64)

            if img == nil {
              fmt.Fprintln(w,"image not found!")
            }

            http.ServeFile(w,r, name + "_face.png")
          }
    })

    m.Get("/:size/:player", func(w http.ResponseWriter,r *http.Request,params martini.Params) {

     size, err := strconv.Atoi(params["size"])

        // something went wrong with the request
      if err != nil {
         size = 64
      }

      if size > 512 || size <= 10 {
        size = 64
      }

      name := params["player"]
      if name == "" {
        fmt.Fprintln(w, "name must not be empty!")
        } else if len(name) > 16 {
          fmt.Fprintln(w, "name to long!")
          } else {
            skin, err := GetSkin(name)

            if err != nil {
              log.Fatal(err)
            }

            img := skin.GetFace(size)

            if img == nil {
              fmt.Fprintln(w,"image not found!")
            }

            http.ServeFile(w,r, name + "_face.png")
          }
      })

    http.ListenAndServe(":80", m)
}
