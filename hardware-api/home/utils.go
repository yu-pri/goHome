package home

import (
	"log"
	"math"
	"os"
	"time"
)

/*
ReportAlert sends message to the system owner "emailTo"
*/
func ReportAlert(b string, s string) error {
	// TODO: implement email alerts
	return nil
}


/*
Exists returns file exists
*/
func Exists(name string) (bool, error) {
	_, err := os.Stat(name)

	//log.Println(v)

	if err != nil {
		log.Println(err.Error())
	}

	if os.IsNotExist(err) {
		log.Println("seems no file: ", name)
		return false, nil
	}

	return err == nil, err
}

/*
IsChipTimeZone - returns true between 23 - 7 local time where we have the chipest electricity
*/
func IsChipTimeZone() bool {
	hour := time.Now().Hour()
	year := time.Now().Year()

	if year < 2000 {
		log.Println("Year: ", year, "Device is not configured")
		return false
	}

	ret := false
	if hour >= ElectroOnFrom || hour < ElectroOnTo {
		ret = true
	}

	return ret
}

/*
Round64 - return Round(price, 0.5, 2)
*/
func Round64(val float64, roundOn float64, places int) float64 {

	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)

	var round float64
	if val > 0 {
		if div >= roundOn {
			round = math.Ceil(digit)
		} else {
			round = math.Floor(digit)
		}
	} else {
		if div >= roundOn {
			round = math.Floor(digit)
		} else {
			round = math.Ceil(digit)
		}
	}

	return round / pow
}

/*
Round32 - return Round(price, 0.5, 2)
*/
func Round32(val float32, roundOn float64, places int) float32 {

	pow := float32(math.Pow(10, float64(places)))
	digit := pow * val
	_, div := math.Modf(float64(digit))

	var round float64
	if val > 0 {
		if div >= roundOn {
			round = math.Ceil(float64(digit))
		} else {
			round = math.Floor(float64(digit))
		}
	} else {
		if div >= roundOn {
			round = math.Floor(float64(digit))
		} else {
			round = math.Ceil(float64(digit))
		}
	}

	return float32(round) / pow
}
