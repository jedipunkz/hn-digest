---
source: "https://reachpad.dev/blog/free-models"
hn_url: "https://news.ycombinator.com/item?id=49291935"
title: "Reachpad now includes free Gemma 4 31B and Gemini 3.6 Flash"
article_title: "Free coding models on every pad · reachpad"
author: "sakuraiben"
captured_at: "2026-08-13T21:35:13Z"
capture_tool: "hn-digest"
hn_id: 49291935
score: 1
comments: 0
posted_at: "2026-08-13T21:12:03Z"
tags:
  - hacker-news
  - translated
---

# Reachpad now includes free Gemma 4 31B and Gemini 3.6 Flash

- HN: [49291935](https://news.ycombinator.com/item?id=49291935)
- Source: [reachpad.dev](https://reachpad.dev/blog/free-models)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T21:12:03Z

## Translation

タイトル: Reachpad に無料の Gemma 4 31B および Gemini 3.6 Flash が含まれるようになりました
記事タイトル: すべてのパッド・リーチパッドの無料コーディング モデル
説明: Gemma 4 31B および Gemini 3.6 Flash は、OpenCode を通じてすべてのリーチパッドに含まれており、プロバイダー アカウントや API キーを構成する必要はありません。

記事本文:
すべてのパッドの無料コーディング モデル · リーチパッド 無料モデル Gemma 4 31B + Gemini 3.6 すべてのパッドのフラッシュ → リーチパッド CLI ドキュメント 価格設定 ブログ お問い合わせ ログイン トライアル ブログを開始
すべてのパッドで無料のコーディング モデル
すべてのリーチパッドには、API キーを持ち込まずに OpenCode を通じて使用できる 2 つのモデル、Gemma 4 31B と Gemini 3.6 Flash が含まれるようになりました。モデルの使用法はパッドに含まれています。
OpenCode は両方のモデルで事前構成されています。それを開いて、リーチパッド プロバイダーを選択し、プロンプトを送信します。 Google アカウントに接続したり、モデルの請求書を設定したりする必要はありません。
Gemma 4 31B は、OpenCode のデフォルトのリーチパッド モデルです。
Gemini 3.6 フラッシュは同じモデル ピッカーから入手でき、同じ付属の許容量を使用します。
Claude Code と Codex は引き続きインストールされ、最新の状態に保たれます。これらのツールは、以前と同様に自分のアカウントを使用します。付属の 2 つのモデルは別のオプションであり、すでに機能しているエージェント セットアップの代替品ではありません。
最初のタスクがプロバイダー アカウントをプロビジョニングし、そのキーをシェルに入力する場合、エージェントは開発環境の準備ができていません。プロバイダーを決定した場合、この設定は合理的です。なじみのないリポジトリをモデルに渡して、その動作を確認したい場合、最初の 5 分間は不十分です。
付属のモデルはパッドの最初のプロンプト部分を構成します。リポジトリのクローンを作成し、OpenCode を開きます。コードの説明や変更を依頼してください。 1 つのモデルが停止した場合は、認証情報を変更したり、作業ツリーを移動したりせずに、もう 1 つのモデルに切り替えます。
プロバイダー キーがパッドから離れた状態を保つ方法
共有アップストリームキーをお客様のマシンにコピーすることはありません。各パッドは、その所有者を対象とした取り消し可能な資格情報を受け取ります。 OpenCode はその資格情報をリーチパッド モデル ゲートウェイに提示し、リクエストの中継中にのみ許可をチェックしてアップストリーム キーを追加します。
ゲートウ

ay は、プロンプトや応答を解析する代わりに、パイプでバイトを計測します。これにより、会話を読み取るアプリケーション コードを追加しなくても、暴走エージェント ループの 1 日あたりの上限が得られます。アップストリーム キーはゲートウェイに残り、パッドを再構築せずに資格情報を取り消すことができます。
これら 2 つのモデルにはトークンごとの料金はかかりません。使用はすべてのパッドの一部であり、別のモデルのサブスクリプションは必要ありません。 1 日あたりの許容量は UTC 午前 0 時にリセットされるため、1 つの壊れたループが無制限の量の推論を消費することはできません。 OpenCode は、パッドがその上限に達すると、クォータ応答を報告します。
含まれているモデルは現在テキストを受け入れます。タスクに別のモデルまたは入力タイプが必要な場合は、独自のプロバイダー アカウントを使用します。
2 つのモデルは、新規および既存のパッドで使用できます。パッドとワークスペースのトライアルの詳細を参照するか、アカウントを作成して開始してください。
コーディングエージェントのための永続的な開発環境。

## Original Extract

Gemma 4 31B and Gemini 3.6 Flash are included with every reachpad through OpenCode, with no provider account or API key to configure.

Free coding models on every pad · reachpad Free models Gemma 4 31B + Gemini 3.6 Flash on every pad → reachpad CLI Docs Pricing Blog Contact Log in Start trial blog
Free coding models on every pad
Every reachpad now includes two models you can use through OpenCode without bringing an API key: Gemma 4 31B and Gemini 3.6 Flash. Model usage is included with the pad.
OpenCode comes preconfigured with both models. Open it, choose the reachpad provider, and send a prompt. There is no Google account to connect and no model bill to set up.
Gemma 4 31B is the default reachpad model in OpenCode.
Gemini 3.6 Flash is available from the same model picker and uses the same included allowance.
Claude Code and Codex are still installed and kept current. Those tools use your own accounts, as they did before. The two included models are another option, not a replacement for an agent setup that already works for you.
A development environment is not ready for an agent if the first task is provisioning a provider account and putting its key into a shell. That setup is reasonable when you have settled on a provider. It is a poor first five minutes when you want to hand an unfamiliar repository to a model and see what it does.
The included models make the first prompt part of the pad. Clone a repository and open OpenCode. Ask it to explain or change the code. If one model gets stuck, switch to the other without changing credentials or moving the working tree.
How the provider key stays off the pad
We do not copy a shared upstream key onto customer machines. Each pad receives a revocable credential scoped to its owner. OpenCode presents that credential to the reachpad model gateway, which checks the allowance and adds the upstream key only while relaying the request.
The gateway meters bytes at the pipe instead of parsing prompts or responses. That gives us a daily ceiling for a runaway agent loop without adding application code that reads the conversation. The upstream key remains at the gateway, and a credential can be revoked without rebuilding the pad.
There is no per-token charge for these two models. Usage is part of every pad and does not require a separate model subscription. A daily allowance resets at midnight UTC, so one broken loop cannot consume an unbounded amount of inference. OpenCode reports a quota response if the pad reaches that ceiling.
The included models accept text today. Bring your own provider account when a task needs a different model or input type.
The two models are available on new and existing pads. See the pad and workspace trial details , or create an account to start.
Persistent development environments for coding agents.
