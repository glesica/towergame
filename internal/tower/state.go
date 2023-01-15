package tower

import "github.com/jakecoffman/cp"

type State struct {
	// The tower's current heading, generally the direction the gun
	// is pointing. This should be a unit vector.
	Heading cp.Vector

	// Position is the location of the tower in 2D space.
	Position cp.Vector

	// The amount of time, in milliseconds, until the tower is ready
	// to fire its gun.
	TimeToFire uint
}
