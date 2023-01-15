package tower

type Instruction struct {
	// Whether the tower should fire the next time it is able to do
	// so. This does not necessarily mean it will fire immediately.
	Fire bool

	// Aim is the direction the tower should rotate where values
	// <0 indicate leftward rotation, values >0 indicate rightward
	// rotation, and 0 indicates no rotation.
	Aim int

	// AimSpeed is the speed, as a fraction of its maximum, at
	// which the tower should rotate. Should be in the range [0,1]
	// but values >1 are treated as 1 and values <0 are treated as 0.
	AimSpeed float64
}
