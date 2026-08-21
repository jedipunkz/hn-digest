---
source: "https://github.com/tX-c0re/xss-grenade"
hn_url: "https://news.ycombinator.com/item?id=49386036"
title: "I Tried build AI web vulnerable scanner"
article_title: "GitHub - tX-c0re/xss-grenade: Modern XSS vulnerability scanner for real-world web applications. · GitHub"
image: "https://opengraph.githubassets.com/dac2780488470c7a56e913fc018b9847a5d992fd1af88d12191ca4fd5c41f7f8/tX-c0re/xss-grenade"
author: "txc0re"
captured_at: "2026-08-21T10:20:04Z"
capture_tool: "hn-digest"
hn_id: 49386036
score: 1
comments: 1
posted_at: "2026-08-21T10:14:11Z"
tags:
  - hacker-news
  - translated
---

# I Tried build AI web vulnerable scanner

- HN: [49386036](https://news.ycombinator.com/item?id=49386036)
- Source: [github.com](https://github.com/tX-c0re/xss-grenade)
- Score: 1
- Comments: 1
- Posted: 2026-08-21T10:14:11Z

## Translation

タイトル: Web脆弱性AIスキャナーを作ってみた
記事のタイトル: GitHub - tX-c0re/xss-grenade: 現実世界の Web アプリケーション用の最新 XSS 脆弱性スキャナー。 · GitHub
説明: 現実世界の Web アプリケーション用の最新の XSS 脆弱性スキャナー。 - GitHub - tX-c0re/xss-grenade: 現実世界の Web アプリケーション用の最新の XSS 脆弱性スキャナー。

記事本文:
GitHub - tX-c0re/xss-grenade: 現実世界の Web アプリケーション用の最新の XSS 脆弱性スキャナー。 · GitHub
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
tX-c0re
/
XSSグレネード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
1 コミット 1 コミット フォルダーとファイル
アセット アセット ベンチ_ファイアリングレンジ ベンチ_ファイアリングレンジ .gitignore .gitignore ライセンス ライセンス README.md README.md _アタック_グラフ_v2_ソース.py _アタック_グラフ_v2_ソース.py _auth.py _au

th.py _blind_xss_oob.py _blind_xss_oob.py _breakout_synth.py _breakout_synth.py _cache_poisoning.py _cache_poisoning.py _checkpoint.py _checkpoint.py _css_injection.py _css_injection.py _dom_clobbering.py _dom_clobbering.py _dom_exploit_engine.py _dom_exploit_engine.py _dom_v6.py _dom_v6.py _dompurify_config.py _dompurify_config.py _dompurify_cve_feed.py _dompurify_cve_feed.py _engine.py _engine.py _exploit_classifier.py _exploit_classifier.py _finding_scorer.py _finding_scorer.py _finding_store.py _finding_store.py _graphql_xss.py _graphql_xss.py _headless_crawler.py _headless_crawler.py _headless_verifier.py _headless_verifier.py _html_analyzer.py _html_analyzer.py _html_report.py _html_report.py _htmx_alpine.py _htmx_alpine.py _js_analyzer.py _js_analyzer.py _library_cve_feed.py _library_cve_feed.py _markdown_xss.py _markdown_xss.py _mutation_xss.py _mutation_xss.py _oob_collector.py _oob_collector.py _open_redirect.py _open_redirect.py _param_wordlist.py _param_wordlist.py _poc_generator.py _poc_generator.py _proto_pollution_analyzer.py _proto_pollution_analyzer.py _render_gate.py _render_gate.py _response_aware.py _response_aware.py _sarif_report.py _sarif_report.py _sourcemap_analyzer.py _sourcemap_analyzer.py _spa_route_extractor.py _spa_route_extractor.py _static_js_analyzer.py _static_js_analyzer.py _stored_xss_tracker.py _stored_xss_tracker.py _template_injection.py _template_injection.py _trusted_types_analyzer.py _trusted_types_analyzer.py _url_analyzer.py _url_analyzer.py benchmark_scoreboard.py benchmark_scoreboard.py context_engine.py context_engine.py destructive_real_test_server.py destructive_real_test_server.py dom_hooks_v6.js dom_hooks_v6.js head2head_dalfox.py head2head_dalfox.py payloads.txt payloads.txt profile_phases.py profile_phases.py repro_le

vel2.py repro_level2.py リ
[切り捨てられた]
最新の Web アプリケーション向けの次世代 XSS 検出エンジン。
ペイロードのファジングだけでなく、実際のセキュリティ テスト用に構築されています。
許可された使用のみ。自分が所有しているシステム、または書面による明示的なテスト許可を得ているシステム (対象範囲内のバグ報奨金プログラムまたは署名された契約) のみをテストします。不正なスキャンはほとんどの管轄区域で違法であり、このツールの使用方法についてはお客様の責任となります。
ほとんどのスキャナは反射で停止します。つまり、「入力がページに戻ってきた」ということです。 XSS Grenade はさらに進んで、実際のブラウザで悪用可能な実際の脆弱性を確認します。
🧠 コンテキスト認識型ペイロード エンジン — HTML / 属性 / URL / JS / CSS / コメント
🌐 最新のアプリ向けに構築 — SPA、React、Angular、Vue、SSR / ハイドレーション
🔍 エコーされたペイロードだけでなく、実際の攻撃ベクトルを検出します
⚡ ヘッドレスブラウザ検証 — 誤検知をほぼゼロに削減します
💣 バグ報奨金に焦点を当てた — 実際に支払われるベクトル
XSS は初めてですか?クロスサイト スクリプティングとは、アプリが攻撃者が制御する入力を被害者のブラウザーでアクティブ コードとしてレンダリングすることです。 「リフレクション」は、入力が応答に現れたことを意味するだけで、実際に実行されるまではバグではありません。 XSS Grenade は、最後の難しいステップを実行します。実際のヘッドレス Chromium にページをロードし、ペイロードが起動するかどうかを確認します。
コンテキストを意識したインジェクション
パラメーターが到着する場所を検出し、そのコンテキスト (同じパラメーターに対する複数のコンテキストを含む) に適合するペイロードのみを起動します。
リアルブラウザでの検証
候補はヘッドレス Chromium (Playwright) に再ロードされます。実際に実行された結果のみが残ります。
静的JSテイント分析
JavaScript を AST に解析し、信頼できないソース ( location.* 、 document.referrer 、 window.name 、 postMessage ) を危険なシンク ( innerHTML 、 eval 、 document.writ ) に追跡します。

e 、フレームワークが沈みます)。
最新の脆弱性クラス
DOM XSS · 突然変異 XSS (mXSS) · プロトタイプ汚染 → XSS · DOM 破壊 · 信頼できるタイプの設定ミス · SSR ハイドレーションの問題。
バグ報奨金ベクトル
postMessage の不正使用 · JSONP コールバック インジェクション · 未解決のマークアップ (スクリプトレス、CSP 耐性) · SVG/XML コンテンツ タイプのリフレクション。
既知の CVE ライブラリの検出
フィンガープリント React、Vue、Angular、Next.js、jQuery、lodash、DOMPurify… および既知の XSS/RCE CVE を使用したバージョンにフラグを立てます。
🖥️ インターフェースのプレビュー
ライブの攻撃対象領域グラフ、リアルタイムの重大度ランク付けされた結果、ブラウザーで検証された結果はすべて、単一の PyQt5 デスクトップ アプリから得られます。
git クローン https://github.com/tX-c0re/xss-grenade.git
CD XSS グレネード
python3 -m venv venv
ソース venv/bin/activate # Windows: venv\Scripts\activate
pip install -r 要件.txt
python -m playwright install chromium # ヘッドレス検証用
Python xss_grenade_gui.py
ヒント
ログインがありませんか?問題ありませんが、興味深い XSS の約 80% が認証の背後に隠れています。セッション Cookie を [設定] → [認証] に貼り付けて、管理パネル、プロファイル、ダッシュボードをスキャンします。 Cookie はメモリ内にのみ残り、レポートやログには書き込まれません。
上部のバーにターゲット URL を入力します。
[設定] タブでモジュールを選択します (または適切なデフォルトのままにします)。
調査結果がリアルタイムで表示されるのを確認します。攻撃対象領域のグラフと調査結果テーブルはライブで更新されます。
レポートをエクスポートします: [保存] → [JSON + HTML + PoC バンドル]。
✔ 結果は実際のブラウザで検証されます。 • ✔ 出力はノイズではなく、悪用可能な脆弱性です。
「スマート ペイロード」を有効にして、高速で高信号のファースト パスを実行します。
⚡ スキャン時間を数時間から数分に短縮
🎯 高信号のコンテキストごとのカバレッジを維持
💡 ディープラン前のファーストパス偵察に最適
完全モジュール式 — 必要なものだけを有効にします。
ワンクリックでエクスポート、3 分でエクスポート

オーマット:
JSON — 構造化され、自動化に対応
HTML — 自己完結型のクライアント対応レポート
PoC バンドル - 再現可能なコピーアンドペーストのエクスプロイト例
✔ 重複排除済み • ✔ 重大度ランク付け • ✔ 安全に開ける (ペイロードはレポート内でエスケープされる)
破壊的テストはデフォルトではオフになっており、明示的にオプトインします。
すべての発見は、報告される前に検証チェックポイントを通過します。
ノイズとリスクを最小限に抑えるように設計されており、さまざまな壁を超えて再現可能な結果が得られます。
xss_grenade.py # コア スキャン エンジン (CLI + オーケストレーション)
xss_grenade_gui.py # PyQt5 デスクトップ GUI
context_engine.py # リフレクションコンテキストの検出
_static_js_analyzer.py # JavaScript ソース → シンク汚染分析
_dom_v6.py # DOM XSS (ヘッドレステイント)
_mutation_xss.py # ミューテーション XSS (mXSS)
_proto_pollution_analyzer.py # プロトタイプ汚染 → ガジェット チェーン
_dom_clobbering.py # DOM の破壊
_trusted_types_analyzer.py # 信頼できるタイプ / CSP 監査
_headless_verifier.py # リアルブラウザの確認 (Playwright)
_html_report.py # 自己完結型 HTML レポート
🧭ロードマップ
ヘッドレス検証の改善
AI 支援によるコンテキスト駆動型のペイロード生成
分散/並列スキャン
プルリクエストは大歓迎です。プロジェクトを鮮明かつ安全に保つには:
リスクのある機能はオプトイン切り替えの背後に置いてください。
新しい検出に対する回帰テストを含めます。
ペイロード数の虚栄心ではなく、現実世界の攻撃ベクトルに焦点を当てます。
GNU GPLv3 — 無料のオープンソース。変更して配布する場合は、公開したままにしておく必要があります。 「ライセンス」を参照してください。
認可されたセキュリティ テストのみ。このツールの使用方法については、お客様が単独で責任を負います。無断使用は違法となる可能性があります。
TX-C0RE セキュリティ調査 · github.com/tX-c0re
現実世界の Web アプリケーション用の最新の XSS 脆弱性スキャナー。
Readme GPL-3.0 ライセンス アクティビティスター
3 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フー

ターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Modern XSS vulnerability scanner for real-world web applications. - GitHub - tX-c0re/xss-grenade: Modern XSS vulnerability scanner for real-world web applications.

GitHub - tX-c0re/xss-grenade: Modern XSS vulnerability scanner for real-world web applications. · GitHub
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
tX-c0re
/
xss-grenade
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
assets assets bench_firingrange bench_firingrange .gitignore .gitignore LICENSE LICENSE README.md README.md _attack_graph_v2_source.py _attack_graph_v2_source.py _auth.py _auth.py _blind_xss_oob.py _blind_xss_oob.py _breakout_synth.py _breakout_synth.py _cache_poisoning.py _cache_poisoning.py _checkpoint.py _checkpoint.py _css_injection.py _css_injection.py _dom_clobbering.py _dom_clobbering.py _dom_exploit_engine.py _dom_exploit_engine.py _dom_v6.py _dom_v6.py _dompurify_config.py _dompurify_config.py _dompurify_cve_feed.py _dompurify_cve_feed.py _engine.py _engine.py _exploit_classifier.py _exploit_classifier.py _finding_scorer.py _finding_scorer.py _finding_store.py _finding_store.py _graphql_xss.py _graphql_xss.py _headless_crawler.py _headless_crawler.py _headless_verifier.py _headless_verifier.py _html_analyzer.py _html_analyzer.py _html_report.py _html_report.py _htmx_alpine.py _htmx_alpine.py _js_analyzer.py _js_analyzer.py _library_cve_feed.py _library_cve_feed.py _markdown_xss.py _markdown_xss.py _mutation_xss.py _mutation_xss.py _oob_collector.py _oob_collector.py _open_redirect.py _open_redirect.py _param_wordlist.py _param_wordlist.py _poc_generator.py _poc_generator.py _proto_pollution_analyzer.py _proto_pollution_analyzer.py _render_gate.py _render_gate.py _response_aware.py _response_aware.py _sarif_report.py _sarif_report.py _sourcemap_analyzer.py _sourcemap_analyzer.py _spa_route_extractor.py _spa_route_extractor.py _static_js_analyzer.py _static_js_analyzer.py _stored_xss_tracker.py _stored_xss_tracker.py _template_injection.py _template_injection.py _trusted_types_analyzer.py _trusted_types_analyzer.py _url_analyzer.py _url_analyzer.py benchmark_scoreboard.py benchmark_scoreboard.py context_engine.py context_engine.py destructive_real_test_server.py destructive_real_test_server.py dom_hooks_v6.js dom_hooks_v6.js head2head_dalfox.py head2head_dalfox.py payloads.txt payloads.txt profile_phases.py profile_phases.py repro_level2.py repro_level2.py re
[truncated]
Next-generation XSS detection engine for modern web applications.
Built for real-world security testing — not just payload fuzzing.
Authorized use only. Test only systems you own or have explicit written permission to test — an in-scope bug-bounty program or a signed engagement. Unauthorized scanning is illegal in most jurisdictions, and you are responsible for how you use this tool.
Most scanners stop at reflections — "your input came back in the page." XSS Grenade goes further and confirms real, exploitable vulnerabilities in a real browser.
🧠 Context-aware payload engine — HTML / attribute / URL / JS / CSS / comment
🌐 Built for modern apps — SPAs, React, Angular, Vue, SSR / hydration
🔍 Finds real attack vectors — not just echoed payloads
⚡ Headless-browser verification — cuts false positives to near zero
💣 Bug-bounty focused — the vectors that actually pay out
New to XSS? Cross-Site Scripting is when an app renders attacker-controlled input as active code in a victim's browser. A "reflection" only means your input appeared in the response — it isn't a bug until it actually executes . XSS Grenade does that last, hard step for you: it loads the page in a real headless Chromium and checks whether the payload fires .
Context-aware injection
Detects where a parameter lands and fires only the payloads that fit that context — including multiple contexts for the same parameter.
Real-browser verification
Candidates are re-loaded in headless Chromium (Playwright). Only findings that actually execute survive.
Static JS taint analysis
Parses JavaScript to an AST and traces untrusted sources ( location.* , document.referrer , window.name , postMessage ) into dangerous sinks ( innerHTML , eval , document.write , framework sinks).
Modern vulnerability classes
DOM XSS · mutation XSS (mXSS) · prototype pollution → XSS · DOM clobbering · Trusted Types misconfig · SSR hydration issues.
Bug-bounty vectors
postMessage abuse · JSONP callback injection · dangling markup (scriptless, CSP-resistant) · SVG/XML content-type reflection.
Known-CVE library detection
Fingerprints React, Vue, Angular, Next.js, jQuery, lodash, DOMPurify… and flags versions with known XSS/RCE CVEs.
🖥️ Interface preview
A live attack-surface graph, real-time severity-ranked findings, and browser-verified results — all from a single PyQt5 desktop app.
git clone https://github.com/tX-c0re/xss-grenade.git
cd xss-grenade
python3 -m venv venv
source venv/bin/activate # Windows: venv\Scripts\activate
pip install -r requirements.txt
python -m playwright install chromium # for headless verification
python xss_grenade_gui.py
Tip
No login? No problem — but ~80% of interesting XSS hides behind authentication. Paste your session cookies in SETTINGS → Authentication to scan admin panels, profiles and dashboards. Cookies stay in memory only — never written to reports or logs.
Enter the target URL in the top bar.
Pick your modules in the SETTINGS tab (or keep the sensible defaults).
Watch findings appear in real time — the attack-surface graph and the FINDINGS table update live.
Export the report: SAVE → JSON + HTML + PoC bundle.
✔ Findings are verified in a real browser • ✔ Output is exploitable vulnerabilities, not noise.
Enable “Smart Payloads” for a quick, high-signal first pass:
⚡ Cuts scan time from hours → minutes
🎯 Keeps high-signal, per-context coverage
💡 Ideal for first-pass recon before a deep run
Fully modular — enable only what you need.
One-click export, three formats:
JSON — structured, automation-ready
HTML — self-contained, client-ready report
PoC bundle — reproducible, copy-paste exploit examples
✔ Deduplicated • ✔ Severity-ranked • ✔ Safe to open (payloads escaped in the report)
Destructive testing is off by default — you opt in explicitly.
Every finding passes a verification checkpoint before it's reported.
Designed to minimise noise and risk — reproducible results over a wall of maybes.
xss_grenade.py # Core scan engine (CLI + orchestration)
xss_grenade_gui.py # PyQt5 desktop GUI
context_engine.py # Reflection-context detection
_static_js_analyzer.py # JavaScript source → sink taint analysis
_dom_v6.py # DOM XSS (headless taint)
_mutation_xss.py # Mutation XSS (mXSS)
_proto_pollution_analyzer.py # Prototype pollution → gadget chains
_dom_clobbering.py # DOM clobbering
_trusted_types_analyzer.py # Trusted Types / CSP audit
_headless_verifier.py # Real-browser confirmation (Playwright)
_html_report.py # Self-contained HTML report
🧭 Roadmap
Headless-verification improvements
AI-assisted, context-driven payload generation
Distributed / parallel scanning
Pull requests welcome. To keep the project sharp and safe:
Keep risky features behind an opt-in toggle .
Include regression tests for new detections.
Focus on real-world attack vectors , not payload-count vanity.
GNU GPLv3 — free and open source. If you modify and distribute it, you must keep it open. See LICENSE .
For authorized security testing only . You are solely responsible for how you use this tool; unauthorized use may be illegal.
TX-C0RE Security Research · github.com/tX-c0re
Modern XSS vulnerability scanner for real-world web applications.
Readme GPL-3.0 license Activity Stars
3 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
