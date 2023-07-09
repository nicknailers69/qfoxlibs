package qfoxlibs

import (
	"fmt"
	_ "github.com/nicknailers69/qfoxlibs/bin"
	"github.com/nicknailers69/qfoxlibs/pkg/nn"
)

func main() {
	fmt.Print("main test qfoxlibs")
	nn.Run()
}
