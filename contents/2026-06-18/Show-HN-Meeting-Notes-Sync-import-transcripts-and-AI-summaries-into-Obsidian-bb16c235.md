---
source: "https://community.obsidian.md/plugins/meeting-notes-sync"
hn_url: "https://news.ycombinator.com/item?id=48582509"
title: "Show HN: Meeting Notes Sync – import transcripts and AI summaries into Obsidian"
article_title: "Meeting Notes Sync - Obsidian Plugin"
author: "andreagrandi"
captured_at: "2026-06-18T10:36:02Z"
capture_tool: "hn-digest"
hn_id: 48582509
score: 1
comments: 1
posted_at: "2026-06-18T08:34:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Meeting Notes Sync – import transcripts and AI summaries into Obsidian

- HN: [48582509](https://news.ycombinator.com/item?id=48582509)
- Source: [community.obsidian.md](https://community.obsidian.md/plugins/meeting-notes-sync)
- Score: 1
- Comments: 1
- Posted: 2026-06-18T08:34:11Z

## Translation

タイトル: HN を表示: 会議メモの同期 – トランスクリプトと AI 概要を Obsidian にインポート
記事のタイトル: 会議メモの同期 - Obsidian プラグイン
説明: 会議の記録、メモ、AI 概要を MacParakeet と Fellow からボールトに同期します。

記事本文:
会議メモの同期 - Obsidian プラグインの検索... プラグインとテーマの検索... ⌘K サインイン 開始する
フォローしてください Discord Twitter Bluesky Threads Mastodon YouTube GitHub © 2026 Obsidian Meeting Notes Sync
14 ダウンロード MacParakeet と Fellow から会議の記録、メモ、AI 概要を保管庫に同期します。
Obsidian プラグインは、MacParakeet と Fellow から会議の記録、メモ、AI 概要をボールトに同期し、会議ごとに 1 つのフォルダーを作成します。両方のソースによって記録された会議は、複製されるのではなく、単一のフォルダーに結合されます。
会議ごとに 1 つのフォルダー。フォルダーメモのインデックスと会議の成果物が含まれます。各アーティファクトはソース ( Summary (MacParakeet).md 、 Summary (Fellow).md ) によってタグ付けされているため、両方によってキャプチャされた会議は明確になります。
クロスソースのマージ : 両方のソースで記録された会議が 1 つのフォルダーに配置され、時間間隔の重複によって一致します (タイトルの類似性により関係が解消されます)。インデックスにはソース ID と会議間隔の両方が含まれます。不確実な一致には、「merge-confidence: low for review」のフラグが立てられます。
増分および非破壊: 各同期は変更されていない会議をスキップし、遅いフェローの要約が MacParakeet が既に作成したフォルダーにマージされます。ボールトはアーカイブです。ソース内の会議を削除してもボールトには影響せず、会議フォルダーに追加したファイルはそのまま残されます。
ミーティング/2026/06 - 6/2 - ウィークリースタンドアップ - 6 月 11 日/
2 - Weekly Standup - 6 月 11 日.md ← インデックス: macparakeet-id + フェロー ID + 間隔
概要 (MacParakeet).md
トランスクリプト (MacParakeet).md
メモ.md
概要（フェロー）.md
アクションアイテム (フェロー).md
成績証明書 (フェロー).md
情報源
各ソースは個別に切り替えられます。無効になったソースは完全に不活性であり、リクエスト、書き込み、通知はありません。
MacParakeet — ローカル CLI、オプトイン
macparakeet-cli から会議を読み取ります (ネットワークなし、AC なし)

カウントされますが、データベースへのアクセスはありません）。 MacParakeet がインストールされた macOS (CLI はアプリ内に同梱) またはスタンドアロン CLI ( brew install Moona3k/tap/macparakeet-cli )、および Obsidian デスクトップが必要です。プラグインは CLI (Homebrew パス、次にアプリバンドル) を自動検出します。必要に応じて設定でパスを上書きすると、ヘルスチェックでそれが確認されます。 [設定] → [会議メモの同期] で有効にします。
Polls Fellow のクラウド AI 用 REST 開発者 API の要約。セットアップ:
有料の Fellow プラン。ワークスペース管理者が API (ワークスペース設定 → セキュリティ) を有効にします。
個人用 API キーを生成します ( [ユーザー設定] → [開発者ツール] ) - 表示できる範囲に限定され、取り消し可能で、監査ログに記録されます。
[設定] → [ミーティング ノートの同期] で、[Fellow] を有効にし、サブドメイン ( acme.fellow.app の acme) とキーを入力し、[接続の確認] を使用して確認します。
⚠️ Fellow キーは data.json (API ベースのプラグインの標準) に平文で保存されます。ボールトが git または同期フォルダーにある場合は、そのファイルを除外し、キーをシークレットとして扱います。漏れた場合はFellowで取り消してください。
コミュニティストアのレビュー待ち — コミュニティストアに到着するまで、BRAT 経由でインストールするか (ベータプラグインを追加 → andreagrandi/obsidian-meeting-notes-sync )、または手動で main.js + manifest.json をリリースから <vault>/.obsidian/plugins/meeting-notes-sync/ にドロップし、プラグインを有効にします。ソースからビルドするには:
git clone https://github.com/andreagrandi/obsidian-meeting-notes-sync
cd obsidian-meeting-notes-sync && npm install && npm run build
# main.js + manifest.json を <vault>/.obsidian/plugins/meeting-notes-sync/ にコピーします
「MacParakeet Sync」からの移行
macparakeet-sync ID で以前のビルドを使用していた場合は、Obsidian を終了し、 <vault>/.obsidian/plugins/macparakeet-sync/ という名前に変更し、meeting-notes-sync/ を再度開き、再度有効にします。これにより、data.json (番号付け、スナップショット、ファイル) が保持されます。

ファイルの所有権）のため、次回の同期は再インポートではなく何も行われません。既存のフォルダーの名前は維持されます。新しいアーティファクトのみがソースサフィックスを取得します。
同期は起動直後および 30 分ごと (構成可能、0 は無効)、またはリボン アイコン / 今すぐ同期 を使用してオンデマンドで実行されます。主な設定:
ソース — MacParakeet および/または Fellow を有効にします (上記を参照)。
ベース フォルダー — 会議フォルダーが置かれる場所 (空 = ボールト ルートなので、既定のテンプレートの Meetings/… がルートです)。
パス テンプレート — 会議ごとのフォルダー パス (以下のトークンを参照)。
コンテンツ — AI 結果 (オン)、会議メモ (オン)、完全なトランスクリプト (オフ、トランスクリプトは長い)。あらゆるソースに適用されます。
次以降同期 — この日付以降の会議のみがインポートされます (デフォルト: インストール日)。バックフィル履歴に戻します。
{year} · {month} · {monthName} (6 月) · {monthShort} (6 月) · {day} · {dayOrdinal} (2 日) · {date} (YYYY-MM-DD) · {n} (月ごとの数値) · {title} 。不明なトークンはそのまま残されます。
デフォルト: Meetings/{year}/{month} - {monthName}/{n} - {title} - {monthShort} {dayOrdinal} → Meetings/2026/06 - June/4 - Core sync-discovery - Jun 2nd 。末尾の日付は、タイトルを共有する定期的な会議を明確にします。
重複しきい値 (短い会議の割合、デフォルトは 0.5) と最小重複 (分) は、2 つの情報源の会議が 1 つとして扱われる場合を制御します。より控えめにマージするには値を上げます。
番号付け ( {n} ) とフォルダー名は最初のインポート時に固定され、決して変更されません。後のソースは再番号付けされずにマージされます。
同期されたファイルはミラーであり、ソースが変更されると上書きされます。自分の考えを別のファイルに保存してください。コンテンツ タイプまたはソースの切り替えは、遡及的にではなく将来に適用されます。
失敗したソース (取り消された Fellow キーなど) は、他の通知をブロックせずに 1 つの通知を表示し、設定によって報告されます。

gs接続チェック。
各ソースは、共通のアダプターとソースに依存しないエンジンの背後にあります。 MacParakeet は macparakeet-cli (meeting list/show/results list --json) 経由で、Fellow は REST API ( /me 、 /recordings 、 /recording/{id} 、 /note/{id} ) 経由で updated_at 変更検出を備えています。会議は間隔の重複によってソース間でマージされ、同期状態 (番号付け、ファイルごとの所有権、ソースごとのスナップショット) は data.json に保存されます。詳細は PLAN.md にあります。
npmインストール
npm run dev # esbuild watch
npm テスト # vitest
npm run build # typecheck + プロダクションバンドル
CI はプッシュと PR ごとにビルドとテストを実行します。
Live Fellow テスト — npx -y tsx scripts/fellow-live-test.mts は、実際のエンジン (ローカル CLI + ライブ Fellow API) を一時ボールトで実行します。ハーネスについてのみ、.env ( .env.example を参照) から FELLOW_SUBDOMAIN / FELLOW_API_KEY を読み取ります。プラグイン自体は環境を読み取ることはありません。エンド ユーザーは設定 UI でサブドメインとキーを設定します。
リリース — npm run release <major.minor.patch> (例: npm run release 0.3.1 ) は package.json 、manifest.json 、versions.json 、および lockfile をバンプし、プレフィックスのない X.Y.Z タグをコミットして作成します。これをプッシュ ( git Push Origin master && git Push Origin <version> ) して .github/workflows/release.yml をトリガーします。これにより、 main.js のビルド来歴が証明され、 main.js + manifest.json が GitHub リリースに添付されます (Obsidian はこれら 2 つだけをダウンロードします。versions.json はリポジトリから読み取られます)。プレーンな npm 実行スクリプト (npm バージョンのライフサイクル フックではない) が使用されるため、グローバルのignore-scripts=true に関係なく機能します。 obsidianmd/obsidian-releases へのコミュニティ ストアの送信は、最初のリリース後に 1 回限りの手動手順です。
インポート フォルダーの同期 詳細 現在のバージョン 0.3.1 最終更新日 3 時間前 作成日 6 日前 更新数 6 リリース ダウンロード 14 Obsidi との互換性

1.5.0 + プラットフォーム デスクトップのみ ライセンス MIT バグの報告 機能の報告 プラグインの作成者 Andrea Grandi andreagrandi www.andreagrandi.it GitHub andreagrandi Community
ライブ カーソルを使用してリアルタイムでコラボレーションします。フォルダーを共有します。アップデートへのアクセスを管理します。
データを Obsidian で使用できる Markdown ファイルに変換します。 Apple Notes、OneNote、Evernote、Notion、Google Keep、その他多くの形式で動作します。
サーバー、モバイル、Web にわたるボールトのリアルタイム同期。誰とでも共有可能。 REST と MCP の統合をサポートし、個人用 AI ナレッジ ベースを構築します。
カスタムの並べ替え、固定、非表示を使用してファイル エクスプローラーを強化します。
複数のノートにプロパティを一度に追加します。フォルダーを右クリックするか、複数のノートを選択して選択内容を右クリックします。
Zotero から引用、参考文献、メモ、PDF 注釈を挿入およびインポートします。
Readwise からハイライトをボールトに同期します。
Notion で提供される機能と同様に、フォルダーを折りたたまずにアクセスできるフォルダー内にメモを作成します。
完全なカレンダー HUB エクスペリエンス。すべてのカレンダーを 1 か所で操作できます。時間を分析して行動を起こしましょう！
コンテナーを自己ホスト型サーバーまたは WEBRTC に安全に同期します。

## Original Extract

Sync meeting transcripts, notes, and AI summaries from MacParakeet and Fellow into your vault.

Meeting Notes Sync - Obsidian Plugin Search... Search plugins and themes... ⌘K Sign in Get started
Follow us Discord Twitter Bluesky Threads Mastodon YouTube GitHub © 2026 Obsidian Meeting Notes Sync
14 downloads Sync meeting transcripts, notes, and AI summaries from MacParakeet and Fellow into your vault.
Obsidian plugin that syncs your meeting transcripts, notes, and AI summaries into your vault — from MacParakeet and Fellow , one folder per meeting. A meeting recorded by both sources merges into a single folder instead of duplicating.
One folder per meeting , with a folder-note index plus the meeting's artifacts. Each artifact is tagged by source ( Summary (MacParakeet).md , Summary (Fellow).md ), so a meeting captured by both stays unambiguous.
Cross-source merge : meetings recorded by both sources land in one folder, matched by time-interval overlap (title similarity breaks ties). The index carries both source ids and the meeting interval; uncertain matches are flagged merge-confidence: low for review.
Incremental & non-destructive : each sync skips unchanged meetings, and a late Fellow recap merges into the folder MacParakeet already made. The vault is your archive — deleting a meeting in a source never touches your vault, and files you add to a meeting folder are left alone.
Meetings/2026/06 - June/2 - Weekly Standup - Jun 11th/
2 - Weekly Standup - Jun 11th.md ← index: macparakeet-id + fellow-id + interval
Summary (MacParakeet).md
Transcript (MacParakeet).md
Notes.md
Summary (Fellow).md
Action Items (Fellow).md
Transcript (Fellow).md
Sources
Each source is toggled independently; a disabled source is completely inert — no requests, writes, or notices.
MacParakeet — local CLI, opt-in
Reads meetings from macparakeet-cli (no network, no accounts, no database access). Requires macOS with MacParakeet installed (the CLI ships inside the app) or the standalone CLI ( brew install moona3k/tap/macparakeet-cli ), and Obsidian desktop. The plugin auto-detects the CLI (Homebrew paths, then the app bundle); override the path in settings if needed, and the health check confirms it. Enable it in Settings → Meeting Notes Sync .
Polls Fellow's REST Developer API for cloud AI recaps. Setup:
A paid Fellow plan , with a workspace admin enabling the API ( Workspace Settings → Security ).
Generate a personal API key ( User Settings → Developer Tools ) — scoped to what you can see, revocable, audit-logged.
In Settings → Meeting Notes Sync : enable Fellow, enter your subdomain (the acme in acme.fellow.app ) and key , then use Check connection to verify.
⚠️ The Fellow key is stored in plaintext in data.json (standard for API-backed plugins). If your vault is in git or a synced folder, exclude that file and treat the key as a secret; revoke it in Fellow if it leaks.
Pending community-store review — until it lands there, install via BRAT ( Add beta plugin → andreagrandi/obsidian-meeting-notes-sync ), or manually drop main.js + manifest.json from a release into <vault>/.obsidian/plugins/meeting-notes-sync/ and enable the plugin. To build from source:
git clone https://github.com/andreagrandi/obsidian-meeting-notes-sync
cd obsidian-meeting-notes-sync && npm install && npm run build
# copy main.js + manifest.json into <vault>/.obsidian/plugins/meeting-notes-sync/
Migrating from "MacParakeet Sync"
If you used an earlier build under the macparakeet-sync id: quit Obsidian, rename <vault>/.obsidian/plugins/macparakeet-sync/ → meeting-notes-sync/ , reopen, and re-enable. This keeps data.json (numbering, snapshots, file ownership) so the next sync is a no-op instead of re-importing. Existing folders keep their names; only new artifacts get the source suffix.
Sync runs shortly after launch and every 30 minutes (configurable; 0 disables), or on demand via the ribbon icon / Sync now . The main settings:
Sources — enable MacParakeet and/or Fellow (see above).
Base folder — where meeting folders go (empty = vault root, so the default template's Meetings/… is the root).
Path template — folder path per meeting (see tokens below).
Content — AI results (on), meeting notes (on), full transcript (off; transcripts are long). Applies to every source.
Sync since — only meetings on/after this date import (default: install date). Move it back to backfill history.
{year} · {month} · {monthName} (June) · {monthShort} (Jun) · {day} · {dayOrdinal} (2nd) · {date} (YYYY-MM-DD) · {n} (per-month number) · {title} . Unknown tokens are left as-is.
Default: Meetings/{year}/{month} - {monthName}/{n} - {title} - {monthShort} {dayOrdinal} → Meetings/2026/06 - June/4 - Core sync-discovery - Jun 2nd . The trailing date disambiguates recurring meetings that share a title.
Overlap threshold (fraction of the shorter meeting, default 0.5) and minimum overlap (minutes) govern when two sources' meetings are treated as one. Raise them to merge more conservatively.
Numbering ( {n} ) and the folder name are frozen on first import and never change — a later source merges in without renumbering.
Synced files are mirrors and get overwritten when the source changes; keep your own thoughts in separate files. Toggling content types or sources on applies going forward, not retroactively.
A failing source (e.g. a revoked Fellow key) shows one Notice without blocking the others, and is reported by the settings connection check.
Each source sits behind a common adapter and a source-agnostic engine: MacParakeet via macparakeet-cli ( meetings list/show/results list --json ), Fellow via its REST API ( /me , /recordings , /recording/{id} , /note/{id} ) with updated_at change detection. Meetings are merged across sources by interval overlap, and sync state (numbering, per-file ownership, per-source snapshots) lives in data.json . Details in PLAN.md .
npm install
npm run dev # esbuild watch
npm test # vitest
npm run build # typecheck + production bundle
CI runs build + tests on every push and PR.
Live Fellow test — npx -y tsx scripts/fellow-live-test.mts runs the real engine (local CLI + live Fellow API) into a temp vault. It reads FELLOW_SUBDOMAIN / FELLOW_API_KEY from .env (see .env.example ) for the harness only; the plugin itself never reads the environment — end users configure the subdomain and key in the settings UI.
Releasing — npm run release <major.minor.patch> (e.g. npm run release 0.3.1 ) bumps package.json , manifest.json , versions.json , and the lockfile, then commits and creates the unprefixed X.Y.Z tag. Push it ( git push origin master && git push origin <version> ) to trigger .github/workflows/release.yml , which attests main.js 's build provenance and attaches main.js + manifest.json to a GitHub release (Obsidian downloads only those two; versions.json is read from the repo). A plain npm run script is used (not the npm version lifecycle hook) so it works regardless of a global ignore-scripts=true . Community-store submission to obsidianmd/obsidian-releases is a one-time manual step after the first release.
Syncing Import Folders Details Current version 0.3.1 Last updated 3 hours ago Created 6 days ago Updates 6 releases Downloads 14 Compatible with Obsidian 1.5.0 + Platforms Desktop only License MIT Report bug Request feature Report plugin Author Andrea Grandi andreagrandi www.andreagrandi.it GitHub andreagrandi Community
Collaborate in real time with live cursors. Share folders. Manage access to updates.
Convert your data to Markdown files you can use in Obsidian. Works with Apple Notes, OneNote, Evernote, Notion, Google Keep, and many other formats.
Real-time sync of your vaults across server, mobile, and web; shareable with anyone; supports REST and MCP integrations to build your personal AI knowledge base.
Enhance the file explorer with custom sorting, pinning, and hiding.
Add properties to multiple notes at once. Either right-click a folder or select multiple notes and right-click the selection.
Insert and import citations, bibliographies, notes, and PDF annotations from Zotero.
Sync highlights from Readwise to your vault.
Create notes within folders that can be accessed without collapsing the folder, similar to the functionality offered in Notion.
Complete Calendar HUB experience. Work with all your calendars in one place. Analyze your time and take action!
Sync vaults securely to self-hosted servers or WEBRTC.
