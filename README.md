## buld project
> go build main.go

## run built project
> ./main.go

## run project without build
> go run main.go

## datetypes
int
int8
int16
int32
int64

### positive int
uint
uint8
uint16
uint32
uint64

## float
float32
float64

## text
string = ""

## bool
bool = true, false

## complex
complex64
complex128

## go packages docs
https://pkg.go.dev

## arrays and slices
array are immutable
slices are mutable

## install packages
> go get <depname>

## search packages
https://awesome-go.com

## run echo server
> go run echoserver.go

## replace a package for a local version
> go mod edit -replace github.com/labstack/echo/v4=./echo

## verify modules
> go mod verify

## remove replacements on go.mod
go mod edit -drop replace

## remove unused dependencies
go mod tidy
