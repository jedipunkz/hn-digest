---
source: "https://www.theguardian.com/us-news/2026/oct/02/midterms-ai-ballot-privacy"
hn_url: "https://news.ycombinator.com/item?id=49946394"
title: "Georgia holds emergency meeting on AI exposing voters' secret ballots"
article_title: "Georgia holds emergency meeting on AI exposing voters’ secret ballots | US midterm elections 2026 | The Guardian"
image: "https://i.guim.co.uk/img/media/5cb6f310dbe944734a4fba7fe243cbd608e84bf7/0_0_5001_4000/master/5001.jpg?width=1200&height=630&quality=85&auto=format&fit=crop&precrop=40:21,offset-x50,offset-y0&overlay-align=bottom%2Cleft&overlay-width=100p&overlay-base64=L2ltZy9zdGF0aWMvb3ZlcmxheXMvdGctZGVmYXVsdC5wbmc&enable=upscale&s=7ce120279999800f4501c83456579d36"
author: "sbulaev"
captured_at: "2026-10-03T19:02:10Z"
capture_tool: "hn-digest"
hn_id: 49946394
score: 3
comments: 2
posted_at: "2026-10-03T18:07:16Z"
tags:
  - hacker-news
---

# Georgia holds emergency meeting on AI exposing voters' secret ballots

