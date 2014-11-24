package main

import (
  "image"
  "image/png"
  "net/http"
  "fmt"
  "image/draw"
  "errors"
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
    return nil, err
  }

  if resp.Header.Get("Content-Type") != "image/png" || resp.StatusCode != 200 {
     return nil, errors.New("An internal error has occurred.")
  }

  img, err1 := png.Decode(resp.Body)

  if err1 != nil {
    return nil, err
  }

  return &Skin{Img: img, Name: userName}, nil
}


// Size -> ratio so 1:1 is 8x8 2:1 16x16
func (skin *Skin) GetFace(size int) image.Image {
  if skin.Img == nil {
    return nil
  }

  m := image.NewRGBA(image.Rect(8,8,16,16))

  draw.Draw(m, m.Bounds(), skin.Img, image.Point{0,0}, draw.Src)

  for minX := 0; minX < m.Bounds().Max.X; minX++ {
    for minY := 0; minY < m.Bounds().Max.Y; minY++ {
      m.Set(minX, minY, skin.Img.At(minX, minY))
    }
  }

  rimg := imaging.Resize(m, size, 0, imaging.NearestNeighbor)

  return rimg
}
