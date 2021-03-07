package oxygen

import (
	"opencensus/core/common"
	"opencensus/core/dkan"
	"strconv"
)

type Record struct {
	Name string

	Institution string
	Code        string

	CutDate      int
	RegisterDate int

	Region   string
	Province string
	District string

	TotalCylinders    int
	TotalOwnCylinders int

	DailyProduction    int
	MaxDailyProduction int

	DailyConsumption int
	MainSourceKind   string
}

func Extract(lapses int) chan []Record {
	channelRecords := make(chan []Record)
	api, err := dkan.NewAPI(common.SuSaludDKANEndpoint)
	if err != nil {
		panic(err)
	}

	oxygenRes := dkan.ResourceWithID("0badb458-b44f-49ad-bd32-51acaaee7a05")
	oxygenRes.First100()

	go extractor(api, oxygenRes, lapses, channelRecords)

	return channelRecords
}

func extractor(api *dkan.API, res *dkan.Resource, lapses int, channel chan []Record) {
	for lapse := 0; lapse < lapses; lapse++ {
		data, err := api.ObtainResource(res)
		if err != nil {
			panic(err) // really?
		}

		records := data["records"].([]interface{})

		recordsArray := []Record{}

		for _, r := range records {
			rec := r.(map[string]interface{})

			name, _ := rec["NOMBRE"].(string)
			// totalCylinders, _ := rec["TOT_CILINDROS"].(string)

			institution, _ := rec["INSTITUCION"].(string)
			code, _ := rec["CODIGO"].(string)

			cutDate, _ := rec["FECHACORTE"].(int)
			registerDate, _ := rec["FECHAREGISTRO"].(int)

			region, _ := rec["REGION"].(string)
			province, _ := rec["PROVINCIA"].(string)
			district, _ := rec["DISTRITO"].(string)

			totalCylinders, _ := rec["TOT_CILINDROS"].(string)
			totalOwnCylinders, _ := rec["TOT_PROPIOS"].(string)

			dailyProduction, _ := rec["PRODUCCION_DIAPLA"].(string)
			maxDailyProduction, _ := rec["PRODUCCION_MAX_PLA"].(string)

			dailyConsumption, _ := rec["CONSUMO_DIA_PLA"].(string)

			// mainSourceKind  , _:= rec[""].(string)
			mainSourceKind := "idk"

			// rec[""].(string)

			totalCylindersNumber, _ := strconv.Atoi(totalCylinders)
			totalOwnCylindersNumber, _ := strconv.Atoi(totalOwnCylinders)
			dailyProductionNumber, _ := strconv.Atoi(dailyProduction)
			maxDailyProductionNumber, _ := strconv.Atoi(maxDailyProduction)
			dailyConsumptionNumber, _ := strconv.Atoi(dailyConsumption)

			recordsArray = append(recordsArray, Record{
				Name:               name,
				Institution:        institution,
				Code:               code,
				CutDate:            cutDate,
				RegisterDate:       registerDate,
				Region:             region,
				Province:           province,
				District:           district,
				TotalCylinders:     totalCylindersNumber,
				TotalOwnCylinders:  totalOwnCylindersNumber,
				DailyProduction:    dailyProductionNumber,
				MaxDailyProduction: maxDailyProductionNumber,
				DailyConsumption:   dailyConsumptionNumber,
				MainSourceKind:     mainSourceKind,
			})
		}

		channel <- recordsArray
		res.NextN(100)
	}

	close(channel)
}
