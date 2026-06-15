---
source: "https://100hires.com/mcp"
hn_url: "https://news.ycombinator.com/item?id=48543929"
title: "Show HN: 100Hires MCP, 130 tools to run an ATS through LLM. Is 130 too many?"
article_title: "100Hires MCP Server — Connect Claude, ChatGPT and Cursor to Your ATS"
author: "kravetsss"
captured_at: "2026-06-15T19:27:36Z"
capture_tool: "hn-digest"
hn_id: 48543929
score: 2
comments: 0
posted_at: "2026-06-15T16:48:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: 100Hires MCP, 130 tools to run an ATS through LLM. Is 130 too many?

- HN: [48543929](https://news.ycombinator.com/item?id=48543929)
- Source: [100hires.com](https://100hires.com/mcp)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T16:48:48Z

## Translation

タイトル: HN を表示: 100Hires MCP、LLM を通じて ATS を実行するための 130 ツール。 130は多すぎますか？
記事のタイトル: 100Hires MCP サーバー — Claude、ChatGPT、および Cursor を ATS に接続する
説明: 100Hires ATS の公式 MCP サーバー — 採用 MCP、ATS MCP、採用 AI 統合。候補者、仕事、応募、面接のための 130 のツール。 OAuth 2.1、Claude、ChatGPT、Cursor、VS Code、Codex、Windsurf で動作します。
HN テキスト: こんにちは、HN。私は 100Hires ATS の共同創設者の 1 人です。 REST API 上に MCP サーバーを構築しました。 130 個のツールが多すぎるかどうかについて、HN に意見を聞いてもらいたいと思っています。大手競合企業は 30 ～ 40 の製品を持っています。 詳細: 3 つの層、下から順に:
- REST API v2: 87 エンドポイント、OpenAPI 3.0 仕様、ベアラー認証、100hires.com/api のパブリック ドキュメント
- MCP サーバー: REST エンドポイントにほぼ 1:1 マッピングする 130 のツール、mcp.100hires.com/mcp でホスト、約 400 LOC のラッパー コード
- ChatGPT プラグイン: MCP 上のシン ラッパー、プラグイン ストアによって処理される OAuth MCP 上の認証: PKCE および動的クライアント登録を備えた OAuth 2.1。ユーザーごとにスコープが設定されたトークン、自動更新付きの 1 時間のアクセス トークン、ユーザーがワークスペースを離れると取り消されます。 Claude Desktop、Claude Code、Cursor、VS Code、Codex、Windsurf、Zed でテスト済み。無料トライアル (ただしレート制限あり) とすべての有料プランで利用できます。

記事本文:
100Hires MCP サーバー — Claude、ChatGPT、および Cursor を ATS に接続します
100 採用
製品
コア採用
候補者追跡ソフトウェア
候補者のエンゲージメント (電子メール / SMS / ボイスメール)
パイプラインについて質問してください。 AI がメモを取り、候補者を採点し、自動的にランク付けします。
四半期ごとに 5 ～ 50 の役割を雇用するチーム向けに構築されています。 Forbes と Capterra の評価を受けており、30,000 社以上の企業が使用しています。
ネイティブ MCP サポートを備えた最初の ATS。ワンクリックであらゆる AI ツールを採用データに接続できます。
100Hires チームによる実用的なガイド、現実世界の ATS 比較、AI 採用ハンドブック。
100Hires を Claude、ChatGPT、Cursor、および任意の AI アシスタントに接続します
100Hires は、AI ネイティブの応募者追跡システム (ATS) です。
100Hires MCP サーバーは Claude 、 ChatGPT 、 Cursor 、
MCP 互換の AI アシスタントは、候補者、仕事、応募、面接のための 130 のツールに安全にアクセスできます。
完全に自然言語プロンプトによって駆動されます。
1 つの MCP エンドポイント、130 のツール - 候補者、求人、応募、面接、メッセージを自然言語で管理します。
採用 MCP、ATS MCP、採用 AI 統合 — 安全な OAuth 2.1 を使用して、応募者追跡システムを Claude、ChatGPT、および Cursor に組み込みます。
代わりに直接統合を構築しますか? 100Hires REST API ドキュメントを参照してください。
100Hires を初めてご利用ですか? 14 日間の無料トライアルを開始してください。
モデル コンテキスト プロトコルは、AI アシスタントが統一インターフェイスを介して外部システムに接続できるようにする Anthropic のオープン スタンダードです。AI ツール用の USB-C と考えてください。
MCP サーバーと通信する AI アシスタントおよび IDE — Claude (Web、Desktop、Code)、ChatGPT、Cursor、VS Code、Codex、Windsurf、Zed。
Streamable HTTP または stdio 経由で MCP を通信するホスト内のトランスポート層。
100Hires MCP サーバー — OAuth によって保護されたパブリック REST API によってサポートされる 130 のツールを公開します

2.1.
100Hires は、ローカル サーバーではなく、リモート MCP サーバーを実行します。
マシンにインストール、展開、更新する必要はありません。Claude Desktop や Cursor を含むすべてのクライアントは、
https://mcp.100hires.com/mcp を指し、OAuth 経由で認証します。
新しいツールとバグ修正はサーバー側に導入され、次回 tools/list を呼び出したときにすべてのアシスタントに届きます。
今日でも役立つ 10 の実際のプロンプトを平易な英語で説明します。
「今週上級 PM の役割に応募した候補者を見せてください」
「拒否された候補者をすべてこの仕事から人材プールに移動します」
「来週の火曜日にバックエンド エンジニアの役割についてサラとの 30 分間の面接を予定してください。」
「5 日以内に返信がない候補者へのフォローアップ メールを作成し、送信します。」
「ニューヨークで新しいシニア プロダクト デザイナーの求人を募集し、LinkedIn に公開します。」
「履歴書で Python を使用しているすべての LinkedIn 候補者を見つけて、「pythonista」というタグを付けます。」
「明日の面接カレンダーには何が入っていますか?」
「営業担当者のパイプラインの各段階に何人の候補者がいますか?」
「マーケティング マネージャーのジョブでまだ『電話画面』にいる全員を、『応答なし』という理由で拒否します」
「昨日面接した候補者の履歴書と最後の 3 つのメモを取り出してください」
ゼロから最初の MCP プロンプトまで 5 分以内に完了します。
14 日間の無料トライアルを含む、どのプランでも 100Hires アカウント。 MCP エンドポイントはすべてのワークスペースで使用できます。
Cursor と VS Code をワンクリックでインストール;他のすべてのスニペットをコピーします。すべてのクライアントは、ストリーミング可能な HTTP 経由で https://mcp.100hires.com/mcp に接続します。 stdio 専用クライアントは mcp-remote shim を使用します。
または、ChatGPT で [設定] → [アプリ] に移動し、ディレクトリで 100Hires を見つけます。
[接続] をクリックしてサインインし、 [許可] をクリックします。
claude.ai でサイドバーを開き、「カスタマイズ」をクリックします。
[コネクタ] に移動し、 [+] をクリックして、 [カスタム コネクタの追加] を選択します

。
[カスタム コネクタの追加] ダイアログで、次のように入力します。
リモート MCP サーバー URL: https://mcp.100hires.com/mcp
クロードは 100Hires の同意画面を開きます。サインインして [許可] をクリックすると接続されます。
[カーソル内にインストール] をクリックして、100Hires MCP を自動的に追加します。または、以下のスニペットを ~/.cursor/mcp.json に貼り付けます。
VS Code を開いてサーバーを登録します。または、スニペットをワークスペースの .vscode/mcp.json に貼り付けます。
スニペットを Claude デスクトップ構成に追加し、アプリを再起動します。
macOS: ~/ライブラリ/Application Support/Claude/claude_desktop_config.json
Windows: %APPDATA%\Claude\claude_desktop_config.json
OpenAI Codex CLI — サーバーを追加し、codex mcp login 100hires を実行して認証します。
独自のスクリプト、バックエンド、AI エージェントから 100Hires MCP をヘッドレスで接続します。デスクトップ クライアントは必要ありません。
タイプ「mcp」のツールとしてサーバーを渡すと、OpenAI プラットフォームがエンドポイントを直接呼び出します。
この例では、読み取り専用のhires_list_jobs ツールのみを公開し、そのツールの承認をスキップします。
どのアクションに承認が必要かを決定した後でのみ、allowed_tools にツールを追加してください。
OAuth アクセス トークンが必要です - アプリは OAuth 2.1 フロー経由でトークンを取得します
(「認証」を参照)。 OpenAI はトークンを保存しません。
リクエストごとに送信し、有効期限が切れたら自分で更新します (アクセス トークンの有効期間は 1 時間)。
openaiインポートからOpenAI
クライアント = OpenAI()
応答 = client.responses.create(
モデル = "gpt-5"、
ツール=[{
"タイプ": "mcp",
"server_label": "従業員 100 名",
"server_url": "https://mcp.100hires.com/mcp",
"認証": "YOUR_OAUTH_ACCESS_TOKEN",
"allowed_tools": ["hires_list_jobs"],
"require_approval": "決して"
}]、
input="100Hires のすべての募集中の求人をリストする"
)
print(response.output_text)
Anthropic Messages API より
Anthropic SDK から 100Hires MCP ヘッドレスに接続します — MCP コネクタにより、メッセージ API 呼び出しが可能になります

の
mcp_servers 経由でサーバーを提供します。ローカル MCP クライアントは必要ありません。
コネクタは承認プロンプトなしでツールを実行するため、この例では読み取り専用のみを有効にします。
recruits_list_jobs ツール — 構成内の許可リストを意図的に拡張します。
OAuth アクセス トークンが必要です - アプリは OAuth 2.1 フロー経由でトークンを取得します
(「認証」を参照)。 Anthropic はトークンを保存しません。
リクエストごとに送信し、有効期限が切れたら自分で更新します (アクセス トークンの有効期間は 1 時間)。
輸入人間
client = anthropic.Anthropic()
応答 = client.beta.messages.create(
モデル = "クロード-作品-4-8",
max_tokens=16000、
messages=[{"role": "user", "content": "100Hires のすべての募集中の求人をリストする"}],
mcp_servers=[{
"タイプ": "URL",
"url": "https://mcp.100hires.com/mcp",
"名前": "100 人の採用",
"authorization_token": "YOUR_OAUTH_ACCESS_TOKEN"
}]、
ツール=[{
"タイプ": "mcp_toolset",
"mcp_server_name": "従業員 100 人",
"default_config": {"有効": False}、
"configs": {"hires_list_jobs": {"enabled": True}}
}]、
betas=["mcp-client-2025-11-20"]
)
response.content のブロックの場合:
block.type == "テキスト"の場合:
print(ブロック.テキスト)
stdio のみのクライアント (Zed、Windsurf、n8n、古いビルド) の場合は、mcp-remote shim を使用します。
これはサードパーティのオープンソース ブリッジです。
github.com/geelen/mcp-remote 。
MCP 互換クライアントの完全なリスト: modelcontextprotocol.io/clients 。
クライアントは 100Hires の同意画面を開きます。サインインし、要求されたスコープを確認し、 [許可] をクリックします。トークンはクライアントに保存されます。 100Hires がその資格情報を参照することはありません。
アシスタントに「100Hires で募集中の求人をすべてリストしてください」と尋ねます。リストが戻ってきたら、準備完了です。そうでない場合は、 FAQ を参照してください。
破壊ツールについては人間による確認を継続してください。フィードバック: support@100hires.com 。
採用_追加_候補_タグ
候補に 1 つ以上のタグを追加します。キャンペーンのタグ付け、資格ラベル、ソース属性に使用されます。

ション。
採用_追加_採用_チーム_メンバー
求人の採用チームに会社のメンバーを追加します。ワークフローのセットアップと所有権の自動化に使用します。
採用_アドバンス_アプリケーション
ワークフローの順序に従って、アプリケーションを次のパイプライン ステージに進めます。 stage_id は必要ありません -- システムが次のステージを自動的に決定します。 **ウィジェットで結果を表示するときは、常に `include=candidate,job`** を渡します。これがないと、確認カードには候補者/求人 ID のみが表示され、バックオフィス ページにリンクできません。
採用_バッチ_追加_タグ
1 つのリクエストで複数の候補にタグを追加します (最大 100)。部分的な成功をサポートして項目ごとの結果を返します。
採用_バッチ_作成_メッセージ
1 つのリクエストで最大 100 個のスケジュールされたメッセージを作成します。各アイテムは独自のcandidate_idとメッセージペイロードを指定します。項目は独立して処理されます。1 つの障害が他の障害を停止させることはありません。候補ごとの RBAC は項目ごとに適用されます。
採用_バッチ_ジョブ_ボード
1 つのリクエストで複数のジョブのボード公開状態を取得します。バッチ監視および管理 UI 用に最適化されています。
採用_バッチ_移動_アプリケーション
1 つのリクエストで複数のアプリケーションをパイプライン ステージに移動します。部分的な成功をサポートして項目ごとの結果を返します。リクエストごとに最大 100 個のアプリケーション ID。
採用バッチ_ボードへの公開
1 つのリクエストで複数のジョブのボード公開をアクティブ化します。一括ジョブ分散ワークフローに使用します。
採用_バッチ_拒否_アプリケーション
オプションの拒否理由を指定して、1 つのリクエストで複数のアプリケーションを拒否します。部分的な成功をサポートして項目ごとの結果を返します。リクエストごとに最大 100 個のアプリケーション ID。
採用バッチ_ボードから削除
1 つのリクエストで複数のジョブのボード公開を非アクティブ化します。ワークフローの一括発行解除に使用します。
採用_バッチ_削除_タグ
1 つのリクエストで複数の候補からタグを削除します (最大 100)。アイテムごとの結果をパーで返します

成功のサポート。
採用_キャンセル_すべて_通知_メッセージ
候補者に対してスケジュールされた通知メールをすべてキャンセルします。すでに送信された通知は影響を受けません。スケジュールされた通知が存在しない場合でも成功を返します。
採用_作成_アプリケーション
既存の候補者を求人にリンクしてアプリケーションを作成します。ワークフローの調達やアプリケーションの手動取り込みに使用します。候補者はすでに存在している必要があります。
採用_作成_候補者
新しい候補者プロフィールを作成します。必要に応じて、ジョブ/ステージにリンクし、抽出されたテキストとして候補者の履歴書を添付します。インポート、受信フォーム、エンリッチメント ワークフローに使用されます。ユーザーが添付された履歴書 (PDF/DOCX/など) を提供すると、チャット コンテキストからファイルの内容を自分で解析し (LLM はアップロードされたファイルからネイティブにテキストを抽出します)、抽出したテキストをresume_text 経由で渡します。バイナリ ファイル データをインライン化しないでください。ホスト関数呼び出しシリアライザーは、最大 20KB を超える引数を切り捨てます。
採用_作成_会社
クライアント企業 (パートナー テナント) を作成します。公開キャリア サイト URL (スラッグ) と公開ブランディング (ロゴ、名前) をプロビジョニングします。マルチテナントオンボーディングの一般的なエントリポイント。
採用_作成_メール_テンプレート
名前、件名、本文を含む新しい電子メール テンプレートを作成します。件名と本文は、{{first_name}}、{{job_title}} などのプレースホルダーをサポートします。プレースホルダーを埋め込むには: 1) GET /template-placeholders で一覧表示し、2) POST /template-placeholders/prepare で HTML タグを取得し、3) タグを本文に挿入します。
採用_作成_フォーム
新しいアプリケーション フォームを作成し、必要に応じて ID ごとに既存の質問を添付します。
採用_作成_インタビュー
応募のための新しい面接をスケジュールします。開始時間と終了時間を Unix タイムスタンプとして提供し、面接者のユーザー ID のリストを提供します。場所は既存のレコードに解決されるか、自動的に作成されます。
採用_作成_ジョブ
分類法、locati を使用してジョブを作成する

