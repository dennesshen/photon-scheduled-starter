package schedule

import "reflect"

type Command func()

var cronCommandStructLists = make([]interface{}, 0)

func RegisterCronAction(commandStruct interface{}) {
	if reflect.TypeOf(commandStruct).Kind() != reflect.Pointer {
		panic("cron action must be pointer")
	}
	cronCommandStructLists = append(cronCommandStructLists, commandStruct)
}
