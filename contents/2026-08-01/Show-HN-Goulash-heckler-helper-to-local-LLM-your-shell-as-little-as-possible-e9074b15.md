---
source: "https://github.com/chaboud/goulash"
hn_url: "https://news.ycombinator.com/item?id=49135861"
title: "Show HN: Goulash – heckler/helper to local-LLM your shell as little as possible"
article_title: "GitHub - chaboud/goulash: Add a little LLM to your shell · GitHub"
author: "chaboud"
captured_at: "2026-08-01T16:49:12Z"
capture_tool: "hn-digest"
hn_id: 49135861
score: 1
comments: 0
posted_at: "2026-08-01T16:33:52Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Goulash – heckler/helper to local-LLM your shell as little as possible

- HN: [49135861](https://news.ycombinator.com/item?id=49135861)
- Source: [github.com](https://github.com/chaboud/goulash)
- Score: 1
- Comments: 0
- Posted: 2026-08-01T16:33:52Z

## Translation

タイトル: HN: Goulash を表示 – シェルのローカル LLM をできるだけ少なくするヤジ/ヘルパー
記事のタイトル: GitHub - Chaboud/goulash: シェルに LLM を追加する · GitHub
説明: シェルに LLM を少し追加します。 GitHub でアカウントを作成して、シャブー/グーラッシュの開発に貢献してください。
HN テキスト: 私はターミナルをよく使いますが、利用可能なコマンド、フラグ、構文をすべて覚えているわけではありません。 ffmpeg、awk へのパイプ、find など...私が十分にクールではないことはわかっていますが、これはクレイジーな作業です。私は、Claude、ChatGPT、またはローカル モデルに、どのようにコマンドを書くか、ファイルやプロンプトで何かを行うかを尋ね、その後、mad-lib コマンドを使用してターミナルに戻り、それを推論するという通常のパターンに陥ります。そこでグーラッシュを作りました。これは、セッションについてコメントし、下矢印を押して「将来の履歴」を選択することで選択できるコマンドを提案する場所としてターミナルの下部 4 行を占め、ほとんどのシェルでは履歴の上矢印で流れを維持します。直接コマンドで質問してみませんか? 「# your question」と入力して Enter キーを押すだけで、そのインライン コメントが取得されて提案が作成されます。提案されたコマンドを実行するかどうかは常にユーザー次第です。下矢印を押して候補をスクロールし、必要に応じて変更し、Enter を押して実行します。これは、Macbook Air 上の gemma4:e4b を備えた Ollama や LM Studio など、ローカルに適したモデルをホストするローカル エンジンで実行するために作成されています。高速な動作を維持するためにコンテキスト キャッシュに大きく依存していますが、Goulash は思考や実行中にブロックしません。助手席に同乗です。あなたの小さなモデルが知らない微調整されたコマンド セットがありますか? 「#@ YourCommandReference.md」を使用して、gemma4:e4b (またはその他) にセッションのピンとしてガイドを与えるか、一時的に生活のスタイルを設定するために使用します (たとえば、ヒエのジョークを追加する FarmAnimals.md ファイルがあります)

任意のコメント。おい...トークンはローカルだ...)。主に Fable と Opus 5 を使用して Rust で書かれています。Mac 上の zsh と bash でテストされました。リポジトリをチェックアウトするか、「cargo install goulash」でインストールしてください。

記事本文:
GitHub - Chaboud/goulash: シェルに小さな LLM を追加する · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
シャブード
/
グーラッシュ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
本支店

s タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
193 コミット 193 コミット .github/ workflows .github/ workflows Formula Formula bench bench docs docs scripts scripts shell シェル src src テスト テスト wiki wiki .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
goulash - シェルの平易な言語ナビゲーター
下矢印をクリックすると、今後の履歴が表示されます。
(もしあなたが生のカールを手動で正規表現してこのページを読んだ末期の灰色ひげなら、今すぐやめて不死の状態に戻ってください。残りの私たちの場合は...)
すべての端末コマンドの構文を常に覚えているわけではありませんし、覚えておく必要はありません。操作をつなぎ合わせようとすると、見慣れたループが発生します。
コマンドを入力し始める → 難解な構文を忘れる → Google/ChatGPT/Claude に行って質問する → メール/Slack/Hacker-News に気づく → いくつかの論文/リンク/スレッドを読む → 昼食をとる → 戻ってきて、この端末で何が起こっているのか困惑するプロンプトを表示する → 状況を再理解する → コマンドを入力し始める → ...
コンテキストの切り替えにはコストがかかります。ブラウザまたはエージェントにアクセスすると、作業フローが中断されます。
しかし、大規模言語モデル (LLM) は、たとえ小規模で大規模な言語モデルであっても、難解な構文に非常に優れています... Brainwave!カチン！ (注: このソフトウェアは無料です。)
Goulash は、すでに使用しているシェル用の LLM 対応オーバーレイです。セッションを監視し、ターミナルにアドバイスや実行可能ファイルの提案を提供しますが、すべてのキーストロークとすべてのコマンドは依然としてあなたのものです。 Goulash はコマンドを実行したり、選択肢をポップアップしたり、ユーザーの作業方法を引き継いだりしません。それはナビゲーターを備えたあなたの端末です。
# 処理にはエンジンが必要です。たとえば、gemma4:e4b を使用した Ollama はうまく機能します。

ル
カーゴインストール goulash # from crates.io
# または、Homebrew を使用して
醸造タップ シャブード/グーラッシュ https://github.com/chaboud/goulash
醸造インストールグーラッシュ
使用法
goulash を実行し、以前とまったく同じようにシェルを使用します。 Goulash はプロンプトの下に提案を表示します。あなたの未来の歴史。気に入ったものが見つかったら、下矢印を押してください。
# 質問して、答えを取得し、プル可能なコマンドを取得します
プロンプトで # と質問を入力します (これは履歴に記録されるシェル コメントであり、実行されることはありません)。
LLM は推奨コマンドを提供します。 (シェル履歴の終わりを超えて) を押すと、編集用のプロンプトが表示されます。 Enter を押すと、独自のコマンドとして実行されます。
グーラッシュはまた、あなたのシェルの使用状況をレビューし、同じ場所にチップを残す場合があります。ローカル モデルを使用すると、端末のコマンド履歴と同様に、すべてプライベートですべてローカルになります。
スタイル - 矢印: 1 つの空間軸
シェル履歴はプロンプトの上に保存されます。グーラッシュの提案は以下にあります。上下は線に沿って移動するだけです。
↑ zsh の歴史 (いつものように)
── 空のプロンプト ────── ニュートラル
↓ 最新の推奨コマンド (↓ 再度: 古い · ↑: バックアップ)
最新を超えて、ルールの右側に表示される位置 ( ↑ 3/7 ↓ ) で推奨コマンドの履歴が表示されます。空のプロンプトまで遡り、それを過ぎると、単純な zsh/bash 履歴になります。
## (または ## question ) はチャット パネルに切り替わります。フォローアップには接頭辞は必要ありません。シェルは上で実行され続けます。プロンプトにコピーするコマンドを選択すると、シェルに戻ります。
モデルにファイルを指定すると、モデルは一般的な Unix だけでなく、ツールを認識します。
#@/path commandRef.md ファイル (またはディレクトリ) を固定します — LLM は関係しません
#@ Synology ドキュメントを使用して言葉で言います。グーラッシュが見つけて固定します
#@ 固定されているもの、大きさ、新鮮さ
#@/unset ドロップ
ビッグフィル

es は切り捨てられません。予算をオーバーフローしたピンは、
構造のみのアウトラインをすぐに作成 (コマンド、フラグ、テーブル - 散文)
ドロップされました)、エンジンが稼働している場合は、LLM 圧縮がクックされます。
背景とその後ろのスワップ。クロムではパーセンテージが表示されますが、
それは実行します。 #@/cancel で停止します。取り込み時に何も待機することはありません。
固定されたテキストはすべての質問に反映され、クロムには固定されている内容が表示されます
( @commandRef.md 、ディスク上で変更された場合は * が付きます。グーラッシュはそれをマークします。
決して黙って再読しないでください)。ベンダーのコマンド リファレンスをその横にドロップします。
CLI と提案は、モデルがこれまでになかったツールに適したものになり始めます
見た。それはまだ示唆するだけです。あなたはまだそれを実行しています。
#@/path … はプレーンなコマンドであるため、モデルはコマンドを提案することができます。
あなた — CMD: #@/path commandRef.md は、通常のプル可能なチップとして到着します。
#/model モーダル モデル ピッカー: フィルターに入力、⏎ 選択して保存
#/model NAME このセッションのモデルを試します (永続化するには `save` を追加します)
#/memory モデルに小さな固定メモリを与えます (REMEMBER/FORGET)
#/memory スロットを参照します: フィルター、↑↓、⏎⏎ を使用してスロットを忘れます
#/思考の推論レベルが低い: オフ |低い |中 |高い
#/settings すべてをライブ調整し、その場で適用して保存します
#/オタクなものをデバッグします (おそらくこれらは必要ありません)
#/ターンごとのヤジを静かにコメントオフ
#/ステータス
#/ヘルプ
注意事項
プラットフォーム。 macOS のターミナルで開発され、毎日使用されています。
ズシュ。一部の Linux と SSH 経由の bash。
エンジン。ほとんどが ollam、かなりの量の LM Studio、まだ何もありません
有料のホステッドプロバイダー。 OpenAI 互換のワイヤーが存在し、機能します
llama.cpp および vLLM に対しては実行しますが、セッションは自分のマシン上に維持します
それがポイントです。
まだ動いています。構成キー、設定名、操作の詳細
リリース間で変更します。 CHANGELOG には、何が動いたのか、何が設定されたのかが記載されています

