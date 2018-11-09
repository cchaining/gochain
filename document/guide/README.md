# 개발 가이드

### 목차

* [코딩]
* [커밋]
* [주석](#주석)
    * [한 줄 주석](#OneLine)
    * [여러 줄 주석](#MultiLine)
    * [교육용 안내 주석](#Notice)
    * [패키지 주석](#Package)
    * [Go 파일 주석](#GoFile)

-------------

## 주석

주석은 코드의 이해를 돕기 위한 보조문이다. 주석에는 한 줄 주석과 여러 줄 주석 두 가지가 있다.
주석은 설명하려는 구문에 맞춰 들여쓰기한다. 문장의 끝에 주석을 작성할 경우, 한 줄 주석을 사용하며
공백을 추가한다. 

### OneLine

__[한 줄 주석]__

- 주석 전에 한 줄 띄워준다.
- 들여쓰기하여 코드와 정렬시킨다.

```
[Good]
function someFunction() {

    // code에 관한 주석
    statements
}

```

### MultiLine

__[여러 줄 주석]__

- 주석의 첫줄에 문장을 기입하지 않는다.
- 표시의 정렬을 맞춘다.

```
[Good]
/*  - 주석의 첫줄에 문장이 안옴
 * '*' 표시의 정렬을 맞춘다.
 */
 
```

### Notice

__[교육용 안내 주석]__

- __[필기]__이라는 안내 펫말을 기입해준다.
- 한줄 주석 + 여러줄 주석을 함께 쓴다.

```
//__[필기]__
/*
 * 
 */
 
```

### Package

__[패키지 주석]__

- 패키지의 가장 첫 go 파일의 앞머리에 해당 패키지에 대한 설명을 기재한다.

```
/*
 * core 패키지 :
 */

package core
```

### GoFile

__[Go 파일 주석]__

- 패키지의 가장 첫 go 파일의 앞머리에 해당 패키지에 대한 설명을 기재한다.

```
/*
 * block go 파일 :
 */
 ```