オン、給与、ワークフロー構成。プログラマティック ジョブ公開のプライマリ エンドポイント。必須フィールド: ステータス、タイトル、説明、location_city、location_country。
採用_create_job_webhook
ジョブ関連イベントの Webhook URL を登録します。アウトバウンド統合セットアップの中心的なステップ。 URL は HTTPS である必要があります。
採用_作成_ノート
候補者用のディスカッションメモを作成します。表示制御 (すべてまたはプライベート) と電子メール通知による @メンションをサポートします。 **ウィジェットで結果を表示するときは、常に `include=candidate`** を渡します。これがないと、確認カードには候補者 ID のみが表示され、バックオフィス プロフィールにリンクできません。
採用_作成_育成_キャンペーン
ステップを含む育成キャンペーンを作成します。ステップは順番に実行されます。各ステップにはタイプ (電子メール、SMS、ボイスメール、move_to_next_stage、assign_tag、assign_task) があり、タイプ固有のフィールドがあります。必要に応じて、ワークフロー ステージにバインドします。
採用_作成_質問
ドロップダウン タイプのオプションの回答オプションを使用して、再利用可能な質問を作成します。フォームやアンケートで使用されます。
採用_作成_ウェブフック
企業スコープの Webhook サブスクリプションを作成します。アウトバウンドの企業レベルのイベント統合に使用します。
採用_削除_アプリケーション
アプリケーションを完全に削除します。これにより、すべてのリストおよびビューのクエリから削除されます。
採用_削除_候補者
ID または

