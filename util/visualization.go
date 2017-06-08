package util

import (
	"fmt"
	"os/exec"
)

func Visualization(outfile string) {

	println("dot", "-Tpng", outfile, "-o", outfile+".png")
	toPng := exec.Command("dot", "-Tpng", outfile, "-o", outfile+".png")
	out, err := toPng.CombinedOutput()
	if err != nil {
		println("err:", err.Error())
	}
	fmt.Println(string(out))

	println("open", outfile+".png")
	openPNG := exec.Command("open", outfile+".png")
	out, err = openPNG.CombinedOutput()
	if err != nil {
		println("err:", err.Error())
	}
	fmt.Println(string(out))
}
