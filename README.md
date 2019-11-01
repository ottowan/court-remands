# court-remands


  [การรัน Intregration Test]

    ใช้คำสั่ง go test -tags=integration 

  [การรัน Unit  Test]
  
    ใช้คำสั่ง go test -v 
    
    หรือ อยากดู test coverage สามารถใช้คำสั่ง  
    go test -cover
    
    หรือต้องการเก็บผลการ test (Test Report) สามารกใช้คำสั่ง
    go test -v -coverprofile cover.out .
    go tool cover -html=cover.out -o cover.html
    open cover.html
    

  [การรัน UAT]
  
    go run main.go
    
    เข้าหน้าเพจ URL : http://localhost
    

  [Reference]
    https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices