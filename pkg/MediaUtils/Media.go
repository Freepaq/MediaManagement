package MediaUtils

import (

	"fmt"
	"github.com/Freepaq/MediaManagement/pkg/Setup"
	_ "github.com/rwcarlsen/goexif/exif"
	_ "github.com/rwcarlsen/goexif/tiff"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func fileStruct(files *[]string, origin string, eligibleFiles string) filepath.WalkFunc {
	fmt.Println("Start scanning folder :" + origin)
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() && origin != path {
			//		filepath.Walk(path, fileStruct(files, path))
		} else {
			if Setup.VIDEO == eligibleFiles {
				if Setup.IsVideoEligible(filepath.Ext(path)) {
					*files = append(*files, path)
				}
			}
			if Setup.PHOTO == eligibleFiles {
				if Setup.IsPhotoEligible(filepath.Ext(path)) {
					*files = append(*files, path)
				}
			}
			if Setup.ALL == eligibleFiles {
				if Setup.IsMediEligible(filepath.Ext(path)) {
					*files = append(*files, path)
				}
			}
		}
		return nil
	}
}
func GetListOfFile(folder string, eligibleFiles string) []string {
	var files []string
	err := filepath.Walk(folder, fileStruct(&files, folder, eligibleFiles))
	if err != nil {
		panic(err)
	}

	return files
}

func Delete(file Setup.FileStruct) {
	err := os.Remove(file.FullName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File [" + file.FullName + "] moved")
	}
}

func Rename(file *Setup.FileStruct) {
	var n = (*file).CreationDate.Format(Setup.TimestampFormat)
	(*file).NewName = n + (*file).Extension
	(*file).NewFullName = (*file).DestinationDir + (*file).NewName
	fmt.Println("New name : [" + (*file).NewName + "]")
}

func Copy(ori *Setup.FileStruct, destFoler string, force bool) bool {
	year, month, _ := (*ori).CreationDate.Date()
	dest := filepath.Join(destFoler, strconv.Itoa(year))
	destMonth := filepath.Join(destFoler, strconv.Itoa(year), Setup.Months[int(month)])
	destFull := filepath.Join(destMonth, (*ori).NewName)
	(*ori).DestinationDir = destMonth
	(*ori).NewFullName = filepath.Join((*ori).DestinationDir, (*ori).NewName)
	if _, err := os.Stat(destFoler); os.IsNotExist(err) {
		os.Mkdir(destFoler, 0755)
	}
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		os.Mkdir(dest, 0755)
	}
	if _, err := os.Stat(destMonth); os.IsNotExist(err) {
		os.Mkdir(destMonth, 0755)
	}
	input, err := ioutil.ReadFile((*ori).FullName)
	if err != nil {
		fmt.Println(err)
		return false
	}
	var result = true
	if !force {
		if _, err := os.Stat(destFull); os.IsNotExist(err) {
			if err, result = writeFile(destFull, input); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Destination file : [" + destFull + "] exists not override")
		}
	} else {
		if err, result = writeFile(destFull, input); err != nil {
			fmt.Println(err)
		}
	}
	input = nil
	return result
}

func writeFile(destFull string, input []byte) (error, bool) {
	err := ioutil.WriteFile(destFull, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destFull)
		fmt.Println(err)
		return nil, false
	}
	fmt.Println("Destination file : [" + destFull + "] copied")
	return err, true
}

func GetMeta(fname string) (Setup.FileStruct, error) {
	fileStr := Setup.FileStruct{}

	fileStr.FullName = fname
	fileStr.NewFullName = fname

	if Setup.IsVideoEligible(filepath.Ext(fname)) {
		ReadVideoMeta(fname, &fileStr)
	}
	if Setup.IsPhotoEligible(filepath.Ext(fname)) {
		ReadPhotoMeta(fname, &fileStr)
	}

	fileStr.OriginDir, fileStr.Name = filepath.Split(fname)
	fileStr.Extension = filepath.Ext(fname)
	fileStr.NewName = fileStr.Name
	return fileStr, nil
}
