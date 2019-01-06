package err_handler

// Base exception interface. You may use any kind of objects in Catch() function.
type Exception interface{}

// TryCatch struct. Finally() method also is available, but in fact it useless. Built-in defer is enough.
type TryCatch struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

// The main endpoint to run try-catch handler.
func (ctl TryCatch) Do() {
	if ctl.Finally != nil {
		defer ctl.Finally()
	}
	if ctl.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				ctl.Catch(r)
			}
		}()
	}
	ctl.Try()
}
