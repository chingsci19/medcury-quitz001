# medcury-quitz001
medcury Developer quitz001

Postman Collection : Quitz001.postman_collection.json<br>
API Document : https://galactic-eclipse-5686.postman.co/workspace/MEDCURY~250dd4ae-e6d6-40f9-bba5-4812fc70fa95/documentation/1580611-b3ce259c-6cc0-4729-827e-b7ad01129480<br>

## Technology
Devlopment Language : Golang<br>
OR Mapping : GORM<br>
Logging : Zero Log<br>
Api Framework : GoFiber<br>
Configuration Utility : Viper<br>
Database : MySQL<br>

## Installation
Install Golang(Download and instruction here) : https://golang.org/doc/install<br>
Set GO111MODULE Environment Variable : on<br>
Pull git to GOPATH (C:\go or canset to Environment Variable ) to folder %GOPATH%\src\medcury\quitz001<br>
Run : run_go_module.bat<br>
Run : run_dev.bat<br>

## To do
[]ควรกำหนดให้คนไข้สามารถทำการนัดได้มากที่สุดไม่เกินกี่ slot ต่อวัน หรือหากนัดไปแล้วก็สามารถทำการเลื่อนนัดได้<br>
[]สถานะของการนัดหมาย จะเป็น [A]ctive,[I]nactive,[M]ove,[R]eject<br>
[]คุณหมอควรจะ ยกเลิกได้หากติดเคส หรือติดธุระอื่น<br>
[]PIN Code ควรต้องทำการเข้ารหัสถ้าหากมีเวลาทำ<br>
[]ควรจะมีการกำหนดให้มีการทำการนัดหมายล่วงหน้าได้ไม่เกินกี่วัน เช่น ต้องนัดหมายก่อน 3 วัน ไวกว่านั้นจะจัดคิวไม่ทัน<br>
[]ควรจะกำหนดวันที่นัดหมายไกลที่สุด เช่นเรามากำหนดให้เป็นในแต่ละเดือน เพื่อให้ข้อมูลไม่กว้างจนเกินไป<br>
[]ควรจะแยกแผนกของคุณหมอ หรือรวมไปถึงโรงพยาบาบที่สังกัด เวลาค้นหาจะสามารถค้นหาได้ไวขึ้น<br>
[]ยังไม่มี primarykey code ของคนไข้ตอนนี้ขอใช้หมายเลขโทรศัพท์ไปก่อน ส่วนตัวคิดว่าถ้าไม่ใช่ HN ก็ควรเป็น CID/Passport Code<br>
[]ยอดสรุปทำเผื่อทุกสถานะให้แล้ว<br>

