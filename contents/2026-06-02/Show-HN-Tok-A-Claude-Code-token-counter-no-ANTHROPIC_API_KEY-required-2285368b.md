---
source: "https://github.com/samjoch/tok"
hn_url: "https://news.ycombinator.com/item?id=48358986"
title: "Show HN: Tok - A Claude Code token counter (no ANTHROPIC_API_KEY required)"
article_title: "GitHub - samjoch/tok: Claude Tokens Counter · GitHub"
author: "samjoch"
captured_at: "2026-06-02T00:36:53Z"
capture_tool: "hn-digest"
hn_id: 48358986
score: 4
comments: 0
posted_at: "2026-06-01T16:23:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tok - A Claude Code token counter (no ANTHROPIC_API_KEY required)

- HN: [48358986](https://news.ycombinator.com/item?id=48358986)
- Source: [github.com](https://github.com/samjoch/tok)
- Score: 4
- Comments: 0
- Posted: 2026-06-01T16:23:55Z

## Translation

タイトル: 表示 HN: Tok - クロード コード トークン カウンター (ANTHROPIC_API_KEY は必要ありません)
記事タイトル: GitHub - samjoch/tok: クロード トークン カウンター · GitHub
説明: クロード トークン カウンター。 GitHub でアカウントを作成して、samjoch/tok の開発に貢献してください。

記事本文:
GitHub - samjoch/tok: クロード トークン カウンター · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
サムジョッホ
/
トク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動

e コード [その他のアクション] メニューを開く フォルダーとファイル
1 コミット 1 コミット src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
端末の正確なクロード トークン カウンター。ローカルのクロード コードを再利用します
macOS キーチェーンからログイン — ANTHROPIC_API_KEY は必要ありません — そして呼び出します
Anthropic の公式 count_tokens エンドポイントなので、数値はクロードのものと一致します。
実は請求書。
カーゴビルド --release
# target/release/tok のバイナリ (オプション: cp target/release/tok ~/.local/bin/)
Claude Code がインストールされており、少なくとも 1 回はログインしている必要があります
(クロード)。 OAuth トークンはキーチェーン エントリから読み取られます
クロードコード認証情報 。
入力なしで実行すると、ライブカウントを含むクロードコードスタイルのフィールドが開きます。
トク
プロンプト (トークン数と推定入力コスト) を入力または貼り付けます。
編集すると更新されます（デバウンス）。
Ctrl+V — スマートペースト: クリップボードに画像がある場合は画像を添付します。
それ以外の場合はテキストを挿入します。 (プレーン ⌘V テキストの貼り付けも機能します。)
貼り付けた画像は、クロードと同様に [画像 1] 、 [画像 2] 、…としてインラインで表示されます。
コード。プレースホルダーは表示専用です。画像自体もプレースホルダーとしてカウントされます。
画像ブロックではありますが、[画像 N] テキストはそうではありません。
Ctrl+X — 最後に添付した画像を削除します · Ctrl+L — すべてをクリアします
テキスト、ファイル、またはイメージの引数は、単一の count-and-print に切り替わります。
tok " Hello, world! " # -> 16 トークン · ≈ $0.000240 入力
エコー「パイプから」 | tok # は標準入力を読み取ります
tok -q " 数字だけ " # -> 11 (整数のみ、スクリプトの場合)
tok --file prompt.txt # ファイルの内容をカウントします
tok " これを説明します " -i Shot.png # テキスト + 画像 (さらに表示するには -i を繰り返します)
tok -m claude-haiku-4-5 " hi " # 価格は選択したモデルを反映します
表示されているコストは入力コスト (このプロンプトを 1 回送信するのに必要なコスト) です。
モデルのトークごとに

n 入力価格 — count_tokens は出力を生成しません。
したがって、出力コストを含めることはありません。価格はモデルごとに一致する概算です
家族（作品/ソネット/俳句）。
旗
意味
-f, --file <パス>
ファイルからプロンプトテキストを読み取る
-i, --image <パス>
画像 (png/jpg/gif/webp) を添付します (繰り返し可能)
-m、--model <ID>
カウント対象のモデル (カウントはクロード モデル間で同一です)
-q、--静か
整数のみを出力します
--インタラクティブ
入力がパイプ接続されている場合でもフィールドを強制します
仕組み
tok は、 claudeAiOauth.accessToken (スコープ user:inference ) を読み取ります。
キーチェーンとテキスト/画像の投稿
https://api.anthropic.com/v1/messages/count_tokens と
anthropic-beta: oauth-2025-04-20 ヘッダー。カウントするためにコンテンツを送信するだけです。
完了は生成されず、トークンに対して何も請求されません。
401 が表示された場合は、ログインの有効期限が切れています。claude を 1 回実行して更新してください。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude Tokens Counter. Contribute to samjoch/tok development by creating an account on GitHub.

GitHub - samjoch/tok: Claude Tokens Counter · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
samjoch
/
tok
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml README.md README.md View all files Repository files navigation
Exact Claude token counter for the terminal. Reuses your local Claude Code
login from the macOS Keychain — no ANTHROPIC_API_KEY required — and calls
Anthropic's official count_tokens endpoint, so the number matches what Claude
actually bills.
cargo build --release
# binary at target/release/tok (optionally: cp target/release/tok ~/.local/bin/)
Requires that Claude Code is installed and you've logged in at least once
( claude ). The OAuth token is read from the Keychain entry
Claude Code-credentials .
Run with no input to open a Claude-Code-style field with a live count:
tok
Type or paste your prompt — the token count and estimated input cost
update as you edit (debounced).
Ctrl+V — smart paste: attaches an image if the clipboard holds one,
otherwise inserts text. (Plain ⌘V text paste also works.)
Pasted images appear inline as [Image 1] , [Image 2] , … just like Claude
Code. The placeholder is display-only — the image itself is counted as an
image block, the [Image N] text is not.
Ctrl+X — drop the last attached image · Ctrl+L — clear everything
Any text, file, or image argument switches to a single count-and-print:
tok " Hello, world! " # -> 16 tokens · ≈ $0.000240 input
echo " from a pipe " | tok # reads stdin
tok -q " just the number " # -> 11 (integer only, for scripts)
tok --file prompt.txt # count a file's contents
tok " describe this " -i shot.png # text + image (repeat -i for more)
tok -m claude-haiku-4-5 " hi " # price reflects the chosen model
The cost shown is the input cost (what it takes to send this prompt once)
at the model's per-token input price — count_tokens doesn't generate output,
so there's no output cost to include. Prices are approximate, matched by model
family (opus/sonnet/haiku).
flag
meaning
-f, --file <PATH>
read prompt text from a file
-i, --image <PATH>
attach an image (png/jpg/gif/webp), repeatable
-m, --model <ID>
model to count against (counts are identical across Claude models)
-q, --quiet
print only the integer
--interactive
force the field even when input is piped
How it works
tok reads the claudeAiOauth.accessToken (scope user:inference ) from the
Keychain and POSTs your text/images to
https://api.anthropic.com/v1/messages/count_tokens with the
anthropic-beta: oauth-2025-04-20 header. It only sends content for counting;
no completion is generated and nothing is billed for tokens.
If you see a 401, your login expired — run claude once to refresh it.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
