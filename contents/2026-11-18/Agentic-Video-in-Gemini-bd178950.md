---
source: "https://blog.google/innovation-and-ai/models-and-research/gemini-models/introducing-agentic-video-in-gemini/"
hn_url: "https://news.ycombinator.com/item?id=50547290"
title: "Agentic Video in Gemini"
article_title: "Introducing Agentic Video in Gemini"
image: "https://storage.googleapis.com/gweb-uniblog-publish-prod/images/agentic-video___keyword__blog-header.width-1300.png"
author: "gmays"
captured_at: "2026-11-18T16:14:23Z"
capture_tool: "hn-digest"
hn_id: 50547290
score: 2
comments: 0
posted_at: "2026-11-18T15:13:45Z"
tags:
  - hacker-news
---

# Agentic Video in Gemini

- HN: [50547290](https://news.ycombinator.com/item?id=50547290)
- Source: [blog.google](https://blog.google/innovation-and-ai/models-and-research/gemini-models/introducing-agentic-video-in-gemini/)
- Score: 2
- Comments: 0
- Posted: 2026-11-18T15:13:45Z

## Translation

Title: Agentic Video in Gemini
Article title: Introducing Agentic Video in Gemini
Description: We’re launching agentic video understanding across our latest Gemini models for improved accuracy and lower costs and token usage.

Article text:
Skip to main content
Introducing agentic video understanding with Gemini
Models & Research
Google DeepMind
Infrastructure & cloud
Global network
Outreach & initiatives
Creating opportunity
Inside Google
Around the globe
Innovation & AI
Innovation & AI
Products & platforms
Products & platforms
Global (English)
Africa (English)
Australia (English)
Brasil (Português)
Canada (English)
Canada (Français)
Česko (Čeština)
Deutschland (Deutsch)
España (Español)
France (Français)
Greece (Ελληνικά)
India (English)
Indonesia (Bahasa Indonesia)
Ireland (English)
Italia (Italiano)
日本 (日本語)
대한민국 (한국어)
Latinoamérica (Español)
Malaysia (English)
الشرق الأوسط وشمال أفريقيا (اللغة العربية)
MENA (English)
Nederlands (Nederland)
New Zealand (English)
Polska (Polski)
Portugal (Português)
România (Română)
Sverige (Svenska)
ประเทศไทย (ไทย)
Türkiye (Türkçe)
台灣 (中文)
Introducing agentic video understanding with Gemini
Our new agentic feature for video analysis cuts token consumption by up to 88%, reduces costs by up to 66%, and boosts quality by up to 7%.
Senior Product Manager, Google DeepMind
Research Director, Google DeepMind
Your browser does not support the audio element.
Google just launched agentic video understanding for Gemini 3.7 Flash, 3.6 Flash, and 3.5 Flash-Lite. This feature allows the model to dynamically scan video segments, which improves accuracy while cutting token usage by up to 88% and costs by up to 66%. You can start using this today by setting your API configuration to "agentic" in Google AI Studio or the Gemini Enterprise Agent Platform.
Today, we’re launching agentic video understanding across our latest models: Gemini 3.7 Flash, 3.6 Flash and 3.5 Flash-Lite. This new capability improves accuracy while dramatically reducing token usage and costs for video analysis. Similar to agentic vision , which combines code execution with Gemini models’ native image understanding, agentic video understanding uses Gemini’s native video tools to improve performance and unlock new capabilities for video processing like sub-second moment retrieval, more accurate anomaly detection, precise counting and more.
The feature is available today for video uploads and YouTube videos via the Gemini API in Google AI Studio and the Gemini Enterprise Agent Platform.
Unlike current ‘static’ processing, where the model ingests the video at a fixed frames-per-second rate (default 1 FPS, adjustable via API), agentic video understanding pairs the model’s core reasoning with native video tools to dynamically search, scan, and inspect target video segments across visual frames, audio, and transcripts. Across standard video analysis benchmarks, Gemini models with agentic video understanding reduce analysis costs by up to 66% and token consumption by up to 88%, while improving accuracy by up to 7%.
These efficiency gains are especially pronounced on long-form video (from 10-minute how-to guides to 90-minute lectures and multi-hour recordings), where static processing forces developers to choose between high token costs or techniques that drop critical details.
Activating agentic video understanding drops token consumption by up to 88% and boosts accuracy by up to 7% with Gemini 3.7 Flash.
While these gains span all three supported models, Gemini 3.7 Flash with agentic understanding offers the best possible quality overall and the best combination of quality and cost efficiency, putting it at the accuracy-to-cost pareto frontier among tested models for video understanding.
Using agentic video understanding places Gemini 3.7 Flash at the accuracy-to-cost pareto frontier for video analysis.
Instead of static processing where the model ingests media streams at a fixed frame rate, agentic video understanding enables Gemini to take an active, goal-directed role in determining what to watch, at what speed, and through which modality (frames, audio, or transcript), fetching only the moments and signals needed. While developers could previously do this manually, with agentic video understanding, Gemini can accomplish it through an agentic loop, invoking an internal tool to load the relevant part of the video file, significantly reducing development overheads.
Agentic video understanding transforms how developers can process long-form video content across a variety of demanding applications.
Sub-second moment retrieval : Pinpoint split-second state changes and tight cut boundaries that are easily missed at 1 FPS, making precise automated video editing possible.
Long-form needle-in-a-haystack search : Answer complex queries across multi-hour videos without consuming millions of tokens.
Anomaly detection : Resample interesting time windows at higher FPS to inspect rapid motion and subtle visual artifacts.
Counting action & object : Accurately track repeated physical movements and distinct objects over time.
Token-efficient long-form video analysis
See how Gemini 3.7 Flash performs with and without agentic video understanding on LongVideoBench, a long-form video understanding benchmark. Notice the large token reductions and accuracy improvements.
Accurate fast action analysis with dynamic FPS
With agentic video understanding, 3.7 Flash is able to accurately count a fast-paced movement by scanning and rewatching the video at different frames per second, as needed.
Token-efficient needle-in-a-haystack search
Using agentic video understanding, Gemini 3.7 is able to accurately answer complex questions based on the content of the video while consuming a significantly lower number of tokens compared to static analysis.
Many of our early access partners saw strong performance while testing with agentic video understanding. Here’s what they have to say:
Agentic video understanding is available via the Gemini API in Google AI Studio and Gemini Enterprise Agent Platform , launching across Gemini 3.7 Flash, 3.6 Flash, and 3.5 Flash-Lite. It uses standard Gemini API token pricing with no additional feature fee.
To enable it, simply set processing to "agentic" in the API configuration. Read our developer guide to get more insights into the feature and how to get started.
from google import genai
client = genai.Client()
interaction = client.interactions.create(
model="gemini-3.7-flash",
input=[
{
"type": "video",
"uri": "https://youtu.be/7Z5Vy9JBANs",
"processing": "agentic"
},
{
"type": "text",
"text": "What are the 3 most important announcements in this keynote?",
},
],
)
print(interaction.output_text)
We are also bringing the efficiency and quality improvements of agentic video understanding to billions of users across Google products. The feature will roll out to all users in the Gemini app across Flash and Flash-Lite models soon. And in the coming months, agentic video understanding will also power YouTube's ‘ Ask YouTube’ feature on the video watch page, leveraging Gemini to deliver higher-quality answers grounded in the visuals.
Acknowledgement for their contribution to this work: Sergi Caelles, Filip Pavetić, Ahmet Iscen, Suhas Yogin, and the Agentic Vision team.
Get the latest news from Google in your inbox
Sign up for our newsletters with product updates, event information, special offers, and more.
Check your inbox to confirm your subscription.
You can also subscribe with a different email address .
Your information will be used in accordance with Google's privacy policy. You may opt out at any time.
Google Search gets a major upgrade with Gemini 4 Hydrogen
Gemini 4 Hydrogen: everyday coding and agentic workflows for developers
Introducing the Multimodal Knowledge Agent benchmark
Gemini 4 Carbon is now available
We’re bringing Adobe’s creative and productivity tools closer to your everyday Google apps.
More of Google
Google Products
About the Blog
Global (English)
Africa (English)
Australia (English)
Brasil (Português)
Canada (English)
Canada (Français)
Česko (Čeština)
Deutschland (Deutsch)
España (Español)
France (Français)
Greece (Ελληνικά)
India (English)
Indonesia (Bahasa Indonesia)
Ireland (English)
Italia (Italiano)
日本 (日本語)
대한민국 (한국어)
Latinoamérica (Español)
Malaysia (English)
الشرق الأوسط وشمال أفريقيا (اللغة العربية)
MENA (English)
Nederlands (Nederland)
New Zealand (English)
Polska (Polski)
Portugal (Português)
România (Română)
Sverige (Svenska)
ประเทศไทย (ไทย)
Türkiye (Türkçe)
台灣 (中文)

## Original Extract

We’re launching agentic video understanding across our latest Gemini models for improved accuracy and lower costs and token usage.

Skip to main content
Introducing agentic video understanding with Gemini
Models & Research
Google DeepMind
Infrastructure & cloud
Global network
Outreach & initiatives
Creating opportunity
Inside Google
Around the globe
Innovation & AI
Innovation & AI
Products & platforms
Products & platforms
Global (English)
Africa (English)
Australia (English)
Brasil (Português)
Canada (English)
Canada (Français)
Česko (Čeština)
Deutschland (Deutsch)
España (Español)
France (Français)
Greece (Ελληνικά)
India (English)
Indonesia (Bahasa Indonesia)
Ireland (English)
Italia (Italiano)
日本 (日本語)
대한민국 (한국어)
Latinoamérica (Español)
Malaysia (English)
الشرق الأوسط وشمال أفريقيا (اللغة العربية)
MENA (English)
Nederlands (Nederland)
New Zealand (English)
Polska (Polski)
Portugal (Português)
România (Română)
Sverige (Svenska)
ประเทศไทย (ไทย)
Türkiye (Türkçe)
台灣 (中文)
Introducing agentic video understanding with Gemini
Our new agentic feature for video analysis cuts token consumption by up to 88%, reduces costs by up to 66%, and boosts quality by up to 7%.
Senior Product Manager, Google DeepMind
Research Director, Google DeepMind
Your browser does not support the audio element.
Google just launched agentic video understanding for Gemini 3.7 Flash, 3.6 Flash, and 3.5 Flash-Lite. This feature allows the model to dynamically scan video segments, which improves accuracy while cutting token usage by up to 88% and costs by up to 66%. You can start using this today by setting your API configuration to "agentic" in Google AI Studio or the Gemini Enterprise Agent Platform.
Today, we’re launching agentic video understanding across our latest models: Gemini 3.7 Flash, 3.6 Flash and 3.5 Flash-Lite. This new capability improves accuracy while dramatically reducing token usage and costs for video analysis. Similar to agentic vision , which combines code execution with Gemini models’ native image understanding, agentic video understanding uses Gemini’s native video tools to improve performance and unlock new capabilities for video processing like sub-second moment retrieval, more accurate anomaly detection, precise counting and more.
The feature is available today for video uploads and YouTube videos via the Gemini API in Google AI Studio and the Gemini Enterprise Agent Platform.
Unlike current ‘static’ processing, where the model ingests the video at a fixed frames-per-second rate (default 1 FPS, adjustable via API), agentic video understanding pairs the model’s core reasoning with native video tools to dynamically search, scan, and inspect target video segments across visual frames, audio, and transcripts. Across standard video analysis benchmarks, Gemini models with agentic video understanding reduce analysis costs by up to 66% and token consumption by up to 88%, while improving accuracy by up to 7%.
These efficiency gains are especially pronounced on long-form video (from 10-minute how-to guides to 90-minute lectures and multi-hour recordings), where static processing forces developers to choose between high token costs or techniques that drop critical details.
Activating agentic video understanding drops token consumption by up to 88% and boosts accuracy by up to 7% with Gemini 3.7 Flash.
While these gains span all three supported models, Gemini 3.7 Flash with agentic understanding offers the best possible quality overall and the best combination of quality and cost efficiency, putting it at the accuracy-to-cost pareto frontier among tested models for video understanding.
Using agentic video understanding places Gemini 3.7 Flash at the accuracy-to-cost pareto frontier for video analysis.
Instead of static processing where the model ingests media streams at a fixed frame rate, agentic video understanding enables Gemini to take an active, goal-directed role in determining what to watch, at what speed, and through which modality (frames, audio, or transcript), fetching only the moments and signals needed. While developers could previously do this manually, with agentic video understanding, Gemini can accomplish it through an agentic loop, invoking an internal tool to load the relevant part of the video file, significantly reducing development overheads.
Agentic video understanding transforms how developers can process long-form video content across a variety of demanding applications.
Sub-second moment retrieval : Pinpoint split-second state changes and tight cut boundaries that are easily missed at 1 FPS, making precise automated video editing possible.
Long-form needle-in-a-haystack search : Answer complex queries across multi-hour videos without consuming millions of tokens.
Anomaly detection : Resample interesting time windows at higher FPS to inspect rapid motion and subtle visual artifacts.
Counting action & object : Accurately track repeated physical movements and distinct objects over time.
Token-efficient long-form video analysis
See how Gemini 3.7 Flash performs with and without agentic video understanding on LongVideoBench, a long-form video understanding benchmark. Notice the large token reductions and accuracy improvements.
Accurate fast action analysis with dynamic FPS
With agentic video understanding, 3.7 Flash is able to accurately count a fast-paced movement by scanning and rewatching the video at different frames per second, as needed.
Token-efficient needle-in-a-haystack search
Using agentic video understanding, Gemini 3.7 is able to accurately answer complex questions based on the content of the video while consuming a significantly lower number of tokens compared to static analysis.
Many of our early access partners saw strong performance while testing with agentic video understanding. Here’s what they have to say:
Agentic video understanding is available via the Gemini API in Google AI Studio and Gemini Enterprise Agent Platform , launching across Gemini 3.7 Flash, 3.6 Flash, and 3.5 Flash-Lite. It uses standard Gemini API token pricing with no additional feature fee.
To enable it, simply set processing to "agentic" in the API configuration. Read our developer guide to get more insights into the feature and how to get started.
from google import genai
client = genai.Client()
interaction = client.interactions.create(
model="gemini-3.7-flash",
input=[
{
"type": "video",
"uri": "https://youtu.be/7Z5Vy9JBANs",
"processing": "agentic"
},
{
"type": "text",
"text": "What are the 3 most important announcements in this keynote?",
},
],
)
print(interaction.output_text)
We are also bringing the efficiency and quality improvements of agentic video understanding to billions of users across Google products. The feature will roll out to all users in the Gemini app across Flash and Flash-Lite models soon. And in the coming months, agentic video understanding will also power YouTube's ‘ Ask YouTube’ feature on the video watch page, leveraging Gemini to deliver higher-quality answers grounded in the visuals.
Acknowledgement for their contribution to this work: Sergi Caelles, Filip Pavetić, Ahmet Iscen, Suhas Yogin, and the Agentic Vision team.
Get the latest news from Google in your inbox
Sign up for our newsletters with product updates, event information, special offers, and more.
Check your inbox to confirm your subscription.
You can also subscribe with a different email address .
Your information will be used in accordance with Google's privacy policy. You may opt out at any time.
Google Search gets a major upgrade with Gemini 4 Hydrogen
Gemini 4 Hydrogen: everyday coding and agentic workflows for developers
Introducing the Multimodal Knowledge Agent benchmark
Gemini 4 Carbon is now available
We’re bringing Adobe’s creative and productivity tools closer to your everyday Google apps.
More of Google
Google Products
About the Blog
Global (English)
Africa (English)
Australia (English)
Brasil (Português)
Canada (English)
Canada (Français)
Česko (Čeština)
Deutschland (Deutsch)
España (Español)
France (Français)
Greece (Ελληνικά)
India (English)
Indonesia (Bahasa Indonesia)
Ireland (English)
Italia (Italiano)
日本 (日本語)
대한민국 (한국어)
Latinoamérica (Español)
Malaysia (English)
الشرق الأوسط وشمال أفريقيا (اللغة العربية)
MENA (English)
Nederlands (Nederland)
New Zealand (English)
Polska (Polski)
Portugal (Português)
România (Română)
Sverige (Svenska)
ประเทศไทย (ไทย)
Türkiye (Türkçe)
台灣 (中文)
