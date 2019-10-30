# court-remands


การรัน Intregration Test
ใช้คำสั่ง go test -tags=integration 
การรัน Unit  Test
ใช้คำสั่ง go test -v 
หรือ อยากดู test coverage สามารถใช้คำสั่ง  go test -cover
หรือต้องการเก็บผลการ test (Test Report) สามารกใช้คำสั่ง
go test -v -coverprofile cover.out ./…
go tool cover -html=covere.out -o cover.html
open cover.html
