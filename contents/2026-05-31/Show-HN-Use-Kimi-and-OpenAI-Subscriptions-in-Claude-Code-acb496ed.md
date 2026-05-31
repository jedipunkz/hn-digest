---
source: "https://github.com/raine/claude-code-proxy"
hn_url: "https://news.ycombinator.com/item?id=48339755"
title: "Show HN: Use Kimi and OpenAI Subscriptions in Claude Code"
article_title: "GitHub - raine/claude-code-proxy: Use Claude Code with your ChatGPT or Kimi subscription via a local Anthropic-compatible proxy · GitHub"
author: "rane"
captured_at: "2026-05-31T00:26:31Z"
capture_tool: "hn-digest"
hn_id: 48339755
score: 1
comments: 0
posted_at: "2026-05-30T19:23:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Use Kimi and OpenAI Subscriptions in Claude Code

- HN: [48339755](https://news.ycombinator.com/item?id=48339755)
- Source: [github.com](https://github.com/raine/claude-code-proxy)
- Score: 1
- Comments: 0
- Posted: 2026-05-30T19:23:51Z

## Translation

タイトル: HN を表示: クロード コードで Kim と OpenAI サブスクリプションを使用する
記事のタイトル: GitHub - Raine/claude-code-proxy: ローカルの Anthropic 互換プロキシ経由で ChatGPT または Kim サブスクリプションでクロード コードを使用する · GitHub
説明: ローカルの Anthropic 互換プロキシ経由で、ChatGPT または Kim サブスクリプションで Claude Code を使用します - Raine/claude-code-proxy

記事本文:
GitHub - Raine/claude-code-proxy: ローカルの Anthropic 互換プロキシ経由で ChatGPT または Kim サブスクリプションで Claude Code を使用する · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
クロードコードプロキシ
公共
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
164 コミット 164 コミット .githooks .githooks .github/ workflows .github/ workflows メタ メタ スクリプト スクリプト src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md bun.lock bun.lock justfile justfile package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロードコードプロキシを使用すると、
ChatGPT を使用した Claude コード
Plus/Pro サブスクリプション、または Kimi Code (kimi.com) アカウント。
クイックスタート · プロバイダー ·
仕組み · 構成 ·
制限事項
時折あるものの、Claude Code が今でも最高のハーネスであると感じています。
アップデートによる不満。ただし、Anthropic は使用法を厳しくし続けています
制限はありますが、OpenAI はさらに寛大です。
OpenAI プランを使用したい場合、最適なオプションは OpenCode と
コーデックス。 OpenCode を試してみましたが、UX には多くの荒削りな部分があり、特に
スキルは二流の機能のように感じます。幸いなことに、それはオープンソースであり、
結局フォークしてパッチを適用することになりましたが、それはやめたほうが良いでしょう。
brew install Raine/claude-code-proxy/claude-code-proxy
スクリプトをインストールします (macOS および Linux):
カール -fsSL https://raw.githubusercontent.com/raine/claude-code-proxy/main/scripts/install.sh |バッシュ
手動: プラットフォーム用の事前構築済みバイナリを次の場所からダウンロードします。
リリースページ 。窓
アーティファクトは claude-code-proxy-windows-amd64.zip として公開されます。
クロードコードプロキシ Windows-arm64.zip ; .exe をコンピュータ上のどこかに抽出します。
パス 。
2. プロバイダーを選択して認証する
プロキシは 2 つのアップストリーム プロバイダーをサポートします。 1 つを選択してログイン フローを実行します。の
プロキシは、トークンが保存されるまでトラフィックの開始を拒否します。
クロードコードプロキシタラ

ex 認証ログイン # ブラウザ OAuth (PKCE)
# または、ヘッドレス マシンの場合:
クロード コード プロキシ コーデックス認証デバイス # デバイス コード フロー
OpenAI API アカウントではなく、ChatGPT Plus/Pro アカウントでサインインします。
claude-code-proxy kimi auth login # device-code flow (URL + コードを出力)
kimi.com アカウントでサインインします。検証 URL が表示されます。開く
任意のブラウザでコードを確認し、完了するまで CLI がポーリングします。
macOS では、認証情報はキーチェーンに移動します。 Windows では、次のように記述されます。
%APPDATA%\claude-code-proxy\<プロバイダー>\auth.json ; Linux では次のように書かれています
${XDG_CONFIG_HOME:-$HOME/.config}/claude-code-proxy/<プロバイダ>/auth.json の下
(サポートされている場合はモード 0600)。
クロードコードプロキシのコーデックス認証ステータス
クロードコードプロキシキミ認証ステータス
3. プロキシを開始します
クロードコードプロキシサーブ # 127.0.0.1:18765 をリッスンします
PORT=11435 claude-code-proxyserve # 待機ポートを変更します
127.0.0.1 のみにバインドします。 1 つのサーブ プロセスですべてのプロバイダーが処理されます。
各リクエストのアップストリームは ANTHROPIC_MODEL から選択されます。
ANTHROPIC_MODEL はプロバイダーを選択します。
gpt-5.5 、 gpt-5.4 、 gpt-5.3-codex 、 gpt-5.3-codex-spark 、 gpt-5.4-mini 、 gpt-5.2 → codex
kimi-for-coding 、 kimi-k2.6 、 k2.6 → kimi
不明なモデルは、サポートされている ID をリストした 400 を返します。ありません
暗黙的なデフォルトプロバイダー。
Claude Code はバックグラウンド リクエストも発行します (セッション タイトルの生成、トークン
カウント) を組み込みの「小さい/速い」俳句モデル ID と照合します。それらのリクエスト
プロバイダーがそれを要求していないため、400 になります。
ANTHROPIC_SMALL_FAST_MODEL も具体的な ID (と同じ値) に
通常は ANTHROPIC_MODEL で問題ありません):
# コーデックス
ANTHROPIC_BASE_URL=http://localhost:18765 \
ANTHROPIC_AUTH_TOKEN=未使用 \
ANTHROPIC_MODEL=gpt-5.4[1m] \
ANTHROPIC_SMALL_FAST_MODEL=gpt-5.4-mini[1m] \
CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC=1 \
CLAUDE_CODE_DISABLE_NONSTRE

AMING_FALLBACK=1 \
クロード
#キミ
ANTHROPIC_BASE_URL=http://localhost:18765 \
ANTHROPIC_AUTH_TOKEN=未使用 \
ANTHROPIC_MODEL=キミ・フォー・コーディング[1m] \
ANTHROPIC_SMALL_FAST_MODEL=キミ・フォー・コーディング[1m] \
CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC=1 \
CLAUDE_CODE_DISABLE_NONSTREAMING_FALLBACK=1 \
クロード
CLAUDE_CODE_DISABLE_NONSTREAMING_FALLBACK=1 をお勧めします。
プロキシは、ストリーミング リクエストを使用して常に上流のプロバイダーと通信します。
クロード コードの非ストリーミング人間応答を蓄積します。クロードの無効化
コードのストリーミングから非ストリーミングへのフォールバックにより、部分的に完了したコードの再試行が回避されます。
ツール呼び出しを複製できる方法でストリーミングします。
または、 ~/.claude/settings.json で永続的に設定します。
{
"環境" : {
"ANTHROPIC_BASE_URL" : " http://127.0.0.1:18765 " 、
"ANTHROPIC_AUTH_TOKEN" : " 未使用 " 、
"ANTHROPIC_MODEL" : " gpt-5.4[1m] " 、
"ANTHROPIC_SMALL_FAST_MODEL" : " gpt-5.4-mini[1m] " ,
"CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC" : 1 、
"CLAUDE_CODE_DISABLE_NONSTREAMING_FALLBACK" : 1
}
}
5. コンテキストウィンドウのサイズ
Claude Code は、モデルのコンテキスト ウィンドウに基づいて自動圧縮を決定します。のために
不明なモデル (プロキシが使用するモデルなど) の場合、デフォルトでは 200K トークンになります。
アップストリーム モデルが実際にサポートするもの (GPT-5.4: 400K+、
キミ: 256K)。これにより、自動圧縮が必要以上に早く実行されます。
モデル名の末尾 [1m] (上記の例に示されている) は、Claude です。
代わりに 1M トークンのコンテキスト ウィンドウを使用するように指示するコード規則。これ
完全に無効にすることなく、自動圧縮のしきい値を上げます。
自動圧縮を完全に無効にしたい場合は、次のように設定します。
env または ~/.claude/settings.json で DISABLE_AUTO_COMPACT=1 。マニュアル
/compact は引き続き機能しますが、その前に実際のアップストリーム制限に達する危険性があります。
Claude Code はあなたのためにコンパクトにすることができます。
プロキシとディアーの切り替え

ct人体
フォールバックしたい Anthropic サブスクリプションがまだある場合は、次のことができます。
プロキシ環境変数のみを注入する小さなラッパーをクロードの前に置きます
フラグ ファイルが存在する場合、フラグを反転するためのトグル スクリプトが追加されます。休暇を取る
~/.claude/settings.json にはプロキシ環境変数が含まれていないため、Anthropic への直接接続が残ります
デフォルト。
~/.local/bin/claude ( PATH 上の実際のクロードの前):
#! /bin/bash
# オプションで claude-code-proxy にルーティングするラッパー。
# ~/.claude/claude-code-proxy-enabled が存在する場合にアクティブになります。
if [ -f " $HOME /.claude/claude-code-proxy-enabled " ] ;それから
エクスポート ANTHROPIC_BASE_URL= " http://localhost:18765 "
エクスポート ANTHROPIC_AUTH_TOKEN= " 未使用 "
エクスポート ANTHROPIC_MODEL= " gpt-5.4[1m] "
エクスポート ANTHROPIC_SMALL_FAST_MODEL= " gpt-5.4-mini[1m] "
エクスポート CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC= " 1 "
エクスポート CLAUDE_CODE_DISABLE_NONSTREAMING_FALLBACK= " 1 "
フィ
exec " $HOME /.local/bin/claude " " $@ "
実際のクロードバイナリがコンピュータ上の別の場所に存在する場合は、実行パスを調整します。
システム (例: $(bun pm bin -g)/claude 、 $HOME/.claude/local/claude )。
claude-proxy-toggle ( PATH 上の任意の場所):
#! /bin/bash
# クロード ラッパーのクロード コード プロキシ ルーティングを切り替えます。
set -euo パイプ失敗
flag= " $HOME /.claude/claude-code-proxy-enabled "
if [ -f " $flag " ] ;それから
rm " $flag "
エコー「プロキシ: オフ」
それ以外の場合
mkdir -p " $( dirname " $flag " ) "
「$flag」をタッチ
エコー「プロキシ: オン」
フィ
claude-proxy-toggle を実行して、プロキシ経由のルーティングを切り替えます (Codex /
キミ）そしてAnthropicと直接話しています。新規または継続のクロードセッションがピックアップされます
すぐに変更します。既存のセッションでは、最初のセッションがすべて保持されます。
アップストリーム: https://chatgpt.com/backend-api/codex/responses (レスポンス API)。
ANTHROPIC_MODEL を、ChatGPT サブスクリプションが使用を許可されているモデルに設定します。
Codex モデル名に -fast を追加します

e そのリクエストに対して Codex 高速モードをリクエストします
プロキシを再起動せずに。たとえば、 gpt-5.4-fast[1m] は次のようにアップストリームに送信されます。
モデル gpt-5.4 と service_tier: "priority" 。明示的な
codex.serviceTier / CCP_CODEX_SERVICE_TIER オーバーライドが引き続き優先されます。
推論の労力: Claude Code の Output_config.effort 値 (
UI は ◐ Medium · /effort ) として Codexreasoning.effort ( low
/中/高/x高)。クロード コードの最大値は次のように上流に送信されます。
xhigh 。明示的な codex.effort / CCP_CODEX_EFFORT オーバーライドにはまだ時間がかかります
が優先され、何も強制することもできます。
解決されたモデルがアカウントでサポートされていない場合、アップストリームは 400 を返します。
好き
「ChatGPT アカウントで Codex を使用する場合、「gpt-4.1」モデルはサポートされません。」 。
代理人はそれをそのまま明らかにします。
アップストリーム: https://api.kimi.com/coding/v1/chat/completions (OpenAI スタイル)
チャット補完)。
公開されているワイヤ モデルは 1 つだけです: kimi-for-coding (kimi-cli での表示名)
は Kim-k2.6 、 256k コンテキスト、推論 + 画像入力 + ビデオ入力をサポートします)。
kimi-k2.6 と k2.6 は、同じワイヤ ID のエイリアスとして受け入れられます。
推論の労力: Claude Code の Output_config.effort 値 (
UI は ◐ 中 · /effort ) として Kim のreasoning_effort ( low
/中/高）。上流モデルからの思考ブロックは次へ転送されます。
クロード コードを思考コンテンツとしてレンダリングします。クロード・コードが思考を無効にするなら、
プロキシは、reasoning_effort と思考の両方を削除します: {type: "enabled"}
転送する前にフラグを立てます。
シーケンス図
自動番号付け
クロード・コードとしての参加者CC
クロードコードプロキシとしての参加者 P
OAuth ホストとしての参加者 AUTH<br/>(auth.openai.com または<br/>auth.kimi.com)
アップストリーム API としての参加者 U<br/>(chatgpt.com/codex または<br/>api.kimi.com)
P,AUTH に関するメモ: ワンタイム: PKCE / de

Vice OAuth<br/>再利用のためにローカルにキャッシュされたトークン
CC->>P: POST /v1/messages (人間の形状、ストリーム: true)
alt アクセス トークンの有効期限が切れます
P->>AUTH: POST /oauth/token (refresh_token)
AUTH-->>P: 新しいアクセス (+ ローテーションされたリフレッシュ)
終わり
P->>P: リクエストの変換<br/>• Anthropic 専用フィールドの削除<br/>• システム ブロック → 命令 / システム メッセージ<br/>• tools_use / tool_result ↔ プロバイダー固有の形状<br/>• プロンプト_キャッシュ_キー = セッション ID
P->>U: POST アップストリーム<br/>ベアラー + プロバイダー固有のヘッダー
U-->>P: プロバイダー SSE<br/>(コーデックス: Output_item.*、output_text.delta、…)<br/>(キミ: chat.completion.chunk、reasoning_content、…)
P->>P: リデューサー: 型付きイベント<br/>(思考 / テキスト / ツールの開始/デルタ/停止、終了)
P-->>CC: 人間性 SSE<br/>(message_start、content_block_*、message_delta、message_stop)
読み込み中
コマンド
コマンド
説明
奉仕する
PORT でプロキシを開始します
コーデックス認証ログイン/デバイス/ステータス/ログアウト
コーデックスOAuth管理
kimi auth ログイン / ステータス / ログアウト
キミ OAuth 管理
奉仕する
HTTP プロキシを開始してブロックします。 127.0.0.1 のみにバインドします。にログを記録します
プラットフォーム状態ディレクトリ (20 MiB でローテーション)。 CCP_LOG_STDERR=1 をミラーリングに設定します
実行中にログ行を標準エラー出力に記録します。
クロードコードプロキシサーブ
PORT=11435 クロードコードプロキシサーブ
CCP_LOG_STD

[切り捨てられた]

## Original Extract

Use Claude Code with your ChatGPT or Kimi subscription via a local Anthropic-compatible proxy - raine/claude-code-proxy

GitHub - raine/claude-code-proxy: Use Claude Code with your ChatGPT or Kimi subscription via a local Anthropic-compatible proxy · GitHub
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
raine
/
claude-code-proxy
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
164 Commits 164 Commits .githooks .githooks .github/ workflows .github/ workflows meta meta scripts scripts src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md bun.lock bun.lock justfile justfile package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
claude-code-proxy lets you use
Claude Code with your ChatGPT
Plus/Pro subscription or your Kimi Code (kimi.com) account.
Quick start · Providers ·
How it works · Configuration ·
Limitations
I feel Claude Code is still the best harness around, despite occasional
frustrations caused by updates. However, Anthropic keeps tightening the usage
limits, while OpenAI is still much more generous.
If you want to use OpenAI plans, your best options seem to be OpenCode and
Codex. I tried OpenCode, but the UX has many rough edges, especially around
skills feeling like a second-class feature. Fortunately it's open source and I
ended up forking it and applying some patches, but would much rather not do it.
brew install raine/claude-code-proxy/claude-code-proxy
Install script (macOS and Linux):
curl -fsSL https://raw.githubusercontent.com/raine/claude-code-proxy/main/scripts/install.sh | bash
Manual: download a prebuilt binary for your platform from the
releases page . Windows
artifacts are published as claude-code-proxy-windows-amd64.zip and
claude-code-proxy-windows-arm64.zip ; extract the .exe somewhere on your
PATH .
2. Pick a provider and authenticate
The proxy supports two upstream providers. Pick one and run its login flow; the
proxy will refuse to start traffic until a token is stored.
claude-code-proxy codex auth login # browser OAuth (PKCE)
# or, on a headless machine:
claude-code-proxy codex auth device # device-code flow
Sign in with your ChatGPT Plus/Pro account , not an OpenAI API account.
claude-code-proxy kimi auth login # device-code flow (prints URL + code)
Sign in with your kimi.com account . The verification URL is displayed; open
it in any browser, confirm the code, and the CLI polls until done.
On macOS credentials go to Keychain. On Windows they are written under
%APPDATA%\claude-code-proxy\<provider>\auth.json ; on Linux they are written
under ${XDG_CONFIG_HOME:-$HOME/.config}/claude-code-proxy/<provider>/auth.json
(mode 0600 where supported).
claude-code-proxy codex auth status
claude-code-proxy kimi auth status
3. Start the proxy
claude-code-proxy serve # listens on 127.0.0.1:18765
PORT=11435 claude-code-proxy serve # change the listen port
Binds to 127.0.0.1 only. One serve process handles all providers — the
upstream for each request is chosen from ANTHROPIC_MODEL .
ANTHROPIC_MODEL selects the provider:
gpt-5.5 , gpt-5.4 , gpt-5.3-codex , gpt-5.3-codex-spark , gpt-5.4-mini , gpt-5.2 → codex
kimi-for-coding , kimi-k2.6 , k2.6 → kimi
An unknown model returns a 400 listing the supported ids. There is no
implicit default provider.
Claude Code also issues background requests (session title generation, token
counts) against its built-in "small/fast" haiku model id. Those requests
would 400 because no provider claims it, so set
ANTHROPIC_SMALL_FAST_MODEL to a concrete id too (the same value as
ANTHROPIC_MODEL is usually fine):
# Codex
ANTHROPIC_BASE_URL=http://localhost:18765 \
ANTHROPIC_AUTH_TOKEN=unused \
ANTHROPIC_MODEL=gpt-5.4[1m] \
ANTHROPIC_SMALL_FAST_MODEL=gpt-5.4-mini[1m] \
CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC=1 \
CLAUDE_CODE_DISABLE_NONSTREAMING_FALLBACK=1 \
claude
# Kimi
ANTHROPIC_BASE_URL=http://localhost:18765 \
ANTHROPIC_AUTH_TOKEN=unused \
ANTHROPIC_MODEL=kimi-for-coding[1m] \
ANTHROPIC_SMALL_FAST_MODEL=kimi-for-coding[1m] \
CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC=1 \
CLAUDE_CODE_DISABLE_NONSTREAMING_FALLBACK=1 \
claude
CLAUDE_CODE_DISABLE_NONSTREAMING_FALLBACK=1 is recommended because the
proxy always talks to upstream providers with streaming requests, even when it
accumulates a non-streaming Anthropic response for Claude Code. Disabling Claude
Code's streaming-to-non-streaming fallback avoids retrying a partially completed
stream in a way that can duplicate tool calls.
Or set it persistently in ~/.claude/settings.json :
{
"env" : {
"ANTHROPIC_BASE_URL" : " http://127.0.0.1:18765 " ,
"ANTHROPIC_AUTH_TOKEN" : " unused " ,
"ANTHROPIC_MODEL" : " gpt-5.4[1m] " ,
"ANTHROPIC_SMALL_FAST_MODEL" : " gpt-5.4-mini[1m] " ,
"CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC" : 1 ,
"CLAUDE_CODE_DISABLE_NONSTREAMING_FALLBACK" : 1
}
}
5. Context window size
Claude Code decides auto-compaction based on the model's context window. For
unknown models (like the ones the proxy uses) it defaults to 200K tokens, which
is smaller than what the upstream models actually support (GPT-5.4: 400K+,
Kimi: 256K). This causes auto-compact to fire earlier than necessary.
The [1m] suffix on the model name (shown in the examples above) is a Claude
Code convention that tells it to use a 1M-token context window instead. This
raises the auto-compact threshold without disabling it entirely.
If you'd rather disable auto-compact completely, set
DISABLE_AUTO_COMPACT=1 in your env or ~/.claude/settings.json . Manual
/compact still works, but you risk hitting real upstream limits before
Claude Code can compact for you.
Toggling between proxy and direct Anthropic
If you still have an Anthropic subscription you want to fall back to, you can
put a small wrapper in front of claude that only injects the proxy env vars
when a flag file exists, plus a toggle script to flip the flag. Leave
~/.claude/settings.json free of proxy env vars so direct-to-Anthropic remains
the default.
~/.local/bin/claude (ahead of the real claude on PATH ):
#! /bin/bash
# Wrapper that optionally routes to claude-code-proxy.
# Active when ~/.claude/claude-code-proxy-enabled exists.
if [ -f " $HOME /.claude/claude-code-proxy-enabled " ] ; then
export ANTHROPIC_BASE_URL= " http://localhost:18765 "
export ANTHROPIC_AUTH_TOKEN= " unused "
export ANTHROPIC_MODEL= " gpt-5.4[1m] "
export ANTHROPIC_SMALL_FAST_MODEL= " gpt-5.4-mini[1m] "
export CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC= " 1 "
export CLAUDE_CODE_DISABLE_NONSTREAMING_FALLBACK= " 1 "
fi
exec " $HOME /.local/bin/claude " " $@ "
Adjust the exec path if the real claude binary lives elsewhere on your
system (e.g. $(bun pm bin -g)/claude , $HOME/.claude/local/claude ).
claude-proxy-toggle (anywhere on your PATH ):
#! /bin/bash
# Toggle claude-code-proxy routing for the claude wrapper.
set -euo pipefail
flag= " $HOME /.claude/claude-code-proxy-enabled "
if [ -f " $flag " ] ; then
rm " $flag "
echo " proxy: off "
else
mkdir -p " $( dirname " $flag " ) "
touch " $flag "
echo " proxy: on "
fi
Run claude-proxy-toggle to flip between routing through the proxy (Codex /
Kimi) and talking to Anthropic directly. New or continued claude sessions pick up
the change immediately; existing sessions keep whatever they started with.
Upstream: https://chatgpt.com/backend-api/codex/responses (Responses API).
Set ANTHROPIC_MODEL to a model your ChatGPT subscription is allowed to use.
Append -fast to a Codex model name to request Codex fast mode for that request
without restarting the proxy. For example, gpt-5.4-fast[1m] is sent upstream as
model gpt-5.4 with service_tier: "priority" . An explicit
codex.serviceTier / CCP_CODEX_SERVICE_TIER override still takes precedence.
Reasoning effort: Claude Code's output_config.effort value (the one you see in
the UI as ◐ medium · /effort ) is forwarded as Codex reasoning.effort ( low
/ medium / high / xhigh ). Claude Code's max value is sent upstream as
xhigh . An explicit codex.effort / CCP_CODEX_EFFORT override still takes
precedence and can also force none .
If the resolved model isn't supported by your account, upstream returns a 400
like
"The 'gpt-4.1' model is not supported when using Codex with a ChatGPT account." .
The proxy surfaces that verbatim.
Upstream: https://api.kimi.com/coding/v1/chat/completions (OpenAI-style
chat-completions).
Only one wire model is exposed: kimi-for-coding (its display name in kimi-cli
is Kimi-k2.6 , 256k context, supports reasoning + image input + video input).
kimi-k2.6 and k2.6 are accepted as aliases for the same wire id.
Reasoning effort: Claude Code's output_config.effort value (the one you see in
the UI as ◐ medium · /effort ) is forwarded as Kimi's reasoning_effort ( low
/ medium / high ). Thinking blocks from the upstream model are forwarded to
Claude Code and rendered as thinking content. If Claude Code disables thinking,
the proxy drops both reasoning_effort and the thinking: {type: "enabled"}
flag before forwarding.
sequenceDiagram
autonumber
participant CC as Claude Code
participant P as claude-code-proxy
participant AUTH as OAuth host<br/>(auth.openai.com or<br/>auth.kimi.com)
participant U as Upstream API<br/>(chatgpt.com/codex or<br/>api.kimi.com)
Note over P,AUTH: One-time: PKCE / device OAuth<br/>tokens cached locally for reuse
CC->>P: POST /v1/messages (Anthropic shape, stream: true)
alt access token expiring
P->>AUTH: POST /oauth/token (refresh_token)
AUTH-->>P: new access (+ rotated refresh)
end
P->>P: translate request<br/>• strip Anthropic-only fields<br/>• system blocks → instructions / system message<br/>• tool_use / tool_result ↔ provider-specific shapes<br/>• prompt_cache_key = session id
P->>U: POST upstream<br/>Bearer + provider-specific headers
U-->>P: provider SSE<br/>(Codex: output_item.*, output_text.delta, …)<br/>(Kimi: chat.completion.chunk, reasoning_content, …)
P->>P: reducer: typed events<br/>(thinking / text / tool start/delta/stop, finish)
P-->>CC: Anthropic SSE<br/>(message_start, content_block_*, message_delta, message_stop)
Loading
Commands
Command
Description
serve
Start the proxy on PORT
codex auth login / device / status / logout
Codex OAuth management
kimi auth login / status / logout
Kimi OAuth management
serve
Starts the HTTP proxy and blocks. Binds to 127.0.0.1 only. Logs to the
platform state directory (rotated at 20 MiB). Set CCP_LOG_STDERR=1 to mirror
log lines to stderr while running.
claude-code-proxy serve
PORT=11435 claude-code-proxy serve
CCP_LOG_STD

[truncated]
