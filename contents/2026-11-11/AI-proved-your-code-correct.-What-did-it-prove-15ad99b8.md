---
source: "https://kellerai.io/resources/ai-proved-your-code-correct"
hn_url: "https://news.ycombinator.com/item?id=50448519"
title: "AI proved your code correct. What did it prove?"
article_title: "AI proved your code correct. What did it actually prove? — Keller AI"
image: "https://kellerai.io/resources/covers/ai-proved-your-code-correct.png"
author: "trehans"
captured_at: "2026-11-11T01:02:55Z"
capture_tool: "hn-digest"
hn_id: 50448519
score: 1
comments: 0
posted_at: "2026-11-11T01:01:30Z"
tags:
  - hacker-news
---

# AI proved your code correct. What did it prove?

- HN: [50448519](https://news.ycombinator.com/item?id=50448519)
- Source: [kellerai.io](https://kellerai.io/resources/ai-proved-your-code-correct)
- Score: 1
- Comments: 0
- Posted: 2026-11-11T01:01:30Z

## Translation

Title: AI proved your code correct. What did it prove?
Article title: AI proved your code correct. What did it actually prove? — Keller AI
Description: An interactive introduction to formal verification: run a sorting function, construct a geometric proof, and assemble an argument about permissions.

Article text:
AI proved your code correct. What did it actually prove? — Keller AI
Skip to main content
Keller AI Product
Engineering notes · Formal verification
AI proved your code correct. What did it actually prove?
How formal verification works, what it can establish about a program, and how its conclusions depend on the specification and assumptions.
Recent work from Anthropic and GitHub has brought more attention to formal verification in AI-assisted programming. In these projects, an agent produces code or proofs that can be checked against a mathematical specification. What does that check establish about the software?
In Anthropic’s work with Theorem, existing fuzz harnesses became the basis for formal statements about compiled programs. Claude constructed proofs that Rocq checked. GitHub Next and Microsoft Research are exploring a workflow in which a human reviews a specification and an agent’s implementation is checked against it using Verus.
Verification of 239 entry points across 123 libraries
The study establishes crash-freedom for inputs below 64 MB, under stated RISC-V, Linux and memory assumptions.
Proof-driven development with PDD Kit
This research prototype uses Verus to check Rust implementations against specifications, alongside tests and human review.
What does it mean to specify a program mathematically, and how can we prove that the code satisfies that specification?
We have spent a lot of time at Keller AI thinking about how software can be checked against the requirements and policies it is supposed to follow. Formal verification gives us a way to establish some of those properties precisely. Understanding how it works and where its guarantees end helps us choose what to prove, interpret the results, and recognize what still needs to be checked.
A sorting function gives us a small example where we can state what we want and examine the reasoning in detail.
What would we prove
about a sort?
Suppose we have some code that is supposed to sort a list of integers. What would convince us that it does so correctly?
There are a few properties we can check to give us confidence. The output should be ordered. It should have the same length as its input. It should contain the same elements as its input, including the same number of duplicates. For [4, 2, 4, 1] , we expect [1, 2, 4, 4] .
The example below checks these three properties for a given input. The alternative implementations each introduce a different error, so we can see which properties detect it.
Three properties of a sorting function
For input [4, 2, 4, 1], insertion sort produces [1, 2, 4, 4]: ordered, same length, same occurrences. Reversing that output breaks ordering. Returning four zeros preserves order and length but loses the original values.
The general argument begins with an empty list. A verified insertion helper preserves order and adds one occurrence. Applying that fact at every iteration establishes the invariant for the processed input prefix, then for the entire input.
Input integers Run → Up to 16 integers, each from −99 to 99. Duplicates Negatives Empty One element
Proof by induction Next step →
The checkmarks describe this input. Running examples is testing; it does not prove the function correct for every input. Nondecreasing order allows duplicates. “Same occurrences” means every value appears exactly as often as it did in the input; it also implies the same length. Each Lean toggle shows the property as the universal statement a proof would establish; the page checks only the current input.
Each incorrect implementation can pass some checks while failing others. A list of zeros can be ordered and the right length while containing the wrong elements; an empty list has no pair of elements out of order. The checks complement one another, covering different aspects of what we mean by sorting. When an implementation passes all of them across a range of inputs, we gain reasonable confidence that it sorts correctly.
Notice that the checks an incorrect implementation passes also depend on the input. Reversing the sorted result fails the ordering check for [4, 2, 4, 1] , but passes every check for [7] . This is the limitation of testing particular cases: a failing example can disprove a universal claim, but passing examples do not establish that it always holds.
But we have only run examples. How would we establish these properties for every input?
This is where proofs come in. For this sorting algorithm, we can use induction: establish that an empty list satisfies the properties, then show that each insertion preserves ordering and adds exactly the intended element. The argument applies no matter how many elements we process, allowing us to establish the properties for every finite list of integers without testing each one.
Induction is the proof technique used here. AI can construct such proofs, while a deterministic proof checker verifies that an accepted conclusion follows from the stated definitions and assumptions. To understand what makes this possible, we need to look more closely at how proofs are put together.
How can a proof cover
every valid case?
How can an argument establish something about inputs we have never tried? To understand that, we need to look at how a proof is built: which facts it starts from, and how those facts combine to establish something new. In the Elements , Euclid organized earlier mathematical knowledge into a system where proofs build on definitions, starting assumptions and previously proved results. A geometric construction makes those steps easier to see.
To illustrate this, let’s look at a geometric construction from the first proposition in Euclid’s Elements . Start with two distinct points, A and B. Draw a circle around each point with radius AB, and choose one intersection, C. We want to prove that the triangle ABC formed by these points is always equilateral.
Constructing an equilateral triangle
r r r C A B ✓ ABC is equilateral
Move A or B to try another construction.
Select a piece, then its partner to combine their equalities.
AC = AB C lies on the circle centered at A. Both lengths are its radius. C ∈ circle(A, |AB|)
⊢ |AC| = |AB| From Euclid’s definition of a circle Mathematics BC = AB C lies on the circle centered at B, with that same radius. C ∈ circle(B, |AB|)
⊢ |BC| = |AB| From Euclid’s definition of a circle Mathematics AC = BC = AB All three sides are equal, so the triangle is equilateral. |AC| = |AB| ∧ |BC| = |AB|
⊢ |AC| = |BC| = |AB| Derived from both equalities · transitivity Mathematics Each side: 150 diagram units.
We could move the points and measure the sides repeatedly. Like testing a function on different inputs, this checks individual cases. Measurements have limited precision, and the cases we try do not establish the claim for every possible placement.
What lets us prove it? Euclid’s postulates permit drawing the lines and circles. The definition of a circle makes every point on its circumference equally distant from its center. The equality principle lets us conclude that quantities equal to the same quantity are equal to each other. The proof combines these facts.
C lies on A’s circle, so AC equals AB. It also lies on B’s circle, so BC equals AB. Therefore AC, BC and AB are equal. No measurement was needed: moving A and B preserves the relationships that establish the result.
What can we do with a proved result? We can use it in another proof without repeating its argument, provided its assumptions apply. Later propositions reuse this construction. This is how a small set of foundations supports a growing body of mathematical knowledge.
Accumulation of verified knowledge
Proved results we can reuse in later proofs.
Combine available statements into a proof and check each step.
A modern proof assistant checks that each step follows from definitions, assumptions or established results. An agent may search for an argument and make mistakes; the checker determines whether the steps are justified. In software, reusable results include function contracts and facts about control flow. The sorting proof already used this approach: the insertion helper’s contract supported the argument about the loop.
Our partner Archway Labs works on translating code into formal mathematics that can be reasoned about in the style of Euclid. They explain how compositional code is important for combining statements to accumulate verified knowledge, and why it is valuable to make contextual assumptions explicit.
Lean checks formal proofs through its kernel; Verus uses automated solvers to verify supported Rust code. Both require precise claims and justified connections to the program.
How does this scale
to a real system?
Suppose we are adding a module that lets users download private documents. One property we want to establish is that every successful download is authorized for the person making the request. How can we prove that about the new code?
We can organize the work into two phases. First, we build a mathematical model of the relevant parts of the application. We define what users, documents and permissions mean, represent the behavior of existing functions, and establish their contracts. Reading source code can suggest a contract; proving that the implementation satisfies it is a further step. For external services, we record the assumptions we rely on .
This model, expressed in a formal language such as Lean, supplies both a vocabulary for specifications and reusable statements about the system. It distinguishes definitions, proved contracts and assumptions. We can extend it as new work requires more of the application to be modeled.
In the second phase, we express the new requirement using that vocabulary and verify the new code against it. The result may be a proof, a demonstrated violation, or an unresolved question with the current model and tools. A successful proof adds another reusable result, together with the assumptions it depends on.
Model behavior and establish contracts
Definitions
Proved contracts
Explicit assumptions
Represent code and formalize the claim
Proved · Violation found · Unresolved
The puzzle below illustrates the second phase. Its contracts for authenticate , lookupOwner and canRead come from our model of existing components. Here, we treat those contracts as assumptions. The guard is part of the new handler, and its behavior follows from the code. Combining these pieces lets us establish the download property.
In this example, permission to read an organization’s documents covers the requested document. The shared variables connect the pieces: the permission check must concern the same caller and the organization that owns that document.
Deriving the download permission
1 user = authenticate(request) 2 org = lookupOwner(document) 3 4 allowed = canRead(user, org) 5 if not allowed: 6 deny the request and stop 7 8 return load(document) Hover over or focus a piece to highlight the code it describes. Selecting a piece keeps that connection visible.
Every successful download is authorized for the authenticated caller.
Select a piece, then its matching partner to derive a new statement. Either branch can be completed first.
Document permission: the ownership lookup identifies the document’s organization; the permission contract establishes that the user may read in that organization.
Authenticated return: the handler returns only after those checks succeed, and the authentication contract identifies the checked user as the caller.
Together, these statements establish an authorized download, provided the component contracts and shared assumptions hold.
Ownership contract: the lookup returns the requested document’s owning organization. Its correctness depends on the query and the ownership records.
Permission contract: an allowed decision means the user may read the organization’s documents in authorization snapshot s.
Authentication contract: the returned user is the authenticated caller of this request. We assume the se

[truncated]

## Original Extract

An interactive introduction to formal verification: run a sorting function, construct a geometric proof, and assemble an argument about permissions.

AI proved your code correct. What did it actually prove? — Keller AI
Skip to main content
Keller AI Product
Engineering notes · Formal verification
AI proved your code correct. What did it actually prove?
How formal verification works, what it can establish about a program, and how its conclusions depend on the specification and assumptions.
Recent work from Anthropic and GitHub has brought more attention to formal verification in AI-assisted programming. In these projects, an agent produces code or proofs that can be checked against a mathematical specification. What does that check establish about the software?
In Anthropic’s work with Theorem, existing fuzz harnesses became the basis for formal statements about compiled programs. Claude constructed proofs that Rocq checked. GitHub Next and Microsoft Research are exploring a workflow in which a human reviews a specification and an agent’s implementation is checked against it using Verus.
Verification of 239 entry points across 123 libraries
The study establishes crash-freedom for inputs below 64 MB, under stated RISC-V, Linux and memory assumptions.
Proof-driven development with PDD Kit
This research prototype uses Verus to check Rust implementations against specifications, alongside tests and human review.
What does it mean to specify a program mathematically, and how can we prove that the code satisfies that specification?
We have spent a lot of time at Keller AI thinking about how software can be checked against the requirements and policies it is supposed to follow. Formal verification gives us a way to establish some of those properties precisely. Understanding how it works and where its guarantees end helps us choose what to prove, interpret the results, and recognize what still needs to be checked.
A sorting function gives us a small example where we can state what we want and examine the reasoning in detail.
What would we prove
about a sort?
Suppose we have some code that is supposed to sort a list of integers. What would convince us that it does so correctly?
There are a few properties we can check to give us confidence. The output should be ordered. It should have the same length as its input. It should contain the same elements as its input, including the same number of duplicates. For [4, 2, 4, 1] , we expect [1, 2, 4, 4] .
The example below checks these three properties for a given input. The alternative implementations each introduce a different error, so we can see which properties detect it.
Three properties of a sorting function
For input [4, 2, 4, 1], insertion sort produces [1, 2, 4, 4]: ordered, same length, same occurrences. Reversing that output breaks ordering. Returning four zeros preserves order and length but loses the original values.
The general argument begins with an empty list. A verified insertion helper preserves order and adds one occurrence. Applying that fact at every iteration establishes the invariant for the processed input prefix, then for the entire input.
Input integers Run → Up to 16 integers, each from −99 to 99. Duplicates Negatives Empty One element
Proof by induction Next step →
The checkmarks describe this input. Running examples is testing; it does not prove the function correct for every input. Nondecreasing order allows duplicates. “Same occurrences” means every value appears exactly as often as it did in the input; it also implies the same length. Each Lean toggle shows the property as the universal statement a proof would establish; the page checks only the current input.
Each incorrect implementation can pass some checks while failing others. A list of zeros can be ordered and the right length while containing the wrong elements; an empty list has no pair of elements out of order. The checks complement one another, covering different aspects of what we mean by sorting. When an implementation passes all of them across a range of inputs, we gain reasonable confidence that it sorts correctly.
Notice that the checks an incorrect implementation passes also depend on the input. Reversing the sorted result fails the ordering check for [4, 2, 4, 1] , but passes every check for [7] . This is the limitation of testing particular cases: a failing example can disprove a universal claim, but passing examples do not establish that it always holds.
But we have only run examples. How would we establish these properties for every input?
This is where proofs come in. For this sorting algorithm, we can use induction: establish that an empty list satisfies the properties, then show that each insertion preserves ordering and adds exactly the intended element. The argument applies no matter how many elements we process, allowing us to establish the properties for every finite list of integers without testing each one.
Induction is the proof technique used here. AI can construct such proofs, while a deterministic proof checker verifies that an accepted conclusion follows from the stated definitions and assumptions. To understand what makes this possible, we need to look more closely at how proofs are put together.
How can a proof cover
every valid case?
How can an argument establish something about inputs we have never tried? To understand that, we need to look at how a proof is built: which facts it starts from, and how those facts combine to establish something new. In the Elements , Euclid organized earlier mathematical knowledge into a system where proofs build on definitions, starting assumptions and previously proved results. A geometric construction makes those steps easier to see.
To illustrate this, let’s look at a geometric construction from the first proposition in Euclid’s Elements . Start with two distinct points, A and B. Draw a circle around each point with radius AB, and choose one intersection, C. We want to prove that the triangle ABC formed by these points is always equilateral.
Constructing an equilateral triangle
r r r C A B ✓ ABC is equilateral
Move A or B to try another construction.
Select a piece, then its partner to combine their equalities.
AC = AB C lies on the circle centered at A. Both lengths are its radius. C ∈ circle(A, |AB|)
⊢ |AC| = |AB| From Euclid’s definition of a circle Mathematics BC = AB C lies on the circle centered at B, with that same radius. C ∈ circle(B, |AB|)
⊢ |BC| = |AB| From Euclid’s definition of a circle Mathematics AC = BC = AB All three sides are equal, so the triangle is equilateral. |AC| = |AB| ∧ |BC| = |AB|
⊢ |AC| = |BC| = |AB| Derived from both equalities · transitivity Mathematics Each side: 150 diagram units.
We could move the points and measure the sides repeatedly. Like testing a function on different inputs, this checks individual cases. Measurements have limited precision, and the cases we try do not establish the claim for every possible placement.
What lets us prove it? Euclid’s postulates permit drawing the lines and circles. The definition of a circle makes every point on its circumference equally distant from its center. The equality principle lets us conclude that quantities equal to the same quantity are equal to each other. The proof combines these facts.
C lies on A’s circle, so AC equals AB. It also lies on B’s circle, so BC equals AB. Therefore AC, BC and AB are equal. No measurement was needed: moving A and B preserves the relationships that establish the result.
What can we do with a proved result? We can use it in another proof without repeating its argument, provided its assumptions apply. Later propositions reuse this construction. This is how a small set of foundations supports a growing body of mathematical knowledge.
Accumulation of verified knowledge
Proved results we can reuse in later proofs.
Combine available statements into a proof and check each step.
A modern proof assistant checks that each step follows from definitions, assumptions or established results. An agent may search for an argument and make mistakes; the checker determines whether the steps are justified. In software, reusable results include function contracts and facts about control flow. The sorting proof already used this approach: the insertion helper’s contract supported the argument about the loop.
Our partner Archway Labs works on translating code into formal mathematics that can be reasoned about in the style of Euclid. They explain how compositional code is important for combining statements to accumulate verified knowledge, and why it is valuable to make contextual assumptions explicit.
Lean checks formal proofs through its kernel; Verus uses automated solvers to verify supported Rust code. Both require precise claims and justified connections to the program.
How does this scale
to a real system?
Suppose we are adding a module that lets users download private documents. One property we want to establish is that every successful download is authorized for the person making the request. How can we prove that about the new code?
We can organize the work into two phases. First, we build a mathematical model of the relevant parts of the application. We define what users, documents and permissions mean, represent the behavior of existing functions, and establish their contracts. Reading source code can suggest a contract; proving that the implementation satisfies it is a further step. For external services, we record the assumptions we rely on .
This model, expressed in a formal language such as Lean, supplies both a vocabulary for specifications and reusable statements about the system. It distinguishes definitions, proved contracts and assumptions. We can extend it as new work requires more of the application to be modeled.
In the second phase, we express the new requirement using that vocabulary and verify the new code against it. The result may be a proof, a demonstrated violation, or an unresolved question with the current model and tools. A successful proof adds another reusable result, together with the assumptions it depends on.
Model behavior and establish contracts
Definitions
Proved contracts
Explicit assumptions
Represent code and formalize the claim
Proved · Violation found · Unresolved
The puzzle below illustrates the second phase. Its contracts for authenticate , lookupOwner and canRead come from our model of existing components. Here, we treat those contracts as assumptions. The guard is part of the new handler, and its behavior follows from the code. Combining these pieces lets us establish the download property.
In this example, permission to read an organization’s documents covers the requested document. The shared variables connect the pieces: the permission check must concern the same caller and the organization that owns that document.
Deriving the download permission
1 user = authenticate(request) 2 org = lookupOwner(document) 3 4 allowed = canRead(user, org) 5 if not allowed: 6 deny the request and stop 7 8 return load(document) Hover over or focus a piece to highlight the code it describes. Selecting a piece keeps that connection visible.
Every successful download is authorized for the authenticated caller.
Select a piece, then its matching partner to derive a new statement. Either branch can be completed first.
Document permission: the ownership lookup identifies the document’s organization; the permission contract establishes that the user may read in that organization.
Authenticated return: the handler returns only after those checks succeed, and the authentication contract identifies the checked user as the caller.
Together, these statements establish an authorized download, provided the component contracts and shared assumptions hold.
Ownership contract: the lookup returns the requested document’s owning organization. Its correctness depends on the query and the ownership records.
Permission contract: an allowed decision means the user may read the organization’s documents in authorization snapshot s.
Authentication contract: the returned user is the authenticated caller of this request. We assume the se

[truncated]
