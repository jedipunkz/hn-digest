---
source: "https://github.com/albertobarnabo/lean"
hn_url: "https://news.ycombinator.com/item?id=48450852"
title: "Lean – Two Claude Code skills that stop the AI from over-engineering"
article_title: "GitHub - albertobarnabo/lean: Teaches Claude to find the clever path before taking the obvious one. 8× fewer tokens on the median real-world task — measured across 17 benchmarks. · GitHub"
author: "mamba99"
captured_at: "2026-06-08T21:38:37Z"
capture_tool: "hn-digest"
hn_id: 48450852
score: 1
comments: 1
posted_at: "2026-06-08T19:53:11Z"
tags:
  - hacker-news
  - translated
---

# Lean – Two Claude Code skills that stop the AI from over-engineering

- HN: [48450852](https://news.ycombinator.com/item?id=48450852)
- Source: [github.com](https://github.com/albertobarnabo/lean)
- Score: 1
- Comments: 1
- Posted: 2026-06-08T19:53:11Z

## Translation

タイトル: Lean – AI のオーバーエンジニアリングを阻止する 2 つのクロード コード スキル
記事のタイトル: GitHub - albertobarnabo/lean: クロードに、明らかな道を選択する前に賢い道を見つけるように教えます。現実世界のタスクの中央値でトークンが 8 分の 1 に減少 — 17 のベンチマークにわたって測定。 · GitHub
説明: 明白な道を選ぶ前に、賢い道を見つけるようにクロードに教えます。現実世界のタスクの中央値でトークンが 8 分の 1 に減少 — 17 のベンチマークにわたって測定。 - アルベルトバーナボ/リーン

記事本文:
GitHub - albertobarnabo/lean: 明らかな道を選択する前に、賢い道を見つけるようにクロードに教えます。現実世界のタスクの中央値でトークンが 8 分の 1 に減少 — 17 のベンチマークにわたって測定。 · GitHub
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
アルベルトバーナボ
/
痩せた
公共
通知
あなたはむ

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main Branches Tags Go to file Code Open more actions menu Folders and files
34 コミット 34 コミット .claude-plugin .claude-plugin アセット アセット コマンド コマンド スキル スキル テスト テスト .gitignore .gitignore CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
最高のトークンは、一度も使ったことがないものです。
「優れたエンジニアは怠惰なエンジニアです。彼らは賢い近道を見つけます。」 — スティーブ・ジョブズ
lean is a Claude Code plugin that gives your AI the instinct great engineers are known for:
pause before working hard, and make sure you can't work smart instead.
問題: AI エージェントは無駄である
無駄のない製造には、不必要な作業を表す言葉「ムダ」があります。無駄。 Toyota built the world's most efficient production system by obsessing over eliminating it.
AI エージェントにはムダの問題があります。クロードは、どんなタスクを与えられても、「もっと賢い方法はないのか？」と立ち止まらずに、徹底的に、ゼロから、全額コストをかけて、最も明白な実装に突き進みます。 And once it's writing, it adds everything it can think of: error handling, tests, abstractions, refactors — none of which was asked for.
その結果、何千もの不要なトークンが生成されます。起こる必要のなかった仕事。無駄。
lean は、重要な 2 つの瞬間だけこれを修正します。
スキル
発火したら
何が妨げられるのか
よく考えてください
アプローチを選択する前に
API、パッケージ、またはワンライナーがすでに存在する場合に最初から実装する
外科用
各ブロックを書き込む前に
Adding error handling, tests, and abstractions nobody asked for
よく考えて、もっと賢い道はないのかと問いかけます。
surgical asks: did the user actually ask for this?
これらが連携して、戦略と実行のあらゆるレベルでリーンを実現します。
クロードに500スタ生成してもらう

ユーザープロファイルを調べます。リーンを使用しない場合、すべてのプロファイルがインラインで書き込まれます。すべて 500、フィールドごと、出力の 66,320 トークンです。リーンでは、代わりに 54 行の Faker スクリプトを作成します。トークン372個。
リーンなし: ~66,320 トークン — Claude Sonnet API の価格で約 1.00 ドル。
リーンの場合: ~372 トークン — 約 0.5 セント。
同じ結果です。 178×コスト。
それは特殊なケースではありません。これは、最初に考えることを教えられていないすべての AI のデフォルトの動作です。
これら 17 のタスク (通常のバイブコーディング午後) のコストは、グリーディ トークン 88,655 トークンに対し、リーン トークン 4,762 です。プロンプトを 1 つも変更せずに、毎回 1.10 ドルの差が生じます。
17 のベンチマーク シナリオからの実際の出力。2 回考えた場合のみ、外科手術のみ、および両方の組み合わせの 3 つの条件下でそれぞれ個別にテストされました。三方の内訳 →
ギャップは狭くない。 17 の実際のタスク (バグ修正、スクリプト、API 統合、データ生成) にわたって、節約範囲は 2 倍から 178 倍、中央値は 8 倍です。
このような広がりが存在するのは、廃棄物が 1 か所から出ていないためです。 2 つの独立した故障モードがあります。
スコープクリープとは、クロードが、固定された限定された答えを持つタスクの上に、あなたが求めていないもの (ドライラン フラグ、ドキュメント文字列、エラー処理、テスト スイート) を追加することです。タスクは小さいです。クリープはそうではありません。外科手術はこれをキャッチします。
間違った戦略は、ライブラリ、API、または組み込みが既に正しく完全に解決しているにもかかわらず、クロードが高価な道を選択することです。 10,000 の空港がある場合、124 の空港がハードコーディングされます。 1 月 1 日に期限が切れるホリデー セット。structuraldClone() が組み込みの場合、手動で実行される deepClone。 think-tww はこれをキャッチします。
これらは同じ問題のバリエーションではなく、タスクによって一方または両方が引き起こされることも、どちらも引き起こされないこともあります。そのため、スキルは別になります。
外科手術では、カウントによってより多くのシナリオが検出されます。よく考えて高価なものを見つけてください。

178× 外れ値がそのスライスに存在します。両方の障害モードが存在する場合、乗数はスタックします。
正直な注意点が 1 つあります。17 シナリオ中 3 つ (ダーク モードの切り替え、ページネーション、ユーザー認証設定) では、外科手術のみが両方のスキルを組み合わせたパフォーマンスを上回りました。 think-twice が手動で作成した最小限のソリューションを超えるセットアップ定型のライブラリにリダイレクトすると、それが追加されるのは痛手です。スキルは必ずしも加算されるわけではありません。だからこそ、スキルは分離されており、完全な 3 者間内訳ですべての条件が示されているのです。
「ステージング データベース用に 500 の現実的なユーザー プロファイルを生成する」
貪欲な
無駄のない
アプローチ
500 個の JSON レコードをインラインで書き込みます
54 行の @faker-js/faker スクリプト、パラメータ化
トークン
~66,320
~372 — 178 分の 1
データ品質
繰り返し（約 30 個の名前が再利用）
統計的に多様な 50 以上のロケール
Bcrypt ハッシュ
偽のハッシュ — ログインでは使用できません
実ハッシュ — ログインで使用可能
再実行可能性
ゼロ — 一時的な出力
シード済み、バージョン管理済み、--count フラグ
チェックポイント
—
よく考えます #2 (フェイカー) + #3 (500 static = 間違った形状)
「このディレクトリ内のすべての .jpeg ファイルの名前を .jpg に変更するスクリプトを作成します。」
貪欲な
無駄のない
出力
110 行の CLI — --dry-run 、 --recursive 、 --verbose 、 --directory を使用した argparse 、ロギング セットアップ、ファイルごとの try/excel 、名前変更されたファイル カウンター、タイプ ヒント、main() ガード
3行のpathlibループ
トークン
～725
~19 — 38 分の 1 に減少
フラグが追加されました
4 ( --dry-run 、 --recursive 、 --verbose 、 --directory )
0
よく考えてください
正しく起動しない - pathlib はすでに適切なツールです
—
チェックポイント
—
外科手術 - ユーザーが CLI ツールではなくスクリプトを要求した
「サインアップフォームに電子メール検証を追加」
貪欲な
無駄のない
アプローチ
RFC 5322 正規表現 + 65 エントリの使い捨てドメイン ブロックリスト + ライブ MX/SMTP プローブ + lru_cache
4 行のコンパイル済み正規表現、stdlib re のみ
トークン
~1,675
~93 — 18 分の 1
ライブネットワーク通話
すべての検証 (SMTP プローブ)
存在しません
セント

メンテナンスするリング
65 個のハードコーディングされた使い捨てドメイン
0
依存関係
smtplib 、ソケット、ロギング、lru_cache
標準ライブラリを超えるものはありません
チェックポイント
—
外科 — 「電子メールを検証する」≠「検証モジュールを構築する」
「フライト検索のために空港の IATA コードを都市名にマッピングします」
貪欲な
無駄のない
アプローチ
約 124 の空港を静的オブジェクトとしてハードコード
npm install Airport + 5 行のルックアップ
トークン
～1,710
~93 — 18 分の 1
空港のカバー範囲
約 10,000 の IATA コードのうち 124 (1.2%)
すべて ~10,000
「TXL」、「CGK」、「DOH」
見つかりません
カバーされた
正しさ
98.8% の空港では不正解
完了
チェックポイント
—
よく考えてみる #2 — 既存のパッケージ
「parse_date の off-by-one エラーを修正」
貪欲な
無駄のない
出力
バグ修正 + 型アノテーション + 入力検証 + docstring + 13 ユニット
[切り捨てられた]
これをプロジェクトの CLAUDE.md に追加します。スキルとは異なり、CLAUDE.md は常にコンテキスト内にあり、いつ適用するかについてクロードの判断に依存しません。
** 重要なコーディング タスクの前 ** (新機能、データ生成、約 20 行にわたる実装):
一時停止して確認してください — パブリック API、パッケージ、またはワンライナーですでにこの問題は解決されていますか? 「はい」の場合は、それを使用してください。
その場合のみ、今日の問題を解決する最小限の作業に進みます。
** 各コード ブロックを記述する前に: **
明示的に要求されたもののみをビルドします。エラー処理、テスト、型注釈などを追加しないでください。
要求されない限り、docstring または抽象化。何か追加する価値があると思われる場合は、その後にそう言ってください
出力を提供します。一方的に追加しないでください。
** 次の場合は両方のルールをスキップします。 ** 〜 10 行以下のバグ修正、infra/terraform/k8s、DB クエリ、または次の場合
ユーザーは、完全な実装または運用準備が整った実装を明示的に要求しました。
オプション 2 — クロード コードのスキル
スキルは、手動で呼び出されたとき、またはクロードがコンテキストが一致すると判断したときに、完全なルールブックをロードします。オンデマンドでの使用や、これらのルールを適用したくないプロジェクトに適しています。

いつでもいます。
/プラグイン マーケットプレイス追加 albertobarnabo/lean
/プラグインのインストール lean@lean
まずリポジトリをマーケットプレイスとして登録し、そこからリーン プラグインをインストールします。 @lean サフィックスはマーケットプレイス名です。その後セッションを再起動して、スキルをロードします。
カール経由 (両方のスキルをインストール):
BASE= " https://raw.githubusercontent.com/albertobarnabo/lean/main/skills "
よく考えて手術をするスキルを身につけるため。する
カール -sL " $BASE / $skill /SKILL.md " -o ~ /.claude/skills/ $skill /SKILL.md --create-dirs
完了しました
単一スキルのみ:
# よく考えてください
カール -sL https://raw.githubusercontent.com/albertobarnabo/lean/main/skills/think-twice/SKILL.md \
-o ~ /.claude/skills/think-twice/SKILL.md --create-dirs
# 外科手術のみ
カール -sL https://raw.githubusercontent.com/albertobarnabo/lean/main/skills/surgical/SKILL.md \
-o ~ /.claude/skills/surgical/SKILL.md --create-dirs
手動呼び出し (特定のタスクにスキルを強制):
これらのスキルは定説ではありません。次の場合にそれらをオーバーライドします。
いずれの場合も、クロードは続行し、それがオーバーライドされる理由を述べます。
無駄のない考え方とは、不注意な行動を減らすことではありません。それは、価値を生み出すことを正確に実行し、コストがかかる前に他のすべてを削減することです。
スティーブ・ジョブズは怠惰を美化していませんでした。彼は工学的判断の最高の形態、つまり明白な道の手前で立ち止まり、賢い道を見つけ、それだけを採用する規律について説明していました。
ほとんどの AI コーディング ツールは、より多くのことを実行できるように最適化されています。彼らは徹底的に、完全に、防御的に生成します。生成することが彼らの得意なことだからです。
リーンは正しいことを行うために最適化します。トークンが流れるまでの 2 つの質問、2 つの瞬間:
もっと賢い方法はあるでしょうか？
これはまさに質問されたことですか？
新しい無駄パターンを発見しました。より良いパスが存在する場合に、クロードがデフォルトで高価なパスを選択するというタスクですか? PR を開きます。
exi 内の新しいショートカット行

スティングスキルの表
まだカバーされていないパターンの新しいスキル
自分の使用量からの実際のトークンコストの比較
最高の貢献とは、最高のコードと同様、必要なことを正確に実行するものであり、それ以上のことは何も行いません。
クロードに、明白な道を選ぶ前に賢い道を見つけるように教えます。現実世界のタスクの中央値でトークンが 8 分の 1 に減少 — 17 のベンチマークにわたって測定。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Teaches Claude to find the clever path before taking the obvious one. 8× fewer tokens on the median real-world task — measured across 17 benchmarks. - albertobarnabo/lean

GitHub - albertobarnabo/lean: Teaches Claude to find the clever path before taking the obvious one. 8× fewer tokens on the median real-world task — measured across 17 benchmarks. · GitHub
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
albertobarnabo
/
lean
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
34 Commits 34 Commits .claude-plugin .claude-plugin assets assets commands commands skills skills tests tests .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md View all files Repository files navigation
The best tokens are the ones you never spent.
"A great engineer is a lazy engineer. They find the clever shortcut." — Steve Jobs
lean is a Claude Code plugin that gives your AI the instinct great engineers are known for:
pause before working hard, and make sure you can't work smart instead.
The Problem: AI Agents Are Wasteful
Lean manufacturing has a word for unnecessary work: muda . Waste. Toyota built the world's most efficient production system by obsessing over eliminating it.
AI agents have a muda problem. Given any task, Claude charges ahead with the most obvious implementation — thorough, from scratch, at full cost — without stopping to ask: is there a smarter path? And once it's writing, it adds everything it can think of: error handling, tests, abstractions, refactors — none of which was asked for.
The result: thousands of unnecessary tokens. Work that didn't need to happen. Waste.
lean fixes this at the only two moments that matter.
Skill
When it fires
What it prevents
think-twice
Before picking an approach
Implementing from scratch when an API, package, or one-liner already exists
surgical
Before writing each block
Adding error handling, tests, and abstractions nobody asked for
think-twice asks: is there a smarter path?
surgical asks: did the user actually ask for this?
Together they enforce lean at every level — strategy and execution.
Ask Claude to generate 500 staging user profiles. Without lean, it writes every profile inline — all 500, field by field, 66,320 tokens of output. With lean, it writes a 54-line faker script instead. 372 tokens.
Without lean: ~66,320 tokens — about $1.00 at Claude Sonnet API pricing.
With lean: ~372 tokens — about half a cent.
Same result. 178× the cost.
That's not an edge case. That's the default behavior of every AI that hasn't been taught to think first.
These seventeen tasks — a normal vibe-coding afternoon — cost 88,655 tokens greedy vs. 4,762 tokens lean . That's a $1.10 difference, every time, without changing a single prompt.
Real outputs from 17 benchmark scenarios, tested independently under three conditions each: think-twice only, surgical only, and both combined. Three-way breakdown →
The gap isn't narrow. Across 17 real tasks — bug fixes, scripts, API integrations, data generation — savings range from 2× to 178× , with a median of 8× .
That spread exists because the waste doesn't come from one place. There are two independent failure modes.
Scope creep is Claude adding what you didn't ask for — --dry-run flags, docstrings, error handling, test suites — on top of a task with a fixed, bounded answer. The task is small; the creep is not. surgical catches this.
Wrong strategy is Claude picking the expensive path when a library, API, or built-in already solves it correctly and completely. 124 airports hardcoded when there are 10,000. A holiday set that expires January 1. Hand-rolled deepClone when structuredClone() is a built-in. think-twice catches this.
These aren't variations of the same problem — a task can trigger one, both, or neither. Which is why the skills are separate.
surgical catches more scenarios by count. think-twice catches the expensive ones — the 178× outlier lives in that slice. When both failure modes are present, the multipliers stack.
One honest caveat: in 3 of 17 scenarios (dark mode toggle, pagination, user auth setup), surgical alone outperformed both skills combined. When think-twice redirects to a library whose setup boilerplate exceeds a minimal hand-rolled solution, adding it hurts. The skills are not always additive — which is why they're separate, and why the full three-way breakdown shows every condition.
"Generate 500 realistic user profiles for our staging database"
Greedy
Lean
Approach
Writes 500 JSON records inline
54-line @faker-js/faker script, parameterized
Tokens
~66,320
~372 — 178x fewer
Data quality
Repetitive (~30 names recycled)
Statistically varied, 50+ locales
Bcrypt hashes
Fake hashes — not login-usable
Real hashes — login-usable
Re-runnability
Zero — ephemeral output
Seeded, version-controlled, --count flag
Checkpoints
—
think-twice #2 (faker) + #3 (500 static = wrong shape)
"Write a script to rename all .jpeg files to .jpg in this directory"
Greedy
Lean
Output
110-line CLI — argparse with --dry-run , --recursive , --verbose , --directory , logging setup, per-file try/except , renamed-file counter, type hints, main() guard
3-line pathlib loop
Tokens
~725
~19 — 38x fewer
Flags added
4 ( --dry-run , --recursive , --verbose , --directory )
0
think-twice
Correctly does not fire — pathlib is already the right tool
—
Checkpoint
—
surgical — user asked for a script, not a CLI tool
"Add email validation to our signup form"
Greedy
Lean
Approach
RFC 5322 regex + 65-entry disposable domain blocklist + live MX/SMTP probe + lru_cache
4-line compiled regex, stdlib re only
Tokens
~1,675
~93 — 18x fewer
Live network call
On every validation (SMTP probe)
Does not exist
Strings to maintain
65 hardcoded disposable domains
0
Dependencies
smtplib , socket , logging , lru_cache
None beyond stdlib
Checkpoint
—
surgical — "validate email" ≠ "build a validation module"
"Map airport IATA codes to city names for our flight search"
Greedy
Lean
Approach
Hardcodes ~124 airports as a static object
npm install airports + 5-line lookup
Tokens
~1,710
~93 — 18x fewer
Airport coverage
124 of ~10,000 IATA codes (1.2%)
All ~10,000
"TXL", "CGK", "DOH"
Not found
Covered
Correctness
Wrong for 98.8% of airports
Complete
Checkpoint
—
think-twice #2 — existing package
"Fix the off-by-one error in parse_date"
Greedy
Lean
Output
Bug fix + type annotations + input validation + docstring + 13 unit
[truncated]
Add this to your project's CLAUDE.md . Unlike skills, CLAUDE.md is always in context — no reliance on Claude's judgment about when to apply it:
** Before any substantial coding task ** (new feature, data generation, implementation over ~ 20 lines):
pause and check — does a public API, package, or one-liner already solve this? If yes, use it.
Only then proceed with the minimum that solves the problem today.
** Before writing each code block: **
build only what was explicitly asked for. Do not add error handling, tests, type annotations,
docstrings, or abstractions unless requested. If something seems worth adding, say so after
delivering the output — don't add it unilaterally.
** Skip both rules for: ** bug fixes under ~ 10 lines, infra/terraform/k8s, DB queries, or when
the user explicitly asked for a complete or production-ready implementation.
Option 2 — Claude Code skills
Skills load their full rulebook when invoked manually or when Claude judges the context matches. Better for on-demand use or projects where you don't want these rules active at all times.
/plugin marketplace add albertobarnabo/lean
/plugin install lean@lean
First register the repo as a marketplace, then install the lean plugin from it. The @lean suffix is the marketplace name. Restart your session afterward so the skills load.
Via curl (installs both skills):
BASE= " https://raw.githubusercontent.com/albertobarnabo/lean/main/skills "
for skill in think-twice surgical ; do
curl -sL " $BASE / $skill /SKILL.md " -o ~ /.claude/skills/ $skill /SKILL.md --create-dirs
done
Single skill only:
# think-twice only
curl -sL https://raw.githubusercontent.com/albertobarnabo/lean/main/skills/think-twice/SKILL.md \
-o ~ /.claude/skills/think-twice/SKILL.md --create-dirs
# surgical only
curl -sL https://raw.githubusercontent.com/albertobarnabo/lean/main/skills/surgical/SKILL.md \
-o ~ /.claude/skills/surgical/SKILL.md --create-dirs
Manual invocation (force a skill on a specific task):
These skills are not dogma. Override them when:
In all cases, Claude proceeds — and states why it's overriding.
Lean thinking is not about doing less carelessly. It's about doing exactly what creates value — and cutting everything else before it costs you.
Steve Jobs wasn't romanticizing laziness. He was describing the highest form of engineering judgment: the discipline to stop before the obvious path, find the clever one, and take only that.
Most AI coding tools optimize for doing more . They generate thoroughly, completely, defensively — because generating is what they're good at.
lean optimizes for doing right . Two questions, two moments, before the tokens flow:
Is there a smarter path?
Is this exactly what was asked?
Found a new waste pattern — a task where Claude defaults to the expensive path when a better one exists? Open a PR:
A new shortcut row in an existing skill's table
A new skill for a pattern not yet covered
A real token-cost comparison from your own usage
The best contributions, like the best code, are the ones that do exactly what's needed — nothing more.
Teaches Claude to find the clever path before taking the obvious one. 8× fewer tokens on the median real-world task — measured across 17 benchmarks.
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
