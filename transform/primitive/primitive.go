package primitive

import (
	"fmt"
	"gophercises/transform/filehandle"
	"os"
	"os/exec"
	"strings"

	"github.com/atrox/homedir"
)

var cmdExecuteFunc = exec.Command

//TransformInfo is a struncture to store the file pointer and an error when an output file is created
type TransformInfo struct {
	File *os.File
	Mode int
	Err  error
}

func transformRoutine(ch <-chan int, file chan<- TransformInfo, filepath string, mode int, number int, ext string) {
	defer func() {
		<-ch
	}()
	var fileData TransformInfo
	fileData.File, fileData.Err = Transform(filepath, mode, number, ext)
	fileData.Mode = mode
	file <- fileData
}

//Transform function would take in a source file path and perform primtive transformatoin operation on the image and return a file pointer and an error if one occurs
func Transform(filepath string, mode int, number int, ext string) (*os.File, error) {
	hDir, _ := homedir.Dir()
	dst, err := filehandle.TempFile(hDir+"/go/src/gophercises/transform/output", ext)
	if err == nil {
		cmdDir := hDir + "/go/bin/primitive"
		cmd := cmdExecuteFunc(cmdDir, strings.Split(fmt.Sprintf("-i %s -o %s -n %d -m %d", filepath, dst.Name(), number, mode), " ")...)
		_, err = cmd.CombinedOutput()
		fmt.Println(strings.Split(fmt.Sprintf("-i %s -o %s -n %d -m %d", filepath, dst.Name(), number, mode), " "))
	}
	return dst, err
}

//BatchTransform function is to perform multiple transformations. it returns slice of file pointers to the output image file and an error if one occurs
func BatchTransform(filepath string, mode []int, number []int, ext string) []TransformInfo {
	var fileInfoSlice []TransformInfo
	var fileData TransformInfo
	channelLen := len(mode) * len(number)
	ch := make(chan int, channelLen)
	file := make(chan TransformInfo, channelLen)
	for _, m := range mode {
		for _, num := range number {
			ch <- 1
			go transformRoutine(ch, file, filepath, m, num, ext)
		}
	}
	for len(ch) != 0 {
	}
	exit := 0
	for {
		select {
		case fileData = <-file:
			fileInfoSlice = append(fileInfoSlice, fileData)
		default:
			exit = 1
		}
		if exit == 1 {
			break
		}
	}
	return fileInfoSlice
}
