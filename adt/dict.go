package adt

type Dict map[interface{}]interface{}

func (d Dict) Contains(k interface{}) (bool, interface{}) {
	val, ok := d[k]
	if ok {
		return true, val
	}
	return false, nil
}

func (d Dict) SetDefault(k interface{}, dv interface{}) interface{} {

	if ok, val := d.Contains(k); ok {
		return val
	}
	d[k] = dv
	return dv
}
