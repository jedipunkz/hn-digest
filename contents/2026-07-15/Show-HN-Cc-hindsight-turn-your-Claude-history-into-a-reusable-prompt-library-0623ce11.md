---
source: "https://github.com/adityaarunsinghal/cc-hindsight"
hn_url: "https://news.ycombinator.com/item?id=48921343"
title: "Show HN: Cc-hindsight – turn your Claude history into a reusable prompt library"
article_title: "GitHub - adityaarunsinghal/cc-hindsight: Mine your Claude Code session history into a oneshot prompt library and CLAUDE.md preferences. The prompt you'd write if you knew then what you know now. · GitHub"
author: "dramebaaz"
captured_at: "2026-07-15T15:20:15Z"
capture_tool: "hn-digest"
hn_id: 48921343
score: 2
comments: 3
posted_at: "2026-07-15T14:26:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Cc-hindsight – turn your Claude history into a reusable prompt library

- HN: [48921343](https://news.ycombinator.com/item?id=48921343)
- Source: [github.com](https://github.com/adityaarunsinghal/cc-hindsight)
- Score: 2
- Comments: 3
- Posted: 2026-07-15T14:26:02Z

## Translation

タイトル: HN を表示: Cc-hindsight – クロード履歴を再利用可能なプロンプト ライブラリに変える
記事のタイトル: GitHub - adityaarunsinghal/cc-hindsight: クロード コードのセッション履歴をワンショット プロンプト ライブラリと CLAUDE.md 設定にマイニングします。今知っていることを当時知っていたら書くだろうプロンプト。 · GitHub
説明: クロード コードのセッション履歴をワンショット プロンプト ライブラリと CLAUDE.md 設定にマイニングします。今知っていることを当時知っていたら書くだろうプロンプト。 - adityaarunsinghal/cc-hindsight
HN テキスト: これを作成したのは、過去のセッションから学び、最初のプロンプトからクロード コードに合わせて調整したかったからです。私はセッションのステアリングに多くの時間を費やして、
同じ好みで、以前に答えたような質問に答えます。によって
セッションの終了時に、その種の作業用のライブラリ エントリをすぐに作成できるようになりました。私も、これらをすぐにスキルに変える手伝いをしてみようと思うかもしれません。

記事本文:
GitHub - adityaarunsinghal/cc-hindsight: クロード コードのセッション履歴をワンショット プロンプト ライブラリと CLAUDE.md 設定にマイニングします。今知っていることを当時知っていたら書くだろうプロンプト。 · GitHub
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
アディティアーンシンハル
/
CC-後知恵
パブ

リック
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
adityaarunsinghal/cc-hindsight
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
35 コミット 35 コミット .github .github scripts scripts src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md biome.json biome.json npm-shrinkwrap.json npm-shrinkwrap.json package.json package.json tsconfig.json tsconfig.json tsdown.config.ts tsdown.config.ts vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
今知っていることを当時知っていたら書くだろうプロンプト。
npx cc-後知恵
cc-hindsight は、クロード コードのセッション履歴をマイニングして、使用できる 2 つのものにします。
明日の朝:
ワンショット プロンプト ライブラリ。あなたが実行した実際のタスクごとに、現実的な
理想的な最初のプロンプト: t=0 の時点で知っていて望んでいたものの、言わなかったすべてのこと、
あなた自身の声で書かれています。
洗練された設定セット。あなたがエージェントに繰り返し言い続けていること
(「行動する前に診断する」、「バージョンを固定する」、「簡潔にする」) として表面化しました。
CLAUDE.md ブロックを貼り付ける準備ができているため、今後のセッションはすべて整列して開始されます。
最初のプロンプトにフロントロードするコンテキストが多いほど、より良い結果が得られます。
それでも、t=0 で過少指定し、セッションのステアリングを費やして、コースを修正します。
好みを言い換えたり、エージェントが尋ねるべきではなかった質問に答えたりします。
そして費用はセッションごとに繰り返されます。あなたの仕事の好みは安定した事実です
どのように働きたいかについて、後でプロジェクトを再導出して再入力する
事前に一度説明するのではなく、プロジェクトを作成し、今後のすべてのセッションで読み込む
最初のメッセージから。
あなたが言うべきだった証拠はすべてすでにディスク上にあります。クロード・コード
イブを救う

~/.claude/projects/ の下に JSONL としてセッションを保存します。たくさんのツールが読み取れます
そのデータはダッシュボードとトランスクリプトに使用されます。誰も歴史のループを閉じない
より良いプロンプトに戻ります。
前 (t=0 で実際に入力した内容):
ホームルーターでデバイスをブロックするスクリプトを手伝ってください
後 (そのセッションが実際にどのように進んだのかを cc-hindsight から抽出したもの):
デバイスのブロックから始めて、ホームルーターを自動化したいと考えています。 APIの取得
まず作業中にアクセスし、私に尋ねる代わりに認証情報を安全に保存してください
クリアに貼り付けます。ネットワークを変更するものはすべて安全でなければなりません
デフォルト: 実行内容を表示し、実行前に確認させます。それから
実際のデバイスでエンドツーエンドでそれを証明します。他の人がこれをどのように行ったかを調べてください
まず、再現できるように設定を書き留めておきます。
2 番目のプロンプトは t=0 で認識可能でした。まだ書いていなかっただけです。
1 つのコマンドでライブラリ全体がビルドされ、完了すると表示されます。
npx cc-hindsight distill # 最初にエクスポートを提案し、次にダイジェスト → クラスター →
# author (クロード呼び出しの前に質問)、その後表示
# あなたは完成したライブラリです
各ステージを自分で運転したいですか?パイプラインは 3 つの動詞だけです。
npx cc-hindsight # 1. scan: プロジェクトのインベントリを作成します (読み取り専用)
npx cc-hindsight エクスポート # 2. エクスポート: セッションごとの人間のみのマークダウン
npx cc-hindsight 蒸留 # 3. 蒸留: ダイジェスト → クラスター → 著者 (最初に質問)
次に、結果を参照して厳選します。
npx cc-hindsight list # あなたのライブラリ (✎ 編集済み / ▲▼ 評価バッジ)
npx cc-hindsight show < slug > # ワンショットを読み取る
npx cc-hindsight copy < slug > # → クリップボード、新しいセッションに貼り付け
npx cc-hindsight edit < slug > # $EDITOR で開きます。編集内容は保護されています
npx cc-hindsight rate < slug > up # 判定を記録する (up|down)
npx cc-hindsight prune # 孤立したエントリを削除します (最初に質問します)
npx

cc-hindsight 設定 # → CLAUDE.md ブロック
npx cc-hindsight ステータス # パイプライン ファネル + 孤立/スキップ フラグ
macOS または Linux ではノード 22 以上が必要です (Windows はテストされておらず、現在
サポートされていません）。スキャンとエクスポートは完全に決定的です: LLM なし、なし
ネットワーク。 distill は、既存の claude CLI を必ず使用します。
尋ねます。
会話履歴全体を読み取るツールにとって、監査可能性は重要です。
製品:
地元限定。すべてがディスクから読み取られ、ディスクに書き込まれます
( ~/.cc-hindsight )。サーバーもアカウントもテレメトリもありません。唯一のネットワーク
call は、すでに実行していることを実行する独自の claude CLI です。
3 つのランタイム依存関係。都会
(依存性ゼロの CLI フレームワーク)、zod (スキーマ
検証)、および @clack/prompts
(進行スピナー、ボンネットの下に 5 つのマイクロパッケージ)。 3つとも、
正確に固定され、推移ツリー全体がロックされた状態で出荷されます。
npm-shrinkwrap.json 。データ パス全体が 1 回で監査できるほど小さい
座っている。
同意ゲート型 LLM の使用。 distill は正確な呼び出し回数を示します。
サブスクリプション/クレジットで何かが実行される前に、[y/N] 待機します。
--dry-run は完全なプランを無料で表示します。拒否すると何もせずに終了します
何かを呼び出します (終了コード 2)。
単純なステートメント: エクスポートには生のプロンプトがすべて含まれます。
これまでにセッションに貼り付けた機密情報 (キー、内部名、その名前など)
怒りのメッセージ）。これらはあなたの許可を得てあなたのホームディレクトリに存在し、
どこにも発送するものはありません。 --redact オプションがロードマップにあります。
フローチャート TD
A["~/.claude/projects/**/*.jsonl"] -->|"スキャン (読み取り専用)"| B[「在庫」]
A -->|"エクスポート (決定的)"| C["exports/*.md + マニフェスト.json"]
A -->|"照応 + 結果の証拠"| D["アナフォラ.json + 結果.json"]
C --> E
D --> E
サブグラフ蒸留 ["蒸留: オプトイン、同意ゲート、再開可能"]
E["ダイジェスト: 1 クロード・カル

l / セッション"] --> F["クラスター: 1 コール"]
F --> G["作成者: 1 コール/タスク"]
終わり
G --> H["ライブラリ/<スラッグ>/ワンショット + 来歴"]
H --> I["リスト/表示/コピー/ステータス"]
H --> J["設定 → CLAUDE.md ブロック"]
読み込み中
抽出層は忠実性契約です。監査された 11 のルールによって内容が決定されます。
人間による入力としてカウントされます (エージェントがビジー状態の間に入力された添付ファイルは復元されました)
スラッシュコマンド、[決定] オプション選択のセリフ、[画像貼り付け]
マーカー)、ルールごとに回帰フィクスチャを使用します。フォーク/再開の重複は次のとおりです。
グローバルに重複排除されます。 「はい」や「オプション 2」などの短い応答は先行詞を取得します
あなたが何を承認したかを作成者ステージに知らせるために添付されます。セッションには、
結果を分類するため、何も成功しなかったタスクはスキップされます
失敗パスを再現する自信のあるプロンプトになるのではなく。
作成されたワンショットは「t=0 で認識可能」テストに合格します。つまり、インテントをフロントロードします。
制約、好み、品質バーを考慮しながら、事実のみを除外します
セッション中に発見されました (パス、根本原因、エラー メッセージ)。そして彼らは尊敬します
努力の予算、やる気のある人間が行うであろう長さ程度にとどまる
700 語の仕様を膨らますのではなく、実際に入力します。
各作業単位の後に各蒸留ステージのチェックポイントをディスクに記録します: Ctrl-C は失われます
何もせず、再開を再実行し、--fresh を意図的にリセットし、ステータス フラグを設定します。
再クラスタリングによって孤立したライブラリ エントリ。
ツール
~/.claude を読み取ります
出力
方向
非難
✓
使用量/コスト表
逆方向 (メトリクス)
鼻をすする
✓
分析ダッシュボード
後ろ向き（行動）
クロードコードの転写
✓
共有可能な HTML トランスクリプト
逆方向（記録）
cctrace / エクスポーター
✓
マークダウン/XML エクスポート
後方 (アーカイブ)
CC-後知恵
✓
ワンショット プロンプト ライブラリ + CLAUDE.md 設定
前に進む (次のセッションがより良くなる)
最も近い親戚は、サイモン・ウィリソンのクロード・コード・トランスクリプトです: 同じ l

ローカルファースト
価値観、転写物が過小評価されている成果物であるという考えと同じです。違い
視線の方向です。セッションを (他の人にとっては後方に) レンダリングします。
cc-hindsight はそれらを抽出します (あなたに向けて)。
私のデータは私のマシンから離れますか?
いいえ。決定論的なコマンドがネットワークに影響することはありません。蒸留管
ローカルにインストールされた claude CLI にコンテンツを追加すると、同じことが起こります
クロードコードを使用するときだけで、それ以外は何も使用しません。
蒸留にかかる費用はいくらですか?
クロード CLI の費用に関係なく、既存のサブスクリプションで実行されます。
API クレジット。同意プロンプトには、正確な呼び出し回数が事前に示されます。
(ダイジェストされたセッションごとに 1 つの呼び出し、クラスターへの呼び出しが 1 つ、作成されたタスクごとに 1 つ)、および
--dry-run は、何も実行せずに同じ計画を出力します。チェックポイントの意味は、
中断された実行はお金の無駄ではありません。入力予算に対してセッションが大きすぎます
(デフォルトは 400k 文字、 --input-budget ) は、何も実行される前に拒否されます。
費やした。 --truncate=extreme を指定して再実行してそれらをカットします (
出所）、または範囲を狭めます。ダイジェスト呼び出しと作成者呼び出しは、一度に 3 つずつ実行されます。
デフォルト ( --concurrency ; 1 = シーケンシャル)。
サブスクリプションかAPI?
どちらか。 cc-hindsight は claude -p にシェルアウトします。ただし、CLI は
認証された場合は、通話の請求方法が決まります。 --モデルは通過します。
エクスポートにセッションが欠落しているのはなぜですか?
人間によるメッセージがゼロのセッションは自然にスキップされます。 --min-messages
薄いものをフィルタリングし、フォーク/再開のコピーを元のコピーに重複除去します。
セッション。重複排除はすべてのプロジェクトにわたってグローバルであることに注意してください。 --project を使用すると、
フォーク/レジュームによって複製されたメッセージは、たとえその最初のセッションに起因すると考えられます
そのセッションがフィルターの外にある場合。すべてを表示するには、export --verbose を実行します。
理由があってドロップした作品。本物の人間による入力が削除された場合、それは問題です。
バグ: 抽出忠実度レポートを提出してください。
すべてを再生できますか

最初から？
cc-hindsight distill --fresh チェックポイントをクリアし (確認後)、
再実行します。前世代の古いライブラリ エントリは孤立としてフラグが付けられます
状態で
既存のマーケットプレイス用の Claude Code プラグイン パッケージ化 (
別のマーケットプレイスを立ち上げることなく配布）
セッション中のスキルのバリアント (「このセッションを進めながら蒸留する」)
エージェント SDK オーケストレーションの調査
parentUuid 対応のアナフォラ (フォークされた分岐の正しい先行詞)
会話)
オプトインコミュニティのワンショットショーケース
npmインストール
npm test # vitest: フィクスチャのみ、~/.claude を読み込むことはありません
npm run lint # バイオーム
npm run typecheck # tsc --noEmit
npm run build # tsdown → dist/
CONTRIBUTING.md 、特に追加のガイドを参照してください。
未処理のトランスクリプト シェイプにヒットしたときのフィクスチャ。
アディティヤ シンハルによって建てられました。 CC の後知恵が表面化した場合
あなたが知らなかったパターン、ぜひ聞きたいです
それ。
Claude Code セッション履歴をワンショット プロンプト ライブラリと CLAUDE.md 設定にマイニングします。今知っていることを当時知っていたら書くだろうプロンプト。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
3
v1.0.2
最新の
2026 年 7 月 15 日
+ 2 リリース
パッケージ
0
読み込み中にエラーが発生しました。これをリロードしてください

[切り捨てられた]

## Original Extract

Mine your Claude Code session history into a oneshot prompt library and CLAUDE.md preferences. The prompt you'd write if you knew then what you know now. - adityaarunsinghal/cc-hindsight

I built this because I wanted to learn from my past session and align with Claude Code better right from the first prompt. I was spending a bunch of my session steering, restating the
same preferences, answering questions that I felt I've answered before. By
the end of a session I can now have a prompt library entry for that kind of work. I might try to help turn these into skills quickly as well.

GitHub - adityaarunsinghal/cc-hindsight: Mine your Claude Code session history into a oneshot prompt library and CLAUDE.md preferences. The prompt you'd write if you knew then what you know now. · GitHub
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
adityaarunsinghal
/
cc-hindsight
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
adityaarunsinghal/cc-hindsight
main Branches Tags Go to file Code Open more actions menu Folders and files
35 Commits 35 Commits .github .github scripts scripts src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md biome.json biome.json npm-shrinkwrap.json npm-shrinkwrap.json package.json package.json tsconfig.json tsconfig.json tsdown.config.ts tsdown.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
The prompt you'd write if you knew then what you know now.
npx cc-hindsight
cc-hindsight mines your Claude Code session history into two things you can use
tomorrow morning:
A oneshot prompt library. For each real task you've done, the realistic
ideal first prompt: everything you knew and wanted at t=0 but didn't say,
written in your own voice.
A distilled preference set. The things you keep re-telling your agent
("diagnose before acting", "pin the versions", "be terse"), surfaced as a
paste-ready CLAUDE.md block so every future session starts aligned.
You get better results the more context you front-load into your first prompt,
yet we underspecify at t=0 and spend the session steering: correcting course,
restating preferences, answering questions the agent shouldn't have had to ask.
And the cost repeats every session. Your working preferences are stable facts
about how you like to work, yet you re-derive and re-type them project after
project instead of stating them once, upfront, where every future session reads
them from the first message.
All the evidence of what you should have said is already on disk. Claude Code
saves every session as JSONL under ~/.claude/projects/ . Plenty of tools read
that data for dashboards and transcripts. Nobody closes the loop from history
back to better prompting.
Before (what you actually typed at t=0):
help me script blocking a device on my home router
After (what cc-hindsight distilled from how that session actually went):
I want to automate my home router, starting with blocking a device. Get API
access working first, and store any credentials securely instead of asking me
to paste them in the clear. Anything that changes the network should be safe
by default: show me what it'll do and let me confirm before it runs. Then
prove it end-to-end on a real device. Research how others have done this
first, and keep the setup written down so I can reproduce it.
The second prompt was knowable at t=0. You just hadn't written it yet.
One command builds the whole library and shows it to you when it's done:
npx cc-hindsight distill # offers to export first, then digest → cluster →
# author (asks before any claude call), then shows
# you the finished library
Prefer to drive each stage yourself? The pipeline is just three verbs:
npx cc-hindsight # 1. scan: inventory your projects (read-only)
npx cc-hindsight export # 2. export: human-only markdown per session
npx cc-hindsight distill # 3. distill: digest → cluster → author (asks first)
Then browse and curate the results:
npx cc-hindsight list # your library (✎ edited / ▲▼ rating badges)
npx cc-hindsight show < slug > # read a oneshot
npx cc-hindsight copy < slug > # → clipboard, paste into a fresh session
npx cc-hindsight edit < slug > # open in $EDITOR; your edits are protected
npx cc-hindsight rate < slug > up # record a verdict (up|down)
npx cc-hindsight prune # remove orphaned entries (asks first)
npx cc-hindsight preferences # → CLAUDE.md block
npx cc-hindsight status # pipeline funnel + orphan/skip flags
Requires Node ≥ 22 on macOS or Linux (Windows is untested and currently
unsupported). scan and export are fully deterministic: no LLM, no
network. distill uses the claude CLI you already have, and never without
asking.
For a tool that reads your entire conversation history, auditability is the
product:
Local-only. Everything is read from your disk and written to your disk
( ~/.cc-hindsight ). No server, no accounts, no telemetry. The only network
call is your own claude CLI doing what it already does.
Three runtime dependencies. citty
(zero-dependency CLI framework), zod (schema
validation), and @clack/prompts
(progress spinners, five micro-packages under the hood). All three are
exact-pinned, and the entire transitive tree ships locked via
npm-shrinkwrap.json . The whole data path is small enough to audit in one
sitting.
Consent-gated LLM use. distill states the exact invocation count and
waits for [y/N] before anything runs on your subscription/credits.
--dry-run shows the full plan for free. If you decline, it exits without
invoking anything (exit code 2).
Plain statement: exports contain your raw prompts, including anything
sensitive you ever pasted into a session (keys, internal names, that one
angry message). They live in your home directory with your permissions, and
nothing ships them anywhere. A --redact option is on the roadmap.
flowchart TD
A["~/.claude/projects/**/*.jsonl"] -->|"scan (read-only)"| B["inventory"]
A -->|"export (deterministic)"| C["exports/*.md + manifest.json"]
A -->|"anaphora + outcome evidence"| D["anaphora.json + outcomes.json"]
C --> E
D --> E
subgraph distill ["distill: opt-in, consent-gated, resumable"]
E["digest: 1 claude call / session"] --> F["cluster: 1 call"]
F --> G["author: 1 call / task"]
end
G --> H["library/&lt;slug&gt;/ oneshot + provenance"]
H --> I["list / show / copy / status"]
H --> J["preferences → CLAUDE.md block"]
Loading
The extraction layer is a fidelity contract: eleven audited rules decide what
counts as human input (attachments typed while the agent was busy, recovered
slash commands, [decision] lines from option picks, [image pasted]
markers), with a regression fixture per rule. Fork/resume duplicates are
deduped globally. Short replies like "yes" and "option 2" get their antecedent
attached so the author stage knows what you approved. Sessions carry an
outcome classification, so a task where nothing ever succeeded gets skipped
instead of becoming a confident prompt that reproduces a failure path.
Authored oneshots pass a "knowable at t=0" test : they front-load intent,
constraints, preferences, and quality bars, while leaving out facts you only
discovered mid-session (paths, root causes, error messages). And they respect
an effort budget , staying around the length a motivated human would
actually type rather than ballooning into a 700-word spec.
Every distill stage checkpoints to disk after each unit of work: Ctrl-C loses
nothing, re-running resumes, --fresh resets deliberately, and status flags
library entries orphaned by re-clustering.
Tool
Reads ~/.claude
Output
Direction
ccusage
✓
usage/cost tables
backward (metrics)
sniffly
✓
analytics dashboard
backward (behavior)
claude-code-transcripts
✓
shareable HTML transcripts
backward (record)
cctrace / exporters
✓
markdown/XML exports
backward (archive)
cc-hindsight
✓
oneshot prompt library + CLAUDE.md preferences
forward (better next session)
Closest kin is Simon Willison's claude-code-transcripts : same local-first
values, same belief that transcripts are undervalued artifacts. The difference
is the direction of gaze: it renders your sessions (backward, for others);
cc-hindsight distills them (forward, for you).
Does my data leave my machine?
No. The deterministic commands never touch the network. distill pipes
content to your locally installed claude CLI, the same thing that happens
when you use Claude Code, and nothing else.
What does distill cost?
Whatever your claude CLI costs you: it runs on your existing subscription or
API credits. The consent prompt states the exact invocation count up front
(one call per session digested, one to cluster, one per task authored), and
--dry-run prints the same plan without running anything. Checkpoints mean an
interrupted run isn't wasted money. Sessions too large for the input budget
(default 400k chars, --input-budget ) are refused before anything is
spent . Re-run with --truncate=extreme to cut them (recorded in
provenance) or narrow the scope. Digest and author calls run 3 at a time by
default ( --concurrency ; 1 = sequential).
Subscription or API?
Either. cc-hindsight shells out to claude -p ; however your CLI is
authenticated is how the calls are billed. --model passes through.
Why is a session missing from my exports?
Sessions with zero human messages are skipped naturally, --min-messages
filters thin ones, and fork/resume copies are deduped into their original
session. Note that dedupe is global across all projects: with --project , a
message duplicated by fork/resume is attributed to its earliest session even
when that session is outside the filter. Run export --verbose to see every
dropped piece with a reason. If genuine human input was dropped, that's a
bug: file an extraction-fidelity report.
Can I regenerate everything from scratch?
cc-hindsight distill --fresh clears the checkpoints (after confirming) and
re-runs. Old library entries from previous generations are flagged as orphans
in status .
Claude Code plugin packaging for existing marketplaces (for
distribution, without standing up a separate marketplace)
An in-session skill variant ("distill this session as we go")
Agent SDK orchestration exploration
parentUuid-aware anaphora (branch-correct antecedents on forked
conversations)
An opt-in community oneshot showcase
npm install
npm test # vitest: fixtures only, never reads your ~/.claude
npm run lint # biome
npm run typecheck # tsc --noEmit
npm run build # tsdown → dist/
See CONTRIBUTING.md , especially the guide to adding a
fixture when you hit an unhandled transcript shape.
Built by Aditya Singhal . If cc-hindsight surfaces
a pattern you didn't know you had, I'd love to hear about
it .
Mine your Claude Code session history into a oneshot prompt library and CLAUDE.md preferences. The prompt you'd write if you knew then what you know now.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
3
v1.0.2
Latest
Jul 15, 2026
+ 2 releases
Packages
0
There was an error while loading. Please reload this

[truncated]
