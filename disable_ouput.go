package nod

var disabledOutputs = make(map[string]bool)

func DisableOutput(output string) {
	if hnd, ok := handlers[output]; ok {
		hnd.Close()
	}
	disabledOutputs[output] = true
}

func EnableOutput(output string) {
	delete(disabledOutputs, output)
}
