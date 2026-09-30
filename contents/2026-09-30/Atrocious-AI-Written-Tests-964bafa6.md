---
source: "https://gruhn.me/blog/2026-09-29/"
hn_url: "https://news.ycombinator.com/item?id=49905603"
title: "Atrocious AI-Written Tests"
article_title: "Niklas Gruhn - Atrocious AI-written tests"
image: ""
author: "ngruhn"
captured_at: "2026-09-30T08:06:16Z"
capture_tool: "hn-digest"
hn_id: 49905603
score: 1
comments: 0
posted_at: "2026-09-30T07:29:27Z"
tags:
  - hacker-news
---

# Atrocious AI-Written Tests

- HN: [49905603](https://news.ycombinator.com/item?id=49905603)
- Source: [gruhn.me](https://gruhn.me/blog/2026-09-29/)
- Score: 1
- Comments: 0
- Posted: 2026-09-30T07:29:27Z

## Translation

Title: Atrocious AI-Written Tests
Article title: Niklas Gruhn - Atrocious AI-written tests

Article text:
Niklas Gruhn - Atrocious AI-written tests
Niklas Gruhn
Atrocious AI-written tests
I stopped reviewing tests.
It's too tempting when there are 10,000 other lines to review.
But recently I got a failing test after changing a comment in a source file.
I had to see what on earth these tests are doing that we're running all day, every day.
That particular test was scanning source files with a regex for the keyword "wall-clock".
My bad for using that in a comment.
Apparently, a certain "wall-clock error" was retired at some point.
I guess we want to make sure it's not re-introduced?
it ( "no source emits wall-clock error text" , ( ) => {
const sourceFiles = [ /* ... */ ] ;
for ( const path of sourceFiles ) {
const content = readFileSync ( path ) ;
expect ( content ) . not . toMatch ( / wall.?clock / i ) ;
}
} ) ;
At least that test has "code coverage".
Digging further, I found this gem:
it ( "default visible statuses are queued, running, drafted, completed" , ( ) => {
const statuses = [ "queued" , "running" , "drafted" , "completed" ] ;
expect ( statuses ) . not . toContain ( "failed" ) ;
expect ( statuses ) . not . toContain ( "cancelled" ) ;
} ) ;
The next one also runs zero production code.
But it "mirrors" production code somewhere else: 👏
it ( "filter description is a comma-joined list in lifecycle order" , ( ) => {
// Mirrors describeFilter() in the component:
const ALL = [ "queued" , "running" , "drafted" , "completed" , "failed" , "cancelled" ] as const ;
const describe = ( visible : Set < string > ) =>
ALL . filter ( ( s ) => visible . has ( s ) ) . join ( ", " ) ;
expect ( describe ( new Set ( ) ) ) . toBe ( "" ) ;
expect ( describe ( new Set ( [ "completed" ] ) ) ) . toBe ( "completed" ) ;
expect ( describe ( new Set ( [ "failed" , "queued" ] ) ) ) . toBe ( "queued, failed" ) ;
expect ( describe ( new Set ( ALL ) ) ) . toBe ( "queued, running, drafted, completed, failed, cancelled" ) ;
} ) ;
Even more common and barely more useful: checking static config values.
What regression is that protecting against?
Sure, config changes can cause bugs, but those emerge in downstream logic.
When this test fails, all it tells me is: you changed the config .
// code:
export const DEFAULT_OPTIONS = {
maxParallel : 75 ,
maxPerUser : 15 ,
queueTimeoutMs : 120_000 ,
} ;
// test:
test ( "defaults: 75 parallel, 15 per user, 120s queue wait" , async ( ) => {
expect ( DEFAULT_OPTIONS ) . toEqual ( {
maxParallel : 75 ,
maxPerUser : 15 ,
queueTimeoutMs : 120_000 ,
} ) ;
} ) ;
Another common pattern: testing for exact phrases in prompts.
Maybe someone thinks this is useful, but I disagree.
How prompts affect model output is highly nondeterministic.
Monitor that with evals, not unit tests.
it ( "keeps every prompt section heading present" , ( ) => {
const prompt = buildSupportPrompt ( makeCtx ( ) , makeConversation ( ) ) ;
for ( const heading of [
"## Grounding — in priority order" ,
"## Routing — questions outside teams ownership" ,
"## Cross-team work requests — cc for visibility" ,
"## Capabilities — describe them accurately" ,
"## The team" ,
"## Memory rules — strictly follow" ,
"## When you cannot help" ,
"## PII discipline — strictly follow" ,
"## Answer style — strictly follow" ,
"## Question" ,
] ) {
expect ( prompt ) . toContain ( heading ) ;
}
} ) ;
Conclusion
These config/prompt tests are plain silly.
But the agents are not stupid.
Ask for "low value tests in the codebase" and Claude/GPT have no trouble identifying those.
I suspect the agents think pairing any change with tests is a good look ,
so they sprinkle them in no matter what.
Also, I never see agents bending the implementation to make it more testable.
The "wall-clock" example hints at an architectural limitation.
Never mind whether asserting error message content makes sense in the first place,
but if you want to do it,
logging could be abstracted into a shared module and intercepted during testing.
In total there were 404 tests like this (of 6,429).
Is this a problem at all?
These tests did not meaningfully slow down the suite, they had no production impact, and removing them was easy.
Sure, you get nonsense red tests on virtually any change and burn some tokens fixing them.
But I think the real issue is loss of confidence.
I have no idea whether these tests cover important requirements, and I can't trust AI to have my back.

## Original Extract

Niklas Gruhn - Atrocious AI-written tests
Niklas Gruhn
Atrocious AI-written tests
I stopped reviewing tests.
It's too tempting when there are 10,000 other lines to review.
But recently I got a failing test after changing a comment in a source file.
I had to see what on earth these tests are doing that we're running all day, every day.
That particular test was scanning source files with a regex for the keyword "wall-clock".
My bad for using that in a comment.
Apparently, a certain "wall-clock error" was retired at some point.
I guess we want to make sure it's not re-introduced?
it ( "no source emits wall-clock error text" , ( ) => {
const sourceFiles = [ /* ... */ ] ;
for ( const path of sourceFiles ) {
const content = readFileSync ( path ) ;
expect ( content ) . not . toMatch ( / wall.?clock / i ) ;
}
} ) ;
At least that test has "code coverage".
Digging further, I found this gem:
it ( "default visible statuses are queued, running, drafted, completed" , ( ) => {
const statuses = [ "queued" , "running" , "drafted" , "completed" ] ;
expect ( statuses ) . not . toContain ( "failed" ) ;
expect ( statuses ) . not . toContain ( "cancelled" ) ;
} ) ;
The next one also runs zero production code.
But it "mirrors" production code somewhere else: 👏
it ( "filter description is a comma-joined list in lifecycle order" , ( ) => {
// Mirrors describeFilter() in the component:
const ALL = [ "queued" , "running" , "drafted" , "completed" , "failed" , "cancelled" ] as const ;
const describe = ( visible : Set < string > ) =>
ALL . filter ( ( s ) => visible . has ( s ) ) . join ( ", " ) ;
expect ( describe ( new Set ( ) ) ) . toBe ( "" ) ;
expect ( describe ( new Set ( [ "completed" ] ) ) ) . toBe ( "completed" ) ;
expect ( describe ( new Set ( [ "failed" , "queued" ] ) ) ) . toBe ( "queued, failed" ) ;
expect ( describe ( new Set ( ALL ) ) ) . toBe ( "queued, running, drafted, completed, failed, cancelled" ) ;
} ) ;
Even more common and barely more useful: checking static config values.
What regression is that protecting against?
Sure, config changes can cause bugs, but those emerge in downstream logic.
When this test fails, all it tells me is: you changed the config .
// code:
export const DEFAULT_OPTIONS = {
maxParallel : 75 ,
maxPerUser : 15 ,
queueTimeoutMs : 120_000 ,
} ;
// test:
test ( "defaults: 75 parallel, 15 per user, 120s queue wait" , async ( ) => {
expect ( DEFAULT_OPTIONS ) . toEqual ( {
maxParallel : 75 ,
maxPerUser : 15 ,
queueTimeoutMs : 120_000 ,
} ) ;
} ) ;
Another common pattern: testing for exact phrases in prompts.
Maybe someone thinks this is useful, but I disagree.
How prompts affect model output is highly nondeterministic.
Monitor that with evals, not unit tests.
it ( "keeps every prompt section heading present" , ( ) => {
const prompt = buildSupportPrompt ( makeCtx ( ) , makeConversation ( ) ) ;
for ( const heading of [
"## Grounding — in priority order" ,
"## Routing — questions outside teams ownership" ,
"## Cross-team work requests — cc for visibility" ,
"## Capabilities — describe them accurately" ,
"## The team" ,
"## Memory rules — strictly follow" ,
"## When you cannot help" ,
"## PII discipline — strictly follow" ,
"## Answer style — strictly follow" ,
"## Question" ,
] ) {
expect ( prompt ) . toContain ( heading ) ;
}
} ) ;
Conclusion
These config/prompt tests are plain silly.
But the agents are not stupid.
Ask for "low value tests in the codebase" and Claude/GPT have no trouble identifying those.
I suspect the agents think pairing any change with tests is a good look ,
so they sprinkle them in no matter what.
Also, I never see agents bending the implementation to make it more testable.
The "wall-clock" example hints at an architectural limitation.
Never mind whether asserting error message content makes sense in the first place,
but if you want to do it,
logging could be abstracted into a shared module and intercepted during testing.
In total there were 404 tests like this (of 6,429).
Is this a problem at all?
These tests did not meaningfully slow down the suite, they had no production impact, and removing them was easy.
Sure, you get nonsense red tests on virtually any change and burn some tokens fixing them.
But I think the real issue is loss of confidence.
I have no idea whether these tests cover important requirements, and I can't trust AI to have my back.
