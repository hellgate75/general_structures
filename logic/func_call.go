package logic

import (
	"runtime"
)

// Get Current Function runtime.Frame
// Returns:
//    runtime.Frame Selected frame or deafult <function name: unknown> fake frame
func GetCurrentFunction() runtime.Frame {
	// Skip GetCurrentFunction
	return getFrame(1)
}

// Get Current Function Caller runtime.Frame
// Returns:
//    runtime.Frame Selected frame or deafult <function name: unknown> fake frame
func GetCallerFunction() runtime.Frame {
	// Skip GetCallerFunction and the function to get the caller of
	return getFrame(2)
}

// Get Current Function Caller runtime.Frame before provided back iterations, starting with 1 as caller
// Parameters:
//    framesBefore (int) Number of previous callers iterations to verify
// Returns:
//    runtime.Frame Selected frame or deafult <function name: unknown> fake frame
func GetCallerFunctionBeforeCurrent(framesBefore int) runtime.Frame {
	// Skip GetCallerFunction before (1 is previous function call) and the function to get the caller of
	return getFrame(1 + framesBefore)
}

func getFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}
