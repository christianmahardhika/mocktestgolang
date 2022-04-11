# Unit testing dengan mock repository pada golang

Halo ini adalah tulisan pertama saya di Medium semoga betah menbacanya :D. 

Kali ini kita akan membahas tentang *unit test* di Golang. Unit test merupakan hal penting dalam men-*develop* sebuah aplikasi, dengan unit test kita dapat mengetahui aplikasi kita berjalan sesuai dengan ekspektasi atau malah banyak error yang dihasilkan :(. 

Unit Test juga dapat diintegrasikan dengan **CICD** (*Continous Integration Continous Deployment*) agar kualitas code kita terjaga dan tetap sesuai dengan ekspektasi bisnis model yang direncanakan. untuk proses integrasi ke **CICD** akan dibahas ditulisan selanjutnya ya ;)

Udah ya basa-basinya..., Yuk *cusss* langsung aja

## Arsitektur Aplikasi
Dalam aplikasi ini terdapat beberapa layer yang memisahkan *entity* *bisnis logic* dengan *repository* (data) 

```text
 ./service
    model.go
    respository.go
    usecase.go
main.go
```

### Model
Model berisi *entity* atau data model pada sebuah domain. *Entity* pada tutorial ini berupa objek yang berisi kumpulan *data structure* yang merepresentasikan domain Aplikasi Todo. 

*Gampangnya* model disini jadi *column* atau *field* pada database/rest-api. 

