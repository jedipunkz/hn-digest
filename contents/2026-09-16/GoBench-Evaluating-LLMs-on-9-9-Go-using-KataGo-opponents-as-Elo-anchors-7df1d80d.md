---
source: "https://rolandgao.com/blog/gobench/"
hn_url: "https://news.ycombinator.com/item?id=49730958"
title: "GoBench: Evaluating LLMs on 9×9 Go using KataGo opponents as Elo anchors"
article_title: "GoBench | Roland Gao"
image: "https://rolandgao.com/social-preview.png"
author: "GodelNumbering"
captured_at: "2026-09-16T18:42:14Z"
capture_tool: "hn-digest"
hn_id: 49730958
score: 1
comments: 1
posted_at: "2026-09-16T18:23:36Z"
tags:
  - hacker-news
---

# GoBench: Evaluating LLMs on 9×9 Go using KataGo opponents as Elo anchors

- HN: [49730958](https://news.ycombinator.com/item?id=49730958)
- Source: [rolandgao.com](https://rolandgao.com/blog/gobench/)
- Score: 1
- Comments: 1
- Posted: 2026-09-16T18:23:36Z

## Translation

Title: GoBench: Evaluating LLMs on 9×9 Go using KataGo opponents as Elo anchors
Article title: GoBench | Roland Gao
Description: Measuring how well frontier language models play 9×9 Go against a calibrated ladder of KataGo opponents.

Article text:
Updated: September 15, 2026 | Original: September 15, 2026 | Author: Roland Gao
GoBench measures how well frontier language models play 9×9 Go against a calibrated ladder of KataGo opponents.
The arXiv release is scheduled for September 17, 2026.
0 1,000 2,000 3,000 4,000 10⁻⁷ 10⁻⁵ 10⁻³ 10⁻¹ KataGo (strongest), 4,427 plus or minus 35 Elo, $0.0049 / move KataGo, 4,385 plus or minus 35 Elo, $0.0051 / move KataGo, 4,384 plus or minus 35 Elo, $0.0049 / move KataGo, 4,368 plus or minus 35 Elo, $0.0019 / move KataGo, 4,346 plus or minus 34 Elo, $0.00051 / move KataGo, 4,332 plus or minus 34 Elo, $0.00049 / move KataGo, 4,308 plus or minus 34 Elo, $0.00049 / move KataGo, 4,278 plus or minus 34 Elo, $0.00019 / move KataGo, 4,216 plus or minus 36 Elo, $2.3 × 10⁻⁵ / move KataGo, 4,199 plus or minus 34 Elo, $8.5 × 10⁻⁶ / move KataGo, 4,191 plus or minus 34 Elo, $8.2 × 10⁻⁶ / move KataGo, 4,160 plus or minus 34 Elo, $8.2 × 10⁻⁶ / move KataGo, 4,131 plus or minus 34 Elo, $3.2 × 10⁻⁶ / move KataGo, 4,102 plus or minus 34 Elo, $1 × 10⁻⁵ / move KataGo, 4,070 plus or minus 34 Elo, $4.1 × 10⁻⁶ / move KataGo, 4,063 plus or minus 34 Elo, $4.1 × 10⁻⁶ / move KataGo, 4,027 plus or minus 34 Elo, $4.1 × 10⁻⁶ / move KataGo, 3,994 plus or minus 34 Elo, $4.1 × 10⁻⁶ / move KataGo, 3,987 plus or minus 33 Elo, $4.1 × 10⁻⁶ / move KataGo, 3,931 plus or minus 34 Elo, $2.1 × 10⁻⁶ / move KataGo, 3,898 plus or minus 33 Elo, $2.1 × 10⁻⁶ / move KataGo, 3,861 plus or minus 33 Elo, $2.1 × 10⁻⁶ / move KataGo, 3,847 plus or minus 33 Elo, $2.1 × 10⁻⁶ / move KataGo, 3,841 plus or minus 33 Elo, $2.7 × 10⁻⁶ / move KataGo, 3,797 plus or minus 33 Elo, $2.2 × 10⁻⁶ / move KataGo, 3,752 plus or minus 33 Elo, $8.7 × 10⁻⁷ / move KataGo, 3,750 plus or minus 33 Elo, $9.1 × 10⁻⁷ / move KataGo, 3,736 plus or minus 33 Elo, $2.1 × 10⁻⁶ / move KataGo, 3,700 plus or minus 33 Elo, $8.9 × 10⁻⁷ / move KataGo, 3,689 plus or minus 33 Elo, $9.5 × 10⁻⁷ / move KataGo, 3,638 plus or minus 33 Elo, $9.4 × 10⁻⁷ / move KataGo, 3,636 plus or minus 33 Elo, $9.2 × 10⁻⁷ / move KataGo, 3,565 plus or minus 32 Elo, $3.2 × 10⁻⁷ / move KataGo, 3,556 plus or minus 32 Elo, $3.2 × 10⁻⁷ / move KataGo, 3,535 plus or minus 32 Elo, $3.5 × 10⁻⁷ / move KataGo, 3,529 plus or minus 32 Elo, $3.2 × 10⁻⁷
[truncated]
Current models show “ jagged intelligence ”: they approach top human performance in math and coding, yet lag in other domains. Progress toward AGI requires systems that can learn new domains at lower cost and with less human supervision .
GoBench measures general reasoning and context-based continual learning through two tracks:
Track 1 uses multi-turn APIs without tools to test LLMs’ general reasoning abilities.
Context-based continual learning
Track 2 gives Codex 0, 1, 2, 4, or 8 hours of continual learning before evaluation. Each evaluation game is capped at 30 minutes. Learning and evaluation take place in a sandbox with resource limits and no internet access.
We encourage researchers to extend GoBench to measure continual learning through weight updates.
GPT-6 Astra · Max 2,568 ± 140 $0.15 $0.15 66s 2-4 GPT-6 Astra · High 2,227 ± 138 $0.033 $0.033 10s 2-6 Claude Opus 5 · High 2,076 ± 139 $0.09 $0.09 26s 2-8 GPT-5.6 Sol · Max 1,929 ± 207 $0.081 $0.081 79s 3-10 GPT-5.6 Sol · High 1,846 ± 207 $0.034 $0.034 27s 3-11 Gemini 3.1 Pro · High 1,797 ± 178 $0.063 $0.063 41s 4-13 DeepSeek V4.1 Flash · Max 1,707 ± 206 $0.022 $0.022 140s 4-13 Gemini 3.8 Flash · High 1,616 ± 206 $0.025 $0.025 13s 5-17 DeepSeek V4 Flash 0731 · Max 1,449 ± 227 $0.016 $0.016 160s 5-17 Muse Spark 1.3 Contributor · High 1,449 ± 227 $0.002 $0.002 46s 6-17 Muse Spark 1.3 Contributor · Extra high 1,415 ± 217 $0.0025 $0.0025 38s 7-18 DeepSeek V4.1 Flash · High 1,380 ± 205 $0.012 $0.012 70s 7-18 Gemini 3.6 Flash · High 1,336 ± 223 $0.02 $0.02 5.3s 9-18 DeepSeek V4 Flash 0731 · High 1,189 ± 197 $0.014 $0.014 150s 9-18 GPT-5.6 Luna · High 1,091 ± 190 $0.0024 $0.0024 17s 9-18 GPT-5.6 Luna · Max 1,058 ± 189 $0.0043 $0.0043 32s 9-18 Grok 4.6 · Extra high 1,044 ± 189 $0.05 $0.05 35s 12-18 Grok 4.6 · High 916 ± 263 $0.044 $0.044 30s KataGo references — KataGo (fastest) 3,332 ± 32 $1.4 × 10⁻⁷ $1.4e-7 0.0072s — KataGo (strongest) 4,427 ± 35 $0.0049 $0.0049 250s
Track 2 Leaderboard
0 1,000 2,000 3,000 4,000 0 1 2 4 8 GPT-6 Astra · High · Codex · 0h, 2,691 plus or minus 186 Elo, 0 hours GPT-6 Astra · High · Codex · 1h, 3,149 plus or minus 196 Elo, 1 hour GPT-6 Astra · High · Codex · 2h, 3,563 plus or minus 181 Elo, 2 hours GPT-6 Astra · High · Codex · 4h, 3,421 plus or minus 175 Elo, 4 hours GPT-6 Astra · High · Codex · 8h, 3,436 plus or minus 183 Elo, 8 hours GPT-5.6 Sol · High · Codex · 0h, 1,580 plus or minus 178 Elo, 0 hours GPT-5.6 Sol · High · Codex · 1h, 2,656 plus or minus 166 Elo, 1 hour GPT-5.6 Sol · High · Codex · 2h, 2,073 plus or minus 165 Elo, 2 hours GPT-5.6 Sol · High · Codex · 4h, 1,910 plus or minus 181 Elo, 4 hours GPT-5.6 Sol · High · Codex · 8h, 1,046 plus or minus 194 Elo, 8 hours Continual learning duration (hours) Elo GPT-6 Astra · Codex GPT-5.6 Sol · Codex View Track 2 results as a table Track 2 Elo ratings by continual learning duration Model Hours Elo ± 95% CI GPT-6 Astra · Codex 0 2,691 ± 186 GPT-6 Astra · Codex 1 3,149 ± 196 GPT-6 Astra · Codex 2 3,563 ± 181 GPT-6 Astra · Codex 4 3,421 ± 175 GPT-6 Astra · Codex 8 3,436 ± 183 GPT-5.6 Sol · Codex 0 1,580 ± 178 GPT-5.6 Sol · Codex 1 2,656 ± 166 GPT-5.6 Sol · Codex 2 2,073 ± 165 GPT-5.6 Sol · Codex 4 1,910 ± 181 GPT-5.6 Sol · Codex 8 1,046 ± 194
Replay every match
Play 9×9 Go against the same calibrated KataGo opponents. Pick an Elo, choose a color, and click Start game to play locally in your browser. Tromp-Taylor rules: area scoring, self-capture allowed, positional superko. 7 komi.
A 9 B 8 C 7 D 6 E 5 F 4 G 3 H 2 J 1 Latest move No moves played Start game Undo Pass Resign Your estimated Elo 1,000 ± 3,920 Past games
Completed games will appear here.

## Original Extract

Measuring how well frontier language models play 9×9 Go against a calibrated ladder of KataGo opponents.

Updated: September 15, 2026 | Original: September 15, 2026 | Author: Roland Gao
GoBench measures how well frontier language models play 9×9 Go against a calibrated ladder of KataGo opponents.
The arXiv release is scheduled for September 17, 2026.
0 1,000 2,000 3,000 4,000 10⁻⁷ 10⁻⁵ 10⁻³ 10⁻¹ KataGo (strongest), 4,427 plus or minus 35 Elo, $0.0049 / move KataGo, 4,385 plus or minus 35 Elo, $0.0051 / move KataGo, 4,384 plus or minus 35 Elo, $0.0049 / move KataGo, 4,368 plus or minus 35 Elo, $0.0019 / move KataGo, 4,346 plus or minus 34 Elo, $0.00051 / move KataGo, 4,332 plus or minus 34 Elo, $0.00049 / move KataGo, 4,308 plus or minus 34 Elo, $0.00049 / move KataGo, 4,278 plus or minus 34 Elo, $0.00019 / move KataGo, 4,216 plus or minus 36 Elo, $2.3 × 10⁻⁵ / move KataGo, 4,199 plus or minus 34 Elo, $8.5 × 10⁻⁶ / move KataGo, 4,191 plus or minus 34 Elo, $8.2 × 10⁻⁶ / move KataGo, 4,160 plus or minus 34 Elo, $8.2 × 10⁻⁶ / move KataGo, 4,131 plus or minus 34 Elo, $3.2 × 10⁻⁶ / move KataGo, 4,102 plus or minus 34 Elo, $1 × 10⁻⁵ / move KataGo, 4,070 plus or minus 34 Elo, $4.1 × 10⁻⁶ / move KataGo, 4,063 plus or minus 34 Elo, $4.1 × 10⁻⁶ / move KataGo, 4,027 plus or minus 34 Elo, $4.1 × 10⁻⁶ / move KataGo, 3,994 plus or minus 34 Elo, $4.1 × 10⁻⁶ / move KataGo, 3,987 plus or minus 33 Elo, $4.1 × 10⁻⁶ / move KataGo, 3,931 plus or minus 34 Elo, $2.1 × 10⁻⁶ / move KataGo, 3,898 plus or minus 33 Elo, $2.1 × 10⁻⁶ / move KataGo, 3,861 plus or minus 33 Elo, $2.1 × 10⁻⁶ / move KataGo, 3,847 plus or minus 33 Elo, $2.1 × 10⁻⁶ / move KataGo, 3,841 plus or minus 33 Elo, $2.7 × 10⁻⁶ / move KataGo, 3,797 plus or minus 33 Elo, $2.2 × 10⁻⁶ / move KataGo, 3,752 plus or minus 33 Elo, $8.7 × 10⁻⁷ / move KataGo, 3,750 plus or minus 33 Elo, $9.1 × 10⁻⁷ / move KataGo, 3,736 plus or minus 33 Elo, $2.1 × 10⁻⁶ / move KataGo, 3,700 plus or minus 33 Elo, $8.9 × 10⁻⁷ / move KataGo, 3,689 plus or minus 33 Elo, $9.5 × 10⁻⁷ / move KataGo, 3,638 plus or minus 33 Elo, $9.4 × 10⁻⁷ / move KataGo, 3,636 plus or minus 33 Elo, $9.2 × 10⁻⁷ / move KataGo, 3,565 plus or minus 32 Elo, $3.2 × 10⁻⁷ / move KataGo, 3,556 plus or minus 32 Elo, $3.2 × 10⁻⁷ / move KataGo, 3,535 plus or minus 32 Elo, $3.5 × 10⁻⁷ / move KataGo, 3,529 plus or minus 32 Elo, $3.2 × 10⁻⁷
[truncated]
Current models show “ jagged intelligence ”: they approach top human performance in math and coding, yet lag in other domains. Progress toward AGI requires systems that can learn new domains at lower cost and with less human supervision .
GoBench measures general reasoning and context-based continual learning through two tracks:
Track 1 uses multi-turn APIs without tools to test LLMs’ general reasoning abilities.
Context-based continual learning
Track 2 gives Codex 0, 1, 2, 4, or 8 hours of continual learning before evaluation. Each evaluation game is capped at 30 minutes. Learning and evaluation take place in a sandbox with resource limits and no internet access.
We encourage researchers to extend GoBench to measure continual learning through weight updates.
GPT-6 Astra · Max 2,568 ± 140 $0.15 $0.15 66s 2-4 GPT-6 Astra · High 2,227 ± 138 $0.033 $0.033 10s 2-6 Claude Opus 5 · High 2,076 ± 139 $0.09 $0.09 26s 2-8 GPT-5.6 Sol · Max 1,929 ± 207 $0.081 $0.081 79s 3-10 GPT-5.6 Sol · High 1,846 ± 207 $0.034 $0.034 27s 3-11 Gemini 3.1 Pro · High 1,797 ± 178 $0.063 $0.063 41s 4-13 DeepSeek V4.1 Flash · Max 1,707 ± 206 $0.022 $0.022 140s 4-13 Gemini 3.8 Flash · High 1,616 ± 206 $0.025 $0.025 13s 5-17 DeepSeek V4 Flash 0731 · Max 1,449 ± 227 $0.016 $0.016 160s 5-17 Muse Spark 1.3 Contributor · High 1,449 ± 227 $0.002 $0.002 46s 6-17 Muse Spark 1.3 Contributor · Extra high 1,415 ± 217 $0.0025 $0.0025 38s 7-18 DeepSeek V4.1 Flash · High 1,380 ± 205 $0.012 $0.012 70s 7-18 Gemini 3.6 Flash · High 1,336 ± 223 $0.02 $0.02 5.3s 9-18 DeepSeek V4 Flash 0731 · High 1,189 ± 197 $0.014 $0.014 150s 9-18 GPT-5.6 Luna · High 1,091 ± 190 $0.0024 $0.0024 17s 9-18 GPT-5.6 Luna · Max 1,058 ± 189 $0.0043 $0.0043 32s 9-18 Grok 4.6 · Extra high 1,044 ± 189 $0.05 $0.05 35s 12-18 Grok 4.6 · High 916 ± 263 $0.044 $0.044 30s KataGo references — KataGo (fastest) 3,332 ± 32 $1.4 × 10⁻⁷ $1.4e-7 0.0072s — KataGo (strongest) 4,427 ± 35 $0.0049 $0.0049 250s
Track 2 Leaderboard
0 1,000 2,000 3,000 4,000 0 1 2 4 8 GPT-6 Astra · High · Codex · 0h, 2,691 plus or minus 186 Elo, 0 hours GPT-6 Astra · High · Codex · 1h, 3,149 plus or minus 196 Elo, 1 hour GPT-6 Astra · High · Codex · 2h, 3,563 plus or minus 181 Elo, 2 hours GPT-6 Astra · High · Codex · 4h, 3,421 plus or minus 175 Elo, 4 hours GPT-6 Astra · High · Codex · 8h, 3,436 plus or minus 183 Elo, 8 hours GPT-5.6 Sol · High · Codex · 0h, 1,580 plus or minus 178 Elo, 0 hours GPT-5.6 Sol · High · Codex · 1h, 2,656 plus or minus 166 Elo, 1 hour GPT-5.6 Sol · High · Codex · 2h, 2,073 plus or minus 165 Elo, 2 hours GPT-5.6 Sol · High · Codex · 4h, 1,910 plus or minus 181 Elo, 4 hours GPT-5.6 Sol · High · Codex · 8h, 1,046 plus or minus 194 Elo, 8 hours Continual learning duration (hours) Elo GPT-6 Astra · Codex GPT-5.6 Sol · Codex View Track 2 results as a table Track 2 Elo ratings by continual learning duration Model Hours Elo ± 95% CI GPT-6 Astra · Codex 0 2,691 ± 186 GPT-6 Astra · Codex 1 3,149 ± 196 GPT-6 Astra · Codex 2 3,563 ± 181 GPT-6 Astra · Codex 4 3,421 ± 175 GPT-6 Astra · Codex 8 3,436 ± 183 GPT-5.6 Sol · Codex 0 1,580 ± 178 GPT-5.6 Sol · Codex 1 2,656 ± 166 GPT-5.6 Sol · Codex 2 2,073 ± 165 GPT-5.6 Sol · Codex 4 1,910 ± 181 GPT-5.6 Sol · Codex 8 1,046 ± 194
Replay every match
Play 9×9 Go against the same calibrated KataGo opponents. Pick an Elo, choose a color, and click Start game to play locally in your browser. Tromp-Taylor rules: area scoring, self-capture allowed, positional superko. 7 komi.
A 9 B 8 C 7 D 6 E 5 F 4 G 3 H 2 J 1 Latest move No moves played Start game Undo Pass Resign Your estimated Elo 1,000 ± 3,920 Past games
Completed games will appear here.
