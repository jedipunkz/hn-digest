---
source: "https://locaible.com/cursor"
hn_url: "https://news.ycombinator.com/item?id=48475954"
title: "Show HN: Locaible – local AI agents for Cursor with pre-tuned marketplace/commu"
article_title: "Use Cursor & Continue.dev with local Ollama agents — Locaible"
author: "locaible"
captured_at: "2026-06-10T16:24:01Z"
capture_tool: "hn-digest"
hn_id: 48475954
score: 2
comments: 0
posted_at: "2026-06-10T13:24:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Locaible – local AI agents for Cursor with pre-tuned marketplace/commu

- HN: [48475954](https://news.ycombinator.com/item?id=48475954)
- Source: [locaible.com](https://locaible.com/cursor)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T13:24:09Z

## Translation

タイトル: Show HN: Locaible – 事前調整されたマーケットプレイス/コミュを備えた Cursor 用のローカル AI エージェント
記事のタイトル: ローカルの Ollama エージェントで Cursor & Continue.dev を使用する — Locaible
説明: 127.0.0.1:11435 の Locaible のローカル OpenAI 互換 API にカーソルまたは Continue.dev をポイントします。事前に調整されたコーディング エージェント、プロジェクトごとのコンテキスト、トークン請求ゼロを利用できます。コードがマシンから離れることはありません。

記事本文:
ローカル Ollama エージェントで Cursor & Continue.dev を使用する — 位置特定可能な LOC AI BLE 機能 価格 セキュリティ コミュニティ ダウンロード サインイン 無料で始める Cursor または Continue.dev を使用している開発者向け Cursor & Continue.dev 内でローカル AI エージェントを実行する
Locaible は、OpenAI 互換 API を 127.0.0.1:11435 で公開します。 Cursor または Continue.dev をそれにポイントすると、マーケットプレイスから事前に調整されたコーディング エージェントを使用して AI が完全にマシン上で実行されます。
セットアップにジャンプ ↓ コードがマシンから離れることはありません トークン請求ゼロ クライアント リポジトリに対して GDPR セーフ 5 分間のウォークスルー — エンドツーエンドでカーソルの後ろに位置確認を参照してください。
Atlas (法的)、Code Reviewer (OWASP 対応)、Tech Support (L2/L3 Runbook) が、カーソルのドロップダウンに「モデル」として表示されます。それぞれが独自のシステム プロンプト + RAG を持ちます。
厳格なデータ処理ルールを持つ 1 つのクライアント リポジトリですか?あなたの規約を理解しているエージェントを立ち上げ、スタイルガイドとアーチドキュメントをフィードします。カーソルモデルメニューから切り替えます。
Cursor Pro はスニペットを Anthropic/OpenAI に送信します。 Locaible が背後にあると、チャット トラフィックは 127.0.0.1 で停止します。 CISO および DPO 監査に対して防御可能。
Cursor と Continue.dev についても同じフローです。エンドポイントを置き換えれば完了です。
1 Locaible をインストールし、コーディング エージェントを作成する
ハードウェアに適合するモデルを選択してください。高速なイテレーションの場合は qwen3:8b、本格的なリファクタリングの場合は qwen3:14b、48 GB リグの場合は deepseek-coder:33b です。 Locaible は、初回起動時に Ollama をインストールします。
2 Locaible でローカル OpenAI API を起動します
設定→ローカルAPI→開始。 Locaible は、OpenAI 仕様を実行するサーバーを 127.0.0.1:11435 にバインドし、エージェントを仮想モデルとして公開します。
3 エンドポイントにカーソル (または続行) をポイントします
OpenAI API キーをオンに切り替えます。空でない文字列をキーとして貼り付けます。
「OpenAI ベース URL をオーバーライド」をクリックし、http://localhost:11435/v1 を貼り付けます。
チャットのモデル ドロップダウンに次のように入力します。

アトラス (またはエージェント名) を選択し、それを選択します。
Continue VS Code 拡張機能をインストールします。
~/. continue/config.json を編集 — 以下を追加します。
{
「モデル」: [
{
"title": "Locaible・Atlas",
"プロバイダー": "openai",
"モデル": "アトラス",
"apiKey": "いいえ",
"apiBase": "http://localhost:11435/v1"
}
】
Continue は、Ollama を介したローカルのタブ オートコンプリートもサポートしています。カーソルのオートコンプリートはクラウド上に残ります。
切り替え後にギャップが見つからないように、残忍な正直さ。
ローカル ローカル インライン編集 (Cmd+K スタイル) ローカル ローカル タブ オートコンプリート ⚠ クラウド ローカル コードベース インデックス作成 (フル リポジトリ) ⚠ クラウド ローカル エージェント / 自律編集 ⚠ 限定的 ⚠ 限定的 Locaible エージェント間の切り替え ローカル ローカル コードがラップトップから離れることはありません ⚠ 部分的ローカル 結論: Cursor の UX を維持し、大量のチャット/編集トラフィックのみをマシンにオフロードしたい場合は、Cursor + Locaible が正しい呼び出しです。完全ローカルが厳しい制約である場合 (規制された産業、防衛、金融)、Continue.dev + Locaible がパスです。
プロのヒント: プロジェクトごとにカーソルを Locaible にロックする
リポジトリのルートに .cursorrules ファイルをドロップして、Cursor (および将来の自分) にどの Locaible エージェントを使用するかを指示します。クライアントの作業に特に役立ちます。クライアントのコードを誤って OpenAI に送信することができません。
# .cursorrules — リポジトリにコミット
# このコードベースに対するすべての AI 支援はローカルを経由する必要があります
# 位置特定可能なエンドポイント。選択したモデル: atlas-acmecorp.
#
# エンドポイント: http://localhost:11435/v1/chat/completions
# 理由: クライアントの NDA + GDPR 条項。 28 サードパーティ製プロセッサを禁止
# ソースコード用。 Locaible は開発者のマシン上で 100% 実行されます。自分のマシン上で Cursor を実行する
無料利用枠は、3 つのエージェントと 1 つのデバイスをカバーします。このページのすべての手順を試すには十分です。
Continue.dev を入手 LOC AI BLE 機能 価格 セキュリティ コミュニティ ダウンロード © 2026 Locaible.すべての人のための主権 AI。

## Original Extract

Point Cursor or Continue.dev at Locaible's local OpenAI-compatible API on 127.0.0.1:11435. Get pre-tuned coding agents, per-project context and zero token bills — your code never leaves your machine.

Use Cursor & Continue.dev with local Ollama agents — Locaible LOC AI BLE Features Pricing Security Community Download Sign in Start free For developers using Cursor or Continue.dev Run local AI agents inside Cursor & Continue.dev
Locaible exposes an OpenAI-compatible API on 127.0.0.1:11435. Point Cursor or Continue.dev at it and your AI runs entirely on your machine, with pre-tuned coding agents from the Marketplace.
Jump to setup ↓ Your code never leaves your machine Zero token bill GDPR-safe for client repos 5-minute walkthrough — see Locaible behind Cursor end-to-end.
Atlas (legal), Code Reviewer (OWASP-aware), Tech Support (L2/L3 runbooks) appear as 'models' in Cursor's dropdown. Each carries its own system prompt + RAG.
One client repo with strict data-handling rules? Spin up an agent that knows your conventions and feed it your style guide + arch docs. Switch via the Cursor model menu.
Cursor Pro sends snippets to Anthropic/OpenAI. With Locaible behind it, the chat traffic stops at 127.0.0.1. Defensible to your CISO and any DPO audit.
Same flow for Cursor and Continue.dev. Substitute your endpoint and you're done.
1 Install Locaible and create a coding agent
Pick a model that fits your hardware — qwen3:8b for fast iteration, qwen3:14b for serious refactors, deepseek-coder:33b if you have a 48 GB rig. Locaible installs Ollama for you on first launch.
2 Start the local OpenAI API in Locaible
Settings → Local API → Start. Locaible binds a server on 127.0.0.1:11435 that speaks the OpenAI spec and exposes your agents as virtual models.
3 Point Cursor (or Continue) at the endpoint
Toggle OpenAI API Key on. Paste any non-empty string as key.
Click Override OpenAI Base URL , paste: http://localhost:11435/v1
In the model dropdown of the chat: type atlas (or your agent name) and select it.
Install Continue VS Code extension.
Edit ~/.continue/config.json — add:
{
"models": [
{
"title": "Locaible · Atlas",
"provider": "openai",
"model": "atlas",
"apiKey": "noop",
"apiBase": "http://localhost:11435/v1"
}
]
} Continue also supports local Tab autocomplete via Ollama — Cursor's autocomplete stays on their cloud.
Brutal honesty so you don't discover the gaps after switching.
local local Inline edit (Cmd+K style) local local Tab autocomplete ⚠ cloud local Codebase indexing (full repo) ⚠ cloud local Agent / autonomous edits ⚠ limited ⚠ limited Switch between Locaible agents local local Your code never leaves the laptop ⚠ partial local Bottom line: Cursor + Locaible is the right call if you want to keep Cursor's UX and only offload the heavy chat/edit traffic to your machine. If full-local is a hard constraint (regulated industry, defence, finance), Continue.dev + Locaible is the path.
Pro tip: lock Cursor to Locaible per project
Drop a .cursorrules file at your repo root telling Cursor (and your future self) which Locaible agent to use. Especially useful for client work — you can't accidentally send their code to OpenAI.
# .cursorrules — committed to the repo
# All AI assistance for this codebase MUST go through the local
# Locaible endpoint. Selected model: atlas-acmecorp.
#
# Endpoint: http://localhost:11435/v1/chat/completions
# Reason: client NDA + GDPR Art. 28 forbid third-party processors
# for source code. Locaible runs 100% on the dev's machine. Get your Cursor running on your own machine
Free tier covers 3 agents and 1 device. Enough to try every step on this page.
Get Continue.dev LOC AI BLE Features Pricing Security Community Download © 2026 Locaible. Sovereign AI for everyone.
