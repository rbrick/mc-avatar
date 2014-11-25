package main

import (
  "log"
  "net/http"
  "github.com/go-martini/martini"
  "strconv"
  "fmt"
  "image/png"
)

func main() {
    m := martini.Classic()

    m.Get("/", func() string {
        return "Hello World!\nappend /face/<playername> or /face/<size>/<playername>\nto get the face you want!"
      })

    m.Get("/face/:player", func(w http.ResponseWriter,r *http.Request,params martini.Params) {
      name := params["player"]
      if name == "" {
         fmt.Fprintln(w,"name must not be empty!")
         //name = "steve"
        } else if len(name) > 16 {
           fmt.Fprintln(w,"name to long!")
//          name = "steve"
          } else {
            skin, err := GetSkin(name)

            if err != nil {
              fmt.Fprintln(w,"An internal error has occurred!")
              return
            }

            img := skin.GetFace(128)

            if img == nil {
              fmt.Fprintln(w,"image not found!")
              return
            }

            w.Header().Set("Content-Type", "image/jpeg")

            err1 := png.Encode(w, img)

            if err1 != nil {
              log.Fatal(err1)
            }


          }
    })

    m.Get("/helm/:player", func(w http.ResponseWriter,r *http.Request,params martini.Params) {
      name := params["player"]
      if name == "" {
        fmt.Fprintln(w,"name must not be empty!")
        //name = "steve"
        } else if len(name) > 16 {
          fmt.Fprintln(w,"name to long!")
          //          name = "steve"
          } else {
            skin, err := GetSkin(name)

            if err != nil {
              fmt.Fprintln(w,"An internal error has occurred!")
              return
            }

            img := skin.GetHelm(128)

            if img == nil {
              fmt.Fprintln(w,"image not found!")
              return
            }

            w.Header().Set("Content-Type", "image/jpeg")

            err1 := png.Encode(w, img)

            if err1 != nil {
              log.Fatal(err1)
            }


          }
          })

    m.Get("/face/:size/:player", func(w http.ResponseWriter,r *http.Request,params martini.Params) {

     size, err := strconv.Atoi(params["size"])

        // something went wrong with the request
      if err != nil {
         size = 128
      }

      if size > 512 || size <= 10 {
        size = 128
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

            w.Header().Set("Content-Type", "image/jpeg")

            err1 := png.Encode(w, img)

            if err1 != nil {
              log.Fatal(err1)
            }
          }
      })

      m.Get("/helm/:size/:player", func(w http.ResponseWriter,r *http.Request,params martini.Params) {

        size, err := strconv.Atoi(params["size"])

        // something went wrong with the request
        if err != nil {
          size = 128
        }

        if size > 512 || size <= 10 {
          size = 128
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

              img := skin.GetHelm(size)

              if img == nil {
                fmt.Fprintln(w,"image not found!")
              }

              w.Header().Set("Content-Type", "image/jpeg")

              err1 := png.Encode(w, img)

              if err1 != nil {
                log.Fatal(err1)
              }
            }
            })

    http.ListenAndServe(":80", m)
}
