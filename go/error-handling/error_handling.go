package erratum

// Use opens a resource, call Frob(input) on the result
// resource and then closes that resource (in all cases).
func Use(o ResourceOpener, input string) (err error) {
	var r Resource
	r, err = o()
	if err != nil {
		switch err.(type) {
		case TransientError:
			return Use(o, input)
		default:
			return err
		}
	}
	defer r.Close()
	defer func() {
		if x := recover(); x != nil {
			switch x.(type) {
			case FrobError:
				r.Defrob(x.(FrobError).defrobTag)
			case error:
				err = x.(error)
			default:
				panic(x)
			}
		}
	}()
	r.Frob(input)
	if err != nil {
		return err
	}
	return nil
}
