---
source: "https://github.com/dshakes/compass"
hn_url: "https://news.ycombinator.com/item?id=48623314"
title: "Compass – guardrails and a hard budget cap for AI coding agents"
article_title: "GitHub - dshakes/compass: Developer-grade Claude Code + Codex configuration: cost-tiered subagents, workflow commands, guardrail hooks, MCP parity, and an installable plugin/marketplace. · GitHub"
author: "chandu1221"
captured_at: "2026-06-21T23:28:36Z"
capture_tool: "hn-digest"
hn_id: 48623314
score: 1
comments: 0
posted_at: "2026-06-21T22:38:35Z"
tags:
  - hacker-news
  - translated
---

# Compass – guardrails and a hard budget cap for AI coding agents

- HN: [48623314](https://news.ycombinator.com/item?id=48623314)
- Source: [github.com](https://github.com/dshakes/compass)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T22:38:35Z

## Translation

タイトル: Compass – AI コーディング エージェントのガードレールとハード予算上限
記事のタイトル: GitHub - dshakes/compass: 開発者グレードの Claude Code + Codex 構成: コスト階層型サブエージェント、ワークフロー コマンド、ガードレール フック、MCP パリティ、およびインストール可能なプラグイン/マーケットプレイス。 · GitHub
説明: 開発者グレードの Claude Code + Codex 構成: コスト階層型サブエージェント、ワークフロー コマンド、ガードレール フック、MCP パリティ、およびインストール可能なプラグイン/マーケットプレイス。 - dshakes/コンパス

記事本文:
GitHub - dshakes/compass: 開発者グレードの Claude Code + Codex 構成: コスト階層型サブエージェント、ワークフロー コマンド、ガードレール フック、MCP パリティ、およびインストール可能なプラグイン/マーケットプレイス。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
シェイク
/
コンパス
公共
通知
あなたは

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
186 コミット 186 コミット .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .github .github Formula Formula アセット アセット ビン ビン クロード クロード コーデックス コーデックス デモ デモ ドキュメント ドキュメント mcp mcp プラグイン プラグイン ルーター ルーター スクリプト スクリプト sdlc sdlc テンプレート テンプレート .gitignore .gitignore .gitleaks.toml .gitleaks.toml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md GEMINI.md GEMINI.md ライセンス ライセンス Makefile Makefile PRIVACY.md PRIVACY.md README.md README.md SECURITY.md SECURITY.md gemini-extension.json gemini-extension.json install.sh install.sh Quickstart.sh Quickstart.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントのガードレールと厳しい予算の上限。
バジェットゲート · ガードレール 100/100 · 最大 61% 安いルーティング · 署名されたリリース · 100% ローカル · テレメトリなし · 常にマージ
実際のセッション、編集なし: コストは 0.35 ドルまで上昇し、その後、次のアクションは上限 0.05 ドルで停止され、それ以上の費用がかかる前に停止されます。
compass は、Claude Code、Codex、Gemini 用のローカルファースト構成レイヤーで、予算の消費、安全でないコマンドの実行、未検証のコードのマージという、エージェントが行うべきではない 3 つのことを阻止します。 COMPASS_MAX_USD=5 に設定すると、セッションは上限でハードストップします。致命的なコマンドは実行前にブロックされ、ガードレール ポリシーは CI で 100/100 のスコア付けされます (アサートされません)。一度インストールすれば、常にマージされます。
# カール|sh なし、完全に元に戻せます。エージェントで任意のリポジトリを開くだけです。
git clone https://github.com/dshakes/compass ~ /compass && cd ~ /compass && ./quickstart.sh
# または、Claude Code: /plugin Marketplace 内に dshakes を追加します

/コンパス
▶ 動作を確認する · 異なる理由 · 自己修正 PR ループ · インストール · 同梱品 · 📚 ドキュメント
⭐ 人々がスクリーンショットした部分: 独自の PR を修正します。
プル リクエストを開くと、compass がレビュー、セキュリティ チェック、テストの実行、2 番目のモデルとの相互監査を行い、グリーンになるまで独自の修正をプッシュします。ただマージするだけです。
アイデアは 1 行にまとめられています。ループは作業単位です。ワンショットエージェントは、最初の間違った答えで停止します。コンパス ループ — 生成 → テスト → 批評 → 修正 → ゲートに対して繰り返す — そのため、品質は 1 回の幸運なプロンプトではなく、反復から生まれます。同じ閉ループで単一の PR またはリポジトリ全体を一晩実行します。 (トークンなしで 30 秒以内にローカルで試してみてください。↓を見てください。)
なぜ違うのか — 雰囲気ではなく測定によるもの
すべての AI エージェント構成は「安全」かつ「安価」であると主張しています。コンパスはその数字をあなたに渡し、懐疑的な人でも 30 秒以内にその数字を再現できるようにします。誰もが同じモデルを持っています。エッジは信頼できる構成であり、別の機能リストではありません。 4 つのクレーム、4 つのコマンド:
🛡 スコア付きのガードレール。壊滅的なコマンドとシークレットの書き込みは実行前にブロックされ、ポリシーはアサートされずに評価ゲートされます。 (人間の言葉で言えば、エージェントがあなたのマシンを削除したり、キーを漏洩させたりすることはなく、それがどれだけうまく機能するかを証明できます。)
コンパスベンチ # → ガードレール 精度/再現率 100% (61 ケースコーパス)、ルーター 96.9% — CI で
# その後、エージェントに `rm -rf /` または .env を書き込むように依頼します → 拒否; `rm -rf ./build` → 許可されます
📉 測定されたコストルーティング。安い仕事は安いモデルに送られます。評価セットに対してスコア付けされ、公正な組み合わせで ~98% の品質を持つ全 Opus よりも ~61% 安くなります。 (人間の言葉で言えば、タイプミスを修正するために Opus の価格を支払うのをやめるということです。)
コンパスルート「認証モデルの再設計」 # → opus
コンパスルート「タイプミスを修正」 # → 俳句
💸A

実際にそれを阻止する予算の上限。使用状況トラッカーは支出を報告します。コンパスはそれを強制します。ドルの上限を設定すると、その上限に達すると次のツール呼び出しの前にセッションが停止されます。 (人間の言葉で言えば、エージェントはあなたの不在中に黙って 40 ドル札を計算することはできません。あなたの番号で止まります。)
import COMPASS_MAX_USD=5 # このセッションは $5 でハードストップします — エージェントは警告されるだけでなくブロックされます
compass put --max-usd 5 # スケジュールされた/フリート実行の場合、元帳と同じ上限
🔏 サプライチェーンを確認できます。リリースにはキーレス SLSA の来歴が含まれるため、改ざんされたダウンロードや類似したダウンロードは拒否されます。 (人間の言葉で言えば、あなたがインストールしたコードが私が出荷したコードであることを証明できます。)
コンパス検証 v0.17.2 # → ✓ 出所検証済み
🧪 レッドチームの抵抗、測定。プロンプトインジェクション (直接/間接/ペースト)、CLAUDE.md ポイズニング、ローカルの安全性オーバーライド、マルウェアと安全でないコード - CI でゲートするラベル付きコーパスに対してスコア付けされ、オプションでマネージド ガードレール サービス (Webhook、Bedrock、Azure) へのエスカレーションが行われます。 (人間の言葉で言えば、毒されたリポジトリや Web ページがエージェントを静かに敵に回すことはできません。)
compass redteam # → インジェクション コーパス 100% P/R、次にこのリポジトリの CLAUDE.md/MCP/settings をスキャンします
サービスもテレメトリも --dangerously-skip-permissions もありません。 git pull して更新します。安全に所有できない作業は差し戻され、マージは維持されます。
まず小さな信念の飛躍、つまり統治の瞬間、それからそれを感じ、次に証拠を見て、それがどのように機能するかを見てください。
0 · 注釈付きの予算上限 ​​— クリーンなウォークスルーとして、ヒーロー クリップと同じハードストップ ($1.80 ✓ → $4.10 ✓ → $5.00 停止)。使用状況トラッカーは支出を報告します。コンパスはそれを強制します:
1 · 日常の雰囲気 - ガードレール、コストを意識したステータス ライン、ループ、乗組員を約 25 秒で再現:
2 · 見出し、再考

al PR — ブロッキング バグとレッド テスト。PR が緑色になるまで独自の修正をプッシュします (その後、待機します)。
3 · そのループの仕組み — レビュー · セキュリティ · テスト · Codex 相互監査を並行して実行します。障害となっている検出結果は、緑色になるまで自動的に修正され、再レビューされ (丸いキャップで表示されます)、緑色になったら停止します。
~/compass/sdlc/orchestrate.sh "<タスク>" (トークンなし) を使用して 30 秒以内にローカルで実行するか、PR ごとに GitHub ループを接続します。 → 仕組み・再現
そして、毎日のステータスラインは静かにスコアを維持するので、それが維持されるのを観察することができます。
Opus 4.8 · myrepo · main* · 45k ctx · $1.23 · 🧭 🛡1 🧹2 💡1 📉~$1.65
セッションの支出、そして今日のコンパスアクティビティ: 🛡 フットガンのブロック · 🧹 ファイルのフォーマット · 💡 ポリシーのナッジ · 📉~$ の推定節約額と全 Opus の比較。各部分は、報告する必要がある場合にのみ表示されます。マシンからは何も残りません。
ここでの自律性は、1 つの大きな魔法のボタンではありません。それは、4 つのスケールで適用される同じ閉ループです。それぞれがゲートに「完了」と表示されるまで実行され、人間のところで停止します。これがすべてのトリックです。ゲートの下での反復は、単一の確信のある推測に打ち勝ちます。
すべてのループは同じ方法で終了します。つまり、マージします。その門は決して動くことはありません。
前提条件 — 必要なもの (そして必要でないもの)
あなたが欲しいのは…
必要です
トークン？
設定、ガードレール、CLI、サブエージェント (ローカル)
クロード コード (または Codex/Gemini) + git
なし
自律的な PR ループ (GitHub Actions)
Actions + gh 、モデル認証 ( CLAUDE_CODE_OAUTH_TOKEN または ANTHROPIC_API_KEY )、および SDLC_BOT_TOKEN (粒度の細かい PAT) を含むリポジトリ。これにより、ループを連鎖させることができます。
はい
キーレス ループ (自己ホスト型ランナー)
コンパス + SDLC_BOT_TOKEN というラベルの付いたランナー
PATのみ
フリート (すべてのリポジトリ)
FLEET_TOKEN + FLEET_MAINTAINER
はい
1 つのコマンドが GitHub ループに接続します: ~/compass/sdlc/setup.sh --all (ラベル + ワークフロー + CODEOWNERS + シークレット + ブランチ保護)。それなし

SDLC_BOT_TOKEN ループはまだ実行されますが、修正後に自動的に再起動されないだけです。 → 完全な SDLC セットアップ
ぴったりのドアをお選びください - すべてリバーシブル、バージョン固定可能、カールなし |し。 AI アシスタント (Claude Code、Codex/Gemini はオプション) + git が必要です。マニュアル、ガードレール、スタッフ、および CLI を取得するための API キーはありません。
🍺 Homebrew — 管理およびバージョン管理
醸造タップ dshakes/compass https://github.com/dshakes/compass
brew install dshakes/compass/compass # 最新リリース · --HEAD でメインを追跡します
compass クイックスタート # プレビューして質問し、~/.claude に接続します
📦 Git クローン — 構成を所有して編集します (推奨)
git clone https://github.com/dshakes/compass ~ /compass && cd ~ /compass
git checkout v0.17.2 # オプション: メインではなくリリースにピン留めする
./quickstart.sh # すべての変更をプレビューし、最初に確認し、完全に元に戻すことができます
🧩 Claude Code プラグイン — ターミナルなし (チームに最適)
/プラグイン マーケットプレイス dshakes/compass を追加
/plugin install core@compass
🛠️ 手動: ドライラン (プレビュー) の作成 → インストールの作成 → ドクターの作成 。 Symlink インストールとは、git pull / brew upgrade によってすべてが更新されることを意味します。 make uninstall は追加したものだけを削除します。 → チーム展開
1 つの構成ですべてのエージェント - ネイティブ インストール
あらゆる種類のユーザー向け: 1 行のマーケットプレイス/拡張機能のインストール (ターミナルなし)、またはファイルを所有したい場合は make install を実行します。同じ操作マニュアル + MCP サーバー、各ツールが期待する方法:
CLAUDE.md、AGENTS.md、GEMINI.md は 1 つのファイル (シンボリックリンク) であり、Claude/Codex プラグイン マニフェストと Gemini 拡張機能は 1 つのソースから生成され、CI チェックされます ( scripts/check-vendor.sh )。そのため、git pull はすべてのエージェントを一度に更新し、マニフェストがドリフトすることはありません。
マーケットプレイス/拡張機能マニフェストは、各ベンダーの文書化されたスキーマと一致しており、CI で構造が検証されています。ライブ インストールは手動で検証されます — g

emini 拡張機能のインストール (gemini 0.26.0) と codex プラグイン マーケットプレイスの追加 (codex 0.130.0) はどちらもこのリポジトリに対して成功しますが、CI では実行されません (これらの CLI はランナーにありません)。
コンパスドクター # インストールを検証します — 「0 エラー」が予想されます
コンパスステータス # ここではコンパスがアクティブになっていますか? 何がロードされていますか?
次に、通常どおり Claude Code を開くだけです。マニュアル、ガードレール、サブエージェント、コマンド、ステータス行はすでに読み込まれています。すぐに実感してください。危険なコマンド (ブロックされています) を要求するか、diff で /review を実行するか、ルート "<task>" をコンパスして、選択される層を確認してください。トークンもサインアップも必要ありません。
以下のすべては 1 回のインストールまたは 1 回のオプトイン後に有効になります。上記の自律ループはその上にあります。 README は売り物です。ドキュメントで説明されています。各行は詳細にリンクしています。
実行前から信頼できるように構築されており、その限界について正直です。
取り返しのつかないものを所有しているのはあなたです。エージェントは準備をします。人間がプッシュ、マージ、デプロイします。必須のチェックとコード所有者の承認によって強制されます。「本番環境にマージ」ボタンはありません。
可読性と可逆性。カールなし |し。インストーラーは置き換えるものをバックアップし、べき等であり、アンインストールする場合は追加したもののみを削除します。 main ではなく、タグを固定します。
ガードレールはフットガンを減らします。それらはセキュリティ境界ではありません。最小権限の資格情報を保持し、差分を確認してください。 (信頼できないコードの場合、コンパス サンドボックスが実際の境界になります。)
レッドチームの強化は多層防御であり、免疫ではありません。プロンプトインジェクション (direct/i) について警告します。

[切り捨てられた]

## Original Extract

Developer-grade Claude Code + Codex configuration: cost-tiered subagents, workflow commands, guardrail hooks, MCP parity, and an installable plugin/marketplace. - dshakes/compass

GitHub - dshakes/compass: Developer-grade Claude Code + Codex configuration: cost-tiered subagents, workflow commands, guardrail hooks, MCP parity, and an installable plugin/marketplace. · GitHub
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
dshakes
/
compass
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
186 Commits 186 Commits .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .github .github Formula Formula assets assets bin bin claude claude codex codex demo demo docs docs mcp mcp plugins plugins router router scripts scripts sdlc sdlc templates templates .gitignore .gitignore .gitleaks.toml .gitleaks.toml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md GEMINI.md GEMINI.md LICENSE LICENSE Makefile Makefile PRIVACY.md PRIVACY.md README.md README.md SECURITY.md SECURITY.md gemini-extension.json gemini-extension.json install.sh install.sh quickstart.sh quickstart.sh View all files Repository files navigation
Guardrails and a hard budget cap for your AI coding agent.
budget gate · guardrails 100/100 · ~61% cheaper routing · signed releases · 100% local · no telemetry · you always merge
Real session, no edits: the cost climbs to $0.35, then the next action is HALTED at the $0.05 cap — before it spends more.
compass is a local-first config layer for Claude Code, Codex & Gemini that stops your agent from doing three things it shouldn't — burning your budget, running unsafe commands, and merging unverified code. Set COMPASS_MAX_USD=5 and the session hard-stops at the cap; catastrophic commands are blocked before they run, and the guardrail policy is scored 100/100 in CI — not asserted. You install it once, and you always merge.
# no curl|sh, fully reversible — then just open any repo in your agent
git clone https://github.com/dshakes/compass ~ /compass && cd ~ /compass && ./quickstart.sh
# or, inside Claude Code: /plugin marketplace add dshakes/compass
▶ See it work · Why it's different · The self-fixing PR loop · Install · What's in the box · 📚 Docs
⭐ The part people screenshot: it fixes its own PRs.
Open a pull request and compass reviews it, security-checks it, runs the tests, cross-audits it with a second model — then pushes its own fixes until it's green. You just merge.
The idea in one line: the loop is the unit of work. A one-shot agent stops at its first wrong answer. compass loops — generate → test → critique → fix → repeat against a gate — so quality comes from iteration, not one lucky prompt. The same closed loop runs a single PR, or your whole fleet of repos overnight. (Try it locally in 30s, no tokens — watch it ↓ .)
Why it's different — measured, not vibes
Every AI-agent config claims "safe" and "cheap." compass is the one that hands you the number — and lets a skeptic reproduce it in 30 seconds. Everyone has the same models; the edge is configuration you can trust , not another feature list. Four claims, four commands:
🛡 Guardrails with a score. Catastrophic commands and secret writes are blocked before they run — and the policy is eval-gated, not asserted. (In human terms: it won't let the agent delete your machine or leak your keys, and it can prove how well.)
compass bench # → guardrail 100% precision/recall (61-case corpus), router 96.9% — in CI
# then ask the agent to `rm -rf /` or write a .env → denied; `rm -rf ./build` → allowed
📉 Cost routing that's measured. Cheap work goes to cheap models — scored against an eval set, ~61% cheaper than all-Opus at ~98% quality on a fair mix. (In human terms: it stops paying Opus prices to fix a typo.)
compass route " redesign the auth model " # → opus
compass route " fix a typo " # → haiku
💸 A budget ceiling that actually stops it. Usage trackers report spend; compass enforces it — set a dollar cap and the session is halted before the next tool call once it's reached, live. (In human terms: an agent can't quietly run up a $40 bill while you're away — it stops at your number.)
export COMPASS_MAX_USD=5 # this session hard-stops at $5 — the agent is blocked, not just warned
compass spend --max-usd 5 # the same ceiling on the ledger, for scheduled / fleet runs
🔏 Supply chain you can verify. Releases carry keyless SLSA provenance, so a tampered or look-alike download is rejected. (In human terms: you can prove the code you installed is the code I shipped.)
compass verify v0.17.2 # → ✓ provenance verified
🧪 Red-team resistance, measured. Prompt-injection (direct/indirect/paste), CLAUDE.md poisoning, local safety-override, malware & insecure-code — scored against a labeled corpus that gates in CI, with optional escalation to a managed guardrails service (webhook · Bedrock · Azure). (In human terms: a poisoned repo or web page can't quietly turn your agent against you.)
compass redteam # → injection corpus 100% P/R, then scans THIS repo's CLAUDE.md/MCP/settings
No service, no telemetry, no --dangerously-skip-permissions ; git pull to update. The work it can't safely own, it hands back — you keep the merge.
Smallest leap of faith first — the governance moment , then feel it , then see the proof , then see how it works.
0 · The budget ceiling, annotated — the same hard-stop as the hero clip, as a clean walkthrough ($1.80 ✓ → $4.10 ✓ → $5.00 HALTED). Usage trackers report spend; compass enforces it:
1 · The day-to-day feel — guardrails, the cost-aware status line, the loop, and the crew, in ~25 seconds:
2 · The headline, on a real PR — a Blocking bug and red tests, and it pushes its own fix until the PR is green (then waits for you):
3 · How that loop works — review · security · tests · Codex cross-audit run in parallel; Blocking findings get auto-fixed and re-reviewed (round-capped) until green, then it stops at you:
Run it locally in 30s with ~/compass/sdlc/orchestrate.sh "<task>" (no tokens), or wire the GitHub loop for every PR. → how it works · reproduce it
And the everyday status line quietly keeps score, so you watch it earn its keep:
Opus 4.8 · myrepo · main* · 45k ctx · $1.23 · 🧭 🛡1 🧹2 💡1 📉~$1.65
session spend, then today's compass activity: 🛡 footguns blocked · 🧹 files formatted · 💡 policy nudges · 📉~$ estimated saved vs all-Opus. Each piece shows only when there's something to report; nothing leaves your machine.
Autonomy here isn't one big magic button — it's the same closed loop applied at four scales. Each runs until a gate says "done," then stops at a human. That's the whole trick: iteration under a gate beats a single confident guess.
Every loop ends the same way — you merge. That gate never moves.
Prerequisites — what you need (and what you don't)
You want…
You need
Tokens?
The config, guardrails, CLI, subagents (local)
Claude Code (or Codex/Gemini) + git
None
The autonomous PR loop (GitHub Actions)
A repo with Actions + gh , model auth ( CLAUDE_CODE_OAUTH_TOKEN or ANTHROPIC_API_KEY ), and SDLC_BOT_TOKEN (fine-grained PAT) so the loop can chain
Yes
Keyless loop (self-hosted runner)
A runner labeled compass + SDLC_BOT_TOKEN
PAT only
The fleet (every repo)
FLEET_TOKEN + FLEET_MAINTAINER
Yes
One command wires the GitHub loop: ~/compass/sdlc/setup.sh --all (labels + workflows + CODEOWNERS + secrets + branch protection). Without SDLC_BOT_TOKEN the loop still runs — it just won't auto-re-fire after a fix. → full SDLC setup
Pick the door that fits — all reversible, version-pinnable, no curl | sh . You need an AI assistant ( Claude Code ; Codex/Gemini optional) + git . No API keys to get the manual, guardrails, crew, and CLI.
🍺 Homebrew — managed & versioned
brew tap dshakes/compass https://github.com/dshakes/compass
brew install dshakes/compass/compass # latest release · --HEAD to track main
compass quickstart # previews, asks, then wires it into ~/.claude
📦 Git clone — own & edit your config (recommended)
git clone https://github.com/dshakes/compass ~ /compass && cd ~ /compass
git checkout v0.17.2 # optional: pin to a release instead of main
./quickstart.sh # previews every change, asks first, fully reversible
🧩 Claude Code plugin — no terminal (ideal for a team)
/plugin marketplace add dshakes/compass
/plugin install core@compass
🛠️ By hand: make dry-run (preview) → make install → make doctor . Symlink install means git pull / brew upgrade updates everything; make uninstall removes only what it added. → Team rollout
One config, every agent — native installs
For every kind of user: a one-line marketplace/extension install (no terminal), or make install if you'd rather own the files. Same operating manual + MCP servers, the way each tool expects them:
CLAUDE.md · AGENTS.md · GEMINI.md are one file (symlinks), and the Claude/Codex plugin manifests + Gemini extension are generated from one source and CI-checked ( scripts/check-vendor.sh ) — so a git pull updates every agent at once and a manifest can't drift.
The marketplace/extension manifests match each vendor's documented schema and are structure-validated in CI. The live install is manually verified — gemini extensions install (gemini 0.26.0) and codex plugin marketplace add (codex 0.130.0) both succeed against this repo — but isn't run in our CI (those CLIs aren't in the runner).
compass doctor # validate the install — expect "0 error"
compass status # is compass active here, and what's loaded?
Then just open Claude Code as usual — the manual, guardrails, subagents, commands, and status line are already loaded. Feel it in a minute: ask for a dangerous command (blocked), run /review on your diff, or compass route "<task>" to see the tier it picks. No tokens, no signup for any of it.
Everything below is on after one install or a single opt-in — the autonomous loops above sit on top of this. The README sells; the docs explain — each row links to the detail.
Built to be trusted before it's run — and honest about its limits.
You own the irreversible. Agents prepare; humans push, merge, deploy. Required checks + a code-owner approval enforce it — there's no "merge to prod" button.
Readable & reversible. No curl | sh . The installer backs up what it replaces, is idempotent, and make uninstall removes only what it added. Pin a tag, not main .
Guardrails reduce footguns; they are not a security boundary. Keep least-privilege credentials and review your diffs. (For untrusted code, compass sandbox is a real boundary.)
Red-team hardening is defense-in-depth, not immunity. It warns on prompt-injection (direct/i

[truncated]
