---
theme: default
marp: true
style: |
  .columns {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 1rem;
  }
---

# Going Full-Stack

## With Go + Templ + HTMX

---

# TL;DR

- Introduction to Templ
- Introduction to HTMX
- Demo
- Why? Pros & Cons?

---

# Templ

## A revolutionary Go templating engine

---

# Before Templ...

<div class="columns">
<div>

```go
{{define "base"}}
    <html>
        <head>
            ...
        </head>
        <body>
            {{template "content" .}}
        </body>
    </html>
{{end}}
```

```go
// content.html
{{template "base" .}}
{{define "content"}}
    <div>
        {{.ID}} - {{.Title}}
    </div>
{{end}}
```

</div><div>

```go
package model

type Todo struct {
    ID int
    Title string
}

```

</div></div>

---

## Well, it's...

- not static-typed
- less expressive
- it's not "Go code"

---

# With Templ...

<div class="columns">
<div>

```go
package layout

templ Base() {
	<html>
		<head>
			...
		</head>
		<body>
			{ children... }
		</body>
	</html>
}
```

</div><div>

```go
package page

type Todo struct {
    ID int
    Title string
}

templ Index(todo *Todo) {
	@layout.Base() {
		<div>
			{ todo.ID } - { todo.Title }
		</div>
	}
}
```

</div></div>

---

- static-typed
- compiled to Go code
- just use Go!
- use LSP under the hood (great IDE supports)

---

![bg fit right](meme.png)

# HTMX

## Level up the power of Hypermedia API

---

# What is Hypermedia?

> ...an extension of the term hypertext, is a nonlinear medium of information that includes graphics, audio, video, plain text and **hyperlinks**... -- Wikipedia

## HATEOAS

> Hypermedia As The Engine Of Application State
> All “state” (that is, both the data and the network actions available on that data) is encoded in hypermedia (e.g. HTML) returned by the server.

---

<div class="columns">
<div>

# Hypermedia APIs

```html
<html>
  <body>
    <div>Account number: 12345</div>
    <div>Balance: $100.00 USD</div>
    <div>Links:
        <a href="/accounts/12345/deposits">deposits</a>
        <a href="/accounts/12345/withdrawals">withdrawals</a>
        <a href="/accounts/12345/transfers">transfers</a>
        <a href="/accounts/12345/close-requests">close-requests</a>
    </div>
  <body>
</html>
```

```html
<html>
  <body>
    <div>Account number: 12345</div>
    <div>Balance: -$50.00 USD</div>
    <div>Links:
        <a href="/accounts/12345/deposits">deposits</a>
    </div>
  <body>
</html>
```

</div><div>

# Data APIs

```json
{
  "account": {
    "account_number": 12345,
    "balance": {
      "currency": "usd",
      "value": -50.0
    },
    "status": "overdrawn"
  }
}
```

</div></div>

---

# Frontend - Backend Dualism Problem

- ## State Management

  - the source of truth of state is in the Backend
  - Frontend tends to duplicate the state

- ## "Specialization" and isolation issues
- ## Large codebase(s)

---

![bg auto](meme3.jpeg)

---

# Remember...

> Everything in software is a trade-offs

- [When should you sse Hypermedia?](https://htmx.org/essays/when-to-use-hypermedia/)
  - If your UI is mostly text & images
  - If your UI is CRUD-y
  - If you need “deep links” & good first-render performance
- Security (?)
  - Always sanitize the request and response.
  - CSRF token mechanism.
- It's not about HTMX vs other frameworks... it's more about use the **_convinient_** tool depending on the situation.

---

# Some good references

- https://templ.guide/
- https://htmx.org/essays/

## And some funny videos...

- [From React To HTMX by ThPrimeagen](https://www.youtube.com/watch?v=wIzwyyHolRs)
- [Another video by ThePrimeagen](https://www.youtube.com/watch?v=fhXyn0Vrwv4)
- [The Truth About HTMX by Theo - t3.gg](https://www.youtube.com/watch?v=NA5Fcgs_viU) or [The react video by ThPrimeagen](https://www.youtube.com/watch?v=2hMrk7A8Wf0)

---

# Thank you!

![bg auto right](meme2.jpeg)
