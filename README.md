# prettystring
make pretty string(like python) golang struct

reflect 를 사용해서 golang object들을 string으로 만들어 줍니다. 

object가 다른 object를 포함 하고 있는 경우 재귀적으로 처리 하며 
이 회수를 제한 하기 위한 인자를 받습니다. 

struct field 중 prettystring의 처리를 생략 하고 싶으면 struct tag 로 `prettystring:"hide"` 을 사용하면 됩니다.

example/main.go 참고 