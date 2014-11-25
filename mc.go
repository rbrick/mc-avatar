package main

import (
  "image"
  "image/png"
  "net/http"
  "fmt"
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

  if len(userName) == 1 {
    url = "https://minecraft.net/images/steve.png"
    resp1, _ := http.Get(url)

    img,_ := png.Decode(resp1.Body)

    return &Skin{Img: img, Name: "Steve"}, nil
  }

  resp, err := http.Get(url) // Gets the data from the url



  if err != nil {
     url = "https://minecraft.net/images/steve.png"
     resp1, _ := http.Get(url)

     img,_ := png.Decode(resp1.Body)

    return &Skin{Img: img, Name: "Steve"}, nil
  }

  if resp.Header.Get("Content-Type") != "image/png" || resp.StatusCode != 200 {
    url = "https://minecraft.net/images/steve.png"
    resp1, _ := http.Get(url)

    img,_ := png.Decode(resp1.Body)

    return &Skin{Img: img, Name: "Steve"}, nil
  }

  img, err1 := png.Decode(resp.Body)

  if err1 != nil {
    return nil, err
  }

  return &Skin{Img: img, Name: userName}, nil
}

func (skin *Skin) GetHelm(size int) image.Image {
   sk := image.NewRGBA(image.Rect(8,8,16,16))
   hat := image.NewRGBA(getRectangle(40,8,8,8))

  for minXh := hat.Bounds().Min.X; minXh < hat.Bounds().Max.X; minXh++ {
    for minYh := hat.Bounds().Min.Y; minYh < hat.Bounds().Max.Y; minYh++ {
      hat.Set(minXh, minYh, skin.Img.At(minXh, minYh))
    }
  }

   draw.Draw(sk, sk.Bounds(), skin.Img, image.Point{0,0}, draw.Src)

   for minX := 0; minX < sk.Bounds().Max.X; minX++ {
     for minY := 0; minY < sk.Bounds().Max.Y; minY++ {
        sk.Set(minX, minY, skin.Img.At(minX, minY))
     }
   }

  if !IsSolidColor(hat) {
    draw.Draw(sk, sk.Bounds(), hat, hat.Bounds().Min, draw.Over)
  }

  rimg := imaging.Resize(sk, size, 0, imaging.NearestNeighbor)

  return rimg
}



func (skin *Skin) GetFace(size int) image.Image {

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

/**
  The following methods are from LapisBlue's project Lapitar.
  Credits to @LapisBlue & @minecrell

  I take no credit for these methods, nor do I claim they are my own. I am only
  using these methods to make my life easier. Again I take no credit. If you would
  like me to take this down I will.
*/
func IsSolidColor(img image.Image) bool {
  base := img.At(img.Bounds().Min.X, img.Bounds().Min.Y)
  for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
    for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
      if img.At(x, y) != base {
        return false
      }
    }
  }
  return true
}

func getRectangle(x,y,width, height int) image.Rectangle {
  return image.Rect(x,y, x+width, y+height)
}
