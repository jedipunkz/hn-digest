---
source: "https://vulang.org/"
hn_url: "https://news.ycombinator.com/item?id=50276431"
title: "Show HN: Vu – a new programming language designed for AI. Rust speeds, Py syntax"
article_title: "Vu — fast as Rust, and you can actually read it"
image: ""
author: "vu0tran"
captured_at: "2026-10-28T17:08:33Z"
capture_tool: "hn-digest"
hn_id: 50276431
score: 1
comments: 0
posted_at: "2026-10-28T16:35:24Z"
tags:
  - hacker-news
---

# Show HN: Vu – a new programming language designed for AI. Rust speeds, Py syntax

- HN: [50276431](https://news.ycombinator.com/item?id=50276431)
- Source: [vulang.org](https://vulang.org/)
- Score: 1
- Comments: 0
- Posted: 2026-10-28T16:35:24Z

## Translation

Title: Show HN: Vu – a new programming language designed for AI. Rust speeds, Py syntax
Article title: Vu — fast as Rust, and you can actually read it
Description: Vu: an experimental programming language for the age of AI. Memory-safe, as fast as Rust, and 21-42% fewer tokens to write.
HN text: hi everyone! I wanted to share what I've been working on these past 6~ years. TL;DR: Rust is great, but somewhere along the way, I learned that I was too dumb to read and write Rust. As I worked on larger and larger Rust projects, I also became frustrated with the developer experience of working in Rust (slow compile time), so I wanted a better way. Vu's language changed a bit over time but the core principles remained the same: 1. It should be easy to read and write
2. It should be memory safe
3. It should be fast to run
4. It should be fast to compile Here's how I approached each: 1. Vu is syntactically very close to Python. The ownership rules are the ones Rust has, but I personally find it a lot easier to parse and understand. I've tested this with various LLMs and found it can read and write Vu out the box because of the similarity to Python. 2. Vu uses the same principles of Rust. Every value has one owner, borrows can't escape or outlive what they borrow, and nothing changes while it's borrowed. 3. Vu compiles down to C at the end of the day, so everything runs approximately at C speeds. 4. vuc check on the compiler’s own source (about 21,500 lines of Vu), takes about 0.4~ s. This check covers parsing, type and borrow checking. cargo check on a Rust project that size usually takes several seconds from scratch. Then there are other opinionated language decisions. Vu has no macros, Vu's borrow checker is lighter weight than Rust's with simpler interfaces and generics. However, in the past few years, as Vu's language evolved and AI began to take over, I expanded Vu's principles to include a couple additional principles: 5. It should be secure
6. It should be efficient for LLMs to write 5. I want to pen test Vu in public. In https://github.com/vu0tran/vu/tree/main/tests , you'll see a battery of tests where I'm trying to basically pwn Vu. There are explicit language decisions in there as well that tries to prevent a host of security flaws that exist in other languages, e.g, integer overflow stops the program instead of quietly wrapping around. 6. The syntax of Vu is designed to be tokenizer efficient. In addition, when Vu checks throws errors, the errors are designed to be token efficient and written in a way that guides LLMs to more easily fix bugs, reducing turns overall. So, in summary, Vu is still experimental. The compiler is 21,500 lines of Vu and compiles itself. The standard library also has TLS, HTTP/2, WebSockets and a PostgreSQL client, all tested against Python’s implementations and hostile input. I think most of the pieces are there. It does also compile down also to WebAssembly so you can play around w/ the checker on the website! Feel free to try to break it! What's next:
1. Figure out how to get Vu to perform faster than Rust
2. Rewrite some OSS libraries in Vu to further prove the language

Article text:
Vu — fast as Rust, and you can actually read it

## Original Extract

Vu: an experimental programming language for the age of AI. Memory-safe, as fast as Rust, and 21-42% fewer tokens to write.

hi everyone! I wanted to share what I've been working on these past 6~ years. TL;DR: Rust is great, but somewhere along the way, I learned that I was too dumb to read and write Rust. As I worked on larger and larger Rust projects, I also became frustrated with the developer experience of working in Rust (slow compile time), so I wanted a better way. Vu's language changed a bit over time but the core principles remained the same: 1. It should be easy to read and write
2. It should be memory safe
3. It should be fast to run
4. It should be fast to compile Here's how I approached each: 1. Vu is syntactically very close to Python. The ownership rules are the ones Rust has, but I personally find it a lot easier to parse and understand. I've tested this with various LLMs and found it can read and write Vu out the box because of the similarity to Python. 2. Vu uses the same principles of Rust. Every value has one owner, borrows can't escape or outlive what they borrow, and nothing changes while it's borrowed. 3. Vu compiles down to C at the end of the day, so everything runs approximately at C speeds. 4. vuc check on the compiler’s own source (about 21,500 lines of Vu), takes about 0.4~ s. This check covers parsing, type and borrow checking. cargo check on a Rust project that size usually takes several seconds from scratch. Then there are other opinionated language decisions. Vu has no macros, Vu's borrow checker is lighter weight than Rust's with simpler interfaces and generics. However, in the past few years, as Vu's language evolved and AI began to take over, I expanded Vu's principles to include a couple additional principles: 5. It should be secure
6. It should be efficient for LLMs to write 5. I want to pen test Vu in public. In https://github.com/vu0tran/vu/tree/main/tests , you'll see a battery of tests where I'm trying to basically pwn Vu. There are explicit language decisions in there as well that tries to prevent a host of security flaws that exist in other languages, e.g, integer overflow stops the program instead of quietly wrapping around. 6. The syntax of Vu is designed to be tokenizer efficient. In addition, when Vu checks throws errors, the errors are designed to be token efficient and written in a way that guides LLMs to more easily fix bugs, reducing turns overall. So, in summary, Vu is still experimental. The compiler is 21,500 lines of Vu and compiles itself. The standard library also has TLS, HTTP/2, WebSockets and a PostgreSQL client, all tested against Python’s implementations and hostile input. I think most of the pieces are there. It does also compile down also to WebAssembly so you can play around w/ the checker on the website! Feel free to try to break it! What's next:
1. Figure out how to get Vu to perform faster than Rust
2. Rewrite some OSS libraries in Vu to further prove the language

Vu — fast as Rust, and you can actually read it
