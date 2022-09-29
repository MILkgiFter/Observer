package main

func main() {

	Artist := newTicket("Instasamka", "Astana", 100)

	observerFirst := &Guest{id: "dora@gmail.com"}
	observerSecond := &Guest{id: "moneylen@gmail.com"}
	obik := &Guest{id: "da ya"}

	Artist.register(observerFirst)
	Artist.register(observerSecond)
	Artist.register(obik)

	Artist.updateAvailability()
}
