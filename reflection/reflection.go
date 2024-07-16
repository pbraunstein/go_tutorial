package reflection

import (
    "reflect"
)

func Info(x interface {}) []string {
    toReturn := make([]string, 0)
    r := reflect.ValueOf(x)
    for i := 0; i < r.NumField(); i++ {
        toReturn = append(toReturn, r.Type().Field(i).Name)
    }
    return toReturn
}
