package tower

type State struct {
	// The tower's current heading, generally the direction the gun
	// is pointing, where 0 is "north" (straight up), decreasing
	// values indicate rotation to the left, and increasing values
	// indicate rotation to the right.
	Heading float64

	// The amount of time, in milliseconds, until the tower is ready
	// to fire its gun.
	TimeToFire uint

	X float64
	Y float64
}
