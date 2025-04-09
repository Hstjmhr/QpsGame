package sz

import (
	"msqp/utils"
	"sort"
	"sync"
)

type Logic struct {
	sync.RWMutex
	cards []int // 52张牌
}

func NewLogic() *Logic {
	return &Logic{
		cards: make([]int, 0),
	}
}

// washCards 方块 梅花 红心 黑桃
func (l *Logic) washCards() {
	l.cards = []int{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d,
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d,
		0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d,
	}
	for i, v := range l.cards {
		random := utils.Rand(len(l.cards))
		l.cards[i] = l.cards[random]
		l.cards[random] = v
	}
}

// getCards 获取三张手牌
func (l *Logic) getCards() []int {
	cards := make([]int, 3)
	l.RLock()
	defer l.RUnlock()
	for i := 0; i < 3; i++ {
		if len(cards) == 0 {
			break
		}
		card := l.cards[len(l.cards)-1]
		l.cards = l.cards[:len(l.cards)-1]
		cards[i] = card
	}
	return cards
}

// CompareCards result 0 he 大于0 win 小于 0  lose
func (l *Logic) CompareCards(from []int, to []int) int {
	//获取牌类型
	fromType := l.getCardsType(from)
	toType := l.getCardsType(to)
	if fromType != toType {
		return int(fromType - toType)
	}
	//对子牌比较
	if fromType == DuiZi {
		duiFrom, danFrom := l.getDuiZi(from)
		duiTo, danTo := l.getDuiZi(to)
		if duiFrom != duiTo {
			return duiFrom - duiTo
		}
		return danFrom - danTo
	}
	//普通牌比较
	valuesFrom := l.getCardsValues(from)
	valuesTo := l.getCardsValues(to)
	if valuesFrom[2] != valuesTo[2] {
		return valuesFrom[2] - valuesTo[2]
	}
	if valuesFrom[1] != valuesTo[1] {
		return valuesFrom[1] - valuesTo[1]
	}
	if valuesFrom[0] != valuesTo[0] {
		return valuesFrom[0] - valuesTo[0]
	}
	return 0
}

func (l *Logic) getCardsType(cards []int) CardsType {
	//1. 豹子 牌面值相等 梅花 方块 红心 黑桃
	one := l.getCardsNumber(cards[0])
	two := l.getCardsNumber(cards[1])
	three := l.getCardsNumber(cards[2])
	if one == two && two == three {
		return BaoZi
	}
	//2. 金花 颜色相同 顺子
	jinhua := false
	oneColor := l.getCardsColor(cards[0])
	twoColor := l.getCardsColor(cards[1])
	threeColor := l.getCardsColor(cards[2])
	if oneColor == twoColor && oneColor == threeColor {
		jinhua = true
	}
	//3 顺子 先排序A 2 3 4 ... J Q K
	shunzi := false
	values := l.getCardsValues(cards)
	oneV := values[0]
	twoV := values[1]
	threeV := values[2]
	if oneV+1 == twoV && twoV+1 == threeV {
		shunzi = true
	}
	if oneV == 2 && twoV == 3 && threeV == 14 {
		shunzi = true
	}
	if jinhua && shunzi {
		return ShunJin
	}
	if jinhua {
		return JinHua
	}
	if shunzi {
		return ShunZi
	}
	if oneV == twoV || twoV == threeV {
		return DuiZi
	}
	return DanZhang
}
func (l *Logic) getCardsValues(cards []int) []int {
	v := make([]int, len(cards))
	for i, card := range cards {
		v[i] = l.getCardsValue(card)
	}
	sort.Ints(v)
	return v
}

func (l *Logic) getCardsValue(card int) int {
	value := card & 0x0f
	return value
}

func (l *Logic) getCardsNumber(card int) int {
	return card & 0x0f
}

// 作用是根据给定的 card 值返回对应的牌色。如果 card 在有效范围内，
// 则返回相应的颜色；否则返回空字符串。这种方法通常用于卡牌游戏中，通过 card 的值来识别和显示牌的颜色。
func (l *Logic) getCardsColor(card int) string {
	colors := []string{"方块", "梅花", "红心", "黑桃"}
	//取模
	if card >= 0x01 && card <= 0x3d {
		return colors[card/0x10]
	}
	return ""
}

func (l *Logic) getDuiZi(cards []int) (int, int) {
	values := l.getCardsValues(cards)
	//判断对子情况
	if values[0] == values[1] {
		//AAB
		return values[0], values[2]
	}
	return values[1], values[0]
}
