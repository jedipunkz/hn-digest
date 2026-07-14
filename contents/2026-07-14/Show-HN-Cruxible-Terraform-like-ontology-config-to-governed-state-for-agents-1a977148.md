---
source: "https://github.com/cruxible-ai/cruxible"
hn_url: "https://news.ycombinator.com/item?id=48907464"
title: "Show HN: Cruxible – Terraform-like ontology config to governed state for agents"
article_title: "GitHub - cruxible-ai/cruxible: Hard, verifiable state layer for AI agents · GitHub"
author: "rmalone1097"
captured_at: "2026-07-14T15:18:16Z"
capture_tool: "hn-digest"
hn_id: 48907464
score: 1
comments: 1
posted_at: "2026-07-14T14:26:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Cruxible – Terraform-like ontology config to governed state for agents

- HN: [48907464](https://news.ycombinator.com/item?id=48907464)
- Source: [github.com](https://github.com/cruxible-ai/cruxible)
- Score: 1
- Comments: 1
- Posted: 2026-07-14T14:26:49Z

## Translation

タイトル: Show HN: Cruxible – エージェントの管理された状態への Terraform のようなオントロジー構成
記事タイトル: GitHub - cruxible-ai/cruxible: AI エージェント用のハードで検証可能な状態レイヤー · GitHub
説明: AI エージェント用のハードで検証可能な状態レイヤー。 GitHub でアカウントを作成して、cruxible-ai/cruxible の開発に貢献してください。
HN テキスト: 汎用テキスト メモリは、明白に述べることができ、めったに変更されず、思い出すだけで済む情報を表現するのに適していますが、いくつかの構造的問題が未解決のまま残されています。 1) 汎用メモリは確定状態を確立せず、コンテキストを返します。テキストベースのメモリは情報を保存および表示できますが、コンテキストに追加されたクレームは、実行されるたびに確率的 LLM によって再解釈されます。私は、特に自然なライフサイクルを持つ知識について、マークダウン Wiki の内容の妥当性を再確認するようエージェントに常に依頼していることに気づきました。さらに、テキスト内の主張の出所や帰属を確立する統一的な方法はありません。git コミット メッセージに依存するのは良い標準とは思えません。 2) LLM は、特にコーパス全体にわたって関係を構築する必要がある場合、それ自体では関係情報に対する決定論的なクエリを実行できません。これは、複数のセッション間および複数エージェントのワークフローにとって非常に重要です。たとえ 1 つのエージェントまたはセッションが正しく対応したとしても、組織の真実が明確かつ決定的に通過可能でない場合、各エージェントまたはセッションが独自の内部表現で動作するバベルの塔の状況が生じます。 3) LLM は、独自の書き込み時間境界を強制することを信頼できません。テキストとして表現されたポリシーは、モデルによってコンテキストとしてのみ解釈されます。より優れたモデルほど、ポリシーに従う可能性が高くなります。しかし、LLM 自体は、どれほど優れていても、それ自体を保証することはできません。

n 無効な移行が拒否されるか、独立したレビューが必要な場合に申し立ての作成者がそれを承認できない場合。重要な真実を得るには、書き込み時の強制をより厳しくする必要があります。私は、エージェント ワークフローに欠けているレイヤーとして Cruxible を構築しました。これは、エージェントと人間が CLI または MCP サーバーを使用して操作する管理された状態エンジンであり、YAML で構成されたオントロジーによってバインドされます。すべての書き込みまたは変更は、完全な出所とともに帰属および受領され、管理された主張には、スキップできない証拠に裏付けられた提案とレビューのフローが必要となる場合があります。書き込みポリシーとガードは、モデル外部の決定論的なチョークポイントによって強制されます。 - 決定論的に構築できるすべての状態は、構造化 (CSV) または非構造化 (ドキュメント) のシード データから決定論的に構築されます。判断が必要な場合、Cruxible には、すべての変更または提案に対する完全な帰属確認を含む厳格なガバナンス レビュー キューが含まれています。承認された提案のみが州に認可されます。 - すべてのクエリは決定的であり、LLM の外部で実行されます。エージェントはグラフを直接走査することも、設定された名前付きクエリを使用して繰り返し読み取りを行うこともできます。クエリは、応答を生成するために使用されたパスの詳細を示す構造化された JSON 形式のレシートを返します。 - 強制は両方向に実行されます。ガードはチョークポイントでステートへの無効な書き込みを拒否し、ゲートはステートが同意するまで外部アクション (CI を介したマージやデプロイなど) を保持します。 - 状態モデルは実行可能です。プログラムで状態を読み取り、それに対して必要なロジックを実行できます。出所によって完全な監査可能性が保証されますが、解決された状態のダウンストリーム読み取りではソース アーティファクトを再解析する必要はありません。 - 読み取りは再現可能です。ワークフロー プラン、プロバイダー コード、アーティファクト、構成はロック固定されており、スナップショットは正確な情報を保存します。

