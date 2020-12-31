package sixthquestion

import "strings"

// MapSum структа состоящая из нашей карты
type MapSum struct {
    MapSum map[string]int
}

// BindingFunc связывающая функция
func BindingFunc() MapSum {
    return MapSum{
        MapSum: make(map[string]int),
    }
}

// Insert вставка нужного значения
func (m *MapSum) Insert(key string, val int)  {
    m.MapSum[key] = val
}

// Sum выводит сумму
func (m *MapSum) Sum(prefix string) int {
    var totalSum int
    for key, value := range m.MapSum {
        if strings.HasPrefix(key, prefix) {
            totalSum += value
        }
    }
    return totalSum
}