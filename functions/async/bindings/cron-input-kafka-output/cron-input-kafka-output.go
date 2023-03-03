package bindings

import (
	"encoding/json"
	"log"

	ofctx "github.com/OpenFunction/functions-framework-go/context"
)

func HandleCronInput(ctx ofctx.Context, in []byte) (ofctx.Out, error) {
	var greeting []byte
	if in != nil {
		log.Printf("binding - Data: %s", in)
		greeting = in
	} else {
		log.Print("binding - Data: Received")
		greeting, _ = json.Marshal(map[string]string{"message": "Hello"})
	}

	if ctx.HasOutputs() {
		for key := range ctx.GetOutputs() {
			_, err := ctx.Send(key, greeting)
	                if err != nil {
		            log.Printf("Error: %v\n", err)
		            return ctx.ReturnOnInternalError(), err
	                }
		}
	
	}
	
	return ctx.ReturnOnSuccess(), nil
}
