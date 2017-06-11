package util

import (
	"fmt"
	"os"
	"os/exec"
)

func Visualization(outfile string) {
	if CheckFileExist(outfile + ".png") {
		os.Remove(outfile + ".png")
	}
	// println("dot", "-Tpng", outfile, "-o", outfile+".png")
	toPng := exec.Command("dot", "-Tpng", outfile, "-o", outfile+".png")
	// toPng.CombinedOutput()
	out, err := toPng.CombinedOutput()
	if err != nil {
		println("err:", err.Error())
	}
	fmt.Println(string(out))

	// println("open", outfile+".png")
	openPNG := exec.Command("open", outfile+".png")
	// openPNG.CombinedOutput()
	out, err = openPNG.CombinedOutput()
	if err != nil {
		println("err:", err.Error())
	}
	fmt.Println(string(out))
}
