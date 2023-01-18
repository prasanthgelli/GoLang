<H1> Learn Go Lang</H1>
<H2>Documentation Links</H2>  

1. Modules/Package Search [here](https://pkg.go.dev/)
2. GoLan Blog [here](https://go.dev/blog/)
3. Editor [here](https://go.dev/play/)
4. Documentation [here](https://go.dev/ref/spec)
<H2>Basic Commands Links</H2>

```
# Get List of all Variable
go env

# Go help command
go help

# Package your code
go mod init <module>

# Get
go get -d <module> #-d for just downloading and not building

# Format the Go code
go fmt or go fmt ./...  #(... for nested files)

# Build an executable file
go build

# Test any Packages
go test # This add a new file called go.sum which maintains all the hash and versions of project

# Install or Remove all modules
go tidy

# Get all the modules/dependencies
go list -m all

#get a specific version of module
go get <module_name>@<version>

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
7. go.sum has the version and hash management.
8. we can't have unused variables
9. <b>_</b> is used as wild operator which can be just assigned and need not be used.
10. To assign value to a new variable we need to use  :=, but if we want to reassign a value we need to use = ,
11. To get the architecture and other configuration of machine we can use runtime package which has goos and gooarch.
12.Strings are immutable  
13. Everything in go is pass by value.
14. For variadic parameters it either 0 to infinite, we can pass any number of values
15. To find the address of a variable where it is stored we use the <b>$</b> operator.
16. To convert a string to json the identifiers inside the struct must be started with capital letters.