それは
移動可能な場合は既存の config.toml から作業を継続します。
しかし、これはまだ構築するための安定した表面ではありません。
カーゴビルド
./target/debug/goulash # $SHELL をラップします
./target/debug/goulash zsh # またはシェルの名前を指定します
シェルの統合は、プレーンなフラグで起動された zsh および bash では自動的に行われます。goulash は、通常の rc ファイルの上にそのフック (ZDOTDIR トリック / --rcfile ラッパー) を挿入するため、編集は必要ありません。それがコマンドブロック、# 脇、そして単純なダウン提案のプルを強化するものです。
手動フォールバック (カスタム シェルまたは auto_integrate = false ):
# ~/.zshrc
[[ -n " $GOULASH " ]] && ソース /path/to/goulash/shell/goulash.zsh
# ~/.bashrc
[[ -n " $GOULASH " ]] && ソース /path/to/goulash/shell/goulash.bash
設定 (オプション): ~/.goulash/config.toml
【状態】
有効 = true
行 = 1
【エンジン】
プロバイダー = " 自動 " # 自動 |オラマ |オープンナイ |オープンナイ生 |なし
# 「openai」は LM Studio、llama.cpp、vLLM および
# ホストされた /v1 エンドポイント; openai-raw はスキップします
# サーバーのチャットテンプレートであり、測定用です
思考 = "オフ" #オフ |低い |中 |高い
max_tokens = 8192 # 推論と回答の 1 つの上限
遅い = " 手動 " # 遅い場合はボランティア: 手動 |クエリ |ウォルドルフ
# (`#?` とピンは常に各段に到達します)
【エンジン。暴露する]
platform = true # モデルに OS、シェル、ユーザーランドを伝えます
# 研究レーン。ここでのすべてのキーはオーバーライドです。 1つを残して、
# その設定は高速レーンに従います - 存在せず、凍結されたコピーではないため、
# 高速デフォルトを改善すると、これが改善されます。モデルのネーミングは、
# 実際にレーンを 2 つのバインディングに分割するもの。
【エンジン。スローレーン]
# Provider = "openai" # 研究目的のみに大物を呼び込む
# モデル = "gpt-oss:20b"
Thinking = " Medium " # デフォルト: 考える余裕のあるレーン
# より新しいモデルのエスケープハッチ

