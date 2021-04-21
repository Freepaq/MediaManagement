package tests

import (
	"testing"
	"time"

	"github.com/Freepaq/MediaManagement/pkg/MediaUtils"
)

func TestRename(t *testing.T) {
	file := MediaUtils.FileStruct{}
	d := time.Now()
	file.CreationDate = d
	file.Extension = ".jpg"
	file.DestinationDir = "tmp/"
	MediaUtils.Rename(&file)
	if file.NewName != d.Format(MediaUtils.TimestampFormat)+file.Extension {
		t.Errorf("Rename file name failed " + file.NewName)
		t.Fail()
	}
	if file.NewFullName != file.DestinationDir+file.NewName {
		t.Errorf("Rename full file name failed " + file.NewName)
		t.Fail()
	}
}

func TestReadVideo(t *testing.T) {
	file := MediaUtils.FileStruct{}
	fname := "testfiles/videotest.mp4"
	err := MediaUtils.ReadVideoMeta(fname, &file)
	if err != nil {
		t.Fail()
		t.Errorf(err.Error())
	}
	if file.CreationDate.Format(MediaUtils.TimestampFormat) != "2017-07-07_101656" {
		t.Fail()
		t.Errorf(file.CreationDate.Format(MediaUtils.TimestampFormat))
	}
}

func TestReadPhoto(t *testing.T) {
	file := MediaUtils.FileStruct{}
	fname := "testfiles/phototest.jpg"
	MediaUtils.ReadPhotoMeta(fname, &file)

	if file.CreationDate.Format(MediaUtils.TimestampFormat) != "2021-04-17_171129" {
		t.Fail()
		t.Errorf(file.CreationDate.Format(MediaUtils.TimestampFormat))
	}
}
