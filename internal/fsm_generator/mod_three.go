package fsm_generator

type State = Mod3State

var mod3FSM = NewMod3FSM()

func ModThree(s string) (int, error) {
	var q, err = mod3FSM.ProcessInput(s)
	if err != nil {
		return 0, err
	}
	return int(q), nil
}