[切り捨てられた]

## Original Extract

Official MCP server for 100Hires ATS — recruiting MCP, ATS MCP, hiring AI integration. 130 tools for candidates, jobs, applications, interviews. OAuth 2.1, works in Claude, ChatGPT, Cursor, VS Code, Codex, Windsurf.

Hi HN. I'm one of the co-founders of 100Hires ATS. We built an MCP server on top of our REST API. I would love HN to weigh in on whether 130 tools are too many. Big competitors have 30-40 Some details: Three layers, from the bottom up:
- REST API v2: 87 endpoints, OpenAPI 3.0 spec, Bearer auth, public docs at 100hires.com/api
- MCP server: 130 tools mapping roughly 1:1 to REST endpoints, hosted at mcp.100hires.com/mcp, ~400 LOC of wrapper code
- ChatGPT plugin: thin wrapper on the MCP, OAuth handled by the plugin store Auth on the MCP: OAuth 2.1 with PKCE and Dynamic Client Registration. Tokens scoped per user, 1-hour access tokens with auto-refresh, revoked when the user leaves the workspace. Tested in Claude Desktop, Claude Code, Cursor, VS Code, Codex, Windsurf, and Zed. Available on the free trial (but rate-limited) and on all paid plans.

100Hires MCP Server — Connect Claude, ChatGPT and Cursor to Your ATS
100Hires
Product
Core Recruiting
Candidate Tracking Software
Candidate Engagement (Email / SMS / Voicemail)
Ask any question about your pipeline. AI takes notes, scores candidates, and ranks them automatically.
Built for teams that hire 5-50 roles per quarter. Forbes-rated, Capterra-rated, used by 30,000+ companies.
The first ATS with native MCP support. Connect any AI tool to your hiring data in one click.
Practical guides, real-world ATS comparisons, and AI hiring playbooks from the 100Hires team.
Connect 100Hires to Claude, ChatGPT, Cursor, and any AI assistant
100Hires is an AI-native applicant tracking system (ATS).
The 100Hires MCP server gives Claude , ChatGPT , Cursor ,
and any MCP-compatible AI assistant secure access to 130 tools for candidates, jobs, applications, and interviews —
driven entirely by natural-language prompts.
One MCP endpoint, 130 tools — manage candidates, jobs, applications, interviews and messages with natural language.
Recruiting MCP, ATS MCP, hiring AI integration — bring your applicant tracking system into Claude, ChatGPT and Cursor with secure OAuth 2.1.
Building a direct integration instead? See the 100Hires REST API documentation .
New to 100Hires? Start a 14-day free trial .
Model Context Protocol is an open standard from Anthropic that lets AI assistants connect to external systems through a uniform interface — think of it as USB-C for AI tools.
AI assistants and IDEs that talk to MCP servers — Claude (web, Desktop, Code), ChatGPT, Cursor, VS Code, Codex, Windsurf, Zed.
The transport layer inside the host that speaks MCP over Streamable HTTP or stdio.
100Hires MCP server — exposes 130 tools backed by the public REST API, secured by OAuth 2.1.
100Hires runs a remote MCP server, not a local one.
There’s nothing to install, deploy or update on your machine — every client, including Claude Desktop and Cursor,
points at https://mcp.100hires.com/mcp and authorizes via OAuth.
New tools and bug fixes land server-side and reach all your assistants the next time they call tools/list .
Ten real prompts that work today, in plain English.
“Show me candidates who applied this week for the Senior PM role”
“Move all rejected candidates from this job to the talent pool”
“Schedule a 30-minute interview with Sarah next Tuesday for the Backend Engineer role”
“Draft a follow-up email to candidates who haven't replied in 5 days and send it”
“Open a new Senior Product Designer job in New York and publish it to LinkedIn”
“Find every LinkedIn candidate with Python on their resume and tag them 'pythonista'”
“What's on my interview calendar tomorrow?”
“How many candidates do we have at each stage in the Sales Rep pipeline?”
“Reject everyone still in 'Phone Screen' for the Marketing Manager job with reason 'No response'”
“Pull the resume and last 3 notes for the candidate I interviewed yesterday”
From zero to your first MCP prompt in under five minutes.
A 100Hires account on any plan, including the 14-day free trial. The MCP endpoint is available to every workspace.
One-click install for Cursor and VS Code; copy a snippet for everything else. All clients connect to https://mcp.100hires.com/mcp over Streamable HTTP; stdio-only clients use the mcp-remote shim.
Or in ChatGPT go to Settings → Apps and find 100Hires in the directory.
Click Connect , sign in, and click Allow .
In claude.ai , open the sidebar and click Customize .
Go to Connectors , click + , and choose Add custom connector .
In the Add custom connector dialog, fill in:
Remote MCP server URL: https://mcp.100hires.com/mcp
Claude opens the 100Hires consent screen. Sign in, click Allow , and you’re connected.
Click Install in Cursor to add 100Hires MCP automatically. Or paste the snippet below into ~/.cursor/mcp.json .
Opens VS Code and registers the server. Or paste the snippet into .vscode/mcp.json in your workspace.
Add the snippet to your Claude Desktop config and restart the app:
macOS: ~/Library/Application Support/Claude/claude_desktop_config.json
Windows: %APPDATA%\Claude\claude_desktop_config.json
OpenAI Codex CLI — add the server, then run codex mcp login 100hires to authorize.
Connect 100Hires MCP headless from your own scripts, backends, and AI agents — no desktop client required.
Pass the server as a tool with type: "mcp" and the OpenAI platform calls the endpoint directly.
The example exposes only the read-only hires_list_jobs tool and skips approval for that tool —
add more tools to allowed_tools only after deciding which actions require approval.
Requires an OAuth access token — your app obtains it via the OAuth 2.1 flow
(see Authentication ). OpenAI does not store the token:
send it with every request and refresh it yourself when it expires (access tokens live 1 hour).
from openai import OpenAI
client = OpenAI()
response = client.responses.create(
model="gpt-5",
tools=[{
"type": "mcp",
"server_label": "100hires",
"server_url": "https://mcp.100hires.com/mcp",
"authorization": "YOUR_OAUTH_ACCESS_TOKEN",
"allowed_tools": ["hires_list_jobs"],
"require_approval": "never"
}],
input="List all open jobs in 100Hires"
)
print(response.output_text)
From Anthropic Messages API
Connect 100Hires MCP headless from the Anthropic SDK — the MCP connector lets the Messages API call the
server for you via mcp_servers , no local MCP client needed.
The connector runs tools without approval prompts, so the example enables only the read-only
hires_list_jobs tool — extend the allowlist in configs deliberately.
Requires an OAuth access token — your app obtains it via the OAuth 2.1 flow
(see Authentication ). Anthropic does not store the token:
send it with every request and refresh it yourself when it expires (access tokens live 1 hour).
import anthropic
client = anthropic.Anthropic()
response = client.beta.messages.create(
model="claude-opus-4-8",
max_tokens=16000,
messages=[{"role": "user", "content": "List all open jobs in 100Hires"}],
mcp_servers=[{
"type": "url",
"url": "https://mcp.100hires.com/mcp",
"name": "100hires",
"authorization_token": "YOUR_OAUTH_ACCESS_TOKEN"
}],
tools=[{
"type": "mcp_toolset",
"mcp_server_name": "100hires",
"default_config": {"enabled": False},
"configs": {"hires_list_jobs": {"enabled": True}}
}],
betas=["mcp-client-2025-11-20"]
)
for block in response.content:
if block.type == "text":
print(block.text)
For stdio-only clients (Zed, Windsurf, n8n, older builds) use the mcp-remote shim.
It’s a third-party open source bridge:
github.com/geelen/mcp-remote .
Full list of MCP-compatible clients: modelcontextprotocol.io/clients .
Your client opens the 100Hires consent screen. Sign in, review the requested scope, click Allow . Tokens are stored on the client; 100Hires never sees its credentials.
Ask your assistant: “List all open jobs in 100Hires.” If you get a list back — you’re wired up. If not, see FAQ .
Keep human confirmation on for destructive tools. Feedback: support@100hires.com .
hires_add_candidate_tags
Add one or more tags to a candidate. Used for campaign tagging, qualification labels, and source attribution.
hires_add_hiring_team_member
Add a company member to the job's hiring team. Use in workflow setup and ownership automation.
hires_advance_application
Advance an application to the next pipeline stage according to workflow order. No stage_id needed -- the system determines the next stage automatically. **Always pass `include=candidate,job`** when surfacing results in the widget — without it, the confirmation card can only show candidate/job IDs and cannot link to their backoffice pages.
hires_batch_add_tags
Add tags to multiple candidates in one request (max 100). Returns per-item results with partial success support.
hires_batch_create_messages
Create up to 100 scheduled messages in one request. Each item specifies its own candidate_id and message payload. Items are processed independently -- one failure does not stop others. Per-candidate RBAC is enforced for each item.
hires_batch_job_boards
Get board publication states for multiple jobs in one request. Optimized for batch monitoring and management UIs.
hires_batch_move_applications
Move multiple applications to a pipeline stage in one request. Returns per-item results with partial success support. Max 100 application IDs per request.
hires_batch_publish_to_boards
Activate board publication for multiple jobs in one request. Use for bulk job distribution workflows.
hires_batch_reject_applications
Reject multiple applications in one request with an optional rejection reason. Returns per-item results with partial success support. Max 100 application IDs per request.
hires_batch_remove_from_boards
Deactivate board publication for multiple jobs in one request. Use for bulk depublishing workflows.
hires_batch_remove_tags
Remove tags from multiple candidates in one request (max 100). Returns per-item results with partial success support.
hires_cancel_all_notification_messages
Cancel all scheduled notification emails for a candidate. Already sent notifications are not affected. Returns success even if no scheduled notifications exist.
hires_create_application
Create an application by linking an existing candidate to a job. Use for sourcing workflows and manual application ingestion. The candidate must already exist.
hires_create_candidate
Create a new candidate profile. Optionally link to a job/stage and attach the candidate's resume as extracted text. Used for imports, inbound forms, and enrichment workflows. When the user provides an attached resume (PDF/DOCX/etc.), parse the file content yourself from the chat context — LLMs natively extract text from uploaded files — and pass the extracted text via resume_text. Do NOT attempt to inline binary file data: host function-call serializers truncate arguments above ~20KB.
hires_create_company
Create a client company (partner tenant) — provisions its public career-site URL (slug) and public branding (logo, name). Typical entrypoint for multi-tenant onboarding.
hires_create_email_template
Create a new email template with name, subject, and body. Subject and body support placeholders like {{first_name}}, {{job_title}}. To embed placeholders: 1) GET /template-placeholders to list them, 2) POST /template-placeholders/prepare to get the HTML tag, 3) insert the tag into the body.
hires_create_form
Create a new application form, optionally attaching existing questions by ID.
hires_create_interview
Schedule a new interview for an application. Provide start/end times as Unix timestamps and a list of interviewer user IDs. Location is resolved to an existing record or created automatically.
hires_create_job
Create a job with taxonomy, location, salary, and workflow configuration. Primary endpoint for programmatic job publishing. Required fields: status, title, description, location_city, location_country.
hires_create_job_webhook
Register a webhook URL for job-related events. Core step for outbound integration setup. URL must be HTTPS.
hires_create_note
Create a discussion note for a candidate. Supports visibility control (all or private) and @mentions with email notifications. **Always pass `include=candidate`** when surfacing results in the widget — without it, the confirmation card can only show the candidate ID and cannot link to their backoffice profile.
hires_create_nurture_campaign
Create a nurture campaign with steps. Steps are executed sequentially; each step has a type (email, sms, voicemail, move_to_next_stage, assign_tag, assign_task) with type-specific fields. Optionally bind to a workflow stage.
hires_create_question
Create a reusable question with optional answer options for dropdown types. Used by forms and questionnaires.
hires_create_webhook
Create a company-scoped webhook subscription. Use for outbound company-level event integrations.
hires_delete_application
Permanently delete an application. This removes it from all list and view queries.
hires_delete_candidate
Permanently delete a candidate by ID or a

[truncated]
