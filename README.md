# โปรเจกต์ MA Backend Training

โปรเจกต์นี้ถูกออกแบบมาเพื่อศึกษาการทำงานของระบบ Backend เพื่อเรียนรู้ในเรื่องที่จำเป็นสำหรับการเป็น Backend Developer โดยมีการกำหนด Requirement และระยะเวลาในการดำเนินการคือ 30 วัน

## คำอธิบาย

โปรเจกต์นี้เกี่ยวข้องกับการสร้าง Backend ด้วย Golang โดยใช้สถาปัตยกรรม Hexagonal (Port Adapter) ซึ่งประกอบด้วยส่วนหลัก ๆ คือ Repository, Service, และ Handler โดย Handler จะใช้ Fiber ในการทำงาน JWT token จะถูกเข้ารหัสด้วย HS256 และ RS256 และระบบต้องรองรับการกำหนดค่าแบบ runtime environment

## ข้อกำหนด

1. **การตั้งค่าโปรเจกต์**
   - สร้างโปรเจกต์ Backend ด้วย Golang
   - ใช้สถาปัตยกรรม Hexagonal (Port Adapter) แบ่งออกเป็น 3 ส่วนหลัก คือ Repository, Service, และ Handler แต่ละส่วนต้องมี Struct Model ของตัวเอง
   - ใช้ Fiber สำหรับ Handler

2. **การเข้ารหัส JWT Token**
   - เข้ารหัส JWT token ด้วย HS256 และ RS256
   - สามารถเลือก Algorithm ที่ใช้ผ่าน runtime environment

3. **การบันทึกและจัดการข้อผิดพลาด**
   - Handler จะเป็นส่วนเดียวที่สามารถดึงข้อมูลจาก Fiber โดยไม่มีการพิมพ์ logging
   - Service จะจัดการข้อมูลและสร้าง response payload โดยไม่เชื่อมต่อกับ Database, Storage หรือ API ภายนอก และต้องจัดการข้อผิดพลาดจาก Repository และบันทึกข้อผิดพลาดตามความเหมาะสม
   - Repository จะติดต่อกับโมดูลภายนอก เช่น Database, Storage โดยไม่มีการพิมพ์ logging

4. **การจัดการรหัสผ่าน**
   - เก็บรหัสผ่านในรูปแบบ bcrypt
   - สร้างผู้ใช้โดยอัตโนมัติเมื่อเริ่มเซิร์ฟเวอร์ครั้งแรก โดยมีผู้ใช้เริ่มต้นเป็น admin (username: admin, password: p@ssw0rd) หากมีผู้ใช้ admin อยู่แล้วจะไม่สร้างใหม่

5. **การตั้งค่าเซิร์ฟเวอร์**
   - ไม่อนุญาตให้เริ่มเซิร์ฟเวอร์หาก runtime environment ไม่ครบถ้วนหรือไม่ถูกต้อง

6. **ข้อกำหนด API**
   - สร้าง API ให้ครบตามข้อกำหนดด้านล่าง

7. **การให้บริการไฟล์แบบ Static**
   - ให้บริการไฟล์ static ผ่าน endpoint `/files/:filename` จากไดเรกทอรี `static`

8. **เอกสารและการจัดการโค้ด**
   - สร้างเอกสาร Swagger โดยใช้ Swaggo
   - อัปเดตโค้ดตามแต่ละฟีเจอร์ตาม Git Workflow

9. **รองรับ Docker**
   - สร้างแอปพลิเคชันเป็น Docker image และตรวจสอบให้แน่ใจว่าสามารถรันและใช้งานได้

10. **การกำหนดค่า Runtime Environment**
    - ใช้ Viper สำหรับกำหนดค่า runtime environment โดยมีตัวแปรดังนี้:
      - `port`: พอร์ตสำหรับเริ่มเซิร์ฟเวอร์ (string, เช่น "3000")
      - `environment`: โหมดการทำงาน (string, เช่น "production")
      - `dbname`: ชื่อฐานข้อมูล (string, เช่น "dbname")
      - `uri`: URI ของฐานข้อมูล (string, เช่น "http://a.com")
      - `path`: ไดเรกทอรีสำหรับเก็บไฟล์ (string, เช่น "/data")
      - `uri_api_service`: URI ของ API service (string, เช่น "http://a.com")
      - `authorization_api_service`: การอนุญาตของ API service (string, เช่น "Bearer xxx")
      - `jwt_algorithm`: Algorithm สำหรับการเข้ารหัส JWT (string, เช่น "HS256,RS256")
      - `jwt_signature`: ลายเซ็นสำหรับ RS256 (string)
      - `jwt_public_key`: Public key สำหรับ RS256 (string)
      - `jwt_private_key`: Private key สำหรับ RS256 (string)

## API Endpoints

