package learn

import (
	"errors"
	"fmt"
)

type QLearner struct {
	Gamma        float64 // 	discount factor
	Alpha        float64 // 	learning rate
	QValuesTable map[string]map[string]float64
	States       map[string]State
	StatesList   []State
	CurrentState State
}

func (qlearner QLearner) AddState(name string) {
	state := State{Name: name}
	qlearner.States[name] = state
	qlearner.StatesList = append(qlearner.StatesList, state)

}
func (qlearner QLearner) setState(name string) {

	if _, ok := qlearner.States[name]; ok {
		qlearner.CurrentState = State{Name: name}
	} else {
		qlearner.CurrentState = qlearner.States[name]
	}

}
func (qlearner QLearner) randomState() State {
	index := generateRandomNumber(0, len(qlearner.StatesList))
	return qlearner.StatesList[index]
}

func (qlearner QLearner) optimalFutureValue(state State) float64 {
	qValues := qlearner.QValuesTable[state.Name]
	max := 0.0
	for _, qValue := range qValues {
		if max < qValue {
			max = qValue
		}
	}
	return max
}

func (qlearner QLearner) knowsAction(state State, action Action) bool {

	if _, stateKnown := qlearner.QValuesTable[state.Name]; stateKnown {
		if _, actionKnown := qlearner.QValuesTable[action.Name]; actionKnown {
			return true
		}
	}
	return false
}

/**
 * From current state, apply the action (name), and set the action's next state as current state.
 */
func (qlearner QLearner) applyAction(name string) {
	actionObject := qlearner.States[qlearner.CurrentState.Name].Actions[name]
	qlearner.CurrentState = actionObject.NextState
}

/**
 * Get the best action from the state (name), which is the one with immediate best Q-value.
 * During the selection process, if states have the same reward, choose one with probability 50%.
 */
func (qlearner QLearner) bestAction(state string) string {
	qValues := qlearner.QValuesTable[state]
	var bestAction string

	for action := range qValues {

		if len(bestAction) == 0 {
			bestAction = action
		} else if qValues[action] == qValues[bestAction] && random() > 0.5 {
			bestAction = action
		} else if qValues[action] > qValues[bestAction] {
			bestAction = action
		}
	}

	return bestAction
}

func (qlearner QLearner) runOnce() (Action, error) {
	bestActionName := qlearner.bestAction(qlearner.CurrentState.Name)
	var err error

	if len(bestActionName) == 0 {
		err = errors.New("path not explored yet")
	} else {
		err = nil
	}
	action := qlearner.States[qlearner.CurrentState.Name].Actions[bestActionName]

	return action, err

}

func (qlearner QLearner) setQValue(stateName, actionName string, reward float64) {

	if _, ok := qlearner.QValuesTable[stateName]; !ok {

		qlearner.QValuesTable[stateName] = make(map[string]float64)
	}
	qlearner.QValuesTable[stateName][actionName] = reward

}
func (qlearner QLearner) getQValue(stateName, actionName string) float64 {

	return qlearner.QValuesTable[stateName][actionName]
}

func (qlearner QLearner) step() {
	state := qlearner.CurrentState
	action, err := state.randomAction()
	if err != nil {
		fmt.Println(err)
	}
	actionReward := action.Reward

	maxQValue := qlearner.optimalFutureValue(action.NextState)
	oldQValue := qlearner.getQValue(state.Name, action.Name)

	newQValue := (1.0-qlearner.Alpha)*oldQValue + qlearner.Alpha*(actionReward+qlearner.Gamma*maxQValue)

	qlearner.setQValue(state.Name, action.Name, newQValue)

	qlearner.CurrentState = action.NextState

}

func (qlearner QLearner) Learn(steps int) {

	if steps < 1 {
		steps = 1
	}

	for ; steps > 0; steps -= 1 {
		qlearner.CurrentState = qlearner.randomState()
		qlearner.step()
	}
}
