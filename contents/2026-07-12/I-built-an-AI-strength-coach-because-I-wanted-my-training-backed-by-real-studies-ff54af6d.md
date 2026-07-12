---
source: "https://github.com/clementrx/Performance-agent/tree/main"
hn_url: "https://news.ycombinator.com/item?id=48879889"
title: "I built an AI strength coach because I wanted my training backed by real studies"
article_title: "GitHub - clementrx/Performance-agent: Open-source AI Strength & Conditioning coach for your terminal. Runs in Claude Code / Gemini CLI / Codex — the LLM narrates, a deterministic engine calculates, and it won't lie about unrealistic goals. · GitHub"
author: "clementrx"
captured_at: "2026-07-12T11:05:58Z"
capture_tool: "hn-digest"
hn_id: 48879889
score: 2
comments: 1
posted_at: "2026-07-12T10:11:34Z"
tags:
  - hacker-news
  - translated
---

# I built an AI strength coach because I wanted my training backed by real studies

- HN: [48879889](https://news.ycombinator.com/item?id=48879889)
- Source: [github.com](https://github.com/clementrx/Performance-agent/tree/main)
- Score: 2
- Comments: 1
- Posted: 2026-07-12T10:11:34Z

## Translation

タイトル: 実際の研究に裏付けられたトレーニングが必要だったので、AI ストレングスコーチを構築しました
記事のタイトル: GitHub - clementrx/Performance-agent: 端末用のオープンソース AI ストレングス & コンディショニング コーチ。 Claude Code / Gemini CLI / Codex で実行 — LLM がナレーションを行い、決定論的エンジンが計算し、非現実的な目標について嘘をつきません。 · GitHub
説明: 端末用のオープンソース AI ストレングス & コンディショニング コーチ。 Claude Code / Gemini CLI / Codex で実行 — LLM がナレーションを行い、決定論的エンジンが計算し、非現実的な目標について嘘をつきません。 - clementrx/パフォーマンスエージェント

記事本文:
GitHub - clementrx/Performance-agent: 端末用のオープンソース AI ストレングス & コンディショニング コーチ。 Claude Code / Gemini CLI / Codex で実行 — LLM がナレーションを行い、決定論的エンジンが計算し、非現実的な目標について嘘をつきません。 · GitHub
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
{{

メッセージ }}
クレメントレックス
/
パフォーマンスエージェント
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
205 コミット 205 コミット .github/ workflows .github/ workflows docs docs 例 例 スキル スキル src/ Performance_agent src/ Performance_agent テスト テスト .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md RELEASING.md RELEASING.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
🏋️ 科学研究に基づいた初のオープンソース AI ストレングス & コンディショニング コーチ。
英語、フランス語、スペイン語、ドイツ語、イタリア語
AI エージェント CLI (Claude Code など、チャットするターミナル プログラム) 内で実行されます。
Gemini CLI、または Codex — これを、設計、説明、
トレーニングを監視し、適応させます。バックエンドなし、API キーなし、ホスティングなし、追加要素なし
サブスクリプション。そして、あなたの目標が非現実的である場合、それは真実を教えてくれます。
$クロード
> 10Kを55分で走ります。 12週間後には35時を目指したいです。
🏋️ コーチ: 実現可能性エンジンを使用してあなたの目標を評価しました。
確率: 0.2% — 正直に言うと、これは非現実的です。
36% の改善が必要です - 週あたり約 3% の継続
12週間。初心者は週に約 1% を維持します。
対案: 12 週間で 46:30 (確率約 78%)、
それから再評価します。そのプログラムを構築してみませんか?
なぜ別の AI フィットネス コーチが必要なのでしょうか?この人はあなたに嘘をつかないから
LLM フィットネス コーチには 2 つの失敗モードがあります。科学的参考資料を発明することと、
聞きたいことを言ってください。 PerformanceAgent は設計されているため、どちらも可能ではありません

ブル:
LLM がナレーションを行い、エンジンが計算します。あらゆる数値 - 実現可能性の確率、
レース予測、トレーニング負荷、ピリオダイゼーションウェーブは、決定論的なものから来ています。
プロパティテスト済みの Python エンジン。エージェントは計算を説明します。決して計算はしません。
引用が幻覚になることはありません。コーチは、チームによって返された研究のみを引用することができます。
ローカル証拠コーパス (等級付け、DOI/PMID 検証済み)。 PDF レンダラーは、どのような場合でもハードフェールします。
コーパスにない参照。
データはクラウドではなくファイルです。プロフィール、プログラム、セッションログ、チェックイン
markdown/YAML のプレーンなディレクトリに存在し、読み取り、編集、差分、同期が可能です。
PerformanceAgent は、ユーザーが開くアプリではありません。AI エージェント CLI にプラグインされます。接続したら
平易な言葉で話しかけるだけです。設定ファイルやコマンドを覚える必要はありません。
クロード コードを使用したことがありませんか?まずインストールしてください:
カール -fsSL https://claude.ai/install.sh |バッシュ
(完全な手順: code.claude.com/docs )。
紫外線も必要です - それ
適切な Python バージョンを自動的に取得します。他にインストールするものはありません。
ステップ 1 — コーチを接続します。任意の端末からこれを 1 回実行します。
クロード mcp add パフォーマンス エージェント -s ユーザー \
--env PERFORMANCE_AGENT_HOME= ~ /athlete-data -- uvx パフォーマンス エージェント
これにより、コーチの「頭脳」（エンジン、科学ライブラリー、あなたの未来）が登録されます。
アスリート プロフィール) をクロード コードが呼び出すことができるツールとして使用します。 -s ユーザーは任意の場所で使用できるようにします
後でクロードを開くフォルダーです。 ~/athlete-data は単なる推奨パスです - 選択してください
どのフォルダーでも、まだ存在する必要はありません。最初に保存するときにコーチによって作成されます。
何か。ここでは、すべてのデータがプレーン ファイルとして保存されます。どこにも何も送信されません。
ステップ 2 — コーチングの方法を教えます。ステップ 1 では、クロードにツール (数学、
データ）。このステップでは、いつ何を質問するか、いつ質問するかというコーチング プロトコルを提供します。
正直に言うと

目標、プログラムの構築方法:
git clone -- Depth 1 https://github.com/clementrx/Performance-agent
mkdir -p ~ /.claude/skills
cp -R パフォーマンスエージェント/スキル/ * ~ /.claude/スキル/
ステップ 3 — Claude Code を完全に終了し、再起動します。新しいツールはクロードの場合にのみロードされます。
セッションの開始: 開いているセッションを完全に閉じて、claude を再度実行します。
機能したことを確認します。新しいセッションで次のように質問します。
> パフォーマンス エージェント ツールをリストします。
47 個のツールが表示されるはずです。そうであれば、もう終わりです。話しかけるだけです。
ターミナルを開いてエージェント (claude) を起動します。
目標をわかりやすい言葉で言いましょう。どんな言語でも構いません。 「10キロを走りたい
50 分以内」、「Hyrox に注ぐ準備をする」。
コーチの質問に答えます。初回は、短いオンボーディングが実行されます (現在の
レベル、履歴、スケジュール、装備）を確認し、プロフィールを保存します。
正直な判決を受けてください。実現可能性エンジンが目標をスコアリングします。それが切れている場合
到達すると、実際の確率と対案が得られます。
「はい」と答えます — プログラムは、ピリオド化されて、 Athletic/programs/program-v1.md に書き込まれます。
デロードとテーパを伴うサイクルに移行し、すべての処方にはその目的があり、
証拠グレード（★★★★★→★☆☆☆☆）と検証された引用。
各ワークアウト後に、「Logged: 8 km easy, RPE 4」という 1 行で記録します。
数週間ごとに「チェックイン」と言います。コーチは計画と現実を比較します。
トレーニング負荷を引用し、フラグをドリフトさせます。
人生で何かが起こったとき - 「足首をひねった」、「次は 2 回しかトレーニングできない」
「week」 - プログラムは、理由とともに適応およびバージョン管理されます (program-v2.md)。
監査証跡に記録されます。
以下の実現可能性の数値は決定論的エンジンから得られたものです。自分で実行して、
同じものを入手できます。結果として得られるプログラム ファイルを確認したいですか?これらのそれぞれ
3 つの会話は完全なプログラムで終わります。
例/ 、すべて f

私たちの言語。
> 10Kを50分以内で走りたいです。 5ヶ月あります。私を指導してください。
🏋️ コーチ: その目標を判断する前に 2 つの事実があります: 現在の 10 キロのタイム、
どれくらいの間、継続的にトレーニングを続けていますか？
> 先月は55:12で走りました。ランニング歴は約1年、週3回。
🏋️ コーチ: 21 週間の期間での実現可能性:
確率: 84% — 現実的であり、取り組む価値があります。
ランナーが
トレーニング年齢は週に約 1% を維持できます。快適なマージン。
プログラムを構築してみませんか?
> はい。
📄 Athletic/programs/program-v1.md に書き込まれたプログラム
21 週間、週 3 回の実行: インターバル セッション 1 回、テンポ 1 回、ロング セッション 1 回
実行 — デロードと 10 日間の漸減で期間化されます。
2. 3 か月で Hyrox を準備
> 3 か月で Hyrox にサインアップしました。計画を立ててください。
🏋️ コーチ: ハイロックスは 8 km のランニングで、8 つの強化ステーションに分かれています —
だから私たちは両方を一緒にトレーニングします。 3 つの事実: あなたの 10,000 時間 (または最高)
推測）、あなたの筋力の背景、トレーニングできる週の日数は何ですか？
> 10K、約 52 分、週に 2 回リフトを行っており、4 セッション行うことができます。
🏋️ コーチ: 力強いファーストフィニッシュに向けた良いベース。正直なところ：12週間
十分に仕上げて楽しむには十分です。表彰台のタイムを追いかけるなら
もっと長い滑走路が必要だ――むしろ今言いたい。
4 日間の計画: ランニングインターバル日 1 日、ステーション強化日 1 日
(そりワーク、ランジ、ウォールボール)、「妥協したランニング」レンガ 1 つ
(走る → 駅 → 走る、Hyrox が実際にテストするスキル)、簡単な 1 つ
ロングラン。 4週目と8週目に負荷を軽減し、レース週ごとに徐々に減らします。建てますか？
> はい。
📄 Athletic/programs/program-v1.md に書き込まれたプログラム
3. 3か月で腹筋が見え、二の腕が大きくなる
> 3か月以内に目に見える腹筋と太い腕が欲しいです。
🏋️ コーチ: これらは 2 つの異なる生理学的目標です — 腕が伸びる
ボリュームと十分なカロリー。ローで腹筋ショー

w体脂肪、これは
通常は赤字を意味します。約束する前に：身長、体重、
トレーニング経験と週あたりのセッション数は?
> 1.78 m、82 kg、リフティング歴 2 年、週 4 回のセッション。
🏋️ コーチ: 正直な判断: 12 週間で真の進歩を遂げることができます。
両方ですが、両方を最大化するわけではありません - 再構成は 2 年と遅いです
トレーニングの。現実的なバージョン:
- 適度なカロリー不足で高たんぱく質目標（
栄養士のスキルによって数値化され、安全対策も施されています）。
- 上半身の優先順位: 腕、重いハードセットを週 12 ～ 16 セット
不足している強度を保護するために複合作業を続けました。
- リフティング 3 日 + コンディショニング 1 日。
- 12週目に予想される: 腕が明らかに太くなり、腹筋が「輪郭を描く」ようになった場合
体脂肪率は約 15 ～ 18% から始まります。それより痩せると時間がかかりますが、
推測させるのではなく、チェックインのたびにそう言います。
プログラムが欲しいですか？
> はい。
📄 Athletic/programs/program-v1.md に書き込まれたプログラム
どの AI ツールで使用できますか?
MCP サーバー (エンジン、サイエンス、データ) は標準プロトコルを通信します。
MCP — ほとんどの AI コーディング ツールが認識します。の
コーチング スキルは、Claude Code のネイティブ スキル形式にパッケージ化されています。どこかであなた
同じ命令をそのツール独自の「カスタム命令」ファイルに貼り付けます。
それぞれのセットアップ コマンド、PDF レポート ( typst が必要)、データ ディレクトリの解決、
およびトラブルシューティング: docs/installing.md 。その他のツール
MCP サーバーは、同じ uvx Performance-agent コマンドで動作することをサポートします。
コーチを使いたいだけですか？これはスキップしてください。好奇心旺盛な人や貢献者向けです。
フローチャートTB
U[あなた] <--> H[あなたのエージェント CLI<br/>クロード コード · ジェミニ CLI · コーデックス<br/>= コーチ: 会話、理由付け、適応]
H <-->|MCP| S[パフォーマンスエージェントサーバー]
H -.follows.-> SK[コーチングスキル<br/>オンボーディング · ニーズ分析 · 深い調査 ·
計画・最適化

· 栄養 · レビュー · チェックイン · 適応]
S --> E[スポーツ サイエンス エンジン<br/>決定論的 / 特性テスト済み / ゼロ LLM]
S --> EV[(証拠コーパス<br/>段階的研究、SQLite FTS5)]
S --> M[(アスリートディレクトリ<br/>プロフィール · プログラム · ログ — プレーンファイル)]
S --> R[Typst PDF レポート<br/>コーチおよびエキスパート モード · en/fr/es]
読み込み中
スキルは、専門的なコーチングのプロトコル (何を質問するか、いつ正直に言うか、どのように質問するか) をエンコードします。
ピリオダイズするか、いつデロードするか)。 MCP ツールはあらゆる事実を所有します。すでに使用しているエージェント
既存のサブスクリプションと一体化します。追加の LLM コストはゼロです。
今日の作業: 決定論的エンジン (1RM 推定、リーゲルのレース予測、
セッション RPE 負荷と ACWR、目標の実現可能性、ピリオダイゼーション ウェーブ。 597 件のテスト (テストを含む)
プロパティベース) · 47 MCP ツール · バージョン管理されたプログラムによるファイルベースのアスリート メモリ
適応監査証跡 · 反捏造を伴う DOI/PMID/ISBN 検証済み証拠コーパス
引用チェック · 生きた証拠の検索 (PubMed、OpenAlex、Crossref、Semantic Scholar)
二重検証ゲートの背後にある · 11 のコーチング スキル (コーチング スキルを含む)義務的な配達
敵対的なセカンドオピニオンを持つゲート · ハードの背後にある Typst PDF レポート (英語/フランス語/英語)
引用ゲート。
ロードマップ: 〜 200 件の研究に向けたコーパスの成長 · 結果シミュレーション (Banister +
モンテカルロ) · より多くのスポーツ分野 (Hyrox 固有のエンジン ツール、サッカー、テニス) ·
オプションの Web フロントエンド R

[切り捨てられた]

## Original Extract

Open-source AI Strength & Conditioning coach for your terminal. Runs in Claude Code / Gemini CLI / Codex — the LLM narrates, a deterministic engine calculates, and it won't lie about unrealistic goals. - clementrx/Performance-agent

GitHub - clementrx/Performance-agent: Open-source AI Strength & Conditioning coach for your terminal. Runs in Claude Code / Gemini CLI / Codex — the LLM narrates, a deterministic engine calculates, and it won't lie about unrealistic goals. · GitHub
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
clementrx
/
Performance-agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
205 Commits 205 Commits .github/ workflows .github/ workflows docs docs examples examples skills skills src/ performance_agent src/ performance_agent tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
🏋️ The first open-source AI Strength & Conditioning Coach powered by scientific research.
English · Français · Español · Deutsch · Italiano
It runs inside an AI agent CLI — a terminal program you chat with, such as Claude Code,
Gemini CLI, or Codex — and turns it into a professional coach that designs, explains,
monitors, and adapts your training. No backend, no API key, no hosting, no extra
subscription. And it will tell you the truth when your goal is unrealistic:
$ claude
> I run 10K in 55:00. I want 35:00 in 12 weeks.
🏋️ Coach: I evaluated your goal with the feasibility engine:
Probability: 0.2% — I have to be honest, this is unrealistic.
You're asking for a 36% improvement — about 3%/week sustained
for 12 weeks. Beginners sustain roughly 1%/week.
Counter-proposal: 46:30 in 12 weeks (~78% probability),
then we reassess. Want me to build that program?
Why another AI fitness coach? Because this one can't lie to you
LLM fitness coaches have two failure modes: they invent scientific references, and they
tell you what you want to hear. PerformanceAgent is architected so neither is possible:
LLMs narrate, the engine calculates. Every number — feasibility probabilities,
race predictions, training loads, periodization waves — comes from a deterministic,
property-tested Python engine. The agent explains the math; it never does the math.
Citations can't be hallucinated. The coach may only cite studies returned by the
local evidence corpus (graded, DOI/PMID-verified). The PDF renderer hard-fails on any
reference that isn't in the corpus.
Your data is files, not a cloud. Profile, programs, session logs, and check-ins
live in a plain directory of markdown/YAML you can read, edit, diff, and sync.
PerformanceAgent isn't an app you open — it plugs into an AI agent CLI. Once plugged
in, you just talk to it in plain language; no config files, no commands to memorize.
Never used Claude Code before? Install it first:
curl -fsSL https://claude.ai/install.sh | bash
(full instructions: code.claude.com/docs ).
You'll also need uv — it
fetches the right Python version by itself, nothing else to install.
Step 1 — plug in the coach. Run this once, from any terminal:
claude mcp add performance-agent -s user \
--env PERFORMANCE_AGENT_HOME= ~ /athlete-data -- uvx performance-agent
This registers the coach's "brain" (the engine, the science library, your future
athlete profile) as a tool Claude Code can call. -s user makes it available in any
folder you later open claude from. ~/athlete-data is just a suggested path — pick
any folder, it doesn't need to exist yet: the coach creates it the first time it saves
something. That's where all your data lives as plain files; nothing is sent anywhere.
Step 2 — teach it how to coach. Step 1 gave Claude the tools (the math, the
data). This step gives it the coaching protocols — when to ask what, when to be
honest about a goal, how to build a program:
git clone --depth 1 https://github.com/clementrx/Performance-agent
mkdir -p ~ /.claude/skills
cp -R Performance-agent/skills/ * ~ /.claude/skills/
Step 3 — fully quit and restart Claude Code. New tools load only when a claude
session starts : close any open session completely and run claude again.
Check it worked — in the fresh session, ask:
> List the performance-agent tools.
You should see 47 tools. If so, you're done — just talk to it.
Open a terminal and start your agent ( claude ).
Say your goal in plain language — any language works. "I want to run a 10K
under 50 minutes" , "Prépare-moi pour un Hyrox" .
Answer the coach's questions. First time, it runs a short onboarding (current
level, history, schedule, equipment) and saves your profile.
Get an honest verdict. The feasibility engine scores your goal; if it's out of
reach, you get the real probability and a counter-proposal.
Say yes — the program is written to athlete/programs/program-v1.md , periodized
into cycles with deloads and taper, every prescription carrying its purpose,
evidence grade (★★★★★ → ★☆☆☆☆), and verified citations.
After each workout, log it in one line — "Logged: 8 km easy, RPE 4" .
Every couple of weeks, say "Check-in" — the coach compares plan vs. reality,
quotes your training load, and flags drift.
When life happens — "I tweaked my ankle" , "I can only train twice next
week" — the program is adapted and versioned ( program-v2.md ) with the reason
recorded in the audit trail.
Feasibility numbers below come from the deterministic engine — run it yourself and
you'll get the same ones. Want to see the resulting program files? Each of these
three conversations ends with a full program — browse them in
examples/ , in all five languages.
> I want to run a 10K under 50 minutes. I have 5 months. Coach me.
🏋️ Coach: Two facts before I judge that goal: your current 10K time,
and how long you've been training consistently?
> I ran 55:12 last month. About a year of running, 3× a week.
🏋️ Coach: Feasibility over your 21-week horizon:
Probability: 84% — realistic, and worth committing to.
You need a 9.5% improvement, i.e. 0.45%/week, while a runner at
your training age can sustain about 1%/week. Comfortable margin.
Want me to build the program?
> Yes.
📄 Program written to athlete/programs/program-v1.md
21 weeks, 3 runs/week: one interval session, one tempo, one long
run — periodized with deloads and a 10-day taper.
2. Prepare a Hyrox in 3 months
> I signed up for a Hyrox in 3 months. Build me a plan.
🏋️ Coach: Hyrox is 8 km of running broken up by 8 strength stations —
so we train both, together. Three facts: your 10K time (or best
guess), your strength background, and days per week you can train?
> 10K around 52 min, I lift twice a week, I can do 4 sessions.
🏋️ Coach: Good base for a strong first finish. Honest note: 12 weeks
is enough to finish well and enjoy it; chasing a podium time would
need a longer runway — I'd rather tell you now.
Plan on 4 days: one running-interval day, one station-strength day
(sled work, lunges, wall balls), one "compromised running" brick
(run → station → run, the skill Hyrox actually tests), one easy
long run. Deloads weeks 4 and 8, race-week taper. Build it?
> Yes.
📄 Program written to athlete/programs/program-v1.md
3. Visible abs and bigger arms in 3 months
> I want visible abs and bigger arms in 3 months.
🏋️ Coach: Those are two different physiological goals — arms grow
with volume and enough calories; abs show at low body fat, which
usually means a deficit. Before I promise anything: height, weight,
training experience, and sessions per week?
> 1.78 m, 82 kg, two years of lifting, 4 sessions a week.
🏋️ Coach: Honest verdict: in 12 weeks you can make real progress on
both, but not maximize both — recomposition is slow at two years
of training. The realistic version:
- Moderate calorie deficit with a high protein target (the
Nutritionist skill quantifies it, with safety guards).
- Upper-body priority: 12–16 hard sets/week for arms, heavy
compound work kept in to protect strength in the deficit.
- 3 lifting days + 1 conditioning day.
- Expected at week 12: visibly fuller arms, abs "outlined" if
you start around 15–18% body fat; leaner than that takes longer,
and I'll say so at every check-in rather than let you guess.
Want the program?
> Yes.
📄 Program written to athlete/programs/program-v1.md
Which AI tool can I use this with?
The MCP server (engine, science, data) speaks a standard protocol —
MCP — that most AI coding tools understand. The
coaching skills are packaged in Claude Code's native skill format; elsewhere you
paste the same instructions into that tool's own "custom instructions" file.
Setup commands for each, PDF reports (requires typst ), data-directory resolution,
and troubleshooting: docs/installing.md . Any other tool that
supports MCP servers works with the same uvx performance-agent command.
Just here to use the coach? Skip this — it's for the curious and for contributors.
flowchart TB
U[You] <--> H[Your agent CLI<br/>Claude Code · Gemini CLI · Codex<br/>= the coach: converses, reasons, adapts]
H <-->|MCP| S[performance-agent server]
H -.follows.-> SK[Coaching skills<br/>onboarding · needs analysis · deep research ·
planning · optimization · nutrition · review · check-ins · adaptation]
S --> E[Sports science engine<br/>deterministic · property-tested · zero LLM]
S --> EV[(Evidence corpus<br/>graded studies, SQLite FTS5)]
S --> M[(Athlete directory<br/>profile · programs · logs — plain files)]
S --> R[Typst PDF reports<br/>coach & expert modes · en/fr/es]
Loading
The skills encode professional coaching protocols (what to ask, when to be honest, how
to periodize, when to deload). The MCP tools own every fact. The agent you already use
glues it together with your existing subscription — zero additional LLM cost .
Working today: deterministic engine (1RM estimation, Riegel race prediction,
session-RPE load & ACWR, goal feasibility, periodization waves; 597 tests incl.
property-based) · 47 MCP tools · file-based athlete memory with versioned programs and
an adaptation audit trail · DOI/PMID/ISBN-verified evidence corpus with anti-fabrication
citation checks · live evidence search (PubMed, OpenAlex, Crossref, Semantic Scholar)
behind a double verification gate · eleven coaching skills incl. a mandatory delivery
gate with an adversarial second opinion · Typst PDF reports (en/fr/es) behind a hard
citation gate.
Roadmap: corpus growth toward ~200 studies · outcome simulation (Banister +
Monte Carlo) · more sports verticals (Hyrox-specific engine tools, football, tennis) ·
optional web front-end r

[truncated]
