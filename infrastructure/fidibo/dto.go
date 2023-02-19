package fidibo

type BooksResponse struct {
	Books Books `json:"books"`
}

type Books struct {
	Hits Hit `json:"hits"`
}

type Hit struct {
	Hits []Obj `json:"hits"`
}

type Obj struct {
	Source Source `json:"_source"`
}

type Publisher struct {
	Title string `json:"title"`
}

type Author struct {
	Name string `json:"name"`
}

type Source struct {
	ImageName  string    `json:"image_name"`
	Publishers Publisher `json:"publishers"`
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Slug       string    `json:"slug"`
	Authors    []Author  `json:"authors"`
}
