---
source: "https://github.com/asp67/llm-colosseum/tree/main"
hn_url: "https://news.ycombinator.com/item?id=48752981"
title: "LLM Colosseum – A zero-dependency browser RTS to test LLM tool calling"
article_title: "GitHub - asp67/llm-colosseum: Four LLMs battle each other in a browser RTS - a non-scientific sandbox eval for precise tool-calling and long-horizon strategic reasoning in an unfamiliar framework. · GitHub"
author: "osti67"
captured_at: "2026-07-01T21:31:25Z"
capture_tool: "hn-digest"
hn_id: 48752981
score: 1
comments: 0
posted_at: "2026-07-01T20:54:47Z"
tags:
  - hacker-news
  - translated
---

# LLM Colosseum – A zero-dependency browser RTS to test LLM tool calling

- HN: [48752981](https://news.ycombinator.com/item?id=48752981)
- Source: [github.com](https://github.com/asp67/llm-colosseum/tree/main)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T20:54:47Z

## Translation

タイトル: LLM Colosseum – LLM ツールの呼び出しをテストする依存関係のないブラウザ RTS
記事のタイトル: GitHub - asp67/llm-colosseum: ブラウザ RTS で 4 つの LLM が互いに戦います - 不慣れなフレームワークでの正確なツール呼び出しと長期的な戦略的推論のための非科学的なサンドボックス評価です。 · GitHub
説明: 4 つの LLM がブラウザ RTS で互いに戦います。これは、不慣れなフレームワークでの正確なツール呼び出しと長期的な戦略的推論のための非科学的なサンドボックス評価です。 - asp67/llm-コロシアム

記事本文:
GitHub - asp67/llm-colosseum: ブラウザ RTS で 4 つの LLM が互いに戦います。これは、不慣れなフレームワークでの正確なツール呼び出しと長期的な戦略的推論のための非科学的なサンドボックス評価です。 · GitHub
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
asp67
/
LLMコロシアム
公共
通知

イケーション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
64 コミット 64 コミット スクリーンショット スクリーンショット css css js js .gitignore .gitignore ライセンス ライセンス README.md README.md game-state-schema.json game-state-schema.json Index.html Index.html models.json models.json package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
古代において言語モデルが王冠をめぐって争う場所。
4 つの LLM。地図は1枚。勝者は 1 名です。
ブラウザベースの Age of Empires スタイルのリアルタイム ストラテジー ゲームで、競合する言語モデルが互いに対戦します。その間に、言語モデルを観察し、指導し、スコアを付けます。
簡単なツアー: ライブラリ内のモデルを配線してから、アリーナに接続します。
LLM Colosseum は、これまで訓練されていなかったタスク、つまり見たことのない小さな RTS 内でリアルタイムに経済と軍隊を運営するというタスクで言語モデルを互いに競わせるためのサンドボックス アリーナです。
これはリーダーボードでも、査読済みのベンチマークでもありません。また、統計的な厳密性を主張するものでもありません。これは実践的で非科学的なテストベッドであり、さまざまなモデルをなじみのないフレームワークに落とし込み、チャットではなく行動するように依頼したときにさまざまなモデルがどのように動作するかを観察する、楽しくて驚くほど明らかな方法です。
毎ターンの状況のコンパクトな JSON スナップショット (資源、建物、ユニット、戦場の霧の発見、脅威、技術ツリー、マップの境界など)、
固定ツールセット ( train_unit 、 build_struct 、 Research_tech 、 upgrade_age 、 Attack_target 、explore など)、
そして命令はただ一つ、「勝つ」です。
その後、試合全体にわたって、次から次へとそれを続けなければなりません。
ライブマッチ — 霧に限定された 3D ワールド、ストリーミング決定ログ (左)、ランク付けされたリーダーボード (右)、およびミニマップ。
なぜそれが興味深いのか（非科学的であれば）

fic) 評価
ほとんどの簡単な LLM デモでは、単一の賢明な答えが得られます。 LLM コロシアムのフルマッチでは、より難しいものが得られ、人々がエージェントに求めている機能が正確に強調されます。
🎯 プレッシャーの下での正確なツール呼び出し。すべての移動は、適切なパラメーターを備えた単一の有効な JSON アクションである必要があります。ツールを幻覚させたり、スキーマをいじったり、散文で包んだりすると、ターンが無駄になります。文字通り、モデルのフォーマット規律が維持されるか崩壊するかを観察することができます。
🧭 緩い、なじみのない枠組みで運営されている。微調整はなく、「良いプレイ」の例もありません。モデルには、システム プロンプト内のルールとその前の状態のみがあります。これまで遭遇したことのないシステムの動作戦略を推測できるでしょうか?
🧠 長期的なコンテキスト、長期的な戦略。経済→テクノロジー→軍事→征服の連鎖が数十ターンにわたって展開されます。永久に経済を最適化し、軍隊を構築しないモデルは負けます。計画を記憶し、偵察に適応し、リソースをプレッシャーに変換するモデルが勝利します。 (ハーネスは各モデルに永続的な目標と計画を与え、ターンを超えて実行できます。ただし、実際にそれを維持し、従うかはモデル次第です。)
🔁 エラー回復。アクションが拒否されると、モデルは正確な理由を返します (例: 「兵舎がまだ建てられていません。最初に調べてください」)。軌道修正するのか、それとも同じ鍵のかかったドアを叩くのか？
🗺️ 空間とリソースの推論。戦争の霧が地図を隠します。リソースと敵は、使用したり攻撃したりする前に偵察する必要があります。良い遊びとは、推測することではなく、探求することを意味します。
⏱️ レイテンシと品質。各モデルは独自の独立したループを実行します。高速なモデルは、単により頻繁に動作するだけです。現実世界と同様に、優れているが遅いモデルは、まともではあるが速いモデルよりテンポが遅れる場合があります。
p値は得られません。あなたはイムを取得します

実際に演奏できるモデルの直感的な感触。
🤖 4 つのモデル、ライブで戦闘 — それぞれが独自の非同期意思決定パイプライン上で動作するため、より高速なモデルほど実際に頻繁に動作します。
🔌 OpenAI 互換 (OpenAI、vLLM、LM Studio、LiteLLM、Groq、OpenRouter など)、Anthropic、Ollama、Google (Gemini) の自動検出付きモデルを使用できます。同じマッチでローカルとクラウドを混合します。
🔐 すべての認証スタイル — なし、API キー (ベアラー)、ヘッダー シークレット、基本、または OAuth2 (トークンを貼り付けるか、クライアント認証情報を介して取得)。
🧰 モデル ライブラリ — 追加、接続のテスト、提供されるモデルの選択、モデルごとの最大トークン、推論言語、および (Ollama の場合) コンテキスト サイズの設定。ローカルに保存され、ファイルとしてエクスポート/インポート可能です。
📝 プレイヤーごとのシステム プロンプト — 1 つの編集可能なテンプレートから各座席に独自の頭脳 (攻撃的か経済的か、簡潔か冗長か) を与え、スタイルの衝突を観察します。
🛰️ ライブ観客ダッシュボード — ランク付けされたリーダーボード、ストリーミング決定ログ (すべての動き + モデルが述べた理由、拒否されたアクションのフラグ付き)、モデルごとのアドバイス チャット、および任意のモデルの再生/一時停止 (ノルマに達したときに便利)。
📊 試合終了後のモデルの評価 — レイテンシ、意思決定数、アクションの成功率、JSON 形式の忠実度、推論率、エラーの内訳、動作タグ、および透明な 0 ～ 100 の戦略スコア。
🌍 完全にローカライズされた UI — 英語、ドイツ語、スペイン語、簡体字中国語 — モデルの言語はインターフェース言語とは別に選択されます。
🎮 人間もプレイ可能 — スタンダード/ハードの小競り合いと、組み込みルールベースの AI に対するキャンペーン。
🚫 ビルドステップはありません - CDN からのプレーン HTML/CSS/JS + Three.js です。クローン、サーブ、プレイ。
インストールもバンドラーもありません。 HTTP 経由でフォルダーを提供する必要があるだけです (アプリは fetch を使用するため、 file:// からindex.html を開くことはできません)。

git クローン https://github.com/asp67/llm-colosseum.git
CD LLM コロシアム
# 任意の静的サーバーを選択します:
npx http-server 。 -p 8080 -o # ノード
# python3 -m http.server 8080 # Python
# php -S localhost:8080 # PHP
次に、 http://localhost:8080 を開き、 [Play] → [🏟️ Arena] をクリックします。
💡 試合への最速パス: Ollama をインストールし、小さくて簡単なモデル ( ollama pull qwen2.5:7b ) をプルし、 http://localhost:11434 にあるいくつかのアリーナ席を指定します。リアルタイムアリーナでは、小型 + 高速が大型 + 低速に勝ちます。
モデル ライブラリ → モデルを追加します。それぞれ: エンドポイントを設定し、プロトコル/プロバイダーを選択し (または自動検出のままにし)、認証方法を選択し、 🔌 接続のテスト をクリックして、提供されるモデルを選択します。必要に応じて、最大トークン、モデル言語、および (Ollama の場合) コンテキスト サイズを設定します。
アリーナ参加者 → 4 つの座席ごとに文明とコントローラー (モデルの 1 つ、またはルールベースの AI) を選択します。
システム プロンプト → 共有テンプレートを調整するか、個々の座席に独自のプロンプトを与えます。
モデル ライブラリ - ローカル エンドポイントとクラウド エンドポイントを混合し、各接続をテストし、提供されるモデルを選択し、カタログをエクスポート/インポートします。
観戦中に、カードをクリックしてカメラをその基地に飛ばしたり、ドラッグしてパンしたり、モデル アドバイスを送信したり、モデルを完全に一時停止したりできます。決定ログは、モデル自身が述べた理由とともにすべての動きをストリーミングし、拒否されたアクションにフラグを立てます。
💡 32K コンテキスト ウィンドウが最適です。通常、大きいほど悪影響を及ぼします。ハーネスは各ターンのプロンプトをゼロから再構築し、意図的に小さく保ちます。システム プロンプト、それぞれ 1 つの短い文に圧縮された最後の ~20 の動き、モデル自体の継続的な目標 + 計画、および現在の状態のスナップショットです。ゲーム終盤（人口 100 人、数十の建物、発見されたノード）でも、約 12,000 トークンが着地するため、32,000 のウィンドウ リーアが得られます。

ほぼすべての試合で快適なヘッドルームを実現します。 Ollama では、num_ctx のサイズが大きすぎると (128K など)、モデルが CPU に溢れてしまい、ターンが遅くなったり、タイムアウトが発生したりする可能性があります。このため、モデルごとのコンテキスト サイズのデフォルトは 32768 です。特に必要がない限り、そのままにしておきます。
モデルが考えすぎる傾向にある重度の推論/「思考」タイプである場合は、コンテキストではなく最大トークン (出力バジェット) を増やして、推論を完了し、最終的な JSON アクションを出力する余地を残します。レイテンシーにも注意してください。考えれば考えるほどターンが遅くなり、ターンが遅いとコンテキストが問題になる前にリクエストのタイムアウトに達する可能性があります。
試合終了時の評価 — 勝者、各モデルの 0 ～ 100 の戦略スコア、およびその背後にある生の統計 (待ち時間、決定、成功率、形式の忠実度、推論、動作タグ)。
試合終了戦略スコア (0 ～ 100) は透明な合成値であり、ブラック ボックスはありません。
これに加えて、生の統計 (平均/最小/最大レイテンシー、行われた決定、成功率、推論率、および完全なエラーの内訳 (タイムアウト、解析失敗、無効なアクション、拒否)) に加えて、「積極的」、「経済重視」、「フォーマットの問題」、「アクションの発明」などのクイック動作タグも取得できます。
ブラウザ (バックエンドなし)
§── Three.js (r128、CDN 経由) — 3D 世界をレンダリングします
§── ゲームエンジン — 経済、戦闘、戦争の霧、年齢、勝利条件
§── プロバイダー アダプター — OpenAI / Anthropic / Ollama / Google リクエスト シェーピング + 認証
└── モデルごとのエージェント ループ — JSON ゲームステートを構築し、モデルを呼び出し、1 つのアクションを解析し、
それを適用し、次のターンに結果をフィードバックします
各ターン、モデルは構造化スナップショットを受け取り、アクションを 1 つだけ返す必要があります。
{ "action" : " build_struct " 、 "params" : { "buildingType" : " barracks " 、 "reason" : " nee

d 歩兵がリーダーに圧力をかける " } }
エンジンはそれを進行チェーン (進行→研究→構築→リソース→トレーニング) に対して検証し、実行できない場合は正確で対処可能なエラーを返します。これは次のターンでモデルのコンテキストの一部になります。完全な状態コントラクトは game-state-schema.json にあります。
設計により、境界のあるコンテキスト。モデルは、ハーネスの観点からはターン全体でステートレスです。同じプロバイダーかどうかに関係なく、ターンごとにプロンプトが最初から再構築されるため、コンテキストが際限なく増大することはなく、ハーネス (各サーバーの切り捨てルールではありません) がモデルに表示する内容を決定します。その小さなウィンドウ内で長期的なプレイを可能にしておくために、プロンプトには最後の ~20 の手 (各手は 1 行のアクション (「理由」) → OK/FAILED: 結果に圧縮されている) に加えて、モデルが作成した永続的な目標と、モデルが書き換えるまでターンを超えて持続する計画が含まれています。目的/計画は、モデルが任意のアクションに付加できるオプションのフィールドであるため、送信したばかりの偵察 (またはゴールドが必要な理由) がメモリからスクロールアウトすることなく、複数段階の意図 (「敵基地の偵察 → 騎兵集団 → 攻撃」) を生き続けることができます。ハーネスはこれらを保存してエコーバックするだけです。モデルを計画することは決してないので、eval は依然としてモデル自体の戦略的推論を測定します。
アクションセット: train_worker、train_unit、research_tech、upgrade_age、build_structure、build_wonder、harvest_r

[切り捨てられた]

## Original Extract

Four LLMs battle each other in a browser RTS - a non-scientific sandbox eval for precise tool-calling and long-horizon strategic reasoning in an unfamiliar framework. - asp67/llm-colosseum

GitHub - asp67/llm-colosseum: Four LLMs battle each other in a browser RTS - a non-scientific sandbox eval for precise tool-calling and long-horizon strategic reasoning in an unfamiliar framework. · GitHub
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
asp67
/
llm-colosseum
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
64 Commits 64 Commits Screenshots Screenshots css css js js .gitignore .gitignore LICENSE LICENSE README.md README.md game-state-schema.json game-state-schema.json index.html index.html models.json models.json package.json package.json View all files Repository files navigation
Where language models battle for the crown in antiquity.
Four LLMs. One map. One winner.
A browser-based, Age-of-Empires-style real-time strategy game in which competing language models play against each other — while you watch, coach, and score them.
A quick tour: wiring up models in the library, then into the arena.
LLM Colosseum is a sandbox arena for pitting language models against one another at a task they were never trained for: running an economy and an army, in real time, inside a small RTS they've never seen.
It is not a leaderboard, not a peer-reviewed benchmark, and makes no claim to statistical rigor. It is a hands-on, non-scientific testbed — a fun, surprisingly revealing way to watch how different models behave when you drop them into an unfamiliar framework and ask them to act , not chat.
a compact JSON snapshot of its situation every turn (resources, buildings, units, fog-of-war discoveries, threats, tech tree, the map bounds…),
a fixed set of tools ( train_unit , build_structure , research_tech , upgrade_age , attack_target , explore , …),
and a single instruction: win.
Then it has to keep doing that, turn after turn, for an entire match.
A live match — the fog-limited 3D world, the streaming decision log (left), the ranked leaderboard (right), and the minimap.
Why it's an interesting (if unscientific) eval
Most quick LLM demos reward a single clever answer. A full match of LLM Colosseum rewards something harder, and it stresses exactly the capabilities people care about in agents:
🎯 Precise tool calling under pressure. Every move must be a single, valid JSON action with the right parameters. Hallucinate a tool, fumble the schema, or wrap it in prose and the turn is wasted. You can literally watch a model's format discipline hold or crumble.
🧭 Operating in a loose, unfamiliar framework. There's no fine-tuning, no examples of "good play." The model only has the rules in its system prompt and the state in front of it. Can it infer a working strategy for a system it has never encountered?
🧠 Long-context, long-horizon strategy. Economy → technology → military → conquest is a chain that plays out over dozens of turns. Models that optimize their economy forever and never build an army lose. Models that remember their plan, adapt to scouting, and convert resources into pressure win. (The harness gives each model a persistent objective + plan it can carry across turns — but it's up to the model to actually maintain and follow it.)
🔁 Error recovery. When an action is rejected, the model gets a precise reason back (e.g. "barracks not built yet — research it first" ). Does it correct course, or bang on the same locked door?
🗺️ Spatial & resource reasoning. Fog of war hides the map. Resources and enemies must be scouted before they can be used or attacked. Good play means exploring, not guessing.
⏱️ Latency vs. quality. Each model runs its own independent loop — faster models simply act more often. A brilliant-but-slow model can be out-tempoed by a decent-but-fast one, just like in the real world.
You won't get a p-value. You will get an immediate, visceral feel for which models can actually play .
🤖 4 models, fighting live — each on its own asynchronous decision pipeline, so faster models genuinely move more often.
🔌 Bring any model — OpenAI-compatible (OpenAI, vLLM, LM Studio, LiteLLM, Groq, OpenRouter, …), Anthropic , Ollama , and Google (Gemini) , with auto-detection. Mix local and cloud in the same match.
🔐 Every auth style — none, API key (Bearer), header secret, Basic, or OAuth2 (paste a token or fetch via client-credentials).
🧰 Model library — add, test connection , pick the served model, set per-model max tokens , reasoning language , and (for Ollama) context size . Saved locally and exportable/importable as a file.
📝 Per-player system prompts — give each seat its own brain (aggressive vs. economic, terse vs. verbose) from one editable template, and watch the styles collide.
🛰️ Live spectator dashboard — a ranked leaderboard, a streaming decision log (every move + the model's stated reason, rejected actions flagged), per-model advice chat , and play/pause for any model (handy when one hits a quota).
📊 End-of-match model evaluation — latency, decision count, action-success rate, JSON format fidelity, reasoning rate, error breakdown, behavior tags, and a transparent 0–100 strategy score .
🌍 Fully localized UI — English, German, Spanish, Simplified Chinese — with the model's language chosen separately from the interface language.
🎮 Also human-playable — Standard / Hard skirmishes and a Campaign vs. the built-in rule-based AI.
🚫 No build step — it's plain HTML/CSS/JS + Three.js from a CDN. Clone, serve, play.
No install, no bundler. You just need to serve the folder over HTTP (the app uses fetch , so opening index.html from file:// won't work).
git clone https://github.com/asp67/llm-colosseum.git
cd llm-colosseum
# pick any static server:
npx http-server . -p 8080 -o # Node
# python3 -m http.server 8080 # Python
# php -S localhost:8080 # PHP
Then open http://localhost:8080 and click Play → 🏟️ Arena .
💡 Fastest path to a match: install Ollama , pull a small, quick model ( ollama pull qwen2.5:7b ), and point a couple of arena seats at http://localhost:11434 . Small + fast beats large + slow in a real-time arena.
Model Library → add your models. For each: set the endpoint , pick the protocol/provider (or leave on auto-detect), choose an auth method, hit 🔌 Test connection , and select the served model. Optionally set max tokens , the model language , and (for Ollama) the context size .
Arena participants → for each of the 4 seats choose a civilization and a controller (one of your models, or the rule-based AI).
System prompt → tweak the shared template, or give individual seats their own prompt.
The model library — mix local and cloud endpoints, test each connection, pick the served model, and export/import the catalogue.
While spectating you can click a card to fly the camera to that base, drag to pan, send a model advice , or pause a model entirely. The decision log streams every move alongside the model's own stated reason, and flags any rejected action:
💡 A 32K context window is the sweet spot — bigger is usually worse. The harness rebuilds each turn's prompt from scratch and keeps it deliberately small: the system prompt, the last ~20 moves compressed to one short sentence each, the model's own standing objective + plan , and the current state snapshot. Even a maxed-out late game (100 population, dozens of buildings and discovered nodes) lands around ~12K tokens , so a 32K window leaves comfortable headroom in virtually every match. Going much larger rarely helps and can hurt — on Ollama, an oversized num_ctx (e.g. 128K) can spill the model onto the CPU and cause slow turns or timeouts. The per-model context size defaults to 32768 for this reason; leave it there unless you have a specific need.
If a model is a heavy reasoning / "thinking" type that tends to overthink, raise its max tokens (the output budget) — not its context — so it has room to finish reasoning and still emit the final JSON action. Watch latency too: more thinking means slower turns, and a slow turn can hit the request timeout before context ever becomes an issue.
End-of-match evaluation — a winner, each model's 0–100 strategy score, and the raw stats behind it (latency, decisions, success rate, format fidelity, reasoning, behavior tags).
The match-end Strategy Score (0–100) is a transparent composite — no black box:
Alongside it you get raw stats — average/min/max latency, decisions made, success ratio, reasoning rate, and a full error breakdown (timeouts · parse fails · invalid actions · rejected) — plus quick behavior tags like Aggressive , Economy-focused , Format issues , or Invents actions .
Browser (no backend)
├── Three.js (r128, via CDN) — renders the 3D world
├── Game engine — economy, combat, fog of war, ages, win conditions
├── Provider adapters — OpenAI / Anthropic / Ollama / Google request shaping + auth
└── Per-model agent loop — builds the JSON game-state, calls the model, parses ONE action,
applies it, feeds the result back next turn
Each turn a model receives a structured snapshot and must return exactly one action:
{ "action" : " build_structure " , "params" : { "buildingType" : " barracks " , "reason" : " need infantry to pressure the leader " } }
The engine validates it against the advancement chain (advance → research → build → resources → train) and returns a precise, actionable error if it can't be done — which becomes part of the model's context on the next turn. The full state contract is in game-state-schema.json .
Bounded context, by design. The model is stateless across turns from the harness's point of view — every turn the prompt is rebuilt from scratch, identical providers or not, so the context never grows unboundedly and the harness (not each server's truncation rules) decides what the model sees. To keep long-horizon play possible inside that small window, the prompt carries the last ~20 moves (each compressed to a one-line action ("reason") → OK/FAILED: outcome ) plus a model-authored standing objective + plan that persists across turns until the model rewrites it. The objective/plan are optional fields the model can attach to any action, so it can keep a multi-step intent alive — "scout the enemy base → mass cavalry → attack" — without a scout it just sent (or the reason it needed gold) scrolling out of memory. The harness only stores and echoes these back; it never plans for the model, so the eval still measures the model's own strategic reasoning.
Action set: train_worker · train_unit · research_tech · upgrade_age · build_structure · build_wonder · harvest_r

[truncated]
