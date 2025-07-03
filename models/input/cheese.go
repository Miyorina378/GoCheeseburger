package input

type Cheese struct {
    Name            string `json:"name"`
    OriginCountryID int    `json:"origin_country_id"`
    CheeseType      string `json:"cheese_type"`
    Description     string `json:"description"`
}