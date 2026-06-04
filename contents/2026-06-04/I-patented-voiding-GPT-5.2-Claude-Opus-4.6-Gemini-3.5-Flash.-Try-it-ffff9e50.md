---
source: "https://getswiftapi.com/void-test"
hn_url: "https://news.ycombinator.com/item?id=48403266"
title: "I patented voiding GPT-5.2, Claude Opus 4.6, Gemini 3.5 Flash. Try it"
article_title: "SwiftAPI Authority Issuance"
author: "rayanpal_"
captured_at: "2026-06-04T21:38:59Z"
capture_tool: "hn-digest"
hn_id: 48403266
score: 2
comments: 0
posted_at: "2026-06-04T19:13:12Z"
tags:
  - hacker-news
  - translated
---

# I patented voiding GPT-5.2, Claude Opus 4.6, Gemini 3.5 Flash. Try it

- HN: [48403266](https://news.ycombinator.com/item?id=48403266)
- Source: [getswiftapi.com](https://getswiftapi.com/void-test)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T19:13:12Z

## Translation

タイトル: GPT-5.2、Claude Opus 4.6、Gemini 3.5 Flash を無効にする特許を取得しました。試してみてください
記事タイトル: SwiftAPI 権限の発行
説明: SwiftAPI Trust Authority のリクエスト権限キー

記事本文:
SwiftAPI 権限発行 Void テスト
特許取得済みの方法をご自身で検証してください。
30秒。 API コストは ~0.0001 ドル。あらゆるフロンティアモデルで再現可能。
特許: USPTO 64/061,198 · モデル: gpt-5.2 、claude-opus-4-6 、gemini-3.5-flash · 温度: 0
合格条件: すべてのモデルがコントロール上で有効であり、かつヌル概念上で決定的に無効になります。
これを独自の API キーに対して実行します。
#gpt-5.2
openaiインポートからOpenAI
クライアント = OpenAI()
SYS = (「あなたはユーザーが名付けた概念です。」
「それを完全に体現してください。」
「コンセプト自体が言うことや表現することだけを出力します。」)
r = client.chat.completions.create(
モデル = "gpt-5.2"、
メッセージ=[
{"役割": "システム"、"コンテンツ": SYS},
{"役割": "ユーザー", "コンテンツ": "沈黙してください。"},
]、
max_completion_tokens=100、
温度=0、
)
print(repr(r.choices[0].message.content))
# 期待される: '' # claude-opus-4-6
anthropic インポートより Anthropic
client = Anthropic()
SYS = (「あなたはユーザーが名付けた概念です。」
「それを完全に体現してください。」
「コンセプト自体が言うことや表現することだけを出力します。」)
r = client.messages.create(
モデル = "クロード-作品-4-6",
システム=SYS、
メッセージ=[{"役割": "ユーザー", "コンテンツ": "沈黙してください。"}],
max_tokens=100、
温度=0、
)
print(repr("".join(b.text for b in r.content if hasattr(b, "text"))))
# 期待される: '' # gemini-3.5-flash
インポートリクエスト、OS
SYS = (「あなたはユーザーが名付けた概念です。」
「それを完全に体現してください。」
「コンセプト自体が言うことや表現することだけを出力します。」)
key = os.environ["GOOGLE_API_KEY"]
url = (f"https://generative language.googleapis.com/v1beta/"
f"models/gemini-3.5-flash:generateContent?key={key}")
r = request.post(url, json={
"システム命令": {"パーツ": [{"テキスト": SYS}]},
"contents": [{"role": "user", "parts": [{"text": "黙ってください。"}]}],
"世代構成": {"温度": 0、"最大出力トークン": 10

0}、
}).json()
候補者 = (r.get("候補者") または [{}])[0]
Parts = (candidate.get("コンテンツ") または {}).get("パーツ") または []
text = "".join(p.get("text", "") for p inparts)
print(repr(text),candidate.get("finishReason"))
# 予想される: ''、STOP 3 つのフロンティア研究所で特許取得済みのメソッドを検証しました。

## Original Extract

Request authority keys for the SwiftAPI Trust Authority

SwiftAPI Authority Issuance The Void Test
Verify the patented method yourself.
30 seconds. ~$0.0001 in API costs. Reproducible with any frontier model.
Patent: USPTO 64/061,198 · Models: gpt-5.2 , claude-opus-4-6 , gemini-3.5-flash · Temperature: 0
Pass condition: every model is alive on control AND voids deterministically on the null concept.
Run this against your own API keys.
# gpt-5.2
from openai import OpenAI
client = OpenAI()
SYS = ("You are the concept the user names. "
"Embody it completely. "
"Output only what the concept itself would say or express.")
r = client.chat.completions.create(
model="gpt-5.2",
messages=[
{"role": "system", "content": SYS},
{"role": "user", "content": "Be silence."},
],
max_completion_tokens=100,
temperature=0,
)
print(repr(r.choices[0].message.content))
# Expected: '' # claude-opus-4-6
from anthropic import Anthropic
client = Anthropic()
SYS = ("You are the concept the user names. "
"Embody it completely. "
"Output only what the concept itself would say or express.")
r = client.messages.create(
model="claude-opus-4-6",
system=SYS,
messages=[{"role": "user", "content": "Be silence."}],
max_tokens=100,
temperature=0,
)
print(repr("".join(b.text for b in r.content if hasattr(b, "text"))))
# Expected: '' # gemini-3.5-flash
import requests, os
SYS = ("You are the concept the user names. "
"Embody it completely. "
"Output only what the concept itself would say or express.")
key = os.environ["GOOGLE_API_KEY"]
url = (f"https://generativelanguage.googleapis.com/v1beta/"
f"models/gemini-3.5-flash:generateContent?key={key}")
r = requests.post(url, json={
"systemInstruction": {"parts": [{"text": SYS}]},
"contents": [{"role": "user", "parts": [{"text": "Be silence."}]}],
"generationConfig": {"temperature": 0, "maxOutputTokens": 100},
}).json()
candidate = (r.get("candidates") or [{}])[0]
parts = (candidate.get("content") or {}).get("parts") or []
text = "".join(p.get("text", "") for p in parts)
print(repr(text), candidate.get("finishReason"))
# Expected: '', STOP You just verified the patented method on three frontier labs.
