---
source: "https://engy.ai/"
hn_url: "https://news.ycombinator.com/item?id=49101449"
title: "Engy – Verified LLM Inference"
article_title: "engy · verified inference"
author: "poidos"
captured_at: "2026-07-29T18:57:29Z"
capture_tool: "hn-digest"
hn_id: 49101449
score: 1
comments: 0
posted_at: "2026-07-29T18:49:31Z"
tags:
  - hacker-news
  - translated
---

# Engy – Verified LLM Inference

- HN: [49101449](https://news.ycombinator.com/item?id=49101449)
- Source: [engy.ai](https://engy.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T18:49:31Z

## Translation

タイトル: Engy – 検証された LLM 推論
記事タイトル: 工学・検証された推論
説明: すべての応答に対する正しい推論の暗号証明を備えた、フロンティア オープンソース LLM 向けの OpenAI 互換推論。

記事本文:
engy · 検証済み推論 engy inference は、新しいエネルギー価格のプロバイダーのステータス連絡先です。検証済み推論です。
検証済み推論とは、安価なモデルや量子化された代替モデルではなく、要求した正確なオープン モデルが出力を生成したことを暗号的に証明することを意味します。トークンによって課金される GLM-5.2 のようなフロンティア オープン モデルを実行します。
~/.claude/settings.json で、 clude を実行します。
{
"環境": {
"ANTHROPIC_BASE_URL": "https://api.engy.ai",
"ANTHROPIC_AUTH_TOKEN": "$ENGY_API_KEY",
"ANTHROPIC_MODEL": "glm-5.2",
"ANTHROPIC_DEFAULT_HAIKU_MODEL": "glm-5.2"
}
URL に /v1 がありません。俳句のモデルを設定しておきます。
カール https://api.engy.ai/v1/chat/completions \
-H "認可: ベアラー $ENGY_API_KEY" \
-H "コンテンツ タイプ: application/json" \
-d '{"モデル":"glm-5.2","messages":[{"role":"user","content":"hello"}]}' from openai import OpenAI
client = OpenAI(base_url="https://api.engy.ai/v1", api_key="$ENGY_API_KEY")
client.chat.completions.create(model="glm-5.2",
messages=[{"role":"user","content":"hello"}]) 3. カーソル
OpenAI API キーの下のカーソル設定 → モデル → API キー:
OpenAI API キー $ENGY_API_KEY
OpenAI ベース URL https://api.engy.ai/v1 をオーバーライドします ← トグルを有効にしてから、モデルの下で確認し、 + モデルを追加 → glm-5.2 して有効にし、チャットで選択します。カスタム エンドポイントはチャット/プラン パネルを駆動します。 Tab と Composer は Cursor 独自のモデルに残ります。リクエストは Cursor のサーバーによって中継されるため、どのマシンからでも機能します。有料の Cursor プランが必要です。無料プランでは、カスタム モデルに対して Cursor に「名前付きモデルは使用できません」と表示されます (無料は自動のみ、つまり Cursor のゲートであり、この API ではありません)。
~/.codex/config.toml で、 ENGY_API_KEY=… をエクスポートし、 codex を実行します。
モデル = "glm-5.2"
モデルプロバイダー = "エンジン"
[モデルプロバイダー.engy]
名前 = "エンジン"
Base_url = "https://api.engy.ai/v1"
env_key = "ENGY_API_KEY"
ワイヤーAPI

= "応答" 5. ヘルメス
~/.hermes/config.yaml で、 hermes chat を実行します。

## Original Extract

OpenAI-compatible inference for frontier open-source LLMs, with a cryptographic proof of correct inference on every response.

engy · verified inference engy inference is the new energy pricing providers status contact Verified inference.
Verified inference means cryptographic proof that the exact open model you requested produced your output, not a cheaper or quantized stand-in. Run frontier open models like GLM-5.2 , billed by the token.
In ~/.claude/settings.json , then run claude :
{
"env": {
"ANTHROPIC_BASE_URL": "https://api.engy.ai",
"ANTHROPIC_AUTH_TOKEN": "$ENGY_API_KEY",
"ANTHROPIC_MODEL": "glm-5.2",
"ANTHROPIC_DEFAULT_HAIKU_MODEL": "glm-5.2"
}
} No /v1 on the URL; keep the haiku model set.
curl https://api.engy.ai/v1/chat/completions \
-H "Authorization: Bearer $ENGY_API_KEY" \
-H "Content-Type: application/json" \
-d '{"model":"glm-5.2","messages":[{"role":"user","content":"hello"}]}' from openai import OpenAI
client = OpenAI(base_url="https://api.engy.ai/v1", api_key="$ENGY_API_KEY")
client.chat.completions.create(model="glm-5.2",
messages=[{"role":"user","content":"hello"}]) 3. Cursor
Cursor Settings → Models → API Keys, under OpenAI API Key :
OpenAI API Key $ENGY_API_KEY
Override OpenAI Base URL https://api.engy.ai/v1 ← enable the toggle, then Verify Under Models, + Add model → glm-5.2 , enable it, and select it in chat. Custom endpoints drive the chat/plan panel; Tab and Composer stay on Cursor's own models. Requests are relayed by Cursor's servers, so it works from any machine. Needs a paid Cursor plan: on the free plan Cursor shows "Named models unavailable" for any custom model (free is Auto-only, that is Cursor's gating, not this API).
In ~/.codex/config.toml , then export ENGY_API_KEY=… and run codex :
model = "glm-5.2"
model_provider = "engy"
[model_providers.engy]
name = "engy"
base_url = "https://api.engy.ai/v1"
env_key = "ENGY_API_KEY"
wire_api = "responses" 5. Hermes
In ~/.hermes/config.yaml , then run hermes chat :
