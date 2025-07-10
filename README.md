# 🎬 Movie API Documentation

API ini digunakan untuk mengelola data film, termasuk fitur list, create, update, dan delete.

## 📌 Base URL
http://localhost:8080

---

## 🔍 List Movies

**Endpoint**:  
`GET /movies`

Digunakan untuk menampilkan daftar film dengan fitur pagination dan filter.

### Query Parameters:

| Parameter     | Tipe   | Deskripsi                                |
|---------------|--------|------------------------------------------|
| page          | number | Halaman yang ingin ditampilkan           |
| size          | number | Jumlah data per halaman                  |
| search        | string | (Opsional) Pencarian umum                |
| title         | string | (Opsional) Filter berdasarkan judul      |
| description   | string | (Opsional) Filter berdasarkan deskripsi  |
| artists       | string | (Opsional) Filter berdasarkan artis      |
| genres        | string | (Opsional) Filter berdasarkan genre      |

### Contoh Request:
GET /movies?page=1&size=1&title=spongebob&artists=irfan

## ➕ Create Movie
**Endpoint**:  
`POST /movies`

Digunakan untuk menambahkan film baru.

### Request Body:
{
  "title": "ABC",
  "description": "Spongebob Movie",
  "duration": 123,
  "artists": [
    "jajang"
    ],
  "genres": [
    "horror"
    ]
  "file": "localhost:8080/storage/movies/1748094679_SampleVideo_1280x720_1mb.mp4"
}

## ✏️ Update Movie
Endpoint:
`PATCH /movies/:id`

Digunakan untuk mengubah data film berdasarkan ID.

### Request Body:
{
  "title": "ABC",
  "description": "Spongebob Movie",
  "duration": 123,
  "artists": ["jajang"],
  "genres": ["horror"]
}

## ❌ Delete Movie
Endpoint:
`DELETE /movies/:id`

Digunakan untuk menghapus data film berdasarkan ID.

## ➕ Upload Movie
Endpoint:
`POST /movies/upload`

Digunakan untuk melakukan upload file movie berupa mp4.

### Request:
- Content-Type: `multipart/form-data`
- Form field:
  - `file`: file video dengan format `.mp4`
