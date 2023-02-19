package fidibo

type Book struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source Source  `json:"_source"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type Hits struct {
	Total    int     `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Book  `json:"hits"`
}

type BooksResponse struct {
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
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
	Weight     int       `json:"weight"`
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Slug       string    `json:"slug"`
	Authors    []Author  `json:"authors"`
}
