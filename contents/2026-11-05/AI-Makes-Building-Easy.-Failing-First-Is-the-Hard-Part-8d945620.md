---
source: "https://brian.biz/blog/fail-fast-fail-often/"
hn_url: "https://news.ycombinator.com/item?id=50383138"
title: "AI Makes Building Easy. Failing First Is the Hard Part"
article_title: "AI Makes Building Easy. Failing First Is the Hard Part. · Brian Woods"
image: "https://brian.biz/static/leeroy_chainkins.png"
author: "BrianWoods"
captured_at: "2026-11-05T19:35:46Z"
capture_tool: "hn-digest"
hn_id: 50383138
score: 1
comments: 0
posted_at: "2026-11-05T19:14:37Z"
tags:
  - hacker-news
---

# AI Makes Building Easy. Failing First Is the Hard Part

- HN: [50383138](https://news.ycombinator.com/item?id=50383138)
- Source: [brian.biz](https://brian.biz/blog/fail-fast-fail-often/)
- Score: 1
- Comments: 0
- Posted: 2026-11-05T19:14:37Z

## Translation

Title: AI Makes Building Easy. Failing First Is the Hard Part
Article title: AI Makes Building Easy. Failing First Is the Hard Part. · Brian Woods
Description: Building is nearly free now, which makes thinking the scarce resource. Write the hypothesis, define success and failure, and pick a kill date before you ship.

Article text:
AI Makes Building Easy. Failing First Is the Hard Part.
Or: how to be wrong quickly, cheaply, and on purpose.
Everyone has a guy. A friend, a cousin, an uncle. Let's call him Dave. Dave is made
up, but you know him anyway.
Last spring Dave discovered AI, and by Thursday he was an “app guy.” He built
an app that matches dog owners who live exactly 0.4 miles apart, in an afternoon, on his
phone, during the game. He has a logo. He has a waitlist: three cousins and a man from his
gym who was being polite.
Ask Dave who the customer is and he says “dog owners.” Ask how he'll know
it's working and he says “we'll see how it goes.” Ask what would make him quit
and he looks at you like you asked him to disown a child. Dave isn't stupid. Dave is
fast . AI taught him how to build. It didn't teach him how to find out whether
anyone wants the thing.
That's the gap this post is about. Building is easy now, but innovation was
never about the building. It's about knowing what to build, for whom, and how to learn
you're wrong before you fall in love with it. The fundamentals didn't get easier. They
got easier to skip.
What “fail fast” actually means
It doesn't mean “screw up constantly and write a thread about it.” The idea
is simple: if something is going to go wrong, find out immediately and
cheaply . It began in engineering. In 1985, Jim Gray described “fail-fast”
modules at Tandem Computers: components that stop the instant they detect a fault instead
of quietly corrupting everything downstream. Jim Shore brought the phrase to everyday
programmers in 2004. Then Silicon Valley turned it into a bumper sticker, with help from
Steve Blank, Eric Ries's The Lean Startup (2011), and Facebook's “move fast
and break things.”
The bumper sticker skips the part that matters: failure only helps if you can
see it and learn from it. A failure nobody notices isn't a lesson. It's just a
failure with good PR.
Zappos is almost insultingly cheap. In 1999, Nick Swinmurn believed people would buy
shoes online without trying them on. The grown-up test was a warehouse and an inventory.
He didn't have the money to be that grown-up, so he photographed shoes in local stores and
put the photos on a website. When someone ordered, he bought the pair at full price and
shipped it himself. He lost money on every sale, and it didn't matter. He was testing
one question, and he had his answer before he had anything to sell. The site became
Zappos, and Amazon bought it in 2009 for a reported $1.2 billion.
Humane's AI Pin bet that people would ditch their phone screen for a voice-driven
device on their lapel. A rough prototype and a few dozen volunteers could have started
answering that in a week. Instead, Humane raised hundreds of millions for custom hardware
and launched in 2024 at $699 plus a subscription. The first moment a customer could say
“no thanks” came after the money was spent. In February 2025, Humane sold its
team and technology to HP for $116 million, and the Pins stopped working within days.
Nothing says “we hear you” like a very expensive lapel ornament.
The idea may well have been wrong, and that's allowed. The tragedy was the price of
finding out. Zappos paid almost nothing to learn its hypothesis was true. Humane paid
everything to learn it was false.
AI adds a new way to fail, because building got cheap. In a
postmortem published in August 2026 ,
solo founder Gus Chiriboga described using AI coding agents to build a tool for supervising
AI coding agents. He shipped 128 releases with CI, tests, bilingual docs and governance
policies. External users: zero. Very thorough governance for a country with no citizens.
His second attempt had 442 passing tests, and broke the moment he used it for real,
because the tests checked his own model of the system.
His diagnosis, after about two years: technical work was how he avoided real users.
“Agents make building so cheap that building becomes the procrastination.” And:
“Tests validate your model of the system. Only strangers validate the system.”
Put both on a mug.
Thinking is the scarce resource
The new loop is: build it, ship it, “learn from the data.” It sounds
responsible. It's a trap wearing a lanyard. You can't learn from data when you never wrote
down what you expected, so every result looks like a signal and nothing ever clearly
fails. A flat chart is “early.” A silent launch is “building
awareness.” Soon you're curating a graveyard of half-used features, each one
“kind of promising.”
Failing fast is cheap now. Failing legibly is still your job. Cheap building
made thinking the only scarce resource left, and the machine will happily build the wrong
thing at astonishing speed without ever asking, “Are you sure?”
Write the hypothesis in one sentence: “We believe [user] will
[behavior] because [reason].” If you can't fill in the blanks, you're not ready to
build. You're ready to think.
Define success and failure before you ship. After is when
everything looks like a win.
Pick a kill date. Put it on a calendar, and tell someone who will
hold you to it.
So write the hypothesis. Set the kill date. You have the most powerful
building tools in history, so aim them at the right thing, and learn something every
time. Every honest failure is a course correction, and course corrections are how you
reach the moon.
Now go put your work in front of a stranger.

## Original Extract

Building is nearly free now, which makes thinking the scarce resource. Write the hypothesis, define success and failure, and pick a kill date before you ship.

AI Makes Building Easy. Failing First Is the Hard Part.
Or: how to be wrong quickly, cheaply, and on purpose.
Everyone has a guy. A friend, a cousin, an uncle. Let's call him Dave. Dave is made
up, but you know him anyway.
Last spring Dave discovered AI, and by Thursday he was an “app guy.” He built
an app that matches dog owners who live exactly 0.4 miles apart, in an afternoon, on his
phone, during the game. He has a logo. He has a waitlist: three cousins and a man from his
gym who was being polite.
Ask Dave who the customer is and he says “dog owners.” Ask how he'll know
it's working and he says “we'll see how it goes.” Ask what would make him quit
and he looks at you like you asked him to disown a child. Dave isn't stupid. Dave is
fast . AI taught him how to build. It didn't teach him how to find out whether
anyone wants the thing.
That's the gap this post is about. Building is easy now, but innovation was
never about the building. It's about knowing what to build, for whom, and how to learn
you're wrong before you fall in love with it. The fundamentals didn't get easier. They
got easier to skip.
What “fail fast” actually means
It doesn't mean “screw up constantly and write a thread about it.” The idea
is simple: if something is going to go wrong, find out immediately and
cheaply . It began in engineering. In 1985, Jim Gray described “fail-fast”
modules at Tandem Computers: components that stop the instant they detect a fault instead
of quietly corrupting everything downstream. Jim Shore brought the phrase to everyday
programmers in 2004. Then Silicon Valley turned it into a bumper sticker, with help from
Steve Blank, Eric Ries's The Lean Startup (2011), and Facebook's “move fast
and break things.”
The bumper sticker skips the part that matters: failure only helps if you can
see it and learn from it. A failure nobody notices isn't a lesson. It's just a
failure with good PR.
Zappos is almost insultingly cheap. In 1999, Nick Swinmurn believed people would buy
shoes online without trying them on. The grown-up test was a warehouse and an inventory.
He didn't have the money to be that grown-up, so he photographed shoes in local stores and
put the photos on a website. When someone ordered, he bought the pair at full price and
shipped it himself. He lost money on every sale, and it didn't matter. He was testing
one question, and he had his answer before he had anything to sell. The site became
Zappos, and Amazon bought it in 2009 for a reported $1.2 billion.
Humane's AI Pin bet that people would ditch their phone screen for a voice-driven
device on their lapel. A rough prototype and a few dozen volunteers could have started
answering that in a week. Instead, Humane raised hundreds of millions for custom hardware
and launched in 2024 at $699 plus a subscription. The first moment a customer could say
“no thanks” came after the money was spent. In February 2025, Humane sold its
team and technology to HP for $116 million, and the Pins stopped working within days.
Nothing says “we hear you” like a very expensive lapel ornament.
The idea may well have been wrong, and that's allowed. The tragedy was the price of
finding out. Zappos paid almost nothing to learn its hypothesis was true. Humane paid
everything to learn it was false.
AI adds a new way to fail, because building got cheap. In a
postmortem published in August 2026 ,
solo founder Gus Chiriboga described using AI coding agents to build a tool for supervising
AI coding agents. He shipped 128 releases with CI, tests, bilingual docs and governance
policies. External users: zero. Very thorough governance for a country with no citizens.
His second attempt had 442 passing tests, and broke the moment he used it for real,
because the tests checked his own model of the system.
His diagnosis, after about two years: technical work was how he avoided real users.
“Agents make building so cheap that building becomes the procrastination.” And:
“Tests validate your model of the system. Only strangers validate the system.”
Put both on a mug.
Thinking is the scarce resource
The new loop is: build it, ship it, “learn from the data.” It sounds
responsible. It's a trap wearing a lanyard. You can't learn from data when you never wrote
down what you expected, so every result looks like a signal and nothing ever clearly
fails. A flat chart is “early.” A silent launch is “building
awareness.” Soon you're curating a graveyard of half-used features, each one
“kind of promising.”
Failing fast is cheap now. Failing legibly is still your job. Cheap building
made thinking the only scarce resource left, and the machine will happily build the wrong
thing at astonishing speed without ever asking, “Are you sure?”
Write the hypothesis in one sentence: “We believe [user] will
[behavior] because [reason].” If you can't fill in the blanks, you're not ready to
build. You're ready to think.
Define success and failure before you ship. After is when
everything looks like a win.
Pick a kill date. Put it on a calendar, and tell someone who will
hold you to it.
So write the hypothesis. Set the kill date. You have the most powerful
building tools in history, so aim them at the right thing, and learn something every
time. Every honest failure is a course correction, and course corrections are how you
reach the moon.
Now go put your work in front of a stranger.