グーラッシュの能力表。
【モデル。 "some-new-reasoner:8b" ]
Thinking = " レベル " # なし |ブール |レベル
推論トークン = 2048
または、ssh 経由およびスクリプト内で動作するコマンド ラインから:
goulash --config print # すべてのキーとそれがあなたのものであるかどうかを表示します
goulash --config エンジンを設定します。高く考えています
goulash --config リセットエンジン. Thinking
個別に考える予算はありません。プロバイダーのメーター推論と
出力は 1 つのカウンターにあり、推論は私たちが行うものではありません - チャット
テンプレートは送信したものをすべて推論し、一部のモデルは推論します。
考える: false 。したがって、max_tokens は 1 つの寛大な上限であり、
エンジンが内部で行うことはエンジンの仕事です。答えは短く残る
プロンプトで 1 行を要求されるためであり、予算が不足しているためではありません。
それら: 測定され、到着した回答は中央値 32 トークンを使用します。
グーラッシュはエンジンを邪魔しないように努めます。コンテキスト サイズは送信されません
ロードされたものが小さすぎて作業できない場合を除き、ollama と LM Studio の両方
これを自分で設定でき、変更すると数秒間のリロードが強制されます。
テスト (実際の PTY でバイナリを駆動します):
カーゴビルド && python3 テスト/e2e.py
ライセンス
オプションで MIT または Apache-2.0 に基づくデュアル ライセンス。あなたが明示的に別段の定めをしない限り、Apache-2.0 ライセンスで定義されているように、グーラッシュに含めるためにあなたが意図的に提出した投稿は、追加の条項や条件なしで上記のように二重ライセンスされるものとします。
シェルに LLM を少し追加します
Readme Apache-2.0、MIT ライセンスが見つかりました。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Add a little LLM to your shell. Contribute to chaboud/goulash development by creating an account on GitHub.

