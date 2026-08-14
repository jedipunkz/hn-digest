---
source: "https://github.com/kanjani-ai-research/Vesta"
hn_url: "https://news.ycombinator.com/item?id=49304849"
title: "Vesta, adaptive ontology for Claude Code"
article_title: "GitHub - kanjani-ai-research/Vesta: An adaptive ontology over a resolved code graph: the vocabulary is learned from this repository, the rules from its own users' corrections. Ask what a change touches, where a kind of work is done, what has already been decided, and what is worth fixing. · GitHub"
author: "aug2uag"
captured_at: "2026-08-14T22:13:05Z"
capture_tool: "hn-digest"
hn_id: 49304849
score: 1
comments: 1
posted_at: "2026-08-14T21:33:11Z"
tags:
  - hacker-news
  - translated
---

# Vesta, adaptive ontology for Claude Code

- HN: [49304849](https://news.ycombinator.com/item?id=49304849)
- Source: [github.com](https://github.com/kanjani-ai-research/Vesta)
- Score: 1
- Comments: 1
- Posted: 2026-08-14T21:33:11Z

## Translation

タイトル: Vesta、クロード コードの適応オントロジー
記事のタイトル: GitHub - kanjani-ai-research/Vesta: 解決されたコード グラフ上の適応オントロジー: 語彙はこのリポジトリから学習され、ルールは独自のユーザーの修正から学習されます。変更が何に影響するのか、どこで何らかの作業が行われるのか、何がすでに決定されているのか、何を修正する価値があるのか​​を尋ねます。 · GitHub
説明: 解決されたコード グラフに対する適応オントロジー: 語彙はこのリポジトリから学習され、ルールは独自のユーザーの修正から学習されます。変更が何に影響するのか、どこで何らかの作業が行われるのか、何がすでに決定されているのか、何を修正する価値があるのか​​を尋ねます。 - 関ジャニ愛リサーチ/Vesta

記事本文:
GitHub - kanjani-ai-research/Vesta: 解決されたコード グラフ上の適応オントロジー: 語彙はこのリポジトリから学習され、ルールは独自のユーザーの修正から学習されます。変更が何に影響するのか、どこで何らかの作業が行われるのか、何がすでに決定されているのか、何を修正する価値があるのか​​を尋ねます。 · GitHub
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
関ジャニ愛リサーチ
/
ベスタ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
本支店T

ags ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
148 コミット 148 コミット .claude-plugin .claude-plugin .claude .claude エージェント エージェント bin bin コマンド コマンド doc doc フック フック スキル/ ベスタ スキル/ ベスタ テスト テスト トライアル トライアル ベスタ ベスタ .claudeignore .claudeignore .gitignore .gitignore ライセンス ライセンス通知.md NOTICE.md README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
/plugin マーケットプレイスに https://gitlab.com/causum/vesta.git を追加します
/プラグインインストールvesta@causum
リポジトリに関する構造的な質問に、解決されたグラフから回答します。
何を指すのか、その作品が何と呼ばれているかのオントロジー、そして以前の何を指すのか
セッションはすでに完了しているため、エージェントは読む代わりに質問することができ、あなたは
同じ理解のために二度お金を払わないでください。
インストールして正常に動作します。実行するものは何もありません。
ベスタはデフォルトでコンパニオンになります。リポジトリのグラフを構築します。
バックグラウンドで、コードが変更されても最新の状態に保ち、既知のものをバックグラウンドに配置します。
重要な瞬間にエージェントの前で:
変更しようとしているファイル内ですでに間違っているものは何か。ファイルに名前を付ける
飲み込まれた失敗を抱えて、一度そう言います。もう一度名前を付けても何も言えません。
この作品は破られるだろうというあなたが設定したルール。ルールはあなたとして記録されます
通常の作業の中でそれらを述べます - 何も実行する必要はありません。
覚えておいてください。
以前のセッションでプロンプト名の定義について解決されたもの、
結論が導き出された地域と合わせて、エージェントが確認できるようにする
彼らを信頼するのではなく。
何も名前を付けないプロンプトには何も表示されませんが、これがほとんどです。
すべてのツールはスラッシュ コマンドです。知っておく価値のあるもの:
/vesta:ファイルを開く前に、このリポジトリが何で構成されているかを確認します
/vesta:変更の範囲と、それをカバーするテストに触れる
/ベスタ:ドー

ある種の仕事が行われる場所、普通の言葉で尋ねる
/vesta:頼まれてもいないのに見つかった、修正する価値のある欠陥
/vesta:あなたが述べた決定されたルール、およびコードがそれらを尊重するかどうか
/vesta:tutorial 自分のリポジトリで一度に 1 ページずつ学習します
/vesta:help には残りがリストされます。
dos は grep ではないものです。作品の語彙を尋ねる —
「失敗したリクエストの再試行」、「送信の重複排除」 — そして、次のような答えが返されます。
コードの語彙は通常異なります。その交差点には
エージェントがリポジトリごとに 1 回抽出する語彙。
プロジェクト全体を空のディレクトリにビルドするように要求すると、次の選択肢が提供されます。
対話的に構築するか、最初に契約に同意して完了まで実行します。選択してください
インタラクティブなので、再度質問されることはありません。
自動モードはユーザーに代わって開始されることはなく、また、自動モードが提供されることもありません。
誰かがすでに作業しているリポジトリ。
Vesta はエージェント独自の推論に基づいて実行されます。 API キーを保持しておらず、
独自のネットワーク呼び出し。
どのモデルが何を行うかは好みではありません: 定義の読み取りとラベル付け
すべての定義に対して 1 回発生するため、小規模なモデルで実行されます。
その体積でより大きなものを使用すると、このアプローチは使用するにはコストが高くなりすぎます。合成
契約や仕様書は、より大きなものに基づいて実行されます。
一度。
派生したものはすべて ~/.vesta の下に保存され、削除できます。 /vesta:開催中
何が保持されているかを報告し、なくなったリポジトリに属するものを再利用します。
マシンからは何も残りません。
隠されたものは読み取られません。ドットで始まるものは決してありません
歩きました — .env 、 .aws 、 .ssh ではありません。依存関係ディレクトリも同様です。
venv 、node_modules 、site-packages 、vendor など。
導き出していないものは主張しません。ルールが記録されます
あなたが実際に言った言葉、

トランスクリプトと照合して検証されます。テンプレートでできること
語彙は提供しますが、あなたの定義のどれがそれを意味するかについては決して主張しません。
仕事。
完全な問題を示唆するのではなく、解決できなかった内容を報告します。
答えてください。伝播セットは、どの参照をたどることができなかったかを示します。グラフ
ワークスペースのプロジェクト間の参照がワークスペース内にないことを示しています。
Python 3.10 以降、および必要な言語の言語サーバー
解決済み — Python の場合は pyright-langserver、Rust の場合は Rust-analyzer、gopls
囲碁など。 Vesta は、最初の使用時に独自のランタイムを構築します。それは尋ねません
自分のものに何でもインストールできます。
サーバーがインストールされていない言語のリポジトリは、むしろそのように報告されます。
静かに半分解決するよりも。
測定されたものと測定されていないもの
このリポジトリに対する読み取り専用の構造的な質問について、
同じプロンプトの実行を制御し、ホスト自身の /cost によってコストが計算されます。
インジェクションは、エージェントが何かを決定する前に 799 文字を配信しました。道具
3 回の通話で 21,905 件を配信しました。残りは doc/measurements.md にあり、
トライアル/ハーネスが付いています。
これによる影響のほとんどは測定されていません。欠陥が表面化しているかどうか
誰かがファイルを編集した瞬間、語彙の有無にかかわらず、作成したものが改善されます。
抽出によって価値のあるルールが回復されるかどうかにかかわらず、crossing は grep では見つけられないものを見つけます。
持っている — そのどれも示されていない。 doc/open-questions.md には何が記録されますか
true である必要があり、それをチェックしたときに何が見つかったのか。
アパッチ2.0。ライセンスおよび通知.md を参照してください。
解決されたコードグラフ上の適応オントロジー: 語彙はこのリポジトリから学習され、ルールは独自のユーザーの修正から学習されます。変更が何に影響するのか、どこで何らかの作業が行われるのか、何がすでに決定されているのか、何を修正する価値があるのか​​を尋ねます。
Readme Apache-2.0 ライセンス アクティビティ カスタム プロパティ スター
0フォーク

s レポートリポジトリのリリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An adaptive ontology over a resolved code graph: the vocabulary is learned from this repository, the rules from its own users' corrections. Ask what a change touches, where a kind of work is done, what has already been decided, and what is worth fixing. - kanjani-ai-research/Vesta

GitHub - kanjani-ai-research/Vesta: An adaptive ontology over a resolved code graph: the vocabulary is learned from this repository, the rules from its own users' corrections. Ask what a change touches, where a kind of work is done, what has already been decided, and what is worth fixing. · GitHub
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
kanjani-ai-research
/
Vesta
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
148 Commits 148 Commits .claude-plugin .claude-plugin .claude .claude agents agents bin bin commands commands doc doc hooks hooks skills/ vesta skills/ vesta tests tests trial trial vesta vesta .claudeignore .claudeignore .gitignore .gitignore LICENSE LICENSE NOTICE.md NOTICE.md README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
/plugin marketplace add https://gitlab.com/causum/vesta.git
/plugin install vesta@causum
Answers structural questions about a repository from a resolved graph of what
refers to what, an ontology of what the work is called, and what earlier
sessions already worked out — so an agent can ask instead of reading, and you
do not pay twice for the same understanding.
Install it and work normally. There is nothing to run.
Vesta is a companion by default. It builds a graph of the repository in the
background, keeps it current as the code changes, and puts what it knows in
front of the agent at the moment it matters:
What is already wrong in the file you are about to change. Naming a file
with a swallowed failure in it says so, once. Naming it again says nothing.
A rule you set that this work would break. Rules are recorded as you
state them in the course of ordinary work — nothing to run, nothing to
remember.
What earlier sessions worked out about a definition the prompt names,
with the regions those conclusions were drawn from, so an agent can check
them rather than take them on trust.
It says nothing on a prompt that names nothing, which is most of them.
Every tool is a slash command. The ones worth knowing:
/vesta:shape what this repository is made of, before opening a file
/vesta:touches what a change reaches, and which tests cover it
/vesta:does where a kind of work happens, asked in ordinary words
/vesta:defects things worth fixing, found without being asked
/vesta:decided rules you have stated, and whether the code honours them
/vesta:tutorial learn it a page at a time, on your own repository
/vesta:help lists the rest.
does is the one that is not a grep. Ask in the vocabulary of the work —
"retrying a failed request", "deduplicating submissions" — and it answers in
the vocabulary of the code, which is usually different. That crossing needs a
vocabulary, which an agent derives once per repository.
Asking for a whole project to be built in an empty directory offers a choice:
build it interactively, or agree a contract first and run to completion. Choose
interactive and you are not asked again.
Automated mode is never entered on your behalf, and never offered in a
repository somebody has already been working in.
Vesta runs on your agent's own inference. It holds no API key and makes no
network calls of its own.
Which model does what is not a preference: reading a definition and labelling
it runs on a small model, because it happens once for every definition and a
larger one at that volume makes the approach too expensive to use. Synthesis
somebody will be held to — a contract, a specification — runs on a larger one,
once.
Everything derived is kept under ~/.vesta and can be deleted. /vesta:held
reports what is held and reclaims what belongs to repositories that are gone.
Nothing leaves your machine.
It does not read anything hidden. Nothing beginning with a dot is ever
walked — not .env , not .aws , not .ssh . Nor are dependency directories:
venv , node_modules , site-packages , vendor , and the rest.
It does not assert what it has not derived. A rule is recorded against
the words you actually said, verified against the transcript; a template can
lend a vocabulary but never a claim about which of your definitions do the
work.
It reports what it could not resolve rather than implying a complete
answer. A propagation set says which references it could not follow; a graph
of a workspace says that references between its projects are not in it.
Python 3.10 or newer, and a language server for the languages you want
resolved — pyright-langserver for Python, rust-analyzer for Rust, gopls
for Go, and so on. Vesta builds its own runtime on first use; it does not ask
you to install anything into yours.
A repository whose language has no server installed is reported as such rather
than silently half-resolved.
What has and has not been measured
On a read-only structural question against this repository, paired with a
control run of the same prompt and costed by the host's own /cost :
Injection delivered 799 characters before the agent decided anything; the tools
delivered 21,905 across three calls. doc/measurements.md has the rest, and
trial/ has the harness.
Most of what this does is unmeasured. Whether surfacing a defect at the
moment somebody edits a file improves what they produce, whether the vocabulary
crossing finds things a grep would not, whether extraction recovers rules worth
having — none of that has been shown. doc/open-questions.md records what
would have to be true, and what was found when it was checked.
Apache 2.0. See LICENSE and NOTICE.md .
An adaptive ontology over a resolved code graph: the vocabulary is learned from this repository, the rules from its own users' corrections. Ask what a change touches, where a kind of work is done, what has already been decided, and what is worth fixing.
Readme Apache-2.0 license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
