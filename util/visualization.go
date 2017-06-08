package util

import (
	"os/exec"
)

func Visualization(outfile string) {

	// println("dot", "-Tpng", outfile, "-o", outfile+".png")
	toPng := exec.Command("dot", "-Tpng", outfile, "-o", outfile+".png")
	toPng.CombinedOutput()
	// out, err := toPng.CombinedOutput()
	// if err != nil {
	// 	println("err:", err.Error())
	// }
	// fmt.Println(string(out))

	// println("open", outfile+".png")
	openPNG := exec.Command("open", outfile+".png")
	openPNG.CombinedOutput()
	// out, err = openPNG.CombinedOutput()
	// if err != nil {
	// 	println("err:", err.Error())
	// }
	// fmt.Println(string(out))
}
