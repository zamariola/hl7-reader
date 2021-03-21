package hl7reader

// Fields is a slice of fields
type Fields []Field

// GetField is used to get the field at the given index.
func (f Fields) GetField(idx int) (Field, bool) {
	if idx >= len(f) {
		return nil, false
	}
	return f[idx], true
}

// GetComponent is used to get the component at the given index.
func (f Fields) GetComponent(fieldIdx, compIdx int) (Component, bool) {
	if field, ok := f.GetField(fieldIdx); ok {
		return field.GetComponent(compIdx)
	}
	return nil, false
}

// GetSubComponent is used to get the component at the given index.
func (f Fields) GetSubComponent(fieldIdx, compIdx, subCompIdx int) (SubComponent, bool) {
	if field, ok := f.GetField(fieldIdx); ok {
		return field.GetSubComponent(compIdx, subCompIdx)
	}
	return nil, false
}

func newFields(compSep, subCompSep, repeat, escape byte, data []byte) Fields {
	var (
		fields Fields
		start  int
	)
	for i := range data {
		if data[i] == repeat {
			fields = append(fields, newField(compSep, subCompSep, escape, data[start:i]))
			start = i + 1
		}
		if i == len(data)-1 {
			fields = append(fields, newField(compSep, subCompSep, escape, data[start:]))
		}
	}
	return fields
}
