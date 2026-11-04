---
source: "https://gizmodo.com/did-claude-help-an-amateur-astronomer-discover-a-new-planet-2000824393"
hn_url: "https://news.ycombinator.com/item?id=50362860"
title: "Did Claude Help an Amateur Astronomer Discover a New Planet?"
article_title: "Did Claude Help an Amateur Astronomer Discover a New Planet?"
image: "https://gizmodo.com/app/uploads/2026/10/claude-new-planet-1200x675.jpg"
author: "nvmdbljstm"
captured_at: "2026-11-04T08:43:26Z"
capture_tool: "hn-digest"
hn_id: 50362860
score: 4
comments: 1
posted_at: "2026-11-04T08:37:59Z"
tags:
  - hacker-news
---

# Did Claude Help an Amateur Astronomer Discover a New Planet?

- HN: [50362860](https://news.ycombinator.com/item?id=50362860)
- Source: [gizmodo.com](https://gizmodo.com/did-claude-help-an-amateur-astronomer-discover-a-new-planet-2000824393)
- Score: 4
- Comments: 1
- Posted: 2026-11-04T08:37:59Z

## Translation

Title: Did Claude Help an Amateur Astronomer Discover a New Planet?
Description: Billions of exoplanets are believed to be out there awaiting discovery. AI—with its ability to sift through mountains of data—could speed that process along.

Article text:
Did Claude Help an Amateur Astronomer Discover a New Planet?
Skip to content
Gizmodo
Latest
Did Claude Help an Amateur Astronomer Discover a New Planet?
Last month, Pavel Rabtsevich—a 28-year-old product manager living in Spain—was hunting for planets. He didn’t have any formal training in astronomy or even a telescope. What he did have was an internet connection and subscriptions to multiple AI tools.
He’d been combing through public astronomical data captured by the Transiting Exoplanet Survey Satellite (TESS), a space telescope operated by MIT astronomers and launched by NASA in 2018. As its name suggests, TESS’ main purpose is to help astronomers find exoplanets (planets outside our home solar system) by measuring changes in stellar luminosity; if a star’s brightness dips at regular intervals, it could mean there’s a planet orbiting it and temporarily blocking some of its light from reaching our telescopes.
The data Rabtsevich was analyzing was taken from a narrow sliver of TESS’s full range, captured from November of last year to early January. He handed the dataset—comprising measurements of over 126,000 stars, according to Rabtsevich—to TypeSafe’s Jev, Anthropic’s Claude Code, and OpenAI’s Codex, prompting the AI systems to look for patterns of dimming that could indicate exoplanets, and iteratively narrowing down a list of candidates. “I didn’t take the models’ first answers,” he told Gizmodo. “I pushed them, round after round, to improve their own methods, and checked what they gave me against other models. “I decided which tests to run, judged whether a result was strong enough, and chose what to keep.”
He eventually homed in on TIC 4206066, a star a little smaller than our own Sun and located roughly 116 light-years from Earth. The measurements were unmistakable: an ever-so-slight dip in the star’s brightness every 3.18 Earth days.
Intrigued but unconvinced, he went deeper into the TESS data, pulling the star’s measurements from 2020 and 2018. Sure enough, the same dips in the star’s brightness appeared in the datasets for both years, at the same frequency.
All the evidence was pointing to the possible presence of an unidentified exoplanet, which he calculated would be just a little over 1.4 times bigger than Earth. With its full solar year lasting just a little over three Earth days, Rabtsevich predicts its surface temperature would be around 1,000 degrees Fahrenheit, significantly hotter than that of Mercury—the planet nearest our Sun—and hot enough to melt lead. (His AI-assisted digging also revealed another dimming pattern for the same star, once every 11.13 Earth days, which could be caused by a second previously unknown exoplanet, but Rabtsevich considers the evidence for that one less convincing than the other.)
He ran a total of seventy-four tests with Claude Code, trying to disprove his hypothesis. In one, he gave the agent two of the TESS datasets and asked it to predict when the dimming would occur in the third. Its predictions were correct in all three cases.
Rabtsevich submitted that report to MIT, requesting that TESS be positioned to observe TIC-4206066 and its possible exoplanet the next time the telescope is scheduled to focus on that sector of space. It was approved. From October 31 to November 26, the telescope will measure the star’s brightness every two minutes. Earlier this week, Rabtsevich publicly posted the times when the star should dim during that observation period, if there is indeed an exoplanet making regular transits across it.
“If the star dims on time in November, this gets a lot more serious,” he wrote in an X post outlining his findings . “If it doesn’t, I’ll post that too.”
The discovery of exoplanets is nothing new: astronomers have already cataloged more than 6,400 . Some (like our own) orbit a star, while other so-called “rogue planets” drift aimlessly through the cosmos untethered to any gravitational center, honoring the Greek root of the word “planet,” which means “to wander.” Billions of exoplanets are believed to still be out there awaiting discovery. In its eight years of scanning the night sky, TESS has already identified more than 1,000 exoplanet candidates that were subsequently confirmed by astronomers in the lab.
Observing subtle changes in luminosity across such vast cosmic distances leaves plenty of room for error. Binary star systems, for example, can be easily mistaken for an exoplanet passing in front of a star. “The biggest challenge with confirming planets is that we have to rule out false-positive scenarios,” says Kevin Hardegree-Ullman, a research scientist at the NASA Exoplanet Science Institute (NExScI). “This is difficult because: (1) there are just so many candidates… (2) we have limited telescopes… and we are generally competing with all other astronomers for time on these… and (3) some targets are not amenable to follow-up measurements with current telescopes (e.g., they are too faint to get a good signal other than a transit-like event).”
AI, with its superhuman ability to parse through galactic quantities of data, could be a gamechanger—at least when it comes to detecting exoplanet candidates. “Without these tools I couldn’t have processed a data set this size and carried it through to a result,” said Rabtsevich. “Even a couple of years ago my job meant building endless spreadsheets and grinding through data sets, and that ate up a huge amount of time. Here, a search I planned went through 126,000 stars, and then dozens of tests ran on a single one, in about two weeks.”
An AI model NASA deployed a little under a year ago, which was designed to comb through data from the TESS telescope, has already identified around 7,000 exoplanets. And since tools like Claude Code and Codex are publicly available, they could also democratize the discovery process beyond the scientific establishment and among amateur astronomers like Rabtsevich. “AI tools won’t necessarily be helpful in the process of confirming planets,” says Kevin Hardegree-Ullman, “but [they] will likely continue to be used to string together publicly available code to help search for planets and run them through the basic vetting processes… ”
Rabtsevich is careful to emphasize that correlation does not necessarily equal causation: the pattern he detected in his AI-assisted investigation may turn out to be rooted in something other than an exoplanet. Regardless of what TESS observations of TIC 4206066 yield starting on the 31st, he’s been happy to see that his findings have catalyzed some interest among fellow amateur astronomers. “Since my [X] post, I’ve seen a lot of people start searching TESS data for planets themselves,” he said. “I hope that enthusiasm comes with careful checks, so professional astronomers aren’t left sorting through poorly vetted signals.”
Others are just relieved to finally be seeing a conversation about AI that’s unrelated to doom, gloom, or slop: “There is an awful lot of absolute shite made with the use of AI,” a Redditor wrote in response to Rabtsevich’s post about his findings in a Claude subreddit. “But this… this is fucking incredible. Well done.”
Subscribe and interact with our community, get up to date with our customised Newsletters and much more.
Lawmakers Can’t Ban VPNs No Matter How Hard They Try
Early Results Suggest 473 Million Dollars’ Worth of Buyer’s Remorse for Elon Musk
Apple Begins Testing Its Controversial Always-Listening AI Feature
Flock Cameras Vandalized in Texas With Message: ‘Release the Epstein Files’
China Details Its Plan to Land Astronauts on the Moon Before 2030
Winter Hasn’t Even Started, and El Niño Is Already Drenching the US From Coast to Coast
Website Selling ‘AI Slop Audiobooks’ Shuts Down Following Fierce Backlash From Authors
UK Finally Removes Masts From Sunken WWII Bomb Ship—and Leaves the Explosives
Proton VPN Just Made Black Friday Easy: One Plan, Every Feature, 75% Off
iPad mini OLED Review: Apple’s Smallest Tablet Is Also Its Best
HP EliteBook X Flip G2i 14 AI Review: Great for the Office, if Not Your Wallet
The Best Gadgets of October 2026
CMF Speaker M Review: This Bluetooth Speaker Sparks Joy
The Best Retro Gaming Gifts Under $100
Lenovo Yoga 9n 2-in-1 Review: More Pen, Less Performance
Jackery HomePower 1000 Plus v2 Review: Pint-Sized Power Pack
The Best Tech Gifts Under $100
Lawmakers Can’t Ban VPNs No Matter How Hard They Try
Governments can make it harder to get around internet restrictions, but eliminating the workarounds is another matter entirely.
China Details Its Plan to Land Astronauts on the Moon Before 2030
The race to the lunar surface is officially on, and China is making sure it stays in the running.
Website Selling ‘AI Slop Audiobooks’ Shuts Down Following Fierce Backlash From Authors
As one Redditor put it: “i will never understand what part of this seemed like a good idea.”
Apple to Place Stricter Age Restrictions on Dating Apps and Some AI Chatbots
Most AI companion chatbots can get an automatic minimum age rating of 18+.
OpenAI Hires DOGE Alumni Who Hacked Into Anthropic to Strengthen Security
The company also said it was also “more than quadrupling our rewards for critical findings” through its bug bounty program.
Meta Removed Disturbing AI Videos After Initially Saying They Didn’t Violate Its Rules
Meta recently removed several accounts sharing AI-generated videos of coaches inappropriately touching and kissing young athletes.
Search
Sign up for our newsletters
Latest
Tech News
Artificial Intelligence
Home Entertainment & Smart Home

## Original Extract

Billions of exoplanets are believed to be out there awaiting discovery. AI—with its ability to sift through mountains of data—could speed that process along.

Did Claude Help an Amateur Astronomer Discover a New Planet?
Skip to content
Gizmodo
Latest
Did Claude Help an Amateur Astronomer Discover a New Planet?
Last month, Pavel Rabtsevich—a 28-year-old product manager living in Spain—was hunting for planets. He didn’t have any formal training in astronomy or even a telescope. What he did have was an internet connection and subscriptions to multiple AI tools.
He’d been combing through public astronomical data captured by the Transiting Exoplanet Survey Satellite (TESS), a space telescope operated by MIT astronomers and launched by NASA in 2018. As its name suggests, TESS’ main purpose is to help astronomers find exoplanets (planets outside our home solar system) by measuring changes in stellar luminosity; if a star’s brightness dips at regular intervals, it could mean there’s a planet orbiting it and temporarily blocking some of its light from reaching our telescopes.
The data Rabtsevich was analyzing was taken from a narrow sliver of TESS’s full range, captured from November of last year to early January. He handed the dataset—comprising measurements of over 126,000 stars, according to Rabtsevich—to TypeSafe’s Jev, Anthropic’s Claude Code, and OpenAI’s Codex, prompting the AI systems to look for patterns of dimming that could indicate exoplanets, and iteratively narrowing down a list of candidates. “I didn’t take the models’ first answers,” he told Gizmodo. “I pushed them, round after round, to improve their own methods, and checked what they gave me against other models. “I decided which tests to run, judged whether a result was strong enough, and chose what to keep.”
He eventually homed in on TIC 4206066, a star a little smaller than our own Sun and located roughly 116 light-years from Earth. The measurements were unmistakable: an ever-so-slight dip in the star’s brightness every 3.18 Earth days.
Intrigued but unconvinced, he went deeper into the TESS data, pulling the star’s measurements from 2020 and 2018. Sure enough, the same dips in the star’s brightness appeared in the datasets for both years, at the same frequency.
All the evidence was pointing to the possible presence of an unidentified exoplanet, which he calculated would be just a little over 1.4 times bigger than Earth. With its full solar year lasting just a little over three Earth days, Rabtsevich predicts its surface temperature would be around 1,000 degrees Fahrenheit, significantly hotter than that of Mercury—the planet nearest our Sun—and hot enough to melt lead. (His AI-assisted digging also revealed another dimming pattern for the same star, once every 11.13 Earth days, which could be caused by a second previously unknown exoplanet, but Rabtsevich considers the evidence for that one less convincing than the other.)
He ran a total of seventy-four tests with Claude Code, trying to disprove his hypothesis. In one, he gave the agent two of the TESS datasets and asked it to predict when the dimming would occur in the third. Its predictions were correct in all three cases.
Rabtsevich submitted that report to MIT, requesting that TESS be positioned to observe TIC-4206066 and its possible exoplanet the next time the telescope is scheduled to focus on that sector of space. It was approved. From October 31 to November 26, the telescope will measure the star’s brightness every two minutes. Earlier this week, Rabtsevich publicly posted the times when the star should dim during that observation period, if there is indeed an exoplanet making regular transits across it.
“If the star dims on time in November, this gets a lot more serious,” he wrote in an X post outlining his findings . “If it doesn’t, I’ll post that too.”
The discovery of exoplanets is nothing new: astronomers have already cataloged more than 6,400 . Some (like our own) orbit a star, while other so-called “rogue planets” drift aimlessly through the cosmos untethered to any gravitational center, honoring the Greek root of the word “planet,” which means “to wander.” Billions of exoplanets are believed to still be out there awaiting discovery. In its eight years of scanning the night sky, TESS has already identified more than 1,000 exoplanet candidates that were subsequently confirmed by astronomers in the lab.
Observing subtle changes in luminosity across such vast cosmic distances leaves plenty of room for error. Binary star systems, for example, can be easily mistaken for an exoplanet passing in front of a star. “The biggest challenge with confirming planets is that we have to rule out false-positive scenarios,” says Kevin Hardegree-Ullman, a research scientist at the NASA Exoplanet Science Institute (NExScI). “This is difficult because: (1) there are just so many candidates… (2) we have limited telescopes… and we are generally competing with all other astronomers for time on these… and (3) some targets are not amenable to follow-up measurements with current telescopes (e.g., they are too faint to get a good signal other than a transit-like event).”
AI, with its superhuman ability to parse through galactic quantities of data, could be a gamechanger—at least when it comes to detecting exoplanet candidates. “Without these tools I couldn’t have processed a data set this size and carried it through to a result,” said Rabtsevich. “Even a couple of years ago my job meant building endless spreadsheets and grinding through data sets, and that ate up a huge amount of time. Here, a search I planned went through 126,000 stars, and then dozens of tests ran on a single one, in about two weeks.”
An AI model NASA deployed a little under a year ago, which was designed to comb through data from the TESS telescope, has already identified around 7,000 exoplanets. And since tools like Claude Code and Codex are publicly available, they could also democratize the discovery process beyond the scientific establishment and among amateur astronomers like Rabtsevich. “AI tools won’t necessarily be helpful in the process of confirming planets,” says Kevin Hardegree-Ullman, “but [they] will likely continue to be used to string together publicly available code to help search for planets and run them through the basic vetting processes… ”
Rabtsevich is careful to emphasize that correlation does not necessarily equal causation: the pattern he detected in his AI-assisted investigation may turn out to be rooted in something other than an exoplanet. Regardless of what TESS observations of TIC 4206066 yield starting on the 31st, he’s been happy to see that his findings have catalyzed some interest among fellow amateur astronomers. “Since my [X] post, I’ve seen a lot of people start searching TESS data for planets themselves,” he said. “I hope that enthusiasm comes with careful checks, so professional astronomers aren’t left sorting through poorly vetted signals.”
Others are just relieved to finally be seeing a conversation about AI that’s unrelated to doom, gloom, or slop: “There is an awful lot of absolute shite made with the use of AI,” a Redditor wrote in response to Rabtsevich’s post about his findings in a Claude subreddit. “But this… this is fucking incredible. Well done.”
Subscribe and interact with our community, get up to date with our customised Newsletters and much more.
Lawmakers Can’t Ban VPNs No Matter How Hard They Try
Early Results Suggest 473 Million Dollars’ Worth of Buyer’s Remorse for Elon Musk
Apple Begins Testing Its Controversial Always-Listening AI Feature
Flock Cameras Vandalized in Texas With Message: ‘Release the Epstein Files’
China Details Its Plan to Land Astronauts on the Moon Before 2030
Winter Hasn’t Even Started, and El Niño Is Already Drenching the US From Coast to Coast
Website Selling ‘AI Slop Audiobooks’ Shuts Down Following Fierce Backlash From Authors
UK Finally Removes Masts From Sunken WWII Bomb Ship—and Leaves the Explosives
Proton VPN Just Made Black Friday Easy: One Plan, Every Feature, 75% Off
iPad mini OLED Review: Apple’s Smallest Tablet Is Also Its Best
HP EliteBook X Flip G2i 14 AI Review: Great for the Office, if Not Your Wallet
The Best Gadgets of October 2026
CMF Speaker M Review: This Bluetooth Speaker Sparks Joy
The Best Retro Gaming Gifts Under $100
Lenovo Yoga 9n 2-in-1 Review: More Pen, Less Performance
Jackery HomePower 1000 Plus v2 Review: Pint-Sized Power Pack
The Best Tech Gifts Under $100
Lawmakers Can’t Ban VPNs No Matter How Hard They Try
Governments can make it harder to get around internet restrictions, but eliminating the workarounds is another matter entirely.
China Details Its Plan to Land Astronauts on the Moon Before 2030
The race to the lunar surface is officially on, and China is making sure it stays in the running.
Website Selling ‘AI Slop Audiobooks’ Shuts Down Following Fierce Backlash From Authors
As one Redditor put it: “i will never understand what part of this seemed like a good idea.”
Apple to Place Stricter Age Restrictions on Dating Apps and Some AI Chatbots
Most AI companion chatbots can get an automatic minimum age rating of 18+.
OpenAI Hires DOGE Alumni Who Hacked Into Anthropic to Strengthen Security
The company also said it was also “more than quadrupling our rewards for critical findings” through its bug bounty program.
Meta Removed Disturbing AI Videos After Initially Saying They Didn’t Violate Its Rules
Meta recently removed several accounts sharing AI-generated videos of coaches inappropriately touching and kissing young athletes.
Search
Sign up for our newsletters
Latest
Tech News
Artificial Intelligence
Home Entertainment & Smart Home
