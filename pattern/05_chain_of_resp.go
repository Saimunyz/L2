package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// status of problem
type problemStatus int

const (
	easy problemStatus = iota
	medium
	hard
	done
	unsolved
)

// problem - some problem with status and descr
type problem struct {
	status      problemStatus
	description string
}

// handler - abstract hadler with 2 methods
type handler interface {
	handleProblem(*problem)
	setNextHandler(handler)
}

// easyProblemHandler - handle all easy problems
type easyProblemHandler struct {
	next handler
}

// handleProblem - solve easy problems
func (e *easyProblemHandler) handleProblem(p *problem) {
	if p.status != easy {
		if e.next != nil {
			e.next.handleProblem(p)
		} else {
			p.status = unsolved
			fmt.Printf("Cannot handle problem: %s", p.description)
		}
		return

	}

	// some work
	p.status = done
	fmt.Printf("Problem: %s\nStatus: easy\nSolved\n", p.description)
}

// setNextHandler - sets next handler
func (e *easyProblemHandler) setNextHandler(next handler) {
	e.next = next
}

// mediumProblemHandler - handle all medium problems
type mediumProblemHandler struct {
	next handler
}

// handleProblem - solve medium problems
func (m *mediumProblemHandler) handleProblem(p *problem) {
	if p.status != medium {
		if m.next != nil {
			m.next.handleProblem(p)
		} else {
			p.status = unsolved
			fmt.Printf("Cannot handle problem: %s", p.description)
		}
		return
	}

	// some work
	p.status = done
	fmt.Printf("Problem: %s\nStatus: medium\nSolved\n", p.description)
}

// setNextHandler - sets next handler
func (m *mediumProblemHandler) setNextHandler(next handler) {
	m.next = next
}

// hardProblemHandler - handle all hard problems
type hardProblemHandler struct {
	next handler
}

// handleProblem - solve medium problems
func (h *hardProblemHandler) handleProblem(p *problem) {
	if p.status != hard {
		if h.next != nil {
			h.next.handleProblem(p)
		} else {
			p.status = unsolved
			fmt.Printf("Cannot handle problem: %s", p.description)
		}
		return
	}

	// some work
	p.status = done
	fmt.Printf("Problem: %s\nStatus: hard\nSolved\n", p.description)
}

// setNextHandler - sets next handler
func (h *hardProblemHandler) setNextHandler(next handler) {
	h.next = next
}

// hardHandler := &hardProblemHandler{}
// mediumHandler := &mediumProblemHandler{next: hardHandler}
// easyHandler := &easyProblemHandler{next: mediumHandler}

// problem := &problem{status: hard, description: "hard problem"}

// easyHandler.handleProblem(problem)
