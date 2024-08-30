package main

import (
  "fmt"
  "strings"
)

type attendee struct {
  party string
  email string
  tickets uint
}

func new_attendee(party, email string, tickets uint) * attendee {
  var a attendee

  a.party = party
  a.email = email
  a.tickets = tickets

  return &a
}

func (a * attendee) validate() [] error {
  errs := make([] error, 0, 3)

  if len(a.party) < 2 {
    errs = append(errs, fmt.Errorf("The name of the party is too short."))
  }

  // Require at least a@b.c
  if len(a.email) < 5 || !strings.Contains(a.email, "@") {
    errs = append(errs, fmt.Errorf("The email address is too short or missint the '@' symbol."))
  }

  if a.tickets == 0 {
    errs = append(errs, fmt.Errorf("At least one ticket must be requested."))
  }

  return errs
}

type event struct {
  name string
  tickets uint
  reserved uint
  attendees [] * attendee
}

func new_event(name string, tickets uint) * event {
  var e event

  e.name = name
  e.tickets = tickets
  e.reserved = 0
  e.attendees = make([] * attendee, 1, tickets)

  return &e
}

func (e * event) book_attendee(a * attendee) error {
  var err error

  available := e.tickets - e.reserved

  if a.tickets <= available {
    e.attendees = append(e.attendees, a)
    e.reserved += a.tickets
  } else {
    err = fmt.Errorf("The attendee's %v ticket(s) requested exceed the available %v ticket(s).", a.tickets, available)
  }

  return err
}

func register_attendee() (* attendee, [] error) {
  var party string
  var email string
  var tickets uint

  fmt.Print("Name of party: ")
  fmt.Scan(&party)

  fmt.Print("Email address of party: ")
  fmt.Scan(&email)

  fmt.Print("Tickets to reserve: ")
  fmt.Scan(&tickets)

  a := new_attendee(party, email, tickets)
  errs := a.validate()

  return a, errs
}

func main() {
  e := new_event("Go Conference", 50)

  fmt.Printf("Welcome to %v!\n", e.name)

  a, errs := register_attendee()

  if len(errs) > 0 {
    for err := range errs {
      fmt.Println(err)
    }
  } else {
    err := e.book_attendee(a)
    
    if err != nil {
      fmt.Println(err)
    } else {
      fmt.Printf("Thank you %v for booking %v ticket(s). A confirmation email has been sent to %v.\n", a.party, a.tickets, a.email)
    }
  }
}
