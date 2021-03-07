package bed

import (
	"context"
	"log"
	"opencensus/core/ent"
	"opencensus/core/ent/district"
	"opencensus/core/ent/place"
	"opencensus/core/ent/province"
	"opencensus/core/ent/region"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func stringInMap(val string, array map[string]int) bool {
	for v := range array {
		if v == val {
			return true
		}
	}
	return false
}

func stringDateToTime(date string) time.Time {
	year := date[0:4]
	month := date[4:6]
	day := date[6:]

	y, _ := strconv.Atoi(year)
	m, _ := strconv.Atoi(month)
	d, _ := strconv.Atoi(day)

	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
}

type record struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func regionsMap(ctx context.Context, client *ent.Client) (map[string]int, error) {
	records := []record{}

	if err := client.Region.Query().Select(region.FieldID, region.FieldName).Scan(ctx, &records); err != nil {
		return nil, errors.WithStack(err)
	}

	regions := map[string]int{}

	for _, r := range records {
		regions[r.Name] = r.ID
	}

	return regions, nil
}

func provincesMap(ctx context.Context, client *ent.Client) (map[string]int, error) {
	records := []record{}

	if err := client.Province.Query().Select(province.FieldID, province.FieldName).Scan(ctx, &records); err != nil {
		return nil, errors.WithStack(err)
	}

	provinces := map[string]int{}

	for _, r := range records {
		provinces[r.Name] = r.ID
	}

	return provinces, nil
}

func districtsMap(ctx context.Context, client *ent.Client) (map[string]int, error) {
	records := []record{}

	if err := client.District.Query().Select(district.FieldID, district.FieldName).Scan(ctx, &records); err != nil {
		return nil, errors.WithStack(err)
	}

	districts := map[string]int{}

	for _, r := range records {
		districts[r.Name] = r.ID
	}

	return districts, nil
}

func placesMap(ctx context.Context, client *ent.Client) (map[string]int, error) {
	records := []record{}

	if err := client.Place.Query().Select(place.FieldID, place.FieldName).Scan(ctx, &records); err != nil {
		return nil, errors.WithStack(err)
	}

	places := map[string]int{}

	for _, r := range records {
		places[r.Name] = r.ID
	}

	return places, nil
}

func processor(ctx context.Context, client *ent.Client, records []Record) error {
	currentRegions, err := regionsMap(ctx, client)
	if err != nil {
		return errors.WithStack(err)
	}

	currentProvinces, err := provincesMap(ctx, client)
	if err != nil {
		return errors.WithStack(err)
	}

	totalPlaces, err := placesMap(ctx, client)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, record := range records {
		var regionID, provinceID, districtID, placeID int
		districtID = 1032
		if stringInMap(record.Region, currentRegions) {
			regionID = currentRegions[record.Region]
		} else {
			reg, err := client.Region.Create().
				SetName(record.Region).
				Save(ctx)
			if err != nil {
				return errors.WithStack(err)
			}

			currentRegions[reg.Name] = reg.ID
			regionID = reg.ID
		}

		if stringInMap(record.Province, currentProvinces) {
			provinceID = currentProvinces[record.Province]
		} else {
			prov, err := client.Province.Create().
				SetName(record.Province).
				Save(ctx)
			if err != nil {
				return errors.WithStack(err)
			}

			currentProvinces[prov.Name] = prov.ID
			provinceID = prov.ID
		}

		if stringInMap(record.Name, totalPlaces) {
			placeID = totalPlaces[record.Name]
		} else {
			place, err := client.Place.Create().
				SetCovidZone(record.CovidZone). //  covid zone
				SetRegionID(regionID).
				SetProvinceID(provinceID).
				SetDistrictID(districtID).
				SetUbigeo(record.Ubigeo).
				SetName(record.Name).
				SetKind(record.Institution).
				Save(ctx)
			if err != nil {
				return errors.WithStack(err)
			}

			totalPlaces[place.Name] = place.ID
			placeID = place.ID
		}

		// kind age adult, pediatra, ucin
		// kind bed hosp, zc, znc  ventilator, uci, ucin

		for _, kindBed := range KindCrossedBeds {
			for _, kindAge := range KindAges {

				var err error
				var bedRecord *ent.BedRecord

				switch kindBed {
				case KindBedUCI:
					switch kindAge {
					case KindAgeAdult:
						bedRecord, err = insertBed(ctx, client, &record, placeID, KindBedUCI, kindAge, record.UCIAdultBusy, record.UCIAdultAvailable, record.UCIAdultTotal)
						if err != nil {
							return errors.WithStack(err)
						}
						log.Printf("new bed record: %d | %s\n", bedRecord.ID, record.Name)

						break
					case KindAgePediatrician:
						bedRecord, err = insertBed(ctx, client, &record, placeID, KindBedUCI, kindAge, record.UCIPediatraBusy, record.UCIPediatraAvailable, record.UCIPediatraTotal)
						if err != nil {
							return errors.WithStack(err)
						}
						log.Printf("new bed record: %d | %s\n", bedRecord.ID, record.Name)
						break
					}
					break
				case KindBedVentilator:
					switch kindAge {
					case KindAgeAdult:
						bedRecord, err = insertBed(ctx, client, &record, placeID, KindBedVentilator, kindAge, record.VentilatorAdultBusy, record.VentilatorAdultAvailable, record.VentilatorAdultTotal)
						if err != nil {
							return errors.WithStack(err)
						}
						log.Printf("new bed record: %d | %s\n", bedRecord.ID, record.Name)
						break
					case KindAgePediatrician:
						bedRecord, err = insertBed(ctx, client, &record, placeID, KindBedVentilator, kindAge, record.VentilatorPediatraBusy, record.VentilatorPediatraAvailable, record.VentilatorPediatraTotal)
						if err != nil {
							return errors.WithStack(err)
						}
						log.Printf("new bed record: %d | %s\n", bedRecord.ID, record.Name)
						break
					}
					break
				}
			}
		}

		for _, kindBed := range KinDistinctBeds {
			var err error
			var bedRecord *ent.BedRecord

			switch kindBed {
			case KindBedZC:
				bedRecord, err = insertBed(ctx, client, &record, placeID, KindBedZC, "", record.ZCBusy, record.ZCAvailable, record.ZCTotal)
				if err != nil {
					return errors.WithStack(err)
				}
				log.Printf("new bed record: %d | %s\n", bedRecord.ID, record.Name)
				break
			case KindBedZNC:
				bedRecord, err = insertBed(ctx, client, &record, placeID, KindBedZNC, "", record.ZNCBusy, record.ZNCAvailable, record.ZNCTotal)
				if err != nil {
					return errors.WithStack(err)
				}
				log.Printf("new bed record: %d | %s\n", bedRecord.ID, record.Name)
				break
			case KindBedUCIN:
				bedRecord, err = insertBed(ctx, client, &record, placeID, KindBedUCIN, "", record.UCINBusy, record.UCINAvailable, record.UCINTotal)
				if err != nil {
					return errors.WithStack(err)
				}
				log.Printf("new bed record: %d | %s\n", bedRecord.ID, record.Name)
				break
			}
		}
	}

	return nil
}

func insertBed(ctx context.Context, client *ent.Client, record *Record, placeID int, kind string, ageKind string, busyBeds int, availableBeds int, totalBeds int) (*ent.BedRecord, error) {

	bedRecordOnCreate := client.BedRecord.Create().
		SetCollectedDate(record.CutDate).
		SetReportedDate(record.RegisterDate).
		SetKindBed(kind).
		SetBusyBeds(busyBeds).
		SetAvailableBeds(availableBeds).
		SetKindAge(ageKind).
		SetTotalBeds(totalBeds).
		AddPlaceIDs(placeID)

	bedRecord, err := bedRecordOnCreate.Save(ctx)

	return bedRecord, err
}

func Processor(ctx context.Context, client *ent.Client, records []Record) error {
	return processor(ctx, client, records)
}
