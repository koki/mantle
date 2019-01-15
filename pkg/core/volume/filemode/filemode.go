package filemode

import (
	"fmt"
	"strconv"

	"mantle/pkg/util"

	"github.com/koki/json"
	serrors "github.com/koki/structurederrors"
)

// FileMode can be unmarshalled from either a number (octal is supported) or a string.
// The json library doesn't allow serializing numbers as octal, so FileMode always marshals to a string.
type FileMode int32

func (m *FileMode) UnmarshalJSON(data []byte) error {
	var i int32
	err := json.Unmarshal(data, &i)
	if err == nil {
		*m = FileMode(i)
		return nil
	}

	var str string
	err = json.Unmarshal(data, &str)
	if err != nil {
		return serrors.InvalidValueErrorf(string(data), "unable to unmarshal data")
	}

	mode, err := strconv.ParseInt(str, 8, 32)
	if err != nil {
		return serrors.InvalidValueErrorf(str, "file mode should be an octal integer, written either as string or number")
	}

	*m = FileMode(int32(mode))
	return nil
}

func (m FileMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("0%o", m))
}

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
