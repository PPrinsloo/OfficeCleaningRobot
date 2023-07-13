package main

type directionVector struct {
	X int
	Y int
}

func (dirVector *directionVector) add(vector directionVector) {
	dirVector.X += vector.X
	dirVector.Y += vector.Y
}

// dirctions N, E, S, W
func makeDirectionsVector(direction string) directionVector {
	switch direction {
	case "N":
		return directionVector{0, 1}
	case "E":
		return directionVector{1, 0}
	case "S":
		return directionVector{0, -1}
	case "W":
		return directionVector{-1, 0}
	}
	return directionVector{0, 0}
}
