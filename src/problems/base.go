package problems

import (
	"fmt"
	"reflect"
	"runtime"
)

type FitnessFunctionPositional func(dimensions int, positions []float64) float64

func GetFunctionName(fn FitnessFunctionPositional) string {
	// Get the function name using runtime.FuncForPC
	fnPtr := runtime.FuncForPC(reflect.ValueOf(fn).Pointer())
	if fnPtr != nil {
		return fnPtr.Name()
	}
	return "unknown"
}

func GetCleanFunctionName(fn FitnessFunctionPositional) string {
	fnPtr := runtime.FuncForPC(reflect.ValueOf(fn).Pointer())
	if fnPtr != nil {
		name := fnPtr.Name()
		// Extract just the function name (remove package path)
		if lastDot := len(name) - 1; lastDot >= 0 {
			for i := lastDot; i >= 0; i-- {
				if name[i] == '.' {
					return name[i+1:]
				}
			}
		}
		return name
	}
	return "unknown"
}

func ClampArray(array []float64, minBound, maxBound float64) {
	for i := range array {
		if array[i] < minBound {
			array[i] = minBound
		} else if array[i] > maxBound {
			array[i] = maxBound
		}
	}
}

func ClampPositionsBounceVelocities(positions, velocities []float64, minBound, maxBound float64) {
	for i := range positions {
		if positions[i] < minBound {
			positions[i] = minBound
			velocities[i] = -velocities[i]
			fmt.Println("Bounced velocities from", -velocities[i], "to", velocities[i])
		} else if positions[i] > maxBound {
			positions[i] = maxBound
			velocities[i] = -velocities[i]
			fmt.Println("Bounced velocities from", -velocities[i], "to", velocities[i])
		}

		if velocities[i] < minBound {
			velocities[i] = minBound
		} else if velocities[i] > maxBound {
			velocities[i] = maxBound
		}
	}
}
