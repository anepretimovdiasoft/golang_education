package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
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
	fmt.Println(Top10(text))
}

func Top10(srcString string) []string {
	words := strings.Fields(srcString)
	wordCounter := make(map[string]int)

	for _, word := range words {
		wordCounter[word]++
	}

	type entrySlice struct {
		key   string
		value int
	}

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

	if isCorrect, err := Verify(srcString); isCorrect && err == nil {
		runes := []rune(srcString)
		size := len(runes)
		for i := 0; i < size-1; i++ {
			if !unicode.IsDigit(runes[i+1]) && !unicode.IsDigit(runes[i]) {
				builder.WriteRune(runes[i])
			}
			if !unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+1]) {
				repeatCount, _ := strconv.Atoi(string(runes[i+1]))
				if repeatCount != 0 {
					builder.WriteString(strings.Repeat(string(runes[i]), repeatCount))
				}
			}
		}
		if !unicode.IsDigit(runes[size-1]) {
			builder.WriteRune(runes[size-1])
		}
	} else {
		return "", err
	}

	return builder.String(), nil
}

func Verify(srcString string) (bool, error) {

	runes := []rune(srcString)
	if unicode.IsDigit(runes[0]) {
		return false, ErrInvalidString
	}

	for i := 1; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i-1]) {
			return false, ErrInvalidString
		}
	}

	return true, nil
}
