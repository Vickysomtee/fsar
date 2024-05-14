package fsar

import (
	"os"

	"github.com/fatih/color"
)

func MountFsar(dir string) {

	_, err := os.Create("./fsar.conf")

	if err != nil {
		panic("Could not mount Fsar")
	}

	if err := os.Mkdir(dir, 0777); err != nil {
		color.Red("Error mounting direction:", err)
		return
	}

	color.Green("Fsar Directory mounted successfully")
}

func UseBucket(bucket string, dir string) {
	files, err := os.ReadDir(dir)

	if err != nil {
		color.Red("Could not access Fsar directory")
		return
	}

	for _, file := range files {
		if file.Name() != bucket {
			color.Red("Bucket does not exit. Try creating the Bucket first")
		}
	}
}

func CreateBucket(name string, dir string) {
	file, err := os.Create(dir + name)

	if err != nil {
		color.Red("Error creating bucket:", err)
		return
	}

	color.Green("Bucket %s created successfully", file)
}
