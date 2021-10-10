# medcury-quitz001
medcury Developer quitz001

Postman Collection : Quitz001.postman_collection.json
API Document : https://galactic-eclipse-5686.postman.co/workspace/MEDCURY~250dd4ae-e6d6-40f9-bba5-4812fc70fa95/documentation/1580611-b3ce259c-6cc0-4729-827e-b7ad01129480

#Technology
Devlopment Language : Golang
OR Mapping : GORM
Logging : Zero Log
Api Framework : GoFiber
Configuration Utility : Viper
Database : MySQL

#Installation
Install Golang(Download and instruction here) : https://golang.org/doc/install
Set GO111MODULE Environment Variable : on
Pull git to GOPATH (C:\go or canset to Environment Variable ) to folder %GOPATH%\src\medcury\quitz001
Run : run_go_module.bat
Run : run_dev.bat

To do
[]ควรกำหนดให้คนไข้สามารถทำการนัดได้มากที่สุดไม่เกินกี่ slot ต่อวัน หรือหากนัดไปแล้วก็สามารถทำการเลื่อนนัดได้
[]สถานะของการนัดหมาย จะเป็น [A]ctive,[I]nactive,[M]ove,[R]eject
[]คุณหมอควรจะ ยกเลิกได้หากติดเคส หรือติดธุระอื่น
[]PIN Code ควรต้องทำการเข้ารหัสถ้าหากมีเวลาทำ
[]ควรจะมีการกำหนดให้มีการทำการนัดหมายล่วงหน้าได้ไม่เกินกี่วัน เช่น ต้องนัดหมายก่อน 3 วัน ไวกว่านั้นจะจัดคิวไม่ทัน
[]ควรจะกำหนดวันที่นัดหมายไกลที่สุด เช่นเรามากำหนดให้เป็นในแต่ละเดือน เพื่อให้ข้อมูลไม่กว้างจนเกินไป
[]ควรจะแยกแผนกของคุณหมอ หรือรวมไปถึงโรงพยาบาบที่สังกัด เวลาค้นหาจะสามารถค้นหาได้ไวขึ้น
[]ยังไม่มี primarykey code ของคนไข้ตอนนี้ขอใช้หมายเลขโทรศัพท์ไปก่อน ส่วนตัวคิดว่าถ้าไม่ใช่ HN ก็ควรเป็น CID/Passport Code
[]ยอดสรุปทำเผื่อทุกสถานะให้แล้ว

