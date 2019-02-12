package model

type BaseModel struct{}

type Modeller interface {
	GetObjArr()         //@todo     public function getObjArr(?array $filter_arr = [], ?array $sort_arr = null, ?int $limit = null, bool $with_trashed = false): Collection;
	GetObjArrPaginate() //@todo     public function getObjArrPaginate(int $per_page, ?array $filter_arr = [], ?array $sort_arr = null, bool $with_trashed = false): LengthAwarePaginator;
}

func (m *BaseModel) AddObj() { // public function addObj(array $data_arr)

}
func (m *BaseModel) SaveObj() { //     public function saveObj(array $data_arr, int $id = null, bool $with_trashed = false)

}

func (m *BaseModel) GetObjByID() { //     public function getObjByID(int $id, bool $with_trashed = false)

}

func (m *BaseModel) DeleteObj() { // public function deleteObj(int $id = null, bool $force_delete = false): bool

}

func (m *BaseModel) RestoreObj() { //     public function restoreObj(int $id = null): bool

}

func (m *BaseModel) CountObjArr() { //     public function countObjArr(?array $filter_arr = [], bool $with_trashed = false): int

}

func (m *BaseModel) DoFilterSortLimit() { //     protected function doFilterSortLimit(?array $filter_arr = [], ?array $sort_arr = null, ?int $limit = null, bool $with_trashed = false)

}

func (m *BaseModel) IsExistObjByID() { //     public function isExistObjByID(int $id, bool $with_trashed = false): bool

}
func (m *BaseModel) shouldInstantiate() { //     private function shouldInstantiate(bool $should, $primary_key_variable = null)

}

func (m *BaseModel) readOnlyGuardian() { //     private function readOnlyGuardian()

}
