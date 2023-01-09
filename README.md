<H1> Learn Go Lang</H1>
<h2>Basic Commands</h2>  

```
# Get List of all Variable
go env

# Go help command
go help

# Package your code
go mod init <module>

# Get
go get -d <url> #-d for just downloading and not building

# Format the Go code
go fmt or go fmt ./...  #(... for nested files)

# Build an executable file
go build

#
```

1. GoPath contains workspace path    
2. GoRoot contains binary execution files of go
3. Directory Structure of go is   
       workspace_dir ->  
       | src  
       | bin  
       | pkg
4. bin contains all the binary build files
5. src contains the code
6. pkg contains all the packages
