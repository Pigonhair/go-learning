# Golang 연습 2021-12-31~
## **Golang 문법**   

> ## 2022-01-11정리   
>* 초기화되지 않은 map에 우리는 어떤 값을 넣을수가 없다. 따라서 비어있는 empty map을 만들고 싶으면 
>   ```go
>      var results = map[string]string{}
>* 이런식으로 만들어 주거나, make()함수를 사용하여 비어있는 empty map을 만든다
>   ```go
>      var results = make(map[string]string)

>* Goroutines이란?(Go를 쓰는이유(멀티스레드 개념). 가장중요★)   
> -기본적으로 다른 함수와 동시에 실행시키는 함수   
> -예를 들어, sexyCount() 라는 메서드를 실행시킬때 앞에 go만 붙여서 실행시켜도 다른 함수와 동시에 실행되서 작업시간을 훨씬 단축시킬 수 있다.
> 그런데, 만약 내가 어떤 작업을 동시에 진행하려고 하면 메인 함수는 다른 goroutines를 기다려 주지 않는다.
> 그래서 메인 함수가 끝나면, goroutines도 소멸됨   
> 
> 예를 들어보겠다.   
>    
> <img src="https://user-images.githubusercontent.com/75151693/148900344-50a3342d-74c8-491c-aa6b-7950b7acff79.png" width="20%" height="10%">
>
>　　　　위의 사진처럼 원래는 하나의 메서드가 for문을 다돌고나서 그다음 메서드가 순차적으로 실행되는것이 일반적이다.      
> 
>　　　　하지만 두개의 메서드중 하나의 메서드에 go를 붙여서 실행하게 되면,
>     
> <img src="https://user-images.githubusercontent.com/75151693/148901064-1ea036b4-f537-4dca-aca6-0ba11bf6b232.png" width="20%" height="10%">
>
>　　　　이렇게 두개의 메서드가 동시에 for문을 돌면서 실행된다.
>
>　　　　그런데 여기서 궁금점! 두개의 메서드에 다 go를 붙여서 실행하면 어떻게 될까?
> 
> <img src="https://user-images.githubusercontent.com/75151693/148901904-0b0d46f5-dab8-4067-91b3-547ca5197562.png" width="20%" height="10%">
> 
>　　　　아무것도 뜨지 않는다. 이유가 뭘까? 메인함수는 다른 goroutines을 기다려 주지 않아서, 메서드가 둘다 goroutines라면, 
>　　　　메인함수가 할 게 없어진다. 따라서 프로그램도 종료된다.    
>
>　　　　하지만 두번째에선 왜 실행이 되었을까? 그 이유는 메인 함수가 goroutines이 아닌 flynn을 카운팅하고 있기 때문이다.   
>　　　　그런데 두개의 메서드를 모두 goroutines로 실행시키는 방법이 있다.
>   ```go
>   func main() {
>       go sexyCount("kunoo")
>       go sexyCount("flynn")
>       time.Sleep(time.Second * 5)
>       }
>   ```
>　　　　이런식으로 time.Sleep을 걸게되면 해당시간동안은 goroutines는 살아있게 되고, 그 이후엔 메인 함수는 종료된다. 
>　　　　또 그와 동시에 goroutines도 같이 끝이난다.
