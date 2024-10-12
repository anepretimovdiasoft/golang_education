package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func main() {
	fmt.Println(Unpack("১4"))
}

type entrySlice struct {
	key   string
	value int
}

func countWordsInText(text string) map[string]int {
	wordCounter := make(map[string]int)

	words := strings.Fields(text)

	for _, word := range words {
		wordCounter[word]++
	}

	return wordCounter
}

func Top10(srcString string) []string {
	wordCounter := countWordsInText(srcString)

	var sortedSlice []entrySlice

	for key, value := range wordCounter {
		sortedSlice = append(sortedSlice, entrySlice{key, value})
	}

	sort.Slice(sortedSlice, func(i, j int) bool {

		if sortedSlice[j].value == sortedSlice[i].value {
			return sortedSlice[j].key > sortedSlice[i].key
		}

		return sortedSlice[j].value < sortedSlice[i].value
	})

	var resSlice []string
	for i := 0; i < min(10, len(sortedSlice)); i++ {
		resSlice = append(resSlice, sortedSlice[i].key)
	}

	return resSlice
}

var ErrInvalidString = errors.New("invalid string")

func Unpack(srcString string) (string, error) {
	var builder strings.Builder

	if srcString == "" {
		return "", nil
	}

	isCorrect, err := verify(srcString)
	if err != nil {
		return "", err
	}

	if isCorrect {
		runes := []rune(srcString)
		size := len(runes)
		for i := 0; i < size-1; i++ {
			if !isDigit(runes[i+1]) && !isDigit(runes[i]) {
				builder.WriteRune(runes[i])
			}
			if !isDigit(runes[i]) && isDigit(runes[i+1]) {
				repeatCount, _ := strconv.Atoi(string(runes[i+1]))
				repeatRuneWriter(runes[i], repeatCount, &builder)
			}
		}
		if !isDigit(runes[size-1]) {
			builder.WriteRune(runes[size-1])
		}
	}

	return builder.String(), nil
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func verify(srcString string) (bool, error) {
	runes := []rune(srcString)
	if isDigit(runes[0]) {
		return false, ErrInvalidString
	}

	for i := 1; i < len(runes); i++ {
		if isDigit(runes[i]) && isDigit(runes[i-1]) {
			return false, ErrInvalidString
		}
	}

	return true, nil
}

func repeatRuneWriter(r rune, repeatCount int, builder *strings.Builder) {
	if repeatCount != 0 {
		builder.WriteString(strings.Repeat(string(r), repeatCount))
	}
}
