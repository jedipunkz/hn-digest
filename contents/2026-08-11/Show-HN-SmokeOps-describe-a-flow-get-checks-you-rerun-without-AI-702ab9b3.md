---
source: "https://github.com/gate3/SmokeOps"
hn_url: "https://news.ycombinator.com/item?id=49262241"
title: "Show HN: SmokeOps – describe a flow, get checks you rerun without AI"
article_title: "GitHub - gate3/SmokeOps: Turn plain-language prompts into approved website checks you can run on a schedule. · GitHub"
author: "crackedev"
captured_at: "2026-08-11T18:47:37Z"
capture_tool: "hn-digest"
hn_id: 49262241
score: 1
comments: 0
posted_at: "2026-08-11T18:12:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: SmokeOps – describe a flow, get checks you rerun without AI

- HN: [49262241](https://news.ycombinator.com/item?id=49262241)
- Source: [github.com](https://github.com/gate3/SmokeOps)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T18:12:40Z

## Translation

タイトル: Show HN: SmokeOps – フローを説明し、AI なしで再実行するチェックを取得します
記事のタイトル: GitHub - Gate3/SmokeOps: 平易な言語のプロンプトを、スケジュールに従って実行できる承認済みの Web サイト チェックに変換します。 · GitHub
説明: 平易な言語のプロンプトを、スケジュールに従って実行できる承認済みの Web サイト チェックに変換します。 - ゲート3/SmokeOps

記事本文:
GitHub - Gate3/SmokeOps: 平易な言語のプロンプトを、スケジュールに従って実行できる承認済みの Web サイト チェックに変換します。 · GitHub
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
ゲート3
/
スモークオプス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
99 コミット 99 コミット .cursor/ rules .cursor/ rules .github .github docs docs サンプル サンプル モニター モニター src/スモークオプス src/スモークオプス テスト テスト .DS_Store .DS_Store .env.example .en

v.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス Makefile Makefile NOTICE NOTICE README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイル ナビゲーション
重要なフローを説明します。 SmokeOps はサイトを駆動し、実際に合格したステップのみを保持し、CI で再実行できる Playwright チェックを残します。承認後にエージェントは必要ありません。
SmokeOps はローカルファーストの合成モニタリングです。ユーザー フローを平易な言葉で説明します。それをブラウザー チェックに変換し、確認してリポジトリに保存し、スケジュールに従って実行します。
AI エージェントはサイトを開いてフローをクリックして実行したり、Playwright スクリプトを作成したりできます。 SmokeOps は、ライブ ページでのより緊密な作成ループを中心に構築されています。インタラクティブなものを確認し、次のステップを提案し、Playwright で実行し、ステップが失敗した場合は、新しいページ コンテキストを取得して再試行します。すでに渡されたものだけが保持されます。一度承認すると、チェックはリポジトリに保存されます。その後、エージェントやホストされたモニターを使用せず、Playwright のみでスモークオプス実行 (および CI/cron) のリプレイが実行されます。
必要ない場合: Playwright をすでに快適に作成および保守している場合、SmokeOps はオプションです。 「この流れはまだ機能するのか？」を知りたい人にとっては勝利です。テスト フレームワークを所有する必要がなく、1 回限りのエージェントの実行だけでなく、ライブ作成後の永続的なチェックを必要とする人にとっても最適です。
SmokeOps はアルファ版です (積極的に開発されており、現実世界に適していますが、荒削りな部分もあります)。 LLM API キーは、小切手を作成する場合にのみ必要です。承認された実行は LLM を呼び出しません。
Python 3.10 以降が必要です。 CLI を環境にインストールします。
pip インストール Smoopops
# または
UVツールでsmoopsをインストール
インストールを確認します。
スモークオプスバージョン
初めての作成または実行時に、SmokeOps は Chromium を自動的にインストールします

足りない場合は味方。 CI 環境では、明示的な Playwright インストール ステップをワークフローに含める必要があります (例は以下に示されています)。
このウォークスルーでは、付属の PeakAir デモ サイト (サンプル反応アプリ) を使用して、ワークフロー全体を約 5 分で示します。
1. デモサイトを開始します（オプション）
以下のクイックスタートでは、このリポジトリに含まれるサンプル アプリである PeakAir を使用します。リポジトリのクローンを作成し、ローカルで開始します。
git クローン https://github.com/gate3/SmokeOps.git
cd SmokeOps/examples/peakair-hvac
npmインストール
npm 実行開発
PeakAir は http://localhost:3000 で実行されます。このターミナルは実行したままにしておきます。
PyPI から SmokeOps をインストールしていて、リポジトリのクローンを作成したくない場合は、この手順をスキップしてください。ステップ 3 では、URL とプロンプトをアクセス可能なサイト (独自のアプリやパブリック ステージング URL など) に置き換えます。
新しいターミナルで、monitors/ を配置するディレクトリから作業します (このチュートリアルでは、SmokeOps リポジトリのルート: cd SmokeOps )。
.env ファイルを作成するか、環境変数をエクスポートします。 SmokeOps は、LiteLLM を介して LLM 呼び出しをルーティングし、Gemini (デフォルト)、OpenAI、およびその他のプロバイダーをサポートします。
export LLM_PROVIDER=ジェミニ
エクスポート DEFAULT_MODEL=gemini-3.6-フラッシュ
GEMINI_API_KEY=your_key_here をエクスポート
3. 最初の小切手を作成します
スモークオプス作成 --url " http://localhost:3000/ " \
" ホームページがロードされ、製品名 PeakAir が表示されることを確認します。"
デフォルトでは、create はヘッドレスで実行されます。Playwright はブラウザ ウィンドウが表示されずにバックグラウンドで Chromium を駆動します。ターミナルで進行状況 (ナビゲーション、DOM プルーニング、ステップの結果) を追跡します。すべてのステップが完了すると、SmokeOps は概要を表示し、承認を待ちます。
ステップの実行中にブラウザを監視するには:
steamops create --headed --url " http://localhost:3000/ " \
" ホームページがロードされ、製品名 PeakAir が表示されることを確認します。"
--headed は実際のブラウザ風を開きます

ow は、実行中の各ステップを出力し、アサート間で少し停止するので、従うことができます。実行が終了したら、Enter キーを押してウィンドウを閉じ、通常どおり承認または拒否します。
SmokeOps がステップを実行すると、概要が表示されます。次のいずれかを選択します。
accept — 現在のディレクトリのmonitors/<name>.jsonに保存します（チェック名はプロンプトのスラッグです。受け入れる前に下書きを編集して名前を変更します）
edit - エディターで JSON を開いて、名前または概要を調整します。
これを受け入れてください。受け入れ後に出力されるパスに注目してください (たとえば、monitors/verify-the-homepage-loads-and-the-product-name-peakair-is-visible.json )。
ここで、受け入れステップからのパスを使用して、LLM を使用せずにチェックを実行します。
スモークオプス実行モニター/ <名前> .json
または、monitors/*.json から対話的に選択するパスを省略します。
スモークオプスラン
この時点からは LLM は必要ありません。チェックは、承認した Playwright の手順のみを使用して、オフラインで実行されます。合格すると、 PASS と表示されます。失敗した場合、SmokeOps はデバッグ用にスクリーンショットとトレースを .smokeops/artifacts/ に保存します。
バンドルされたジャーニー チェックを試してください (クローンのみ): このリポジトリには、作成せずに実行できる承認済みの複数ページの PeakAir チェックも含まれています。
スモークオプス実行モニター/complete-the-signup-form-with-sample-contact-details-choose-the-seasonal-tune-up.json
これが中心となるループです: 作成 (LLM が提案) → 承認 (あなたが決定) → 実行 (オフライン検証、スケジュール通り)。
シナリオ
代替案
スモークオプス
Playwright を手書きで書く
時間がかかる、脆いセレクター、ドリフトしやすい
LLM がステップを提案し、一度確認すれば、JSON はリポジトリに残ります
ホスト型モニター (DataDog、New Relic など) を使用している場合
不透明、ベンダーロックイン、UX をコードから分離
リポジトリ内の透過的な JSON、ポータブル、GitHub Actions によるスケジュール
合成モニタリングをスキップします
重要なフローが可視化されない
5分秒

タップ、オフライン実行、所有アーティファクト
カスタムセレクターまたはロジックが必要です
ホスト型サービスによるカスタマイズの制限
JSON を直接編集します。人間が読める形式です
SmokeOps ではないもの
完全な E2E テスト フレームワークではありません。数百のシナリオを含む包括的なテスト スイートの場合は、Playwright Test を使用してください。
視覚的な回帰ツールではありません — ピクセル完璧な比較ではなく、フローの検証とアサーションに重点を置いています
CI のローカルホスト用に設計されていない - GitHub アクションにはパブリック ステージング URL またはトンネル サービスを使用してください
監視プラットフォームではありません。ダッシュボード、アラート、履歴傾向分析はありません。そのためには DataDog や Grafana などのツールを使用します
create 中に、SmokeOps はライブ ページ上のインタラクティブな内容を確認し、モデルに次のステップを要求し、それらのステップを Playwright で実行します。ステップが失敗した場合は、新しいページ コンテキストを取得して再試行します。渡されたセグメントはコミットされたままになります。十分な失敗があった後は、悪いチェックを強制するのではなく、ストップとエスカレーションを作成します。一度承認すると、チェックはリポジトリに保存されます。その後、実行と CI は Playwright のみを使用します (LLM は使用しません)。現在のデフォルト: 戦略変更前に 5 回の失敗、作成ごとに最大 5 つのセグメント。
CLI — 作成、承認/拒否/編集、実行、表示、エクスポート
DOM プルーナー — ライブ ページ上のインタラクティブな要素を検出し、モデルに作業の基礎となる小さな固定セットを提供します。
ループの作成 - 提案→検証→Playwrightで実行→渡されたものをコミット、または再試行/エスカレーション
セレクターとステップ実行 — 耐久性のあるロケーターと Playwright アクション/アサート
承認＆チェックストア — ヒューマンゲート;ドラフトは監視下でコミットされた JSON になります/
リプレイ — スモークオプス実行 (および CI) は、LLM を使用せずに保存されたステップを実行します。失敗してもスクリーンショット/トレースアーティファクトを保存できる
フローチャート TD
プロンプト[プロンプト + URL] --> create[ループの作成]
作成 --> 承認[あなたが承認します]
承認 --> json[Che

リポジトリ内の JSON を確認してください]
json --> run[smokeops run / CI]
読み込み中
ループの作成
フローチャート LR
prune[インタラクティブな要素を取り除く] -->proposal[ステップを提案する]
提案 --> 検証[提案の検証]
検証 -->|ok| verify[Playwrightで実行]
検証 -->|パス、プロンプトはカバーされています|完了[承認の準備ができました]
検証 -->|無効|提案する
verify -->|パス、プロンプトはまだ終了していません|提案する
検証 -->|失敗| reground[新しいページのコンテキスト]
やり直し --> 提案する
提案 -->|失敗予算ヒット|エスカレート[エスカレート]
読み込み中
設計上の決定
モデルがページを参照する前にプルーニングします。完全な DOM はノイズが多くなります。インタラクティブな要素と一時的な ID により、提案は実際に画面上にあるものに基づいたものになります。
コミット前に検証 — Playwright がライブ ページで正常に実行するまで、プロポーザルは保持されません。
ライブ失敗コンテキストで再試行します。失敗すると、次のプロポーザルでは、推測された DOM ではなく、ページがそのまま表示されます (URL、タイトル、表示テキスト)。
耐久性のあるアーティファクトの前に人間が承認します。承認されたチェックのみが CI およびスケジュールのためにモニターに表示されます。
作成と実行を分割する - 作成には LLM を使用できます。監視はどちらかに依存してはなりません。
エリア
場所
ステートマシンの作成
src/smokeops/core/agents/smoke_check_proposer.py 、proposal_session.py
DOM プルーニング
src/smokeops/core/dom_pruner.py
セレクターと実行
src/smokeops/core/selector_builder.py 、 step_executor.py
承認と保管
src/smokeops/core/approval.py 、check_store.py
CLI エントリポイント
src/smokeops/cli.py
デザインに関する質問
質問
SmokeOps の答え
理性はどこに存在するのでしょうか？
LLM は手順を提案します。ハーネスはそれらを実行し、失敗した場合は新しいページ コンテキストで再試行し、渡されたもののみを保持します。耐久性が得られる前に承認してください。
LLM はどのような場合に使用されますか?
create 時のみ。 run、CI、export はモデルを呼び出しません。
真実の情報源は何ですか?
承認されたチェック JSON が

リポジトリ (スタンドアロン Playwright ファイルへのオプションのエクスポート)。
作成が失敗するとどうなりますか?
失敗したステップはライブ ページで再実行され、（失敗の予算まで）再試行され、その後、サイレントにチェックを作成するのではなくエスカレーションされます。
利用ガイド
クイックスタートはそのループを 1 回実行します。これらのセクションでは、各ステップをさらに詳しく説明します。
何をテストするかを平易な言葉で説明します。
サインアップ ページで、電子メールを入力してフォームを送信し、アカウント作成メッセージが表示されることを確認します。
SmokeOps create を実行すると、SmokeOps は次のようになります。
指定した URL に移動します
ページ上のインタラクティブな要素 (ボタン、入力、リンク、フォーム コントロール) を検索します。
LLM に、プロンプトに一致する Playwright ステップを提案するよう依頼します
これらのステップを実際のブラウザで実行して、機能することを確認します。
概要を表示し、決定を待ちます
デフォルトでは、create はヘッドレス (ブラウザ ウィンドウが表示されない) で実行されます。 --headed を追加すると、ブラウザーを監視し、ターミナルの各ステップを確認できます。
スモークオプス作成 --url " https://example.com/signup " \
" 電子メールとパスワードを使用してサインアップを完了し、ダッシュボードが読み込まれることを確認します。"
steamops create --headed --url " https://example.com/signup " \
" 電子メールとパスワードを使用してサインアップを完了し、ダッシュボードが読み込まれることを確認します。"
--headed を使用すると、SmokeOps はブラウザ ウィンドウを開き、実行中の各ステップを出力し、プロンプトを表示します。

[切り捨てられた]

## Original Extract

Turn plain-language prompts into approved website checks you can run on a schedule. - gate3/SmokeOps

GitHub - gate3/SmokeOps: Turn plain-language prompts into approved website checks you can run on a schedule. · GitHub
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
gate3
/
SmokeOps
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
99 Commits 99 Commits .cursor/ rules .cursor/ rules .github .github docs docs examples examples monitors monitors src/ smokeops src/ smokeops tests tests .DS_Store .DS_Store .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Describe a critical flow. SmokeOps drives your site, keeps only steps that actually passed, and leaves you a Playwright check you can rerun in CI — no agent required after you approve.
SmokeOps is local-first synthetic monitoring. You describe a user flow in plain language; it turns that into a browser check you review, save in your repository, and run on a schedule.
An AI agent can open your site and click through a flow, or write you a Playwright script. SmokeOps is built around a tighter create loop on the live page: look at what’s interactive, propose the next steps, run them in Playwright, and if a step fails, take fresh page context and try again. It only keeps what already passed. You approve once; the check is saved in your repo. After that, smokeops run (and CI/cron) replays with Playwright alone — no agent, no hosted monitor.
When you might not need it: If you already write and maintain Playwright comfortably, SmokeOps is optional. The win is for people who want “does this flow still work?” without owning a test framework — and for anyone who wants a durable check after a live create, not only a one-off agent run.
SmokeOps is alpha (actively developed, real-world friendly, some rough edges). An LLM API key is needed only when you create a check; approved runs do not call an LLM.
Requires Python 3.10 or later . Install the CLI into your environment:
pip install smokeops
# or
uv tool install smokeops
Confirm the installation:
smokeops version
On your first create or run , SmokeOps installs Chromium automatically if it's missing. In CI environments, you need to keep an explicit playwright install step in your workflows (examples are provided below).
This walkthrough shows the entire workflow in about 5 minutes using the included PeakAir demo site (a sample react app).
1. Start the demo site (optional)
The quickstart below uses PeakAir , a sample app included in this repository. Clone the repo and start it locally:
git clone https://github.com/gate3/SmokeOps.git
cd SmokeOps/examples/peakair-hvac
npm install
npm run dev
PeakAir runs at http://localhost:3000 . Leave this terminal running.
If you installed SmokeOps from PyPI and do not want to clone the repo, skip this step. In step 3, replace the URL and prompt with a site you can reach (for example your own app or a public staging URL).
In a new terminal , work from the directory where you want monitors/ to live (for this walkthrough, the SmokeOps repo root: cd SmokeOps ).
Create a .env file or export environment variables. SmokeOps routes LLM calls through LiteLLM and supports Gemini (default), OpenAI, and other providers:
export LLM_PROVIDER=gemini
export DEFAULT_MODEL=gemini-3.6-flash
export GEMINI_API_KEY=your_key_here
3. Create your first check
smokeops create --url " http://localhost:3000/ " \
" Verify the homepage loads and the product name PeakAir is visible "
By default, create runs headless : Playwright drives Chromium in the background with no visible browser window. You follow progress in the terminal (navigation, DOM pruning, step results). When all steps pass, SmokeOps shows a summary and waits for your approval.
To watch the browser while steps run:
smokeops create --headed --url " http://localhost:3000/ " \
" Verify the homepage loads and the product name PeakAir is visible "
--headed opens a real browser window, prints each step as it runs, and pauses briefly between asserts so you can follow along. When the run finishes, press Enter to close the window, then approve or reject as usual.
After SmokeOps runs the steps, you'll see a summary. Choose one of:
accept — saves to monitors/<name>.json in your current directory (the check name is a slug from your prompt; edit the draft to rename before accept)
edit — opens the JSON in your editor to adjust the name or summary
Accept this one. Note the path printed after accept (for example monitors/verify-the-homepage-loads-and-the-product-name-peakair-is-visible.json ).
Now run the check without the LLM, using the path from the accept step:
smokeops run monitors/ < name > .json
Or omit the path to pick from monitors/*.json interactively:
smokeops run
No LLM is needed from this point. The check runs offline, using only the Playwright steps you approved. If it passes, you see PASS . If it fails, SmokeOps saves a screenshot and trace under .smokeops/artifacts/ for debugging.
Try a bundled journey check (clone only): this repo also includes an approved multipage PeakAir check you can run without creating one:
smokeops run monitors/complete-the-signup-form-with-sample-contact-details-choose-the-seasonal-tune-up.json
That's the core loop: create (LLM proposes) → approve (you decide) → run (offline verification, on schedule).
Scenario
Alternative
SmokeOps
You write Playwright by hand
Time-consuming, brittle selectors, easy to drift
LLM proposes steps, you review once, JSON stays in your repo
You use a hosted monitor (e.g., DataDog, New Relic)
Opaque, vendor lock-in, separate UX from your code
Transparent JSON in your repo, portable, schedule with GitHub Actions
You skip synthetic monitoring
No visibility into critical flows
5-min setup, offline execution, owned artifact
You need custom selectors or logic
Hosted services limit customization
Edit the JSON directly; it's human-readable Playwright
What SmokeOps is NOT
Not a full E2E test framework — for comprehensive test suites with hundreds of scenarios, use Playwright Test
Not a visual regression tool — focuses on flow verification and assertions, not pixel-perfect comparisons
Not designed for localhost in CI — use a public staging URL or tunnel service for GitHub Actions
Not a monitoring platform — no dashboards, alerts, or historical trend analysis; use tools like DataDog or Grafana for that
During create , SmokeOps looks at what’s interactive on the live page, asks the model for the next steps, runs those steps in Playwright, and if a step fails, takes fresh page context and tries again. Passed segments stay committed. After enough failures, create stops and escalates instead of forcing a bad check. You approve once; the check is saved in your repo. After that, run and CI use Playwright alone — no LLM. Defaults today: five failures before strategy shift, five segments max per create.
CLI — create , approve/reject/edit, run , show , export
DOM pruner — finds interactive elements on the live page and gives the model a small, grounded set to work from
Create loop — propose → validate → run in Playwright → commit what passed, or retry / escalate
Selectors & step execution — durable locators plus Playwright actions/asserts
Approval & check store — human gate; draft becomes committed JSON under monitors/
Replay — smokeops run (and CI) executes stored steps with no LLM; failures can save screenshot/trace artifacts
flowchart TD
prompt[Prompt + URL] --> create[Create loop]
create --> approve[You approve]
approve --> json[Check JSON in repo]
json --> run[smokeops run / CI]
Loading
Create loop
flowchart LR
prune[Prune interactive elements] --> propose[Propose steps]
propose --> validate[Validate proposal]
validate -->|ok| verify[Run in Playwright]
verify -->|pass, prompt covered| done[Ready for approve]
validate -->|invalid| propose
verify -->|pass, prompt not finished yet| propose
verify -->|fail| reground[Fresh page context]
reground --> propose
propose -->|failure budget hit| escalate[Escalate]
Loading
Design decisions
Prune before the model sees the page — full DOM is noisy; interactive elements plus temporary ids keep proposals grounded in what is actually on screen.
Verify before commit — a proposal is not kept until Playwright runs it successfully on the live page.
Retry with live failure context — on failure, the next proposal sees the page as it was (URL, title, visible text), not a guessed DOM.
Human approve before durable artifact — only accepted checks land in monitors/ for CI and scheduling.
Split create from run — creation can use an LLM; monitoring must not depend on one.
Area
Location
Create state machine
src/smokeops/core/agents/smoke_check_proposer.py , proposal_session.py
DOM pruning
src/smokeops/core/dom_pruner.py
Selectors & execution
src/smokeops/core/selector_builder.py , step_executor.py
Approval & storage
src/smokeops/core/approval.py , check_store.py
CLI entrypoints
src/smokeops/cli.py
Design questions
Question
SmokeOps’s answer
Where does reasoning live?
The LLM proposes steps. The harness runs them, retries with fresh page context on failure, and only keeps what passed. You approve before anything is durable.
When is the LLM used?
Only during create . run , CI, and export do not call a model.
What is the source of truth?
An approved check JSON in your repo (optional export to a standalone Playwright file).
What happens when create fails?
Failed steps re-ground on the live page and retry (up to a failure budget), then escalate instead of silently inventing a check.
Usage Guides
The quickstart runs through that loop once. These sections go deeper on each step.
Describe what to test in plain language:
On the signup page, fill in an email, submit the form, and verify the account-created message appears.
When you run smokeops create , SmokeOps:
Navigates to the URL you provide
Finds interactive elements on the page (buttons, inputs, links, form controls)
Asks the LLM to propose Playwright steps that match your prompt
Runs those steps in a real browser to verify they work
Shows you a summary and waits for your decision
By default, create runs headless (no visible browser window). Add --headed to watch the browser and see each step in the terminal.
smokeops create --url " https://example.com/signup " \
" Complete a signup with email and password, then verify the dashboard loads "
smokeops create --headed --url " https://example.com/signup " \
" Complete a signup with email and password, then verify the dashboard loads "
With --headed , SmokeOps opens a browser window, prints each step as it runs, and prompts yo

[truncated]
