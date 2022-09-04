package book

type Service interface {
	FindByID(ID int) (Book, error)
	FindAll() ([]Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
	Delete(ID int, bookRequest BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {

	books, err := s.repository.FindAll()
	return books, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	return s.repository.FindByID(ID)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	discount, _ := bookRequest.Discount.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Discount:    int(discount),
		Rating:      int(rating),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, _ := s.repository.FindByID(ID)
	price, _ := bookRequest.Price.Int64()
	discount, _ := bookRequest.Discount.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Discount = int(discount)
	book.Rating = int(rating)
	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int, bookRequest BookRequest) (Book, error) {
	book, _ := s.repository.FindByID(ID)
	newBook, err := s.repository.Delete(book)
	return newBook, err
}
