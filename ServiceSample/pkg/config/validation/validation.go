package validation

import "errors"

type Warning error

func AppendError(err error, errs ...error) error {
	if len(errs) == 0 {
		return err
	}
	return errors.Join(append([]error{err}, errs...)...)
}