州。レシートは常に書き込みを完全に説明するため、md ファイルの git 履歴を調べる必要はありません。 - YAML オントロジーは、状態が常に代表的であることを保証するために、コードのように進化できます。リポジトリはここにあります: https://github.com/cruxible-ai/cruxible または pip install cruxible でインストールします。これは Apache-2.0、内部で SQLite、Python デーモンを介して 100% ローカルです。 GitHub リポジトリにあるドメインの完全なウォークスルーとガイド。私は、たとえそれがエージェントにとって何が真実であるかについて、あなたやあなたのチームがどのように一貫性を維持しているかということであっても、どんな種類のフィードバックでも非常に重視しています。貢献は大歓迎です!

記事本文:
GitHub - cruxible-ai/cruxible: AI エージェント用のハードで検証可能な状態レイヤー · GitHub
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
読み込み中にエラーが発生しました。このページをリロードしてください。
るつぼai
/
るつぼ
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
246 コミット 246 コミット .github/ workflows .github/ workflows 資産 アセット デプロイ デプロイ ドキュメント ドキュメント サンプル/ wiki-import-demo サンプル/ wiki-インポート-デモ キット キット パッケージ/ cruxible-client パッケージ/ cruxible-client パッケージ化/ cruxible-core-stub パッケージ化/ cruxible-core-stub スクリプト スクリプト スキル スキル src/ cruxible_core src/ cruxible_core テスト テスト .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_DATA.md THIRD_PARTY_DATA.md context7.json context7.json Glama.json Glama.json mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml server.json server.json uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
るつぼ.ai ·
クイックスタート ·
ドキュメント ·
キット・
スキル
Cruxible は、オントロジーに基づいた AI エージェント向けの管理された状態エンジンです。
必要なクレーム用の汎用エージェント メモリの代替
解決され、施行され、監査可能です。
すべての人間とエージェントは、実際のドメインの 1 つの型付きモデル、つまりすべてのクレームを共有します。
宣言されたルールに照らして検証され、審査を通じて判断が下されます。
そして、計算されたすべての答えにはその受領書が含まれます。これをハード状態と呼びます。
実際には、これは CLI と MCP サーバーを備えた Python デーモンであり、
SQLite の型付きグラフ。ドメインのオントロジーを宣言します。
Terraform のような YAML 構成 (エンティティ タイプ、関係、書き込みルール、
クエリ);人間とエージェントは管理された書き込みを通じてグラフを構築します。
名前付きクエリが横断する: 爆発範囲、ダウンストリームの影響、マルチホップ
質問の類似性検索は追跡できません。
pip install cruxible —
クイックスタート
行きます

