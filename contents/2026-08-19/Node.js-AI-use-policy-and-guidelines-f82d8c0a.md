---
source: "https://github.com/nodejs/node/blob/main/doc/contributing/ai-guidelines.md"
hn_url: "https://news.ycombinator.com/item?id=49367779"
title: "Node.js AI use policy and guidelines"
article_title: "node/doc/contributing/ai-guidelines.md at main · nodejs/node · GitHub"
image: "https://opengraph.githubassets.com/6fb9575e608e106f16e31c50ca6ab3d73732f72e3113f9372965300f92fa0570/nodejs/node"
author: "thesdev"
captured_at: "2026-08-19T22:14:38Z"
capture_tool: "hn-digest"
hn_id: 49367779
score: 1
comments: 0
posted_at: "2026-08-19T22:03:24Z"
tags:
  - hacker-news
  - translated
---

# Node.js AI use policy and guidelines

- HN: [49367779](https://news.ycombinator.com/item?id=49367779)
- Source: [github.com](https://github.com/nodejs/node/blob/main/doc/contributing/ai-guidelines.md)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T22:03:24Z

## Translation

タイトル: Node.js AI 使用ポリシーとガイドライン
記事のタイトル: メインのnode/doc/contributing/ai-guidelines.md · nodejs/node · GitHub
説明: Node.js JavaScript ランタイム ✨🐢🚀✨。 GitHub でアカウントを作成して、nodejs/node の開発に貢献します。

記事本文:
メインのnode/doc/contributing/ai-guidelines.md · nodejs/node · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
パスをコピーする もっとファイル アクションを責める もっとファイル アクションを責める 最新のコミット
歴史 歴史 98 行 (75 loc) · 4.92 KB メイン ブレッドクラム
コピー パス トップ ファイルのメタデータとコントロール
RAWファイルをコピー RAWファイルをダウンロード アウトラインエディ

AI の使用に関するポリシーとガイドライン
AIを貢献に活用する場合
AIをコミュニケーションに活用すると
このドキュメントは、OpenJS Foundation AI コーディング アシスタント ポリシーに準拠しています。
ツールが人間の判断に取って代わるべきではありません。
AI を活用。
Node.js では、コントリビューターが次のことを理解し、全責任を負う必要があります。
彼らが提案するあらゆる変化。 「なぜ X が改善されるのか?」に対する答えできる
「よくわかりません。AI がやったのです」ということは決してありません。
AI ツールが貢献の生成に貢献した場合は、それを正直に認めてください。
AIによってどれだけのコードが生成されたとしても、開示は役に立たない
責任の免責として。
営利目的の商標または商用ブランドについての言及があることに注意してください。
コードベースの一部であるコミットメッセージは、次の目的で悪用される可能性があります。
利益主導のマーケティング。開示に営利目的の商標が含まれる場合、または
商業ブランドの場合は、ブランドを匿名化することをお勧めします (例:
フロンティア推論モデル、代わりにクローズドソースのコーディングエージェント
<ブランド> )、または PR で営利目的のブランド/商標のみを言及する
ただし、メッセージに説明が含まれていない場合を除き、コミット メッセージには含まれません。
特定のブランド/商標に言及しなくても意味は通じます。これら
推奨事項は営利目的のツール/モデルにのみ適用され、非営利のものには適用されません。
コントリビューターが作成していない AI 生成コードを含むプル リクエスト
個人的に理解し、テストし、検証すると、協力者の時間と時間を無駄にします。
追加の審査がなければ閉鎖される可能性があります。貢献者
このような変更を繰り返し提出したり、プロジェクトに理解を示さなかったり、
そのプロセス、または自動化された支援の使用について不誠実である
今後の投稿がブロックされる場合があります。
特に指定しない限り、プル リクエストを自動ツールで開いてはなりません
で承認されました

プロジェクトによって前進します。承認をリクエストするには、次のいずれかの方法で問題を開きます
nodejs/admin 、または自動化が可能な場合
GitHub ワークフローの形式で実行するには、プル リクエストを送信して、
ワークフローを作成し、通常のプル リクエストのレビュー プロセスを使用してコンセンサスを求めます。
AIを貢献に活用する場合
寄稿者は AI ツールを使用して寄稿を支援する場合がありますが、そのようなツールは
人間の判断を決して置き換えないでください。
AI をコーディング アシスタントとして使用する場合:
まずコードベースを理解してください。必ず理解してください
関連するサブシステム。ツールによって生成された分析を常に検証します。
実際のソースコードは人間の判断によって作成されます。
あなたが送信したすべての行を所有します。あなたはあなたのコード内のすべてのコードに対して責任を負います
作成方法に関係なく、プル リクエスト。提出された変更
プロジェクトの開発者の原産地証明書とライセンスを満たしている必要があります
要件。レビュー中に変更があれば詳細に説明できるように準備してください。
コミットを論理的に保ちます。コミットメッセージのガイドライン
そしてスカッシュガイドラインをコミットする
プル リクエストでどのツールが使用されているかに関係なく、従う必要があります。
徹底的にテストしてください。既存のテストは削除または変更しないでください
人間による検証なしで。人間の判断で検証することが重要です。
機能の意図された動作に対する新しいテストの正確性
実装がどのように動作するかとは関係なく、テストされます。
消えないでください。 PR を開いた場合は、それに従ってください。に応答する
フィードバックを受け取り、作業が完了するか明示的に終了するまで繰り返します。もしあなたが
もう追求できないので、PR を閉じます。停止した PR は進行をブロックします。
AI を使用して「最初の課題が良好」なタスクを主張しないでください。これらの問題は、
新しい貢献者がコードベースとプロセスを実践的に学ぶのに役立ちます。
コメントは役に立つものにしておいてください。人間の判断で検証してください。
コメントは必要かつ正確です。 R

単にコメントを削除する
コードの動作をもう一度説明します。ロジックが明確ではない場合にのみコメントを追加してください。
AIをコミュニケーションに活用すると
Node.js は、共同作業者と協力者を尊重した簡潔で正確なコミュニケーションを重視します。
投稿者の時間。
AI によって完全に生成されたメッセージをプル リクエスト、問題、
またはプロジェクトのコミュニケーションチャネル。このようなコミュニケーションは次の段階で削除される場合があります。
Node.js モデレーション ポリシーに従っています。
コードに関する主張は、コードを使用する前に人間の判断で検証してください。
通信。 AI ツールの結果は仮説としてのみ扱う必要があります。
信頼できる情報源として、実際のコード、ドキュメント、仕様へのリンク。
文法ツールやスペルチェック ツールは、明瞭さを向上させ、
簡潔さ。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Node.js JavaScript runtime ✨🐢🚀✨. Contribute to nodejs/node development by creating an account on GitHub.

node/doc/contributing/ai-guidelines.md at main · nodejs/node · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 98 lines (75 loc) · 4.92 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions AI use policy and guidelines
When AI is used in contributions
When AI is used in communications
This document aligns with the OpenJS Foundation AI Coding Assistants Policy .
Tools should never replace human judgment, regardless of whether they are
powered by AI.
Node.js requires contributors to understand and take full responsibility for
every change they propose. The answer to "Why is X an improvement?" can
never be "I'm not sure. The AI did it."
If AI tools assisted in generating a contribution, acknowledge that honestly.
Regardless of how much code is generated by AI, disclosure does not serve
as a disclaimer of responsibility.
Be aware that the mention of for-profit trademarks or commercial brands in
commit messages, which are part of the code base, can be abused for
profit-driven marketing. If the disclosure involves for-profit trademarks or
commercial brands, it's recommended to either anonymize the branding (e.g. say
a frontier reasoning model , a closed-source coding agent instead of
<brand> ), or only mention the for-profit brand/trademark in the PR
description, but not in the commit message, unless the message would not have
made sense without mentioning the specific brand/trademark. These
recommendations only apply to for-profit tools/models, not any non-profit ones.
Pull requests that contain AI-generated code the contributor has not
personally understood, tested, and verified waste collaborator time and
will be subject to closure without additional review. Contributors who
repeatedly submit such changes, show no understanding of the project or
its processes, or are dishonest about the use of automated assistance
may be blocked from further contributions.
Pull requests must not be opened by automated tooling, unless specifically
approved in advance by the project. To request approval, either open an issue in
nodejs/admin , or if the automation can
be done in the form of a GitHub workflow, submit a pull request to add the
workflow and use the usual pull request review process to seek consensus.
When AI is used in contributions
Contributors may use AI tools to assist with contributions, but such tools
never replace human judgment.
When using AI as a coding assistant:
Understand the codebase first. Do not skip familiarizing yourself with
the relevant subsystem. Always verify analysis generated by tools against
the actual source code with human judgement.
Own every line you submit. You are responsible for all code in your
pull request, regardless of how it was created. The submitted changes
must satisfy the project's Developer's Certificate of Origin and licensing
requirements. Be prepared to explain any change in detail during review.
Keep the commits logical. The commit message guidelines
and commit squashing guidelines
must be followed regardless of what tool is used in the pull request.
Test thoroughly. Existing tests should not be removed or modified
without human verification. It is crucial to verify, with human judgement,
the correctness of new tests against the intended behavior of the feature
being tested, independently of how the implementation happens to behave.
Do not disappear. If you open a PR, follow it through. Respond to
feedback and iterate until the work lands or is explicitly closed. If you
can no longer pursue it, close the PR. Stalled PRs block progress.
Do not use AI to claim "good first issue" tasks. These issues exist to
help new contributors learn the codebase and processes hands-on.
Keep the comments useful. Verify with human judgement that the
comments are necessary and accurate. Remove comments that simply
restate what the code does. Add comments only where the logic is non-obvious.
When AI is used in communications
Node.js values concise, precise communication that respects collaborator and
contributor time.
Do not paste messages generated entirely by AI in pull requests, issues,
or the project's communication channels. Such communication may be removed in
accordance to the Node.js moderation policy .
Verify claims about the code with human judgement before using them in
communications . Results from AI tools should only be treated as hypothesis.
Link to actual code, documentation and specifications as source of truth.
Grammar and spell-check tools are acceptable when they improve clarity and
conciseness.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