Untuk lebih jelasnya bisa kepo [kesini](https://www.tutorialspoint.com/dbms/er_model_basic_concepts.htm) atau [sini](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

kita buat modelnya dulu ```model.go```
```go
package service

type Todo struct {
  ID int
  Title string
}

type TodoDetail struct {
  ID int
  TodoID int
  Item string
}

type TodoAll struct {
  todo Todo
  todoDetail []*TodoDetail
}
```

Model yang kita buat terdiri dari *entity* **Todo**, **TodoDetail**, dan **TodoAll**

### Repository

Bertugas sebagai sumber data di aplikasi Todo ini. Kita bisa berinteraksi dengan database melalui Repository ini. Kalo pengen tau mendalam tentang apa itu repository bisa [kemari](https://medium.com/@Dewey92/repository-pattern-what-e47ddee3364d) 

yuk kita buat repositorynya ```repository.go```

```go
package service

type Repository interface {
  CreateTodo(todo *Todo) error 
  CreateTodoDetail(todoDetail *TodoDetail) error
  FindTodoById(id int) (*Todo, error)
  FindTodoDetailById(id int) ([]*TodoDetail, error)
}
```

Tutorial kali ini kita cuma define [*contract*](https://refactoring.guru/design-patterns/abstract-factory) atau *interface*-nya aja ya. Lengkapnya di tutorial selanjutnya ;)

### Use Case

Nah disini kita jabarin tuh aplikasinya mau ngapain aja. Jadi *use case* disini merupakan representasi dari bisnis logic dari aplikasi (dalam hal ini Aplikasi Todo)

kita *namain* ```usecase.go```
```go
package service

import(
  "strconv"
)

type UseCase interface {
  SaveTodo(numberOfItems int) (string, error)
}

func NewUseCase(repo Repository) UseCase {
  return &useCase{repo: repo}
}

type useCase struct {
  repo Repository
}

// Menyimpan Todo List berdasarkan jumlah detail item
func (uc *useCase) SaveTodo(numberOfItems int) (string, error){
  todo := Todo{
    Title: "this is title",
  }
  err := uc.repo.CreateTodo(&todo)
  if err != nil {
    return "", err
  }

  for i := 0; i < numberOfItems; i++ {
    todoDetail := TodoDetail{
      ID: todo.ID,
      Item: "item "+ strconv.Itoa(i),
    }
    err := uc.repo.CreateTodoDetail(&todoDetail)
    if err != nil {
    return "", err
  }
  }

  return "success", nil
  
}

// Menamplikan Todo List sesuai ID
func (uc *useCase) GetTodoById(id int) (*TodoAll, error){
  res, err := uc.repo.FindTodoDetailById(id)
  if err != nil {
    return nil, err
  }
  resTodo, err := uc.repo.FindTodoById(res[0].ID)
  if err != nil {
    return nil, err
  }
  var todoResult TodoAll
  todoResult.todoDetail = res
  todoResult.todo = *resTodo

  return &todoResult, nil
  
}
```

Aplikasi ini bisa *nyimpen* dan *tampilin* **Todo List**

## Testing Usecase

Nah sampe juga di Part yang ditunggu. Disini kita akan pake [testify](https://github.com/stretchr/testify) karena lengkap sudah ada package [assertion](https://github.com/stretchr/testify#assert-package) dan [mock](https://github.com/stretchr/testify#mock-package) kamu bisa install dengan cara berikut yaaa!

```bash
go get github.com/stretchr/testify
```

### Assertion
fungsi nya untuk bandingin hasil test sesuai ekspektasi atau engga. untuk yang belum tau apa itu assertion bisa [kunjungin](https://www.tutorialspoint.com/software_testing_dictionary/assertion_testing.htm)

contoh assertion

```go
  // assert equality
  assert.Equal(t, 123, 123, "they should be equal")

  // assert inequality
  assert.NotEqual(t, 123, 456, "they should not be equal")

  // assert for nil (good for errors)
  assert.Nil(t, object)

  // assert for not nil (good when you expect something)
  if assert.NotNil(t, object) {

    // now we know that object isn't nil, we are safe to make
    // further assertions without causing any errors
    assert.Equal(t, "Something", object.Value)

  }
```

### Mock
Kalo mau ngelakuin **unit test** kan ga mungkin kita test di ```database production```, bisa aja sih bikin/pisahin jadi ```database testing``` tapi boros banget sama *resource* :'( 

Jadi kita pake Mock ini buat *niruin* data di dalam database

### Implement Mock Repository

Kita bikin dulu "seakan-akan" MockRepository ini konek ke MySQL/postgres/mongodb etc...

```repository.go```

```go
package service

import(
  "errors"
  "github.com/stretchr/testify/mock"
)
type TestRepositoryMock struct {
	Mock mock.Mock
}

func (repository *TestRepositoryMock) CreateTodo(todo *Todo) error {
	arguments := repository.Mock.Called(todo)
	if arguments.Get(0) == nil {
		return errors.New("Error CreateTodo")
	} else {
		return nil
	}
}

func (repository *TestRepositoryMock) CreateTodoDetail(todo *TodoDetail) error {
	arguments := repository.Mock.Called(todo)
	if arguments.Get(0) == nil {
		return errors.New("Error CreateTodoDetail")
	} else {
		return nil
	}
}

func (repository *TestRepositoryMock) FindTodoById(id int) (*Todo, error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New("Error")
	} else {
    todo := arguments.Get(0).(*Todo)
		return todo, nil
	}
}

func (repository *TestRepositoryMock) FindTodoDetailById(id int) ([]*TodoDetail, error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New("Error")
	} else {
		todo := arguments.Get(0).([]*TodoDetail)
		return todo, nil
	}
}
```

Pertama kita import dulu package ```github.com/stretchr/testify/mock```, terus kita bikin fungsi-fungsi sesuai dengan *contract* [repository di atas](###repository)

### Implement Test Scenario

Nah *kelar bikin* repository mock nya kita bikin skenario test.

kita kasih nama ```service_usecase_test.go``` harus ada ```_test.go``` dibelakang nama filenya ya biar nanti golang test bisa bedain file antara bisnis logic dan skenario test

```go
package service

import(
  "testing"
  "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repository = &TestRepositoryMock{Mock: mock.Mock{}}
var testService = useCase{repo: repository}

func TestService_SaveTodoSuccess(t *testing.T) {
  // Mock Entity Todo
  todo := Todo{
    Title: "this is title",
  }

  // Mock Entity TodoDetail
  todoDetail1 := TodoDetail{
      ID: todo.ID,
      Item: "item 0",
    }
  todoDetail2 := TodoDetail{
      ID: todo.ID,
      Item: "item 1",
    }
  todoDetail3 := TodoDetail{
      ID: todo.ID,
      Item: "item 2",
    }
  repository.Mock.On("CreateTodo", &todo).Return(todo)
  repository.Mock.On("CreateTodoDetail", &todoDetail1).Return(todoDetail1)
  repository.Mock.On("CreateTodoDetail", &todoDetail2).Return(todoDetail2)
  repository.Mock.On("CreateTodoDetail", &todoDetail3).Return(todoDetail3)
  result, err := testService.SaveTodo(3)
  assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestService_FindTodoSuccess(t *testing.T) {
  // Mock Entity Todo
  todo := Todo{
    ID: 1,
    Title: "this is title",
  }

  // Mock Entity TodoDetail Array
  todoDetail := []*TodoDetail{
    {
      ID: 1,
    TodoID: todo.ID,
      Item: "item 1",
    },
  {
      ID: 2,
    TodoID: todo.ID,
      Item: "item 2",
    },
  {
      ID: 3,
    TodoID: todo.ID,
      Item: "item 3",
    },
  }

  var todoAll TodoAll
  todoAll.todo = todo
  todoAll.todoDetail = todoDetail
  
  repository.Mock.On("FindTodoDetailById", 1).Return(todoDetail)
  repository.Mock.On("FindTodoById", 1).Return(&todo)
  result, err := testService.GetTodoById(1)
  assert.Nil(t, err)
	assert.NotNil(t, result)
  assert.Equal(t, &todoAll, result, "the result should be equal")
}

func TestService_FindTodoFailed(t *testing.T) {
  repository.Mock.On("FindTodoDetailById", 5).Return(nil)
  repository.Mock.On("FindTodoById", 5).Return(nil)
  result, err := testService.GetTodoById(5)
  assert.Nil(t, result)
	assert.NotNil(t, err)
}
```

Nah saya bakal jelasin dikit tentang apa aja yang terjadi di ```service_usecase_test.go```

Pertama kita initiate dulu ```RepositoryMock``` biar ```usecase``` kenal sama "database tiruan" yang udah kita bikin

```go
var repository = &TestRepositoryMock{Mock: mock.Mock{}}
var testService = useCase{repo: repository}
```
#### Skenario test
Disini ada 6 skenario yang dijalankan kalian bisa bikin skenario sendiri sesuai kebutuhan project

```text
TestService_SaveTodoSuccess
TestService_SaveTodoSuccess 
TestService_FindTodoSuccess
TestService_FindTodoSuccess
TestService_FindTodoFailed
TestService_FindTodoFailed
```

#### Code ini ngapain sih?
```go
repository.Mock.On("CreateTodoDetail", &todoDetail1).Return(todoDetail1)
```
Code tersebut memanggil *contract repsository* ```CreateTodoDetail``` dengan kondisi jika dapet input ```&todoDetail1``` akan mengembalikan nilai ```todoDetail1```. 

Jika ```Return(nil)``` berarti "database" tidak menampilkan apa-apa seakan data tidak ditemukan di database.

```go
repository.Mock.On("FindTodoById", 10).Return(nil)
```

#### Loop test mock data
Karena pada usecase *si-aplikasi* ngelakuin looping terhadap fungsi ```CreateTodoDetail```, jadi harus bikin **Mock data TodoDetail** yang jumlahnya sesuai dengan jumlah loop yang dilakukan di aplikasi (pada kasus kali ini dibikin 3 kali loop)

Mock yang dibikin sejumlah loop yang dilakukan
```go
repository.Mock.On("CreateTodoDetail", &todoDetail1).Return(todoDetail1)
repository.Mock.On("CreateTodoDetail", &todoDetail2).Return(todoDetail2)
repository.Mock.On("CreateTodoDetail", &todoDetail3).Return(todoDetail3)
```

Fungsi loop pada Use Case

```go
for i := 0; i < numberOfItems; i++ {
    todoDetail := TodoDetail{
      ID: todo.ID,
      Item: "item "+ strconv.Itoa(i),
    }
    err := uc.repo.CreateTodoDetail(&todoDetail)
    if err != nil {
    return "", err
  }
  }
```


### Menjalankan Test

Kamu bisa jalanin aplikasinya dengan command

```bash
go test ./service/... -v -cover 
```
```-v``` berfungsi untuk melihat detail skenario yang kita testing

```-cover``` berfungsi untuk melihat coverage testing 

nanti kalo berhasil hasilnya bakal kaya gini:
```bash
go test ./service/... -v -cover
=== RUN   TestService_SaveTodoSuccess
--- PASS: TestService_SaveTodoSuccess (0.00s)
=== RUN   TestService_FindTodoSuccess
--- PASS: TestService_FindTodoSuccess (0.00s)
=== RUN   TestService_FindTodoFailed
--- PASS: TestService_FindTodoFailed (0.00s)
PASS
coverage: 82.1% of statements
ok      github.com/christianmahardhika/mocktestgolang/service   0.020s  coverage: 82.1% of statements
```

## Penutup
code lengkapnya bisa clone disini ya
https://github.com/christianmahardhika/mocktestgolang/


Kamu bisa coba dengan skenario testing / usecase yang sesuai dengan project kamu ;)

Setelah tutorial unit-test ini akan dibuat tutorial Integration-Test yang pastinya bakal lebih seru :D



Happy Testing!!! :D
