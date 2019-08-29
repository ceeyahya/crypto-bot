package types

// Currency models the currency type returned by the rest API
type Currency struct {
	Success   bool               `json:"success"`
	Terms     string             `json:"terms"`
	Privacy   string             `json:"privacy"`
	Timestamp int                `json:"timestamp"`
	Target    string             `json:"target"`
	Rates     map[string]float64 `json:"rates"`
}
