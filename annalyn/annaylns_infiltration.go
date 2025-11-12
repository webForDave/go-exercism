package annalyn

func CanFastAttack(knightIsAwake bool) bool {
	if !knightIsAwake{
		return true
	}
	return false
}

func CanSpy(knightIsAwake, archerisAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerisAwake || prisonerIsAwake{
		return true
	}
	return false
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if archerIsAwake{
		return false
	} else if !prisonerIsAwake{
		return false
	}
	return true
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if petDogIsPresent && !archerIsAwake{
		return true
	} else if prisonerIsAwake && !archerIsAwake && !knightIsAwake{
		return true
	}
	return false
}