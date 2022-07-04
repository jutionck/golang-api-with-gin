## Practice API

### Soal

Buatlah sebuah API untuk CRUD Product dengan ketentuan sebagai berikut:

#### Model

##### Category
```go
type Category struct {
	CategoryId string   `json:"id" binding:"required"`
	CategoryName string `json:"name" binding:"required"`
}
```

#### Caregory Data
```json
{
  "message": "OK",
  "data": [
    {
      "id": "C0001",
      "name": "Food"
    },
    {
      "id": "C0002",
      "name": "Drink"
    },
    {
      "id": "C0003",
      "name": "Other"
    }
  ]
}
```

##### Product
```go
type Product struct {
	ProductId string    `json:"id" binding:"required"`
	ProductName string  `json:"name" binding:"required"`
	Category Category   `json:"category" binding:"required"`
	IsActive bool       `json:"is_active"`
}
```

#### Storage
Penyimpanan sementara hanya di `slice`

#### Method
1. GET Category (List)
2. POST Product
3. GET Product (params: id)
4. GET Product By Category (List)
5. PUT Product Category
6. DELETE Product (dengan melakukan update is_active)

#### All Request
```
Use json
```

#### All Response
```json
{
  "messagee": "string",
  "data": interface{}
}
```
`Note`: `interface{}` disini dapat menampung data sesuai yang di request, misalnya product atau catgery. Atau bisa di lihat contoh sebagai berikut:
```json
{
  "messagee": "string",
  "data": [
    {
      "id": "string",
      "name": "string",
      "category": {
        "id": "string",
        "name": "string"
      },
      "is_active": true
    },
    {
      "id": "string",
      "name": "string",
      "category": {
        "id": "string",
        "name": "string"
      },
      "is_active": false
    }
  ]
}
```
