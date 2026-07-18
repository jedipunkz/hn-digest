---
source: "https://marketplace.visualstudio.com/items?itemName=sync-md.sync-md"
hn_url: "https://news.ycombinator.com/item?id=48961628"
title: "Sync.md – keep AGENTS.md, Claude.md, .cursorrules in sync by meaning, not diff"
article_title: "sync.md - Visual Studio Marketplace"
author: "anzilzedex"
captured_at: "2026-07-18T19:58:56Z"
capture_tool: "hn-digest"
hn_id: 48961628
score: 1
comments: 0
posted_at: "2026-07-18T19:49:49Z"
tags:
  - hacker-news
  - translated
---

# Sync.md – keep AGENTS.md, Claude.md, .cursorrules in sync by meaning, not diff

- HN: [48961628](https://news.ycombinator.com/item?id=48961628)
- Source: [marketplace.visualstudio.com](https://marketplace.visualstudio.com/items?itemName=sync-md.sync-md)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T19:49:49Z

## Translation

タイトル: Sync.md – AGENTS.md、Claude.md、.cursorrules を diff ではなく意味によって同期させます。
記事のタイトル: sync.md - Visual Studio マーケットプレイス
説明: Visual Studio Code の拡張機能 - AI モデルを使用してテキストだけでなく意味を比較することにより、AI コーディング エージェント命令ファイル (AGENTS.md、CLAUDE.md、.cursorrules、.github/copilot-instructions.md) の同期を維持します。ローカルのみの長さとアクティブ ファイルのチェックも含まれます。テレメトリもバンドルもありません
[切り捨てられた]

記事本文:
sync.md - Visual Studio マーケットプレイス
コンテンツにスキップ
|マーケットプレイス
サインイン
Visual Studio Code > その他 > sync.md Visual Studio Code は初めてですか?今すぐ入手してください。同期.md
複数の AI コーディング ツールを使用していますか? AGENTS.md 、 CLAUDE.md 、および
.cursorrules は静かにバラバラになり、互いに矛盾し始めます。
sync.md は AI 命令ファイルをテキストではなく意味で比較し、
それらを整列させ、きれいに保ち、ワークスペース全体で見えるようにします。
ネストされたサブディレクトリも含みます。
AGENTS.md 、 CLAUDE.md 、 .cursorrules 、および
.github/copilot-instructions.md 。ソースを変更することはありません
コード — ルール ファイル自体のみ。テレメトリーはありません。バンドルされたクラウドなし
API キー。
「問題」パネル - 競合、ルールの欠落、および長さの警告が表示されます
通常の VS Code 診断として、sync.md ソースを使用してファイルごとにグループ化
ラベルを付けると、いつでもどこから来たのかがわかります。
テキストの差分ではなく、AI を活用した意味の比較 — これら 2 つのファイルは次のように述べています
同じことをまったく異なる言語/表現 (ドイツ語、絵文字、
日本語、中国語と普通の英語)、sync.md はそれらを次のように正しく扱います。
同期中、本当に欠落しているセキュリティ セクションにのみフラグを立てます。
コマンド パレットから「sync.md: Reconcile」を実行します。
「Reconcile」Web ビュー — 紛争や紛争についての穏やかでネイティブな感覚のビュー
2 つのファイル間に欠落しているルール。ワンクリックで「...にコピー」アクションを実行し、
提案されたすべての修正を一度に適用する「整合性を保つ」ボタン:
調整後 - 欠落していたセキュリティセクションが書き込まれるようになりました。
CLAUDE.md は自動的に明確にマークされるので、内容を正確に確認できます。
コミットする前に変更されました:
2 つ以上のルール ファイルが同じディレクトリ スコープを共有する場合、sync.md は次のように尋ねます。
テキストではなく意味を比較するための言語モデル。ファイル
同じことを別の言葉で言うものは、次のように正しく扱われます。
シン

c — 純粋な矛盾 (衝突) または一方の中に存在する規則のみ
ファイルが存在するが、別のファイルにはまったく存在しない ( missing ) というフラグが立てられます。
診断はインライン (波線) および「問題」パネルに表示され、次のラベルが付けられます。
したがって、それらが sync.md からのものであることが常にわかります。
<他のファイル> と競合します: <トピック>。 (警告)
<他のファイル> が見つかりません: <トピック>。 (警告)
意味を比較するので、同じルールをドイツ語、日本語、
絵文字や平易な英語はすべて同等のものとして認識されます。フラグを立てるだけです。
本当に何が違うのか。
モデルは競合ごとに、どのバージョンを保持するかを推奨します。
( "推奨: <file> — <reason> を維持" )、実際のプロジェクトに基づいた
利用可能な場合は証拠 — たとえば、実際の package.json テスト
スクリプトを使用すると、どのテスト コマンドを実行するかについての意見の相違を解決できます。もし
モデルは自信を持って判断できず、推奨事項は示されておらず、最初の
ファイルの文言は決定的なフォールバックとして使用されます。
sync.md を実行: Reconcile して、落ち着いたネイティブ感覚の Web ビューを開きます
競合と欠落しているルールをアイテムごとに明確にグループ化して表示します。
「コピー先…」アクションと、すべての項目に適用される「整合性を保つ」ボタン
すぐに修正することを提案しました。個々の項目を独自の項目でオーバーライドできます。
このボタンを使用します。
チェックは保存時とファイルを開くときに再実行され (デバウンス ~800ms)、キャッシュされます。
ファイル ペアごとの最後の結果 — ファイル ペアごとにモデルが再度呼び出されるのは、
監視ファイルの内容が実際に変更されました。
注意: 同期チェックは AI を活用しており、言語モデルが必要です —
GitHub Copilot (デフォルト) または LM Studio 経由のローカル モデルのいずれか
/ オラマ。これがなくても、長さとアクティブ ファイル機能は引き続き機能します。
同期チェックは、落ちるのではなくモデルが必要であることを丁寧に示します
信頼性の低いテキストの差分分析に戻ります。
auto (デフォルト) — VS Code の言語モデル API、通常はサポートされています
GitHub コパイロットによる。
ローカル — O

LM Studio や
Ollama、クラウドに依存しない完全にオフラインのセットアップ用。
ルール ファイルが rulesSync.maxLines を超える場合、空でない行 (デフォルト)
200)、1 行目の情報診断では、次のように分割することが提案されています。
ネストされたルール ファイル。 AIは必要ありません。
ステータス バーの項目には、現在のファイルを管理するルール ファイルが表示されます。
現在編集中、「最も近いファイルが優先」(最も近いルール ファイルが優先) を使用
ディレクトリ ツリーを上に歩いていきます): ルール: <相対パス> 、または
ルール: なし。それをクリックしてそのファイルを開きます。切り替えると更新されます
編集者たち。 AIは必要ありません。
ローカル モデルの有効化 (完全なオフライン同期チェック)
LM Studio または Ollama をインストールし、ローカル サーバーを起動します。
(OpenAI 互換の /v1/chat/completions エンドポイント)。
rulesSync.modelProvider を "local" に設定します。
サーバーが http://localhost:1234/v1 にない場合は、次のように設定します。
それに応じて、 rulesSync.localEndpoint 。
トリガーするルール ファイルを保存します (または sync.md: Re-scan workspace を実行します)。
ローカルモデルに対するチェック。
sync.md: 調整 — 競合/欠落ルールの Web ビューを開きます
現在の (または最初に利用可能な) 非同期ファイル ペアの場合。
sync.md: ワークスペースの再スキャン — 新しいスキャンと AI チェックを強制します
(モデルは、最後のチェック以降にコンテンツが変更された場合にのみ再呼び出されます)。
調整 Web ビューは、VS Code テーマ変数を排他的に読み取ります (
色をハードコードすることはありません）ので、明るい場所でも暗い場所でもネイティブに見えます。
テーマ。抑制されたスリーサイズタイプのスケール、ヘアライン 1px を使用します
ボックスやシャドウの代わりに境界線、最大 8 ピクセルの丸い角、単一の
アクセントカラーはプライマリの一貫性を保つボタン用に予約されています —
静かで広々とした、ネイティブな雰囲気。
ソースからの貢献/構築
拡張機能の実行については DEVELOPMENT.md を参照してください (F5)。
vsce によるパッケージ化、ファイル構造、およびすべてのチュートリアル
サンプル/ワークスペースのエッジケース。

## Original Extract

Extension for Visual Studio Code - Keeps AI coding-agent instruction files (AGENTS.md, CLAUDE.md, .cursorrules, .github/copilot-instructions.md) in sync by using an AI model to compare their meaning, not just their text. Also includes local-only length and active-file checks. No telemetry, no bundle
[truncated]

sync.md - Visual Studio Marketplace
Skip to content
| Marketplace
Sign in
Visual Studio Code > Other > sync.md New to Visual Studio Code? Get it now. sync.md
Using more than one AI coding tool? Your AGENTS.md , CLAUDE.md , and
.cursorrules quietly drift apart and start contradicting each other.
sync.md compares your AI instruction files by meaning — not text — and
keeps them aligned, clean, and visible across your whole workspace,
including nested subdirectories.
It watches AGENTS.md , CLAUDE.md , .cursorrules , and
.github/copilot-instructions.md . It never modifies your source
code — only the rule files themselves. No telemetry. No bundled cloud
API keys.
Problems panel — conflicts, missing rules, and length warnings appear
as normal VS Code diagnostics, grouped by file, with the sync.md source
label so you always know where they came from:
AI-powered meaning comparison, not text diffing — these two files say
the same thing in totally different languages/wording (German, emoji,
Japanese, Chinese vs. plain English) and sync.md correctly treats them as
in sync, only flagging a genuinely missing Security section:
Run "sync.md: Reconcile" from the Command Palette:
The Reconcile webview — a calm, native-feeling view of conflicts and
missing rules between two files, with one-click "Copy to..." actions and a
"Make consistent" button that applies every suggested fix at once:
After reconciling — the missing Security section is now written into
CLAUDE.md automatically, clearly marked so you can review exactly what
changed before committing:
When two or more rule files share the same directory scope, sync.md asks
a language model to compare their meaning, not their text . Files
that say the same thing in different words are correctly treated as in
sync — only genuine contradictions ( conflicts ) or rules present in one
file but entirely absent from another ( missing ) get flagged.
Diagnostics appear inline (squiggles) and in the Problems panel, labeled
so you always know they came from sync.md:
Conflicts with <other file>: <topic>. (Warning)
Missing vs <other file>: <topic>. (Warning)
Because it compares meaning, the same rules written in German, Japanese,
emoji, or plain English are all recognized as equivalent — it only flags
what's genuinely different.
For each conflict, the model also recommends which version to keep
( "Suggested: keep <file> — <reason> " ), grounded in real project
evidence where available — for example, an actual package.json test
script can settle a disagreement about which test command to run. If the
model can't judge confidently, no recommendation is shown and the first
file's wording is used as a deterministic fallback.
Run sync.md: Reconcile to open a calm, native-feeling webview
showing conflicts and missing rules grouped clearly, with per-item
Copy to… actions and a Make consistent button that applies every
suggested fix at once. You can override any individual item with its own
Use this button.
The check re-runs on save and on file open (debounced ~800ms) and caches
the last result per file pair — it only calls the model again when a
watched file's content actually changed.
Heads up: the sync check is AI-powered and needs a language model —
either GitHub Copilot (default) or a local model via LM Studio
/ Ollama. Without one, the length and active-file features still work;
the sync check politely tells you it needs a model rather than falling
back to unreliable text diffing.
auto (default) — VS Code's Language Model API, typically backed
by GitHub Copilot.
local — an OpenAI-compatible endpoint such as LM Studio or
Ollama, for a fully offline setup with no cloud dependency.
If a rule file exceeds rulesSync.maxLines non-empty lines (default
200), an Information diagnostic on line 1 suggests splitting it into
nested rule files. No AI required.
A status bar item shows which rule file governs the file you're
currently editing, using "nearest file wins" (the closest rule file
walking up the directory tree): Rules: <relative path> , or
Rules: none . Click it to open that file. Updates as you switch
editors. No AI required.
Enabling a local model (fully offline sync checks)
Install LM Studio or Ollama and start its local server
(OpenAI-compatible /v1/chat/completions endpoint).
Set rulesSync.modelProvider to "local" .
If your server isn't on http://localhost:1234/v1 , set
rulesSync.localEndpoint accordingly.
Save a rule file (or run sync.md: Re-scan workspace ) to trigger
a check against the local model.
sync.md: Reconcile — opens the conflict/missing-rules webview
for the current (or first available) out-of-sync file pair.
sync.md: Re-scan workspace — forces a fresh scan and AI check
(the model is only re-called if content changed since the last check).
The reconcile webview reads VS Code theme variables exclusively (it
never hardcodes colors), so it looks native in both light and dark
themes. It uses a restrained three-size type scale, hairline 1px
borders instead of boxes or shadows, ~8px rounded corners, and a single
accent color reserved for the primary Make consistent button — for
a calm, spacious, native feel.
Contributing / building from source
See DEVELOPMENT.md for running the extension (F5),
packaging with vsce , the file structure, and a walkthrough of every
edge case in the sample/ workspace.
