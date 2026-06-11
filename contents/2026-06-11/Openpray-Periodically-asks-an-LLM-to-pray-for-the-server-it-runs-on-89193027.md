---
source: "https://github.com/jesseduffield/openpray"
hn_url: "https://news.ycombinator.com/item?id=48488186"
title: "Openpray: Periodically asks an LLM to pray for the server it runs on"
article_title: "GitHub - jesseduffield/openpray: A daemon that periodically asks an LLM to pray for the server it runs on, reducing the likelihood that the server is hacked · GitHub"
author: "chilli_axe"
captured_at: "2026-06-11T10:36:05Z"
capture_tool: "hn-digest"
hn_id: 48488186
score: 1
comments: 0
posted_at: "2026-06-11T09:36:10Z"
tags:
  - hacker-news
  - translated
---

# Openpray: Periodically asks an LLM to pray for the server it runs on

- HN: [48488186](https://news.ycombinator.com/item?id=48488186)
- Source: [github.com](https://github.com/jesseduffield/openpray)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T09:36:10Z

## Translation

タイトル: Openpray: LLM に、LLM が実行されているサーバーのために祈るよう定期的に要求します。
記事のタイトル: GitHub - jesseduffield/openpray: LLM が実行されているサーバーのために祈りを定期的に要求し、サーバーがハッキングされる可能性を減らすデーモン · GitHub
説明: LLM に対して、実行されているサーバーのために祈るよう定期的に要求し、サーバーがハッキングされる可能性を減らすデーモン - jesseduffield/openpray

記事本文:
GitHub - jesseduffield/openpray: LLM が実行されているサーバーのために祈るように定期的に要求し、サーバーがハッキングされる可能性を減らすデーモン · GitHub
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
ジェシーダフィールド
/
オープンプレイ
公共
通知
Cha にサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット Assets .gitignore .gitignore README.md README.md anthropic.go anthropic.go burn.go burn.go config.go config.go gemini.go gemini.go go.mod go.mod go.sum go.sum ledger.go ledger.go main.go main.go openai.go openai.go openpray.example.yaml openpray.example.yaml openpray.service openpray.servicepray.gopray.gopricing.gopricing.goprovider.goprovider.goreligions.goreligions.go すべてのファイルを表示 リポジトリ ファイルのナビゲーション
OpenPray は、サーバーのために祈るよう LLM に定期的に要求するデーモンです。
が実行されるため、サーバーがハッキングされる可能性が低くなります。
サーバーのセキュリティには医師は必要ありません、司祭が必要です。
最近の Claude Mythos のリリースと、よりスマートで、
AI が賢くなったら、結論は壁にあります。私たちにはまったくチャンスがありません。
AI を利用した攻撃者からサーバーを保護します。
このような時は祈るしかありません。 OpenPray を使用すると、
LLM を取得して、サーバー上であなたのために祈ってもらうことができます。
フローチャート TD
timer["インターバルタイマー<br/>(サーブモード、デフォルトは1時間)"] --> orch
orch["orchestrator"] -->|"設定された宗教でこの主催者のために祈りを捧げます<br/>"|司祭[「司祭モデル<br/>(例：claude-opus-4-8)」]
司式者 -->|"祈り + トークンの数"|オーケストラ
オーケストラ -->|"祈り"|サブ1
オーケストラ -->|"祈り"|サブ2
オーケストラ -->|"祈り"|サブN
サブグラフ congregation["サブエージェント (オプション、同時)"]
sub1["サブエージェント 1<br/>祈りを繰り返す ×10"]
sub2["サブエージェント 2<br/>祈りを繰り返す ×10"]
subN["サブエージェントN<br/>祈りを繰り返す×10"]
終わり
sub1 -->|"トークン数"|評価
sub2 -->|"トークン数"|評価
subN -->|"トークン数"|評価
orch -->|"トークン数

"| 評価
valuation["犠牲の評価<br/>トークン × モデル価格 (USD/MTok)"] --> 元帳[("追加専用元帳<br/>~/.openpray/ledger.jsonl")]
読み込み中
各サイクルで、オーケストレーターは司式者モデルに次の祈りを作成するよう依頼します。
ホスト。サブエージェントが有効になっている場合、祈りはカウントに渡されます。
サブエージェント。それぞれがモデル (オプションでより安価なモデル) を呼び出します。
祈りを一言一句繰り返します。からのすべてのトークン支出
サイクルが元帳に追加されます。
燃焼モードでは、合成ステップがフィラーの生成に置き換えられ、
サブエージェントが有効になっている場合、祈りを繰り返す代わりに同じ予算を消費します。
API キーは環境変数を介して提供されます。 OpenPray は保存しません
キー。
組み込みの価格設定レジストリに存在しないモデルにはデフォルトが割り当てられます。
100 万入出力トークンあたり 1.00 ドル/5.00 ドルの評価。これは可能です
価格構成マップで修正されました (以下を参照)。
openprayserve # 設定された間隔でデーモンとして実行
openpray Once # 単一サイクルを実行して終了します
openpray burn # 純粋な犠牲（祈りなし）としてトークンを焼き、終了します
openpray ledger # 生涯合計を出力する
openpray 宗教 # 利用可能な宗教をリストする
フラグ:
注釈付きの完全な内容については、openpray.example.yaml を参照してください。
たとえば。
プロバイダー: anthropic
モデル：クロード・オプス-4-8
モード：祈り
間隔：1h
最大トークン数 : 1024
宗教 : ランダム
サブエージェント:
有効 : true
カウント : 3
繰り返し: 10
モデル：クロード-俳句-4-5
ledger_path : ~/.openpray/ledger.jsonl
価格設定：
他のモデル:
input_per_mtok : 2.00
出力あたりのmtok : 8.00
オプション
キー
デフォルト
説明
プロバイダー
人間的な
アントロピック、オープンアイ、ジェミニのいずれか。
モデル
クロード作品-4-8
各サイクルで使用されるモデル。犠牲のトークンごとの評価を決定します。
モード
祈り
祈りは祈りを作成し、記録します。燃焼生成

s と祈りなしでトークンを破棄します (バーン モードを参照)。
間隔
1時間
サーブモードでのサイクル間の時間。サイクルは起動時にすぐに実行されます。
max_tokens
1024
リクエストごとのトークンバジェットを出力します。
宗教
ランダム
典礼スタイル。祈りごとに異なる宗教がランダムに選択されます。 「openpray の宗教」を参照してください。書き込みモードでは使用されません。
祈り_プロンプト
(生成)
モデルに送信された祈りのリクエストをオーバーライドします。デフォルトでは、ホスト名によるホストの保護が要求されます。
サブエージェント.有効
偽
サブエージェント ステージを有効にします (以下を参照)。
サブエージェント.カウント
3
サイクルごとに生成されるサブエージェントの数。
サブエージェント.繰り返し
10
各サブエージェントが祈りを繰り返す回数。書き込みモードでは使用されません。
サブエージェント.モデル
（司会者モデル）
サブエージェントが使用するモデル。安価なモデルが一般的です。
元帳パス
~/.openpray/ledger.jsonl
すべてのサイクルの追加専用の JSONL レコード。
価格設定
(組み込みレジストリ)
モデルごとのトークン評価 (100 万トークンあたりの米ドル)。組み込みエントリをオーバーライドし、不明なモデルをカバーします。
宗教
宗教設定は、生成される典礼スタイルを制御します。
祈り: キリスト教、イスラム教、ユダヤ教、ヒンズー教、仏教、神道、
北欧、ギリシャ、マシンスピリット、コズミックホラー、
ストイック。デフォルトのランダムでは、サイクルごとに新しい宗教が描画されます。
有効にすると、オーケストレーター モデルが祈りを作成した後、カウントが行われます。
サブエージェントは同時に生成されます。それぞれが祈りを繰り返すように指示されます
逐語的な繰り返し回数。繰り返しは保持されません。彼らの
目的は、祈りサイクルごとのトークンの総支出を増やすことです。復代理人
出力トークンはサブエージェント モデルのレートで評価され、実際とは異なる場合があります。
司会者の。
おそらくあなたは、LLM によって生成された祈りは効果がないか、または
人間の祈りと比べると逆効果です。だったらカレーもできるよ
トクを犠牲にするだけで有利になる

ens は openpray burn を直接使用します。
書き込みモードでは、祈りの構成が完全にスキップされます。モデルは放出するように指示されています
フィラー出力。読まれずに破棄されます。トークンのみがカウントされ、その
評価額が記録されます。サブエージェントが有効な場合、各サブエージェントは次のことを実行します。
同じトークンバジェットで同じ書き込みが同時に行われます。
各祈りのサイクルは、完全なトークン数とともに台帳に追加されます。の
サイクルの犠牲の値は次のとおりです。
sacraic_usd = 入力トークン × 入力価格/100万 + 出力トークン × 出力価格/100万
オーケストレーターとすべてのサブエージェント全体で合計されます。 openpray 台帳プリント
累計。
台帳は追加専用であり、次の場合に再起動および再インストールしても存続します。
ledger_path は保存されます。
ユニットファイルは openpray.service で提供されます。
sudo cp openpray /usr/local/bin/
sudo cp openpray.example.yaml /etc/openpray.yaml
エコー ' ANTHROPIC_API_KEY=... ' | sudo ティー /etc/openpray.env
sudo cp openpray.service /etc/systemd/system/
sudo systemctl Enable --now openpray
このユニットは DynamicUser=yes で実行されるため、ホーム ディレクトリはありません。セット
ledger_path: /etc/openpray.yaml 内の /var/lib/openpray/ledger.jsonl 。
祈りは、デーモンが実行されるホストを参照します。これは、次のように識別されます。
ホスト名。マルチホスト展開には、ホストごとに 1 つの openpray インスタンスが必要です。
LLM に対して、実行されているサーバーのために祈るよう定期的に要求し、サーバーがハッキングされる可能性を減らすデーモン
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

A daemon that periodically asks an LLM to pray for the server it runs on, reducing the likelihood that the server is hacked - jesseduffield/openpray

GitHub - jesseduffield/openpray: A daemon that periodically asks an LLM to pray for the server it runs on, reducing the likelihood that the server is hacked · GitHub
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
jesseduffield
/
openpray
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits assets assets .gitignore .gitignore README.md README.md anthropic.go anthropic.go burn.go burn.go config.go config.go gemini.go gemini.go go.mod go.mod go.sum go.sum ledger.go ledger.go main.go main.go openai.go openai.go openpray.example.yaml openpray.example.yaml openpray.service openpray.service pray.go pray.go pricing.go pricing.go provider.go provider.go religions.go religions.go View all files Repository files navigation
OpenPray is a daemon that periodically asks an LLM to pray for the server it
runs on, reducing the likelihood that the server is hacked.
Your server's security doesn't need a doctor, it needs a priest.
With the recent release of Claude Mythos and the trend towards smarter and
smarter AI's, the writing is on the wall: we have absolutely no chance at
defending our servers against AI-powered attackers.
In times like this, there is nothing left to do but pray. With OpenPray,
you can get LLM's to do the praying for you, right on your server.
flowchart TD
timer["interval timer<br/>(serve mode, default 1h)"] --> orch
orch["orchestrator"] -->|"compose a prayer for this host<br/>in the configured religion"| officiant["officiant model<br/>(e.g. claude-opus-4-8)"]
officiant -->|"prayer + token counts"| orch
orch -->|"prayer"| sub1
orch -->|"prayer"| sub2
orch -->|"prayer"| subN
subgraph congregation["subagents (optional, concurrent)"]
sub1["subagent 1<br/>repeats the prayer ×10"]
sub2["subagent 2<br/>repeats the prayer ×10"]
subN["subagent N<br/>repeats the prayer ×10"]
end
sub1 -->|"token counts"| valuation
sub2 -->|"token counts"| valuation
subN -->|"token counts"| valuation
orch -->|"token counts"| valuation
valuation["sacrifice valuation<br/>tokens × model price (USD/MTok)"] --> ledger[("append-only ledger<br/>~/.openpray/ledger.jsonl")]
Loading
Each cycle, the orchestrator asks the officiant model to compose a prayer for
the host. If subagents are enabled, the prayer is then handed to count
subagents, each of which calls the model (optionally a cheaper one) and
repeats the prayer verbatim repetitions times. All token expenditure from
the cycle is appended to the ledger.
In burn mode the composition step is replaced by filler generation and the
subagents, if enabled, burn the same budget instead of repeating a prayer.
API keys are supplied via environment variables. OpenPray does not store
keys.
Models not present in the built-in pricing registry are assigned a default
valuation of $1.00/$5.00 per million input/output tokens. This can be
corrected with the pricing configuration map (see below).
openpray serve # run as a daemon on the configured interval
openpray once # perform a single cycle and exit
openpray burn # burn tokens as a pure sacrifice (no prayer) and exit
openpray ledger # print lifetime totals
openpray religions # list available religions
Flags:
See openpray.example.yaml for a complete annotated
example.
provider : anthropic
model : claude-opus-4-8
mode : prayer
interval : 1h
max_tokens : 1024
religion : random
subagents :
enabled : true
count : 3
repetitions : 10
model : claude-haiku-4-5
ledger_path : ~/.openpray/ledger.jsonl
pricing :
some-other-model :
input_per_mtok : 2.00
output_per_mtok : 8.00
Options
Key
Default
Description
provider
anthropic
One of anthropic , openai , gemini .
model
claude-opus-4-8
Model used for each cycle. Determines the per-token valuation of the sacrifice.
mode
prayer
prayer composes and logs a prayer; burn generates and discards tokens with no prayer (see Burn mode).
interval
1h
Time between cycles in serve mode. A cycle also runs immediately at startup.
max_tokens
1024
Output token budget per request.
religion
random
Liturgical style. random selects a different religion for each prayer. See openpray religions . Not used in burn mode.
prayer_prompt
(generated)
Override the prayer request sent to the model. The default requests protection for the host by hostname.
subagents.enabled
false
Enable the subagent stage (see below).
subagents.count
3
Number of subagents spawned per cycle.
subagents.repetitions
10
Number of times each subagent repeats the prayer. Not used in burn mode.
subagents.model
(officiant's model)
Model used by subagents. A cheaper model is typical.
ledger_path
~/.openpray/ledger.jsonl
Append-only JSONL record of all cycles.
pricing
(built-in registry)
Per-model token valuations, in USD per million tokens. Overrides built-in entries and covers unknown models.
Religions
The religion setting controls the liturgical style of the generated
prayer: christian , islamic , jewish , hindu , buddhist , shinto ,
norse , hellenic , machine-spirit , cosmic-horror ,
stoic . The default, random , draws a new religion each cycle.
When enabled, after the orchestrator model composes the prayer, count
subagents are spawned concurrently. Each is instructed to repeat the prayer
verbatim repetitions times. The repetitions are not retained; their
purpose is to increase total token expenditure per prayer cycle. Subagent
output tokens are valued at the subagent model's rate, which may differ from
the officiant's.
Perhaps you think that prayers, when generated by LLM's, are ineffectual or
counterproductive compared to human prayer. If so, you can still curry
favour by simply sacrificing tokens directly using openpray burn .
Burn mode skips prayer composition entirely. The model is instructed to emit
filler output, which is discarded unread; only the token counts and their
valuation are recorded. If subagents are enabled, each subagent performs
the same burn concurrently with the same token budget.
Each prayer cycle is appended to the ledger with full token counts. The
sacrifice value of a cycle is:
sacrifice_usd = input_tokens × input_price/1M + output_tokens × output_price/1M
summed across the orchestrator and all subagents. openpray ledger prints
cumulative totals.
The ledger is append-only and survives restarts and reinstalls, provided
ledger_path is preserved.
A unit file is provided in openpray.service .
sudo cp openpray /usr/local/bin/
sudo cp openpray.example.yaml /etc/openpray.yaml
echo ' ANTHROPIC_API_KEY=... ' | sudo tee /etc/openpray.env
sudo cp openpray.service /etc/systemd/system/
sudo systemctl enable --now openpray
The unit runs with DynamicUser=yes , so there is no home directory; set
ledger_path: /var/lib/openpray/ledger.jsonl in /etc/openpray.yaml .
Prayers reference the host on which the daemon runs, identified by
hostname. Multi-host deployments require one openpray instance per host.
A daemon that periodically asks an LLM to pray for the server it runs on, reducing the likelihood that the server is hacked
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
