package dependencies

import "github.com/fatih/color"
//TODO: Download
//TODO: Use the external package to print out the s string in BgHiRed color
func ColorPrint(s string) {
	color.New(color.BgHiRed).Println(s)
}
