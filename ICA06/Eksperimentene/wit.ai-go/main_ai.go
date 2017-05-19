package main

import (
    "fmt"
    witai "github.com/meinside/wit.ai-go"
)
const (
    Token = "553XNQUSZEO7E3XXVMOSYMDW55YKX75R"
)

func main(){

c := witai.NewClient(Token)

c.Verbose = true

if result, err := c.QuerySpeechMp3("test3.mp3", nil, "", "", 1); err == nil {
  fmt.Printf("query speech result: %+v\n", result)
} else {
  fmt.Printf("%s\n", err)
}
}
