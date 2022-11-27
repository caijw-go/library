package base

import (
    "github.com/spf13/viper"
)

var listConfig = make(map[string]*viper.Viper)
var defaultName string

func initConfig(nameSet []string) {
    if len(nameSet) > 0 {
        defaultName = nameSet[0]
        for _, name := range nameSet {
            listConfig[name] = viper.New()
            listConfig[name].SetConfigType("yaml")
            listConfig[name].SetConfigName(name)
            listConfig[name].AddConfigPath("./config")
            err := listConfig[name].ReadInConfig()
            if err != nil {
                panic(err)
            }
        }
        return
    }
    panic("init config nameSet is nil")
}

func Config(name ...string) *viper.Viper {
    viperName := defaultName
    if len(name) > 0 {
        viperName = name[0]
    }

    if i, ok := listConfig[viperName]; ok {
        return i
    }
    return nil
}
