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

	CutDate      string
	RegisterDate string

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

	oxygenRes := dkan.ResourceWithID("b1791142-8fcb-4766-9b8c-b0ee0ffc6dff")
	oxygenRes.First100()

	go extractor(api, oxygenRes, lapses, channelRecords)

	return channelRecords
}
func subExtractorMax(record map[string]interface{}, nameSlice [5]string) (maxProducer string, production int) {
	maxProduction := 0
	for _, producer := range nameSlice {
		data, _ := strconv.Atoi(record[producer].(string))
		production = production + data
		if data > maxProduction {
			maxProducer = producer
		}
	}
	return maxProducer, production
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

			institution, _ := rec["INSTITUCION"].(string)
			code, _ := rec["CODIGO"].(string)

			cutDate, _ := rec["FECHACORTE"].(string)
			registerDate, _ := rec["FECHAREGISTRO"].(string)

			region, _ := rec["REGION"].(string)
			province, _ := rec["PROVINCIA"].(string)

			totalCylinders, _ := rec["TOT_CILINDROS"].(string)
			totalOwnCylinders, _ := rec["TOT_PROPIOS"].(string)

			productionPlantNames := [5]string{"PRODUCCION_DIAPLA", "PRODUCCION_DIA_CRIO", "PRODUCCION_DIA_GEN", "PRODUCCION_DIA_ISO",
				"PRODUCCION_DIA_OTR"}
			_, dailyProductionNumber := subExtractorMax(rec, productionPlantNames)

			productionMaxPlantNames := [5]string{"PRODUCCION_MAX_PLA", "PRODUCCION_MAX_CRIO", "PRODUCCION_MAX_GEN", "PRODUCCION_MAX_ISO", "PRODUCCION_MAX_OTR"}
			mainSourceKind, maxDailyProductionNumber := subExtractorMax(rec, productionMaxPlantNames)

			consumptionPlantNames := [5]string{"CONSUMO_DIA_CRIO", "CONSUMO_DIA_GEN", "CONSUMO_DIA_ISO", "CONSUMO_DIA_OTR", "CONSUMO_DIA_PLA"}
			_, dailyConsumptionNumber := subExtractorMax(rec, consumptionPlantNames)

			//dailyProduction, _ := rec["PRODUCCION_DIAPLA"].(string)
			//maxDailyProduction, _ := rec["PRODUCCION_MAX_PLA"].(string)
			//dailyConsumption, _ := rec["CONSUMO_DIA_PLA"].(string)

			//mainSourceKind := "idk"

			totalCylindersNumber, _ := strconv.Atoi(totalCylinders)
			totalOwnCylindersNumber, _ := strconv.Atoi(totalOwnCylinders)
			//dailyProductionNumber, _ := strconv.Atoi(dailyProduction)
			//maxDailyProductionNumber, _ := strconv.Atoi(maxDailyProduction)
			///dailyConsumptionNumber, _ := strconv.Atoi(dailyConsumption)

			recordsArray = append(recordsArray, Record{
				Name:               name,
				Institution:        institution,
				Code:               code,
				CutDate:            cutDate,
				RegisterDate:       registerDate,
				Region:             region,
				Province:           province,
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
