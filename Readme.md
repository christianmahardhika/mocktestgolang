# Implement Graceful Shutdown pada Golang dengan Fiber Http Server

## **Pendahuluan**

Kita perlu untuk memastikan aplikasi sudah menyelesaikan setiap request sebelum aplikasi tersebut restart, scale down, atau mati agar user *experience tetap terjaga*. Pada post kali ini saya ingin menjelaskan bagaimana cara menjalaknan proses tersebut dengan graceful shutdown yang terintegrasi dengan signal shutdown process. *Gracefull shutdown* adalah sebuah kondisi dimana sebuah sistem dimatikan secara sengaja oleh user / maintener melalui rangkaian proses yang aman sehingga data / proses yang sedang diolah oleh sistem tidak rusak.

## **Signal Shutdown**

Pada saat sebuah sistem akan dimatikan, sistem operasi akan mengirimkan sinyal yang menandakan bahwa proses dari sistem akan berakhir. Beberapa tanda atau signal dalam proses shutdown pada operating system dapat dilihat seperti dabel dibawah:

```text
Signal	    Value	Action	Comment

SIGTERM	    15	    Term	Termination signal
SIGINT  	2	    Term	Famous CONTROL+C interrupt from keyboard
SIGQUIT	    3	    Core	CONTRL+\ or quit from keyboard
SIGABRT	    6	    Core	Abort signal
SIGKILL	    9	    Term	Kill signal
```

Penjelasan lebih lanjut terkait tentang signal OS dapat *mampir* ke sumber bawah ini:

[sumber 1](https://github.com/drbeco/killgracefully) | [sumber 2](https://man7.org/linux/man-pages/man7/signal.7.html)

## **Urutan Mematikan Sistem**

Aplikasi ini terhubung dengan beberapa *service* yaitu `database` sebagai peyimpanan data, `http server` untuk men-*delivery* aplikasi ke *client*

Dalam kasus aplikasi Todo, berikut urutan dalam mematikan aplikasi:

- Matikan *incoming request* dari `http server`
- Tunggu proses yang sedang berjalan setelah *request* dimatikan
- Matikan koneksi `database`
- Aplikasi mati / berakhir
## **Code**

 Ketika signal yang menginformasikan bahwa aplikasi akan mati, proses ***graceful shutdown*** akan dimulai. Aplikasi akan melakukan *listen* terhadap `os signal`,dengan memanfaatkan fitur `goroutines` kita dapat mengimplementasikan ***graceful shutdown*** 

*code* dapat dilihat seperti dibawah ini:

`main.go`

```go
package main

import (
	"github.com/christianmahardhika/mocktestgolang/server"
	"github.com/gofiber/fiber/v2"
)

var FiberApp *fiber.App

func init() {

	dbString := "mongodb://root:root@localhost:27017"
	dbName := "mocktestgolang"
	FiberApp = server.InitiateServer(dbString, dbName)
}

func main() {
	port := "8080"

	server.ShutdownApplication(FiberApp)

	server.StartApplication(FiberApp, port)

	server.CleanUpApplication()
}

```

`runner.go`

```go
package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func StartApplication(r *fiber.App, port string) {

	// Start the server
	err := r.Listen(":" + port)
	if err != nil {
		panic(err)
	}

}

func ShutdownApplication(r *fiber.App) {

	// Implement graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		_ = <-ctx.Done()
		log.Println("Gracefully shutting down...")
		_ = r.Shutdown()
		stop()
	}()
}

func CleanUpApplication() {
	log.Println("Running cleanup tasks...")
	// wait 2 seconds for the server to shutdown
	time.Sleep(2 * time.Second)
	context.WithTimeout(context.Background(), 5*time.Second)
	DBCloseConnection()
	log.Println("Finish cleanup tasks...")
}
```

`database.go`

```go
package server

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbConn *mongo.Database

func InitDB(dbString string, dbName string) {
	// Initialize the database
	var ctx = context.Background()
	clientOptions := options.Client()
	clientOptions.ApplyURI(dbString)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	dbConn = client.Database(dbName)
}

func GetDBConnection(dbString string, dbName string) *mongo.Database {
	if dbConn == nil {
		InitDB(dbString, dbName)
		return dbConn
	} else {
		return dbConn
	}
}

func DBCloseConnection() {
	if dbConn != nil {
		dbConn.Client().Disconnect(nil)
	}
	log.Println("database connection closed")
}
```

## **Implementasi selanjutnya**

Jika aplikasi terhubung dengan koneksi selain `databas`e misal `message broker` atau koneksi ke *service* lain bisa dimasukan juga ke prosedur / tahapan ***Graceful Shutdown*** aplikasi sesuai dengan kebutuhan.

## **Kesimpulan**

Berikut repository yang bisa teman-teman clone jika ingin eksplore lebih dalam atau mau *kulik-kulik* sendiri terkait dengan ***Graceful Shutdown***. Sampai jumpa pada artikel selanjutnya tentang Software Engineer ing ;D