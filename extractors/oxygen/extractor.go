package oxygen

import (
	"opencensus/core/common"
	"opencensus/core/dkan"
	"strconv"
)

type Record struct {
	Name           string
	TotalCylinders int64
}

func Extract(lapses int) chan []Record {
	channelRecords := make(chan []Record)
	api, err := dkan.NewAPI(common.SuSaludDKANEndpoint)
	if err != nil {
		panic(err)
	}

	oxygenRes := dkan.ResourceWithID("0badb458-b44f-49ad-bd32-51acaaee7a05")
	oxygenRes.First100()

	go func(res *dkan.Resource, channel chan []Record) {
		for lapse := 0; lapse < lapses; lapse++ {
			data, err := api.ObtainResource(oxygenRes)
			if err != nil {
				panic(err) // really?
			}

			records := data["records"].([]map[string]interface{})

			recordsArray := []Record{}

			for _, r := range records {
				name, _ := r["NOMBRE"].(string)
				totalCylinders, _ := r["TOT_CILINDROS"].(string)

				totalCylindersNumber, _ := strconv.Atoi(totalCylinders)

				recordsArray = append(recordsArray, Record{
					Name:           name,
					TotalCylinders: int64(totalCylindersNumber),
				})
			}

			channel <- recordsArray
		}
		oxygenRes.NextN(100)
	}(oxygenRes, channelRecords)

	return channelRecords
}
