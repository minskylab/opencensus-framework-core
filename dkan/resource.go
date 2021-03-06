package dkan

type Resource struct {
	id     string
	limit  int64
	offset int64
}

func ResourceWithID(id string) *Resource {
	return &Resource{
		id: id,
	}
}

func (dkan *Resource) Limit(val int64) *Resource {
	dkan.limit = val
	return dkan
}

func (dkan *Resource) Offset(val int64) *Resource {
	dkan.offset = val
	return dkan
}

func (dkan *Resource) First100() *Resource {
	dkan.limit = 100 // static limit (DKAN API Datastore)
	dkan.offset = 0

	return dkan
}

func (dkan *Resource) NextN(n int64) *Resource {
	dkan.offset = dkan.limit + dkan.offset + 1

	return dkan
}
