package todo

import (
    "io/ioutil"
    "encoding/json"
)

type Item struct {
    Text string
    Priority int
}

func SaveItems(filename string, items []Item) error {
    b, err := json.Marshal(items)
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(filename, b, 0644)

    return nil
}

func ReadItems(filename string) ([]Item, error) {
    b, err := ioutil.ReadFile(filename)
    if err != nil {
        return []Item{}, err
    }

    var items []Item
    if err := json.Unmarshal(b, &items); err != nil {
        return []Item{}, err
    }

    return items, nil
}

func (i *Item) SetPriority(pri int) {
    pri = 3  // Default
    if 5 >= pri  && pri >= 0 {
        i.Priority = pri
    }
}