最初のクエリにインストールします。以下の「Get Started」を実行すると、
シードされたデモワールドはトークンなしで約 3 分で作成できます。
ドメインをモデル化します。キットから開始する 、または
オーサリングスキルを身につける
あなたのデータをエージェントに送信します。エージェントは YAML オントロジーの草案を作成します。あなたはそれを確認します。
それが提案するもの。
ソースを固定します。あなたの真実を示すエクスポート、表、文書
コンテンツハッシュされたアーティファクトとして登録されます。主張がそれらに引用されています。
厳しい事実を理解してください。決定論的なワークフローはソース行と一致します
入力されたエンティティとエッジに変換され、コミットする前にプレビューされます。
判決を提案する。管理対象と宣言したクレーム タイプは適用できません
ワークフローではなく、直接書き込むこと。彼らはプロポーズされる
それらと一致する証拠を持ったレビューグループ。
オプトインした場合に限り、レビューとミントが行われます。 ガバナンスされたクレームが適用されます。
レビュー: あなたが宣言した信頼ルールに基づく人間またはエージェントが承認します。
または拒否します。それ以外のものはすべて、書かれた瞬間に生きています。
承認または修正は、クエリ結果から直接得られる 1 つの動詞です。
質問し、その答えに基づいて行動します。エージェントは MCP または CLI を介して動作します。
あなたが彼らに与える許可レベル。クエリは応答とともに領収書を返します。
警備員はルールに違反する書き込みを拒否し、ゲートは屋外で保管されます
状態が同意するまでアクション (マージ、デプロイ) を実行します。
構成が宣言されると、すべてのミューテーションと計算されたクエリは
受信した CLI または MCP 動詞:
以下のチュートリアルでは、このループを次のように実行します
実際のコマンド。
エンジン内にはLLMはありません。見たところにのみ判断が入る
(モデリング、提案、レビュー);エンジン自体が行うすべてのこと —
検証、拒否、記録、回答 - 決定的であり、どのエージェントでも機能します
またはハーネスを使用すると、所有する SQLite ファイルに到達します。
仕事は、あなたが真実であると判断したことの永続的な記録を形成します。
高価な問題が発生したとき、どの資産が公開されるのでしょうか?何が壊れるのか
する

ダウンストリーム？この権限は依然として良い法律なのでしょうか？ — 答えは計算されます
そのレコードは、生のコンテキストの山から各セッションを再抽出したものではありません。
pip インストールるつぼ
独自のドメインをモデル化する: エージェントにオーサリング スキルを教えます。
スキル/
( prepare-data → create-state → review-state ) とエクスポート
(ウィキからステートへは、既存のチーム ウィキまたは Obsidian Vault を変換します)、または
モデリング状態から開始
および構成テンプレート。
または、デモを実行します。シードされたサプライチェーンの世界、約 3 分、トークンなし
(サンドボックスは組み込みのオペレーター ID に属性を書き込みます):
# シェル 1 — ローカル サンドボックス デーモン
CRUXIBLE_SERVER_STATE_DIR= " $HOME /.cruxible/sandbox " cruxible サーバーの起動
# シェル 2 — キット バンドルがリリースからフェッチされ、ダイジェスト検証されます
# (agent-operationはオプションのagent-opsレイヤーです。ドメインのみでも機能します)
cruxible --server-url http://127.0.0.1:8100 init --kit Agent-operation --kit Supply-chain-blast-radius
cruxible context connect --server-url http://127.0.0.1:8100 --instance-id < インスタンス ID >
# 決定的取り込み: プレビューしてからコミット
cruxible run --workflow build_seed_state && cruxible apply --workflow build_seed_state --from-last-preview
cruxible run --workflow ingest_incidents && cruxible apply --workflow ingest_incidents --from-last-preview
# インシデント フィードはインパクト エッジのみを提案できます。記録上の判断はあなたのものです
cruxible 提案 --workflow提案_インシデント_影響_サプライヤー
Crucible グループ リスト --status pending_review
cruxible グループの解決 --group < GRP-id > --action accept \
--rationale " サプライヤーの地理的条件に照らして確認 " --expected-pending-version 1
# 件の回答を、あなたが認めたエッジを通じて受信しました
重要なクエリ実行 open_incident_impacts --json
重要なクエリの実行 Incident_impacted_suppliers --param Incident_id=INC-TW-RAIL-2026-

