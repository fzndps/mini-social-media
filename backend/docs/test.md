## 📝 API Endpoint: `POST /auth/register`

### Deskripsi:

Mendaftarkan user baru ke dalam sistem.

### Request

**URL**
`POST /auth/register`

**Headers**

```http
Content-Type: application/json
```

**Body**

```json
{
  "username": "fizonenda",
  "email": "fizo@mail.com",
  "password": "rahasia123"
}
```

### Response

**201 Created**

```json
{
  "id": 1,
  "username": "fizonenda",
  "email": "fizo@mail.com"
}
```

**400 Bad Request**

```json
{
  "error": "Username or email already exists"
}
```

---

## 📝 API Endpoint: `POST /auth/login`

### Deskripsi:

Autentikasi user dan mengembalikan JWT token untuk digunakan dalam permintaan selanjutnya.

### Request

**URL**
`POST /auth/login`

**Headers**

```http
Content-Type: application/json
```

**Body**

```json
{
  "username": "fizonenda",
  "password": "rahasia123"
}
```

### Response

**200 OK**

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "user": {
    "id": 1,
    "username": "fizonenda",
    "email": "fizo@mail.com"
  }
}
```

**401 Unauthorized**

```json
{
  "error": "Invalid username or password"
}
```

---

## 🔐 Catatan Penting:

* Token bertipe **Bearer** → disimpan oleh frontend, lalu dikirim di setiap permintaan yang butuh autentikasi.
* Contoh Authorization header:

```http
Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
```

---

## 📝 API Endpoint: `POST /posts`

### Deskripsi:

Membuat postingan baru berupa teks dan/atau gambar.

### Request

**URL**
`POST /posts`

**Headers**

```http
Authorization: Bearer <token>
Content-Type: multipart/form-data
```

**Body (form-data)**

* `content` (string, opsional) – isi teks dari postingan
* `image` (file, opsional) – file gambar

### Response

**201 Created**

```json
{
  "id": 10,
  "user_id": 1,
  "content": "Hari ini aku belajar Go lang!",
  "image_url": "https://yourcdn.com/uploads/post10.jpg",
  "created_at": "2025-05-18T14:22:35Z"
}
```

**400 Bad Request**

```json
{
  "error": "Content or image must be provided"
}
```

**401 Unauthorized**

```json
{
  "error": "Unauthorized"
}
```

---

## 📝 API Endpoint: `GET /posts`

### Deskripsi:

Mengambil semua postingan terbaru dari semua user (feed umum).

### Request

**URL**
`GET /posts`

**Headers**

```http
Authorization: Bearer <token>
```

**Query Parameters (opsional)**

* `limit` (int): jumlah post per halaman (default: 10)
* `offset` (int): offset pagination (default: 0)

### Response

**200 OK**

```json
[
  {
    "id": 10,
    "user": {
      "id": 1,
      "username": "fizonenda"
    },
    "content": "Hari ini aku belajar Go lang!",
    "image_url": "https://yourcdn.com/uploads/post10.jpg",
    "likes_count": 5,
    "comments_count": 2,
    "created_at": "2025-05-18T14:22:35Z"
  }
]
```

---

## 📝 API Endpoint: `GET /posts/{id}`

### Deskripsi:

Mengambil detail 1 postingan berdasarkan ID-nya.

### Request

**URL**
`GET /posts/10`

**Headers**

```http
Authorization: Bearer <token>
```

### Response

**200 OK**

```json
{
  "id": 10,
  "user": {
    "id": 1,
    "username": "fizonenda"
  },
  "content": "Hari ini aku belajar Go lang!",
  "image_url": "https://yourcdn.com/uploads/post10.jpg",
  "likes_count": 5,
  "comments": [
    {
      "id": 1,
      "user": {
        "id": 2,
        "username": "teman"
      },
      "comment": "Semangat belajarnya!",
      "created_at": "2025-05-18T15:10:00Z"
    }
  ],
  "created_at": "2025-05-18T14:22:35Z"
}
```

**404 Not Found**

```json
{
  "error": "Post not found"
}
```

---

## 📝 API Endpoint: `DELETE /posts/{id}`

### Deskripsi:

Menghapus postingan berdasarkan ID. Hanya pemilik postingan yang dapat menghapus.

### Request

**URL**
`DELETE /posts/10`

**Headers**

```http
Authorization: Bearer <token>
```

### Response

**200 OK**

```json
{
  "message": "Post deleted successfully"
}
```

**403 Forbidden**

```json
{
  "error": "You are not authorized to delete this post"
}
```

**404 Not Found**

```json
{
  "error": "Post not found"
}
```

---

## 📝 API Endpoint: `POST /posts/{id}/like`

### Deskripsi:

Memberi like pada postingan tertentu. Jika user sudah like, maka akan menghapus like tersebut (toggle).

### Request

**URL**
`POST /posts/10/like`

**Headers**

```http
Authorization: Bearer <token>
```

### Response

**200 OK**

```json
{
  "message": "Post liked"
}
```

**200 OK** (jika sebelumnya sudah like dan like dibatalkan)

```json
{
  "message": "Post unliked"
}
```

**404 Not Found**

```json
{
  "error": "Post not found"
}
```

**401 Unauthorized**

```json
{
  "error": "Unauthorized"
}
```

---

## 📝 API Endpoint: `POST /posts/{id}/comments`

### Deskripsi:

Menambahkan komentar pada postingan tertentu.

### Request

**URL**
`POST /posts/10/comments`

**Headers**

```http
Authorization: Bearer <token>
Content-Type: application/json
```

**Body**

```json
{
  "comment": "Postinganmu keren banget!"
}
```

### Response

**201 Created**

```json
{
  "id": 15,
  "post_id": 10,
  "user": {
    "id": 2,
    "username": "teman"
  },
  "comment": "Postinganmu keren banget!",
  "created_at": "2025-05-18T16:12:45Z"
}
```

**400 Bad Request**

```json
{
  "error": "Comment cannot be empty"
}
```

**404 Not Found**

```json
{
  "error": "Post not found"
}
```

**401 Unauthorized**

```json
{
  "error": "Unauthorized"
}
```

---

## 📝 API Endpoint: `GET /users/{id}`

### Deskripsi:

Mengambil profil user berdasarkan ID.

### Request

**URL**
`GET /users/1`

**Headers**

```http
Authorization: Bearer <token>
```

### Response

**200 OK**

```json
{
  "id": 1,
  "username": "fizonenda",
  "email": "fizo@mail.com",
  "bio": "Belajar backend Golang",
  "profile_picture_url": "https://yourcdn.com/uploads/user1.jpg",
  "followers_count": 10,
  "following_count": 5
}
```

**404 Not Found**

```json
{
  "error": "User not found"
}
```

**401 Unauthorized**

```json
{
  "error": "Unauthorized"
}

```
