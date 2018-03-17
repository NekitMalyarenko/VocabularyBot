package telegramData

import "log"


type ButtonsDataHolder struct {
	data map[int]func(actionData ActionData)bool
}


var  (
	buttonsDataHolder *ButtonsDataHolder
)


func GetButtonsDataHolder() *ButtonsDataHolder {

	if buttonsDataHolder == nil {
		buttonsDataHolder = &ButtonsDataHolder{
			data : make(map[int]func(actionData ActionData)bool, 0),
		}
		log.Println("New Instance of ButtonsData holder")
	}

	return buttonsDataHolder
}


func (holder *ButtonsDataHolder) RegisterButtonData(id int, function func(actionData ActionData)bool)  {
	holder.data[id] = function
}


func (holder *ButtonsDataHolder) GetButtonData(id int) func(actionData ActionData)bool {
	return holder.data[id]
}
