/**
 * @Author: Vitali Saroka
 * @Description: Hash calc command line tool
 * @Date: 2022/3/29 16:10
 */
package main
 
import (
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"crypto/sha256"
	"io"
	"sync"
)
 
func main() {
	var mode string
	var folder string
 
	flag.StringVar (& mode, "m", "", "Mode of hash calculation, possible options are single and parallel.")
	flag.StringVar (& folder, "f", "", "Path to folder with files.")
	flag.Parse()
	
	if folder == "" {
        fmt.Println("There is no path to folder defined, please use -f parameter and folder name, e.g. -f myfolder")
        os.Exit(1)
    }
    
    files := getOnlyFiles(folder)
    
	if len(files) < 1 {
        fmt.Println("No files found! Exiting!")
        os.Exit(1)
    }
	
	fmt.Println("Starting md5sum calculations ...")
	
	if mode == "parallel" {
		fmt.Println("Calculating in parallel mode!")
		shaResponses := make(chan string)
		var wg sync.WaitGroup
		wg.Add(len(files))
		
		for _, file := range files {
			path := getPath(folder,file.Name())
			go func(path string) {
				defer wg.Done()
				shaResponses <- path + getSHA256(path)
			}(path)
		}
		go func() {
				for response := range shaResponses {
					fmt.Println(response)
				}
			}()

		wg.Wait()
	} else {
		fmt.Println("Calculating in sequential mode!")
		fmt.Println("If you need paralle mode please use -m parallel")
		for _, file := range files {
			path := getPath(folder,file.Name())
			calcAndPrintHash(path)
		}
	} 
	fmt.Scanln() 
}

func calcAndPrintHash (path string) {
	checksum := getSHA256(path)
	fmt.Println(path, checksum)
}

func getPath(folder string, file string) string {
	return folder+"/"+file
}

func getOnlyFiles(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	onlyFiles := files[:0]
	for _, file := range files {
		if file.IsDir() == false {
		    onlyFiles = append(onlyFiles, file)
		}	
	}
	return onlyFiles;
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

