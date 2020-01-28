package main

import "fmt"

func main() {
	bp := NewBannerPrinter("Hello")
	bp.Weak()
	bp.Strong()
}

type Banner struct {
	text string
}

func (b *Banner) Paren() {
	fmt.Println("(" + b.text + ")")
}
func (b *Banner) Aster() {
	fmt.Println("*" + b.text + "*")
}

type Printer interface {
	Weak()
	Strong()
}

type BannerPrinter struct {
	banner *Banner
}

func NewBannerPrinter(text string) Printer {
	return &BannerPrinter{
		banner: &Banner{
			text: text,
		},
	}
}

func (bp *BannerPrinter) Weak() {
	bp.banner.Paren()
}

func (bp *BannerPrinter) Strong() {
	bp.banner.Aster()
}
