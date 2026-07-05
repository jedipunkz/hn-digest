---
source: "https://github.com/adam-s/goldseam"
hn_url: "https://news.ycombinator.com/item?id=48796042"
title: "Show HN: Goldseam – heal broken Cypress selectors with a local LLM"
article_title: "GitHub - adam-s/goldseam: Bring-your-own-model healing and authoring for Cypress · GitHub"
author: "dataviz1000"
captured_at: "2026-07-05T17:57:21Z"
capture_tool: "hn-digest"
hn_id: 48796042
score: 2
comments: 0
posted_at: "2026-07-05T17:27:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Goldseam – heal broken Cypress selectors with a local LLM

- HN: [48796042](https://news.ycombinator.com/item?id=48796042)
- Source: [github.com](https://github.com/adam-s/goldseam)
- Score: 2
- Comments: 0
- Posted: 2026-07-05T17:27:15Z

## Translation

タイトル: Show HN: Goldseat – 壊れた Cypress セレクターをローカル LLM で修復する
記事のタイトル: GitHub - adam-s/goldseat: Cypress の独自モデルの修復とオーサリング · GitHub
説明: Cypress の持ち込みモデルの修復とオーサリング - adam-s/goldseat

記事本文:
GitHub - adam-s/goldseat: Cypress 用のBring-your-own-model 修復とオーサリング · GitHub
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
アダムさん
/
ゴールドシーム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブランチ Tags Go

ファイルへ コード 「その他のアクション」メニューを開く フォルダーとファイル
64 コミット 64 コミット .agents .agents .claude .claude .github/ workflows .github/ workflows .goldseat-prompts .goldseat-prompts .vscode .vscode ベンチ ベンチ サイプレス サイプレス デモ デモ docs/ メディア docs/ メディア パッケージ パッケージ 証明 証明スクリプト スクリプト selfhost selfhost .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md cypress.config.ts cypress.config.ts goldseat.config.example.mjs goldseat.config.example.mjs package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Cypress の持ち込みモデルの修復とオーサリング - すべての修復は差分をレビューされます
セレクターが壊れます。失敗は編集されたキャプチャとなり、モデルは最小限の修正を提案し、スイートがそれを検証し、レビューした git diff として修復が完了します。ベンダー クラウドでは何も実行されません。
完全なドキュメント · 試験場 · セルフホスト レシピ · 設計ノート
goldseat は、Cypress のクラウドでホストされる cy.prompt() に代わる、オープンな独自モデルの持ち込みです。モデルは、デフォルトですでに持っている Claude Code CLI ( claude -p )、ローカルの Ollama、OpenAI 互換のエンドポイント、またはコマンドライン プログラムを使用します。すべての結果はレビュー可能なファイルとして保存されます。
セレクターが壊れる → 編集されたキャプチャ (DOM + aria ツリー + エラー)
→ モデルは最小限のセレクター修正を提案します
→ 6段はしごがそれを証明する
→ 一行の差分がレビュー用に表示されます
金継ぎにちなんで名付けられました。修理は目に見え、確認され、オブジェクトのストーリーの一部です。
ゴールドシーム
cy.prompt() (サイプレス クラウド)
ランタイムヒーラー
モデルが実行される場所
✅ あなたのマシンまたはエンドポイント
❌ ベンダークラウド
⚠️ ML または BYO のバンドル
修理が完了する場所
✅再

git diff を表示しました
❌ 実行時に再解決される
❌ 実行時に置換される
それは主張を弱める可能性がありますか？
✅ なし — セレクター文字列のみ
⚠️ 歩数を回復します
⚠️ ロケーターをサイレントに交換します
失敗の正直さ
✅ 諦めは報告された第一級の結果です
⚠️ 自信を持ったフォールスグリーンが可能
❌ 癒されました ≠ 確認されました
グリーンラン
✅ 変更なし、書き込みゼロ
—
⚠️ インストルメント化
オープンソース
✅ マサチューセッツ工科大学
❌ 独自の
さまざま
「ランタイム ヒーラー」は、Healenium/CodeceptJS-heal クラスです。全景と出典: Competition.md 。
🩹 修復 — 壊れたセレクターは、レビューされた 1 行の差分になります。スイートは着陸前にそれを検証します
✍️ 著者 — 平易な英語のステップは実際の Cypress コマンドに一度変換され、コミット可能なファイルとしてキャッシュされます
🧠 独自のモデルの持ち込み — クロード コード CLI (デフォルト)、ローカル Ollama、任意の OpenAI 互換エンドポイント、任意の実行可能ファイル
🕵️ 編集されたキャプチャ — モデルには編集された DOM + アクセシビリティ ツリー + エラーが表示されますが、アプリケーション ソースは表示されません
🪜 6 段階の検証 — オフラインでトリアージ、提案、解決、オラクルジャッジを行います。 2 つの再実行ランでアプリを再実行します
🏳️ 正直な諦め — ページはロードされず、キャプチャは低下し、信頼性は低くなります: 報告され、非表示にはなりませんでした
👻 透過的プラグイン — スイートは、 goldseat がインストールされている場合とインストールされていない場合とで同じように動作します。
📼 CI でのモデル呼び出しはゼロ — 作成されたステップは .goldseat-prompts/ のコミットされたキャッシュ ファイルから再生されます。
🗂️ すべてがファイル — キャプチャ、修復、翻訳はレビュー可能なアーティファクトとして到着し、実行時の魔法ではありません
セレクターが壊れます。 1 つのコマンドで、それを 1 行のレビュー済みの差分に修復します (実際のモデル、実際の実行、68 秒)。
npm install --save-dev goldseat
npx ゴールドシーム初期化
それが全体の統合です。緑のランは手つかずです。失敗はキャプチャ アーティファクトを .goldseat/failures/ に書き込みます。
npx ゴールドシーム ヒー

l # すべてのキャプチャの修正を提案し、検証する
git diff # セレクターのみの編集を確認してからコミットします
npx goldseat レポート # キャプチャと修復のテストごとの概要
3. 著者
サイ。訪問 ( '/shop' ) ;
サイ。ゴールドシーム ( [
'エンバーマグをカートに追加' ,
'カート数は 1 と表示されるはずです' ,
]);
ステップは、モデルを介して固定コマンド語彙 (決して eval されないコード) に変換され、コミット可能な .goldseat-prompts/ ファイルにキャッシュされ、モデル呼び出しなしで CI で再生されます。
すべてのヒールは 6 段のはしごを通過します。キャプチャされた DOM に対してオフラインでトリアージ、提案、解決、オラクルジャッジを行います。 2 つの再実行ランはアプリを再実行します。諦めと失敗は第一級であり、決して隠されることはありません。
フローチャート TD
C([壊れたテストのキャプチャ]) --> T
T{"トリアージ<br/><i>セレクターはまだキャプチャされた DOM と<br/>一致していますか?</i>"}
T -->|"はい — ドリフトではなくタイミングです"| G1([諦め・報告])
T -->|いいえ| P
P["提案<br/><i>モデルはセレクター文字列のみを編集します</i>"]
P --> R
R{"解決<br/><i>修正は<br/>キャプチャされた DOM で解決されますか?</i>"}
R -->|"0 は一致 / あいまい · 再試行 (キャップ 3)"| P
R -->|"解決または延期"| ○
O{"オラクル<br/><i>緑色のときと同じアイデンティティ<br/>？</i>"}
O -->|「詐欺師 / アイデンティティが失われた」| G2([諦め・報告])
O -->|"一致するか、ファイルにオラクルがありません"| RT
RT{"再実行テスト<br/><i>修復されたテストは単独で合格しますか?</i>"}
RT -->|いいえ| F([失敗・回復を元に戻す])
RT -->|はい| RS
RS{"仕様を再実行<br/><i>仕様全体が合格しましたか?</i>"}
RS -->|いいえ| F
RS -->|はい| H([修復済み・セレクターのみレビュー済みの差分])
classDef オフライン fill:#eef2ff、ストローク:#6366f1、color:#1e1b4b;
classDef オンライン 塗りつぶし:#ecfeff、ストローク:#0891b2、色:#083344;
classDef 良好な塗りつぶし:#dcfce7、ストローク:#16a34a、色:#052e16;
classDef ストップ フィル:#fef2f2、ストローク:#dc2626、カラー:#450a0a;
クラス P、T、R、O はオフラインです。
オンラインクラスRT、RS。
クラスH良好。
クラス G1、G2、F 停止。
読み込み中
回復は選択のみを変更できます

弦に問題があり、修正が嘘になると大声であきらめます。
ランナー
旗
どこで実行されるか
クロードコード CLI
クロード (デフォルト)
あなたのマシン、すでに持っている claude -p 経由
オラマ
オラマ:<モデル>
ローカル、ゼロエグレス — 14B Qwen で実証済みのエンドツーエンド
OpenAI対応
openai:<モデル>
あらゆるエンドポイント — 実際の OpenAI に対して実証済み。モーダル レシピを使用して独自の GPU でセルフホストする
任意の実行可能ファイル
cmd:<実行可能ファイル>
標準入力でプロンプトを入力し、標準出力で応答します
共有モデルの選択とデフォルトは、オプションの goldseat.config.mjs に存在し、CLI とプラグインの両方によって読み取られます。
なぜ 1 つの自己修復プロンプトではなく 2 つのツールがあるのでしょうか?
このマージされたループは、cy.prompt() が提供するものです。平易な英語のステップが中断されると、アプリが後退したのか、文章が曖昧すぎたのか、あるいは再解決すると、もっともらしいが間違った類似したものにたどり着くのかがわかりません。それらが結合され、自信に満ちた偽の緑色が生成されます。分離しておけば、障害モードは分離可能であり、各ツールの信頼性は維持されます。完全な推論: authored-self-healing.md 。
パス
そこには何が住んでいるのか
パッケージ/ゴールドシーム/
プラグイン + CLI — フルオプション、アーティファクトスキーマ、モデルランナー、および保証
パッケージ/aria-スナップショット/
Playwright の aria スナップショット + スタンドアロン パッケージとしてのターゲット ユーティリティ ( npm )
デモ/ + サイプレス/
備品店とドッグフードスイート
証明/
実際のアプリ、誘発されたドリフト、実際の修復 (レシート)
.agents/reference/
設計の参考資料
開発する
git clone https://github.com/adam-s/goldseat && cd goldseat
npmインストール
npm run build:パッケージ
npm 実行テスト:ユニット && npm 実行テスト:システム && npm 実行テスト:強化 && npm 実行テスト:修復
ライセンス
MIT — aria-snapshot は Apache-2.0 です (注意)。
金継ぎにちなんで名付けられました。修理は目に見えて確認でき、オブジェクトのストーリーの一部です。
Cypress の持ち込みモデルの修復とオーサリング
読み込み中にエラーが発生しました

g.このページをリロードしてください。
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

Bring-your-own-model healing and authoring for Cypress - adam-s/goldseam

GitHub - adam-s/goldseam: Bring-your-own-model healing and authoring for Cypress · GitHub
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
adam-s
/
goldseam
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
64 Commits 64 Commits .agents .agents .claude .claude .github/ workflows .github/ workflows .goldseam-prompts .goldseam-prompts .vscode .vscode bench bench cypress cypress demo demo docs/ media docs/ media packages packages proving proving scripts scripts selfhost selfhost .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md cypress.config.ts cypress.config.ts goldseam.config.example.mjs goldseam.config.example.mjs package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Bring-your-own-model healing and authoring for Cypress — every repair is a reviewed diff
A selector breaks. The failure becomes a redacted capture, your model proposes a minimal fix, the suite verifies it, and the repair lands as a git diff you review. Nothing runs in a vendor cloud.
Full Docs · Proving Ground · Self-Host Recipes · Design Notes
goldseam is an open, bring-your-own-model alternative to Cypress's Cloud-hosted cy.prompt() . You bring the model: the Claude Code CLI you already have ( claude -p ) by default, a local Ollama, any OpenAI-compatible endpoint, or any command-line program. Every result lands as a reviewable file.
A selector breaks → redacted capture (DOM + aria tree + error)
→ your model proposes a minimal selector fix
→ a six-rung ladder verifies it
→ a one-line diff lands for review
Named for kintsugi : the repair is visible, reviewed, and part of the object's story.
goldseam
cy.prompt() (Cypress Cloud)
Runtime healers
Where the model runs
✅ Your machine or endpoint
❌ Vendor cloud
⚠️ Bundled ML or BYO
Where the repair lands
✅ Reviewed git diff
❌ Re-resolved at runtime
❌ Substituted at runtime
Can it weaken assertions?
✅ Never — selector strings only
⚠️ Regenerates steps
⚠️ Swaps locators silently
Failure honesty
✅ Give-up is a first-class, reported outcome
⚠️ Confident false greens possible
❌ Healed ≠ verified
Green runs
✅ Untouched, zero writes
—
⚠️ Instrumented
Open source
✅ MIT
❌ Proprietary
Varies
"Runtime healers" is the Healenium / CodeceptJS- heal class. The full landscape, with sources: competition.md .
🩹 Heal — A broken selector becomes a one-line reviewed diff; the suite verifies it before it lands
✍️ Author — Plain-English steps translate once into real Cypress commands, cached as committable files
🧠 Bring Your Own Model — Claude Code CLI (default), local Ollama, any OpenAI-compatible endpoint, any executable
🕵️ Redacted Captures — The model sees a redacted DOM + accessibility tree + error, never your application source
🪜 Six-Rung Verification — Triage, propose, resolve, and oracle judge offline; two rerun rungs re-run the app
🏳️ Honest Give-Ups — Page never loaded, degraded capture, low confidence: reported, never hidden
👻 Transparent Plugin — A suite behaves identically with and without goldseam installed
📼 Zero Model Calls in CI — Authored steps replay from committed cache files in .goldseam-prompts/
🗂️ Everything Is a File — Captures, heals, and translations arrive as reviewable artifacts, never runtime magic
A selector breaks; one command heals it into a one-line reviewed diff (real model, real run, 68 seconds):
npm install --save-dev goldseam
npx goldseam init
That's the whole integration. Green runs are untouched; failures write capture artifacts to .goldseam/failures/ .
npx goldseam heal # propose + verify a fix for every capture
git diff # review the selector-only edit, then commit
npx goldseam report # per-test summary of captures and heals
3. Author
cy . visit ( '/shop' ) ;
cy . goldseam ( [
'Add the Ember Mug to the cart' ,
'The cart count should show 1' ,
] ) ;
Steps translate once through your model into a fixed command vocabulary (never eval 'd code), cached in a committable .goldseam-prompts/ file that replays in CI with zero model calls:
Every heal passes a six-rung ladder. Triage, propose, resolve, and oracle judge offline against the captured DOM; the two rerun rungs re-run the app; give-up and fail are first-class, never hidden:
flowchart TD
C([broken-test capture]) --> T
T{"triage<br/><i>selector still matches<br/>the captured DOM?</i>"}
T -->|"yes — timing, not drift"| G1([give up · reported])
T -->|no| P
P["propose<br/><i>model edits selector strings only</i>"]
P --> R
R{"resolve<br/><i>fix resolves in<br/>the captured DOM?</i>"}
R -->|"0 matches / ambiguous · retry (cap 3)"| P
R -->|"resolves, or deferred"| O
O{"oracle<br/><i>same identity it had<br/>while green?</i>"}
O -->|"impostor / identity gone"| G2([give up · reported])
O -->|"matches, or no oracle on file"| RT
RT{"rerun-test<br/><i>healed test passes alone?</i>"}
RT -->|no| F([fail · heal reverted])
RT -->|yes| RS
RS{"rerun-spec<br/><i>whole spec passes?</i>"}
RS -->|no| F
RS -->|yes| H([healed · selector-only reviewed diff])
classDef offline fill:#eef2ff,stroke:#6366f1,color:#1e1b4b;
classDef online fill:#ecfeff,stroke:#0891b2,color:#083344;
classDef good fill:#dcfce7,stroke:#16a34a,color:#052e16;
classDef stop fill:#fef2f2,stroke:#dc2626,color:#450a0a;
class P,T,R,O offline;
class RT,RS online;
class H good;
class G1,G2,F stop;
Loading
A heal can only ever change selector strings, and gives up loudly when a fix would be a lie.
Runner
Flag
Where it runs
Claude Code CLI
claude (default)
Your machine, via the claude -p you already have
Ollama
ollama:<model>
Local, zero egress — proven end-to-end on a 14B Qwen
OpenAI-compatible
openai:<model>
Any endpoint — proven against real OpenAI; self-host on your own GPU with the Modal recipe
Any executable
cmd:<executable>
Prompt in on stdin, reply out on stdout
Shared model choice and defaults live in an optional goldseam.config.mjs , read by both the CLI and the plugin.
Why Two Tools Instead of One Self-Healing Prompt?
That merged loop is what cy.prompt() markets. When a plain-English step breaks you can't tell whether the app regressed, the sentence was too vague, or a re-resolve would land on a plausible-but-wrong look-alike — merged, those compound into a confident false green. Kept separate, the failure modes stay separable and each tool stays trustworthy. Full reasoning: authored-self-healing.md .
Path
What lives there
packages/goldseam/
The plugin + CLI — full options, artifact schema, model runners, and guarantees
packages/aria-snapshot/
Playwright's aria snapshot + targeting utilities as a standalone package ( npm )
demo/ + cypress/
The fixture shop and the dogfood suite
proving/
Real apps, induced drift, real heals ( receipts )
.agents/reference/
Design references
Develop
git clone https://github.com/adam-s/goldseam && cd goldseam
npm install
npm run build:packages
npm run test:unit && npm run test:system && npm run test:hardening && npm run test:heal
License
MIT — aria-snapshot is Apache-2.0 with NOTICE.
Named for kintsugi — the repair is visible, reviewed, and part of the object's story.
Bring-your-own-model healing and authoring for Cypress
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
