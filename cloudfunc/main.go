package cloudfunc

import (
	"fmt"
	"net/http"
)

var message = "Hello world"

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, message)
}
