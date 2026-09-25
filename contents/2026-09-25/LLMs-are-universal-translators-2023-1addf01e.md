---
source: "https://lil.law.harvard.edu/blog/2023/11/29/llms-are-universal-translators/"
hn_url: "https://news.ycombinator.com/item?id=49844062"
title: "LLMs are universal translators (2023)"
article_title: "LLMs are universal translators: on building my own translation tools for a foreign language conference | Library Innovation Lab"
image: "https://lil.law.harvard.edu/assets/images/statue-of-liberty.png"
author: "droidjj"
captured_at: "2026-09-25T13:17:36Z"
capture_tool: "hn-digest"
hn_id: 49844062
score: 1
comments: 1
posted_at: "2026-09-25T13:02:26Z"
tags:
  - hacker-news
---

# LLMs are universal translators (2023)

- HN: [49844062](https://news.ycombinator.com/item?id=49844062)
- Source: [lil.law.harvard.edu](https://lil.law.harvard.edu/blog/2023/11/29/llms-are-universal-translators/)
- Score: 1
- Comments: 1
- Posted: 2026-09-25T13:02:26Z

## Translation

Title: LLMs are universal translators (2023)
Article title: LLMs are universal translators: on building my own translation tools for a foreign language conference | Library Innovation Lab
Description: The Library Innovation Lab is growing knowledge and community by bringing library principles to technological frontiers.

Article text:
* We're hiring — Perma.cc Legal Digital Scholarship Project Manager
* We're hiring — Perma.cc Legal Digital Scholarship Project Manager
* We're hiring — Perma.cc Legal Digital Scholarship Project Manager
* We're hiring — Perma.cc Legal Digital Scholarship Project Manager
* We're hiring — Perma.cc Legal Digital Scholarship Project Manager
Home
Library Innovation Lab
Our Work
LIL on Bluesky
LIL on LinkedIn
AI Research
LLMs are universal translators: on building my own translation tools for a foreign language conference
Here is a picture of the Statue of Liberty doing a TikTok dance, as painted by van Gogh, as interpreted by ChatGPT. This is very relevant to my point and we’ll come back to it.
One of the best ways to think about large language models is as universal, personal translators. When I gave a talk at a Spanish-language library conference in Argentina recently, it was an excellent chance to test what LLMs currently offer as translators and what they might become. The answer made me optimistic for how LLMs can work as humanistic knowledge tools, in concert with library values .
This is long, so I’ve broken it up into a few sections that might be helpful to different audiences:
LLMs are universal translators . This section explains LLM embedding spaces and argues that many of LLM’s most successful applications are essentially translation tasks. I argue that LLMs are “universal translators,” not in the sense that they are perfect but in the sense that they try to translate between any input and any output.
How I built my own personal translation tools . When I spoke in Argentina, I built my own tools to translate my conference slides to Spanish and to translate other talks to English. This section gets into the weeds of what I did and how I did it. It will be most useful if you are a programmer interested in making more practical use of LLMs, or if you are interested in what might be possible for everyone as LLM tools get easier to use.
Building my own tools, part 2: real time translation . After my own talk, I watched other talks using a multimodal model to translate slides, and voice recognition and text completion APIs to translate talks.
What a universal translator means for an innovation lab . The ability to make individual, personalized translation tools changes what all of us should work on next — things that once could have been entire companies are now afternoon projects. This part considers, on the one hand, how my trip made me imagine a bunch of tools I could make and share, and on the other hand whether making and scaling tools still makes sense at all.
The cooperative principle, AI translators, and human connection . This part reflects on my experience of using technical tools to try to connect with people. I find that they highlight the “cooperative principle” — when two people communicate via an accessibility tool, they have to be more attentive, rather than less, to each other’s social signals, making me optimistic that tools can help to bring us together rather than alienate us.
LLMs are universal translators
LLMs are, in a literal sense, universal translators. They take all of their training data and embed it in a single high dimensional space, an embedding space, and then produce outputs by moving around this embedding space.
The goal of an embedding space is that similar concepts end up near each other, and different concepts end up far away. And the goal of a “large” language model is to embed everything — the space is trained using trillions of tokens representing all of the world’s digital knowledge.
A classic example to understand embedding spaces is this: we take a bunch of data and train an encoder so that if we put in similar words, they encode to a similar location in space. Words like “king” and “queen” each end up encoded as locations somewhere in the embedding space. And then, miraculously, it turns out we can do math on those locations and it makes sense. If you encode “king” into a location, and then subtract the location of “man” and add the location of “woman”, you arrive at the location of “queen.”
This is already a kind of “translation” — we’re literally moving, or translating, from the location of “king” to the location of “queen.” But we can do other kinds of translation with this same technique. We can subtract English and add Spanish, and move from “king” to “rey.” Or we can build encoders that embed pictures and sound as well as text, and encode a picture of King Arthur and come out with the word “king,” or encode the word “king” and come out with an audio file of someone saying “king.”
Embedding spaces translate from everything to everything .
Not surprisingly, a lot of the most promising applications of LLMs can be thought of as translation problems:
A programmer inputs a comment describing what a function should do in English, and it is translated to an implementation of the function in Python.
A doctor inputs an image of an x-ray, and it is translated to an English-language diagnosis.
A user inputs a text description of an image, and it is translated to an image matching the description.
A lawyer inputs a list of summaries of case holdings and facts provided by a client, and it is translated to a legal brief.
A social network inputs images uploaded by users, and they are translated to text descriptions for screenreader users.
And of course literal translation — you click “Translate text” in the Firefox browser and your computer translates it to another language.
This brings us back to the image of the Statue of Liberty doing a TikTok dance, as painted by van Gogh, that opened the article. How did the program “know” what the Statue of Liberty looks like, what dancing looks like, how van Gogh paints, or how those would all go together? It started at a random point in a high-dimensional embedding space, and then translated toward the spot that had the highest overlap of Statue-of-Liberty-ness, dance-ness, and van Gogh-ness, which it could do because it was able to encode and decode both text and images in and out of that space. It could just as easily have navigated to nearby spaces — from Statue of Liberty to Napoleon, or from van Gogh to Monet:
All of the concepts of the world are embedded in the same space and available for translation.
The idea of large language models is that we want the same model to do all of these tasks, because with human problems there’s no way of predicting what’s relevant to what. The lawyer’s brief or the programmer’s code or the Firefox translation could all require a concept map that includes Napoleon or TikTok trends for an accurate translation; large language models are willing to absorb it all and remix in any form.
That’s what I mean by “universal” translator — we don’t have to decide, up front, which facts are necessary for a successful translation, what inputs and outputs to use, because every available idea can be translated in and out of the same embedding space.
Being a universal translator doesn’t make something an accurate translator, or a social benefit. I’m not using “universal” as a superlative or saying it can do any particular translation task well. But a universal translator is a very different tool from a special-purpose translator, and it’s worth experimenting to see what it means to have one.
How I built my own personal translation tools
So, I believe that LLMs are universal translators. And I also believe, as the head of an innovation lab, that getting our hands messy is the best way to improve our intuitions about what’s coming next. So when I was invited to give a talk on disruptive innovation in libraries (adapted transcript) at the Universidad Católica de Argentina for a Spanish audience — a language I don’t speak — it was the perfect chance to experiment with what it means to have a universal translator.
To be clear, I was able to attend this Spanish-language conference not because of the tools described below, but because of the resourcefulness, patience, and enthusiasm of UCA library director Maria Soledad Lago, language professor Mercedes Rego Perlas, and the other speakers and attendees who welcomed me. Many thanks for all of their support, including with these experiments!
The scenario I decided to test was: I’m attending a conference in a foreign language, and I’m going to use low-level APIs to see if it’s possible to build my own tools to solve problems while I’m there.
My first goal was to see if it was possible to translate my slides. I knew my talk would be offered with simultaneous translation, but I wanted it to be easier to follow the text on the slides as well. That is, I wanted to show each block of text in the slides in both English and Spanish, like this:
PowerPoint already has a translator built in — you can click a text box and get a translation, like this:
I wanted to see if I could save time by automatically inserting translations for all the text boxes. I also thought I could improve on the PowerPoint feature in a couple of ways:
I could include round-trip translations in each box, English -> Spanish -> English, which would give me a way to check the translation accuracy without speaking Spanish.
I could translate entire slides at once, instead of just one text box, which would give the translation program more context to work with.
I could keep the internal formatting of the text boxes, so the same word would end up highlighted in both versions of the text.
And because the goal was to test whether universal translation can make translation tools more personal and customizable, I wanted to try to do all this in a few hours.
I started by asking ChatGPT to write a program to edit a PowerPoint deck for me:
With a little back and forth, I had a starting point — a simple program that capitalizes each word in a PowerPoint. I then started copying and pasting in code to call the OpenAI API. All I’d have to do is take the text blocks for each page, ask GPT4 to translate them to Argentine Spanish, and put the results back in. This gave me a chance to try out OpenAI’s function calling API for structured output , which I had a hunch would help with translation.
I had the fun experience at this point of having Copilot, a GPT-powered coding tool, start to recommend instructions to supply to its sibling in the translation prompt:
Here you can see that I’ve written some code myself to make a “translate” function that takes a string in English and returns Spanish, and I’m writing the instructions that will be sent off to the model. Copilot sees what I’m doing, and suggests the completion of the instruction in gray italic text — including, itself, translating English to Spanish.
The upshot was a script that edited slides to look like this:
This looks like a mess, but it’s just what I wanted! For example, here’s the text of the first block:
“The Library is to us what a laboratory is to the chemist or the physicist, and what the museum is to the naturalist.” -> “La biblioteca es para nosotros lo que el laboratorio es para el químico o el físico, y lo que el museo es para el naturalista.” -> “The library is for us what the laboratory is for the chemist or the physicist, and what the museum is for the naturalist.”
Since the round trip language looks good, I can guess that the Spanish is at least intelligible, and delete the round trip translation and move on.
Checking the round trip translations was a fascinating game, and changed how I think about machine translation. One slide I was suspicious about, for example, translated the English “patron” as “mecenas”:
I wondered if “patrons” came through correctly, or was confusing, so checked what ChatGPT thought, without tipping my hand about the word “mecenas”:
Me: what word is most common for library users in spanish? is there a word like “patrons” that denotes something distinct from commercial customers?
ChatGPT: In Spanish, the term “usuarios” is commonly used to refer to library users. “Usuarios” simply means “users.”

[truncated]

## Original Extract

The Library Innovation Lab is growing knowledge and community by bringing library principles to technological frontiers.

* We're hiring — Perma.cc Legal Digital Scholarship Project Manager
* We're hiring — Perma.cc Legal Digital Scholarship Project Manager
* We're hiring — Perma.cc Legal Digital Scholarship Project Manager
* We're hiring — Perma.cc Legal Digital Scholarship Project Manager
* We're hiring — Perma.cc Legal Digital Scholarship Project Manager
Home
Library Innovation Lab
Our Work
LIL on Bluesky
LIL on LinkedIn
AI Research
LLMs are universal translators: on building my own translation tools for a foreign language conference
Here is a picture of the Statue of Liberty doing a TikTok dance, as painted by van Gogh, as interpreted by ChatGPT. This is very relevant to my point and we’ll come back to it.
One of the best ways to think about large language models is as universal, personal translators. When I gave a talk at a Spanish-language library conference in Argentina recently, it was an excellent chance to test what LLMs currently offer as translators and what they might become. The answer made me optimistic for how LLMs can work as humanistic knowledge tools, in concert with library values .
This is long, so I’ve broken it up into a few sections that might be helpful to different audiences:
LLMs are universal translators . This section explains LLM embedding spaces and argues that many of LLM’s most successful applications are essentially translation tasks. I argue that LLMs are “universal translators,” not in the sense that they are perfect but in the sense that they try to translate between any input and any output.
How I built my own personal translation tools . When I spoke in Argentina, I built my own tools to translate my conference slides to Spanish and to translate other talks to English. This section gets into the weeds of what I did and how I did it. It will be most useful if you are a programmer interested in making more practical use of LLMs, or if you are interested in what might be possible for everyone as LLM tools get easier to use.
Building my own tools, part 2: real time translation . After my own talk, I watched other talks using a multimodal model to translate slides, and voice recognition and text completion APIs to translate talks.
What a universal translator means for an innovation lab . The ability to make individual, personalized translation tools changes what all of us should work on next — things that once could have been entire companies are now afternoon projects. This part considers, on the one hand, how my trip made me imagine a bunch of tools I could make and share, and on the other hand whether making and scaling tools still makes sense at all.
The cooperative principle, AI translators, and human connection . This part reflects on my experience of using technical tools to try to connect with people. I find that they highlight the “cooperative principle” — when two people communicate via an accessibility tool, they have to be more attentive, rather than less, to each other’s social signals, making me optimistic that tools can help to bring us together rather than alienate us.
LLMs are universal translators
LLMs are, in a literal sense, universal translators. They take all of their training data and embed it in a single high dimensional space, an embedding space, and then produce outputs by moving around this embedding space.
The goal of an embedding space is that similar concepts end up near each other, and different concepts end up far away. And the goal of a “large” language model is to embed everything — the space is trained using trillions of tokens representing all of the world’s digital knowledge.
A classic example to understand embedding spaces is this: we take a bunch of data and train an encoder so that if we put in similar words, they encode to a similar location in space. Words like “king” and “queen” each end up encoded as locations somewhere in the embedding space. And then, miraculously, it turns out we can do math on those locations and it makes sense. If you encode “king” into a location, and then subtract the location of “man” and add the location of “woman”, you arrive at the location of “queen.”
This is already a kind of “translation” — we’re literally moving, or translating, from the location of “king” to the location of “queen.” But we can do other kinds of translation with this same technique. We can subtract English and add Spanish, and move from “king” to “rey.” Or we can build encoders that embed pictures and sound as well as text, and encode a picture of King Arthur and come out with the word “king,” or encode the word “king” and come out with an audio file of someone saying “king.”
Embedding spaces translate from everything to everything .
Not surprisingly, a lot of the most promising applications of LLMs can be thought of as translation problems:
A programmer inputs a comment describing what a function should do in English, and it is translated to an implementation of the function in Python.
A doctor inputs an image of an x-ray, and it is translated to an English-language diagnosis.
A user inputs a text description of an image, and it is translated to an image matching the description.
A lawyer inputs a list of summaries of case holdings and facts provided by a client, and it is translated to a legal brief.
A social network inputs images uploaded by users, and they are translated to text descriptions for screenreader users.
And of course literal translation — you click “Translate text” in the Firefox browser and your computer translates it to another language.
This brings us back to the image of the Statue of Liberty doing a TikTok dance, as painted by van Gogh, that opened the article. How did the program “know” what the Statue of Liberty looks like, what dancing looks like, how van Gogh paints, or how those would all go together? It started at a random point in a high-dimensional embedding space, and then translated toward the spot that had the highest overlap of Statue-of-Liberty-ness, dance-ness, and van Gogh-ness, which it could do because it was able to encode and decode both text and images in and out of that space. It could just as easily have navigated to nearby spaces — from Statue of Liberty to Napoleon, or from van Gogh to Monet:
All of the concepts of the world are embedded in the same space and available for translation.
The idea of large language models is that we want the same model to do all of these tasks, because with human problems there’s no way of predicting what’s relevant to what. The lawyer’s brief or the programmer’s code or the Firefox translation could all require a concept map that includes Napoleon or TikTok trends for an accurate translation; large language models are willing to absorb it all and remix in any form.
That’s what I mean by “universal” translator — we don’t have to decide, up front, which facts are necessary for a successful translation, what inputs and outputs to use, because every available idea can be translated in and out of the same embedding space.
Being a universal translator doesn’t make something an accurate translator, or a social benefit. I’m not using “universal” as a superlative or saying it can do any particular translation task well. But a universal translator is a very different tool from a special-purpose translator, and it’s worth experimenting to see what it means to have one.
How I built my own personal translation tools
So, I believe that LLMs are universal translators. And I also believe, as the head of an innovation lab, that getting our hands messy is the best way to improve our intuitions about what’s coming next. So when I was invited to give a talk on disruptive innovation in libraries (adapted transcript) at the Universidad Católica de Argentina for a Spanish audience — a language I don’t speak — it was the perfect chance to experiment with what it means to have a universal translator.
To be clear, I was able to attend this Spanish-language conference not because of the tools described below, but because of the resourcefulness, patience, and enthusiasm of UCA library director Maria Soledad Lago, language professor Mercedes Rego Perlas, and the other speakers and attendees who welcomed me. Many thanks for all of their support, including with these experiments!
The scenario I decided to test was: I’m attending a conference in a foreign language, and I’m going to use low-level APIs to see if it’s possible to build my own tools to solve problems while I’m there.
My first goal was to see if it was possible to translate my slides. I knew my talk would be offered with simultaneous translation, but I wanted it to be easier to follow the text on the slides as well. That is, I wanted to show each block of text in the slides in both English and Spanish, like this:
PowerPoint already has a translator built in — you can click a text box and get a translation, like this:
I wanted to see if I could save time by automatically inserting translations for all the text boxes. I also thought I could improve on the PowerPoint feature in a couple of ways:
I could include round-trip translations in each box, English -> Spanish -> English, which would give me a way to check the translation accuracy without speaking Spanish.
I could translate entire slides at once, instead of just one text box, which would give the translation program more context to work with.
I could keep the internal formatting of the text boxes, so the same word would end up highlighted in both versions of the text.
And because the goal was to test whether universal translation can make translation tools more personal and customizable, I wanted to try to do all this in a few hours.
I started by asking ChatGPT to write a program to edit a PowerPoint deck for me:
With a little back and forth, I had a starting point — a simple program that capitalizes each word in a PowerPoint. I then started copying and pasting in code to call the OpenAI API. All I’d have to do is take the text blocks for each page, ask GPT4 to translate them to Argentine Spanish, and put the results back in. This gave me a chance to try out OpenAI’s function calling API for structured output , which I had a hunch would help with translation.
I had the fun experience at this point of having Copilot, a GPT-powered coding tool, start to recommend instructions to supply to its sibling in the translation prompt:
Here you can see that I’ve written some code myself to make a “translate” function that takes a string in English and returns Spanish, and I’m writing the instructions that will be sent off to the model. Copilot sees what I’m doing, and suggests the completion of the instruction in gray italic text — including, itself, translating English to Spanish.
The upshot was a script that edited slides to look like this:
This looks like a mess, but it’s just what I wanted! For example, here’s the text of the first block:
“The Library is to us what a laboratory is to the chemist or the physicist, and what the museum is to the naturalist.” -> “La biblioteca es para nosotros lo que el laboratorio es para el químico o el físico, y lo que el museo es para el naturalista.” -> “The library is for us what the laboratory is for the chemist or the physicist, and what the museum is for the naturalist.”
Since the round trip language looks good, I can guess that the Spanish is at least intelligible, and delete the round trip translation and move on.
Checking the round trip translations was a fascinating game, and changed how I think about machine translation. One slide I was suspicious about, for example, translated the English “patron” as “mecenas”:
I wondered if “patrons” came through correctly, or was confusing, so checked what ChatGPT thought, without tipping my hand about the word “mecenas”:
Me: what word is most common for library users in spanish? is there a word like “patrons” that denotes something distinct from commercial customers?
ChatGPT: In Spanish, the term “usuarios” is commonly used to refer to library users. “Usuarios” simply means “users.”

[truncated]