07 --json
エージェントが参加すると、ID がオンになります。 CRUXIBLE_SERVER_AUTH=true で再起動します。
ブートストラップ資格情報を要求し、各エージェントごとに独自のトークンを作成します。
書き込みが帰属されます。詳細、権限層、強化:
クイックスタート ·
実行時認証とエージェントの役割。
LLM はポリシーを解釈できます。 Crucible はそれを強制します。解釈は
サンプリング: LLM のコンプライアンスは確率であり、呼び出しごとに新たに抽出されます。
モデルがどれほど優れていても、メモリに保存されたポリシーは変わりません
それ;記憶されたテキストは依然として解釈され、強制されません。宣言されたルール
in config は不変式であり、ライターがスキップできないチョークポイントで実行されます。
2 つの強制方向、1 つの目的 - 共有状態を真実に保つ
スタックの残りの部分が信頼できるレイヤー:
警備員は内側を向いています。受け入れられた書き込み境界を保護します。
状態: 宣言された条件が満たされた場合にのみ、クレームが入力または変更されます。
保留 — 行為者は許可されており、移行は合法であり、証拠
逆参照すると、凍結されたフィールドは凍結されたままになります。
ゲートは外側を向いています。 Cruxible の外側のアクションは、前の状態をチェックします。
重要なエクスポージャがオープンしている間はデプロイが保留され、ファイリングが行われます。
ステップは引用チェックを待ち、マージは承認を待ちます
レビュー。チェックは、アクションが存在する場所 (CI ジョブ、
内部プロトコル、git フック)。
ゲート :
マージレビュー:
種類: git-pre-push
エンティティタイプ : ReviewRequest
match_property :change_head
条件 : {ステータス: 承認済み}
アダプター: {ブランチパターン: refs/heads/main}
最初に出荷されたゲートの種類は、Git リポジトリをその状態に保持します: 1 行
.git/hooks/pre-push にあり、すべてのブランチがマージされない限りプッシュは拒否されます
main への移行は承認されたレビューによって固定されています - フェールクローズされ、必須です
同じチェックがプッシュでスキップできない場所に配線されると (CI、保護されています)
インフラストラクチャ）。このリポジター

yはその上を走ります：ゲートは私たちのものを拒否しました
0.2.2のリリースはレビュー記録が修正されるまでプッシュされる状態です。私たち
フックではなく状態を修正しました。
判断はモデルに委ねられます。ルールはそれなしでも実行されます。
管理対象ドメインの概要
このセクションの行き着く先: 検討された判決 - この事件は影響を与える
このサプライヤー — そしてそこから、クエリは再帰的な部品表をたどります。
下流にある 5 つのタイプ指定されたホップを通過するすべての公開貨物に名前を付け、領収書を添付します。
その真実がどのように構築されるかは次のとおりです。サプライチェーンの最小限のスライス
キット構成で作成されたオントロジー:
エンティティタイプ:
サプライヤー :
プロパティ:
サプライヤー ID : { タイプ: 文字列、プライマリキー: true }
名前: { タイプ: 文字列、インデックス付き: true }
Primary_geography : { タイプ: 文字列、オプション: true }
コンポーネント:
プロパティ:
コンポーネント ID : { タイプ: 文字列、プライマリキー: true }
名前: { タイプ: 文字列、インデックス付き: true }
重要度: { タイプ: 文字列、オプション: true、enum_ref: 重要度 }
事件 :
プロパティ:
Incident_id : { タイプ: 文字列、プライマリキー: true }
タイトル : { タイプ: 文字列、インデックス付き: true }
重大度: { タイプ: 文字列、オプション: true、enum_ref: Incident_severity }
関係 :
- 名前：supplies_supplies_component
提供元: サプライヤー
to : コンポーネント
# 統治された判断: インシデントはサプライヤーに重大な影響を与えます。
- 名前 : Incident_impacts_supplier
より: 事件
宛先: サプライヤー
名前付きクエリ :
# 爆発範囲: インシデントから、影響を受けたサプライヤーを横断して、
彼らが提供するコンポーネントは # 個です。
コンポーネント_エクスポーズド_バイ_インシデント :
モード: トラバーサル
エントリポイント : インシデント
戻り値: コンポーネント
トラバース :
- 関係 : Incident_impacts_supplier
方向: 発信
- 関係 : サプライヤー_サプライヤー_コンポーネント
方向: 発信
オントロジーは構成の一部にすぎません。同じファイルで列挙型が宣言されます。
語彙、ガード、プロポーザルルーティング、ワークフロー

