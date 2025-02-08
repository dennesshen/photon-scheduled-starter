package schedule

import (
	"context"
	"reflect"
	
	"github.com/dennesshen/photon-core-starter/log"
	"github.com/robfig/cron/v3"
)

var cronServer *cron.Cron

func Start(ctx context.Context) error {
	log.Logger().Info(ctx, "Start cron server")
	
	cronServer = cron.New()
	for _, commandStruct := range cronCommandStructLists {
		
		commandStructType := reflect.TypeOf(commandStruct).Elem()
		commandStructValue := reflect.ValueOf(commandStruct).Elem()
		
		for i := 0; i < commandStructValue.NumField(); i++ {
			commandType := commandStructType.Field(i)
			if commandType.Type.Kind() != reflect.Func || commandType.Type != reflect.TypeOf(func() {}) {
				continue
			}
			spec := commandType.Tag.Get("spec")
			
			entryId, err := cronServer.AddFunc(spec, func() {
				defer func() {
					if r := recover(); r != nil {
						log.Logger().Error(context.Background(),
							"Cron job panic", "job_name", commandType.Name, "error", r)
					}
				}()
				commandStructValue.Field(i).Interface().(func())()
			})
			if err != nil {
				return err
			}
			log.Logger().Info(ctx, "Add cron job", "spec", spec, "entryId", entryId, "command", commandType.Name)
		}
		
	}
	
	cronServer.Start()
	return nil
}

func Shutdown(ctx context.Context) error {
	cronServer.Stop()
	log.Logger().Info(ctx, "Stop cron server")
	return nil
}
