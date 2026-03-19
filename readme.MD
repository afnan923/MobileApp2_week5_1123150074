# Testing API Golang with Postman
## Struktur Project
catalog_1123150074│
├── main.go
├── go.mod
├── .env
├── .gitignore
│
├── config
│   ├── database.go
│   └── firebase.go
├── images
│   
│
├── models
│   ├── user.go
│   └── product.go
│
├── middleware
│   └── auth_middleware.go
│
├── handlers
│   ├── auth_handler.go
│   └── product_handler.go
│
├── services
│   ├── auth_service.go
│   └── product_service.go
│
├── repositories
│   ├── user_repository.go
│   └── product_repository.go
│
└── routes
    └── router.go

### Penjelasan
| Folder       | Fungsi                                         |
| ------------ | ---------------------------------------------- |
| config       | konfigurasi database dan firebase              |
| models       | struktur data yang terhubung ke tabel database |
| repositories | layer akses database                           |
| services     | business logic aplikasi                        |
| handlers     | menerima HTTP request                          |
| middleware   | autentikasi JWT                                |
| routes       | definisi endpoint API                          |
| images       | image untuk di readme                          |

## API Endpoint
Berikut Endpoint yang tersedia pada sistem.
### Health Check
```
GET /v1/health
```
![alt text](images/healthpostman.png)
![alt text](images/healthweb.png)