、プロバイダーなど、
ドメインのモデル、ルール、およびプロシージャは、バージョン管理された 1 つの状態として一緒に出荷されます。
組み立て可能なキット。
この状態を手動で入力する人はいません。この状態は経路を通って入力されます。
config が宣言し、状態が異なると扱いも異なります。
確かな事実は決定論的な取り込みです。 BOM ワークフローはエクスポートを
アーティファクトを抽出し、その行をサプライヤー、コンポーネント、サプライエッジと照合します。
コミットする前にプレビューされる:
cruxible run --workflow ingest_bom --input-file ./exports/bom-2026-07.csv # プレビュー
cruxible apply --workflow ingest_bom --from-last-preview # commit
Incident_impacts_supplier は判断を求めるものであるため、次のように管理されます。
ライブ直接書き込みは拒否されます — CLI、MCP、バッチ、どの権限層でも
(直接書き込みでは、せいぜいレビュー用にエッジをステージングすることができます)。それ
構成が宣言する管理された動詞を介してのみ入力されます。
提案とレビューの領域です。インシデントフィードのワークフローレコード
事件自体は厳然たる事実だが、その影響は薄れる可能性がある
提案するだけです。それらの候補者は検討グループに集まり、それぞれが
それと一致する信号と証拠:
cruxible 提案 --workflow提案_インシデント_影響_サプライヤー --入力ファイル ./exports/incidents.json
判断自体は人間かエージェントに委ねられますw

[切り捨てられた]

## Original Extract

Hard, verifiable state layer for AI agents. Contribute to cruxible-ai/cruxible development by creating an account on GitHub.

General-purpose text memory works well to represent information that can be plainly stated, rarely changes, and only needs to be recalled, but it leaves several structural problems unresolved: 1) Generic memory does not establish settled state, it returns context. Text-based memory can preserve and present information, but a claim added to context is reinterpreted by a probabilistic LLM every time it runs. I found myself constantly asking my agents to re-check the validity of what was in my markdown wikis, especially for knowledge that has a natural lifecycle. Additionally, there's no uniform way of establishing provenance or attribution for a claim in text - relying on git commit messages doesn't feel like a good standard. 2) An LLM can't do deterministic queries for relational information on its own, especially when relationships have to be constructed across a corpus. This really matters across sessions and for multi-agent workflows - even if one agent or session gets it right, if your organization's truth is not plainly and deterministically traversable, it creates a tower of Babel situation where each agent or session is operating from its own internal representation. 3) An LLM can't be trusted to enforce its own write-time boundary. A policy represented as text can only be interpreted by the model as context - better models have a better probability of following them, but an LLM by itself, no matter how good, can never guarantee that an invalid transition is refused, or that an author of a claim can't approve it when you want an independent review. For truth that matters, there has to be stricter write-time enforcement. I built Cruxible to be that missing layer in agentic workflows. It's a governed state engine that agents and humans operate using a CLI or MCP server, bound by an ontology that is configured in YAML. Every write or mutation is attributed and receipted with full provenance, and governed claims can require evidence-backed proposal and review flows that cannot be skipped. Write policies and guards are enforced by a deterministic chokepoint outside the model. - All state that can be built deterministically is built deterministically from your seeded data, which can be structured (CSVs) or unstructured (documents). When judgment is needed, Cruxible includes strict governance review queues with full attribution receipts for all changes or proposals. Only proposals that are approved get minted into state. - Every query is deterministic and executes outside the LLM. Agents can traverse the graph directly or use configured named queries for recurring reads. Queries return receipts in structured JSON that detail the paths used to produce their answers. - Enforcement runs in both directions: guards refuse invalid writes to state at the chokepoint, and gates hold outside actions (e.g. a merge or a deploy through CI) until state agrees. - The state model is executable: you can read the state programmatically and do whatever logic you want with it. Downstream reads of settled state never have to re-parse source artifacts, although provenance ensures full auditability. - Reads are reproducible: workflow plans, provider code, artifacts, and configuration are lock-pinned, and snapshots preserve exact state. Receipts always explain a write fully, so no more going through the git history of an md file. - The YAML ontology can evolve like code to ensure that state is always representative. Repo is here: https://github.com/cruxible-ai/cruxible or install with pip install cruxible It's Apache-2.0, SQLite under the hood, and 100% local through a Python daemon. Full walkthroughs and guides for domains on the GitHub repo. I'd really value feedback of any kind, even if it's just how you and/or your team are maintaining consistency about what's true for your agents to use. Contributions welcome!

