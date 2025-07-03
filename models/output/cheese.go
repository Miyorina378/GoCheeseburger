package output

type Cheese struct {
    ID              int    `json:"id"`
    Name            string `json:"name"`
    OriginCountryID int    `json:"origin_country_id"`
    CheeseType      string `json:"cheese_type"`
    Description     string `json:"description"`
}