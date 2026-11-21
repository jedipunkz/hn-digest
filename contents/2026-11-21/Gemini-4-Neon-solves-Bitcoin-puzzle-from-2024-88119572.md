---
source: "https://echo.ls/zden-level-halv/"
hn_url: "https://news.ycombinator.com/item?id=50592648"
title: "Gemini 4 Neon solves Bitcoin puzzle from 2024"
article_title: "Solving Zden's BTCrypto Puzzle Level HALV"
image: ""
author: "maplet"
captured_at: "2026-11-21T19:57:29Z"
capture_tool: "hn-digest"
hn_id: 50592648
score: 4
comments: 0
posted_at: "2026-11-21T19:19:03Z"
tags:
  - hacker-news
---

# Gemini 4 Neon solves Bitcoin puzzle from 2024

- HN: [50592648](https://news.ycombinator.com/item?id=50592648)
- Source: [echo.ls](https://echo.ls/zden-level-halv/)
- Score: 4
- Comments: 0
- Posted: 2026-11-21T19:19:03Z

## Translation

Title: Gemini 4 Neon solves Bitcoin puzzle from 2024
Article title: Solving Zden's BTCrypto Puzzle Level HALV

Article text:
Solving Zden's BTCrypto Puzzle Level HALV
Zden announced BTCrypto Puzzle Level HALV , often simply called “Level HALV” or “HALV”, on April 20, 2024. The puzzle consists of a single image with an encouraging description: “Level HALV - my new crypto puzzle to celebrate the fourth Bitcoin Halving. This level is way easier than LVL 5. It shouldn’t take long until it’s solved.” 943 days later (just 7 short of the image’s width and height of 950), it was solved.
After I solved Zden’s BTCrypto Puzzle Level 5 on September 22, Will Adcox and I continued working to solve HALV. Together, we tried a wide range of approaches and directed countless AI agents running state-of-the-art models as they were released. OpenAI models (GPT-6 Astra, GPT-6 Sol, GPT-6.1 Sol, GPT-6.2 Astra, GPT-6.2 Sol), Anthropic models (Fable 5.1, Opus 5.5, Sonnet 5.5, Fable 5.2), Gemini models (Argon, Carbon, Hydrogen, Neon), and others (GLM-5.3 Flash, DeepSeek V4.1 Flash) were all tested. I alone used almost 6 billion tokens in sessions targeting HALV, about 83% of which came from OpenAI models in Codex.
The winning setup was Orchestrator (Neon), using Hydrogen subagents, in Antigravity. The successful run was my first Gemini 4 Neon session, and it produced the solution after 9 hours and 19 minutes of work using 103 million tokens. My input was merely high-level guidance, e.g. “I think the solution must use the theme of bitcoin halving, the vertical stripe (possibly as a decoding rule?), and waveform simultaneously to solve the puzzle.” This write-up and explanation are my primary contributions to the problem.
The first step in decoding the puzzle is extracting the waveform’s peaks. We can define an intensity function which maps from pixel coordinates in the image to their 8-bit luminance, denoted by I : { 0 , 1 , … , 949 } × { 0 , 1 , … , 949 } → { 0 , 1 , … , 255 } I : \{ 0, 1, \ldots, 949 \} \times \{ 0, 1, \ldots, 949 \} \to \{ 0, 1, \ldots, 255 \} since the image has dimensions 950 × 950 950 \times 950 .
Then, we define the quantized level L ( s , j ) L(s, j) for each halving section s ∈ { 0 , 1 , … , 7 } s \in \{ 0, 1, \ldots, 7 \} and each intra-section column j ∈ { 0 , 1 , … , 31 } j \in \{ 0, 1, \ldots, 31 \} by
Since the waveform is symmetric, we could equivalently use 475 − v ⋅ 2 7 − s 475 - v \cdot 2^{7-s} for the y value. In addition, L L is well-defined as v v is unique for each s , j s, j .
Most of the apparent complexity of this assignment is mapping pixel coordinates to the encoded data. The idea here is that for any section s s and any column in that section j j , its quantized level L ( s , j ) L(s, j) is simply how many 128 / 2 s 128 / 2^s -pixel steps separate its peak from the centerline at y = 475 y = 475 . In fact, the image is constructed so that L ( s , j ) ∈ { 0 , 1 , 2 , 3 } , ∀ s , ∀ j L(s, j) \in \{ 0,1,2,3 \}, \forall s, \forall j .
Now that we have extracted the waveform data, we turn our attention to the narrow vertical stripe on the left side of the puzzle image; it is a 64 pixel stripe that corresponds to an 8 × 8 8 \times 8 circular gradient ( ( r + c ) m o d 8 ) ((r + c) \bmod 8) . We apply this rotation to the quantized columns with a step size of 32 / 8 = 4 32 / 8 = 4 .
The puzzle was funded with 312,500 sats, exactly one-thousandth of the block subsidy after the 4th bitcoin halving, which Zden explicitly said the puzzle was designed to celebrate. Each of the 8 sections in the image corresponds to a Bitcoin network epoch with a block subsidy of 50 / 2 s 50 / 2^s . We use these values to perform a weighted sum over the sections. This idea is also motivated by the halving staircase on the bottom right of the image, which is a modification to Zden’s typical puzzle logo.
Finally, we round the sums, represent each by two hexadecimal digits, and concatenate them in order of p p to obtain the private key.
The private key’s uncompressed secp256k1 public key yields the puzzle’s vanity Bitcoin P2PKH address, 1crypto24HCr178iMcKd5iUi5D4rsg1nK, verifying the solve.
The decoding process can be implemented as a small script that constructs the private key directly from the image. For instance, here is a Python solver which uses Pillow:
import sys
from PIL import Image
img = Image.open(sys.argv[1]).load()
key = bytearray(32)
for p in range(32):
total = 0.0
for s in range(8):
j = (p - 4 * s) % 32
x = 92 + 3 * (32 * s + j)
for v in (1, 2, 3):
y = 475 + v * 2**(7 - s)
if img[x, y] == 255:
total += v * 50 / 2**s
key[p] = round(total)
print(key.hex())
Once again, thank you to Zden for creating this puzzle series!

## Original Extract

Solving Zden's BTCrypto Puzzle Level HALV
Zden announced BTCrypto Puzzle Level HALV , often simply called “Level HALV” or “HALV”, on April 20, 2024. The puzzle consists of a single image with an encouraging description: “Level HALV - my new crypto puzzle to celebrate the fourth Bitcoin Halving. This level is way easier than LVL 5. It shouldn’t take long until it’s solved.” 943 days later (just 7 short of the image’s width and height of 950), it was solved.
After I solved Zden’s BTCrypto Puzzle Level 5 on September 22, Will Adcox and I continued working to solve HALV. Together, we tried a wide range of approaches and directed countless AI agents running state-of-the-art models as they were released. OpenAI models (GPT-6 Astra, GPT-6 Sol, GPT-6.1 Sol, GPT-6.2 Astra, GPT-6.2 Sol), Anthropic models (Fable 5.1, Opus 5.5, Sonnet 5.5, Fable 5.2), Gemini models (Argon, Carbon, Hydrogen, Neon), and others (GLM-5.3 Flash, DeepSeek V4.1 Flash) were all tested. I alone used almost 6 billion tokens in sessions targeting HALV, about 83% of which came from OpenAI models in Codex.
The winning setup was Orchestrator (Neon), using Hydrogen subagents, in Antigravity. The successful run was my first Gemini 4 Neon session, and it produced the solution after 9 hours and 19 minutes of work using 103 million tokens. My input was merely high-level guidance, e.g. “I think the solution must use the theme of bitcoin halving, the vertical stripe (possibly as a decoding rule?), and waveform simultaneously to solve the puzzle.” This write-up and explanation are my primary contributions to the problem.
The first step in decoding the puzzle is extracting the waveform’s peaks. We can define an intensity function which maps from pixel coordinates in the image to their 8-bit luminance, denoted by I : { 0 , 1 , … , 949 } × { 0 , 1 , … , 949 } → { 0 , 1 , … , 255 } I : \{ 0, 1, \ldots, 949 \} \times \{ 0, 1, \ldots, 949 \} \to \{ 0, 1, \ldots, 255 \} since the image has dimensions 950 × 950 950 \times 950 .
Then, we define the quantized level L ( s , j ) L(s, j) for each halving section s ∈ { 0 , 1 , … , 7 } s \in \{ 0, 1, \ldots, 7 \} and each intra-section column j ∈ { 0 , 1 , … , 31 } j \in \{ 0, 1, \ldots, 31 \} by
Since the waveform is symmetric, we could equivalently use 475 − v ⋅ 2 7 − s 475 - v \cdot 2^{7-s} for the y value. In addition, L L is well-defined as v v is unique for each s , j s, j .
Most of the apparent complexity of this assignment is mapping pixel coordinates to the encoded data. The idea here is that for any section s s and any column in that section j j , its quantized level L ( s , j ) L(s, j) is simply how many 128 / 2 s 128 / 2^s -pixel steps separate its peak from the centerline at y = 475 y = 475 . In fact, the image is constructed so that L ( s , j ) ∈ { 0 , 1 , 2 , 3 } , ∀ s , ∀ j L(s, j) \in \{ 0,1,2,3 \}, \forall s, \forall j .
Now that we have extracted the waveform data, we turn our attention to the narrow vertical stripe on the left side of the puzzle image; it is a 64 pixel stripe that corresponds to an 8 × 8 8 \times 8 circular gradient ( ( r + c ) m o d 8 ) ((r + c) \bmod 8) . We apply this rotation to the quantized columns with a step size of 32 / 8 = 4 32 / 8 = 4 .
The puzzle was funded with 312,500 sats, exactly one-thousandth of the block subsidy after the 4th bitcoin halving, which Zden explicitly said the puzzle was designed to celebrate. Each of the 8 sections in the image corresponds to a Bitcoin network epoch with a block subsidy of 50 / 2 s 50 / 2^s . We use these values to perform a weighted sum over the sections. This idea is also motivated by the halving staircase on the bottom right of the image, which is a modification to Zden’s typical puzzle logo.
Finally, we round the sums, represent each by two hexadecimal digits, and concatenate them in order of p p to obtain the private key.
The private key’s uncompressed secp256k1 public key yields the puzzle’s vanity Bitcoin P2PKH address, 1crypto24HCr178iMcKd5iUi5D4rsg1nK, verifying the solve.
The decoding process can be implemented as a small script that constructs the private key directly from the image. For instance, here is a Python solver which uses Pillow:
import sys
from PIL import Image
img = Image.open(sys.argv[1]).load()
key = bytearray(32)
for p in range(32):
total = 0.0
for s in range(8):
j = (p - 4 * s) % 32
x = 92 + 3 * (32 * s + j)
for v in (1, 2, 3):
y = 475 + v * 2**(7 - s)
if img[x, y] == 255:
total += v * 50 / 2**s
key[p] = round(total)
print(key.hex())
Once again, thank you to Zden for creating this puzzle series!
