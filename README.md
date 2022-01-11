# Golang 연습 2021-12-31~
## **Golang 문법**   

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

