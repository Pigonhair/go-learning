# Golang 연습 2021-12-31~
## **Golang 문법**   

> ## 2022-01-03정리   
>* &는 메모리주소를 참조함(ex. b:= &a 하면 a의 메모리주소값을 참조한다.)
>* *는 살펴보기 기능(ex. *b하면 a의 메모리 주소에 들어있는 실제 값을 참조한다.)      
>* Array만드는 형식 ex.
>   * ```go   
>       names := []string{"K", "L", "P"}
>       names = append(names, "J") // List에 추가하기
>       kunwoo := map[string]string{"name": "kunwoo", "age": "29"} // key, value형 리스트 
>       ```
> ### Constructor 생성법
>* Go는 private를 보호하기 위해 소문자는 export할수없게 했다. 따라서 대문자로 시작해야 export 할 수있다.
>   * ```go   
>       package accounts
>       
>       // Account struct
>       type Account struct {
>	         Owner   string
>	         Balance int
>        }
>        ```
>   * ```go
>        package main
>        
>        import "github.com/Pigonhair/go-learning/accounts"
>        
>        func main() {
>            account := banking.Account{Owner:"kunwoo", Balance: 1000}
>        }
>        ``` 
>* 하지만 이런식으로 만들게 되면 private하지 않기 때문에 사용자가 임의로 값을 변경시킬수가 있다. 
>* 따라서 struct을 소문자로 다시 바꾸고, 아래와 같이 함수를 만들어 사용한다
>   * ```go
>        // NewAccount creates Account
>        func NewAccount(owner string) *Account {
>           account := Account{owner: owner, balance: 0}
>           return &account
>        }
>     ```
>* 값을 변경하고 싶을땐, method를 만들어서 변경시키자.
>   * ```go
>        // Deposit x amount on your account
>        func (a *Account) Deposit(amount int) { //*을 붙여준 이유는 Deposit메서드를 호출한 account를 사용하라는 의미에서
>        a.balance += amount
>        } 
>     ```
>   * ```go
>        // Balance of your account
>        func (a Account) Balance() int {
>           return a.balance
>         }
>     ```
>* 그리고 이렇게 main에서 호출해서 사용한다.
>   * ```go
>        func main() {
>           account := accounts.NewAccount("kunwoo")
>           account.Deposit(10)
>           fmt.Println(account.Balance())
>        }
>     ```
> ### error handling
>* Go에는 try catch같은 예외처리가 없다. 따라서 메서드에서 error처리를 해줘야한다.   
> 
> 예시)
>   * ```go
>        var errNoMoney = errors.New("Can't withdraw because you are poorT_T")
>        
>        // Withdraw x amount from your account
>        func (a *Account) Withdraw(amount int) error {
>           if a.balance < amount {
>               return errNoMoney
>           }
>           a.balance -= amount
>           return nil  //error도 bool처럼 두가지 value가 있는데, errors와 nil이 있음
>        }
>* main.go에서 error적용 
>   * ```go
>        func main() {
>           err := account.Withdraw(20)
>           if err != nil {
>               fmt.Println(err)
>           }
>        }
> ## 2022-01-04정리
>* Go가 struct를 자동호출 해주는 method가 있는데, 그중 하나가 'String()'이다   
> 예시)
>   * ```Go
>        func (a Account) String() string {
>           return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
>        }
>     ```   
> ## 2022-01-11정리
>* Map을 이용하여 dict에 key,value를 넣고 search로 값을 검색했을때, 해당하는 값이 없을경우 err메세지 출력
>   * ```go
>        package mydict
>       
>        import "errors"
>        
>        // Dictionary type
>        type Dictionary map[string]string
>        var errNotFound = errors.New("Not Found")
>         
>        // Search for a word
>        func (d Dictionary) Search(word string) (string, error) {
>           value, exists := d[word]
>	          if exists {
>      	      return value, nil
> 	        }
>        return "", errNotFound
>        }
>     ```   
>   * ```go
>        package main
>
>        import (
>	        "fmt"
>
>	        "github.com/Pigonhair/go-learning/mydict"
>        )
>
>        func main() {
>	          dictionary := mydict.Dictionary{"first": "First word"}
>        	  definition, err := dictionary.Search("second")
>         	if err != nil {
>		          fmt.Println(err)
>   	      } else {
>		          fmt.Println(definition)
>   	      }
>        }
>     ```   
>* 실행결과   
>![go-learning1](https://user-images.githubusercontent.com/75151693/148868662-0ee7077b-ce37-4ad4-8876-5b2370cfd6fa.jpg)   
>* dict 찾기, 추가, 업데이트, 삭제   
>    err관련 변수
>   * ```go
>        var (
>             errNotFound   = errors.New("Not Found")
>             errCantUpdate = errors.New("Cant update non-existing word")
>             errWordExists = errors.New("That word already exists")
>             )
>    찾기
>   * ```go
>        // Search for a word
>        func (d Dictionary) Search(word string) (string, error) {
>            value, exists := d[word]
>            if exists {
>                 return value, nil
>            }
>            return "", errNotFound
>        } 
>    추가
>   * ```go
>        // Add a word to the dictionary
>        func (d Dictionary) Add(word string, def string) error {
>           _, err := d.Search(word) // 추가하려는 단어가 이미 있으면 err
>           switch err {
>           case errNotFound:
>		            d[word] = def
>           case nil:
>               return errWordExists
>           }
>	          return nil
>        }
>    수정
>   * ```go
>        // Update a word
>        func (d Dictionary) Update(word, definition string) error {
>           _, err := d.Search(word)
>           switch err {
>           case nil:
>               d[word] = definition
>           case errNotFound:
>               return errCantUpdate
>           }
>        return nil
>        }  
>    삭제
>    * ```go
>         // Delete a word
>         func (d Dictionary) Delete(word string) {
>             delete(d, word)
>         }      
