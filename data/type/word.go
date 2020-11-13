package dtype

// Case 用例
type Case struct {
	Ri  string `json:"日" bson:"日"`
	Han string `json:"汉" bson:"汉"`
}

// Mean 日语词句释义
type Mean struct {
	Mean string   `json:"义" bson:"义"`
	Desc []string `json:"注,omitempty" bson:"注,omitempty"`
	Case []Case   `json:"例,omitempty" bson:"例,omitempty"`
}

// Words 日语词句
type Words struct {
	Han      string            `json:"汉,omitempty" bson:"汉,omitempty"`
	Ri       string            `json:"日" bson:"日"`
	Label    []string          `json:"标,omitempty" bson:"标,omitempty"`
	Desc     []string          `json:"注,omitempty" bson:"注,omitempty"`
	Mean     []Mean            `json:"释" bson:"释"`
	Relation map[string]string `json:"联,omitempty" bson:"联,omitempty"`
	ID       string            `json:"id,omitempty" bson:"id"`
}
