package qfoxlibs

import (
	"fmt"
	"log"

	_ "github.com/nicknailers69/qfoxlibs/pkg/nn"
	"github.com/open-quantum-safe/liboqs-go/oqs"
)

func init() {
	log.Print("main qfoxlibs")
	fmt.Println(oqs.EnabledSigs())

}
