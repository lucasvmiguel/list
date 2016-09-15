# List
[![Build Status](https://travis-ci.org/lucasvmiguel/list.svg?branch=master)](https://travis-ci.org/lucasvmiguel/list)
[![GoDoc](https://godoc.org/github.com/lucasvmiguel/list?status.svg)](https://godoc.org/github.com/lucasvmiguel/list)

## Description 

List is a library to handle with slice in golang easily, it's inspired by the Class List .NET 

https://msdn.microsoft.com/pt-br/library/6sh2ey19(v=vs.110).aspx

## Installation

```bash
go get github.com/lucasvmiguel/list
```

## How to use

``` go
import "github.com/lucasvmiguel/list"

func main() {
  l := Generic{}
  l.Add("1")
  l.Add("2")
  l.Add("3")

  l.Me() //[]interface{"1", "2", "3"}
  l.ToSliceString() //[]string{"1", "2", "3"}
}
```

## RoadMap

-Método DeleteRange
-Outros tipos de list(string, int...)
-Gerar benchs
-Verificar outros métodos do .Net(List)

## License

  [MIT](LICENSE)