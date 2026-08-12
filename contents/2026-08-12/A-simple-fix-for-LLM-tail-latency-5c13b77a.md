---
source: "https://engineering.myhoai.com/posts/a-simple-fix-for-llm-tail-latency/"
hn_url: "https://news.ycombinator.com/item?id=49278732"
title: "A simple fix for LLM tail latency"
article_title: "A simple fix for LLM tail latency | HOAi"
author: "zhixuan"
captured_at: "2026-08-12T21:35:28Z"
capture_tool: "hn-digest"
hn_id: 49278732
score: 1
comments: 0
posted_at: "2026-08-12T21:22:03Z"
tags:
  - hacker-news
  - translated
---

# A simple fix for LLM tail latency

- HN: [49278732](https://news.ycombinator.com/item?id=49278732)
- Source: [engineering.myhoai.com](https://engineering.myhoai.com/posts/a-simple-fix-for-llm-tail-latency/)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T21:22:03Z

## Translation

タイトル: LLM テール レイテンシの簡単な修正
記事のタイトル: LLM テール レイテンシの簡単な修正 |ホアイ
説明: すべてのリクエストを 2 回送信し、より速い応答を受け取ります。

記事本文:
LLM テール レイテンシの簡単な修正 |ホアイ、
したがって、FOUC はなく、インライン テーマ スクリプトも必要ありません。 -->
コンテンツにスキップ
li>a]:block [&>li>a]:px-4 [&>li>a]:py-3 [&>li>a]:text-center [&>li>a]:hover:text-accent [&>li>a]:hover:underline [&>li>a]:hover:underline-offset-4 sm:[&>li>a]:px-2 sm:[&>li>a]:py-1 hidden sm:mt-0 sm:flex sm:w-auto sm:gap-x-6 sm:gap-y-0"> エージェント
戻る LLM テール レイテンシーの簡単な修正
2026 年 6 月 10 日 | LLM の応答がリアルタイムのユースケースに対して遅すぎる場合、より高速なサービス層のために 2 倍のコストを支払おうとする誘惑に駆られるかもしれません。 Anthropic の優先度層、OpenAI の優先度処理、Gemini の優先度推論、LLM プロバイダーがそれをどのように呼んでも。もっと簡単な解決策があります。すべてのリクエストを 2 回送信し、より速い応答を受け取ります。
音声エージェントにとってテール遅延が重要な理由
HOAi の音声エージェントが電話に応答します。会話の各ターンで LLM リクエストが行われます。ほとんどの応答は 1.5 秒以内に返されますが、場合によっては 10 ～ 20 秒かかることがあります。電話の場合、気まずい沈黙が 10 秒続きます。十分な沈黙の後、システムは発信者との通話を切ります。
これはあなたが思っているよりも頻繁に起こります。通常の電話通話は 20 ～ 30 回かかります。 LLM リクエストの 1% が壊滅的に遅い場合、25 ターンの呼び出しで長い沈黙に陥る可能性は約 22% になります。
優先順位と各リクエストを 2 回送信する場合
OpenAI の優先レベルにアップグレードすると、トークンあたり 2 倍のコストを支払うことで、より高速で一貫性のある応答が得られます。
標準レベルを維持しますが、すべてのリクエストを 2 回送信し、より高速な応答を受け取ります。
両方のセットアップに対して 50 件の実際の運用リクエストを再生し、最初のトークンまでの時間 (エージェントが話し始めたとき) と応答が完了するまでの時間 (エージェントの呼び出しに対応できるとき) という 2 つのメトリクスを追跡しました。
リクエストを 2 回送信するのは明らかに外側です

優先層を形成しました。最悪の場合の応答完了までの時間は 9.8 秒から 3.5 秒に短縮されました。最初のトークンまでの最悪の場合の時間は 4.2 秒から 1.2 秒に短縮されました。個々のリクエストごとに標準レベルの方が遅いにもかかわらず、中央値でさえ優先レベルと正確に一致していました。
これは、応答が遅いことがまれで独立している場合に機能します。リクエストを 2 回送信すると、同じターンで両方のコピーが遅くなる可能性は低くなります。これにより、発信者が経験していた 10 秒間の沈黙が大幅に減少しました。
リアルタイムの対話型 LLM 製品を構築している場合は、より高速なサービス層の料金を支払う前に、リクエストを 2 回送信することに対してベンチマークを行ってください。同じコストでレイテンシーを改善できる可能性があります。
トップに戻る
この投稿を共有する: この投稿を X で共有する この投稿を LinkedIn で共有する この投稿を電子メールで共有する 前の投稿 モデルが考えすぎるとき 次の投稿 モデルに理由を聞いてください このページについて

## Original Extract

Send every request twice and take the faster response.

A simple fix for LLM tail latency | HOAi ,
so there is no FOUC and no inline theme script is needed. -->
Skip to content
li>a]:block [&>li>a]:px-4 [&>li>a]:py-3 [&>li>a]:text-center [&>li>a]:hover:text-accent [&>li>a]:hover:underline [&>li>a]:hover:underline-offset-4 sm:[&>li>a]:px-2 sm:[&>li>a]:py-1 hidden sm:mt-0 sm:flex sm:w-auto sm:gap-x-6 sm:gap-y-0"> Agents
Go back A simple fix for LLM tail latency
June 10, 2026 | When LLM responses are too slow for your realtime use case, you may be tempted to pay double the cost for a faster service tier. Anthropic’s Priority tier , OpenAI’s priority processing , Gemini’s priority inference , whatever your LLM provider calls it. There’s a simpler solution: send every request twice and take the faster response.
Why tail latency matters for voice agents
Our voice agent at HOAi answers phone calls. Every turn in a conversation makes an LLM request. Most responses come back within 1.5 seconds, but occasionally one takes 10 to 20 seconds. On a phone call, that’s 10 seconds of awkward silence, and after enough silence, our system hangs up on the caller.
This happens more often than you’d think. A typical phone call has 20 to 30 turns. If 1% of LLM requests are catastrophically slow, a 25-turn call has roughly a 22% chance of hitting a long silence.
Priority tier vs. sending each request twice
Upgrade to OpenAI’s priority tier and pay 2x cost per token for faster, more consistent responses.
Stay on standard tier, but send every request twice and take the faster response.
We replayed 50 real production requests against both setups and tracked two metrics: time to first token (when the agent starts speaking) and time to complete response (when it can act on tool calls).
Sending the request twice clearly outperformed the priority tier. Worst-case time to complete response dropped from 9.8s to 3.5s. Worst-case time to first token dropped from 4.2s to 1.2s. Even the median matched the priority tier exactly, despite the standard tier being slower per individual request.
This works when slow responses are rare and independent. Sending the request twice makes it unlikely both copies are slow on the same turn. This significantly reduced the 10-second silences our callers were experiencing.
If you are building a realtime interactive LLM product, before you pay for the faster service tier, benchmark it against sending the request twice. You may get better latency at the same cost.
Back To Top
Share this post on: Share this post on X Share this post on LinkedIn Share this post via email Previous Post When the model overthinks Next Post Just ask the model why On this page
