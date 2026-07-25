---
source: "https://github.com/firasd/mcpclock"
hn_url: "https://news.ycombinator.com/item?id=49046289"
title: "MCP Clock: live public hosted tool for Claude, Codex, any MCP Client"
article_title: "GitHub - firasd/mcpclock · GitHub"
author: "firasd"
captured_at: "2026-07-25T11:02:16Z"
capture_tool: "hn-digest"
hn_id: 49046289
score: 1
comments: 1
posted_at: "2026-07-25T10:17:12Z"
tags:
  - hacker-news
  - translated
---

# MCP Clock: live public hosted tool for Claude, Codex, any MCP Client

- HN: [49046289](https://news.ycombinator.com/item?id=49046289)
- Source: [github.com](https://github.com/firasd/mcpclock)
- Score: 1
- Comments: 1
- Posted: 2026-07-25T10:17:12Z

## Translation

タイトル: MCP クロック: クロード、コーデックス、任意の MCP クライアント用のライブ パブリック ホスト ツール
記事タイトル: GitHub - firasd/mcp Clock · GitHub
説明: GitHub でアカウントを作成して、firasd/mcp Clock の開発に貢献します。

記事本文:
GitHub - firasd/mcp Clock · GitHub
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
フィラスド
/
マッククロック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 開く

その他のアクション メニュー フォルダーとファイル
35 コミット 35 コミット src src .gitignore .gitignore README.md README.md biome.json biome.json package-lock.json.bak package-lock.json.bak package.json package.json tsconfig.json tsconfig.json worker-configuration.d.ts worker-configuration.d.ts wrangler.jsonc wrangler.jsonc すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ライブ MCP エンドポイント: https://mcp Clock.firasd.workers.dev/mcp
Claude.ai、Claude iOS、Claude Code、OpenAI Codex、および任意の MCP 互換クライアントで動作します。
1 つ以上のタイムゾーンの現在時刻を返します。
timezones (オプション) — IANA ゾーン名の配列 (例: "America/New_York" ) に特殊なリテラル "UTC" と "Alphadec" を加えたもの。デフォルトは ["UTC"] です。最大15ゾーン。
offsetSeconds (オプション) — フォーマット前に適用される秒単位の符号付きオフセット。例えば。 24 時間前は -86400、1 分前は +60。
adec_canonical_only (オプション) — 「true」を渡すと、Alphadec 単位の説明が抑制され、正規文字列のみが返されます。
クロックゲット{}
Clock_get{"タイムゾーン": ["アジア/東京", "アメリカ/ニューヨーク"]}
Clock_get{"タイムゾーン": ["Alphadec"]、"adec_canonical_only": "true"}
時計の日情報
UTC 日付のカレンダー情報 (平日、月の日数、年間の日、年間進捗率 %、および ISO 週番号) を返します。
date (オプション) — YYYY-MM-DD 形式の日付。デフォルトは今日 (UTC) です。
時計の日の情報{}
Clock_day_info{"日付": "2025-09-09"}
時計変換
タイムスタンプを 1 つのタイム ゾーンから 1 つ以上のターゲット ゾーンに変換します。
source_zone — IANA ゾーン名または "UTC" 。
iso — ISO-8601 文字列。ソースが "UTC" の場合は、Z サフィックスを含めます。ソースが IANA ゾーンの場合は、オフセット (実時間) を省略します。
target_zones — IANA ゾーン名または "UTC" の配列。最小 1、最大 15。
時計_変換{
"source_zone" : " アメリカ/メキシコシティ " ,
"iso" : " 2025-10-28T07:30:00 " ,
「ターゲットゾーン

es" : [ " UTC " , " アジア/ドバイ " ]
}
クロック_変換_alphadec
UTC ISO タイムスタンプと Alphadec 文字列の間で変換します。
方向 — "utc_to_alphadec" (デフォルト) または "alphadec_to_utc" 。
value — UTC ISO 文字列 (例: "2025-06-09T16:30:00.000Z" ) または Alphadec 文字列 (例: "2025_L3T5_000000" )。
Clock_convert_alphadec{"方向": "utc_to_alphadec", "値": "2025-06-09T16:30:00.000Z"}
Clock_convert_alphadec{"方向": "alphadec_to_utc", "値": "2025_L3T5_000000"}
クロック_変換_unixtime
UTC ISO タイムスタンプと Unix タイムスタンプ (エポックからの秒数) の間で変換します。
方向 — "utc_to_unixtime" (デフォルト) または "unixtime_to_utc" 。
value — UTC ISO 文字列 (例: "2025-06-15T12:00:00Z" ) または Unix タイムスタンプ文字列 (例: "1749988800" )。
Clock_convert_unixtime{"方向": "utc_to_unixtime", "値": "2025-06-15T12:00:00Z"}
Clock_convert_unixtime{"方向": "unixtime_to_utc", "値": "1749988800"}
クロックデルタ_utc
2 つの UTC ISO タイムスタンプ間の時間差を計算します。現在の時刻 (つまり、「からの時間」または「までの時間」) を使用するには、どちらかの端を省略します。
start (オプション) — UTC ISO タイムスタンプ (例: "2022-01-15T10:30:00Z" )。デフォルトは現在です。
end (オプション) — UTC ISO タイムスタンプ。デフォルトは現在です。開始または終了の少なくとも 1 つは必須です。
total_seconds 、内訳 (年/日/時/分/秒)、および読み取り可能な文字列を返します。開始が終了より後である場合は負です。
Clock_delta_utc{"開始": "2022-01-15T10:30:00Z"}
Clock_delta_utc{"終了": "2025-12-31T23:59:59Z"}
Clock_delta_utc{"開始": "2022-01-15T10:30:00Z"、"終了": "2025-08-31T14:45:30Z"}
クロックデルタアルファデック
現在の年の 2 つの 4 文字 Alphadec コード間の時差を計算します。現在の時刻を使用するには、どちらかの端を省略します。
alphadec_start (オプション) — 4 文字の Alphadec コード (例: "A2B3" )。デフォルトは現在です。
アルプス

hadec_end (オプション) — 4 文字の Alphadec コード (例: "Z8Y9" )。デフォルトは現在です。少なくとも 1 つは必要です。
ISO 時間差と Alphadec 単位デルタ (ピリオド/アーク/バー/ビート) の両方を返します。
Clock_delta_alphadec{"alphadec_start": "A2B3"}
Clock_delta_alphadec{"alphadec_end": "Z8Y9"}
Clock_delta_alphadec{"alphadec_start": "A2B3", "alphadec_end": "C1Y9"}
アルファデック
Alphadec は、コンパクトで人間が判読できるタイムスタンプ形式です。完全な正規文字列は 2026_I2J9_382995 のようになります。
Alphadec 文字列は K ソート可能であり、辞書編集順は時系列順と一致します。より少ない文字に切り詰めると、自然な時間グループが作成されます (たとえば、2026_I2 は I2 アーク全体をカバーします)。
季節アンカー (おおよそ): 期間 F = 3 月分点、期間 M = 6 月至点、期間 S = 9 月分点、期間 Z = 12 月至点。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to firasd/mcpclock development by creating an account on GitHub.

GitHub - firasd/mcpclock · GitHub
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
firasd
/
mcpclock
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
35 Commits 35 Commits src src .gitignore .gitignore README.md README.md biome.json biome.json package-lock.json.bak package-lock.json.bak package.json package.json tsconfig.json tsconfig.json worker-configuration.d.ts worker-configuration.d.ts wrangler.jsonc wrangler.jsonc View all files Repository files navigation
Live MCP endpoint: https://mcpclock.firasd.workers.dev/mcp
Works in Claude.ai, Claude iOS, Claude Code, OpenAI Codex, and any MCP-compatible client.
Returns current time in one or more time zones.
timezones (optional) — Array of IANA zone names (e.g. "America/New_York" ) plus the special literals "UTC" and "Alphadec" . Defaults to ["UTC"] . Max 15 zones.
offsetSeconds (optional) — Signed offset in seconds applied before formatting. E.g. -86400 for 24 h ago, +60 for one minute ahead.
adec_canonical_only (optional) — Pass "true" to suppress the Alphadec unit explanation and return only the canonical string.
clock_get{}
clock_get{"timezones": ["Asia/Tokyo", "America/New_York"]}
clock_get{"timezones": ["Alphadec"], "adec_canonical_only": "true"}
clock_day_info
Returns calendar information for a UTC date — weekday, days in month, day of year, year progress %, and ISO week number.
date (optional) — Date in YYYY-MM-DD format. Defaults to today (UTC).
clock_day_info{}
clock_day_info{"date": "2025-09-09"}
clock_convert
Converts a timestamp from one time zone to one or more target zones.
source_zone — IANA zone name or "UTC" .
iso — ISO-8601 string. If source is "UTC" , include the Z suffix. If source is an IANA zone, omit the offset (wall-clock time).
target_zones — Array of IANA zone names or "UTC" . Min 1, max 15.
clock_convert{
"source_zone" : " America/Mexico_City " ,
"iso" : " 2025-10-28T07:30:00 " ,
"target_zones" : [ " UTC " , " Asia/Dubai " ]
}
clock_convert_alphadec
Converts between a UTC ISO timestamp and an Alphadec string.
direction — "utc_to_alphadec" (default) or "alphadec_to_utc" .
value — UTC ISO string (e.g. "2025-06-09T16:30:00.000Z" ) or Alphadec string (e.g. "2025_L3T5_000000" ).
clock_convert_alphadec{"direction": "utc_to_alphadec", "value": "2025-06-09T16:30:00.000Z"}
clock_convert_alphadec{"direction": "alphadec_to_utc", "value": "2025_L3T5_000000"}
clock_convert_unixtime
Converts between a UTC ISO timestamp and a Unix timestamp (seconds since epoch).
direction — "utc_to_unixtime" (default) or "unixtime_to_utc" .
value — UTC ISO string (e.g. "2025-06-15T12:00:00Z" ) or Unix timestamp string (e.g. "1749988800" ).
clock_convert_unixtime{"direction": "utc_to_unixtime", "value": "2025-06-15T12:00:00Z"}
clock_convert_unixtime{"direction": "unixtime_to_utc", "value": "1749988800"}
clock_delta_utc
Calculates the time difference between two UTC ISO timestamps. Omit either end to use the current time (i.e. "time since" or "time until").
start (optional) — UTC ISO timestamp (e.g. "2022-01-15T10:30:00Z" ). Defaults to now.
end (optional) — UTC ISO timestamp. Defaults to now. At least one of start or end is required.
Returns total_seconds , a breakdown (years/days/hours/minutes/seconds), and a readable string. Negative if start is after end.
clock_delta_utc{"start": "2022-01-15T10:30:00Z"}
clock_delta_utc{"end": "2025-12-31T23:59:59Z"}
clock_delta_utc{"start": "2022-01-15T10:30:00Z", "end": "2025-08-31T14:45:30Z"}
clock_delta_alphadec
Calculates the time difference between two 4-character Alphadec codes within the current year. Omit either end to use the current time.
alphadec_start (optional) — 4-character Alphadec code (e.g. "A2B3" ). Defaults to now.
alphadec_end (optional) — 4-character Alphadec code (e.g. "Z8Y9" ). Defaults to now. At least one is required.
Returns both ISO time difference and Alphadec unit delta (periods/arcs/bars/beats).
clock_delta_alphadec{"alphadec_start": "A2B3"}
clock_delta_alphadec{"alphadec_end": "Z8Y9"}
clock_delta_alphadec{"alphadec_start": "A2B3", "alphadec_end": "C1Y9"}
Alphadec
Alphadec is a compact, human-readable timestamp format. A full canonical string looks like 2026_I2J9_382995 .
Alphadec strings are K-sortable — lexicographic order matches chronological order. Truncating to fewer characters creates natural time groupings (e.g. 2026_I2 covers the entire I2 arc).
Seasonal anchors (approximate): Period F = March equinox · Period M = June solstice · Period S = September equinox · Period Z = December solstice.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
