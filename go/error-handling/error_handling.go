package erratum

// Use opens a resource, call Frob(input) on the result
// resource and then closes that resource (in all cases).
func Use(o ResourceOpener, input string) (err error) {
	var r Resource
	r, err = o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		r, err = o()
	}
	defer r.Close()
	defer func() {
		if x := recover(); x != nil {
			if frob, ok := x.(FrobError); ok {
				r.Defrob(frob.defrobTag)
			}
			err = x.(error)
		}
	}()
	r.Frob(input)
	if err != nil {
		return err
	}
	return nil
}
