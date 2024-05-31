### 1. "boss all homestay room"

1. route definition

- Url: /travel/v1/homestay/businessList
- Method: POST
- Request: `BusinessListReq`
- Response: `BusinessListResp`

2. request definition



```golang
type BusinessListReq struct {
	LastId int64 `json:"lastId"` // ID of the last item fetched
	PageSize int64 `json:"pageSize"` // Number of items to fetch
	HomestayBusinessId int64 `json:"homestayBusinessId"` // Store ID to filter by
}
```


3. response definition



```golang
type BusinessListResp struct {
	List []Homestay `json:"list"` // List of homestays
}
```

### 2. "guess you like homestay room"

1. route definition

- Url: /travel/v1/homestay/guessList
- Method: POST
- Request: `GuessListReq`
- Response: `GuessListResp`

2. request definition



```golang
type GuessListReq struct {
}
```


3. response definition



```golang
type GuessListResp struct {
	List []Homestay `json:"list"` // List of guessed homestays
}
```

### 3. "homestay room detail"

1. route definition

- Url: /travel/v1/homestay/homestayDetail
- Method: POST
- Request: `HomestayDetailReq`
- Response: `HomestayDetailResp`

2. request definition



```golang
type HomestayDetailReq struct {
	Id int64 `json:"id"` // ID of the homestay
}
```


3. response definition



```golang
type HomestayDetailResp struct {
	Homestay Homestay `json:"homestay"` // Detailed information of a homestay
}

type Homestay struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	SubTitle string `json:"subTitle"`
	Banner string `json:"banner"`
	Info string `json:"info"`
	PeopleNum int64 `json:"peopleNum"` // Number of people the homestay can accommodate
	HomestayBusinessId int64 `json:"homestayBusinessId"` // Store ID
	UserId int64 `json:"userId"` // Host ID
	RowState int64 `json:"rowState"` // 0: Inactive, 1: Active
	RowType int64 `json:"rowType"` // Selling type 0: by room, 1: by person
	FoodInfo string `json:"foodInfo"` // Meal standards
	FoodPrice float64 `json:"foodPrice"` // Meal price
	HomestayPrice float64 `json:"homestayPrice"` // Homestay price
	MarketHomestayPrice float64 `json:"marketHomestayPrice"` // Market price of the homestay
}
```

### 4. "homestay room list"

1. route definition

- Url: /travel/v1/homestay/homestayList
- Method: POST
- Request: `HomestayListReq`
- Response: `HomestayListResp`

2. request definition



```golang
type HomestayListReq struct {
	Page int64 `json:"page"` // Page number
	PageSize int64 `json:"pageSize"` // Number of items per page
}
```


3. response definition



```golang
type HomestayListResp struct {
	List []Homestay `json:"list"` // List of homestays
}
```

### 5. "good boss"

1. route definition

- Url: /travel/v1/homestayBussiness/goodBoss
- Method: POST
- Request: `GoodBossReq`
- Response: `GoodBossResp`

2. request definition



```golang
type GoodBossReq struct {
}
```


3. response definition



```golang
type GoodBossResp struct {
	List []HomestayBusinessBoss `json:"list"` // List of homestay business bosses
}
```

### 6. "boss detail"

1. route definition

- Url: /travel/v1/homestayBussiness/homestayBussinessDetail
- Method: POST
- Request: `HomestayBussinessDetailReq`
- Response: `HomestayBussinessDetailResp`

2. request definition



```golang
type HomestayBussinessDetailReq struct {
	Id int64 `json:"id"` // ID of the homestay business
}
```


3. response definition



```golang
type HomestayBussinessDetailResp struct {
	Boss HomestayBusinessBoss `json:"boss"` // Detailed information of the business boss
}

type HomestayBusinessBoss struct {
	Id int64 `json:"id"`
	UserId int64 `json:"userId"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Info string `json:"info"` // Introduction of the homestay owner
	Rank int64 `json:"rank"` // Ranking
}
```

### 7. "business list"

1. route definition

- Url: /travel/v1/homestayBussiness/homestayBussinessList
- Method: POST
- Request: `HomestayBussinessListReq`
- Response: `HomestayBussinessListResp`

2. request definition



```golang
type HomestayBussinessListReq struct {
	LastId int64 `json:"lastId"` // ID of the last item fetched
	PageSize int64 `json:"pageSize"` // Number of items to fetch
}
```


3. response definition



```golang
type HomestayBussinessListResp struct {
	List []HomestayBusinessListInfo `json:"list"` // List of homestay businesses with additional info
}
```

### 8. "homestay comment list"

1. route definition

- Url: /travel/v1/homestayComment/commentList
- Method: POST
- Request: `CommentListReq`
- Response: `CommentListResp`

2. request definition



```golang
type CommentListReq struct {
	LastId int64 `json:"lastId"` // ID of the last comment fetched
	PageSize int64 `json:"pageSize"` // Number of comments to fetch
}
```


3. response definition



```golang
type CommentListResp struct {
	List []HomestayComment `json:"list"` // List of comments
}
```

