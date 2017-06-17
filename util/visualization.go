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
	toPng := exec.Command("dot", "-Tpng", outfile, "-o", outfile+".png")
	// toPng.CombinedOutput()
	out, err := toPng.CombinedOutput()
	if err != nil {
		fmt.Println(string(out) + "err: " + err.Error())
	}
	Debug(string(out))

	openPNG := exec.Command("open", outfile+".png")
	// openPNG.CombinedOutput()
	out, err = openPNG.CombinedOutput()
	if err != nil {
		fmt.Println(string(out) + "err: " + err.Error())
	}
	Debug(string(out))

	os.Remove(outfile)
}
