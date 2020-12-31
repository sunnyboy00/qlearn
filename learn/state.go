package learn

import "errors"

type State struct {
	Name        string
	Actions     map[string]Action
	ActionsList []Action
}

func (s State) addAction(nextState State, name string, reward float64) {
	action := Action{Name: name, Reward: reward, NextState: nextState}
	s.Actions[name] = action
	s.ActionsList = append(s.ActionsList, action)
}

func (s State) randomAction() (Action, error) {
	if len(s.ActionsList) == 0 {
		err := errors.New("state doesn't contain actions ")
		return Action{}, err
	}
	index := generateRandomNumber(0, len(s.ActionsList))
	return s.ActionsList[index], nil
}
