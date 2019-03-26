package handlers

import (
	"os"
	"strings"
	"encoding/base64"
)

func SaveFile(name, data string) string {
	if data == "" {
		return "/public/images/avatar.png"
	}
	if _, err := os.Stat("." + data); err == nil {
		return data
	} else {
		b64data := data[strings.IndexByte(data, ',')+1:]
		dec, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			panic(err)
		}
		f, err := os.Create("./public/uploads/" + name)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err := f.Write(dec); err != nil {
			panic(err)
		}
		if err := f.Sync(); err != nil {
			panic(err)
		}
		return "/public/uploads/" + name
	}
}