GitHub - cruxible-ai/cruxible: Hard, verifiable state layer for AI agents · GitHub
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
cruxible-ai
/
cruxible
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
246 Commits 246 Commits .github/ workflows .github/ workflows assets assets deploy deploy docs docs examples/ wiki-import-demo examples/ wiki-import-demo kits kits packages/ cruxible-client packages/ cruxible-client packaging/ cruxible-core-stub packaging/ cruxible-core-stub scripts scripts skills skills src/ cruxible_core src/ cruxible_core tests tests .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_DATA.md THIRD_PARTY_DATA.md context7.json context7.json glama.json glama.json mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml server.json server.json uv.lock uv.lock View all files Repository files navigation
cruxible.ai ·
quickstart ·
docs ·
kits ·
skills
Cruxible is a governed state engine for AI agents — an ontology-backed
alternative to general-purpose agent memory for claims that must be
settled, enforced, and auditable.
All humans and agents share one typed model of a real domain: every claim
is validated against declared rules, judgment calls route through review,
and every computed answer carries its receipt. We call it hard state .
In practice it's a Python daemon with a CLI and an MCP server, storing a
typed graph in SQLite. You declare your domain's ontology in a
Terraform-like YAML config (entity types, relationships, write rules,
queries); humans and agents build the graph through governed writes, and
named queries traverse it: blast radius, downstream impact, the multi-hop
questions similarity retrieval can't follow.
pip install cruxible — the
Quickstart
goes install to first query; Get Started below runs a
seeded demo world in ~3 minutes, no tokens.
Model the domain. Start from a kit , or
hand the authoring skills
to your agent with your data: it drafts the YAML ontology, you review
what it proposes.
Pin the sources. The exports, tables, and documents your truth
comes from register as content-hashed artifacts; claims cite into them.
Ingest the hard facts. Deterministic workflows match source rows
into typed entities and edges, previewed before they commit.
Propose the judgment calls. Claim types you declared governed can't
be written directly, not even by a workflow; they're proposed into
review groups carrying the evidence that matched them.
Review and mint, only where you opted in. Governed claims land in
review: a human, or an agent under trust rules you declared, approves
or rejects. Everything else is live the moment it's written, and
approving or correcting it is one verb, straight from a query result.
Ask, and act on the answer. Agents work through MCP or the CLI at
the permission tier you give them; queries return answers with receipts;
guards refuse writes that break the rules, and gates hold outside
actions (a merge, a deploy) until state agrees.
Once the config is declared, every mutation and computed query is a
receipted CLI or MCP verb: the
walkthrough below runs this loop as
the actual commands.
There is no LLM inside the engine. Judgment enters only where you saw it
(modeling, proposals, review); everything the engine itself does —
validate, refuse, record, answer — is deterministic, works with any agent
or harness, and lands in a SQLite file you own.
Work compounds into a durable record of what you've determined to be true.
When the expensive question arrives — which assets are exposed? what breaks
downstream? is this authority still good law? — the answer is computed over
that record, not re-derived each session from a pile of raw context.
pip install cruxible
Model your own domain : hand your agent the authoring skills in
skills/
( prepare-data → create-state → review-state ) with your exports
( wiki-to-state converts an existing team wiki or Obsidian vault), or
start from Modeling State
and the config template .
Or run the demo — a seeded supply-chain world, ~3 minutes, no tokens
(sandbox writes attribute to a built-in operator identity):
# shell 1 — local sandbox daemon
CRUXIBLE_SERVER_STATE_DIR= " $HOME /.cruxible/sandbox " cruxible server start
# shell 2 — kit bundles are fetched from the release and digest-verified
# (agent-operation is the optional agent-ops layer; domain-only works too)
cruxible --server-url http://127.0.0.1:8100 init --kit agent-operation --kit supply-chain-blast-radius
cruxible context connect --server-url http://127.0.0.1:8100 --instance-id < instance-id >
# deterministic ingest: preview, then commit
cruxible run --workflow build_seed_state && cruxible apply --workflow build_seed_state --from-last-preview
cruxible run --workflow ingest_incidents && cruxible apply --workflow ingest_incidents --from-last-preview
# the incident feed can only PROPOSE impact edges; the judgment is yours, on the record
cruxible propose --workflow propose_incident_impacts_supplier
cruxible group list --status pending_review
cruxible group resolve --group < GRP-id > --action approve \
--rationale " Confirmed against supplier geography " --expected-pending-version 1
# receipted answers through the edges you just admitted
cruxible query run open_incident_impacts --json
cruxible query run incident_impacted_suppliers --param incident_id=INC-TW-RAIL-2026-07 --json
When agents join, identity turns on: restart with CRUXIBLE_SERVER_AUTH=true ,
claim the bootstrap credential, and mint each agent its own token — every
write is attributed. Details, permission tiers, and hardening:
Quickstart ·
Runtime Auth And Agent Roles .
An LLM can interpret a policy. Cruxible enforces it. Interpretation is
sampling: an LLM's compliance is a probability, drawn fresh on every call,
however good the model — and storing the policy in memory doesn't change
that; remembered text is still interpreted, not enforced. A rule declared
in config is an invariant, run at a chokepoint no writer can skip.
Two enforcement directions, one purpose — keeping the shared state a truth
layer the rest of your stack can trust:
Guards face inward. They protect the write boundary of accepted
state: a claim enters or changes it only when the declared conditions
hold — the actor is authorized, the transition is legal, the evidence
dereferences, frozen fields stay frozen.
Gates face outward. An action outside Cruxible checks state before
it proceeds: a deploy holds while critical exposures are open, a filing
step waits for the citation check, a merge waits for its approved
review. The check wires in wherever the action lives (a CI job, an
internal protocol, a git hook).
gates :
merge-review :
kind : git-pre-push
entity_type : ReviewRequest
match_property : change_head
condition : {status: approved}
adapter : {branch_pattern: refs/heads/main}
The first shipped gate kind holds a git repository to its state: one line
in .git/hooks/pre-push , and a push is refused unless every branch merged
into main is pinned by an approved review — fail closed, and mandatory
once the same check is wired where pushes can't skip it (CI, protected
infrastructure). This repository runs on it: the gate refused our own
0.2.2 release push until the review record was corrected in state. We
fixed the state, not the hook.
Judgment stays with the model; the rules run without it.
What A Governed Domain Looks Like
Where this section ends up: one reviewed judgment — this incident impacts
this supplier — and from it, a query walks a recursive bill of materials
to name every exposed shipment, five typed hops downstream, with a receipt.
Here is how that truth gets built. A minimal slice of a supply-chain
ontology, as authored in a kit config:
entity_types :
Supplier :
properties :
supplier_id : { type: string, primary_key: true }
name : { type: string, indexed: true }
primary_geography : { type: string, optional: true }
Component :
properties :
component_id : { type: string, primary_key: true }
name : { type: string, indexed: true }
criticality : { type: string, optional: true, enum_ref: criticality }
Incident :
properties :
incident_id : { type: string, primary_key: true }
title : { type: string, indexed: true }
severity : { type: string, optional: true, enum_ref: incident_severity }
relationships :
- name : supplier_supplies_component
from : Supplier
to : Component
# Governed judgment: an incident materially impacts a supplier.
- name : incident_impacts_supplier
from : Incident
to : Supplier
named_queries :
# Blast radius: from an incident, traverse impacted suppliers to the
# components they supply.
components_exposed_by_incident :
mode : traversal
entry_point : Incident
returns : Component
traversal :
- relationship : incident_impacts_supplier
direction : outgoing
- relationship : supplier_supplies_component
direction : outgoing
The ontology is only part of the config: the same file declares the enum
vocabularies, guards, proposal routing, workflows, and providers, so a
domain's model, rules, and procedures ship together as one versioned,
composable kit.
Nobody types this state in by hand: it enters through the pathways the
config declares, and different state earns different treatment.
Hard facts are deterministic ingest. A BOM workflow pins the export as an
artifact and matches its rows into suppliers, components, and supply edges,
previewed before it commits:
cruxible run --workflow ingest_bom --input-file ./exports/bom-2026-07.csv # preview
cruxible apply --workflow ingest_bom --from-last-preview # commit
incident_impacts_supplier is a judgment call, so it is governed: every
live direct write is refused — CLI, MCP, batch, at any permission tier
(a direct write can at most stage the edge for review). It
enters only through the governed verbs the config declares, and in this
domain that is proposal and review. The incident feed's workflow records
the incidents themselves as hard facts, but the impact edges it can
only propose . Those candidates land in a review group, each carrying the
signals and evidence that matched it:
cruxible propose --workflow propose_incident_impacts_supplier --input-file ./exports/incidents.json
The judgment itself stays with a human, or with an agent w

[truncated]
