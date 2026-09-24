---
source: "https://github.com/savarin/lean-agent"
hn_url: "https://news.ycombinator.com/item?id=49829771"
title: "Show HN: Link proof assistant Lean to Claude, fix your code's hidden assumptions"
article_title: "GitHub - savarin/lean-agent · GitHub"
image: "https://opengraph.githubassets.com/74710363081d32fa0ef55eaaa256f617327971a4f9ed20db8d2d7f75478cc2e4/savarin/lean-agent"
author: "kurinikku"
captured_at: "2026-09-24T13:17:51Z"
capture_tool: "hn-digest"
hn_id: 49829771
score: 2
comments: 0
posted_at: "2026-09-24T12:37:45Z"
tags:
  - hacker-news
---

# Show HN: Link proof assistant Lean to Claude, fix your code's hidden assumptions

- HN: [49829771](https://news.ycombinator.com/item?id=49829771)
- Source: [github.com](https://github.com/savarin/lean-agent)
- Score: 2
- Comments: 0
- Posted: 2026-09-24T12:37:45Z

## Translation

Title: Show HN: Link proof assistant Lean to Claude, fix your code's hidden assumptions
Article title: GitHub - savarin/lean-agent · GitHub
Description: Contribute to savarin/lean-agent development by creating an account on GitHub.

Article text:
GitHub - savarin/lean-agent · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
savarin
/
lean-agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
.claude-plugin .claude-plugin skills/ enforce skills/ enforce src/ lean_agent src/ lean_agent tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
lean-agent finds the assumptions your code makes but doesn't enforce — then fixes them. It measures each assumption as an Invariant Enforcement Score (IES), iterates until the score plateaus, and leaves you a branch to review. Your tests run every iteration; failures are discarded.
# before — IES: 0.00
def transfer ( self , from_id , to_id , amount ):
self . accounts [ from_id ] -= amount
self . accounts [ to_id ] += amount
# after — IES: 1.00
def transfer ( self , from_id : AccountId , to_id : AccountId , amount : PositiveAmount ) -> None :
# ... validates existence, positivity, sufficient balance ...
self . accounts = {
** self . accounts ,
from_id : self . accounts [ from_id ] - amount ,
to_id : self . accounts [ to_id ] + amount ,
}
uv pip install lean-agent
lean-agent enforce < repo >
lean-agent score < repo > --min-ies 0.80 # CI gate
Here's a 13-line ledger.
class Ledger :
def __init__ ( self ):
self . accounts = {}
def create_account ( self , account_id , opening_balance ):
self . accounts [ account_id ] = opening_balance
def transfer ( self , from_id , to_id , amount ):
self . accounts [ from_id ] -= amount
self . accounts [ to_id ] += amount
def balance ( self , account_id ):
return self . accounts [ account_id ]
Read the transfer method. Two lines. Say from_id is "alice" and to_id is "bob" . Alice has $100, Bob has $50. We transfer $30.
Line 1: self.accounts["alice"] -= 30 . Alice now has $70. This is in the dict.
Line 2: self.accounts["bob"] += 30 . Bob now has $80. Total was $150, still $150.
Now: what if to_id doesn't exist? Line 1 succeeds — Alice has $70. Line 2 throws KeyError . Total was $150. Total is now $120. The $30 is gone.
If the caller catches the exception, they see the KeyError . They don't see that line 1 already succeeded. The error says nothing about the debit. The ledger is silently unbalanced — money debited, never credited.
This is the class of bug that matters most. Not the ones that crash loudly — the ones where the damage is done before anything visibly fails. The system keeps running. The ledger keeps serving balances. The numbers just don't add up anymore.
Every codebase is full of these. The question is: how do you find them?
What code assumes but doesn't enforce
Look at transfer again. It makes five assumptions:
from_id exists in self.accounts
amount is positive (negative reverses the direction — silently)
amount is numeric (pass a string, get TypeError deep in arithmetic)
The source has enough balance (no overdraft protection)
None are checked. None are documented. None are enforced by the type system. They're just... hoped for.
We'll call these invariants — properties that must be true for the system to produce correct results. Each invariant sits at some enforcement level:
The Invariant Enforcement Score (IES) is the average, normalized to 0–1:
IES = Σ(scores) / (3 × N)
The ledger scores 0.00. Every invariant unguarded.
The interesting thing isn't the number. It's what happens when you try to move it. I want to follow one of these — transfer atomicity — all the way through: how formalization surfaces it, how it scores, and what the fix looks like at each enforcement level.
What Lean reveals about atomicity
You could stare at transfer and eventually notice the atomicity bug. For 13 lines, maybe. For 13,000 lines across multiple modules, probably not.
lean-agent uses Lean 4 — a proof assistant — to formalize the domain. Not to prove theorems. The proofs are all sorry (Lean's way of saying "trust me on this"). The value is in writing the types.
To formalize transfer, you start by defining its preconditions — what must be true before a transfer can happen?
structure TransferPrecondition (l : Ledger) (from_id to_id : AccountId) (amt : PosAmount) where
sourceExists : l.hasAccount from_id
destExists : l.hasAccount to_id
sufficientBalance : ∃ (bal : Amount), l.lookup from_id = some bal ∧ bal ≥ amt.val
Three fields. Three assumptions. To construct a TransferPrecondition , you must provide all three. The Python code checks none of them.
Now try to write the conservation theorem — total balance is unchanged after transfer:
theorem transfer_conserves_total
(l : Ledger) (from_id to_id : AccountId) (amt : PosAmount)
(pre : TransferPrecondition l from_id to_id amt) :
True := by
sorry
The proof is sorry . But look at the signature. It requires TransferPrecondition as an argument. Why? Because without it, the theorem doesn't type-check. If the source account might not exist, l.lookup from_id returns none , and you can't do arithmetic on none . Conservation literally cannot be stated without the precondition that both accounts exist.
This is the discovery moment. You weren't looking for the atomicity bug. You were trying to state what "correct transfer" means. And the type checker told you: you can't state it unless both sides are guaranteed to succeed. If either side might fail, conservation is not a property of this code — it's a wish.
You don't need Lean for this. You could write the preconditions on a whiteboard. The value of the tool is that it won't let you be vague. hasAccount from_id — yes or no? The code either checks or it doesn't.
The invariant: transfer must be atomic. Both sides happen, or neither does. The code doesn't enforce this. Score: 0 .
Validated (score 2). Add checks before either mutation:
def transfer ( self , from_id , to_id , amount ):
if from_id not in self . accounts :
raise ValueError ( f"source account { from_id } does not exist" )
if to_id not in self . accounts :
raise ValueError ( f"destination account { to_id } does not exist" )
if self . accounts [ from_id ] < amount :
raise ValueError ( f"insufficient balance" )
self . accounts [ from_id ] -= amount
self . accounts [ to_id ] += amount
All checks above the mutations. If any fails, we raise before touching the dict. The partial mutation can't happen. Score: 2 .
Good defensive programming. But notice what it relies on: the ordering of statements. The checks must come before the mutations. A future refactor that reorders them breaks the invariant. The enforcement is in the control flow, not in the structure of the data.
Structural (score 3). Replace two mutations with one expression:
self . accounts = {
** self . accounts ,
from_id : self . accounts [ from_id ] - amount ,
to_id : self . accounts [ to_id ] + amount ,
}
One expression. No intermediate state. The old dict has Alice at $100, the new dict has Alice at $70 and Bob at $80. self.accounts points to the old dict, then the new dict. There's no moment where Alice is at $70 and Bob is still at $50.
Score: 3 . The invariant isn't checked — it's structurally guaranteed. You can't write a partial mutation because there's only one mutation.
That's one invariant, start to finish: surfaced by formalization, scored at 0, hardened to 2, promoted to 3. The same process runs on every invariant lean-agent finds.
On the ledger, this takes the 13-line class from the top of this page to:
from types import MappingProxyType
from typing import NewType , TypeAlias
AccountId = NewType ( "AccountId" , str )
PositiveAmount = NewType ( "PositiveAmount" , float )
Account : TypeAlias = int | float
class Ledger :
"""Double-entry ledger. Balance is NonNeg — enforced by transfer validation."""
def __init__ ( self ) -> None :
self . accounts : dict [ AccountId , Account ] = {}
@ property
def accounts_view ( self ) -> MappingProxyType :
"""Read-only view of accounts. External code should use this."""
return MappingProxyType ( self . accounts )
def create_account ( self , account_id : AccountId , opening_balance : int | float ) -> None :
if not isinstance ( opening_balance , ( int , float )):
raise TypeError ( f"opening_balance must be numeric, got { type ( opening_balance ). __name__ } " )
if account_id in self . accounts :
raise ValueError ( f"account { account_id } already exists" )
self . accounts [ account_id ] = opening_balance
def transfer ( self , from_id : AccountId , to_id : AccountId , amount : PositiveAmount ) -> None :
if amount <= 0 :
raise ValueError ( f"transfer amount must be positive, got { amount } " )
if from_id not in self . accounts :
raise ValueError ( f"source account { from_id } does not exist" )
if to_id not in self . accounts :
raise ValueError ( f"destination account { to_id } does not exist" )
if self . accounts [ from_id ] < amount :
raise ValueError (
f"insufficient balance: account { from_id } has { self . accounts [ from_id ] } , need { amount } "
)
self . accounts = {
** self . accounts ,
from_id : self . accounts [ from_id ] - amount ,
to_id : self . accounts [ to_id ] + amount ,
}
def balance ( self , account_id : AccountId ) -> float | None :
return self . accounts . get ( account_id )
IES: 0.00 → 1.00. Eight invariants, all structural. The dict-spread on the transfer is the atomicity fix — one expression, no intermediate state where the debit happened but the credit hasn't. The MappingProxyType view means external code can't mu
[truncated]
How the score moves. IES: 0.00 → 1.00 in 8 iterations (6 kept, 2 discarded).
The big jump is iteration 1 — going from "nothing is checked" to "everything is checked at runtime" is the highest-leverage change. Iterations 2–8 are promotions: runtime checks (score 2) → type-level enforcement (score 3). Each promotion is smaller in score but deeper in safety. A runtime check catches the bug. A type makes the bug impossible to write.
Three phases. Each is a different kind of work.
Analyze. Read the codebase. Optionally formalize in Lean. Identify 5–15 critical invariants, prioritized by "invisible when broken" — things that fail silently, not things that crash loudly. Build a frozen evaluation harness ( prepare.py ) that scores each invariant by examining the source code. This harness is the ground truth for the entire run. It doesn't execute the code — it reads it, using AST parsing and pattern matching to detect enforcement levels. Once frozen, it never changes.
Execute. Work through pre-planned improvements. Each iteration: implement the change, commit, run the harness, run the test suite. If tests pass and the score improves, keep. If not, git reset --hard and move on. The simplicity criterion applies: prefer a 3-line runt

[truncated]

## Original Extract

Contribute to savarin/lean-agent development by creating an account on GitHub.

GitHub - savarin/lean-agent · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
savarin
/
lean-agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
.claude-plugin .claude-plugin skills/ enforce skills/ enforce src/ lean_agent src/ lean_agent tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
lean-agent finds the assumptions your code makes but doesn't enforce — then fixes them. It measures each assumption as an Invariant Enforcement Score (IES), iterates until the score plateaus, and leaves you a branch to review. Your tests run every iteration; failures are discarded.
# before — IES: 0.00
def transfer ( self , from_id , to_id , amount ):
self . accounts [ from_id ] -= amount
self . accounts [ to_id ] += amount
# after — IES: 1.00
def transfer ( self , from_id : AccountId , to_id : AccountId , amount : PositiveAmount ) -> None :
# ... validates existence, positivity, sufficient balance ...
self . accounts = {
** self . accounts ,
from_id : self . accounts [ from_id ] - amount ,
to_id : self . accounts [ to_id ] + amount ,
}
uv pip install lean-agent
lean-agent enforce < repo >
lean-agent score < repo > --min-ies 0.80 # CI gate
Here's a 13-line ledger.
class Ledger :
def __init__ ( self ):
self . accounts = {}
def create_account ( self , account_id , opening_balance ):
self . accounts [ account_id ] = opening_balance
def transfer ( self , from_id , to_id , amount ):
self . accounts [ from_id ] -= amount
self . accounts [ to_id ] += amount
def balance ( self , account_id ):
return self . accounts [ account_id ]
Read the transfer method. Two lines. Say from_id is "alice" and to_id is "bob" . Alice has $100, Bob has $50. We transfer $30.
Line 1: self.accounts["alice"] -= 30 . Alice now has $70. This is in the dict.
Line 2: self.accounts["bob"] += 30 . Bob now has $80. Total was $150, still $150.
Now: what if to_id doesn't exist? Line 1 succeeds — Alice has $70. Line 2 throws KeyError . Total was $150. Total is now $120. The $30 is gone.
If the caller catches the exception, they see the KeyError . They don't see that line 1 already succeeded. The error says nothing about the debit. The ledger is silently unbalanced — money debited, never credited.
This is the class of bug that matters most. Not the ones that crash loudly — the ones where the damage is done before anything visibly fails. The system keeps running. The ledger keeps serving balances. The numbers just don't add up anymore.
Every codebase is full of these. The question is: how do you find them?
What code assumes but doesn't enforce
Look at transfer again. It makes five assumptions:
from_id exists in self.accounts
amount is positive (negative reverses the direction — silently)
amount is numeric (pass a string, get TypeError deep in arithmetic)
The source has enough balance (no overdraft protection)
None are checked. None are documented. None are enforced by the type system. They're just... hoped for.
We'll call these invariants — properties that must be true for the system to produce correct results. Each invariant sits at some enforcement level:
The Invariant Enforcement Score (IES) is the average, normalized to 0–1:
IES = Σ(scores) / (3 × N)
The ledger scores 0.00. Every invariant unguarded.
The interesting thing isn't the number. It's what happens when you try to move it. I want to follow one of these — transfer atomicity — all the way through: how formalization surfaces it, how it scores, and what the fix looks like at each enforcement level.
What Lean reveals about atomicity
You could stare at transfer and eventually notice the atomicity bug. For 13 lines, maybe. For 13,000 lines across multiple modules, probably not.
lean-agent uses Lean 4 — a proof assistant — to formalize the domain. Not to prove theorems. The proofs are all sorry (Lean's way of saying "trust me on this"). The value is in writing the types.
To formalize transfer, you start by defining its preconditions — what must be true before a transfer can happen?
structure TransferPrecondition (l : Ledger) (from_id to_id : AccountId) (amt : PosAmount) where
sourceExists : l.hasAccount from_id
destExists : l.hasAccount to_id
sufficientBalance : ∃ (bal : Amount), l.lookup from_id = some bal ∧ bal ≥ amt.val
Three fields. Three assumptions. To construct a TransferPrecondition , you must provide all three. The Python code checks none of them.
Now try to write the conservation theorem — total balance is unchanged after transfer:
theorem transfer_conserves_total
(l : Ledger) (from_id to_id : AccountId) (amt : PosAmount)
(pre : TransferPrecondition l from_id to_id amt) :
True := by
sorry
The proof is sorry . But look at the signature. It requires TransferPrecondition as an argument. Why? Because without it, the theorem doesn't type-check. If the source account might not exist, l.lookup from_id returns none , and you can't do arithmetic on none . Conservation literally cannot be stated without the precondition that both accounts exist.
This is the discovery moment. You weren't looking for the atomicity bug. You were trying to state what "correct transfer" means. And the type checker told you: you can't state it unless both sides are guaranteed to succeed. If either side might fail, conservation is not a property of this code — it's a wish.
You don't need Lean for this. You could write the preconditions on a whiteboard. The value of the tool is that it won't let you be vague. hasAccount from_id — yes or no? The code either checks or it doesn't.
The invariant: transfer must be atomic. Both sides happen, or neither does. The code doesn't enforce this. Score: 0 .
Validated (score 2). Add checks before either mutation:
def transfer ( self , from_id , to_id , amount ):
if from_id not in self . accounts :
raise ValueError ( f"source account { from_id } does not exist" )
if to_id not in self . accounts :
raise ValueError ( f"destination account { to_id } does not exist" )
if self . accounts [ from_id ] < amount :
raise ValueError ( f"insufficient balance" )
self . accounts [ from_id ] -= amount
self . accounts [ to_id ] += amount
All checks above the mutations. If any fails, we raise before touching the dict. The partial mutation can't happen. Score: 2 .
Good defensive programming. But notice what it relies on: the ordering of statements. The checks must come before the mutations. A future refactor that reorders them breaks the invariant. The enforcement is in the control flow, not in the structure of the data.
Structural (score 3). Replace two mutations with one expression:
self . accounts = {
** self . accounts ,
from_id : self . accounts [ from_id ] - amount ,
to_id : self . accounts [ to_id ] + amount ,
}
One expression. No intermediate state. The old dict has Alice at $100, the new dict has Alice at $70 and Bob at $80. self.accounts points to the old dict, then the new dict. There's no moment where Alice is at $70 and Bob is still at $50.
Score: 3 . The invariant isn't checked — it's structurally guaranteed. You can't write a partial mutation because there's only one mutation.
That's one invariant, start to finish: surfaced by formalization, scored at 0, hardened to 2, promoted to 3. The same process runs on every invariant lean-agent finds.
On the ledger, this takes the 13-line class from the top of this page to:
from types import MappingProxyType
from typing import NewType , TypeAlias
AccountId = NewType ( "AccountId" , str )
PositiveAmount = NewType ( "PositiveAmount" , float )
Account : TypeAlias = int | float
class Ledger :
"""Double-entry ledger. Balance is NonNeg — enforced by transfer validation."""
def __init__ ( self ) -> None :
self . accounts : dict [ AccountId , Account ] = {}
@ property
def accounts_view ( self ) -> MappingProxyType :
"""Read-only view of accounts. External code should use this."""
return MappingProxyType ( self . accounts )
def create_account ( self , account_id : AccountId , opening_balance : int | float ) -> None :
if not isinstance ( opening_balance , ( int , float )):
raise TypeError ( f"opening_balance must be numeric, got { type ( opening_balance ). __name__ } " )
if account_id in self . accounts :
raise ValueError ( f"account { account_id } already exists" )
self . accounts [ account_id ] = opening_balance
def transfer ( self , from_id : AccountId , to_id : AccountId , amount : PositiveAmount ) -> None :
if amount <= 0 :
raise ValueError ( f"transfer amount must be positive, got { amount } " )
if from_id not in self . accounts :
raise ValueError ( f"source account { from_id } does not exist" )
if to_id not in self . accounts :
raise ValueError ( f"destination account { to_id } does not exist" )
if self . accounts [ from_id ] < amount :
raise ValueError (
f"insufficient balance: account { from_id } has { self . accounts [ from_id ] } , need { amount } "
)
self . accounts = {
** self . accounts ,
from_id : self . accounts [ from_id ] - amount ,
to_id : self . accounts [ to_id ] + amount ,
}
def balance ( self , account_id : AccountId ) -> float | None :
return self . accounts . get ( account_id )
IES: 0.00 → 1.00. Eight invariants, all structural. The dict-spread on the transfer is the atomicity fix — one expression, no intermediate state where the debit happened but the credit hasn't. The MappingProxyType view means external code can't mu
[truncated]
How the score moves. IES: 0.00 → 1.00 in 8 iterations (6 kept, 2 discarded).
The big jump is iteration 1 — going from "nothing is checked" to "everything is checked at runtime" is the highest-leverage change. Iterations 2–8 are promotions: runtime checks (score 2) → type-level enforcement (score 3). Each promotion is smaller in score but deeper in safety. A runtime check catches the bug. A type makes the bug impossible to write.
Three phases. Each is a different kind of work.
Analyze. Read the codebase. Optionally formalize in Lean. Identify 5–15 critical invariants, prioritized by "invisible when broken" — things that fail silently, not things that crash loudly. Build a frozen evaluation harness ( prepare.py ) that scores each invariant by examining the source code. This harness is the ground truth for the entire run. It doesn't execute the code — it reads it, using AST parsing and pattern matching to detect enforcement levels. Once frozen, it never changes.
Execute. Work through pre-planned improvements. Each iteration: implement the change, commit, run the harness, run the test suite. If tests pass and the score improves, keep. If not, git reset --hard and move on. The simplicity criterion applies: prefer a 3-line runt

[truncated]
