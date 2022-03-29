# golang-sha256
This is a command line tool which calculates sha256 sum for all files from a given folder.  
There are two variants:  
I. With waitgroup - run example_with_wait_group.go file  
Examples of usage:  
1) Sequential mode: go run example_with_wait_group.go -f myfolder  
2) Parallel mode: go run example_with_wait_group.go -f myfolder -m parallel  
II. Without waitgroup - run sha256.go file  
Examples of usage:  
1) Sequential mode: go run sha256.go -f myfolder  
2) Parallel mode: go run sha256.go -f myfolder -m parallel  

