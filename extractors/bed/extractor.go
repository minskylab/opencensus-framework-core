package bed

import (
	"opencensus/core/common"
	"opencensus/core/dkan"
	"strconv"
)

const (
	KindAgeAdult        = "Adult"
	KindAgePediatrician = "Pediatrician"
)

var KindAges = []string{
	KindAgeAdult,
	KindAgePediatrician,
}

const (
	KindBedZNC  = "ZNC"
	KindBedZC   = "ZC"
	KindBedUCIN = "UCIN"

	KindBedUCI        = "UCI"
	KindBedVentilator = "Ventilator"
)

var KindCrossedBeds = []string{
	KindBedUCI,
	KindBedVentilator,
}
var KinDistinctBeds = []string{
	KindBedZNC, KindBedZC,
	KindBedUCIN,
}

type Record struct {
	Name      string
	CovidZone bool

	Institution string
	Code        string
	Category    string
	Correlative string

	CutDate      string
	RegisterDate string

	Region   string
	Province string
	// District string
	Ubigeo string

	ZNCBusy      int
	ZNCAvailable int
	ZNCTotal     int

	ZCBusy      int
	ZCAvailable int
	ZCTotal     int

	HospBusy      int
	HospAvailable int
	HospTotal     int

	UCINBusy      int
	UCINAvailable int
	UCINTotal     int

	UCIPediatraBusy      int
	UCIPediatraAvailable int
	UCIPediatraTotal     int

	UCIAdultBusy      int
	UCIAdultAvailable int
	UCIAdultTotal     int

	VentilatorAdultBusy      int
	VentilatorAdultAvailable int
	VentilatorAdultTotal     int

	VentilatorPediatraBusy      int
	VentilatorPediatraAvailable int
	VentilatorPediatraTotal     int

	MainSourceKind string
}

const dateLayout = "20060102"

