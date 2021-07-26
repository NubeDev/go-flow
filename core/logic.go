package core

func And() Spec {
	return Spec{
		Name:    "and",
		Inputs:  []Pin{Pin{"in", BOOLEAN}, Pin{"in", BOOLEAN}, Pin{"in", BOOLEAN}, Pin{"aidan new in", BOOLEAN}},
		Outputs: []Pin{Pin{"out", BOOLEAN}},
		Kernel: func(in, out, internal MessageMap, s Source, i chan Interrupt) Interrupt {
			x, ok := in[0].(bool)
			if !ok {
				out[0] = NewError("need boolean")
				return nil
			}
			y, ok := in[1].(bool)
			if !ok {
				out[0] = NewError("need boolean")
				return nil
			}
			y, ok = in[2].(bool)
			if !ok {
				out[0] = NewError("need boolean")
				return nil
			}
			out[0] = x && y
			return nil
		},
	}
}

func Or() Spec {
	return Spec{
		Name:    "or",
		Inputs:  []Pin{Pin{"in", BOOLEAN}, Pin{"in", BOOLEAN}},
		Outputs: []Pin{Pin{"out", BOOLEAN}},
		Kernel: func(in, out, internal MessageMap, s Source, i chan Interrupt) Interrupt {
			x, ok := in[0].(bool)
			if !ok {
				out[0] = NewError("need boolean")
				return nil
			}
			y, ok := in[1].(bool)
			if !ok {
				out[0] = NewError("need boolean")
				return nil
			}
			out[0] = x || y
			return nil
		},
	}
}

var toggle int

func Toggle() Spec {
	return Spec{
		Name:    "toggle",
		Inputs:  []Pin{},
		Outputs: []Pin{Pin{"draw", NUMBER}},
		Kernel: func(in, out, internal MessageMap, s Source, i chan Interrupt) Interrupt {
			toggle = (toggle + 1) % 2
			out[0] = toggle
			return nil
		},
	}
}


func Not() Spec {
	return Spec{
		Name:    "not",
		Inputs:  []Pin{Pin{"in", BOOLEAN}},
		Outputs: []Pin{Pin{"out", BOOLEAN}},
		Kernel: func(in, out, internal MessageMap, s Source, i chan Interrupt) Interrupt {
			x, ok := in[0].(bool)
			if !ok {
				out[0] = NewError("need boolean")
				return nil
			}
			out[0] = !x
			return nil
		},
	}
}