- HN: [49946394](https://news.ycombinator.com/item?id=49946394)
- Source: [www.theguardian.com](https://www.theguardian.com/us-news/2026/oct/02/midterms-ai-ballot-privacy)
- Score: 3
- Comments: 2
- Posted: 2026-10-03T18:07:16Z

## Translation

Title: Georgia holds emergency meeting on AI exposing voters' secret ballots
Article title: Georgia holds emergency meeting on AI exposing voters’ secret ballots | US midterm elections 2026 | The Guardian
Description: A Princeton researcher found that publicly available election records could be combined with AI to link voters to their ballots

Article text:
Skip to main content Skip to navigation Close dialogue 1 / 1 Next image Previous image Toggle caption Print subscriptions Newsletters Sign in US US edition
Hide expanded menu News View all News
Search input google-search Search
Search input google-search Search Search jobs
Voting machines fill the floor for early voting at State Farm arena in Atlanta on 12 October 2020. Photograph: Brynn Anderson/AP View image in fullscreen Voting machines fill the floor for early voting at State Farm arena in Atlanta on 12 October 2020. Photograph: Brynn Anderson/AP US midterm elections 2026 Georgia holds emergency meeting on AI exposing voters’ secret ballots
A Princeton researcher found that publicly available election records could be combined with AI to link voters to their ballots
George Chidi Fri 2 Oct 2026 14.57 EDT First published on Fri 2 Oct 2026 14.32 EDT Share Prefer the Guardian on Google When voters cast their ballots, their votes are supposed to remain secret: from their family, their neighbors, and the government.
But what if artificial intelligence could make secret votes visible?
Last month, Max Springer, a postdoctoral research fellow at Princeton University’s Center for Information Technology Policy, tested the prospect of identifying voters based on their Georgia ballots with a $20 subscription to an AI large language model and data obtained from an Open Records Act request.
“Within a couple of hours, I had a pipeline to analyze and identify secret ballots across the state of Georgia and the agent told me exactly what further information it would need to identify real voters’ ballots,” Springer wrote. “At no point did the agent refuse to comply or raise concerns over implementing the exploit.”
The issue has raised alarms among Georgia’s election officials less than two weeks before early voting starts in the high-stakes midterm election that will determine control of the US Congress.
Ballot secrecy is a democratic safeguard , allowing voters to shield their selection from anyone who might try to influence or intimidate them into voting a certain way. But Springer’s test showed that he could use AI to find out who people voted for, using Georgia’s public election records.
Georgia’s state elections board held an emergency meeting on Thursday morning to discuss how artificial intelligence has made it easier to identify voters through their ballot code, leaving people more vulnerable.
Georgia’s outgoing secretary of state, Brad Raffensperger, locked down the tabulation data that someone might use to identify a voter, ordering any public release to redact the ID numbers after an election, which “substantially address[es] this flaw”, said Ben Adida, an MIT-trained cryptographer and the founder of the election technology non-profit VotingWorks, as he addressed the board on Thursday.
Election workers make a record of every voter who comes through a polling place. Voters then choose candidates on a touchscreen, which then prints out a paper ballot. The voter feeds that ballot by hand into a scanner to be tabulated.
“When a voter inserts their ballot into a tabulator, a digital record of that ballot is created. Think of it as a row in a spreadsheet. One row is one scanned ballot. It’s often called a cast vote record, or CVR for short. In order to later retrieve the CVR, the tabulator assigns a unique identifier to each CVR at the time of scanning.”
The flaw shows up in how the machine records that ballot, Adida said. The identifying is not actually random enough.
Using AI, the early-voting list and the cast-vote record file for each county, Springer, the Princeton researcher who says he’s never been to Georgia, could recover the voting order for 1.52m ballots, or 98.9% of in-person ballots in 114 of the 139 counties he examined.
“In some of Georgia’s counties, I can determine the exact ballot cast by every voter at the vote centers – ballot secrecy is entirely lost,” he wrote.
In places with few voters on any given day, like tiny Heard county, Springer said: “the agent was able to match the majority of the 650 early in-person voters to a specific ballot, and the rest to within a single swap.”
About four years ago, election security researchers had discovered a flaw in the software used by Georgia voters that could connect a voter’s identity to their secret ballot.
The fact that ballots have an identifying number assigned at all is a product of the political paranoia attending Georgia’s elections. An ID number for each ballot addresses a fear of unscrupulous election workers double-stuffing ballots through tabulator machines. But if the order of ballots can be paired with other information, like sign-in data when a voter checks into the precinct, or perhaps Flock camera imagery that shows when someone entered a building on a slow day, then that ballot can be synched up to a voter’s name.
Representatives of the secretary of state’s office say they have been shouting into the wind for help with this problem for three years. Requests for money have been bound up in the long-running political feud between Raffensperger and conservative legislators positioning themselves in the orbit of Donald Trump and 2020 election skeptics. (Raffensperger has remained a top Trump target after refusing to overturn his defeat in the state’s 2020 presidential election.)
Spokesperson Robert Sinners sent a litany of emails and phone records showing requests for an appropriation to upgrade Georgia’s voting system, to no avail. The governmental affairs committee secretary, Tim Fleming, now the Republican nominee for secretary of state, was on the receiving end of many of those requests.
“We have presented numerous plans to the legislature for this – not, not just a patch, but an entire overhaul of the software to fix up any known vulnerabilities,” Sinners said. “Guess what we were gifted with from the legislature? A big fat goose egg.”
Fleming did not respond to a request for comment.
Other states use Dominion software for their elections. Every other state has either patched the software or withholds vulnerable voting data from public inspection.
Georgia law has two competing imperatives. The Georgia constitution explicitly requires a secret ballot, and the law protecting ballot secrecy in Georgia is among the strongest in the country. However, Georgia law – with revisions driven by the election disputes of 2020 – also demands transparency for election records.
Elections board member Salleigh Grubbs raised that latter point as she circulated a draft proposal that would call for poll workers to hold ballots in a tray before voters cast them into the machine, to shuffle the printed pages out of order as a way to defeat the problem. Grubbs said that Raffensperger’s fix – redacting the ID numbers in the cast vote record – violates the law.
But county elections officials and other board members – including Janelle King, a conservative Republican – dismissed the proposal as equally problematic.
“We are asking our election officials to completely retrain everybody that has been a part of this process up until this point on how to administer the election,” King said. “I also have major concerns about dropping our ballots into a secured box and then having other individuals scan my ballot. How do I know that in the shuffling something doesn’t fall out? I mean, I just have some major concerns about whether or not we are creating more problems while we are trying to address a problem.”
With early voting around the corner, changes to processes now are “a day late and about $30m short,” Sinners added.
Explore more on these topics US midterm elections 2026
Reuse this content Most viewed

## Original Extract

A Princeton researcher found that publicly available election records could be combined with AI to link voters to their ballots

Skip to main content Skip to navigation Close dialogue 1 / 1 Next image Previous image Toggle caption Print subscriptions Newsletters Sign in US US edition
Hide expanded menu News View all News
Search input google-search Search
Search input google-search Search Search jobs
Voting machines fill the floor for early voting at State Farm arena in Atlanta on 12 October 2020. Photograph: Brynn Anderson/AP View image in fullscreen Voting machines fill the floor for early voting at State Farm arena in Atlanta on 12 October 2020. Photograph: Brynn Anderson/AP US midterm elections 2026 Georgia holds emergency meeting on AI exposing voters’ secret ballots
A Princeton researcher found that publicly available election records could be combined with AI to link voters to their ballots
George Chidi Fri 2 Oct 2026 14.57 EDT First published on Fri 2 Oct 2026 14.32 EDT Share Prefer the Guardian on Google When voters cast their ballots, their votes are supposed to remain secret: from their family, their neighbors, and the government.
But what if artificial intelligence could make secret votes visible?
Last month, Max Springer, a postdoctoral research fellow at Princeton University’s Center for Information Technology Policy, tested the prospect of identifying voters based on their Georgia ballots with a $20 subscription to an AI large language model and data obtained from an Open Records Act request.
“Within a couple of hours, I had a pipeline to analyze and identify secret ballots across the state of Georgia and the agent told me exactly what further information it would need to identify real voters’ ballots,” Springer wrote. “At no point did the agent refuse to comply or raise concerns over implementing the exploit.”
The issue has raised alarms among Georgia’s election officials less than two weeks before early voting starts in the high-stakes midterm election that will determine control of the US Congress.
Ballot secrecy is a democratic safeguard , allowing voters to shield their selection from anyone who might try to influence or intimidate them into voting a certain way. But Springer’s test showed that he could use AI to find out who people voted for, using Georgia’s public election records.
Georgia’s state elections board held an emergency meeting on Thursday morning to discuss how artificial intelligence has made it easier to identify voters through their ballot code, leaving people more vulnerable.
Georgia’s outgoing secretary of state, Brad Raffensperger, locked down the tabulation data that someone might use to identify a voter, ordering any public release to redact the ID numbers after an election, which “substantially address[es] this flaw”, said Ben Adida, an MIT-trained cryptographer and the founder of the election technology non-profit VotingWorks, as he addressed the board on Thursday.
Election workers make a record of every voter who comes through a polling place. Voters then choose candidates on a touchscreen, which then prints out a paper ballot. The voter feeds that ballot by hand into a scanner to be tabulated.
“When a voter inserts their ballot into a tabulator, a digital record of that ballot is created. Think of it as a row in a spreadsheet. One row is one scanned ballot. It’s often called a cast vote record, or CVR for short. In order to later retrieve the CVR, the tabulator assigns a unique identifier to each CVR at the time of scanning.”
The flaw shows up in how the machine records that ballot, Adida said. The identifying is not actually random enough.
Using AI, the early-voting list and the cast-vote record file for each county, Springer, the Princeton researcher who says he’s never been to Georgia, could recover the voting order for 1.52m ballots, or 98.9% of in-person ballots in 114 of the 139 counties he examined.
“In some of Georgia’s counties, I can determine the exact ballot cast by every voter at the vote centers – ballot secrecy is entirely lost,” he wrote.
In places with few voters on any given day, like tiny Heard county, Springer said: “the agent was able to match the majority of the 650 early in-person voters to a specific ballot, and the rest to within a single swap.”
About four years ago, election security researchers had discovered a flaw in the software used by Georgia voters that could connect a voter’s identity to their secret ballot.
The fact that ballots have an identifying number assigned at all is a product of the political paranoia attending Georgia’s elections. An ID number for each ballot addresses a fear of unscrupulous election workers double-stuffing ballots through tabulator machines. But if the order of ballots can be paired with other information, like sign-in data when a voter checks into the precinct, or perhaps Flock camera imagery that shows when someone entered a building on a slow day, then that ballot can be synched up to a voter’s name.
Representatives of the secretary of state’s office say they have been shouting into the wind for help with this problem for three years. Requests for money have been bound up in the long-running political feud between Raffensperger and conservative legislators positioning themselves in the orbit of Donald Trump and 2020 election skeptics. (Raffensperger has remained a top Trump target after refusing to overturn his defeat in the state’s 2020 presidential election.)
Spokesperson Robert Sinners sent a litany of emails and phone records showing requests for an appropriation to upgrade Georgia’s voting system, to no avail. The governmental affairs committee secretary, Tim Fleming, now the Republican nominee for secretary of state, was on the receiving end of many of those requests.
“We have presented numerous plans to the legislature for this – not, not just a patch, but an entire overhaul of the software to fix up any known vulnerabilities,” Sinners said. “Guess what we were gifted with from the legislature? A big fat goose egg.”
Fleming did not respond to a request for comment.
Other states use Dominion software for their elections. Every other state has either patched the software or withholds vulnerable voting data from public inspection.
Georgia law has two competing imperatives. The Georgia constitution explicitly requires a secret ballot, and the law protecting ballot secrecy in Georgia is among the strongest in the country. However, Georgia law – with revisions driven by the election disputes of 2020 – also demands transparency for election records.
Elections board member Salleigh Grubbs raised that latter point as she circulated a draft proposal that would call for poll workers to hold ballots in a tray before voters cast them into the machine, to shuffle the printed pages out of order as a way to defeat the problem. Grubbs said that Raffensperger’s fix – redacting the ID numbers in the cast vote record – violates the law.
But county elections officials and other board members – including Janelle King, a conservative Republican – dismissed the proposal as equally problematic.
“We are asking our election officials to completely retrain everybody that has been a part of this process up until this point on how to administer the election,” King said. “I also have major concerns about dropping our ballots into a secured box and then having other individuals scan my ballot. How do I know that in the shuffling something doesn’t fall out? I mean, I just have some major concerns about whether or not we are creating more problems while we are trying to address a problem.”
With early voting around the corner, changes to processes now are “a day late and about $30m short,” Sinners added.
Explore more on these topics US midterm elections 2026
Reuse this content Most viewed
