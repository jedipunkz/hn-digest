---
source: "https://github.com/AIStoryHub/etincel"
hn_url: "https://news.ycombinator.com/item?id=49248600"
title: "Show HN: Étincel trains Claude's voice for PR reviews and code comments"
article_title: "GitHub - AIStoryHub/etincel: Non-fiction writing connector for Claude Code, Claude Desktop, and MCP-enabled tools: trainable voice, premade tone presets, and a deterministic anti-AI-writing-tells audit. · GitHub"
author: "jeeps1911"
captured_at: "2026-08-10T19:48:11Z"
capture_tool: "hn-digest"
hn_id: 49248600
score: 1
comments: 0
posted_at: "2026-08-10T19:35:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Étincel trains Claude's voice for PR reviews and code comments

- HN: [49248600](https://news.ycombinator.com/item?id=49248600)
- Source: [github.com](https://github.com/AIStoryHub/etincel)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T19:35:38Z

## Translation

タイトル: HN を表示: エタンセルが PR レビューとコード コメント用にクロードの声をトレーニング
記事のタイトル: GitHub - AIStoryHub/etincel: Claude Code、Claude Desktop、および MCP 対応ツール用のノンフィクション執筆コネクタ: トレーニング可能な音声、事前に作成されたトーン プリセット、および決定論的な反 AI ライティング テル監査。 · GitHub
説明: Claude Code、Claude Desktop、および MCP 対応ツール用のノンフィクション執筆コネクタ: トレーニング可能な音声、事前に作成されたトーン プリセット、および決定論的な反 AI ライティング テル監査。 - AIStoryHub/エティンセル

記事本文:
GitHub - AIStoryHub/etincel: Claude Code、Claude Desktop、および MCP 対応ツール用のノンフィクション執筆コネクタ: トレーニング可能な音声、事前に作成されたトーン プリセット、決定論的な反 AI ライティング テル監査。 · GitHub
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
AISストーリーハブ
/
エチンセル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
この GitHub アクションをプロジェクトで使用する このアクションを既存のワークフローに追加するか、新しいワークフローを作成します

マーケットプレイスのメイン ブランチで表示 タグ ファイルに移動 コード さらにアクション メニューを開く フォルダーとファイル
46 コミット 46 コミット .circleci .circleci .claude-plugin .claude-plugin .github/ workflows .github/ workflows scripts scripts skill/ etincel-nonfiction skill/ etincel-nonfiction src src .gitattributes .gitattributes .gitignore .gitignore .mcp.json .mcp.json COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md action.yml action.yml 効力ベースライン.json 有効性ベースライン.json orb.yml orb.yml パッケージロック.json パッケージロック.json パッケージ.json package.json server.json server.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エティンセル: ノンフィクション執筆コネクター
___ ___ ___ ___ ___ ___
/\ \ /\ \ ___ /\__\ /\ \ /\ \ /\__\
/::\ \ \:\ \ /\ \ /::| | /::\ \ /::\ \ /:/ /
/:/\:\ \ \:\ \ \:\ \ /:|:| | /:/\:\ \ /:/\:\ \ /:/ /
/::\~\:\ \ /::\ \ /::\__\ /:/|:| |__ /:/ \:\ \ /::\~\:\ \ /:/ /
/:/\:\ \:\__\ /:/\:\__\ __/:/\/__/ /:/ |:| /\__\ /:/__/ \:\__\ /:/\:\ \:\__\ /:/__/
\:\~\:\ \/__/ /:/ \/__/ /\/:/ / \/__|:|/:/ / \:\ \ \/__/ \:\~\:\ \/__/ \:\ \
\:\ \:\__\ /:/ / \::/__/ |:/:/ / \:\ \ \:\ \:\__\ \:\ \
\:\ \/__/ \/__/ \:\__\ |::/ / \:\ \ \:\ \/__/ \:\ \
\:\__\ \/__/ /:/ / \:\__\ \:\__\ \:\__\
\/__/ \/__/ \/__/ \/__/ \/__/
これは、このリポジトリのローカル/stdio エンジンではなく、ホストされているリモート サーバー ( etincel.ai/api/mcp ) を監査します。 2 つは同じツールを公開しますが、別のデプロイメントとして実行されます。
Claude Code、Claude Desktop、および既存の書き込み内容の上に人間のようなノンフィクション オーサリングのレイヤーを追加する MCP 対応ツール用のコネクタ。電子メール クライアント、エディター、または CMS を置き換えるものではありません。それは、散文がそこに到達する前に、自分の文章から訓練するか、あらかじめ用意された一連の音声から選択する音声で、散文を形成します。

モーショントーンのプリセットを使用すると、AI があなたの言葉を黙って書き直すのではなく、透過的に伝えるようフラグを立てます。
AI が起草した散文は、認識可能な形をしています。統一された段落、権威のヘッジ、コンマで済むところの全角ダッシュ、欠陥のないケーススタディ、あまりにもきれいに解決された結びです。その形状により、たとえ事実に問題がないとしても、テキストが AI によって書かれたように感じられます。このコネクタは、その形状を回避するルールをエンコードします。また、同様に重要なことに、音声を静かに上書きするのではなく、検出された内容とその理由を表示します。あなたは著者のままです。
MCP サーバー ( src/server.ts ) は 19 のツールを公開します。
list_styles : 事前に作成されたトーン プリセット、トレーニングした音声、および (リポジトリ ローカルの .etincelrc で定義されている場合) 共有チーム スタイル
get_style_guide : 1 つのスタイルの製図手順
train_style : 自分のライティングサンプルから音声を学習します (文のリズム、短縮率、全角ダッシュの習慣、段落の差異、繰り返しのフレーズ: 推測ではなく測定)
create_style_from_dials : サンプルの代わりに、明示的な形式性/暖かさ/直接性と機械式ダイヤルからスタイルを構築します。
update_style : トレーニングされた音声の名前を変更するか、そのダイヤルを適切な位置に調整します
fork_style : プリセットのダイヤルをコピーして、再トレーニングまたは手動調整できる新しいトレーニングされた音声にガイドします。または、ホストされているギャラリーで公開された別のインストーラーのスタイルをフォークします ( handle/slug としてアドレス指定されます。例: jpleblanc/blunt-memo 。etincel.ai/v/handle/slug の公開ページに表示されているのと同じアドレスです)。パブリック スタイルのフォークは、etincel.ai を取得するためにネットワーク呼び出しを 1 回行います。プリセット フォークはこのインストールから離れることはありません。
delete_style : トレーニングされた音声を完全に削除します
set_default_style : 繰り返すことなく、どのスタイルを使用するかを覚えてください。
check_voice_match : ドラフトの測定されたリズムをトレーニングされた音声のベースラインと比較します。リズム

m/mechanics チェック。著者名や AI 検出チェックではなく、短い入力に対する信頼度は低い
check_self_repetition : AI が教えるのではなく、草稿と音声自身の最近のトレーニング サンプルを比較して習慣を確認します。過去のいくつかの曲で繰り返される同じオープナーまたはフレーズ (「最近の 6 曲のうち 4 曲でこのように開いています」)。今のところローカルインストールのみ
Audit_text : AI の決定論的でルールベースのスキャンは、段階、重大度を伴う特定の結果、および強度シグナル (特異性、具体性と抽象性の比率、文章とリズムのバリエーション) を返すため、修正によって散文が平坦になることはありません。オプションのレジスター ( email / blog / memo / speech / social / docs / general 、デフォルトの general ) を使用して、テキストの種類に対する厳密性を調整します。 docs は、Markdown 構造の誤検知 (見出し、太字の用語) を抑制し、よりパンチの効いた短い形式のコピーではなく、長い形式の参照散文に対してリズム/語彙の検出を再調整します。
add_banned_word / Remove_banned_word : 独自の禁止語彙リストを管理し、組み込みコーパスとともに Audit_text によってチェックされます。
add_custom_word / Remove_custom_word : 「決してフラグを立てない」リストを維持します: 組織独自の頭字語や社内用語、企業辞書の場合
list_dictionary : スコープの禁止語/カスタム語、および (スタイルの場合) グローバル リストとマージされた後に実際に適用される内容を確認します。
set_style_instructions / clear_style_instructions / get_style_instructions : スコープ (必須要素、禁止トピック、書式制約) のフリーテキストのドラフト ルールを保存、削除、または読み取ります。辞書が Audit_text にマージされるのと同じ方法で get_style_guide にマージされます。
クロード コード / クロード デスクトップ スキル ( skill/etincel-nonfiction/ ) は、意味のある長さのノンフィクションの散文を草案または改訂するときにこれらのツールを使用します。
訓練された声、ディ

アクションとデフォルトのスタイルは ~/.etincel/ にローカルに存在します。どこにも何も送信されません。 Audit_text は、モデル呼び出しではなく、プレーンな決定論的コード (文字列分析 + AI 記述の精選されたコーパス) です。 1 つの例外は、 fork_style を介してパブリック スタイルをフォークすることです。これは、 etincel.ai のパブリック ギャラリーからそのスタイルのガイドをフェッチします (決して送信しません)。プリセットやこのリスト内の他のものをフォークしても、ネットワークにはまったく接続されません。
組み込みの AI-Tell コーパスを超えて、独自の禁止単語リストや「常に許可」単語リストを管理できます。「禁止単語リストに [単語] を追加してください」または「カスタム単語リストに [単語] を追加してください。それは私たちの単語の 1 つです」などのことをクロード (または任意の MCP クライアント) に伝えるだけです。各リストはスコープ: グローバル (どこにでも適用され、スタイル名が指定されていない場合のデフォルト) または特定のスタイル ID で存在します。そのスタイルに対して監査を行うと、そのリストはグローバルの上にマージされます。 list_dictionary は、スコープに対して保存されているものと、スタイルに対して有効に結合されたリストを示します。グローバル リストの編集は、すべてのスタイル間で単語の同期を保つための方法としてすでに使用されています。つまり、audit_text または list_dictionary が実行されるたびに、ライブで自動的にマージされます。
/プラグイン マーケットプレイス AIStoryHub/etincel を追加
/プラグインインストール エチンセルノンフィクション
または、ローカル クローンから: /plugin マーケットプレイスに /path/to/etincel を追加します。
クロード デスクトップ / 他の MCP ホスト
構築されたサーバーを MCP 構成に指定します。
{
"mcpサーバー": {
"エチンセル-ノンフィクション" : {
"コマンド" : "ノード" ,
"args" : [ " /path/to/etincel/dist/server.js " ]
}
}
}
最初にビルドします: npm install && npm run build 。
ホストされたバージョンも https://etincel.ai/api/mcp で入手でき、公開されています
stdio の代わりにアカウントごとの認証を使用して、Streamable HTTP 経由で同じツールを使用します。
MCP クライアントを直接指定します。
{
"mcpサーバー": {
"エチンセル-ノンフィクション" : {
"url" : " https://etincel.ai/api/mcp "
}
}
}
Th

ホストされているサーバーはこのリポジトリの一部ではありません。このリポジトリはローカル/stdio です
エンジン、CLI、およびホストされたバージョンがその上に構築されるスキル。
インストールしたら、Claude Code または Claude Desktop 内で、「遅延についてチームにメールを下書きする」、「X についてブログ投稿を書く」、「このメモをクリーンアップする」など、通常要求することを要求するだけです。このスキルは、意味のある長さのノンフィクションの散文に対して自動的に習得されます。自分の声をトレーニングするには:
私が書いた 3 つのメールから「me」というスタイルをトレーニングします: [サンプルを貼り付け]
次に、リクエストごとに名前を付けるか (「これを私の声で書いてください」)、デフォルトとして設定します。
12 の既成プリセットが箱から出してすぐに出荷されます。6 つの感情的なトーン (Direct & Warm、Executive Brief、Reflective Essayist、Founder Memo、Plainspoken Analyst、Wry & Candid) と 6 つのユースケース プリセット (PR レビュー、コード コメント、Slack メッセージ、LinkedIn 投稿、Web サイトのコピー、ブログ投稿) です。それぞれに、形式的/温かさ/率直さのダイヤルに加えて、空白を埋めるテンプレートではなく、製図コンテキストとしてモデルに入力される文のリズムと音声の説明が含まれています。サーバーはこれらを src/data/presets.json から読み取ります。 fork_style を使用して、任意のプリセットをトレーニング済みの音声にフォークして、独自の音声にします。
Audit_text は内部の純粋な関数であるため、チャット クライアントの外部で散文をリンティングするための CLI としても出荷されます (README、ドキュメント、CI の PR 説明)。
npx etincel lint 'docs/**/*.md'
npx etincel lint README.md --register docs --threshold yellow
一致したファイルの層が --threshold (デフォルトはオレンジ色) 以上である場合、ゼロ以外で終了します。 .md / .mdx ファイルはデフォルトでドキュメントに自動的に登録されます (実際の見出しはチャットボットの指示ではないため、Markdown 構造の誤検知が抑制されます)。 --register を渡してオーバーライドします。機械可読レポートの場合は --json を追加します。完全なオプション リストについては、npx etincel lint --help を実行してください。
GitHub A

ction は同じ CLI をラップします (実際の例については、このリポジトリの action.yml および .github/workflows/lint.yml を参照してください)。
- 使用: AIStoryHub/etincel@main
付き:
パターン: " docs/**/*.md README.md "
しきい値：オレンジ
リポジトリローカル構成: 辞書、手順、共有チームスタイル
チームのルールは、各人のローカル ~/.etincel/ にのみ存在する必要はありません。 .etincelrc (または .etincelrc.json / etincel.config.json ) をリポジトリのルートにドロップすると、CLI とローカル (stdio) サーバーによって自動的に取得され、コード レビューでレビュー可能になり、誰かが去ったときに表示されなくなってしまうのではなく、バージョン管理されます。
{
"bannedWords" : [ " Acme Cloud Platform " ],
"allowedWords" : [ " レバレッジ " ],
"登録" : "ドキュメント" ,
"しきい値" : " オレンジ " 、
"instructions" : " 必ず最後に 1 行の CTA を含めてください。 " ,
"スタイル" : {
"名前" : "ハウスボイス" ,
「ダイヤル」: {
「形式的」：6、
「温かさ」：4、
「直接性」：7、
"文の長さ" : 40 、
"センテンスリズムバリアンス" : 50 、
"段落分散" : 30 、
"収縮使用" : 20 、
"emDashUse" : 0 、
"フラグメント許容値" : 10 、
"質問の使用" : 5 、
「エントロピー」：60
}
}
}
BandWords / allowedWords は、アカウント/スタイル辞書にあるものと統合されます。レジスタ/しきい値はリポジトリとして機能します

[切り捨てられた]

## Original Extract

Non-fiction writing connector for Claude Code, Claude Desktop, and MCP-enabled tools: trainable voice, premade tone presets, and a deterministic anti-AI-writing-tells audit. - AIStoryHub/etincel

GitHub - AIStoryHub/etincel: Non-fiction writing connector for Claude Code, Claude Desktop, and MCP-enabled tools: trainable voice, premade tone presets, and a deterministic anti-AI-writing-tells audit. · GitHub
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
AIStoryHub
/
etincel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Folders and files
46 Commits 46 Commits .circleci .circleci .claude-plugin .claude-plugin .github/ workflows .github/ workflows scripts scripts skills/ etincel-nonfiction skills/ etincel-nonfiction src src .gitattributes .gitattributes .gitignore .gitignore .mcp.json .mcp.json CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md action.yml action.yml efficacy-baselines.json efficacy-baselines.json orb.yml orb.yml package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json View all files Repository files navigation
Étincel: Non-Fiction Writing Connector
___ ___ ___ ___ ___ ___
/\ \ /\ \ ___ /\__\ /\ \ /\ \ /\__\
/::\ \ \:\ \ /\ \ /::| | /::\ \ /::\ \ /:/ /
/:/\:\ \ \:\ \ \:\ \ /:|:| | /:/\:\ \ /:/\:\ \ /:/ /
/::\~\:\ \ /::\ \ /::\__\ /:/|:| |__ /:/ \:\ \ /::\~\:\ \ /:/ /
/:/\:\ \:\__\ /:/\:\__\ __/:/\/__/ /:/ |:| /\__\ /:/__/ \:\__\ /:/\:\ \:\__\ /:/__/
\:\~\:\ \/__/ /:/ \/__/ /\/:/ / \/__|:|/:/ / \:\ \ \/__/ \:\~\:\ \/__/ \:\ \
\:\ \:\__\ /:/ / \::/__/ |:/:/ / \:\ \ \:\ \:\__\ \:\ \
\:\ \/__/ \/__/ \:\__\ |::/ / \:\ \ \:\ \/__/ \:\ \
\:\__\ \/__/ /:/ / \:\__\ \:\__\ \:\__\
\/__/ \/__/ \/__/ \/__/ \/__/
This audits the hosted remote server ( etincel.ai/api/mcp ), not the local/stdio engine in this repo; the two expose the same tools but run as separate deployments.
A connector for Claude Code, Claude Desktop, and any MCP-enabled tool that adds a layer of human-like non-fiction authoring on top of whatever you already write in. It does not replace your email client, editor, or CMS. It shapes the prose before it gets there, in a voice you either train from your own writing or pick from a set of premade emotional-tone presets, and it flags AI writing tells transparently instead of silently rewriting your words.
AI-drafted prose has a recognizable shape: uniform paragraphs, hedged authority, em dashes where a comma would do, case studies with no flaws, closings that resolve too neatly. That shape is what makes text feel AI-written even when it's factually fine. This connector encodes the rules that avoid that shape, and, just as important, it shows you what it found and why, instead of quietly overwriting your voice. You stay the author.
An MCP server ( src/server.ts ) exposing nineteen tools:
list_styles : premade tone presets, any voices you've trained, and (if a repo-local .etincelrc defines one) a shared team style
get_style_guide : the drafting instructions for one style
train_style : learn a voice from your own writing samples (sentence rhythm, contraction rate, em-dash habits, paragraph variance, recurring phrasing: measured, not guessed)
create_style_from_dials : build a style from explicit formality/warmth/directness and mechanical dials instead of samples
update_style : rename a trained voice or adjust its dials in place
fork_style : copy a preset's dials and guide into a new trained voice you can retrain or hand-tune, or fork another installer's style once they've published it publicly on the hosted gallery (addressed as handle/slug , e.g. jpleblanc/blunt-memo , the same address shown on its public page at etincel.ai/v/handle/slug ); a public-style fork makes one network call to etincel.ai to fetch it, a preset fork never leaves this install
delete_style : permanently remove a trained voice
set_default_style : remember which style to use without repeating yourself
check_voice_match : compare a draft's measured rhythm against a trained voice's baseline. A rhythm/mechanics check, not an authorship or AI-detection check, and low-confidence on short input
check_self_repetition : compare a draft against a voice's own recent training samples for habits, not AI tells: the same opener, or a phrase, recurring across several past pieces ("you've opened this way in 4 of your last 6 pieces"). Local install only for now
audit_text : a deterministic, rules-based scan for AI tells, returning a tier, specific findings with severity, and a strengths signal (specificity, concrete-vs-abstract ratio, sentence-rhythm variation) so fixes don't flatten the prose. Takes an optional register ( email / blog / memo / essay / social / docs / general , default general ) to calibrate strictness against the kind of text it is: docs suppresses Markdown-structure false positives (headings, bolded terms) and recalibrates rhythm/vocabulary detection against long-form reference prose instead of punchier short-form copy
add_banned_word / remove_banned_word : maintain your own banned-vocabulary list, checked by audit_text alongside the built-in corpus
add_custom_word / remove_custom_word : maintain a "never flag this" list: an org's own acronyms or house terms, the corporate-dictionary case
list_dictionary : see a scope's banned/custom words, and (for a style) what actually applies once merged with the global list
set_style_instructions / clear_style_instructions / get_style_instructions : save, remove, or read free-text drafting rules for a scope (required elements, forbidden topics, format constraints), merged into get_style_guide the same way dictionaries merge into audit_text
A Claude Code / Claude Desktop skill ( skills/etincel-nonfiction/ ) that uses those tools when you're drafting or revising non-fiction prose of any meaningful length.
Trained voices, dictionaries, and your default style live locally in ~/.etincel/ : nothing is sent anywhere. audit_text is plain deterministic code (string analysis + a curated corpus of AI-writing tells), not a model call. The one exception is forking a public style via fork_style , which fetches (never sends) that style's guide from etincel.ai 's public gallery; forking a preset, or anything else in this list, still touches the network not at all.
Beyond the built-in AI-tell corpus, you can maintain your own banned and "always allowed" word lists: just tell Claude (or any MCP client) things like "add [word] to my banned words list" or "add [word] to my custom words list, it's one of ours." Each list lives at a scope : global (applies everywhere, the default when no style is named) or a specific style id, whose list is merged on top of global when you audit against that style. list_dictionary shows what's saved for a scope, plus the effective merged list for a style. Editing the global list is already the way to keep a word in sync across every style: it's merged in automatically, live, every time audit_text or list_dictionary runs.
/plugin marketplace add AIStoryHub/etincel
/plugin install etincel-nonfiction
Or from a local clone: /plugin marketplace add /path/to/etincel .
Claude Desktop / other MCP hosts
Point your MCP config at the built server:
{
"mcpServers" : {
"etincel-nonfiction" : {
"command" : " node " ,
"args" : [ " /path/to/etincel/dist/server.js " ]
}
}
}
Build first: npm install && npm run build .
A hosted version is also available at https://etincel.ai/api/mcp , exposing
the same tools over Streamable HTTP with per-account auth instead of stdio.
Point any MCP client at it directly:
{
"mcpServers" : {
"etincel-nonfiction" : {
"url" : " https://etincel.ai/api/mcp "
}
}
}
The hosted server isn't part of this repo; this repo is the local/stdio
engine, CLI, and skill that the hosted version is built on top of.
Once installed, just ask for what you'd normally ask for, like "draft an email to the team about the delay," "write a blog post about X," or "clean up this memo," inside Claude Code or Claude Desktop. The skill picks up automatically for non-fiction prose of meaningful length. To train your own voice:
Train a style called "me" from these three emails I wrote: [paste samples]
Then either name it per-request ("write this in my voice") or set it as default:
Twelve premade presets ship out of the box: six emotional tones (Direct & Warm, Executive Brief, Reflective Essayist, Founder Memo, Plainspoken Analyst, Wry & Candid) plus six use-case presets (PR Review, Code Comment, Slack Message, LinkedIn Post, Website Copy, Blog Post). Each carries formality/warmth/directness dials plus a sentence-rhythm and voice description that gets fed to the model as drafting context, not a template that fills in blanks. The server reads these from src/data/presets.json . Fork any preset into a trained voice with fork_style to make it your own.
audit_text is a pure function under the hood, so it also ships as a CLI, for linting prose outside a chat client (READMEs, docs, PR descriptions in CI):
npx etincel lint 'docs/**/*.md'
npx etincel lint README.md --register docs --threshold yellow
Exits non-zero if any matched file's tier is at or above --threshold (default orange ). .md / .mdx files default to the docs register automatically (suppresses the Markdown-structure false positives, since a real heading isn't a chatbot tell); pass --register to override. Add --json for a machine-readable report. Run npx etincel lint --help for the full option list.
A GitHub Action wraps the same CLI (see action.yml , and .github/workflows/lint.yml in this repo for a working example):
- uses : AIStoryHub/etincel@main
with :
patterns : " docs/**/*.md README.md "
threshold : orange
Repo-local config: dictionary, instructions, and a shared team style
A team's rules don't have to live only in each person's local ~/.etincel/ . Drop a .etincelrc (or .etincelrc.json / etincel.config.json ) at the repo root and it's picked up automatically by the CLI and by the local (stdio) server, reviewable in code review and versioned instead of invisible and gone when someone leaves:
{
"bannedWords" : [ " Acme Cloud Platform " ],
"allowedWords" : [ " leverage " ],
"register" : " docs " ,
"threshold" : " orange " ,
"instructions" : " Always include a one-line CTA at the end. " ,
"style" : {
"name" : " House Voice " ,
"dials" : {
"formality" : 6 ,
"warmth" : 4 ,
"directness" : 7 ,
"sentenceLength" : 40 ,
"sentenceRhythmVariance" : 50 ,
"paragraphVariance" : 30 ,
"contractionUse" : 20 ,
"emDashUse" : 0 ,
"fragmentTolerance" : 10 ,
"questionUse" : 5 ,
"entropy" : 60
}
}
}
bannedWords / allowedWords merge alongside whatever's in your account/style dictionary; register / threshold act as repo-w

[truncated]
