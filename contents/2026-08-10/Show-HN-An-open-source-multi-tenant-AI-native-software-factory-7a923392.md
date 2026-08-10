---
source: "https://github.com/missingstudio/eva"
hn_url: "https://news.ycombinator.com/item?id=49241928"
title: "Show HN: An open-source multi-tenant, AI-native software factory"
article_title: "GitHub - missingstudio/eva: An autonomous, multi-tenant, AI-native software factory. · GitHub"
author: "missingstack"
captured_at: "2026-08-10T10:58:08Z"
capture_tool: "hn-digest"
hn_id: 49241928
score: 1
comments: 1
posted_at: "2026-08-10T10:40:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: An open-source multi-tenant, AI-native software factory

- HN: [49241928](https://news.ycombinator.com/item?id=49241928)
- Source: [github.com](https://github.com/missingstudio/eva)
- Score: 1
- Comments: 1
- Posted: 2026-08-10T10:40:40Z

## Translation

タイトル: Show HN: オープンソースのマルチテナント、AI ネイティブのソフトウェア ファクトリ
記事のタイトル: GitHub - missingstudio/eva: 自律型、マルチテナント、AI ネイティブのソフトウェア ファクトリ。 · GitHub
説明: 自律型、マルチテナント型、AI ネイティブのソフトウェア ファクトリ。 - missingstudio/eva

記事本文:
GitHub - missingstudio/eva: 自律型、マルチテナント型、AI ネイティブのソフトウェア ファクトリ。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
行方不明のスタジオ
/
エヴァ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
91 コミット 91 コミット .github/ workflows .github/ workflows cmd/ eva cmd/ eva docs docs 内部 内部 .gitignore .gitignore .golangci.yml .g

olangci.yml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTEXT.md CONTEXT.md LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイル ナビゲーション
自律型、マルチテナント型、AI ネイティブのソフトウェア ファクトリ
ほとんどの AI ツールは、そのツールが何をしたかを教えてくれます。エヴァが最初にそれを書き留めます。
あなたが尋ねたすべての質問、返されたすべての回答、すべての再試行、およびカウントされたすべてのトークンは、発生時にディスク上のファイルに追加されます。画面に表示されている内容は、そのファイルから読み取られたものです。 Eva が答えを示した場合、ファイルにはすでに答えが含まれています。イベントの 2 番目のバージョンはありません。
それは小さなことのように聞こえます。全体のデザインです。
$エヴァ
エヴァ
主張ではなく証拠
バージョン 0.1.0+e839c8a
モデル クロード ソネット 4-5
ブランチメイン
cwd ~/code/eva
スラッシュコマンドの場合は「/help」と入力します。
› キャッシュ書き込みとキャッシュ読み取りの違いは何ですか?
キャッシュ書き込みにはプロンプト プレフィックスが保存されるため、後の呼び出しで再読み取りをスキップできます。
それ。キャッシュ読み取りは、保存されたコピーをヒットする後の呼び出しの 1 つです。
› /コスト
セッション 1.2k 入力 / 340 出力 · キャッシュ 2.0k 書き込み / 1.1k 読み取り · コスト未報告
コストが報告されていないことに注意してください。 Anthropic も OpenAI も、応答とともにドルの数字を返しません。そこでエヴァは、トークンにどこかで調べた価格を掛けるのではなく、そう言いました。法案について議論できる数字は法案から得られなければなりません。
エヴァは早いですね。現在、これは非常に慎重に構築された基盤を備えた優れたターミナル チャット クライアントです。あなたの質問を読んで答えることができます。ファイルを読み取ったり、テストを実行したり、シェルにアクセスしたりすることはできません。ツールはまだありません。
Eva はコーディング エージェント用のコントロール プレーンを目指して構築されています。作業は、機械がチェックできる許容基準を備えた仕様として到着します。いくつかのハーネスは、隔離された環境で同じ仕様で競合します。 Eva が所有する検証者が実際に何を決定するか

パスすると、他のすべてのスコアと同じ記録からレース全体がスコアされます。
最初にこれほど慎重に基礎を構築する理由は、 docs/explanation/the-ladder.md のはしごです。通常のストーリーは、モデル→エージェント→ハーネス→工場という流れになります。このチェーンは 5 つの横線をスキップしており、それぞれの省略が失敗する既知の方法です。
19 のステージには、失敗する可能性がある終了テストがあります。そのうちの1つが構築されています。計画は草案です。出荷された段階ではありません。
Go 1.26 以降が必要です。他には何もありません。
git clone git@github.com:missingstudio/eva.git
CDエヴァ
go build -o eva ./cmd/eva
これにより、現在のディレクトリに単一のバイナリが生成されます。どこにでも置きたい場合は、PATH に置きます。
Eva は Anthropic と OpenAI に話しかけます。 1 つ選んでください。
エクスポート ANTHROPIC_API_KEY=sk-ant-...
./エヴァ
それが全体のセットアップです。 Anthropic がデフォルトなので、何も設定する必要はありません。
OpenAI の場合、設定ファイルを作成し、プロバイダーに名前を付けます。
./eva init # ~/.eva/config.toml を書き込みます
エクスポート OPENAI_API_KEY=sk-...
【プロバイダー】
名前 = " オープンアイ "
1 行で十分です。モデルとキー変数は選択したプロバイダーに従うため、大声で言わなくても gpt-5.6-terra が OPENAI_API_KEY を読み取ります。
ChatGPT または Codex サブスクリプションを使用する場合
OpenAI を毎月支払う場合は、API キーの代わりにそれを使用できます。
./eva ログイン
URL とショート コードが出力され、ブラウザで承認すると、資格情報が ~/.eva/auth.json に保存されます。次に、 ~/.eva/config.toml でモードを設定します。
【プロバイダー】
名前 = " オープンアイ "
認証 = " サブスクリプション "
Eva が実際に何を使用するかをいつでも確認してください:
./eva 認証ステータス
プロバイダー：オープンアイ
認証: サブスクリプション
ストア: /Users/you/.eva/auth.json
ログイン: アカウント acct_1a2b、2026 年 8 月 11 日月曜日 09:14:00 IST まで有効
重要
auth が決定し、それをオーバーライドするものはありません。 subscription と表示されている場合は、エクスポートされた OPENAI_API

_KEY は無視され、それを黙って使用するのではなく、eva 認証ステータスによってそのことがわかります。ほとんどのツールは最初に環境を試します。これにより、人々は気付かずに間違ったアカウントに 1 か月間請求することになります。 (なぜ)
キーが設定ファイルに書き込まれることはありません。 Eva は環境からそれを読み取るか、ログイン時にそれを取得します。これは履歴ファイル、ログ、またはモデルに送信されるものには決して表示されません。
引数と型を指定せずに eva を実行します。回答は到着すると次々と届きます。
中断しても安全です。会話は引き続き使用可能であり、履歴ファイルには会話を停止したことが記録されます。
行の先頭に / を入力します。これらはローカルで処理され、モデルに届くことはないため、費用はかかりません。
/model は、コンテキストを失わずに会話中にモデルを交換するため、次の回答にはあなたが話した内容が引き続きわかります。 Eva は有効なモデル名のリストを保持していません。先月まとめられたリストでは、先週リリースされたモデルが拒否されてしまうからです。プロバイダーが名前を認識できない場合、その回答は失敗し、その旨が通知されます。
/clear は、現在の会話からメッセージを削除するのではなく、新しい会話を開始します。どちらの方法でも、古いメッセージは履歴ファイルに残ります。 (なぜ)
eva -p は 1 つの質問に答え、それを標準出力に出力して終了します。
eva -p " このエラーの説明 " > Answer.md ||エコー「失敗しました」
応答が失敗した場合はゼロ以外で終了し、その理由を stderr に書き込みます。これにより、パイプラインで安全に使用できるようになります。stdout が答えであり、他には何もありません。
> これは何ですか？
応答なし - 資格情報が拒否されました
Provider.auth は「api_key」なので、Anthropic が拒否したのは次のキーです。
$ANTHROPIC_API_KEY
2 行、どちらも真実です。最初の部分では、ベンダーのエラー文書ではなく、Eva 自身の言葉で失敗の種類を挙げています。 2 番目は、Eva があなたのマシンについて何かをチェックしたとき、そしてその事実がちょうど 1 つのネックを残したときにのみ表示されます。

ステップ。ログインが見つからない場合は、eva ログインを実行してくださいと表示されます。これは間違いなく修正です。拒否されたキーは、どのキーが送信されたかを示して停止します。これは、取り消された、間違った組織、および一時停止されたアカウントはすべて、ここからは同じに見えるためです。決して壊れなかったものを修理するためにあなたを派遣すると、後で正しいはずだったヒントがすべて犠牲になります。 (なぜ)
コマンド
何をするのか
エヴァ
チャットを開く
eva -p "<質問>"
1 回応答し、標準出力に出力して終了します
エヴァ初期化
スターター設定ファイルを書き込む
エヴァログイン
サブスクリプションにサインインする
EVA認証ステータス
Eva がどのように認証するかを表示する
エヴァ助けて
このリストを表示
2 つのフラグですが、これは意図的なものです。 --config <path> は設定ファイルを選択し、-p <question> は 1 つの質問をします。設定は確認可能ですが、フラグは確認できないため、その他はすべて設定です。
eva init は、存在するすべてのオプションをコメントアウトして ~/.eva/config.toml を書き込みます。また、各オプションの横に、そのオプションがない場合に何が起こるかを示すメモが書かれています。あなたのために何も選ばれていません。行のコメントを解除して変更します。
何も書かれていない状態で、エヴァが何をしているかは次のとおりです。
model = " claude-sonnet-4-5 " # 省略した場合はプロバイダーに従います
【プロバイダー】
name = " anthropic " # anthropic または openai
auth = " api_key " # api_key またはサブスクリプション
api_key_env = " ANTHROPIC_API_KEY "
base_url = " " # プロキシ、ゲートウェイ、またはローカル サーバー
max_tokens = 0 # 0 プロバイダーが決定します
【痕跡】
path = " ~/.eva/trace.jsonl " # 履歴の保存先
kind = " jsonl " # どのライターが保持するか
【アイデンティティ】
テナント = "ローカル"
俳優 = "ローカル"
Actor_kind = " human " # 人間、エージェント、またはシステム
色、グリフ、間隔、キー バインディングは [theme] と [keymap.bind] の下にあり、スターター ファイルにはそれらもリストされています。どれも設定しないと、Eva は設定可能になる前とまったく同じように表示されます。色は端末の背景と表示内容に従うため、ウィットに適合します。

ハウトが言われてる。
設定は 4 つの場所から読み取られ、それぞれが以前の場所 (組み込みのデフォルト、ユーザーのファイル、プロジェクトのファイル、 --config ) を上書きします。 Eva はどれを使用しても問題なく動作します。
タイプミスはエラーであって、肩をすくめるものではありません。 modl = "..." と書き込むと、Eva は起動を拒否し、キーに名前を付けます。あなたが書いた内容を黙って無視する設定ファイルは、読み込まれない設定ファイルよりも悪いです。
リポジトリには、どこからでもアクセスできる .eva/config.toml を含めることができます。チームの外観を共有するのに便利です:
【テーマ。カラー】
人 = " #7AA6DC "
【テーマ。記号]
プロンプト = " › "
[ キーマップ .バインド]
フォロー = [ " ctrl+g " ]
警告
リポジトリは、Eva の外観を変えることはできますが、その内容を変えることはできません。インターネットからリポジトリのクローンを作成すると、Eva は最初の質問の前に、API キーを保持するプロセスでそのファイルを読み取ります。したがって、設定される可能性のあるリストは、外観とキー バインディングという短い許可リストになります。モデルプロバイダーを選択したり、トラフィックを別のサーバーに向けたり、キーの読み取り元の変数の名前を変更したり、履歴ファイルを移動したりすることはできません。そのファイル内のその他のものは、名前によって拒否されます。 (なぜ)
Eva は名前付けに厳格です。3 つの名前の下で同じ概念がコードベースの腐敗につながるためです。そのうちの 5 つはドキュメントとコードに表示されます。
完全なリストは CONTEXT.md で、試行され廃止された単語も含まれています。
Go モジュール 1 つ。データは一方向に流れ、コンパイラはそれを強制します。
あなたは入力します、あなたは読みます
│▲
▼ │
┌──────┐ ┌──────┐ ┌──────┐ ┌────┐
│ tui │───▶│ ループ │───▶│ プロバイダー │ │ レンダリング │
━━━━┘ ━━┬───┘ ━━━━┘ ━━━▲────┘
│ │
▼ 最初にコミットする │
┌───────┐───────

─────────┘
│ トレース │ その後表示
━───────┘
重要な部分は底面です。最初にファイルを通過しないものは何も画面に表示されません。レンダリング レイヤーは物理的にモデル、会話、または履歴ファイルと通信できません。レコードを取得して文字列を返します。インポートできるのはそれだけです。
これらの境界は、 .golangci.yml 内のパッケージごとの許可リストであり、厳密モードで実行されます。誰も明示的に許可されていないインポートはビルドに失敗します。リストを拡張することは、diff 内の目に見える行であり、その横に理由が示されており、偶然に起こることではありません。
ほとんどの作業は 3 つの決定によって行われます。
プロバイダーが知っているのは、ダイヤル方法、チャンクの読み取り方法、そして電話を切る方法だけです。キューイング、再試行、およびカウントは一度書き込まれて共有されます。プロバイダーの追加は数百行であり、機械のコピーではありません。 ( 0034 )
物事は自動的に登録されます。プロバイダーとファイル ライターは、設定が選択するセットに自分自身を追加します。したがって、配線層は実装名を指定せず、オプションをリストするエラーが古くなることはありません。 ( 0028 )
画面は読み取り専用ビューです。レコードだけをレンダリングします。 ( 0015 )
小切手を作る
それはまさに CI が実行することです: フォーマット、ビルド、精査、lint、テスト、すべてのパッケージ。
リンター バージョンは固定されており、 go run 経由で実行されるため、何もインストールしなくても、結果は CI と一致します。
テストでは、実際のプロトコルを使用するローカル サーバーに対して実際の Anthropic および OpenAI コードを実行します。モックは 2 番目の実装であるため、モック プロバイダーはありません。

[切り捨てられた]

## Original Extract

An autonomous, multi-tenant, AI-native software factory. - missingstudio/eva

GitHub - missingstudio/eva: An autonomous, multi-tenant, AI-native software factory. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
missingstudio
/
eva
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
91 Commits 91 Commits .github/ workflows .github/ workflows cmd/ eva cmd/ eva docs docs internal internal .gitignore .gitignore .golangci.yml .golangci.yml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTEXT.md CONTEXT.md LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum View all files Repository files navigation
An autonomous, multi-tenant, AI-native software factory
Most AI tools tell you what they did. Eva writes it down first.
Every question you ask, every answer that comes back, every retry, and every token counted is appended to a file on your disk as it happens. What you see on screen is read back out of that file. If Eva shows you an answer, the file already had it. There is no second version of events.
That sounds like a small thing. It is the whole design.
$ eva
EVA
Evidence, not claims
version 0.1.0+e839c8a
model claude-sonnet-4-5
branch main
cwd ~/code/eva
type /help for slash commands
› what's the difference between a cache write and a cache read?
A cache write stores your prompt prefix so later calls can skip re-reading
it. A cache read is one of those later calls hitting the stored copy…
› /cost
session 1.2k in / 340 out · cache 2.0k write / 1.1k read · cost unreported
Notice cost unreported . Neither Anthropic nor OpenAI returns a dollar figure with a response. So Eva says so, rather than multiplying tokens by a price it looked up somewhere. A number you can argue with a bill about has to come from the bill.
Eva is early. Today it is a good terminal chat client with a very carefully built foundation. It can read your question and answer it. It cannot read your files, run your tests, or touch your shell — there are no tools yet.
Eva is being built toward a control plane for coding agents. Work arrives as a spec with acceptance criteria a machine can check. Several harnesses race the same spec in isolated environments. A verifier Eva owns decides what actually passed, and the whole race is scored from the same record everything else is scored from.
The reason for building the foundation this carefully first is the ladder in docs/explanation/the-ladder.md . The usual story goes model → agent → harness → factory. That chain skips five rungs, and each omission is a known way this fails:
Nineteen stages, each with an exit test it can fail. One of them is built. The plan is a draft; the stage that shipped is not.
You need Go 1.26 or newer. Nothing else.
git clone git@github.com:missingstudio/eva.git
cd eva
go build -o eva ./cmd/eva
That produces a single binary in the current directory. Put it on your PATH if you want it everywhere.
Eva talks to Anthropic and OpenAI. Pick one.
export ANTHROPIC_API_KEY=sk-ant-...
./eva
That's the whole setup. Anthropic is the default, so nothing needs configuring.
For OpenAI, create a settings file and name the provider:
./eva init # writes ~/.eva/config.toml
export OPENAI_API_KEY=sk-...
[ provider ]
name = " openai "
One line is enough. The model and the key variable follow the provider you picked, so you get gpt-5.6-terra reading OPENAI_API_KEY without saying either out loud.
With a ChatGPT or Codex subscription
If you pay OpenAI monthly, you can use that instead of an API key:
./eva login
It prints a URL and a short code, you approve it in a browser, and the credential is saved to ~/.eva/auth.json . Then set the mode in ~/.eva/config.toml :
[ provider ]
name = " openai "
auth = " subscription "
Check what Eva will actually use at any time:
./eva auth status
provider: openai
auth: subscription
store: /Users/you/.eva/auth.json
login: account acct_1a2b, valid until Mon, 11 Aug 2026 09:14:00 IST
Important
auth decides, and nothing overrides it. If it says subscription , an exported OPENAI_API_KEY is ignored, and eva auth status will tell you so rather than quietly using it. Most tools try the environment first, which is how people bill the wrong account for a month without noticing. ( why )
Your key is never written to a settings file. Eva reads it from the environment, or gets it when you log in. It never appears in the history file, a log, or anything sent to a model.
Run eva with no arguments and type. Answers stream in as they arrive.
Interrupting is safe. The conversation stays usable and the history file records that you stopped it.
Type / at the start of a line. These are handled locally and never reach a model, so they cost nothing.
/model swaps the model mid-conversation without dropping context, so the next answer still knows what you talked about. Eva doesn't keep a list of valid model names, because a list compiled last month would reject a model released last week. If the provider doesn't recognise the name, that answer fails and tells you.
/clear starts a new conversation rather than deleting messages from the current one. Your old messages are still in the history file either way. ( why )
eva -p answers one question, prints it to stdout, and exits.
eva -p " explain this error " > answer.md || echo " that failed "
It exits non-zero when the answer failed, and writes the reason to stderr. That makes it safe to use in a pipeline: stdout is the answer and nothing else.
› what is this?
No response — the credential was refused
provider.auth is "api_key", so what anthropic refused is the key in
$ANTHROPIC_API_KEY
Two lines, both true. The first names the kind of failure, in Eva's own words rather than the vendor's error document. The second appears only when Eva checked something about your machine, and only when that fact leaves exactly one next step. A missing login says run eva login because that is certainly the fix. A refused key says which key was sent and stops, because revoked, wrong organisation, and suspended account all look identical from here. Sending you to fix something that was never broken costs you every later hint that would have been right. ( why )
Command
What it does
eva
Open the chat
eva -p "<question>"
Answer once, print to stdout, exit
eva init
Write a starter settings file
eva login
Sign in to a subscription
eva auth status
Show how Eva will authenticate
eva help
Show this list
Two flags, and that's deliberate: --config <path> picks a settings file, -p <question> asks one question. Everything else is a setting, because settings are reviewable and flags are not.
eva init writes ~/.eva/config.toml with every option present but commented out, and a note beside each saying what happens without it. Nothing is chosen for you. Uncomment a line to change it.
Here is what Eva does with none of it written down:
model = " claude-sonnet-4-5 " # follows the provider if you leave it out
[ provider ]
name = " anthropic " # anthropic or openai
auth = " api_key " # api_key or subscription
api_key_env = " ANTHROPIC_API_KEY "
base_url = " " # a proxy, gateway, or local server
max_tokens = 0 # 0 lets the provider decide
[ trace ]
path = " ~/.eva/trace.jsonl " # where the history goes
kind = " jsonl " # which writer keeps it
[ identity ]
tenant = " local "
actor = " local "
actor_kind = " human " # human, agent, or system
Colours, glyphs, spacing, and key bindings live under [theme] and [keymap.bind] , and the starter file lists those too. Set none of them and Eva looks exactly as it did before any of it was configurable. Colours follow your terminal's background and what it can display, so it fits in without being told.
Settings are read from four places, each overriding the one before: built-in defaults, your file, the project's file, then --config . Eva works fine with none of them.
A typo is an error, not a shrug. Write modl = "..." and Eva refuses to start and names the key. A settings file that silently ignores what you wrote is worse than one that won't load.
A repo can carry .eva/config.toml , found by walking up from wherever you are. Handy for sharing a team's look:
[ theme . colors ]
person = " #7AA6DC "
[ theme . symbols ]
prompt = " › "
[ keymap . bind ]
follow = [ " ctrl+g " ]
Warning
A repo can change how Eva looks, never what it does. You clone a repo from the internet, and Eva reads that file before your first question, in a process holding your API key. So the list of what it may set is a short allow-list: appearance and key bindings. It cannot pick the model provider, point traffic at another server, rename the variable your key is read from, or move your history file. Anything else in that file is refused by name. ( why )
Eva is strict about naming, because the same concept under three names is how a codebase rots. Five of them show up in the docs and the code:
The full list is CONTEXT.md , including the words that were tried and retired.
One Go module. Data flows one way, and the compiler enforces it.
you type you read
│ ▲
▼ │
┌──────┐ ┌──────┐ ┌───────────┐ ┌────────┐
│ tui │───▶│ loop │───▶│ providers │ │ render │
└──────┘ └──┬───┘ └───────────┘ └───▲────┘
│ │
▼ committed first │
┌───────┐ ───────────────────────────┘
│ trace │ then shown
└───────┘
The important part is the bottom. Nothing reaches your screen that didn't go through the file first. The rendering layer physically cannot talk to a model, a conversation, or the history file. It takes records and returns strings, and that's all it's allowed to import.
Those boundaries are an allow-list per package in .golangci.yml , running in strict mode. An import nobody explicitly permitted fails the build. Widening a list is a visible line in a diff with a reason next to it, not something that happens by accident.
Three decisions do most of the work:
A provider only knows how to dial, read a chunk, and hang up. Queueing, retrying, and counting are written once and shared. Adding a provider is a few hundred lines, not a copy of the machinery. ( 0034 )
Things register themselves. Providers and file writers add themselves to the set that settings choose from. So the wiring layer names no implementation, and the error listing your options cannot go stale. ( 0028 )
The screen is a read-only view. It renders records and nothing else. ( 0015 )
make check
That is exactly what CI runs: formatting, build, vet, lint, tests, every package.
The linter version is pinned and run via go run , so you don't install anything and your results match CI.
Tests drive the real Anthropic and OpenAI code against a local server speaking the real protocol. There's no mock provider, because a mock is a second implementation that can disa

[truncated]
