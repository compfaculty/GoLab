package carman

func Manufacturer(vin string) string {
	manufacturer := vin[:3]
	if manufacturer[2] == '9' {
		manufacturer += vin[11:14]
	}
	return manufacturer

}
