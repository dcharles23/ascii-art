#!/bin/bash

go run . "hello"
go run . "HELLO"
go run . "HeLlo HuMaN"
go run . "1Hello 2There"
go run . "Hello\nThere"
go run . "Hello\n\nThere"
go run . "{Hello & There #}"
go run . 'hello There 1 to 2!'
go run . "MaD3IrA&LiSboN"
go run . "1a\"#FdwHywR&/()="
go run . "{|}~"
go run . "[\]^_ 'a"
go run . "RGB"
go run . ":;<=>?@"
go run . '\!" #$%&'"'"'()*+,-./'
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
go run . "abcdefghijklmnopqrstuvwxyz"

# first time launch = chmod +x ./audit.sh after that just ./audit.sh