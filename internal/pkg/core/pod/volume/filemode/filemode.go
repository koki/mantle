package filemode

import (
	"mantle/pkg/util"
)

// FileMode can be unmarshalled from either a number (octal is supported) or a string.
// The json library doesn't allow serializing numbers as octal, so FileMode always marshals to a string.
type FileMode int32

func FileModePtr(m FileMode) *FileMode {
	return &m
}

func ConvertFileModeToInt32Ptr(mode *FileMode) *int32 {
	if mode == nil {
		return nil
	}

	return util.Int32Ptr(int32(*mode))
}

func NewFileModeFromKubeV1(kubeMode *int32) *FileMode {
	if kubeMode == nil {
		return nil
	}

	return FileModePtr(FileMode(*kubeMode))
}
