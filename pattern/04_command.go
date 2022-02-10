package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Command interface
type Command interface {
	execute()
}

// seller interface
type seller interface {
	sell(amount int)
	buy(amount int)
}

// sellCommand struct
type sellCommand struct {
	seller seller
	amount int
}

// newSellCommand - create new instane with all needed params
func newSellCommand(seller seller, amount int) *sellCommand {
	return &sellCommand{
		amount: amount,
		seller: seller,
	}
}

func (s *sellCommand) execute() {
	s.seller.sell(s.amount)
}

// buyCommand struct
type buyCommand struct {
	seller seller
	amount int
}

// newBuyCommand - create new instane with all needed params
func newBuyCommand(seller seller, amount int) *buyCommand {
	return &buyCommand{
		amount: amount,
		seller: seller,
	}
}

func (b *buyCommand) execute() {
	b.seller.buy(b.amount)
}

// Sender
type Bot struct {
	cmd Command
}

// newBot - creates new instanse of Bot
func newBot(cmd Command) *Bot {
	return &Bot{
		cmd: cmd,
	}
}

// setCommand - sets new command in bot
func (b *Bot) setCommand(cmd Command) {
	b.cmd = cmd
}

func (b *Bot) start() {
	b.cmd.execute()
}

// Ivan - reciever
type Ivan struct {
	items int
}

func (i *Ivan) sell(amount int) {
	i.items -= amount
}

func (i *Ivan) buy(amount int) {
	i.items += amount
}

/*
 *	reciver := &Ivan{items: 10}
 *	buyCommand := newBuyCommand{seller: reciever, amount: 5}
 *	sellCommand := newSellCommand{seller: reciever, amount: 3}
 *	buyerBot := Bot{cmd: buyCommand}
 *	sellerBot := Bot{cmd :sellCommand}
 *
 *	sellerBot.start()
 *	buyerBot.start()
 */
