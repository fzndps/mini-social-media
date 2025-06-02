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
