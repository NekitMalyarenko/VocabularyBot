package telegramData

import "log"

type ButtonsHolder struct {
	data map[int]func(ActionData)bool
}

var (
	buttonsHolder *ButtonsHolder
)

func GetButtonsHolder() *ButtonsHolder {

	if buttonsHolder == nil {
		buttonsHolder = &ButtonsHolder{
			data: make(map[int]func(ActionData)bool, 0),
		}
		log.Println("New Instance of ButtonsData holder")
	}

	return buttonsHolder
}

func (holder *ButtonsHolder) RegisterButton(id int, function func(ActionData)bool) {
	holder.data[id] = function
}

func (holder *ButtonsHolder) GetButton(id int) func(ActionData)bool {
	return holder.data[id]
}