I use the terminal a lot, but I don't remember every command, flag, and syntax available. ffmpeg, piping to awk, find, etc... I know I'm not cool enough, but it's crazy-making. I get into a regular pattern of asking Claude, ChatGPT, or a local model how I'd write a command or do something with some files or a prompt, then going back to the terminal with a mad-lib command and reasoning through it. So I built goulash. It takes up the bottom four lines of your terminal as a place to comment on your session and suggest commands that you can select by hitting down arrow to select your " future history ", staying in flow with up arrow for history in most shells. Want to ask it a direct command question? Just type "# your question" and hit enter, and it will take that inline comment and make a suggestion. Running a suggested command is always up to the user. Hit down arrow to scroll through suggestions, modify if you want, hit enter to execute. This is made for running with local engines hosting local-friendly models, like Ollama or LM Studio with gemma4:e4b on a Macbook Air. It relies heavily on context caching to stay snappy, but goulash doesn't block as it thinks or runs. It's a ride-along in the passenger seat. Have a tweaky command set that your tiny model doesn't know? Use "#@ YourCommandReference.md" to give that gemma4:e4b (or whatever) a guide as a pin for the session, or use it to temporarily style your life (e.g., I have a FarmAnimals.md file that adds a barnyard joke to any comment. Hey... the tokens are local...). Written in rust, largely with Fable and Opus 5. Tested with zsh and bash on Mac. Check out the repo, or install with "cargo install goulash".

