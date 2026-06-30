---
source: "https://github.com/tigerless-labs/autoharness"
hn_url: "https://news.ycombinator.com/item?id=48732795"
title: "Show HN: Autoharness – a self-learning, maintaining skill layer for Claude Code"
article_title: "GitHub - tigerless-labs/autoharness: A self-learning skill layer for Claude Code — distills skills from your real sessions, updates them as you work, and prunes the ones that stop getting used. No daemon, no benchmark. · GitHub"
author: "Tigerless_ailab"
captured_at: "2026-06-30T14:51:53Z"
capture_tool: "hn-digest"
hn_id: 48732795
score: 3
comments: 0
posted_at: "2026-06-30T13:53:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Autoharness – a self-learning, maintaining skill layer for Claude Code

- HN: [48732795](https://news.ycombinator.com/item?id=48732795)
- Source: [github.com](https://github.com/tigerless-labs/autoharness)
- Score: 3
- Comments: 0
- Posted: 2026-06-30T13:53:28Z

## Translation

タイトル: Show HN: Autoharness – クロード コードの自己学習、維持スキル レイヤー
記事のタイトル: GitHub - Tigerless-labs/autoharness: Claude Code の自己学習スキル レイヤー — 実際のセッションからスキルを抽出し、作業中に更新し、使用されなくなったスキルを削除します。デーモンもベンチマークもありません。 · GitHub
説明: Claude Code の自己学習スキル レイヤー — 実際のセッションからスキルを抽出し、作業中に更新し、使用されなくなったスキルを削除します。デーモンもベンチマークもありません。 - タイガーレスラボ/オートハーネス

記事本文:
GitHub - Tigerless-labs/autoharness: クロード コードの自己学習スキル レイヤー — 実際のセッションからスキルを抽出し、作業中に更新し、使用されなくなったスキルを削除します。デーモンもベンチマークもありません。 · GitHub
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
ああ、ああ！
の

読み込み中にエラーが発生しました。このページをリロードしてください。
タイガーレスラボ
/
オートハーネス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows エージェント エージェント docs/ アセット docs/ アセット フック フック src/ オートハーネス src/ オートハーネス テスト テスト ツール ツール .gitignore .gitignore .mcp.json .mcp.json ライセンス ライセンス README.md README.md pytest.ini pytest.ini ruff.toml ruff.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
autoharness は、Claude Code の自己学習スキル レイヤーです。本当の自分からスキルを学びます
セッション、ほぼ重複したものを積み重ねるのではなく、同じシナリオのものをマージし、使用中のそれらを更新します。
そして、使用されなくなったものはすべて削除します。そのため、レイヤーはそれ自体でクリーンな状態を維持し、接触のみを維持します。
それ自体が書いたスキル。
同じモデル、異なるハーネス - CORE-Bench (HAL) で 42% → 78%。
ハーネスは多くの作業を行いますが (swyx の Big Model と Big Harness の比較)、それでもまだ再構築されています。
すべてのモデルの世代を手にします。オートハーネスは、その一部であるスキル層が自身を維持できることに賭けています。
PATH に python3 が必要です — オートハーネスは完全に Python として実行されます (サードパーティは使用しません)
依存関係);そのフックと MCP サーバーはそれなしでは起動しません。
/プラグイン マーケットプレイス追加 Tigerless-Labs/autoharness
/プラグインのインストール autoharness@autoharness
Claude Code を再起動し、ライブであることを確認します。
/plugin # autoharness@autoharness — 有効
/mcp # stage_skill — 接続済み
設定ゼロ。セッションを監視し、学習したスキルを .claude/skills/ に格納するようになりました。
背景。ケイデンスとライフサイクルのしきい値は、AUTOHARNESS_* 環境変数を介して調整できます。
クロードプラグインのアンインストールautoh

arness@autoharness # フック + MCP サーバーを停止します
クロード プラグイン マーケットプレイス オートハーネスを削除 # オプション — インストール ソースも削除します
アンインストールしても実行が停止されるだけです。アンインストールしたスキルとそれ自体の状態は、外部に存在します。
プラグインしてディスク上に残ります。これらもクリアするには、その状態ディレクトリ (~/.claude/autoharness/ global,
<repo>/.claude/autoharness/ プロジェクトごと) および .claude/skills/ の下の自己作成スキル (それぞれ
には自分で作成した台帳マーカーが付いているので、自分のものと簡単に見分けることができます。あなた自身のスキルは、
決して触れなかった。
学習パイプラインはホストの横で実行され、そのリコール パスから外れます。シンボルはプレーン ネイティブです。
スキルは、あたかも人間が書いたかのように、ホスト自身の名前と説明のメカニズムによって呼び出されます。
図のソース: docs/assets/pipeline.mmd — 編集後に Pipeline.svg に再レンダリングします。
自己学習スキル レイヤーは、保持されているベンチマークに対して、またはそれ自体の使用に対して検証できます。
オートハーネスは 2 番目の方法で、より安価で、制限のない場所で無制限の作業を行うライブ ホスト上で動作します。
ベンチマークが存在します。
NousResearch/hermes-agent — 研究中
自動スキル作成と記憶統合の設計は、オートハーネスの遵守ベースの強化に役立ちました。
デーモンフリーのテイク。
Claude Code の自己学習スキル レイヤー — 実際のセッションからスキルを抽出し、作業中に更新し、使用されなくなったスキルを削除します。デーモンもベンチマークもありません。
github.com/tigerless-labs/autoharness
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
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

A self-learning skill layer for Claude Code — distills skills from your real sessions, updates them as you work, and prunes the ones that stop getting used. No daemon, no benchmark. - tigerless-labs/autoharness

GitHub - tigerless-labs/autoharness: A self-learning skill layer for Claude Code — distills skills from your real sessions, updates them as you work, and prunes the ones that stop getting used. No daemon, no benchmark. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
tigerless-labs
/
autoharness
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows agents agents docs/ assets docs/ assets hooks hooks src/ autoharness src/ autoharness tests tests tools tools .gitignore .gitignore .mcp.json .mcp.json LICENSE LICENSE README.md README.md pytest.ini pytest.ini ruff.toml ruff.toml View all files Repository files navigation
autoharness is a self-learning skill layer for Claude Code. It learns skills from your real
sessions, merges same-scenario ones instead of stacking near-duplicates, updates them in use,
and prunes any that stop getting used — so the layer stays clean on its own , touching only
the skills it wrote itself .
Same model, different harness — 42% → 78% on CORE-Bench ( HAL ).
The harness does much of the work (swyx's Big Model vs Big Harness ), yet it's still rebuilt by
hand every model generation. autoharness bets one slice of it — the skill layer — can maintain itself.
Requires python3 on your PATH — autoharness runs entirely as Python (zero third-party
dependencies); its hooks and MCP server won't fire without it.
/plugin marketplace add tigerless-labs/autoharness
/plugin install autoharness@autoharness
Restart Claude Code, then check it's live:
/plugin # autoharness@autoharness — enabled
/mcp # stage_skill — connected
Zero config. It now watches your sessions and lands learned skills into .claude/skills/ in the
background. Cadence and lifecycle thresholds are tunable via AUTOHARNESS_* environment variables.
claude plugin uninstall autoharness@autoharness # stops the hooks + MCP server
claude plugin marketplace remove autoharness # optional — also drops the install source
Uninstalling only stops it from running — the skills it landed and its own state live outside the
plugin and stay on disk. To clear those too, delete its state dir ( ~/.claude/autoharness/ global,
<repo>/.claude/autoharness/ per project) and the self-authored skills under .claude/skills/ (each
carries a self-authored ledger marker, so they're easy to tell from yours). Your own skills are
never touched.
A learning pipeline runs beside the host and stays off its recall path — symbols are plain native
skills, recalled by the host's own name-and-description mechanism as if a human had written them.
Diagram source: docs/assets/pipeline.mmd — re-render to pipeline.svg after editing.
A self-learning skill layer can be validated against a held-out benchmark, or against its own use.
autoharness takes the second — cheaper, and it works on a live host doing open-ended work where no
benchmark exists.
NousResearch/hermes-agent — studying its
auto-skill-creation and memory-consolidation design helped sharpen autoharness's adherence-based,
daemon-free take.
A self-learning skill layer for Claude Code — distills skills from your real sessions, updates them as you work, and prunes the ones that stop getting used. No daemon, no benchmark.
github.com/tigerless-labs/autoharness
Topics
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
