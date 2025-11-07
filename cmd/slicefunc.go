package main

// straight map
func Map[FROM any, TO any](src []FROM, fn func(FROM) TO) []TO {
	if nil == src {
		return nil
	}
	var dest = make([]TO, len(src))
	for i, v := range src {
		dest[i] = fn(v)
	}
	return dest
}

// map and filter (filters out if ok bool = false) )
func MapFilter[FROM any, TO any](src []FROM, fn func(FROM) (TO, bool)) []TO {
	if nil == src {
		return nil
	}
	var dest = make([]TO, 0, len(src))
	for _, v := range src {
		var result, ok = fn(v)
		if ok {
			dest = append(dest, result)
		}
	}
	return dest
}
