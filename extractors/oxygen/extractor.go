package oxygen

import (
	"fmt"
	"opencensus/core/common"
	"opencensus/core/dkan"
)

func Extract() {
	api, err := dkan.NewAPI(common.SuSaludDKANEndpoint)
	if err != nil {
		panic(err)
	}

	oxygenRes := dkan.ResourceWithID("0badb458-b44f-49ad-bd32-51acaaee7a05")

	data, err := api.ObtainResource(oxygenRes.First100())
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
