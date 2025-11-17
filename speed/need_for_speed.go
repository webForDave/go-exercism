package speed


type Car struct{
	Battery int
	BatteryDrain int
	Speed int
	Distance int 
}

type Track struct{
	distance int
}

func NewCar(carSpeed, batteryDrainPercentage int) Car{

	return Car{
		Speed: carSpeed,
		BatteryDrain: batteryDrainPercentage,
		Battery: 100,
	}
}

func NewTrack(trackDistance int) Track{
	return Track{
		distance: trackDistance,
	}
}

func Drive(car Car) Car{
	
	if car.BatteryDrain > car.Battery{
		return car
	}

	car.Battery = car.Battery - car.BatteryDrain

	if car.Distance > 0{
		car.Distance = car.Distance + car.Speed
	} else {
		car.Distance = car.Speed
	}

	return car
}

func CanFinish(car Car, track Track) bool {
    currentCar := car
    
    for currentCar.Distance < track.distance {
        if currentCar.Battery < currentCar.BatteryDrain {
            return false
        }
        currentCar = Drive(currentCar)
    }
    
    return true
}