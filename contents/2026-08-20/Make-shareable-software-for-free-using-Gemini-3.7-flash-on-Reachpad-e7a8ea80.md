---
source: "https://reachpad.dev/blog/gemini-3-7-flash"
hn_url: "https://news.ycombinator.com/item?id=49368932"
title: "Make shareable software for free using Gemini 3.7 flash on Reachpad"
article_title: "Gemini 3.7 Flash offered for free · reachpad"
image: "https://reachpad.dev/blog/gemini-3-7-flash/opengraph-image?63b5690b3ad1a641"
author: "sakuraiben"
captured_at: "2026-08-20T00:38:52Z"
capture_tool: "hn-digest"
hn_id: 49368932
score: 1
comments: 0
posted_at: "2026-08-20T00:21:24Z"
tags:
  - hacker-news
  - translated
---

# Make shareable software for free using Gemini 3.7 flash on Reachpad

- HN: [49368932](https://news.ycombinator.com/item?id=49368932)
- Source: [reachpad.dev](https://reachpad.dev/blog/gemini-3-7-flash)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T00:21:24Z

## Translation

タイトル: Reachpad の Gemini 3.7 フラッシュを使用して共有可能なソフトウェアを無料で作成する
記事のタイトル: Gemini 3.7 Flash が無料で提供 · リーチパッド
説明: Gemini 3.7 フラッシュは、Gemma 4 31B および Gemini 3.6 フラッシュとともに、OpenCode を通じてすべてのリーチパッドに含まれています。得られるもの、実行方法、含まれる手当の範囲。

記事本文:
Gemini 3.7 フラッシュが無料で提供 · リーチパッド 無料モデル Gemma 4 31B + Gemini 3.7 すべてのパッドにフラッシュ → リーチパッド CLI ドキュメント 価格 ブログ お問い合わせ ログイン トライアル ブログを開始
Gemini 3.7 Flash が無料で提供される
Gemini 3.7 Flash がすべてのパッドに含まれるようになりました。 Gemma 4 31B および Gemini 3.6 Flash に加わります。OpenCode の 3 つのモデル、プロバイダー アカウントなし、貼り付けるキーなし、使用法はパッドに含まれています。
各モデルは OpenCode の独自のプロバイダー エントリであり、agent.reachpad.dev のリーチパッド モデル ゲートウェイから提供されます。ピッカーは 3 つすべてをパッド上に表示します。
インストールするものや貼り付けるものは何もありません。パッドはプロバイダーが構成され、独自の資格情報とともに届くため、シェルを開く前にモデルの準備ができています。パッドの機能を OpenCode に尋ねます。
$ opencode models モデルごとに 1 行、プロバイダーを最初に出力します。同じ ID で 3.7 に対してワンショットを実行します。
$ opencode run -m reversepad-gemini-3-7/gemini-3-7-flash "explain this repo" OpenCode TUI では、/model によってピッカーが開き、ステータス バーにどのプロバイダーが応答したかが表示されます。各モデルが個別のエントリであり、間違ったモデルも同じように流暢に応答するため、一見の価値があります。デフォルトを固定するには、 ~/.config/opencode/opencode.json でモデルを設定します。パッドはその選択肢を保持します。
Claude Code と Codex はすべてのパッドにインストールされ、独自のアカウントを使用し続けます。付属のモデルは別のオプションであり、すでに機能しているエージェント セットアップの代替品ではありません。
トークンごとの料金や個別のモデルのサブスクリプションはありません。使用量はパッドの一部であり、1 日あたりの許容量は UTC 午前 0 時にリセットされるため、1 つの暴走エージェント ループが無制限の量の推論を消費することはありません。 OpenCode は、パッドがその上限に達すると、クォータ応答を表示します。
プロバイダー キーがパッドにコピーされることはありません。各パッドは独自の取り消し可能な資格情報を取得し、reachpa

d ゲートウェイは、リクエストの中継中にのみアップストリーム キーを追加します。トラフィックを読み取るのではなく、トラフィックを測定します。許可を機能させるためにプロンプ​​トや応答を解析したり保存したりすることはありません。
含まれているモデルは現在テキストを受け入れます。タスクに別のモデルまたは入力タイプが必要な場合は、独自のプロバイダー アカウントを使用します。
3 つのモデルはすべて、新規および既存のパッドで使用でき、パッドは次のコンバージでそれらを選択します。パッドとワークスペースのトライアルの詳細を参照するか、アカウントを作成して開始してください。
コーディングエージェントのための永続的な開発環境。

## Original Extract

Gemini 3.7 Flash is included with every reachpad through OpenCode, alongside Gemma 4 31B and Gemini 3.6 Flash. What you get, how to run it, and what the included allowance covers.

Gemini 3.7 Flash offered for free · reachpad Free models Gemma 4 31B + Gemini 3.7 Flash on every pad → reachpad CLI Docs Pricing Blog Contact Log in Start trial blog
Gemini 3.7 Flash offered for free
Gemini 3.7 Flash is now included with every pad. It joins Gemma 4 31B and Gemini 3.6 Flash: three models in OpenCode, no provider account, no key to paste, usage included with the pad.
Each model is its own provider entry in OpenCode, served from the reachpad model gateway at agent.reachpad.dev . The picker shows all three on a pad.
There is nothing to install and nothing to paste. A pad arrives with the providers configured and a credential of its own, so the models are ready before you open a shell. Ask OpenCode what the pad has:
$ opencode models That prints one line per model, provider first. Run a one-shot against 3.7 with the same id:
$ opencode run -m reachpad-gemini-3-7/gemini-3-7-flash "explain this repo" In the OpenCode TUI, /model opens the picker and the status bar shows which provider answered — worth a glance, since each model is a separate entry and the wrong one answers just as fluently. To pin a default, set model in ~/.config/opencode/opencode.json ; a pad keeps that choice.
Claude Code and Codex are installed on every pad and keep using your own accounts. The included models are another option, not a replacement for an agent setup that already works for you.
There is no per-token charge and no separate model subscription. Usage is part of the pad, with a daily allowance that resets at midnight UTC so one runaway agent loop cannot consume an unbounded amount of inference. OpenCode surfaces a quota response if a pad reaches that ceiling.
No provider key is ever copied onto a pad. Each pad gets its own revocable credential, and the reachpad gateway adds the upstream key only while relaying the request. We meter the traffic rather than reading it: no prompt or response is parsed or stored to make the allowance work.
The included models accept text today. Bring your own provider account when a task needs a different model or input type.
All three models are available on new and existing pads, and a pad picks them up on its next converge. See the pad and workspace trial details , or create an account to start.
Persistent development environments for coding agents.
