# IGG-GO-ACL

## Requirement
- ชื่อ Module คือ ```github.com/igeargeek/igg-go-acl```
- Flexible ต่อการใช้งานและใช้งานได้ง่าย
- สามารถกำหนด Rules ได้ผ่าน Constant
- สามารถนำมาใช้งานเป็น Middleware ระหว่าง Route ได้ยกตัวอย่างเช่น ```Permission(["Admin","Seller","..."])``` หรือคัดลอกท่าจาก Adonis ที่เราเคยทำก็ได้
- หรือถ้ามี Solution อื่น ก็สามารถเปลี่ยนท่าได้นะครับ
- มี Unit test
- อัพเป็น tag version เช่น v1.0.0
___
## ใช้ Casbin RESTful Adapter on Fiber Web Framework
https://github.com/prongbang/fiber-casbinrest/blob/master/README.md
- ใช้แบบ Adapter File (แก้ไข Config File ที่ policy.csv)

___
# Starter of Golang Fiber Framework

## How to run this project
1. go mod tidy
2. air run server.go

## Connection
- MongoDB => mongodb+srv://golang-starter:lubgNuy0zTBv2yzT@flowstock.jpkb8.mongodb.net/golang-starter?retryWrites=true&w=majority