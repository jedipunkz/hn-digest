---
source: "https://prod.codes/blog/reverse-engineering-with-ai-your-decompiler-can-lie-to-you/"
hn_url: "https://news.ycombinator.com/item?id=50400753"
title: "Reverse Engineering with AI: Your Decompiler Can Lie to You"
article_title: "Reverse Engineering with AI: Your Decompiler Can Lie to You"
image: "https://prod.codes/img/xref/card-xref-engineering-experience.jpg"
author: "alex09x"
captured_at: "2026-11-07T01:39:39Z"
capture_tool: "hn-digest"
hn_id: 50400753
score: 2
comments: 0
posted_at: "2026-11-07T01:31:19Z"
tags:
  - hacker-news
---

# Reverse Engineering with AI: Your Decompiler Can Lie to You

- HN: [50400753](https://news.ycombinator.com/item?id=50400753)
- Source: [prod.codes](https://prod.codes/blog/reverse-engineering-with-ai-your-decompiler-can-lie-to-you/)
- Score: 2
- Comments: 0
- Posted: 2026-11-07T01:31:19Z

## Translation

Title: Reverse Engineering with AI: Your Decompiler Can Lie to You
Description: A field name, 2,000 look-alike functions and a wrong pseudo-C branch: an XRef walkthrough in which every conclusion is checked against machine instructions.

Article text:
Skip to content everything before it matters Projects Blog About RSS On this page What the public example actually is
A key string is a starting point, not a field
A field’s address and the value stored there are different things
An access range is not a struct declaration
Two thousand functions that look like the answer
What ranking gives you, and what it doesn’t
Ask where the base pointer came from
One function, different call sites, different questions
Look at the candidate from another angle
An empty answer should explain itself
Fast repeats: what can be cached and what can’t
A workflow you can hand to an agent
Reverse Engineering with AI: Your Decompiler Can Lie to You
A field name, 2,000 look-alike functions and a wrong pseudo-C branch: an XRef walkthrough in which every conclusion is checked against machine instructions.
What the public example actually is
A key string is a starting point, not a field
A field’s address and the value stored there are different things
An access range is not a struct declaration
Two thousand functions that look like the answer
What ranking gives you, and what it doesn’t
Ask where the base pointer came from
One function, different call sites, different questions
Look at the candidate from another angle
An empty answer should explain itself
Fast repeats: what can be cached and what can’t
A workflow you can hand to an agent
TL;DR. Serialized field names can lead you to their offsets. The harder question is which code changes those fields: in a public example, 2,000 decoys write the same offsets as the real function. The real function and the displayed decoys earn the same 40 points for their writes to those fields. The real function gets ten more because it shares a caller, main , with the serializer; the decoys have no callers at all. The candidate is then checked against its instructions and call sites. The walkthrough uses XRef 0.4.0 and a synthetic C and assembly program; it demonstrates a method and is not a benchmark.
Before the 0.4.0 release, one xref field query came back with not_found_under_limits : no matching layouts. The response flagged incomplete discovery under its analysis limits, but it gave no way to tell which keys the tool had seen or where it had lost the connection to the fields. That case is recorded in issue 76 .
An empty list makes the next wrong move easy: widen the search at random, conclude that the structure is not there, or start taking the whole binary apart by hand. What I needed was an answer to a much narrower question: which observations get you from a field name to the code that changes it?
Put in familiar terms: a program reports "status": "failed" —which function changed that field? That is an illustrative way to frame the question. The recorded example below uses the keys c , m and e , preserved exactly so the source, instructions and output can be checked together.
When an AI agent does this work, it needs the same checkable chain a human does. In this setup, XRef is a native analyzer: the agent asks it questions through the CLI or MCP and works with the data that comes back. Finding a plausible function, understanding its role and reading its pseudo-C correctly are three separate steps, and the third is where the title comes from: later on, a historical XRef reproducer shows pseudo-C that, read literally, returns 99 for input the machine code evaluates to 4 . The recording below contains CLI commands and their captured outputs; the closing section turns the same sequence into a checklist an agent can follow.
For an executable without symbols, this is a surprisingly awkward question. You can know a field’s name because it shows up in serialized output and know almost nothing about how it is laid out in memory. The source type names are gone. The function names are gone. A single small displacement can appear in thousands of completely unrelated instructions. A plausible decompilation can make an early guess look convincing long before the facts back it up.
While I was working on XRef , this question gradually got sharper: first find code from the name of a serialized field, then match a set of neighboring keys and find the functions that write those fields. Short keys exposed the next gap: even a negative result has to explain what the tool actually managed to check. You can follow that path through issues 58 , 74 and 76 .
The xref field command connects key uses to offsets relative to an incoming pointer, then finds and ranks candidate writers. A string reference, an offset and a matching function are different leads. To turn them into an answer, you have to follow the whole chain and keep the instructions each conclusion rests on.
This article works through one question step by step: which code changes the values associated with these keys? Along the way it shows the false leads that keep the problem from being a simple search for memory stores.
What the public example actually is
To follow this chain to the end, I use a small, published test program with a known answer, so every conclusion can be checked against the source and the machine instructions. Real-world tasks shaped what the tool needed to do; the walkthrough itself runs on the open example.
The existing recording uses the published XRef 0.4.0 Linux executable. It was made on November 5, 2026; I reuse its commands and outputs here rather than claiming a new run. The exact build revision and file hashes are in the recording’s environment record (linked at the end).
The target program is generated C and assembly. The part of the C that matters fits on one screen:
struct record { uint64_t c, m, e; };
void serializer ( struct record * p ) {
emit ( "c" , 1 , & p -> c );
emit ( "m" , 1 , & p -> m );
emit ( "e" , 1 , & p -> e );
}
void writer ( struct record * p , uint64_t value ) {
p -> m += value;
p -> e ^= value;
}
This is an excerpt from the full generated source . Function attributes and build flags preserve the call pattern the walkthrough relies on. main calls writer first, then serializer .
The emit function does nothing but mix the key, the length and the value at the given address into a volatile variable. That reproduces the machine-level pattern “key + value address → shared function” without a real JSON encoder or Serde. It is a synthetic C and assembly example for studying the mechanism. Assessing accuracy on real Rust applications requires a separate evaluation.
The assembly part adds 2,000 decoy functions. Each one writes at offsets +8 and +16 from its first argument. None of them is ever called. That is what keeps the question from being trivial: an analysis that just looks for matching offsets will produce plenty of convincing answers.
The program was built with GCC 13.3.0 using -O2 -g0 -fno-inline -fno-ipa-icf , then run through objcopy --strip-all . XRef reports 2,007 functions and zero symbols; in the recorded example, the boundaries of those functions come from .eh_frame , the stack-unwinding data. The analysis runs on the stripped file; the symbols of a separate reference build are used only to check the answer.
A key string is a starting point, not a field
Searching for m does not immediately tell you which memory location holds the value named m . A short string is an especially weak lead: its bytes can be part of a shared literal or be used in some other context. Those were exactly the cases that had to be handled after the request in issue 76. PR 77 added shared-literal recognition for short keys, per-key diagnostics and a way to inspect partial matches.
What I need is the relationship between two arguments at the use site: the key being passed and the value passed along with it. Only then does the name acquire a useful local meaning.
The first query in the recording looks for layout evidence only:
xref -b target-stripped field --keys c,m,e --layout-only
The report finds one candidate function, 0x11c0..0x1204 , with offsets relative to its incoming rdi . For the three keys it reports c -> 0x0 , m -> 0x8 and e -> 0x10 ; the corresponding instructions read eight bytes.
These values are useful because the report also lists the call sites where each key and value are passed. In the recorded image they are 0x11d7 , 0x11ec and 0x1202 ; the last one is a tail jump to the same callee. You can inspect the function that passes the keys instead of accepting the layout as an unexplained result.
Here is a shortened disassembly excerpt from the recording. The omitted instructions load the keys and their lengths:
0x11c5 mov rdx , rdi
0x11c8 mov rbx , rdi
0x11d7 call 0x11a0
0x11dc lea rdx , [ rbx + 8 ]
0x11ec call 0x11a0
0x11f1 lea rdx , [ rbx + 0x10 ]
0x1202 jmp short 0x11a0
On entry, rdi holds the pointer p . The copy saved in rbx survives the first call. So the values passed in rdx are p , p + 8 and p + 16 . In the full listing, the matching key string addresses and lengths are passed right next to them. That is how key names get tied to offsets.
Figure 1. Each arrow is a separate investigative step. Finding a later link does not automatically prove the earlier ones.
A field’s address and the value stored there are different things
One distinction shaped the problem more than any other: following a pointer and following a value loaded through that pointer are not the same thing.
Consider two schematic instructions. They illustrate the distinction and are not an additional result from the recorded test program:
lea rax , [ rdi + 8 ] ; address: incoming pointer + 8
mov rax , [ rdi + 8 ] ; value: eight bytes loaded there
After lea , rax still holds an address relative to the original input. After mov , it holds whatever was stored at that address: a number, a pointer to a different object, or a value the current analysis cannot classify. A later write through [rax+16] does not automatically become a write to rdi+24 .
That is why it matters where a base pointer came from: whether it arrived as an argument, was returned from a call or was loaded from memory. Each leads to a different conclusion about which object you are looking at. Lose track of a pointer’s origin, and later offsets can be wrongly attributed to the earlier object.
There is another practical trap here: searching only for lea can miss a field that is accessed through a direct load or store. An empty result from such a search does not prove the field was never used. It means the chosen instruction pattern was not found under those search conditions.
Figure 2. A schematic example: after a load from memory, you cannot keep tracking the original base by simply adding offsets.
An access range is not a struct declaration
The three offsets and read widths in this example give an addressed range of at least 0x18 bytes. It is tempting to write sizeof(record) = 24 next to it and move on.
Here, the source lets us look up the real declaration. The stripped binary on its own supports a narrower claim: those positions were accessed with those widths. It does not follow that there are no fields beyond them, that no field was missed, or that the allocated memory ends there. Alignment, tail padding and how the memory was allocated are separate questions.
The report expresses this as addressed-span>=0x18 and exact-size=unknown . It also keeps the lexical order of call sites separate from the source declaration order or any universal execution order. These caveats are a useful part of the result, not an apology tacked on after the analysis.
The same goes for names. We observed the byte strings c , m and e . We know the type name struct record from the source and the reference build’s symbols, not from the query result.
Explicit boundaries make the result more useful. You can search for writes at the observed offsets without first declaring an invented complete type.
Two thousand functions that look like the answer
Now let’s ask for functions that write both named fields through the same pointer:
xref -b target-stripped field --keys c,m,e \
--write --writes-all-of m,e --limit 5
The --

[truncated]

## Original Extract

A field name, 2,000 look-alike functions and a wrong pseudo-C branch: an XRef walkthrough in which every conclusion is checked against machine instructions.

Skip to content everything before it matters Projects Blog About RSS On this page What the public example actually is
A key string is a starting point, not a field
A field’s address and the value stored there are different things
An access range is not a struct declaration
Two thousand functions that look like the answer
What ranking gives you, and what it doesn’t
Ask where the base pointer came from
One function, different call sites, different questions
Look at the candidate from another angle
An empty answer should explain itself
Fast repeats: what can be cached and what can’t
A workflow you can hand to an agent
Reverse Engineering with AI: Your Decompiler Can Lie to You
A field name, 2,000 look-alike functions and a wrong pseudo-C branch: an XRef walkthrough in which every conclusion is checked against machine instructions.
What the public example actually is
A key string is a starting point, not a field
A field’s address and the value stored there are different things
An access range is not a struct declaration
Two thousand functions that look like the answer
What ranking gives you, and what it doesn’t
Ask where the base pointer came from
One function, different call sites, different questions
Look at the candidate from another angle
An empty answer should explain itself
Fast repeats: what can be cached and what can’t
A workflow you can hand to an agent
TL;DR. Serialized field names can lead you to their offsets. The harder question is which code changes those fields: in a public example, 2,000 decoys write the same offsets as the real function. The real function and the displayed decoys earn the same 40 points for their writes to those fields. The real function gets ten more because it shares a caller, main , with the serializer; the decoys have no callers at all. The candidate is then checked against its instructions and call sites. The walkthrough uses XRef 0.4.0 and a synthetic C and assembly program; it demonstrates a method and is not a benchmark.
Before the 0.4.0 release, one xref field query came back with not_found_under_limits : no matching layouts. The response flagged incomplete discovery under its analysis limits, but it gave no way to tell which keys the tool had seen or where it had lost the connection to the fields. That case is recorded in issue 76 .
An empty list makes the next wrong move easy: widen the search at random, conclude that the structure is not there, or start taking the whole binary apart by hand. What I needed was an answer to a much narrower question: which observations get you from a field name to the code that changes it?
Put in familiar terms: a program reports "status": "failed" —which function changed that field? That is an illustrative way to frame the question. The recorded example below uses the keys c , m and e , preserved exactly so the source, instructions and output can be checked together.
When an AI agent does this work, it needs the same checkable chain a human does. In this setup, XRef is a native analyzer: the agent asks it questions through the CLI or MCP and works with the data that comes back. Finding a plausible function, understanding its role and reading its pseudo-C correctly are three separate steps, and the third is where the title comes from: later on, a historical XRef reproducer shows pseudo-C that, read literally, returns 99 for input the machine code evaluates to 4 . The recording below contains CLI commands and their captured outputs; the closing section turns the same sequence into a checklist an agent can follow.
For an executable without symbols, this is a surprisingly awkward question. You can know a field’s name because it shows up in serialized output and know almost nothing about how it is laid out in memory. The source type names are gone. The function names are gone. A single small displacement can appear in thousands of completely unrelated instructions. A plausible decompilation can make an early guess look convincing long before the facts back it up.
While I was working on XRef , this question gradually got sharper: first find code from the name of a serialized field, then match a set of neighboring keys and find the functions that write those fields. Short keys exposed the next gap: even a negative result has to explain what the tool actually managed to check. You can follow that path through issues 58 , 74 and 76 .
The xref field command connects key uses to offsets relative to an incoming pointer, then finds and ranks candidate writers. A string reference, an offset and a matching function are different leads. To turn them into an answer, you have to follow the whole chain and keep the instructions each conclusion rests on.
This article works through one question step by step: which code changes the values associated with these keys? Along the way it shows the false leads that keep the problem from being a simple search for memory stores.
What the public example actually is
To follow this chain to the end, I use a small, published test program with a known answer, so every conclusion can be checked against the source and the machine instructions. Real-world tasks shaped what the tool needed to do; the walkthrough itself runs on the open example.
The existing recording uses the published XRef 0.4.0 Linux executable. It was made on November 5, 2026; I reuse its commands and outputs here rather than claiming a new run. The exact build revision and file hashes are in the recording’s environment record (linked at the end).
The target program is generated C and assembly. The part of the C that matters fits on one screen:
struct record { uint64_t c, m, e; };
void serializer ( struct record * p ) {
emit ( "c" , 1 , & p -> c );
emit ( "m" , 1 , & p -> m );
emit ( "e" , 1 , & p -> e );
}
void writer ( struct record * p , uint64_t value ) {
p -> m += value;
p -> e ^= value;
}
This is an excerpt from the full generated source . Function attributes and build flags preserve the call pattern the walkthrough relies on. main calls writer first, then serializer .
The emit function does nothing but mix the key, the length and the value at the given address into a volatile variable. That reproduces the machine-level pattern “key + value address → shared function” without a real JSON encoder or Serde. It is a synthetic C and assembly example for studying the mechanism. Assessing accuracy on real Rust applications requires a separate evaluation.
The assembly part adds 2,000 decoy functions. Each one writes at offsets +8 and +16 from its first argument. None of them is ever called. That is what keeps the question from being trivial: an analysis that just looks for matching offsets will produce plenty of convincing answers.
The program was built with GCC 13.3.0 using -O2 -g0 -fno-inline -fno-ipa-icf , then run through objcopy --strip-all . XRef reports 2,007 functions and zero symbols; in the recorded example, the boundaries of those functions come from .eh_frame , the stack-unwinding data. The analysis runs on the stripped file; the symbols of a separate reference build are used only to check the answer.
A key string is a starting point, not a field
Searching for m does not immediately tell you which memory location holds the value named m . A short string is an especially weak lead: its bytes can be part of a shared literal or be used in some other context. Those were exactly the cases that had to be handled after the request in issue 76. PR 77 added shared-literal recognition for short keys, per-key diagnostics and a way to inspect partial matches.
What I need is the relationship between two arguments at the use site: the key being passed and the value passed along with it. Only then does the name acquire a useful local meaning.
The first query in the recording looks for layout evidence only:
xref -b target-stripped field --keys c,m,e --layout-only
The report finds one candidate function, 0x11c0..0x1204 , with offsets relative to its incoming rdi . For the three keys it reports c -> 0x0 , m -> 0x8 and e -> 0x10 ; the corresponding instructions read eight bytes.
These values are useful because the report also lists the call sites where each key and value are passed. In the recorded image they are 0x11d7 , 0x11ec and 0x1202 ; the last one is a tail jump to the same callee. You can inspect the function that passes the keys instead of accepting the layout as an unexplained result.
Here is a shortened disassembly excerpt from the recording. The omitted instructions load the keys and their lengths:
0x11c5 mov rdx , rdi
0x11c8 mov rbx , rdi
0x11d7 call 0x11a0
0x11dc lea rdx , [ rbx + 8 ]
0x11ec call 0x11a0
0x11f1 lea rdx , [ rbx + 0x10 ]
0x1202 jmp short 0x11a0
On entry, rdi holds the pointer p . The copy saved in rbx survives the first call. So the values passed in rdx are p , p + 8 and p + 16 . In the full listing, the matching key string addresses and lengths are passed right next to them. That is how key names get tied to offsets.
Figure 1. Each arrow is a separate investigative step. Finding a later link does not automatically prove the earlier ones.
A field’s address and the value stored there are different things
One distinction shaped the problem more than any other: following a pointer and following a value loaded through that pointer are not the same thing.
Consider two schematic instructions. They illustrate the distinction and are not an additional result from the recorded test program:
lea rax , [ rdi + 8 ] ; address: incoming pointer + 8
mov rax , [ rdi + 8 ] ; value: eight bytes loaded there
After lea , rax still holds an address relative to the original input. After mov , it holds whatever was stored at that address: a number, a pointer to a different object, or a value the current analysis cannot classify. A later write through [rax+16] does not automatically become a write to rdi+24 .
That is why it matters where a base pointer came from: whether it arrived as an argument, was returned from a call or was loaded from memory. Each leads to a different conclusion about which object you are looking at. Lose track of a pointer’s origin, and later offsets can be wrongly attributed to the earlier object.
There is another practical trap here: searching only for lea can miss a field that is accessed through a direct load or store. An empty result from such a search does not prove the field was never used. It means the chosen instruction pattern was not found under those search conditions.
Figure 2. A schematic example: after a load from memory, you cannot keep tracking the original base by simply adding offsets.
An access range is not a struct declaration
The three offsets and read widths in this example give an addressed range of at least 0x18 bytes. It is tempting to write sizeof(record) = 24 next to it and move on.
Here, the source lets us look up the real declaration. The stripped binary on its own supports a narrower claim: those positions were accessed with those widths. It does not follow that there are no fields beyond them, that no field was missed, or that the allocated memory ends there. Alignment, tail padding and how the memory was allocated are separate questions.
The report expresses this as addressed-span>=0x18 and exact-size=unknown . It also keeps the lexical order of call sites separate from the source declaration order or any universal execution order. These caveats are a useful part of the result, not an apology tacked on after the analysis.
The same goes for names. We observed the byte strings c , m and e . We know the type name struct record from the source and the reference build’s symbols, not from the query result.
Explicit boundaries make the result more useful. You can search for writes at the observed offsets without first declaring an invented complete type.
Two thousand functions that look like the answer
Now let’s ask for functions that write both named fields through the same pointer:
xref -b target-stripped field --keys c,m,e \
--write --writes-all-of m,e --limit 5
The --

[truncated]
