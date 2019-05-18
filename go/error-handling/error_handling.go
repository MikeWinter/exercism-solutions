package erratum

import "errors"

func Use(o ResourceOpener, input string) (err error) {
	res, e := tryOpen(&o)
	if e != nil {
		return e
	}
	defer alwaysClose(&res)
	defer onError(&res, &err)

	res.Frob(input)
	return nil
}

func tryOpen(o *ResourceOpener) (Resource, error) {
	for {
		if res, err := (*o)(); err == nil {
			return res, nil
		} else if _, transient := err.(TransientError); !transient {
			return nil, err
		}
	}
}

func alwaysClose(res *Resource) {
	_ = (*res).Close()
}

func onError(res *Resource, err *error) {
	if recovered := recover(); recovered != nil {
		switch e := recovered.(type) {
		case FrobError:
			*err = e
			(*res).Defrob(e.defrobTag)
		case error:
			*err = e
		default:
			*err = errors.New(`an error occurred`)
		}
	}
}
