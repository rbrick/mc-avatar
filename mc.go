package main

import (
  "image"
  "image/png"
  "net/http"
  "fmt"
  "os"
  "log"
  "image/draw"
  "github.com/disintegration/imaging"
)

type Skin struct {
  Img image.Image // The image retrieved from the URL
  Name string
}

var mcurl string = "http://skins.minecraft.net/MinecraftSkins/%s.png"

/**
Returns a Skin.
*/
func GetSkin(userName string) (skin *Skin, err error) {
  url := fmt.Sprintf(mcurl, userName)
  resp, err := http.Get(url) // Gets the data from the url

  if err != nil {
    return
  }


  img, err1 := png.Decode(resp.Body)

  if err1 != nil {
    return
  }

  return &Skin{Img: img, Name: userName}, nil
}


// Size -> ratio so 1:1 is 8x8 2:1 16x16
func (skin *Skin) GetFace(size int) image.Image {
  m := image.NewRGBA(image.Rect(8,8,16,16))

  draw.Draw(m, m.Bounds(), skin.Img, image.Point{0,0}, draw.Src)

  for minX := 0; minX < m.Bounds().Max.X; minX++ {
    for minY := 0; minY < m.Bounds().Max.Y; minY++ {
      m.Set(minX, minY, skin.Img.At(minX, minY))
    }
  }

  rimg := imaging.Resize(m, size, 0, imaging.NearestNeighbor)

  newImg, _ := os.Create(skin.Name + "_face.png")

  defer newImg.Close()



  err := png.Encode(newImg, rimg)

  if err != nil {
    log.Fatal(err)
  }
  return m
}