GitHub - chaboud/goulash: Add a little LLM to your shell · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
chaboud
/
goulash
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
193 Commits 193 Commits .github/ workflows .github/ workflows Formula Formula bench bench docs docs scripts scripts shell shell src src tests tests wiki wiki .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md View all files Repository files navigation
goulash - a plain-language navigator for your shell
Just Down-arrow for your future history.
(If you are a terminal graybeard who hand-regexed raw curl to read this page, you can stop now and go back to being immortal. For the rest of us...)
We don't always remember the syntax for all terminal commands, and we shouldn't have to. When we try to string together operations, there's a familiar loop:
Start typing command → forget esoteric syntax → go to Google/ChatGPT/Claude to ask → notice email/slack/Hacker-News → read a few papers/links/threads → get lunch → come back to prompt confused about what was happening in this terminal → re-understand the situation → start typing command → ...
Context switching has a cost. Going to a browser or an agent breaks working flow.
But large language models (LLMs) are pretty good at esoteric syntax , even small large language models... brainwave! ka-ching! (note: this software is free.)
Goulash is an LLM-aware overlay for the shell you already use. It watches the session and offers advice and executable suggestions in your terminal , but every keystroke and every command is still yours. Goulash doesn't run commands, pop up choices, or take over how you work. It's your terminal... with a navigator.
# You need an engine for processing. Ollama with gemma4:e4b works well, for example
cargo install goulash # from crates.io
# or, with Homebrew
brew tap chaboud/goulash https://github.com/chaboud/goulash
brew install goulash
Usage
Run goulash and use your shell exactly as before. Goulash provides suggestions below your prompt; your future history . Just press the down arrow if you see something you like.
# ask a question, get an answer + a pullable command
Type # and a question at your prompt (it's a shell comment recorded in history, never executed).
The LLM will provide a suggested command. Press down (past the end of shell history) and it lands on your prompt for editing. Enter runs it as your own command.
Goulash also reviews your shell use as you go and may leave a tip in the same spot. Use a local model, and it's all private, all local, just like your terminal command history.
Style - Arrows: one spatial axis
Shell history lives above your prompt; goulash's suggestions live below. Up and Down just move along the line:
↑ zsh history (as always)
── your empty prompt ──────────────── neutral
↓ newest suggested command ( ↓ again: older · ↑: back up )
Down past the newest keeps going into the history of suggested commands with the position ( ↑ 3/7 ↓ ) shown at the right of the rule. Up retraces to your empty prompt, and past that it's plain zsh/bash history.
## (or ## question ) flips into a chat panel; follow-ups need no prefix. The shell keeps running above. When you select a command to copy to the prompt, you're back in your shell.
Point the model at a file and it knows your tools, not just common Unix:
#@/path commandRef.md pin a file (or a directory) — no LLM involved
#@ use the synology doc say it in words; goulash finds and pins it
#@ what's pinned, how big, how fresh
#@/unset drop it
Big files don't get truncated. A pin that overflows its budget serves a
structure-only outline immediately (commands, flags, tables — prose
dropped), and if an engine is up, an LLM compression cooks in the
background and swaps in behind it. The chrome shows a percentage while
that runs; #@/cancel stops it. Nothing ever waits on an ingest.
The pinned text rides in every ask, and the chrome shows what's anchored
( @commandRef.md , with a * when it changed on disk — goulash marks it,
never silently re-reads). Drop a vendor's command reference next to their
CLI and suggestions start coming out right for a tool the model has never
seen. It still only ever suggests ; you still run it.
Because #@/path … is a plain command, the model can suggest one back at
you — CMD: #@/path commandRef.md arrives as a normal pullable chip.
#/model modal model picker: type to filter, ⏎ selects & saves
#/model NAME try a model for this session (add `save` to persist)
#/memory on give the model a small pinned memory (REMEMBER/FORGET)
#/memory browse the slots: filter, ↑↓, ⏎⏎ to forget one
#/thinking low reasoning level: off | low | medium | high
#/settings live-tune everything, applied and saved on the spot
#/debug nerd stuff (you probably don't need these)
#/commentary off quiet the per-turn heckling
#/status
#/help
Caveats
Platforms. Developed and used daily on macOS, in Terminal, under
zsh. Some Linux and bash over ssh.
Engines. Mostly ollama, a fair amount of LM Studio, nothing yet with
a paid hosted provider. The OpenAI-compatible wire is there and works
against llama.cpp and vLLM, but keeping your session on your own machine
is the, like... the point.
Still moving. Config keys, setting names and interaction details
change between releases. The CHANGELOG says what moved and settings that
move keep working from an existing config.toml where that is possible,
but this is not yet a stable surface to build on.
cargo build
./target/debug/goulash # wraps $SHELL
./target/debug/goulash zsh # or name a shell
Shell integration is automatic for zsh and bash launched with plain flags — goulash injects its hooks (ZDOTDIR trick / --rcfile wrapper) on top of your normal rc files, no editing required. That's what powers command blocks, # asides, and the plain-Down suggestion pull.
Manual fallback (custom shells or auto_integrate = false ):
# ~/.zshrc
[[ -n " $GOULASH " ]] && source /path/to/goulash/shell/goulash.zsh
# ~/.bashrc
[[ -n " $GOULASH " ]] && source /path/to/goulash/shell/goulash.bash
Config (optional): ~/.goulash/config.toml
[ status ]
enabled = true
rows = 1
[ engine ]
provider = " auto " # auto | ollama | openai | openai-raw | none
# "openai" covers LM Studio, llama.cpp, vLLM and
# hosted /v1 endpoints; openai-raw skips the
# server's chat template and is for measurement
thinking = " off " # off | low | medium | high
max_tokens = 8192 # ONE cap over reasoning and answer together
slow = " manual " # when slow volunteers: manual | query | waldorf
# (`#?` and pins always reach it, at every rung)
[ engine . divulge ]
platform = true # tell the model your OS, shell and userland
# The research lane. Every key here is an override; leave one out and
# that setting follows the fast lane — absent, not a frozen copy, so
# improving the fast default improves this with it. Naming a model is
# what actually splits the lanes onto two bindings.
[ engine . slow_lane ]
# provider = "openai" # call in the big guns for research only
# model = "gpt-oss:20b"
thinking = " medium " # the default: the lane that can afford to think
# Escape hatch for a model newer than goulash's capability table.
[ models . "some-new-reasoner:8b" ]
thinking = " levels " # none | bool | levels
reasoning_tokens = 2048
Or from the command line, which works over ssh and in scripts:
goulash --config print # every key, and whether it is yours
goulash --config set engine.thinking high
goulash --config reset engine.thinking
There is no separate thinking budget. Providers meter reasoning and
output on one counter, and reasoning is not ours to ration — a chat
template reasons whatever we send, and some models reason through
think:false . So max_tokens is one generous ceiling and whatever the
engine does inside it is the engine's business. Answers stay short
because the prompt asks for one line, not because the budget starves
them: measured, answers that arrive use a median of 32 tokens.
Goulash tries not to disturb your engine. It sends no context size
unless what is loaded is too small to work in — both ollama and LM Studio
let you set that yourself, and changing it forces a multi-second reload.
Tests (drives the binary under a real PTY):
cargo build && python3 tests/e2e.py
License
Dual-licensed under MIT or Apache-2.0 , at your option. Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in goulash by you, as defined in the Apache-2.0 license, shall be dual licensed as above, without any additional terms or conditions.
Add a little LLM to your shell
Readme Apache-2.0, MIT licenses found Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
