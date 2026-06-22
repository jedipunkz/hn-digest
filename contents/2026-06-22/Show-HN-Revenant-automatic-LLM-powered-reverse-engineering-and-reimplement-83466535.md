---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48630450"
title: "Show HN: Revenant – automatic LLM powered reverse engineering and reimplement"
article_title: ""
author: "sylwester"
captured_at: "2026-06-22T14:53:12Z"
capture_tool: "hn-digest"
hn_id: 48630450
score: 3
comments: 0
posted_at: "2026-06-22T14:15:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Revenant – automatic LLM powered reverse engineering and reimplement

- HN: [48630450](https://news.ycombinator.com/item?id=48630450)
- Score: 3
- Comments: 0
- Posted: 2026-06-22T14:15:32Z

## Translation

タイトル: Show HN: Revenant – LLM による自動リバース エンジニアリングと再実装
HN テキスト: 私はハードウェア エンジニアでありセキュリティ研究者であり、他のトピックにも集中できるように自分の仕事を部分的に自動化できないかと考えていました。そこで、radare2、ghidra などを中心に構築され、ファームウェアを完全に自動的に分析し、オープン ソース スケルトンを実装できる LLM 搭載 (Claude、OpenAI、ローカル AI) ツールキットである revenant を構築します。ピン配置、ハードウェアの立ち上げ、ペリフェラルの立ち上げなど、または既存のファームウェアを 1:1 で複製することもできるため、古いハードウェアを最新のツールチェーンで復活させることができます。いくつかのアプリケーションは次のとおりです: - 古いハードウェアに新しい命を吹き込みます - 怪しいファームウェアのセキュリティ分析 次の場所で確認してください。
https://github.com/DatanoiseTV/revenant これに関する意見や、改善のための推奨事項をいただければ幸いです。

## Original Extract

I am a hardware engineer and security researcher and I've been wondering whether my work could be partially automated, so I can focus on other topics as well, so I build revenant - a LLM powered (Claude, OpenAI, local AI) toolkit that builds around radare2, ghidra etc and can fully automatically analyze firmware, implement open source skeletons incl. pinouts, hardware bringup, peripheral bringup etc. or can even 1:1 replicate existing firmware so old hardware can be resurrected with modern toolchains. Some applications are: - Give old hardware new life - Security Analysis of shady firmware Check it at:
https://github.com/DatanoiseTV/revenant I would love some input on this and maybe some recommendations for improvements.

