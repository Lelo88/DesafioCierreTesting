package products

type repositoryFake struct{
	ProductMem []Product
	Error error
}

func NewFakeRepository() *repositoryFake { 
	return &repositoryFake{}
}

func (r *repositoryFake) GetAllBySeller(sellerID string) ([]Product, error) {
	return r.ProductMem, r.Error
}

func (r *repositoryFake) Reset() {
	r.ProductMem = make([]Product, 0)
	r.Error = nil
}