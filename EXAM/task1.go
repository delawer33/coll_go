package main

import "fmt"

type PepeSchnele struct {
	Speed, Charisma, Wisdom int
}

func NewPepeSchnele(speed, charisma, wisdom int) *PepeSchnele {
	return &PepeSchnele{
		Speed: speed,
		Charisma: charisma,
		Wisdom: wisdom,
	}
}

func (p PepeSchnele) GetRating() int {
	return (p.Speed * 2) + (p.Charisma * 3) + p.Wisdom
}

func (p PepeSchnele) String() string {
	return fmt.Sprintf("Пепе Шнеле [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d",
        p.Speed, p.Charisma, p.Wisdom, p.GetRating())
}

func main() {
    pepe1 := NewPepeSchnele(50, 30, 100)
    pepe2 := NewPepeSchnele(20, 90, 40)

    fmt.Println(pepe1)
    fmt.Println(pepe2)

    if pepe1.GetRating() > pepe2.GetRating() {
        fmt.Printf("У первого Пепе Шнеле рейтинг выше: %d > %d\n", pepe1.GetRating(), pepe2.GetRating())
    } else if pepe2.GetRating() > pepe1.GetRating() {
        fmt.Printf("У второго Пепе Шнеле рейтинг выше: %d > %d\n", pepe2.GetRating(), pepe1.GetRating())
    } else {
        fmt.Println("Рейтинги равны")
    }
}

