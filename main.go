// 4個のダイスと１人につき10本のマッチ棒が必要で、それらは中央にまとめて置く
// マッチする目を振るたびにマッチ棒を中央から得る
// 目的：11本のマッチ棒を一番最初に得ること
// ルール：
// ①２個のダイスを振り、２個の目が等しい場合、２本のマッチを受け取ります。２本のマッチ棒を受け取り、再度投げる。
// ②２個の目が等しいダブルを振り、それが中央の１個の目とマッチする場合、２本のマッチ棒を得る。そして再び振る。
// ③もし出目が中央の2個と完全にマッチするなら3本のマッチ棒を受け取る。そして中央の2個のダイスを振り直し、新たな、目標値を確立。
// ④もしどの目標値ともマッチしなければ、ダイスを次のプレイヤーに渡す。
// 例：目標値は２と４
// １回目：４と５→１本マッチ棒獲得
// ２回目：２と２→２本マッチ棒獲得
// ３回目：３と５→終了し、次のプレイヤーへ渡す
// パッケージ名main
package main

//必要なimport
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

//player構造体
type Player struct {
	matchStickCount int
}

func (p *Player) setMatchStickCount(matchStickCount int) {
	// +=で足していく
	p.matchStickCount += matchStickCount
}

func (p *Player) getMatchStickCount() int {
	return p.matchStickCount
}

var players = []*Player{}

//エントリポイントとは、コンピュータプログラムを実行する際に、一番最初に実行することになっている箇所のこと。
//mainパッケージのmain関数がエントリポイントとなる
//func main関数定義
//引数を受け取らないので引数は省略
//戻り値も指定しないので省略(戻り値の型、return文は省略)

var sc = bufio.NewScanner(os.Stdin)

//マッチ棒
var MatchStick int

func main() {
	//for {
	//[①How many player?]
	fmt.Println("MatchStickDiceGame")
	fmt.Println("11本のマッチ棒を一番最初に取得したプレイヤーの勝ちです")
	fmt.Println(`プレイヤーは何人ですか？`)
	// var totalPlayer int
	// fmt.Scan(&totalPlayer)
	var totalPlayer int
	if sc.Scan() {
		var x string
		x = sc.Text()
		//totalPlayer, _ := strconv.Atoi(x) //:=は既に型を定義してあるためerrorになる。：はいらない。
		//strconv.Atoiがより高速
		totalPlayer, _ = strconv.Atoi(x)
	}
	if totalPlayer < 4 {
		fmt.Println(totalPlayer, `人で対戦します`)
	} else {
		fmt.Println(`3人以上では対戦できません`)
		return //関数から抜ける
	}
	var playerArrayList []int
	playerArrayList = startPlayer(totalPlayer)
	for i := 0; i < totalPlayer; i++ {
		fmt.Println("プレイヤー", (i + 1), playerArrayList[i], "番目")
		players = append(players, &Player{})
	}
	//ダイスを投げる順番でソート
	sort.Ints(playerArrayList)
	//メインの数字を決める
	fmt.Println("メインの数字を決めます。エンターキーを押してください")
	bufio.NewScanner(os.Stdin).Scan()
	main1, main2 := mainNumber()

	//順番にダイスを振る
	for i := 0; i < len(playerArrayList); i++ {
		fmt.Println("プレイヤー", playerArrayList[i], "エンターキーを押してダイスを投げてください")
		//throw Dice呼び出し
		throwDiceResult1, throwDiceResult2 := throwDice()

		//判定する
		judge(players[i], throwDiceResult1, throwDiceResult2, main1, main2)

	}

}

//2.who start?
// func startPlayer(totalPlayer int) []int {
// 	//順番を決める
// 	//var playerArray [totalPlayer]int //配列は数字していしかできないため、error
// 	//variablesを使いたい場合は以下 可変長配列
// 	playerArray := make([]int, totalPlayer)
// 	fmt.Println(playerArray)
// 	for i := 0; i < totalPlayer; i++ {
// 		num := rand.Intn(totalPlayer + 1)
// 		playerArray[i] = num
// 	}
// 	fmt.Println(playerArray)
// 	return playerArray
// }
//totalPlayer数の連番スライスを作っておき、その内容をシャッフル

//プレイヤー順番決め
func startPlayer(totalPlayer int) []int {
	playerArray := []int{1, 2, 3, 4, 5, 6, 7, 8}[:totalPlayer]
	rand.Shuffle(len(playerArray), func(i, j int) {
		playerArray[i], playerArray[j] = playerArray[j], playerArray[i]
	})
	return playerArray
}

//メインの数字を決める
func mainNumber() (int, int) {
	//Mainの数字生成
	var main1 int
	var main2 int
	main1 = rand.Intn(6)
	main2 = rand.Intn(6)
	fmt.Println("メインの数字は", main1+1, "と", main2+1, "です")

	return main1, main2
}

//throw Dice結果
func throwDice() (int, int) {
	bufio.NewScanner(os.Stdin).Scan()
	var throwDice_1 int
	var throwDice_2 int
	throwDice_1 = rand.Intn(6)
	throwDice_2 = rand.Intn(6)
	fmt.Println(throwDice_1+1, "と", throwDice_2+1, "が出ました")
	return throwDice_1, throwDice_2
}

func judge(player *Player, throwDiceResult1 int, throwDiceResult2 int, main1 int, main2 int) {
	// ①２個のダイスを振り、２個の目が等しい場合、２本のマッチを受け取ります。２本のマッチ棒を受け取り、再度投げる。
	if throwDiceResult1 == throwDiceResult2 {
		player.setMatchStickCount(1)
		fmt.Println("1本マッチ棒獲得")
	} else if throwDiceResult1 == main1 || throwDiceResult1 == main2 || throwDiceResult2 == main1 || throwDiceResult2 == main2 {
		// ②２個の目が等しいダブルを振り、それが中央の１個の目とマッチする場合、２本のマッチ棒を得る。そして再び振る。
		player.setMatchStickCount(2)
		fmt.Println("2本マッチ棒獲得")
	} else if throwDiceResult1 == main1 || throwDiceResult1 == main2 && throwDiceResult2 == main1 || throwDiceResult2 == main2 {
		// ③もし出目が中央の2個と完全にマッチするなら3本のマッチ棒を受け取る。そして中央の2個のダイスを振り直し、新たな、目標値を確立。
		player.setMatchStickCount(3)
		fmt.Println("3本マッチ棒獲得")
	}

	fmt.Println("player match stick count: ", player.getMatchStickCount())
}

// // You can edit this code!
// // Click here and start typing.
// package main

// import "fmt"

// type Human struct {
// 	name string
// 	age  int
// }

// func (h *Human) setName(name string) {
// 	h.name = name
// }

// func (h *Human) getName() string {
// 	return h.name
// }

// func (h *Human) setAge(age int) {
// 	h.age = age
// }

// func (h *Human) getAge() int {
// 	return h.age
// }

// func main() {

// 	humans := []*Human{}

// 	human1 := &Human{}
// 	human1.setName("yui")
// 	human1.setAge(60)

// 	humans = append(humans, human1)

// 	human2 := &Human{}
// 	human2.setName("momo")
// 	human2.setAge(20)

// 	humans = append(humans, human2)

// 	fmt.Println(humans[0].getName())
// 	fmt.Println(humans[0].getAge())
// 	fmt.Println(humans[1].getName())
// 	fmt.Println(humans[1].getAge())
// }
