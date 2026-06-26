---
source: "https://github.com/omkar-foss/noslop-oss"
hn_url: "https://news.ycombinator.com/item?id=48688267"
title: "No-Slop OSS, a checklist of contribution best practices when using AI (or not)"
article_title: "GitHub - omkar-foss/noslop-oss: A checklist of no-slop contribution best practices when using AI (or not) in open source projects. · GitHub"
author: "omkar-foss"
captured_at: "2026-06-26T16:38:06Z"
capture_tool: "hn-digest"
hn_id: 48688267
score: 3
comments: 0
posted_at: "2026-06-26T16:08:10Z"
tags:
  - hacker-news
  - translated
---

# No-Slop OSS, a checklist of contribution best practices when using AI (or not)

- HN: [48688267](https://news.ycombinator.com/item?id=48688267)
- Source: [github.com](https://github.com/omkar-foss/noslop-oss)
- Score: 3
- Comments: 0
- Posted: 2026-06-26T16:08:10Z

## Translation

タイトル: No-Slop OSS、AI を使用する (または使用しない) 場合の貢献のベスト プラクティスのチェックリスト
記事のタイトル: GitHub - omkar-foss/noslop-oss: オープンソース プロジェクトで AI を使用する (または使用しない) 場合の、スロップなしのコントリビューションのベスト プラクティスのチェックリスト。 · GitHub
説明: オープンソース プロジェクトで AI を使用する (または使用しない) 場合の、無駄のない貢献のベスト プラクティスのチェックリスト。 - omkar-foss/noslop-oss

記事本文:
GitHub - omkar-foss/noslop-oss: オープンソース プロジェクトで AI を使用する (または使用しない) 場合のノースロップ貢献のベスト プラクティスのチェックリスト。 · GitHub
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
オムカルフォス
/
ノスロップオス
公共
通知
通知設定を変更するにはサインインする必要があります
あ

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md ライセンス ライセンス NOSLOP.md NOSLOP.md README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI を使用する場合（または使用しない場合）の、無駄のない貢献のベスト プラクティスのチェックリスト
オープンソースプロジェクトで。
最近では、明確な所有権のない、AI によって生成されたリポジトリや PR が非常に多くあります。
「AI スロップ」とも呼ばれます。 AI活用が停滞の一因だが、場合によっては
それだけが理由ではありません。多くの新しい、本物の (人間の) 貢献者は単に
オープンソースのワークフローやプロジェクトの期待に慣れておらず、AI ツールは
時にはそれらのギャップを増幅させます。その結果、彼らの貢献が失われることになります
ノイズの中に紛れ込んでいて、メンテナから無視されています。
このチェックリストは、貢献者が信頼と信頼を築くのに役立つベスト プラクティスに焦点を当てています。
AI を使用しているかどうかにかかわらず、質の高い投稿を投稿してください。チェックリストです
メンテナーがレビューを楽しむような貢献者になってくれることを感謝します。
チェックリスト
0. 目的を持って貢献する
1. プロジェクトをよく理解する
2. プロジェクトコミュニティとの関わり
3. プロジェクト環境を徹底的にセットアップする
4. プロジェクトの AI ポリシーを理解する
5. AI を代替品としてではなく、ツールとして使用する
6. テスト、検証、文書化
8. 問題への取り組みを開始する
9. 早い段階から頻繁にコミュニケーションをとる
10. 投稿の送信
高品質で手間のかからない OSS コントリビューターになるには、次のガイドラインに従ってください。 🌟
プロジェクトに価値を生み出すという目標を持って貢献してください。
学び、問題を解決し、コミュニティに貢献すること。
可視性を最適化するのではなく、有意義な貢献を行うことに重点を置く
またはメトリクス。
他の人とではなく、自分自身の進歩と競争してください。コラボレーションを選択する
競争。

☐ 1. プロジェクトをよく理解する
プロジェクトの Readme ファイル、寄稿ファイル、および行動規範ファイルを読みます。
プロジェクトの目標、アーキテクチャ、技術スタックを調査します。
既存の問題、ディスカッション、プルリクエストを調査してコミュニティを把握する
優先順位。
最近のプル リクエストを読んでプロジェクトのコーディング スタイルを理解し、レビューします。
期待と現在の開発の優先事項。
プロジェクト固有のガイドライン (設計ドキュメント、ロードマップ、
アーキテクチャの決定)。
☐ 2. プロジェクトコミュニティと交流する
プロジェクトのコミュニティに適している場合は、公開の場で自己紹介をしてください
コミュニケーション チャネル (Slack、Discord、メーリング リストなど)。
例: こんにちは！私は[名前]です。[プロジェクト]に貢献したいと思っています。私は
特に[特定の分野]に興味がある。
関連する質問は公開フォーラムで行い、個人的に質問しないでください。
議論の重複。
質問する前に、数分かけて既存の問題を検索してください。
ディスカッションとドキュメント。
最初に読んで、後で投稿する、つまりコミュニティのディスカッションを観察する前に
飛び込む。
プロジェクトのコミュニケーション規範に従ってください (例: 応答時間、トーン、
会議のスケジュール）。
☐ 3. プロジェクト環境を徹底的にセットアップする
リポジトリをフォークしてローカルにクローンを作成します。
プロジェクトに文書化されているとおりに開発環境をセットアップします。
依存関係をインストールし、プロジェクトがローカルで実行されることを確認します。
推奨されている場合は、コミット前のフック、リンター、またはフォーマッタをセットアップします。
☐ 4. プロジェクトの AI ポリシーを理解する
プロジェクトに AI 貢献ポリシー (AI の開示など) があるかどうかを確認する
要件）。
プロジェクトで AI の使用について言及されていない場合は、それが許容されると想定しないでください。
開示やその他の期待が明確でない場合は、メンテナに尋ねることを検討してください。
できれば、PR/問題の中で AI の使用状況を開示するようにしてください。
プロジェクトの AI ポリシーの需要

座る。
☐ 5. AI を代替品としてではなく、ツールとして使用する
AI を使用して次のことを行います。
アイデアをブレインストーミングするか、ドラフト コード スニペットを生成します。
反復的なタスク (書式設定、スキャフォールディングのテストなど) を自動化します。
複雑なプロセス (セットアップ手順、トラブルシューティングなど) を文書化します。
AI が生成したコードは、以下の場合を除いて決して送信しないでください。
一行ずつ見直していきます。
プロジェクトのスタイルに合わせて変更します。
送信するすべての変更を、それが作成されたものであるかどうかに関係なく説明できるようにしてください。
あなた自身、または AI によって支援されます。
☐ 6. テスト、検証、文書化する
プロジェクトのテスト スイートを実行し、エラーがないことを確認します。
変更に対する新しいテストを作成します。
変更が意図したとおりに機能することを手動で確認します。
AI によって生成されたテストを使用する場合は、エッジ ケースがないか慎重にレビューしてください。
変更があった場合は、ドキュメント、サンプル、または変更ログのエントリを更新します。
ユーザーに影響を与えます。
小規模な機能の実装 (「良好な最初の問題」というラベルが付いています)。
プロジェクトに完全に慣れるまでは、大規模な変更や曖昧な変更は避けてください。
問題トラッカーで「ヘルプが必要」または「レビューが必要」のラベルを確認します。
☐ 8. 問題への取り組みの開始
問題の PR を提起する前に、既存の PR が 1 つ以上あるかどうかを確認してください。
それをPRする。既存の PR の場合:
これらの PR のいずれかが積極的に更新されている場合は、その問題を次のような問題として選択しないでください。
重複した作業が発生する可能性があります。取り組む別の問題を見つけてください。
これらの PR がすべて非アクティブな場合は、次の後に問題が発生する可能性があります。
メンテナーに確認中。
作業を開始する前に問題についてコメントし、メンテナに確認してください
重複を避けるため。
例: 「これに取り組みたいのですが、これを手に取ってもいいですか?
次のように実装してください: [あなたの簡単な計画]。」
☐ 9. 早い段階で頻繁にコミュニケーションを取る
より大きな変更を行う場合は、提案するアプローチについてフィードバックを求めてください。
PRを開く。
提案した解決策が拒否されることを覚悟してください。良いアイデアばかりではありません
プロジェクトと一致する

ロードマップ。
思考プロセスを号/PR に文書化します。
フィードバックに敏感に反応し、反復してください。
☐ 10. 投稿の送信
プロジェクトにプル リクエスト テンプレートがある場合は、それに徹底的に従ってください。
可能な限り、プル リクエストは単一の論理変更に焦点を当ててください。
変更用に新しいブランチを作成します (例: fix/issue-123 )。
明確で説明的なコミット メッセージを作成します。
プロジェクトの規則に従ってください (例: 修正 #123: [説明] )。
AI ツールを使用していて、プロジェクトがコミットレベルの AI 開示を奨励している場合、
プロジェクトの優先形式 (例: Assisted-by または Co-authored-by ) を使用します。
プロジェクトが DCO サインオフを推奨する場合は、次を使用してコミットに DCO サインオフを指定します。
によるサインオフ。
コミット時に Assisted-By を使用した AI 開示を伴う DCO サインオフの例:
ユーザー認証モジュールのエラー処理を改善しました。
- 認証失敗に対するカスタム エラー クラスを追加しました
- 一時的な障害に対する再試行ロジックの実装
- 新しいエラー シナリオをカバーするためにテストを更新しました
支援: GitHub Copilot (新しいエラー シナリオのテスト用)
サインオフ者: あなたの名前 <youremail@address.here>
PR の説明に関連する問題への言及を含めます。
PR の説明では次のようになります。
AI の使用状況を開示します (例: 「この PR には、AI を利用したコード生成が含まれています)
[特定の部分]」)。
プロセスと推論を説明してください。
関連するディスカッションや問題へのリンク。
PR のフィードバックを監視し、変更を加える準備を整えてください。
査読者に応答する時間を与え、更新を繰り返し求めることは避けてください（例:
数時間ごと)。
時間を割いてフィードバックを提供してくれた査読者に感謝します。
PR が非公開の場合は、改善方法についてフィードバックを求めてください。
あなたの間違いから。
あなたの貢献を祝い、他の人から受けた助けに感謝しましょう
コミュニティメンバー。オープンソースはコミュニティの取り組みです。
AI が生成したコードをレビューせずに送信しないでください。

理解とテスト
それ。
コミュニティのガイドラインやコミュニケーション規範を無視しないでください。
事前の議論と承認なしに大規模な変更を送信しないでください。
問題トラッカーや PR キューに低品質の投稿をスパム送信しないでください。
コードベースを理解したり、AI に関与したりする代わりに AI に依存しないでください。
コミュニティ。
プロジェクトはこのチェックリストを NOSLOP.md として採用することを歓迎します。
その一部を CONTRIBUTING.md に組み込むか、単にこのリポジトリを共有する
貢献者と一緒に。
何か追加したい場合は、お気軽に PR を上げてください。を参照してください。
貢献ガイドラインについては上記のチェックリストをご覧ください。すべての貢献
この行動規範を遵守する必要があります。
AI を使用してオープンソースに責任を持って効果的に貢献する方法
生成型 AI の時代におけるオープンソースの維持: メンテナとコントリビュータへの推奨事項
オープンソースはコードだけでなく人を通じて成功します。信頼関係を築き、
明確なコミュニケーションと良好なコラボレーションは、
ライティングソフト。
このチェックリストは、トーンを改善するために ChatGPT の支援を受けて見直されました。
そして文法の一貫性。すべての提案は、修正される前に手動でレビューされました。
組み込まれています。
オープンソース プロジェクトで AI を使用する (または使用しない) 場合の、無駄のない貢献のベスト プラクティスのチェックリスト。
CC0-1.0ライセンス
行動規範
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

A checklist of no-slop contribution best practices when using AI (or not) in open source projects. - omkar-foss/noslop-oss

GitHub - omkar-foss/noslop-oss: A checklist of no-slop contribution best practices when using AI (or not) in open source projects. · GitHub
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
omkar-foss
/
noslop-oss
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md LICENSE LICENSE NOSLOP.md NOSLOP.md README.md README.md View all files Repository files navigation
A checklist of no-slop contribution best practices when using AI (or not)
in open source projects.
These days there are so many AI-generated repos and PRs without clear ownership,
often referred to as "AI slop". AI use is a reason for the slop, but some cases
it's not the only reason. Many new, genuine (human) contributors are simply
unfamiliar with open source workflows and project expectations, and AI tools can
sometimes amplify those gaps. It results in their contributions getting lost
among the noise and getting ignored by maintainers.
This checklist focuses on best practices that help contributors build trust and
submit high-quality contributions, whether they use AI or not. It's a checklist
for becoming the kind of contributor maintainers enjoy reviewing!
Checklist
0. Contribute with purpose
1. Understand the project well
2. Engage with the project community
3. Set up your project environment thoroughly
4. Understand the project's AI policies
5. Use AI as your tool, not as your replacement
6. Test, validate and document
8. Starting to work on an issue
9. Communicate early and often
10. Submitting your contribution
Follow these guidelines to become a high quality, no-slop OSS contributor! 🌟
Contribute with the goal of creating value for the project, whether you're
learning, solving a problem, or giving back to the community.
Focus on making meaningful contributions rather than optimizing for visibility
or metrics.
Compete with your own progress, not with others. Choose collaboration over
competition.
☐ 1. Understand the project well
Read the project's readme, contributing, and code of conduct files.
Study the project's goals, architecture, and tech stack.
Explore existing issues, discussions, and pull requests to grasp community
priorities.
Read recent pull requests to understand the project's coding style, review
expectations, and current development priorities.
Check for project-specific guidelines (e.g., design docs, roadmaps, or
architecture decisions).
☐ 2. Engage with the project community
If appropriate for the project's community, introduce yourself in its public
communication channels (Slack, Discord, mailing lists, etc).
Example: Hi! I'm [Name], and I'd like to contribute to [Project]. I'm
particularly interested in [specific area].
Ask relevant questions in public forums, and not privately to avoid
duplicating discussions.
Before asking a question, spend a few minutes searching existing issues,
discussions, and documentation.
Read first, contribute later i.e. observe community discussions before
jumping in.
Follow the project's communication norms (e.g., response times, tone,
meeting schedules).
☐ 3. Set up your project environment thoroughly
Fork the repository and clone it locally.
Set up the development environment exactly as documented in the project.
Install dependencies and verify the project runs locally.
Set up pre-commit hooks, linters, or formatters if recommended.
☐ 4. Understand the project's AI policies
Check if the project has AI contribution policies (e.g. AI disclosure
requirements).
If the project doesn't mention AI usage, avoid assuming it's acceptable.
Consider asking maintainers if disclosure or other expectations aren't clear.
Preferably disclose AI usage in your PR/issue irrespective of whether the
project's AI policy demands it.
☐ 5. Use AI as your tool, not as your replacement
Use AI to:
Brainstorm ideas or generate draft code snippets.
Automate repetitive tasks (e.g., formatting, testing scaffolding).
Document complex processes (e.g., setup instructions, troubleshooting).
Never submit AI-generated code without:
Reviewing it line-by-line.
Modifying it to fit the project's style.
Make sure you can explain every change you submit, whether it was written by
you or assisted by AI.
☐ 6. Test, validate and document
Run the project's test suite and ensure no errors.
Write new tests for your changes.
Manually verify your changes work as intended.
If using AI-generated tests, review them carefully for edge cases.
Update documentation, examples, or changelog entries if your changes
affect users.
Small feature implementations (labeled "good first issue").
Avoid large or ambiguous changes until you're deeply familiar with the project.
Check for "help wanted" or "needs review" labels in the issue tracker.
☐ 8. Starting to work on an issue
Before raising a PR for an issue, check if there's one or more existing
PRs for it. In case of existing PRs:
If any of those PRs are actively being updated, don't pick the issue as
it could lead to duplicate effort. Find another issue to work on.
If all those PRs are inactive, then you could pick up the issue after
confirming with maintainers.
Comment on the issue and confirm with the maintainers before starting work
to avoid duplication.
Example: "I'd like to work on this. May I pick this up? I plan to
implement it like this: [your brief plan]."
☐ 9. Communicate early and often
For larger changes, ask for feedback on your proposed approach before
opening a PR.
Be prepared for your proposed solution to be declined. Not every good idea
aligns with a project's roadmap.
Document your thought process in the issue/PR.
Be responsive to feedback and iterate.
☐ 10. Submitting your contribution
If the project has a pull request template, follow it thoroughly.
Keep your pull request focused on a single logical change whenever possible.
Create a new branch for your changes (e.g., fix/issue-123 ).
Write clear, descriptive commit messages.
Follow the project's conventions (e.g., Fix #123: [description] ).
If you used AI tools and if project encourages commit-level AI disclosure,
use the project's preferred format (e.g. Assisted-by or Co-authored-by ).
If project recommends DCO sign-off, provide one in your commits using
Signed-off-by .
Example DCO sign-off along with AI disclosure with Assisted-By in commit:
Improve error handling in user authentication module
- Added custom error classes for authentication failures
- Implemented retry logic for transient failures
- Updated tests to cover new error scenarios
Assisted-by: GitHub Copilot (for new error scenario tests)
Signed-off-by: Your Name <youremail@address.here>
Include references to relevant issues in your PR description.
In the PR description:
Disclose AI usage (e.g., "This PR includes AI-assisted code generation for
[specific part]").
Explain your process and reasoning.
Link to any related discussions or issues.
Monitor your PR for feedback and be ready to make changes.
Give the reviewers time to respond, avoid repeatedly asking for updates (e.g.
every few hours).
Thank reviewers for their time and feedback.
If your PR is closed, ask for feedback on how to improve, so you can learn
from your mistakes.
Celebrate your contribution, appreciate any help you received from other
community members; open source is a community effort!
Do not submit AI-generated code without reviewing, understanding, and testing
it.
Do not ignore community guidelines or communication norms.
Do not submit large changes without prior discussion and approval.
Do not spam the issue tracker or PR queue with low-quality contributions.
Do not rely on AI in place of understanding the codebase or engaging with
the community.
Projects are welcome to adopt this checklist as NOSLOP.md ,
incorporate parts of it into CONTRIBUTING.md , or simply share this repository
with contributors.
Feel free to raise a PR in case you'd like to add something. Refer to the
checklist above for contribution guidelines. All contributions
should abide by this code of conduct .
How to Responsibly and Effectively Contribute to Open Source Using AI
Maintaining open source in the age of generative AI: Recommendations for maintainers and contributors
Open source succeeds through people as much as code. Building trust,
communicating clearly, and collaborating well are just as important as
writing software.
This checklist has been reviewed with assistance from ChatGPT to improve tone
and grammar consistency. All suggestions were manually reviewed before being
incorporated.
A checklist of no-slop contribution best practices when using AI (or not) in open source projects.
CC0-1.0 license
Code of conduct
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
