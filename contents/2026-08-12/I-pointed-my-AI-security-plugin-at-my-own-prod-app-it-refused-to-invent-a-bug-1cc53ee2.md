---
source: "https://github.com/kenjichristopherv-del/securitymaxxing"
hn_url: "https://news.ycombinator.com/item?id=49272217"
title: "I pointed my AI security plugin at my own prod app – it refused to invent a bug"
article_title: "GitHub - kenjichristopherv-del/securitymaxxing: Evidence-based application security for Claude Code — audit commands that prove their findings, plus skills that make Claude write secure code by default. · GitHub"
author: "kenjitubera"
captured_at: "2026-08-12T14:14:51Z"
capture_tool: "hn-digest"
hn_id: 49272217
score: 3
comments: 0
posted_at: "2026-08-12T13:36:46Z"
tags:
  - hacker-news
  - translated
---

# I pointed my AI security plugin at my own prod app – it refused to invent a bug

- HN: [49272217](https://news.ycombinator.com/item?id=49272217)
- Source: [github.com](https://github.com/kenjichristopherv-del/securitymaxxing)
- Score: 3
- Comments: 0
- Posted: 2026-08-12T13:36:46Z

## Translation

タイトル: AI セキュリティ プラグインを自分の製品アプリに向けました – バグを生み出すことを拒否しました
記事のタイトル: GitHub - kenjichristopherv-del/securitymaxxing: クロード コードの証拠に基づくアプリケーション セキュリティ — 結果を証明する監査コマンドと、デフォルトでクロードに安全なコードを作成させるスキル。 · GitHub
説明: クロード コードの証拠に基づくアプリケーション セキュリティ — 結果を証明する監査コマンドと、デフォルトでクロードに安全なコードを作成させるスキル。 - kenjichristopherv-del/securitymaxxing

記事本文:
GitHub - kenjichristopherv-del/securitymaxxing: クロード コードの証拠に基づくアプリケーション セキュリティ — 結果を証明する監査コマンドと、デフォルトでクロードに安全なコードを作成させるスキル。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ケンジクリストファーブデル
/
セキュリティ最大化
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .claude-plugin .claude-plugin コマンド コマンド スキー

lls スキル .gitignore .gitignore ライセンス ライセンス PROMPTS.md PROMPTS.md README.md README.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コードのアプリケーション セキュリティ。結果を証明する監査コマンドとスキル
あなたが求める前にクロードに安全なコードを書かせるのです。
2 人向けに構築: IDOR が何であるかを知っていて、厳密な 2 番目のペアを必要とするエンジニア
の目、そして AI の助けを借りて何かを出荷し、次のことに対する率直な答えを必要としているビルダー
「これをインターネットに公開しても安全ですか?」
AI に「コードにセキュリティ上の問題がないかチェックしてください」と依頼すると、自信に満ちたもっともらしいページが得られます。
ほとんどが間違った発見。 JSON API にヘッダーがありません。パラメータ化されたクエリでの SQL インジェクション。
存在しない CVE 番号。 2 回読んでも真実は何も見つからず、二度と実行することはありません —
したがって、顧客のアドレスを漏洩する 1 つのエンドポイントはとにかく出荷されます。
問題はモデルではありません。それはプロンプトです。
securitymaxxing は、本物のセキュリティ エンジニアが適用する規律をエンコードしています。
攻撃者が制御するソースから危険なシンクまでの経路を追跡しなければ、発見はありません。
両方とも名前が付けられているか、事実として報告されるのではなく「検証が必要」というフラグが付けられています。
まずは緩和策を確認してください。 「ORM での SQL インジェクション」に関する調査結果のほとんどは、一度読むと消えてしまいます。
そのバージョンでフレームワークが実際に行うこと。
重大度は、バグの種類がどれほど恐ろしいかではなく、現実的な悪用可能性 × 影響によって決まります。
検出結果がゼロであれば、有効な結果となります。詳細を確認するためにレポートを水増しする必要はありません。
すべての発見にはエクスプロイト パスとパッチ、つまり攻撃者が送信するリクエストが含まれています。
そして修正したコード。
/プラグイン マーケットプレイス追加 kenjichristopherv-del/securitymaxxing
/plugin install securitymaxxing
または、ローカル開発の場合:
/プラグイン マーケットプレイス追加 ~/Developer/securitymaxxing
/plugin install securitymaxxing
の

コマンド
クロード コードに「/sec」と入力すると、すべてのフィルターが表示されます。
fix (パッチを書き込む) と redteam (パッチを送信する) を除くすべてのコマンドは読み取り専用です。
あなたが所有するアプリに対する積極的な攻撃）。
クロードが関連する作業を検出すると、スキルが自動的にロードされます。コマンドは必要ありません。
セキュアコーディングは最も重要な静かなものです。監査を待つことはありません。それは
クエリが初めて作成されるときに、クエリの所有権チェック部分。
/securitymaxxing:vibe-check # 何を持っているかわからない場合は、ここから始めてください
/securitymaxxing:authz src/api # であればここから始めてください
/securitymaxxing:diff # ルートに関わるすべてのコミットの前
優れたワークフロー: 実際に出荷したコードの認証とインジェクション、シークレットは 1 回だけ
全履歴を照合し、調査結果が本物であることを証明するためにレッドチームを結成し、すべての PR を比較し、
進水前の出荷検査。
発見→証明→修正。コマンドはループとして提供されます。レビュー コマンドにより仮説が浮上し、
redteam は独自のアプリに対して動作する概念実証を構築します (そして既存のアプリを攻撃します)
防御機能が保持されていることを確認し、PoC を維持する回帰テストで問題を解決します。
二度と戻ってこないから。あなたが悪用され、その後ブロックされるのを目撃した発見は、あなた自身の発見です。
信頼できる。
PROMPTS.md — 75 を超えるセキュリティ プロンプトを、コピー＆ペースト可能な生のテキストとして表示します。それらを使用してください
カーソル、Copilot、ChatGPT、またはその他の場所。スタック用にそれらを編集します。または読んで確認してください
まさにコマンドが要求しているものです。
セクション 0 の前文から始めます。最高のレバレッジです
リポジトリ内の段落 — ツールのセキュリティ プロンプトと出力品質の前に貼り付けます。
すぐに変わります。
OWASP トップ 10 (2021)、OWASP API セキュリティ トップ 10 (2023)、LLM アプリケーションの OWASP トップ 10、
さらに、これらのリストでは十分にカバーされていないもの: ビジネス ロジックと競合状態、CI/CD、および
サプライチャイ

n、クラウド IAM、コストベースのサービス拒否、およびインシデントへの対応。
これは非常に良い一次合格であり、本当に良い先生です。侵入テストではありません。
コードを読み取ります。アプリケーションを実行したり、ファジングしたり、ライブでテストしたりしないため、コードを見つけることはできません。
実行時、負荷時、またはクラウド コンソールにのみ表示されるもの。調査結果は仮説です
確認するまでは。本物のお金や実際の個人データを保持するものはこれに属します
人間のレビューの代わりではなく、人間のレビューと並行して。
新しいプロンプト、より良いプロンプト、および修正はすべて歓迎されます。特に実際の誤検知は大歓迎です。
これらは security-review-method の誤検知防止ルールになるため、ヒットします。開く
誤って報告されたコードの問題とレビューで何を言うべきだったか。
クロード コードのための証拠に基づくアプリケーション セキュリティ — 結果を証明する監査コマンドと、デフォルトでクロードに安全なコードを作成させるスキル。
Readme Apache-2.0 ライセンス セキュリティ ポリシー
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Evidence-based application security for Claude Code — audit commands that prove their findings, plus skills that make Claude write secure code by default. - kenjichristopherv-del/securitymaxxing

GitHub - kenjichristopherv-del/securitymaxxing: Evidence-based application security for Claude Code — audit commands that prove their findings, plus skills that make Claude write secure code by default. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
kenjichristopherv-del
/
securitymaxxing
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .claude-plugin .claude-plugin commands commands skills skills .gitignore .gitignore LICENSE LICENSE PROMPTS.md PROMPTS.md README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Application security for Claude Code. Audit commands that prove their findings, plus skills
that make Claude write secure code before you ask.
Built for two people: the engineer who knows what an IDOR is and wants a rigorous second pair
of eyes, and the builder who shipped something with AI help and needs a straight answer to
"is this safe to put on the internet?"
Ask any AI to "check my code for security issues" and you get a page of confident, plausible,
mostly-wrong findings. Missing headers on a JSON API. SQL injection in a parameterized query.
A CVE number that doesn't exist. You read it twice, find nothing real, and never run it again —
so the one endpoint that does leak your customers' addresses ships anyway.
The problem isn't the model. It's the prompt.
securitymaxxing encodes the discipline a real security engineer applies:
No finding without a traced path from an attacker-controlled source to a dangerous sink.
Both ends named, or it's flagged "needs verification" instead of reported as fact.
Check the mitigation first. Most "SQL injection in an ORM" findings die once you read
what the framework actually does in that version.
Severity by realistic exploitability × impact , not by how frightening the bug class sounds.
Zero findings is a valid result. No padding a report to look thorough.
Every finding carries the exploit path and the patch — the requests an attacker sends,
and the corrected code.
/plugin marketplace add kenjichristopherv-del/securitymaxxing
/plugin install securitymaxxing
Or, for local development:
/plugin marketplace add ~/Developer/securitymaxxing
/plugin install securitymaxxing
The commands
Type /sec in Claude Code and they'll all filter into view.
Every command is read-only except fix (which writes patches) and redteam (which sends
active attacks against an app you own).
Skills load automatically when Claude detects the relevant work — no command needed.
secure-coding is the quiet one that matters most. It doesn't wait for an audit; it makes the
ownership check part of the query the first time the query gets written.
/securitymaxxing:vibe-check # start here if you're not sure what you have
/securitymaxxing:authz src/api # start here if you are
/securitymaxxing:diff # before every commit that touches a route
A good workflow: authz and injection on the code you actually shipped, secrets once
against full history, redteam to prove the findings are real, diff on every PR,
ship-check before launch.
Find → prove → fix. The commands come as a loop: a review command surfaces a hypothesis,
redteam builds a working proof-of-concept against your own app (and attacks your existing
defenses to confirm they hold), and fix closes it with a regression test that keeps the PoC
from ever coming back. A finding you've watched get exploited and then blocked is a finding you
can trust.
PROMPTS.md — 75+ security prompts as raw, copy-pasteable text. Use them in
Cursor, Copilot, ChatGPT, or anywhere else; edit them for your stack; or read them to see
exactly what the commands are asking for.
Start with Section 0, the preamble . It's the highest-leverage
paragraph in the repo — paste it before any security prompt in any tool and the output quality
changes immediately.
OWASP Top 10 (2021), OWASP API Security Top 10 (2023), and OWASP Top 10 for LLM Applications,
plus the things those lists don't cover well: business logic and race conditions, CI/CD and
supply chain, cloud IAM, cost-based denial of service, and incident readiness.
This is a very good first pass and a genuinely good teacher. It is not a penetration test.
It reads code — it doesn't run your application, fuzz it, or test it live, so it can't find
what only appears at runtime, under load, or in your cloud console. Findings are hypotheses
until you verify them. For anything holding real money or real personal data, this belongs
alongside a human review, not instead of one.
New prompts, better prompts, and corrections are all welcome — especially real false positives
you hit, since those become anti-false-positive rules in security-review-method . Open an
issue with the code that got misreported and what the review should have said.
Evidence-based application security for Claude Code — audit commands that prove their findings, plus skills that make Claude write secure code by default.
Readme Apache-2.0 license Security policy
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
