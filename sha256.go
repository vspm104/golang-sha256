/**
 * @Author: Vitali Saroka
 * @Description: Hash calc command line tool
 * @Date: 2022/3/28 11:20
 */
package main
 
import (
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"log"
	"crypto/sha256"
	"io"
)
 
func main() {
	var mode string
	var folder string
 
	flag.StringVar (& mode, "m", "", "Mode of hash calculation, possible options are single and parallel.")
	flag.StringVar (& folder, "f", "", "Path to folder with files.")
	flag.Parse()
	
	if folder == "" {
        fmt.Println("There is no folder defined, please use -f parameter and folder name, e.g. -f myfolder")
        os.Exit(1)
    }
    
    files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Starting md5sum calculations ...")
	
	if mode == "" {
		fmt.Println("Calculating in sequential mode!")
		fmt.Println("If you need paralle mode please use -m parallel")
	}
	
	for _, file := range files {
		path := folder+"/"+file.Name()
		if file.IsDir() == false {
		    if mode == "parallel" {
				go CalcAndPrintHash(path)
		    } else {
				CalcAndPrintHash(path)
		    }  
		}	
	}

    fmt.Scanln() 
}

func CalcAndPrintHash (path string) {
	checksum := getSHA256(path)
	fmt.Println(path, checksum)
} 

func getSHA256(path string) string {
    hash := sha256.New()
	file, error := os.Open(path)
	if error != nil {
		panic(error)
	}
	defer file.Close()

	if _, error := io.Copy(hash, file); error != nil {
		panic(error)
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}

