package util

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

// ConvertToGIF converts a folder of pngs into an animated gif. Requires imagemagick convert.
func ConvertToGIF(folder, outFileName string, fps float64) {
	delay := fmt.Sprintf("%f", 1000.0/fps/10.0)
	path := folder + "/*.png"
	cmd := exec.Command("convert", "-delay", delay, path, outFileName)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// ViewImage displays an image using installed image viewer.
func ViewImage(imagePath string) {
	cmd := exec.Command("eog", imagePath)
	if runtime.GOOS == "darwin" {
		cmd = exec.Command("qlmanage", "-p", imagePath)
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
