---
source: "https://www.deldiosglasshouse.com/stuff-weve-made/ai-g-code"
hn_url: "https://news.ycombinator.com/item?id=50046050"
title: "AI-Powered Plasma Cutting"
article_title: "Glass House Arts - AI G-Code"
image: "https://lh7-us.googleusercontent.com/sitesv-images-rt/AMxu72tpj_gQ6o65tKy8fNWzZIjvy40520XR2ZTu0PMirgmLz5dUMnCOQhI5W_H7bA_alwRIoK8wO6SI2QhCWBoipzyH7jb_THunJjgceQJtfv9yxMOKWF9YBRqjzYqznf8PNbKYpxcyHJTygLCwAd_Pp5LPLCZfcBJ2KCp-aW0oblQf_ig61GkUSl_Z0agmG0hDISK-6mdntlTejureAtII3ibs6jwtYBJKTDp7I6BglYQ=w1280"
author: "cjemmott"
captured_at: "2026-10-11T18:38:31Z"
capture_tool: "hn-digest"
hn_id: 50046050
score: 1
comments: 1
posted_at: "2026-10-11T18:11:19Z"
tags:
  - hacker-news
---

# AI-Powered Plasma Cutting

- HN: [50046050](https://news.ycombinator.com/item?id=50046050)
- Source: [www.deldiosglasshouse.com](https://www.deldiosglasshouse.com/stuff-weve-made/ai-g-code)
- Score: 1
- Comments: 1
- Posted: 2026-10-11T18:11:19Z

## Translation

Title: AI-Powered Plasma Cutting
Article title: Glass House Arts - AI G-Code
Description: So today (April, 2023) I had this idea - maybe I should let AI (GPT-4) control my CNC plasma cutter!

Article text:
Glass House Arts - AI G-Code Search this site Embedded Files Skip to main content Skip to navigation Glass House Arts Home
Stuff We've Made Our Own Devices
Temple of Floating Compression
Stuff We've Made Our Own Devices
Temple of Floating Compression
Stuff We've Made Our Own Devices
Temple of Floating Compression
So today (April, 2023) I had this idea - maybe I should let AI (GPT- 4 ) control my CNC plasma cutter!
To back up, a CNC plasma cutter is an industrial robot that cuts sheets of steel like they are butter. The AI is OpenAI's GPT-4, which is not specifically trained to do this - it is just a large language model that might have seen some examples on the internet. The AI is also known to be confidently wrong and hallucinate outputs.
No harm in giving control of an industrial robot to a language model, right? I mean, that gif above shows what it does to 10 Gauge steel.
The typical workflow is to design your part in drafting software (CAD) and then use different software (CAM) to make the tool paths (G-Code). The plasma cutter then just follows the instructions in that G-Code.
So I asked GPT to skip all that and directly make some G-Code for a sign:
Thankfully GPT has some guard rails, and was unwilling to entertain this request.
But reading between the lines, it seemed like the objection wasn't that the request was insane, but rather that the sign was too detailed. After rewording my requests a few times, I got it to generate some G-Code:
Probably the right next step would be to do some validation work, but I ran up to the shop and started setting up the machine! Surprisingly, the code loaded on the first try:
If the code is syntactically valid then it must be safe, right? Let's hit go!
It was liberal with the interpretation of "GPT", but I kind of love it.
One reason it is off is that GPT-4 has no vision system - it doesn't know the shape of things - and it works entirely by predicting what word is most likely to follow other words. It also did have a few technical errors - not setting the initial height at each pierce caused the "T" to be kinda messed up, and there was a weird 30 second pause in the middle of the cut path for no reason. But nobody got hurt, so I am going to call it a win!
Prompt handed to Claude Opus 4.5 Extended:
Happy Friday to you! Attached are two files. The PDF is an experiment I ran about three years ago using an LLM to directly write G-code for my plasma cutter. It went ok - you can see in the output that it really struggled with the shapes and even failed to get height control on the T, but it did write valid G-code.
We have come a LONG way in the last three years! Second file is a simple example shape I cut last week, 16 Ga steel. Hypertherm 45XP on a Langmuir Crossfire Pro table. Can you make a sign that shows off how much better LLMs are now? I feel like just writing Claude is too easy - maybe have some fun with it?
The sign has some issues (the chamfered corners are inverted), but is clearly legible - a big improvement over the past three years. I especially like that the sign is a callback to my original sign proposal from 2023: "This is a terrible idea"...
This time I just sent Opus 5.5 Max the URL for this page with the prompt:
Ok, time to show off. Update for: https://www.deldiosglasshouse.com/stuff-weve-made/ai-g-code
Folks argue if LLMs are improving, but in my mind the huge jumps were reasoning models back at the end of 2025, and whatever happened in March 2026 when they got better at coding than me. I feel like this sign captures where we are - it included tool use (writing and running python code to make and debug the g-code), and made sure to include stenciled letters complete with lead ins. It parsed the sign, made a good joke, and even looked up they specs (web search) for FireControl and my Hypertherm 45XP, because this time I didn't give it any example code.
Aside: the letter shift in "terrible" was my plasma cutter bumping the material - the kerning is great in the file.

## Original Extract

So today (April, 2023) I had this idea - maybe I should let AI (GPT-4) control my CNC plasma cutter!

Glass House Arts - AI G-Code Search this site Embedded Files Skip to main content Skip to navigation Glass House Arts Home
Stuff We've Made Our Own Devices
Temple of Floating Compression
Stuff We've Made Our Own Devices
Temple of Floating Compression
Stuff We've Made Our Own Devices
Temple of Floating Compression
So today (April, 2023) I had this idea - maybe I should let AI (GPT- 4 ) control my CNC plasma cutter!
To back up, a CNC plasma cutter is an industrial robot that cuts sheets of steel like they are butter. The AI is OpenAI's GPT-4, which is not specifically trained to do this - it is just a large language model that might have seen some examples on the internet. The AI is also known to be confidently wrong and hallucinate outputs.
No harm in giving control of an industrial robot to a language model, right? I mean, that gif above shows what it does to 10 Gauge steel.
The typical workflow is to design your part in drafting software (CAD) and then use different software (CAM) to make the tool paths (G-Code). The plasma cutter then just follows the instructions in that G-Code.
So I asked GPT to skip all that and directly make some G-Code for a sign:
Thankfully GPT has some guard rails, and was unwilling to entertain this request.
But reading between the lines, it seemed like the objection wasn't that the request was insane, but rather that the sign was too detailed. After rewording my requests a few times, I got it to generate some G-Code:
Probably the right next step would be to do some validation work, but I ran up to the shop and started setting up the machine! Surprisingly, the code loaded on the first try:
If the code is syntactically valid then it must be safe, right? Let's hit go!
It was liberal with the interpretation of "GPT", but I kind of love it.
One reason it is off is that GPT-4 has no vision system - it doesn't know the shape of things - and it works entirely by predicting what word is most likely to follow other words. It also did have a few technical errors - not setting the initial height at each pierce caused the "T" to be kinda messed up, and there was a weird 30 second pause in the middle of the cut path for no reason. But nobody got hurt, so I am going to call it a win!
Prompt handed to Claude Opus 4.5 Extended:
Happy Friday to you! Attached are two files. The PDF is an experiment I ran about three years ago using an LLM to directly write G-code for my plasma cutter. It went ok - you can see in the output that it really struggled with the shapes and even failed to get height control on the T, but it did write valid G-code.
We have come a LONG way in the last three years! Second file is a simple example shape I cut last week, 16 Ga steel. Hypertherm 45XP on a Langmuir Crossfire Pro table. Can you make a sign that shows off how much better LLMs are now? I feel like just writing Claude is too easy - maybe have some fun with it?
The sign has some issues (the chamfered corners are inverted), but is clearly legible - a big improvement over the past three years. I especially like that the sign is a callback to my original sign proposal from 2023: "This is a terrible idea"...
This time I just sent Opus 5.5 Max the URL for this page with the prompt:
Ok, time to show off. Update for: https://www.deldiosglasshouse.com/stuff-weve-made/ai-g-code
Folks argue if LLMs are improving, but in my mind the huge jumps were reasoning models back at the end of 2025, and whatever happened in March 2026 when they got better at coding than me. I feel like this sign captures where we are - it included tool use (writing and running python code to make and debug the g-code), and made sure to include stenciled letters complete with lead ins. It parsed the sign, made a good joke, and even looked up they specs (web search) for FireControl and my Hypertherm 45XP, because this time I didn't give it any example code.
Aside: the letter shift in "terrible" was my plasma cutter bumping the material - the kerning is great in the file.
