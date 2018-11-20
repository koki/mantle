package initialize

import (
	"io"
	"os"

	"mantle/pkg/codec"
)

func MantleInit() error {
	out, err := codec.Decode(os.Stdin)
	if err != nil {
		return err
	}
	_, err = io.Copy(os.Stdout, out)
	return err
}
