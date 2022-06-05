package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func numTranslate(num string) string {
	switch num {
	case "zero": return "ноль"
	case "one": return "один"
	case "two": return "два"
	case "three": return "три"
	case "four": return "четыре"
	case "five": return "пять"
	case "six": return "шесть"
	case "seven": return "семь"
	case "eight": return "восемь"
	case "nine": return "девять"
	case "ten": return "десять"
	default: return "ошибка"
	}
}

func thesaurus(unsortedStrings ...string) map[string][]string{
	keys:=make([]rune,0,len(unsortedStrings))
	sortedStrings:=make([]string,0,len(unsortedStrings))
	for i:=range unsortedStrings {
		keys=append(keys, keyMaker(unsortedStrings[i]))
		sortedStrings=append(sortedStrings,unsortedStrings[i])
	}

	sort.Strings(sortedStrings)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for i:=1; i<len(keys);i++ {
		if keys[i]==keys[i-1] {
			keys = append(keys[0:(i-1)], keys[i:]...)
		}
	}

	sameSurnames:=make([]string,0,len(sortedStrings))
	thesaurusMap:=make(map[string][]string)
	for i:=range keys {
		sameSurnames=nil
		for j:=range sortedStrings {
			if rune(sortedStrings[j][0])==keys[i] {
				sameSurnames=append(sameSurnames, sortedStrings[j])
			}
			thesaurusMap[string(keys[i])]=sameSurnames
		}
	}
	return thesaurusMap
}

func keyMaker (stringsForKey string) rune {
	key:=rune(stringsForKey[0])
	return key
}

func getJokes(number int, nouns []string, adverbs []string, adjectives []string) []string {
	jokes:=make([]string,0,0)
	for i:=0; i<number; i++ {
		joke:=nouns[rand.Intn(len(nouns)-1)]+" "+adverbs[rand.Intn(len(adverbs)-1)]+" "+adjectives[rand.Intn(len(adjectives)-1)]
		jokes=append(jokes, joke)
	}
	return jokes
}


