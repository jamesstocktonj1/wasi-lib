package keyvalue

import "encoding/json"

func (b *Bucket) GetDocument(key string, v any) error {
	data, err := b.GetBytes(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (b *Bucket) SetDocument(key string, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return b.SetBytes(key, data)
}