### การรับรองความถูกต้อง

- **SignIn**
  - **Endpoint**: `/api/v1/signin`
  - **Method**: POST
  - **Request Body**:
    - `username` (string, required)
    - `password` (string, required)
  - **Response**:
    - `code` (int)
    - `status` (boolean)
    - `message` (string)
    - `data` (object)

- **SignUp**
  - **Endpoint**: `/api/v1/signup`
  - **Method**: POST
  - **Request Body**:
    - `first_name` (string, required)
    - `last_name` (string, required)
    - `username` (string, required)
    - `password` (string, required)
  - **Response**:
    - `code` (int)
    - `status` (boolean)
    - `message` (string)
    - `data` (object)

### การจัดการผู้ใช้

- **Get Users**
  - **Endpoint**: `/api/v1/users`
  - **Method**: GET
  - **Request Headers**:
    - `Authorization` (JWT, required)
  - **Query Params**:
    - `page` (int, optional, default: 1)
    - `row` (int, optional, default: 10)
    - `keyword` (string, optional)
  - **Response**:
    - `code` (int)
    - `status` (boolean)
    - `message` (string)
    - `data` (object)

- **Update User**
  - **Endpoint**: `/api/v1/user/<user_id>`
  - **Method**: POST
  - **Request Headers**:
    - `Authorization` (JWT, required)
  - **Request Params**:
    - `user_id` (string, required)
  - **Request Body**:
    - `first_name` (string, required)
    - `last_name` (string, required)
    - `username` (string, required)
  - **Response**:
    - `code` (int)
    - `status` (boolean)
    - `message` (string)
    - `data` (object)

- **Delete User**
  - **Endpoint**: `/api/v1/user/<user_id>`
  - **Method**: DELETE
  - **Request Headers**:
    - `Authorization` (JWT, required)
  - **Request Params**:
    - `user_id` (string, required)
  - **Response**:
    - `code` (int)
    - `status` (boolean)
    - `message` (string)
    - `data` (object)

- **Reset Password**
  - **Endpoint**: `/api/v1/user/<user_id>/password`
  - **Method**: POST
  - **Request Headers**:
    - `Authorization` (JWT, required)
  - **Request Params**:
    - `user_id` (string, required)
  - **Request Body**:
    - `password` (string, required)
  - **Response**:
    - `code` (int)
    - `status` (boolean)
    - `message` (string)
    - `data` (object)

### การจัดการไฟล์

- **Upload File**
  - **Endpoint**: `/api/v1/files/upload`
  - **Method**: POST
  - **Request Headers**:
    - `Authorization` (JWT, required)
  - **Request Body**:
    - `files` (file, required)
  - **Response**:
    - `code` (int)
    - `status` (boolean)
    - `message` (string)
    - `data` (object)

- **Get File Lists**
  - **Endpoint**: `/api/v1/files`
  - **Method**: GET
  - **Request Headers**:
    - `Authorization` (JWT, required)
  - **Query Params**:
    - `page` (int, optional, default: 1)
    - `row` (int, optional, default: 10)
    - `keyword` (string, optional)
  - **Response**:
    - `code` (int)
    - `status` (boolean)
    - `message` (string)
    - `data` (object)

- **Get File**
  - **Endpoint**: `/api/v1/image/profile/<filename>`
  - **Method**: GET
  - **Response**: file content

- **Delete File**
  - **Endpoint**: `/api/v1/file/<file_id>`
  - **Method**: DELETE
  - **Request Headers**:
    - `Authorization` (JWT, required)
  - **Request Params**:
    - `file_id` (string, required)
  - **Response**:
    - `code` (int)
    - `status` (boolean)
    - `message` (string)
    - `data` (object)

### การเชื่อมต่อ API Service

- **Get API Service**
  - **Endpoint**: `/api/v1/intregation/post/<post_id>`
  - **Method**: GET
  - **Request Headers**:
    - `Authorization` (JWT, required)
  - **Request Params**:
    - `post_id` (string, required)
  - **Response**:
    - `code` (int)
    - `status` (boolean)
    - `message` (string)
    - `data` (object)

- **Post API Service**
  - **Endpoint**: `/api/v1/intregation/post`
  - **Method**: POST
  - **Request Headers**:
    - `Authorization` (JWT, required)
  - **Request Body**:
    - `title` (string, required)
    - `body` (string, required)
    - `user_id` (string, required)
  - **Response**:
    - `code` (int)
    - `status` (boolean)
    - `message` (string)
    - `data` (object)

## โครงสร้างฐานข้อมูล

### คอลเลคชั่น: Users

```json
{
  "user_id": "string",
  "first_name": "string",
  "last_name": "string",
  "username": "string",
  "password": "string",
  "create_date": "timestamp",
  "create_by": "string",
  "update_date": "timestamp"
}
