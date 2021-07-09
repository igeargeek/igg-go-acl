# IGG-GO-ACL

## Requirement
- ชื่อ Module คือ ```github.com/igeargeek/igg-go-acl```
- Flexible ต่อการใช้งานและใช้งานได้ง่าย
- สามารถกำหนด Rules ได้ผ่าน Constant
- สามารถนำมาใช้งานเป็น Middleware ระหว่าง Route ได้ยกตัวอย่างเช่น ```Permission(["Admin","Seller","..."])``` หรือคัดลอกท่าจาก Adonis ที่เราเคยทำก็ได้
- หรือถ้ามี Solution อื่น ก็สามารถเปลี่ยนท่าได้นะครับ
- มี Unit test
- อัพเป็น tag version เช่น v1.0.0

### Architecture
- เป็น middlware lib
- config โยนค่า roles และ permission เข้าไปใน middleware
- ส่งผ่านความสมารถไปยัง contorller โดยส่งผ่าน context ว่าสมารถทำอะไรได้บ้าง
- ถ้าไม่มีสิทธิ์ return 401