# GOCHAIN 이란

Go lang으로 구현한 기본적인 블록체인

[자세한 설명](https://github.com/gochain/documents/guide)

# 구현 목적

블록체인의 기본 구조를 파악하여 직접구현하고, 향후 이더리움 블록체인 오픈소스에 기여하기 위한
기본기를 다져놓는다.

# 블록체인 스펙

- UTXO 기반 트렌젝션
- go routine 기반 다중 프로세스 지원
- 채굴 프로세스 / 합의 프로세스
- VOLTDB
- 거래 후 잔여 UTXO SET 지원

# gochain GOPATH 설정 방법
1. 우선 go path를 설정한다.

```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
2. go 디렉토리로 이동한다.
```
cd go
```
3. github 레퍼지토리를 clone 한다.
```
git clone https://github.com/cchaining/gochain.git
```

# gochain 커멘드 실행

```
go run main.go [명령어]
```

# 기여자
@ParkDongJo(Charles)[github](https://github.com/ParkDongJo) 
@myunghui[github](https://github.com/myunghui)

# 참고 코드
[naive coin 깃헙 링크](https://github.com/lhartikk/naivecoin)
<br>
[go lang blockchain core ex](https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc)
<br>
[go lang 문법 쉽게 풀어쓴 사이트](https://gobyexample.com/)
<br>
[go lang tutorial 사이트](https://www.tutorialspoint.com/developers_best_practices/index.htm)
<br>
[Go로 구현하는 블록체인](https://mingrammer.com/building-blockchain-in-go-part-2/#%EC%84%9C%EB%A1%A0)
<br>
[Building Blockchain in Go](https://jeiwan.cc/posts/building-blockchain-in-go-part-1/)