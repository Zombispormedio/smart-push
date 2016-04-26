package config

import (
    "os"
    "strconv"
)


func PacketFrequency() int{
    var freq int
	envFreq:=os.Getenv("PACKET_FREQUENCY")
	
	if envFreq!=""{
		var ConversionError error
		freq, ConversionError=strconv.Atoi(envFreq)
		
		if ConversionError!=nil{
			freq=50
		}
	}else{
		freq=50
	}
	
    return freq
}