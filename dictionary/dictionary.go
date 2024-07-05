package dictionary

const (
    ErrNotFound = DictionaryErr("No such word")
    ErrWordExists = DictionaryErr("Cannot add a definition for existing word")
)

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
    value, ok := d[key]

    if ok {
        return value, nil
    } else{
        return "", ErrNotFound
    }
}

func (d Dictionary) Add(key, value string) error {
    _, ok := d[key]

    if ok {
        return ErrWordExists
    } else {
        d[key] = value
        return nil
    }
}