func Extract(lapses int) chan []Record {
	channelRecords := make(chan []Record)
	api, err := dkan.NewAPI(common.SuSaludDKANEndpoint)
	if err != nil {
		panic(err)
	}

	bedRes := dkan.ResourceWithID("0badb458-b44f-49ad-bd32-51acaaee7a05")
	bedRes.Sort("FECHA_CORTE", false)
	bedRes.First100()

	go extractor(api, bedRes, lapses, channelRecords)

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
			covidZone, _ := rec["ZONA_COVID"].(string)

			institution, _ := rec["INSTITUCION"].(string)
			category, _ := rec["CATEGORIA"].(string)
			correlative, _ := rec["CORRELATIVO"].(string)

			code, _ := rec["CODIGO"].(string)

			cutDate, _ := rec["FECHA_CORTE"].(string)
			registerDate, _ := rec["FECHA_REGISTRO"].(string)

			region, _ := rec["REGION"].(string)
			province, _ := rec["PROVINCIA"].(string)
			// district, _ := rec["DISTRITO"].(string)
			ubigeo, _ := rec["UBIGEO"].(string)

			hospAvailable, _ := rec["CAMAS_HOSP_DISPONIBLE"].(string)
			hospBusy, _ := rec["CAMAS_HOSP_OCUPADAS"].(string)
			hospTotal, _ := rec["CAMAS_HOSP_TOTAL"].(string)

			ucinAvailable, _ := rec["UCIN_CAMAS_DISPONIBLE"].(string)
			ucinBusy, _ := rec["UCIN_CAMAS_OCUPADAS"].(string)
			ucinTotal, _ := rec["UCIN_CAMAS_TOTAL"].(string)

			uciAdultAvailable, _ := rec["UCI_ADULTOS_CAMAS_DISPONIBLE"].(string)
			uciAdultBusy, _ := rec["UCI_ADULTOS_CAMAS_OCUPADAS"].(string)
			uciAdultTotal, _ := rec["UCI_ADULTOS_CAMAS_TOTAL"].(string)

			uciPediatraAvailable, _ := rec["UCI_PEDIATRIA_CAMAS_DISPONIBLE"].(string)
			uciPediatraBusy, _ := rec["UCI_PEDIATRIA_CAMAS_OCUPADAS"].(string)
			uciPediatraTotal, _ := rec["UCI_PEDIATRIA_CAMAS_TOTAL"].(string)

			ventilatorAdultAvailable, _ := rec["VENTILADORES_UCI_ADULTO_DISPONIBLE"].(string)
			ventilatorAdultBusy, _ := rec["VENTILADORES_UCI_ADULTO_OCUPADOS"].(string)
			ventilatorAdultTotal, _ := rec["VENTILADORES_UCI_ADULTO_TOTAL"].(string)

			ventilatorPediatraAvailable, _ := rec["VENTILADORES_UCI_PEDIATRIA_DISPONIBLE"].(string)
			ventilatorPediatraBusy, _ := rec["VENTILADORES_UCI_PEDIATRIA_OCUPADOS"].(string)
			ventilatorPediatraTotal, _ := rec["VENTILADORES_UCI_PEDIATRIA_TOTAL"].(string)

			zcAvailable, _ := rec["CAMAS_ZC_DISPONIBLES"].(string)
			zcBusy, _ := rec["CAMAS_ZC_OCUPADOS"].(string)
			zcTotal, _ := rec["CAMAS_ZC_TOTAL"].(string)

			zncAvailable, _ := rec["CAMAS_ZNC_DISPONIBLE"].(string)
			zncBusy, _ := rec["CAMAS_ZNC_OCUPADOS"].(string)
			zncTotal, _ := rec["CAMAS_ZNC_TOTAL"].(string)

			mainSourceKind := "idk"

			var covidZoneBool bool
			if covidZone == "Si" {
				covidZoneBool = true
			} else {
				covidZoneBool = false
			}
			hospAvailableNumber, _ := strconv.Atoi(hospAvailable)
			hospBusyNumber, _ := strconv.Atoi(hospBusy)
			hospTotalNumber, _ := strconv.Atoi(hospTotal)

			ucinAvailableNumber, _ := strconv.Atoi(ucinAvailable)
			ucinBusyNumber, _ := strconv.Atoi(ucinBusy)
			ucinTotalNumber, _ := strconv.Atoi(ucinTotal)

			uciAdultAvailableNumber, _ := strconv.Atoi(uciAdultAvailable)
			uciAdultBusyNumber, _ := strconv.Atoi(uciAdultBusy)
			uciAdultTotalNumber, _ := strconv.Atoi(uciAdultTotal)

			uciPediatraAvailableNumber, _ := strconv.Atoi(uciPediatraAvailable)
			uciPediatraBusyNumber, _ := strconv.Atoi(uciPediatraBusy)
			uciPediatraTotalNumber, _ := strconv.Atoi(uciPediatraTotal)

			zcAvailableNumber, _ := strconv.Atoi(zcAvailable)
			zcBusyNumber, _ := strconv.Atoi(zcBusy)
			zcTotalNumber, _ := strconv.Atoi(zcTotal)

			zncAvailableNumber, _ := strconv.Atoi(zncAvailable)
			zncBusyNumber, _ := strconv.Atoi(zncBusy)
			zncTotalNumber, _ := strconv.Atoi(zncTotal)

			ventilatorAdultAvailableNumber, _ := strconv.Atoi(ventilatorAdultAvailable)
			ventilatorAdultBusyNumber, _ := strconv.Atoi(ventilatorAdultBusy)
			ventilatorAdultTotalNumber, _ := strconv.Atoi(ventilatorAdultTotal)

			ventilatorPediatraAvailableNumber, _ := strconv.Atoi(ventilatorPediatraAvailable)
			ventilatorPediatraBusyNumber, _ := strconv.Atoi(ventilatorPediatraBusy)
			ventilatorPediatraTotalNumber, _ := strconv.Atoi(ventilatorPediatraTotal)

			recordsArray = append(recordsArray, Record{
				Name:         name,
				CovidZone:    covidZoneBool,
				Institution:  institution,
				Code:         code,
				Category:     category,
				Correlative:  correlative,
				CutDate:      cutDate,
				RegisterDate: registerDate,
				Region:       region,
				Province:     province,
				// District:           district,
				Ubigeo: ubigeo,

				ZNCBusy:      zncBusyNumber,
				ZNCAvailable: zncAvailableNumber,
				ZNCTotal:     zncTotalNumber,

				ZCBusy:      zcBusyNumber,
				ZCAvailable: zcAvailableNumber,
				ZCTotal:     zcTotalNumber,

				HospBusy:      hospBusyNumber,
				HospAvailable: hospAvailableNumber,
				HospTotal:     hospTotalNumber,

				UCINAvailable: ucinAvailableNumber,
				UCINBusy:      ucinBusyNumber,
				UCINTotal:     ucinTotalNumber,

				UCIAdultBusy:      uciAdultBusyNumber,
				UCIAdultAvailable: uciAdultAvailableNumber,
				UCIAdultTotal:     uciAdultTotalNumber,

				UCIPediatraBusy:      uciPediatraBusyNumber,
				UCIPediatraAvailable: uciPediatraAvailableNumber,
				UCIPediatraTotal:     uciPediatraTotalNumber,

				VentilatorAdultBusy:      ventilatorAdultBusyNumber,
				VentilatorAdultAvailable: ventilatorAdultAvailableNumber,
				VentilatorAdultTotal:     ventilatorAdultTotalNumber,

				VentilatorPediatraBusy:      ventilatorPediatraBusyNumber,
				VentilatorPediatraAvailable: ventilatorPediatraAvailableNumber,
				VentilatorPediatraTotal:     ventilatorPediatraTotalNumber,

				MainSourceKind: mainSourceKind,
			})
		}

		channel <- recordsArray
		res.NextN(100)
	}

	close(channel)
}
