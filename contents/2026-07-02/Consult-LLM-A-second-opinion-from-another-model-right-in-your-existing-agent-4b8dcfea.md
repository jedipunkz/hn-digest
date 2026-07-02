---
source: "https://github.com/raine/consult-llm"
hn_url: "https://news.ycombinator.com/item?id=48768213"
title: "Consult-LLM – A second opinion from another model right in your existing agent"
article_title: "GitHub - raine/consult-llm: Get a second opinion from another AI model · GitHub"
author: "tiahura"
captured_at: "2026-07-02T23:09:41Z"
capture_tool: "hn-digest"
hn_id: 48768213
score: 1
comments: 0
posted_at: "2026-07-02T22:31:24Z"
tags:
  - hacker-news
  - translated
---

# Consult-LLM – A second opinion from another model right in your existing agent

- HN: [48768213](https://news.ycombinator.com/item?id=48768213)
- Source: [github.com](https://github.com/raine/consult-llm)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T22:31:24Z

## Translation

タイトル: Consult-LLM – 既存のエージェントの別のモデルからのセカンドオピニオン
記事のタイトル: GitHub - Raine/consult-llm: 別の AI モデルからセカンドオピニオンを得る · GitHub
説明: 別の AI モデルからセカンドオピニオンを取得します。 GitHub でアカウントを作成して、raine/consult-llm の開発に貢献してください。

記事本文:
GitHub - Raine/consult-llm: 別の AI モデルからセカンドオピニオンを得る · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
レイン
/
相談-llm
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動

コード 「その他のアクション」メニューを開く フォルダーとファイル
842 コミット 842 コミット .github/ workflows .github/ workflows crates crates 例 例 フック フック メタ メタ スクリプト スクリプト スキル スキル src src .gitignore .gitignore .workmux.yaml .workmux.yaml CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.tomlライセンス ライセンス README.md README.md RELEASE.md RELEASE.md build.rs build.rs checkle.toml checkle.toml justfile justfile すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クイックスタート ·
使い方・
プロバイダーと構成 ·
スキル・
ロギング ·
モニター・
変更履歴
Consult-llm は、別の AI モデルからセカンドオピニオンを取得するためのツールです。
既存のエージェントのワークフロー内ですぐに実行できます。建築計画に使用し、
変更をレビューしたり、アプローチについて議論したり、厄介なバグから抜け出したりできます。 GPT-5.5、Gemini 3.1 Pro、Claude Opus 4.7、
DeepSeek V4 Pro、MiniMax M2.7、Grok 4.3、および GLM 5.2、API およびローカル CLI バックエンドを備えた任意の openrouter/* モデル、
マルチターン スレッド、git diff コンテキスト、Web モード クリップボード エクスポート、およびライブ モニター TUI。
モデル自身の作業をレビューすることは、独立したチェックではありません。フレッシュでも
コンテキストでは、同じトレーニング、事前経験、同じ失敗の多くを共有しています。
モード。異なるモデルは異なる方法でトレーニングされ、異なるモデルを作成します
間違いがあるため、反発したり、弱い推論に異議を唱えたり、
盲点を暴露する。
モデルの強度にはばらつきがあり、タスクごとに予測するのは困難です。簡単なクロスチェック
は安価であり、場合によっては他のモデルが活路を見出すこともあります。
エージェントにスキルをインストールする (Claude Code、Codex、OpenCode)
スラッシュコマンドでトリガー: /consult 、 /debate 、 /collab
スキルはプロンプトをConsult-llmにパイプし、バックエンドを呼び出して応答をインラインでストリーミングバックします。
より詳細なフローについては、以下に進みます。
>

修正後もヒントのロック解除は壊れたままです。ジェミニに質問する
⏺ スキル(相談-llm)
⎿ スキルが正常にロードされました · 1 つのツールが許可されます
⏺ このセグメンテーションの不一致を理解するために、Gemini に相談してみましょう。
⏺ Bash(cat issue.md |Consult-llm -m gemini -f HintsText.tsx -fhints.ts)
⎿ [モデル:gemini-3.1-pro-preview]
根本的な原因は、「handleconfirmGenerateHints」の競合状態です。
「await generatedHints()」の後、コンポーネントの状態が更新されます。
バックグラウンドですが、関数は古い `activeSegment` で続行します。
変数を使用して、「ではない」ではなく「ではない」というヒントを要求します。
⏺ Gemini は本当の問題を特定しました: 古い状態による競合状態
非同期操作後。完全に間違ったものを見ていました。
注記
このプロジェクトは以前、トランスポート層として MCP を使用していました。そのバージョンを使用していた場合は、以下の移行ガイドを参照してください。
関連するファイル コンテキスト ( /consult ) を使用して、エージェントから別のモデルからセカンドオピニオンを取得します。
モデルに最適なアプローチについて議論させ、コンセンサスを合成させます ( /debate )
API キーなしで CLI バックエンド経由で既存のサブスクリプションを使用する ( Gemini CLI 、 Codex CLI 、 Cursor CLI 、 Claude CLI 、 OpenCode )
thread_id を使用してリクエスト間で会話を継続する
ブラウザベースの LLM のプロンプトをクリップボードにコピー ( --web )
リアルタイム TUI でアクティブな実行と過去の実行を監視します
#自作
brew install Raine/consult-llm/consult-llm
# シェルスクリプト (macOS/Linux)
カール -fsSL https://raw.githubusercontent.com/raine/consult-llm/main/scripts/install.sh |バッシュ
バックエンドを構成します。既存のサブスクリプションがある場合は、CLI バックエンドが最も簡単です (API キーは必要ありません)。
Consult-llm config set gemini.backend gemini-cli # 必要なもの: Gemini ログイン
Consult-llm config set openai.backend codex-cli # 必要なもの: codex ログイン
または、API キーを設定します。
コンサルト-llm設定セットopenai.api_key your_openai_ke

y
コンサルト-llm構成セットgemini.api_key your_gemini_key
コンサルト-llm 設定 grok.api_key your_xai_key
コンサルト-llm設定セットzai.api_key your_zai_key
# または環境変数経由: OPENAI_API_KEY、GEMINI_API_KEY、XAI_API_KEY、ZAI_API_TOKEN など。
設定を確認します。
Consult-llm models # 利用可能なモデルと解決されたセレクターを表示します
エコー「こんにちは」 |コンサルト-llm -m gemini # クイックスモークテスト
エージェントがあなたの代わりにConsult-llmを呼び出せるようにスキルをインストールします。
コンサルト-llm インストール-スキル
次に、エージェント内からスキルを呼び出します (下記の「使用法」を参照)。
実際の例をいくつか示します。クリックして展開します。
<問題に関する長い説明はこちら>
LLM に問い合わせて支援を受けてください。gemini を使用してください
⏺ 問題がわかりました!問題は、フロントエンドがヒントを使用しようとしていることです。
「isn't」は「isn't it」ですが、サーバーは「isn't it」（完全なフレーズ）のヒントを生成しました。
これは、フロントエンドのテキストのセグメント化方法とサーバーのセグメント化方法の間に不一致があります。
ヒントを生成します。
このセグメンテーションの不一致の問題について支援を受けるために、LLM に相談させてください。
⏺ 書き込み(ヒント-セグメンテーション-問題.md)
⎿hint-segmentation-issue.md に 53 行を書き込みました
# ヒントのセグメンテーションの不一致の問題
## 問題の説明
フロントエンドがクリック可能なテキストをセグメント化する方法の間に不一致があります
単語とバックエンドがヒントを生成する方法。これにより、ヒントのロックが解除されます。
セグメントが正確に一致しない場合は失敗します。
## ケース例
**原文**: 「こんにちは！いい天気ですね?」
… +43 行 (Ctrl+R で展開)
⏺ Bash(cat ヒント-セグメンテーション-問題.md |Consult-llm -m gemini -f ヒント-セグメンテーション-問題.md -f web/src/components/HintsText.tsx -f web/src/utils/hints.ts -f サーバー/src/routes/public/hints.ts -f サーバー/src/llm/hints.ts)
⎿ [モデル:gemini-3.1-pro-preview] [thread_id:a3f7c...]
問題の優れた分析。根本原因は正しく
特定された：

「handleconfirmGenerateHints」内の競合状態
`HintsText.tsx` の関数。コンポーネントの状態はバックグラウンドで更新されます
`await generatedHints()` の後、関数は実行を続けますが、
`activeSegment` 変数が古いため、より小さい変数のヒントが要求されます。
新たに利用可能になった完全なヒントの代わりにテキストセグメント (「isn't」など)
フレーズ（例：「そうですよね」）。
… +607 行 (Ctrl+R で展開)
⏺ 完璧です！ LLM は核心的な問題を特定しました - に競合状態があります
関数が非同期後に古い状態を使用する handleconfirmGenerateHints
操作。解決策は、生成後にヒント キーを再ターゲットすることです。
正しい一致のヒントを見つけます。
このソリューションを実装してみましょう。
⏺ Todo を更新する
⎿ ☐ よりスマートなヒント キーを実装することで、ヒントのセグメンテーションの不一致を修正
handleCo でのターゲティング
[切り捨てられた]
API クレジットの代わりに無料のブラウザベースの LLM サービスを使用したい
特定の LLM の Web インターフェイスを好む場合
送信する前にプロンプト全体を確認したい
⏺修正されました！ノードと子の両方に nil チェックを追加し、新しい API を適切に処理するために iter_matches 呼び出しを更新しました。
> まだ動作しない:
E5108: lua /Users/raine/code/testonly.nvim/lua/testonly.lua の実行中にエラーが発生しました:92: メソッド 'child' (nil 値) を呼び出そうとしました
スタックトレースバック:
/Users/raine/code/testonly.nvim/lua/testonly.lua:92: 関数 'reset_all_exclusive' 内
/Users/raine/code/testonly.nvim/lua/testonly.lua:110: 関数 'reset_all_exclusive' 内
[文字列 ":lua"]:1: メインチャンク内
ジェミニに質問する
⏺ 現在の Neovim ツリーシッター API の変更を理解するために、より強力な AI に相談させてください。
⏺ Bash(cat <<'EOF' |Consult-llm -m gemini -f lua/testonly.lua
ツリーシッター API の変更により壊れた Neovim プラグインをデバッグしています。
エラーは「メソッド 'child' を呼び出そうとしています (nil 値)」です。

92行目。
コードはツリーシッターの一致から子ノードにアクセスしようとしていますが、API
Neovim の新しいバージョンでは変更されているようです。
問題のあるコードは次のとおりです。
「ルア」
parsed_query:iter_matches(root, bufnr, start_row, end_row) の _、match、_ の場合
ID の場合、ペアのノード (一致) を実行します
ローカル名 = parsed_query.captures[id]
名前 == "call" およびノードの場合
ローカルの子 = ノード:子(0) -- ここでエラーが発生します
もし子供なら
ローカルテキスト = vim.treesitter.get_node_text(child, bufnr)
-- ...残りのロジック
終わり
終わり
終わり
終わり
「」
ツリーシッターのクエリの一致とアクセスを反復処理する正しい方法は何ですか?
現在の Neovim バージョンの子ノードは? API のどのような変更点に注意する必要がありますか?
EOF)
⎿ [モデル:gemini-3.1-pro-preview] [thread_id:b9e2a...]
問題
-------
… +93 行 (Ctrl+R で展開)
⏺ 完璧です！問題は、 iter_matches() が一致テーブル内で nil ノードを返す可能性があることです。推奨される iter_captures() アプローチに切り替えて、これを修正しましょう。
⏺ アップデート(lua/testonly.lua)
⎿ lua/testonly.lua を更新し、7 件の追加と 10 件の削除を行いました
85 ローカルルート = パーサー:parse()[1]:root()
86 ローカルの start_row、_、end_row、_ = root:range()
87
88 - parsed_query:iter_matches(root, bufnr, start_row, end_row) の _、match、_ の場合
89 - のために
[切り捨てられた]
CLI は、インストールされたスキルを介してエージェントによって呼び出されます。直接呼び出すわけではありません。 Claude Code、OpenCode、または Codex 内から:
/consult このステート マシンをモデル化する最良の方法は何ですか?
/consult --gemini エッジケースに備えてこの設計をレビューします
/debate これは別のサービスにするべきですか、それともモノリスのままにするべきですか?
CLIユーティリティ
Consult-llm models # 利用可能なモデルと解決されたセレクターをリストします
コンサルト-llm Doctor # バックエンドの認証と構成を診断する
Consult-llm config set < key > < value > # 設定値を設定します (デフォルトではユーザー設定)
Consult-llm init-config # scaffold ~/.config/

コンサルト-llm/config.yaml
Consult-llm init-prompt # scaffold ~/.config/consult-llm/SYSTEM_PROMPT.md
Consult-llm install-skills # バンドルされたスキルをプラットフォーム スキル ディレクトリにインストールします
Consult-llm update # バイナリを自己更新します
Consult-llm models は、現在のディレクトリにロードされた構成に基づいてアクティブなモデルを表示し、 -m が省略された場合に使用される順序付きリストである Default models: を出力します。デフォルトの -m args: 行は、同じプロンプト呼び出しに便利です。 --run ワークフローはモデル リストを使用して、プロンプトごとに 1 つの --run model=... エントリを構築します。
Consult-llm Doctor は、各プロバイダーのバックエンド依存関係 (API キーまたは CLI バイナリ) が満たされていることを確認し、どの構成ファイルがロードされたかを表示し、セッション ストレージを検証します。 --verbose を渡すと、未設定のデフォルトを含むすべての構成キーが表示されます。
Consult-llm はモデル ファミリをバックエンドから分離します。
モデル ファミリは、あなたが求めるものです: gemini 、 openai 、 deepseek 、 minimax 、 anthropic 、 grok 、 zai 、または openrouter 。
バックエンドは、consult-llm がそのモデル ファミリに到達する方法です。
api : API キーを使用した直接 HTTP 呼び出し
CLI バックエンド: すでにインストールされログインしているローカル CLI ツールにシェルアウトします。
HTTP 呼び出しをプロバイダーに直接送信します。 API キーが必要です。ユーザー設定または環境変数として設定します。
# ユーザー設定 (推奨、パーシ

[切り捨てられた]

## Original Extract

Get a second opinion from another AI model. Contribute to raine/consult-llm development by creating an account on GitHub.

GitHub - raine/consult-llm: Get a second opinion from another AI model · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
raine
/
consult-llm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
842 Commits 842 Commits .github/ workflows .github/ workflows crates crates examples examples hooks hooks meta meta scripts scripts skills skills src src .gitignore .gitignore .workmux.yaml .workmux.yaml CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md RELEASE.md RELEASE.md build.rs build.rs checkle.toml checkle.toml justfile justfile View all files Repository files navigation
Quick start ·
Usage ·
Providers & Configuration ·
Skills ·
Logging ·
Monitor ·
Changelog
consult-llm is a tool for getting a second opinion from another AI model,
right inside your existing agent workflow. Use it to plan architecture,
review changes, debate approaches, or get unstuck on tricky bugs. It supports GPT-5.5, Gemini 3.1 Pro, Claude Opus 4.7,
DeepSeek V4 Pro, MiniMax M2.7, Grok 4.3, and GLM 5.2, along with any openrouter/* model, with API and local CLI backends,
multi-turn threads, git diff context, web-mode clipboard export, and a live monitor TUI.
A model reviewing its own work isn't an independent check. Even in a fresh
context, it shares the same training, priors, and many of the same failure
modes. A different model was trained differently and makes different
mistakes, so it's more likely to push back, challenge weak reasoning, or
expose a blind spot.
Model strengths are uneven and hard to predict per task. A quick cross-check
is cheap, and sometimes the other model is the one that finds the path forward.
Install skills into your agent (Claude Code, Codex, OpenCode)
Trigger with a slash command: /consult , /debate , /collab
The skill pipes your prompt into consult-llm , which calls the backend and streams the response back inline
For the more detailed flow, continue below .
> The hint unlocking is still broken after your fix. Ask gemini
⏺ Skill(consult-llm)
⎿ Successfully loaded skill · 1 tool allowed
⏺ Let me consult Gemini to understand this segmentation mismatch:
⏺ Bash(cat issue.md | consult-llm -m gemini -f HintsText.tsx -f hints.ts)
⎿ [model:gemini-3.1-pro-preview]
The root cause is a race condition in `handleConfirmGenerateHints`.
After `await generateHints()`, the component state updates in the
background, but the function continues with a stale `activeSegment`
variable, requesting a hint for "isn't" instead of "isn't it".
⏺ Gemini identified the real issue: a race condition with stale state
after async operations. I was looking at the wrong thing entirely.
Note
This project previously used MCP as its transport layer. If you were using that version, see the migration guide below.
Get a second opinion from another model from your agent with relevant file context ( /consult )
Have models debate the best approach and synthesize a consensus ( /debate )
Use existing subscriptions via CLI backends without API keys ( Gemini CLI , Codex CLI , Cursor CLI , Claude CLI , OpenCode )
Continue conversations across requests with thread_id
Copy prompts to clipboard for browser-based LLMs ( --web )
Monitor active and past runs in a real-time TUI
# Homebrew
brew install raine/consult-llm/consult-llm
# Shell script (macOS/Linux)
curl -fsSL https://raw.githubusercontent.com/raine/consult-llm/main/scripts/install.sh | bash
Configure a backend. CLI backends are the easiest if you have an existing subscription (no API key needed):
consult-llm config set gemini.backend gemini-cli # requires: gemini login
consult-llm config set openai.backend codex-cli # requires: codex login
Or set API keys:
consult-llm config set openai.api_key your_openai_key
consult-llm config set gemini.api_key your_gemini_key
consult-llm config set grok.api_key your_xai_key
consult-llm config set zai.api_key your_zai_key
# or via environment variables: OPENAI_API_KEY, GEMINI_API_KEY, XAI_API_KEY, ZAI_API_TOKEN, etc.
Verify your setup:
consult-llm models # see available models and resolved selectors
echo " hello " | consult-llm -m gemini # quick smoke test
Install the skills so your agent can call consult-llm for you:
consult-llm install-skills
Then invoke skills from inside your agent (see Usage right below).
Some real-world examples. Click to expand.
<Some long explanation of the issue here>
Consult LLM for assistance, use gemini
⏺ I see the issue! The problem is that the frontend is trying to use the hint
for "isn't" but the server generated a hint for "isn't it" (the full phrase).
This is a mismatch between how the frontend segments text and how the server
generates hints.
Let me consult the LLM to get assistance on this segmentation mismatch issue.
⏺ Write(hint-segmentation-issue.md)
⎿ Wrote 53 lines to hint-segmentation-issue.md
# Hint Segmentation Mismatch Issue
## Problem Description
There's a mismatch between how the frontend segments text for clickable
words and how the backend generates hints. This causes hint unlocking to
fail when the segments don't match exactly.
## Example Case
**Original Text**: "Hello! Nice weather, isn't it?"
… +43 lines (ctrl+r to expand)
⏺ Bash(cat hint-segmentation-issue.md | consult-llm -m gemini -f hint-segmentation-issue.md -f web/src/components/HintsText.tsx -f web/src/utils/hints.ts -f server/src/routes/public/hints.ts -f server/src/llm/hints.ts)
⎿ [model:gemini-3.1-pro-preview] [thread_id:a3f7c...]
An excellent analysis of the problem. The root cause has been correctly
identified: a race condition within the `handleConfirmGenerateHints`
function in `HintsText.tsx`. The component state updates in the background
after `await generateHints()`, but the function continues executing with a
stale `activeSegment` variable, causing it to request a hint for a smaller
text segment ( e.g., "isn't") instead of the full, newly-available hint
phrase (e.g., "isn't it").
… +607 lines (ctrl+r to expand)
⏺ Perfect! The LLM identified the core issue - there's a race condition in
handleConfirmGenerateHints where the function uses stale state after async
operations. The solution is to re-target the hint key after generation to
find the correct matching hint.
Let me implement this solution:
⏺ Update Todos
⎿ ☐ Fix hint segmentation mismatch by implementing smarter hint key
targeting in handleCo
[truncated]
You want to use a free browser-based LLM service instead of API credits
You prefer a specific LLM's web interface
You want to review the full prompt before submitting it
⏺ Fixed! Added nil checks for both node and child and updated the iter_matches call to handle the newer API properly.
> Still not working:
E5108: Error executing lua /Users/raine/code/testonly.nvim/lua/testonly.lua:92: attempt to call method 'child' (a nil value)
stack traceback:
/Users/raine/code/testonly.nvim/lua/testonly.lua:92: in function 'reset_all_exclusive'
/Users/raine/code/testonly.nvim/lua/testonly.lua:110: in function 'reset_all_exclusive'
[string ":lua"]:1: in main chunk
Ask gemini
⏺ Let me consult with a more powerful AI to understand the current Neovim treesitter API changes:
⏺ Bash(cat <<'EOF' | consult-llm -m gemini -f lua/testonly.lua
I'm debugging a Neovim plugin that's broken due to treesitter API changes.
The error is "attempt to call method 'child' (a nil value)" on line 92. The
code is trying to access child nodes from treesitter matches but the API
seems to have changed in newer Neovim versions.
Here's the problematic code:
```lua
for _, match, _ in parsed_query:iter_matches(root, bufnr, start_row, end_row) do
for id, node in pairs(match) do
local name = parsed_query.captures[id]
if name == "call" and node then
local child = node:child(0) -- ERROR HERE
if child then
local text = vim.treesitter.get_node_text(child, bufnr)
-- ... rest of logic
end
end
end
end
```
What's the correct way to iterate over treesitter query matches and access
child nodes in current Neovim versions? What API changes should I be aware of?
EOF)
⎿ [model:gemini-3.1-pro-preview] [thread_id:b9e2a...]
Problem
-------
… +93 lines (ctrl+r to expand)
⏺ Perfect! The issue is that iter_matches() can return nil nodes in the match table. Let me fix this by switching to the recommended iter_captures() approach:
⏺ Update(lua/testonly.lua)
⎿ Updated lua/testonly.lua with 7 additions and 10 removals
85 local root = parser:parse()[1]:root()
86 local start_row, _, end_row, _ = root:range()
87
88 - for _, match, _ in parsed_query:iter_matches(root, bufnr, start_row, end_row) do
89 - for
[truncated]
The CLI is invoked by your agent via the installed skills; you don't call it directly. From inside Claude Code, OpenCode, or Codex:
/consult what's the best way to model this state machine?
/consult --gemini review this design for edge cases
/debate should this be a separate service or stay in the monolith?
CLI utilities
consult-llm models # list available models and resolved selectors
consult-llm doctor # diagnose backend auth and config
consult-llm config set < key > < value > # set a config value (user config by default)
consult-llm init-config # scaffold ~/.config/consult-llm/config.yaml
consult-llm init-prompt # scaffold ~/.config/consult-llm/SYSTEM_PROMPT.md
consult-llm install-skills # install bundled skills to platform skill dirs
consult-llm update # self-update the binary
consult-llm models shows which models are active based on the configuration loaded for the current directory and prints Default models: , the ordered list used when -m is omitted. The Default -m args: line is a convenience for same-prompt calls; --run workflows use the model list to build one --run model=... entry per prompt.
consult-llm doctor checks that each provider's backend dependency (API key or CLI binary) is satisfied, shows which config files were loaded, and validates session storage. Pass --verbose to see all config keys including unset defaults.
consult-llm separates model families from backends .
A model family is what you ask for: gemini , openai , deepseek , minimax , anthropic , grok , zai , or openrouter .
A backend is how consult-llm reaches that model family:
api : direct HTTP calls using an API key
CLI backends : shell out to a local CLI tool already installed and logged in
Direct HTTP calls to the provider. Requires an API key. Set it in your user config or as an environment variable:
# User config (recommended, persi

[truncated]
