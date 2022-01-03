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
