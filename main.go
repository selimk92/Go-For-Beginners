package main

import "fmt"

type insan struct {
	isim string
	yas  int
}

func main() {
	kisi := insan{"selim", 25}
	kisi.tanıt()
}

func (i insan) tanıt() {
	fmt.Println("Merhaba benim adım ", i.isim, "ve", i.yas, "yasındayım")

}
