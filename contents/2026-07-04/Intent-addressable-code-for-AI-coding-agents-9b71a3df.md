---
source: "https://github.com/croviatrust/causari"
hn_url: "https://news.ycombinator.com/item?id=48784012"
title: "Intent-addressable code for AI coding agents"
article_title: "GitHub - croviatrust/causari: Intent-addressable code for AI agents. Causari records every action an agent takes on your codebase - the prompt, model, reads, writes and reasoning - and lets you trace, diff, revert and bisect them like git commits. With bidirectional causal provenance and a built-in\n[truncated]"
author: "CroviaTrust"
captured_at: "2026-07-04T10:00:19Z"
capture_tool: "hn-digest"
hn_id: 48784012
score: 1
comments: 0
posted_at: "2026-07-04T09:19:00Z"
tags:
  - hacker-news
  - translated
---

# Intent-addressable code for AI coding agents

- HN: [48784012](https://news.ycombinator.com/item?id=48784012)
- Source: [github.com](https://github.com/croviatrust/causari)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T09:19:00Z

## Translation

タイトル: AI コーディング エージェント用のインテント アドレス指定可能なコード
記事のタイトル: GitHub - croviatrust/causari: AI エージェント用のインテント アドレス指定可能なコード。 Causari は、エージェントがコードベースに対して実行するすべてのアクション (プロンプト、モデル、読み取り、書き込み、推論) を記録し、git コミットのようにそれらをトレース、差分、元に戻し、二分することができます。双方向の因果関係と組み込み
[切り捨てられた]
説明: AI エージェント用のインテントアドレス指定可能なコード。 Causari は、エージェントがコードベースに対して実行するすべてのアクション (プロンプト、モデル、読み取り、書き込み、推論) を記録し、git コミットのようにそれらをトレース、差分、元に戻し、二分することができます。双方向の因果関係と内蔵 MCP サーバーを備えています。 - クロビアトラスト/原因
[切り捨てられた]

記事本文:
GitHub - croviatrust/causari: AI エージェント用のインテント アドレス指定可能なコード。 Causari は、エージェントがコードベースに対して実行するすべてのアクション (プロンプト、モデル、読み取り、書き込み、推論) を記録し、git コミットのようにそれらをトレース、差分、元に戻し、二分することができます。双方向の因果関係と内蔵 MCP サーバーを備えています。 · GitHub
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
アカウントを切り替えました

他のタブまたはウィンドウ。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
クロビアトラスト
/
原因
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
39 コミット 39 コミット .github .github アセット アセット スクリプト スクリプト サイト サイト src src テスト/ベクトル テスト/ベクトル .gitattributes .gitattributes .gitignore .gitignore CLA.md CLA.md COTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス通知 通知 README.md README.md wrangler.jsonc wrangler.jsonc すべてのファイルを表示 リポジトリ ファイルのナビゲーション
意図を追跡します。デバッグの因果関係。
AI エージェント用のインテントでアドレス指定可能なコード。
causari.dev
·
リリース
·
ディスカッション
·
MCP
·
ライセンス (BSL 1.1)
Causari (ラテン語、非難動詞): 大義を主張し、その理由を主張すること。なぜなら
AI が生成したコードのすべての行は、防御され、追跡され、
分かりました。
Causari は、AI エージェントがコードベースに対して実行するすべてのアクションを記録します。
変更されたバイト、ただし尋ねられたプロンプト、モデル
回答、読み取られたファイル、および変更の背後にある理由。
そして、エージェントの許可を求めずにこれを実行します。
キャプチャ エンジン (再プロキシ + 再ウォッチ + 再フック) は LLM を監視します
トラフィックとファイルシステムを独立して、コンテンツごとに結合します。
ファイル内に表示されるコードは、補完内にあります。
数秒前に作成しました。出自は自己報告ではなく事実になります。
次に、バージョン管理システムがこれまでに答えたことのない質問をすることができます。
re proxy # ローカル LLM プロキシ: すべてのプロンプト、トークン、ドル
# プロバイダーに向かう途中で Causari を通過します
re watch # パッシブレコーダー + 因果結合: ファイル変更の取得
# 実際のプロンプトに起因する

、モデルと価格
rehook claude-code # エージェントのライフサイクルフックを介したネイティブキャプチャ
なぜ src/auth.ts:42 # この正確な行を作成したのは誰ですか?
re トレース src/auth.ts:42 # 完全な UPSTREAM コーザル コーン: すべてのイベント
# 読み取り/書き込みを通じて推移的に貢献
re Impact <event-id> # フル DOWNSTREAM コーン: このアクションから流れたもの、
# 推移的 (因果関係を考慮した爆発範囲)
re Lens src/auth.ts # 行ごとの出所注釈を付けてファイルをレンダリングする
re find " the JWT refactor " # すべてのプロンプト、推論、メッセージを検索します
re bisect --test " npm test " # ビルドを中断したエージェント アクションを検索します
re churn # AI コードの生存を測定する: どれだけ生き残ったか
# エージェントごとに書き換えられ、無駄な費用が発生しました
re report --open # AI 廃棄物の共有可能な HTML ダッシュボードを生成します
re skill distill # 検証されたイベントを署名された再利用可能なスキルに変換する
re スキル エクスポート < id > # チームメイト用のポータブル Ed25519 バンドル
re skill pull <チームディレクトリ> # 共有フォルダーを同期する (Dropbox、git、NFS — サーバーなし)
re skill trust add < label > # 組織署名キーを信頼します。不明な署名者は拒否されました
re fork Experiment-claude # 並列タイムラインに分岐します
re revert < id > # 他のアクションの因果関係のプレビューを使用してアクションを元に戻します
# 暗黙的に元に戻すことになります
エージェントが 30 個のファイルに触れて何かが壊れた場合、読む必要はありません
4000行のチャット。あなたはカウザリに、なぜ、そしていつなのか尋ねます。
キャプチャ エンジン — 協力なしの来歴
Causari 以前のすべての来歴ツールには同じ致命的な依存関係がありました。
エージェントが自らの歴史を自発的に提供した場合に機能しました。エージェントはそうではありません。ハーネス
理屈を暴露しないでください。誰もコストを報告しません。
Causari は依存関係を削除します。 2 つの独立した観測ストリーム、1 つ
因果結合:
┌─────────────┐ ┌───

─────────┐
│ リプロキシ │ │ リウォッチ │
│ │ │ │
│ はすべてのプロンプトを確認し、 │ はすべてのバイトを確認します │
│ 完了、トークン、および │ │ ディスク上の変更 │
│ ドル (OpenAI- および │ │ (スナップショット、差分) │
│ 人間互換） │ │ │
━───────┬───────┘ └───────┬───────┘
│ │
│ コンテンツベースの結合 │
━───────►◄───────┘
ファイルに挿入された行
補完内で検索されます
直前にキャプチャされた瞬間 - 試合は
自信を持って因果関係のある指紋を取得する
実際のセッション、エンドツーエンド:
$リプロキシ
原因: http://127.0.0.1:4242 をリッスンしている LLM キャプチャ プロキシ
• gpt-4o 42→18 tok $0.0003 「24 時間ごとにローテーションする JWT リフレッシュ ロジックを追加」
$ 別の端末で # 再視聴
• 0d47599550 認証.py
↳ インテント: 「24 時間ごとにローテーションする JWT リフレッシュ ロジックを追加する」 gpt-4o (信頼性 100%、3/3 行)
$ なぜ auth.py:2 なのか
認証.py:2
トークン = issue_token(ユーザー、スコープ="セッション")
0d47599550 によって紹介されました
エージェント: カーソル
モデル: gpt-4o
プロンプト: 24 時間ごとにローテーションする JWT 更新ロジックを追加します
エージェントにプロキシを指定すれば完了です。
OPENAI_BASE_URL=http://127.0.0.1:4242/openai/v1
ANTHROPIC_BASE_URL=http://127.0.0.1:4242/anthropic
エージェント ランタイムがライフサイクル フックを公開する場合、キャプチャはネイティブであり、
正確 — 推論は必要ありません:
claude-code を再フック # UserPromptSubmit + PostToolUse を接続します
# .claude/settings.json: キャプチャされたすべてのプロンプト、
# すべての編集は完全な Causari イベントとして記録されます
すべてがマシン上に残ります。.causari/capture/ はローカルです。
追加専用の台帳。クラウド、テレメトリ、API キーのタッチなし

編
Crovia Seals — すべての完了に対する暗号化された領収書
Causari は、
Crovia シール — オープン、
IETF が草案した AI 出力のレシート形式
( ドラフト-crovia-seal-01 )。
1 つのフラグがプロキシをシーリング ゲートウェイに変えます。
$ リプロキシ --seal
原因: Crovia シール発行者がアクティブ — 公開鍵 3fa9c2…
• gpt-4o 42→18 tok $0.0003 「JWT 更新ロジックの追加」 🔏 cs_2026_Q7RM2KJ3VWXA5YBN4CDEFGH2I6
すべての完了には、Ed25519 で署名され、ハッシュチェーンされ、オフラインで検証可能なデータが取得されます。
.causari/seal/seals.jsonl の領収書。シールは SHA-256 ハッシュにコミットします
正確なリクエストバイトとレスポンスバイトのデータ — コンテンツがユーザーの外に出ることはありません
機械。あなたの公開鍵を持っている人は誰でもチェーン全体を検証できます
サーバー、アカウント、または Causari 自体を使用しない場合:
$再シール確認
✓ 128 個の印鑑が検証済み - すべての署名が有効で、起源から連続したチェーン
$ シール発行者 # 監査人と共有するために公開鍵を印刷します
$ 再封印リスト # 発行された領収書を閲覧する
実装は規範的な適合ベクトルに対して証明されています
croviatrust/crovia-seal より
(CSC-1 正規化、ドメイン分離ペイロード、フェールクローズ
検証）。規制当局、顧客、または裁判所が「どのモデルか」と尋ねたとき、
このコードを書いたのですが、それを証明できますか?」 — 答えは 1 つのファイルと 1 つです
公開鍵。
エクスペリエンス層 — 信頼を獲得したスキル
過去を記録することが仕事の半分です。残りの半分はエージェントがいないことを確認しています
同じレッスン料を二度支払うことはありません。
再スキル蒸留は台帳を調べ、完了したすべてのタスクを圧縮します。
それをトリガーしたプロンプト、実行された手順、実行されたファイル
変更 — スキル : エージェントが思い出すことができる経験の単位
演技する前に。各スキルはリポジトリの Ed25519 キーで署名されています
蒸留の瞬間。その後 1 バイト編集して、
再スキル検証

それを暴露します。
信頼は獲得するものであり、決して主張するものではありません。
● 台帳から抽出された記録 — 成功のシグナルはまだありません
◆ 検証済みの証拠が添付されています: 終了コード 0、または作業はまだです
タイムラインの先端で生きている（生き残った）
★ 新しい作業を行うエージェントによって 3 回以上検証され、リコールされたことが証明されています
$再スキル蒸留
蒸留: 128 個のイベントがスキャンされ、7 個の新しいスキル、12 個がすでに蒸留されています
◆ 2ce0c7bbda が指数バックオフを使用した再試行を追加することを確認しました
$再スキル検証
ok 2ce0c7bbda 指数バックオフを使用した再試行を追加します
検証: 7 つのスキル、すべての署名が有効
ループは MCP を通じて終了します。エージェントが署名された causari_recall を呼び出すと、
スキルが最初に返され、信頼度に基づいてランク付けされます (実証済み×4、検証済み×2)。
思い出すたびにスキルの使用カウンターが増加します。これがまさにその方法です。
検証されたスキルは★を獲得します。エージェントは時間の経過とともに目に見えて安くなり、
再解約では、節約額がドルで表示されます。
チーム スキル メッシュ — サーバーもアカウントも必要ありません
1 人のエンジニアが検証した修正は、すべてのエージェントの直感になります。
中心的なSaaS:
re スキルのエクスポート 2ce0c7bbda --output jwt-fix.json # ポータブル バンドル
スキル信頼公開鍵 # Ed25519 キーを共有します
re skill trust add platform < their-pubkey > # チームメイト/組織キーを信頼する
re skill import jwt-fix.json # 署名を検証 + 受け入れる
re skill pull ~ /Dropbox/causari-skills/ # チーム フォルダー全体を同期する
不明なキーで署名されたスキルは拒否され、インポートされません。メッシュは
暗号化: Dropbox、git、NFS、S3 — どのフォルダーでも機能します。カウザリ氏は検証する
すべてのファイルに Ed25519。改ざんされたバンドルはフェールクローズされます。
Causari のすべてと同様、スキルはローカル ファイル ( .causari/skills/ ) です。
自己完結型でポータブル。署名はスキルを共有できることを意味し、
中央サーバーを必要とせず、リポジトリ、チーム、組織全体で誰でも検証できます。
Causari Proof — 検証可能な AI の出所、トラッスル

s
すべてのリポジトリは、AI の出所を証明する署名済みの証明書を作成できます (エージェントの数)。
アクション、どのエージェントとモデル、どの程度の検証済みエクスペリエンスか — にバインドされています。
コンテンツダイジェストによる正確な台帳であり、リポジトリの Ed25519 キーで署名されています。
再証明生成 # → causari-proof.json + causari-proof.svg バッジ
reproof verify # 署名をオフラインでチェックします — サーバーもアカウントもありません
reproof verify --against-repo # …そしてライブ台帳とまだ一致していることを確認します
査読者、監査人、自己PRを読んでいる見知らぬ人など、誰でも実行できます。
証明を検証し、署名後に証明が変更されていないことを確認します。いいえ
Causari アカウント、ネットワーク通話なし、私たちへの信頼なし。単一の番号を改ざんする
そして検証は失敗して終了します。
README にバッジをドロップすると、すべての訪問者に表示されます。
[ ![ AI の出所 — Causari によって検証されました ] ( causari-proof.svg )] ( https://causari.dev/verify )
構造上エージェントに依存しません。証明は元帳を集約するため、
カウザリが捕らえたすべてのエージェントを網羅 - クロード・コード、カーソル、クライン、ウィンドサーフィン、
raw リプロキシ — 単なる 1 つのランタイムではありません。
永久無料: プルーフをオフラインで生成および検証します。コマーシャル（信託）
Plane): causari.dev でホストされている公開検証ページ、組織全体
証明レジストリ、RFC 3161 タイムスタンプアンカーリング、および監査グレードのコンプライアンス
輸出。
既存のツールは、テキスト (git) を追跡するか、セッション (IDE チェックポイント) を追跡するか、
会話を追跡します (LangSmith、Helicone)。それらのどれも線を接続しません
コードを生成したインテントに合わせてコードを作成します。
エージェントの協力。カウザリは次の両方を行います。
双方向因果グラフ
ほとんどのバージョン管理は 1 次元、つまり一連のコミットです。コーザー

[切り捨てられた]

## Original Extract

Intent-addressable code for AI agents. Causari records every action an agent takes on your codebase - the prompt, model, reads, writes and reasoning - and lets you trace, diff, revert and bisect them like git commits. With bidirectional causal provenance and a built-in MCP server. - croviatrust/caus
[truncated]

GitHub - croviatrust/causari: Intent-addressable code for AI agents. Causari records every action an agent takes on your codebase - the prompt, model, reads, writes and reasoning - and lets you trace, diff, revert and bisect them like git commits. With bidirectional causal provenance and a built-in MCP server. · GitHub
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
croviatrust
/
causari
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
39 Commits 39 Commits .github .github assets assets scripts scripts site site src src tests/ vectors tests/ vectors .gitattributes .gitattributes .gitignore .gitignore CLA.md CLA.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE NOTICE NOTICE README.md README.md wrangler.jsonc wrangler.jsonc View all files Repository files navigation
Trace intent. Debug causality.
Intent-addressable code for AI agents.
causari.dev
·
Releases
·
Discussions
·
MCP
·
License (BSL 1.1)
Causari (Latin, deponent verb): to plead a cause, to argue why. Because
every line of AI-generated code deserves to be defended, traced, and
understood.
Causari records every action an AI agent takes on your codebase — not just
the bytes that changed, but the prompt that asked , the model that
answered , the files it read , and the reasoning behind the change .
And it does so without asking the agent's permission : the built-in
capture engine ( re proxy + re watch + re hook ) observes the LLM
traffic and the filesystem independently, then joins them by content —
the code that appears in your files is found inside the completion that
produced it seconds earlier. Provenance becomes a fact, not a self-report.
You can then ask questions no version control system has ever answered:
re proxy # local LLM proxy: every prompt, token and dollar
# flows through Causari on its way to the provider
re watch # passive recorder + causal join: file changes get
# attributed to the real prompt, model and cost
re hook claude-code # native capture via agent lifecycle hooks
re why src/auth.ts:42 # who/what produced this exact line?
re trace src/auth.ts:42 # full UPSTREAM causal cone: every event that
# contributed transitively, through reads/writes
re impact < event-id > # full DOWNSTREAM cone: what flowed from this action,
# transitively (causality-aware blast radius)
re lens src/auth.ts # render a file with per-line provenance annotations
re find " the JWT refactor " # search every prompt, reasoning and message
re bisect --test " npm test " # find the agent action that broke the build
re churn # measure AI code survival: how much survived vs
# was rewritten, per agent, with wasted spend
re report --open # generate a shareable HTML dashboard of AI waste
re skill distill # turn verified events into signed, reusable skills
re skill export < id > # portable Ed25519 bundle for teammates
re skill pull < team-dir > # sync a shared folder (Dropbox, git, NFS — no server)
re skill trust add < label > # trust an org signing key; unknown signers rejected
re fork experiment-claude # branch into a parallel timeline
re revert < id > # undo an action with causal preview of what else
# you are implicitly undoing
When an agent touches 30 files and something breaks, you don't need to read
4 000 lines of chat. You ask Causari why and when .
The Capture Engine — provenance without cooperation
Every provenance tool before Causari had the same fatal dependency: it only
worked if the agent volunteered its own history. Agents don't. Harnesses
don't expose reasoning. Nobody reports costs.
Causari removes the dependency. Two independent observation streams, one
causal join:
┌─────────────────────────┐ ┌─────────────────────────┐
│ re proxy │ │ re watch │
│ │ │ │
│ sees every prompt, │ │ sees every byte that │
│ completion, token and │ │ changes on disk │
│ dollar (OpenAI- and │ │ (snapshots, diffs) │
│ Anthropic-compatible) │ │ │
└────────────┬────────────┘ └────────────┬────────────┘
│ │
│ CONTENT-BASED JOIN │
└────────────────►◄────────────────┘
the lines inserted in your files
are searched inside the completions
captured moments before — a match is
a causal fingerprint, with confidence
A real session, end to end:
$ re proxy
causari: LLM capture proxy listening on http://127.0.0.1:4242
• gpt-4o 42→18 tok $0.0003 "Add JWT refresh logic that rotates every 24h"
$ re watch # in another terminal
• 0d47599550 auth.py
↳ intent: "Add JWT refresh logic that rotates every 24h" gpt-4o (confidence 100%, 3/3 lines)
$ re why auth.py:2
auth.py:2
token = issue_token(user, scope="session")
introduced by 0d47599550
agent: cursor
model: gpt-4o
prompt: Add JWT refresh logic that rotates every 24h
Point any agent at the proxy and you're done:
OPENAI_BASE_URL=http://127.0.0.1:4242/openai/v1
ANTHROPIC_BASE_URL=http://127.0.0.1:4242/anthropic
Where the agent runtime exposes lifecycle hooks, capture is native and
exact — no inference needed:
re hook claude-code # wires UserPromptSubmit + PostToolUse into
# .claude/settings.json: every prompt captured,
# every edit recorded as a full Causari event
Everything stays on your machine: .causari/capture/ is a local,
append-only ledger. No cloud, no telemetry, no API keys touched.
Crovia Seals — a cryptographic receipt for every completion
Causari is the first production issuer of
Crovia Seals — the open,
IETF-drafted receipt format for AI outputs
( draft-crovia-seal-01 ).
One flag turns the proxy into a sealing gateway:
$ re proxy --seal
causari: Crovia Seal issuer active — pubkey 3fa9c2…
• gpt-4o 42→18 tok $0.0003 "Add JWT refresh logic" 🔏 cs_2026_Q7RM2KJ3VWXA5YBN4CDEFGH2I6
Every completion gets an Ed25519-signed, hash-chained, offline-verifiable
receipt in .causari/seal/seals.jsonl . The seal commits to SHA-256 hashes
of the exact request and response bytes — content never leaves your
machine . Anyone holding your public key can verify the whole chain
without a server, an account, or Causari itself:
$ re seal verify
✓ 128 seal(s) verified — every signature valid, chain contiguous from genesis
$ re seal issuer # print the pubkey to share with auditors
$ re seal list # browse issued receipts
The implementation is proven against the normative conformance vectors
from croviatrust/crovia-seal
(CSC-1 canonicalization, domain-separated payloads, fail-closed
verification). When a regulator, a customer or a court asks "which model
wrote this code, and can you prove it?" — the answer is one file and one
public key.
The Experience Layer — skills with earned trust
Recording the past is half the job. The other half is making sure no agent
ever pays for the same lesson twice.
re skill distill walks the ledger and compresses every completed task —
the prompt that triggered it, the steps that were taken, the files that
changed — into a skill : a unit of experience an agent can recall
before acting. Each skill is signed with the repository's Ed25519 key
at the moment of distillation; edit one byte afterwards and
re skill verify exposes it.
Trust is earned, never claimed:
● recorded distilled from the ledger — no success signal yet
◆ verified evidence attached: exit code 0, or the work is still
alive at the tip of the timeline (it survived)
★ proven verified AND recalled 3+ times by agents doing new work
$ re skill distill
distill: 128 event(s) scanned, 7 new skill(s), 12 already distilled
◆ verified 2ce0c7bbda add retry with exponential backoff
$ re skill verify
ok 2ce0c7bbda add retry with exponential backoff
verify: 7 skill(s), every signature valid
The loop closes through MCP: when an agent calls causari_recall , signed
skills are returned first, ranked by trust (proven ×4, verified ×2), and
every recall bumps the skill's use counter — which is exactly how a
verified skill earns the ★. Agents get measurably cheaper over time, and
re churn shows you the savings in dollars.
Team skill mesh — no server, no accounts
One engineer's verified fix becomes every agent's instinct — without a
central SaaS:
re skill export 2ce0c7bbda --output jwt-fix.json # portable bundle
re skill trust pubkey # share your Ed25519 key
re skill trust add platform < their-pubkey > # trust a teammate/org key
re skill import jwt-fix.json # verify signature + accept
re skill pull ~ /Dropbox/causari-skills/ # sync a whole team folder
Skills signed by unknown keys are rejected , not imported. The mesh is
cryptographic: Dropbox, git, NFS, S3 — any folder works. Causari verifies
Ed25519 on every file; tampered bundles fail closed.
Like everything in Causari, skills are local files ( .causari/skills/ ),
self-contained and portable. The signature means a skill can be shared and
verified by anyone — across repos, teams, and orgs, with no central server.
Causari Proof — verifiable AI provenance, trustless
Every repo can mint a signed proof of its AI provenance — how many agent
actions, which agents and models, how much verified experience — bound to the
exact ledger by a content digest and signed with the repo's Ed25519 key.
re proof generate # → causari-proof.json + causari-proof.svg badge
re proof verify # checks the signature offline — no server, no account
re proof verify --against-repo # …and confirms it still matches the live ledger
Anyone — a reviewer, an auditor, a stranger reading your PR — can run
re proof verify and confirm the proof was not altered after signing . No
Causari account, no network call, no trust in us. Tamper with a single number
and verification fails closed.
Drop the badge in your README and every visitor sees it:
[ ![ AI provenance — verified by Causari ] ( causari-proof.svg )] ( https://causari.dev/verify )
It is agent-agnostic by construction: the proof aggregates the ledger , so it
covers every agent Causari captured — Claude Code, Cursor, Cline, Windsurf, a
raw re proxy — not just one runtime.
Free forever: generating and verifying proofs offline. Commercial (Trust
Plane): the hosted public verification page on causari.dev , the org-wide
proof registry, RFC 3161 timestamp anchoring, and audit-grade compliance
exports.
Existing tools either track text (git), track sessions (IDE checkpoints), or
track conversations (LangSmith, Helicone). None of them connect a line of
code to the intent that produced it — and none of them can do it without
the agent's cooperation. Causari does both:
The bidirectional causal graph
Most version control is one-dimensional: a chain of commits. Causar

[truncated]
