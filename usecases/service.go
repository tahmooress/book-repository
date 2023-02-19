package usecases

type service struct {
	ur UserRepository
	bs BookSearchService
	bc BookCacheService
}

func New(userRepo UserRepository, bookSearcher BookSearchService,
	bookCacher BookCacheService) Service {
	return &service{
		ur: userRepo,
		bs: bookSearcher,
		bc: bookCacher,
	}
}