func main() {
	var duration int
	fmt.Scan(&duration)
	if duration <60 {
		fmt.Println("duration =",duration,": ", duration, "сек")
	} else if duration >=60 && duration <3600 {
		fmt.Println("duration =",duration,": ", duration/60, "мин", duration%60, "сек")
	} else if duration >=3600 && duration <86400 {
		fmt.Println("duration =", duration, ": ", duration/3600, "час", (duration%3600)/60, "мин", (duration%3600)%60, "сек")
	} else {
		fmt.Println("duration =", duration, ": ",duration/86400, "дн", (duration%86400)/3600, "час", ((duration%86400)%3600)/60, "мин", ((duration%86400)%3600)%60, "сек")
	}
	fmt.Println()

	var sum, degree, accSum int
	for number:=1; number<=1000; number++ {
		sum=0
		if number>=1&&number%2!=0{
			degree=number*number*number
			degreeS:=degree
			for degreeS>=1 {
				sum+=degreeS%10
				degreeS=degreeS/10
			}
			if sum%7==0&&sum!=0 {
				accSum+=number
				fmt.Println(number,"^3=", degree, "[", sum, "] накоп. сумма: ", accSum)
			}
		}
	}
	fmt.Println()


	var declension string
	for percent:=0; percent<=200; percent++ {
		switch{
			case (percent==1||percent%10==1)&&percent%100!=11: declension="процент"
			case percent%10>=2&&percent%10<=4&&(percent<10||(percent>20&&percent<110)||percent>120): declension="процента"
			case percent==0||((percent%100>=11&&percent%100<=19)||(percent%10>=5&&percent%10<=9)): declension="процентов"
		}
	fmt.Println(percent,declension)
	}
	fmt.Println()

	fmt.Println(reflect.TypeOf(15*3),15*3)
	fmt.Println(reflect.TypeOf(15/3),15/3)
	a:=float64(16)
	b:=float64(5)
	fmt.Println(reflect.TypeOf(a/b),a/b)
	fmt.Println(reflect.TypeOf(math.Pow(15,2)), math.Pow(15,2))
	fmt.Println()

	strArr1:=[]string{"в", "5", "часов", "17", "минут", "температура", "воздуха", "была", "+5", "градусов"}
	strArr2:=[]string{"примерно в", "23", "часа", "8", "минут", "03", "секунд", "температура", "воздуха", "была", "-5", "градусов Цельсия", "темп", "воды", "+12", "градусов", "Цельсия"}
	strArr3:=[]string{"+9", "примерно", "в", "23", "часа", "8", "минут", "03", "05", "секунд", "температура", "воздуха", "была", "5", "градусов Цельсия", "темп", "воды", "+2.0", "градусов", "Цельсия", "-02", "11"}
	strArrApp := make([]string, 0, 0)
	stringsArr :=[3] []string {strArr1,strArr2,strArr3}
	for x:=range stringsArr {
		strArrApp=nil
		for i:=range stringsArr[x] {
			if strings.ContainsAny(stringsArr[x][i], "0123456789+-") {
				strArrApp = append(strArrApp, "\"")
				strArrApp = append(strArrApp, stringsArr[x][i])
				strArrApp = append(strArrApp, "\"")
			} else {
				strArrApp = append(strArrApp, stringsArr[x][i])
			}
		}
		for i:=range strArrApp {
			if strings.ContainsAny(strArrApp[i], "0123456789+-")&&len(strArrApp[i])<2 {
				strArrApp[i]="0"+ strArrApp[i]
			}
			if (strings.ContainsAny(strArrApp[i], "+")|| strings.ContainsAny(strArrApp[i], "-"))&&len(strArrApp[i])<3 {
				for j:=range strArrApp[i]{
					if strArrApp[i][j]=='+' {
						strArrApp[i]=string(strArrApp[i][j])+"0"+string(strArrApp[i][j+1])
					}
					if strArrApp[i][j]=='-' {
						strArrApp[i]=string(strArrApp[i][j])+"0"+string(strArrApp[i][j+1])
					}
				}
			}
		}
		strArrApp:= strArrApp
		fmt.Println(x+1,"Исходный список: ", stringsArr[x])
		fmt.Println("\t1.Новый список + добавление элементов-кавычек: ", strArrApp)
		fmt.Print("\t2.Окончательная строка: ")
		for i:=range strArrApp {
			if i<(len(strArrApp)-1)&& strArrApp[i]=="\""&& strings.ContainsAny(strArrApp[i+1], "0123456789+-")|| strings.ContainsAny(strArrApp[i], "0123456789+-") {
				fmt.Print(string(strArrApp[i]))
			} else {fmt.Print(string(strArrApp[i])+ " ")}
		}
		fmt.Println()
	}


	igorStr:="инженер-конструктор Игорь"
	marinaStr:="главный бухгалтер МАРИНА"
	nikolaiStr:="токарь высшего разряда нИКОЛАй"
	aelitaStr:="директор аэлита"
	workersArr :=[4]string {igorStr,marinaStr,nikolaiStr,aelitaStr}
	for i:=range workersArr {
		workerNameArr:=strings.Split(workersArr[i], " ")
		workerNameRuneArr :=[]rune(workerNameArr[len(workerNameArr)-1])
		for j:=range workerNameRuneArr {
			if j==0 {
				workerNameRuneArr[j]=unicode.ToUpper(workerNameRuneArr[j])
			} else {
				workerNameRuneArr[j]=unicode.ToLower(workerNameRuneArr[j])}
		}
		fmt.Println("Привет, "+string(workerNameRuneArr)+"!")
	}
	fmt.Println()

	prices:= []float64{57.8, 46.40, 97, 12.3, 67.54, 8.07, 982.12}

	for i:=range prices {
		pricesEven, pricesFrac := math.Modf(prices[i])
		pricesDotStr:=strconv.Itoa(int(pricesFrac *100))
		if len(pricesDotStr)==1 {
			pricesDotStr="0"+pricesDotStr
		}
		fmt.Printf("%.0f руб %s коп, ", pricesEven, pricesDotStr)
	}
	fmt.Print("\n")
	fmt.Print("id элементов перед сортировкой: ")
	for i:=range prices {
		fmt.Print(&prices[i]," ")
	}

	sort.Float64s(prices)

	fmt.Print("\n")
	fmt.Print("id элементов после сортировки: ")
	for i:=range prices {
		fmt.Print(&prices[i]," ")
	}
	fmt.Print("\n")

	fmt.Println(prices)
	fmt.Println("5 самых дорогих товаров:")
	for i:=len(prices)-1; i>len(prices)-6; i-- {
		fmt.Println(prices[i])
	}
	fmt.Println()

	fmt.Println(numTranslate("zero"))
	fmt.Println(numTranslate("one"))
	fmt.Println(numTranslate("two"))
	fmt.Println(numTranslate("three"))
	fmt.Println(numTranslate("four"))
	fmt.Println(numTranslate("five"))
	fmt.Println(numTranslate("six"))
	fmt.Println(numTranslate("seven"))
	fmt.Println(numTranslate("eight"))
	fmt.Println(numTranslate("nine"))
	fmt.Println(numTranslate("ten"))
	fmt.Println(numTranslate("`12weddsfadfw`"))
	fmt.Println()

	surnamesMap:=thesaurus("Miheev","Pavlova", "Maltsev", "Grigorjev", "Smirnov", "Kirillov", "Pylnova", "Livneva")
	for letter, surname:=range surnamesMap {
		fmt.Printf("\"%s\": %q\n", letter, surname)
	}
	fmt.Println()

	nouns:=[]string {"автомобиль", "лес", "огонь", "город", "дом"}
	adverbs:=[]string {"сегодня", "вчера", "завтра", "позавчера", "ночью"}
	adjectives:=[]string {"веселый", "яркий", "зеленый", "утопичный", "мягкий"}
	jokes:= getJokes(10, nouns ,adverbs ,adjectives)
	for i:=range jokes {
		fmt.Printf("\"%s\"\n", jokes[i])
	}
}



