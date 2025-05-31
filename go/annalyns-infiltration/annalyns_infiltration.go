package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	//panic("Please implement the CanFastAttack() function")
	return !knightIsAwake
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	//panic("Please implement the CanSpy() function")
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	//panic("Please implement the CanSignalPrisoner() function")
	return !archerIsAwake && prisonerIsAwake
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	//panic("Please implement the CanFreePrisoner() function")
	//if petDogIsPresent {
	//	return !archerIsAwake
	//} else {
	//	return prisonerIsAwake && !archerIsAwake && !knightIsAwake
	//}
	return prisonerIsAwake && !knightIsAwake && !archerIsAwake || petDogIsPresent && !archerIsAwake
}
