package makeDir

import (
	"fmt"
	"os"
	_ "os/exec"
	"path/filepath"
)

func detectNumLev() uint8 { //Detect numbers levels to ~(home)
	locationDir, err := filepath.Abs(filepath.Dir(os.Args[0])) //Save location of directory
	var numLev uint8 = 0                                       //save Number levels
	//	fmt.Printf("\n%s\n", locationDir)
	if err == nil {
	}
	for i := 1; i < len(locationDir); i++ {
		if locationDir[i] == filepath.Separator { //If locationDir[i] is equal to "/"(or "\"), then increase the numLev
			numLev++
		}
	}
	return numLev //Return number od levels
}

func MakeDir(numEpisodes, series int, path string) {
	var pathDir string
	numLev := detectNumLev() //Save value of detectNumLev()

	for i := 0; i < int(numLev)-1; i++ { //Add "../" to end of pathDir
		pathDir += fmt.Sprintf("..%c", filepath.Separator)
	}
	pathDir += path[2:] //Add path[2:] to end of pathDir
	fmt.Printf("path: %s\n", pathDir)
	var err error
	for i := 1; i <= numEpisodes; i++ {
		err = os.Mkdir(fmt.Sprintf("%s%cS%02dE%02d", pathDir, filepath.Separator, series, i), 0777)
	}

	fmt.Println(err)
}
