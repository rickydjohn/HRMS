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
