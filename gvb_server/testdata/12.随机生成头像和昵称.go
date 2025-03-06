package main

import (
	"fmt"
	"image/png"
	"os"
	"path"
	"unicode/utf8"

	"github.com/DanPlayer/randomname"
	"github.com/disintegration/letteravatar"
	"github.com/golang/freetype"
)

func main() {
	// 生成随机名称
	// name := randomname.GenerateName()
	// fmt.Println(name)

	// 生成头像
	// names := []rune(name)
	// dir := "uploads/chat_avatar"
	// DrawImage(string(names[0]), dir)

	// 生成所有头像
	GenerateNameAvatar()
}
func GenerateNameAvatar() {
  dir := "uploads/chat_avatar"
  for _, s := range randomname.AdjectiveSlice {
    DrawImage(string([]rune(s)[0]), dir)
  }
  for _, s := range randomname.PersonSlice {
    DrawImage(string([]rune(s)[0]), dir)
  }
}

func DrawImage(name string, dir string) {
  fontFile, err := os.ReadFile("uploads/system/方正清刻本悦宋简体.TTF")
  font, err := freetype.ParseFont(fontFile)
  if err != nil {
    fmt.Println(err)
    return
  }
  options := &letteravatar.Options{
    Font: font,
  }
  // 绘制文字
  firstLetter, _ := utf8.DecodeRuneInString(name)
  img, err := letteravatar.Draw(140, firstLetter, options)
  if err != nil {
    fmt.Println(err)
    return
  }
  // 保存
  filePath := path.Join(dir, name+".png")
  file, err := os.Create(filePath)
  if err != nil {
    fmt.Println(err)
    return
  }
  err = png.Encode(file, img)
  if err != nil {
    fmt.Println(err)
    return
  }
}
