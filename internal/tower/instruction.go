package tower

type Instruction struct {
	// Whether the tower should fire the next time it is able to do
	// so. This does not necessarily mean it will fire immediately.
	Fire bool

	// Aim is the direction the tower should rotate where values
	// <0 indicate rightward rotation, values >0 indicate leftward
	// rotation, and 0 indicates no rotation.
	Aim int
}
