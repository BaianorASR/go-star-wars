package entities

// Planet represents a planet in the solar system
type Planet struct {
	ID      string `json:"id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	Climate string `json:"climate" bson:"climate"`
	Terrain string `json:"terrain" bson:"terrain"`
	Films   int    `json:"films" bson:"films"`
}
