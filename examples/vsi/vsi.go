package main

import (
	"fmt"
	"log"

	"github.com/dewberry/gdal"
)

func main() {
	driver, err := gdal.GetDriverByName("GTiff")
	if err != nil {
		log.Fatal(err)
	}

	ds := driver.Create("/vsimem/tmp.tif", 256, 256, 1, gdal.Byte, []string{})
	ds.Close()

	fmt.Println("Globbing all vsimem files:")
	fmt.Println(gdal.VSIReadDirRecursive("/vsimem/"))

	fmt.Println("Adding a second vsimem file...")
	ds2 := driver.Create("/vsimem/tmp2.tif", 256, 256, 1, gdal.Byte, []string{})
	ds2.Close()

	fmt.Println("Globbing all vsimem files:")
	fmt.Println(gdal.VSIReadDirRecursive("/vsimem/"))

	fmt.Println("Unlinking the first vsimem file...")
	err = gdal.VSIUnlink("/vsimem/tmp.tif")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Globbing all vsimem files again:")
	fmt.Println(gdal.VSIReadDirRecursive("/vsimem/"))
}
