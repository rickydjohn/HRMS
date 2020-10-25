package db

func (d *db) funcmap(name Queries) (queryFunc, error) {
	fmap := map[Queries]queryFunc{
		PERSONAL: d.personal,
		BANKEDU:  d.bankedu,
		LEAVES:   d.leaves,
	}
	v, ok := fmap[name]

	if !ok {
		return nil, ErrNoQuery
	}
	return v, nil
}

func (d *db) ApiFuncs(val string, fName Queries) ([]byte, error) {
	fmap := map[Queries]apidata{
		API_DESIGNATIONS: d.designations,
	}
	v, ok := fmap[fName]
	if !ok {
		return nil, ErrNoQuery
	}
	data, err := v(val)
	if err != nil {
		return nil, err
	}
	return data, nil
}
