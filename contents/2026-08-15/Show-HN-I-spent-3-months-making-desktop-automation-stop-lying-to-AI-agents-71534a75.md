---
source: "https://github.com/lahfir/agent-desktop/tree/main"
hn_url: "https://news.ycombinator.com/item?id=49307819"
title: "Show HN: I spent 3 months making desktop automation stop lying to AI agents"
article_title: "GitHub - lahfir/agent-desktop: Native desktop automation CLI for AI agents. Control any application through OS accessibility trees with structured JSON output and deterministic element refs. · GitHub"
author: "lahfir"
captured_at: "2026-08-15T05:16:32Z"
capture_tool: "hn-digest"
hn_id: 49307819
score: 1
comments: 0
posted_at: "2026-08-15T05:07:25Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I spent 3 months making desktop automation stop lying to AI agents

- HN: [49307819](https://news.ycombinator.com/item?id=49307819)
- Source: [github.com](https://github.com/lahfir/agent-desktop/tree/main)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T05:07:25Z

## Translation

タイトル: Show HN: デスクトップ オートメーションが AI エージェントに嘘をつくのをやめさせるのに 3 か月かかりました
記事のタイトル: GitHub - lahfir/agent-desktop: AI エージェント用のネイティブ デスクトップ オートメーション CLI。構造化された JSON 出力と決定的な要素参照を備えた OS アクセシビリティ ツリーを通じて、あらゆるアプリケーションを制御します。 · GitHub
説明: AI エージェント用のネイティブ Desktop Automation CLI。構造化された JSON 出力と決定的な要素参照を備えた OS アクセシビリティ ツリーを通じて、あらゆるアプリケーションを制御します。 - lahfir/エージェントデスクトップ
HN text: それは大胆な主張ですね。しかし、実際にコンピューターの使用を解決できたかもしれないと心から感じています (デモ: https://x.com/mdlahfir/status/2088109763783700827?s=20 ) コンテキストとして、私はデスクトップ アプリ用の自動化 CLI である Agent-desktop (Vercel Labs の Agent-browser からインスピレーションを得た) を構築しています。これは Playwright に似ていますが、ネイティブだけでなく、デスクトップ用だけでなく、Chromium アプリにも対応しています。信じてください、はい、アクセシビリティ ツリーが密集している Chromium アプリです。 MacOS は GA です。 Windows と Linux 向けのリリースがもうすぐ始まります。それで、どうやって解決したのでしょうか？基本的には相互運用性です。コンピューターの使用に関する最大の問題は、Playwright、エージェント ブラウザーなど、ブラウザーで使用するための信頼できるフレームワークがあるのに、デスクトップでは同じではないことです。 TryCua のような、本当に優れたソリューションが登場しており、私は大ファンです。エージェント デスクトップに関する私のビジョンは、エージェントが長期的なタスクに使用できる最も信頼性の高いフレームワークを構築することです。 Agent-desktop は軽量で、Rust 上に構築されており、高速で、トークンを大量に消費しません (コンテキスト ウィンドウを超えずに何時間も実行できます) これを可能にするために私が使用したアプローチは次のとおりです: a) スケルトン スナップショット - ウィンドウ/アプリのスナップショットを作成する場合、アクセシビリティ ツリー全体ではなく、親コンテナのスナップショットのみが作成され、ref ID が返されます。 b) スケルトンドリル - エージェントがそれを取得したら

ツリーにアクセスすると、特定の領域にドリルダウンすることを決定できます。 --find、--click、--wait... などのすべてのサブコマンドは、すべてその特定の参照対応領域で機能します。つまり、アプリ全体で要素を見つけたい場合、その要素をアプリ全体で永遠に検索する必要はありません。むしろ、エージェントは、トークンコストの一部でその要素がどこにあるかについて正確な手がかりを得ることができます。 c) 連鎖インタラクション フォールバック - シングル クリックは 1 つの API 呼び出しではなく、順序付けられたメカニズムの連鎖です (AXPress -> AXOpen -> 内部セルによるアクティブ化 -> 選択の書き込み -> AXconfirm)。各ステップは要素が通知する場合にのみ実行され、成功は戻りコードではなくアプリの状態変化を観察することによって判断されます。これは、アプリは両方向に存在するためです。Finder は、機能したアクションにはエラーを返し、何もしなかったアクションには成功を返します。最初に観察された効果が優先されます。応答では試行されたすべてのステップが報告されるため、エージェントはどのメカニズムが到達したかを正確に知ることができます。 d) アクション後のフィードバック - すべてのアクションは、その性質 (配信済みで検証済み、配信済みだが未検証、未配信) と開いたサーフェイス (ダイアログ、メニュー、シート) を報告するため、エージェントはアプリ全体を再スキャンすることなく、何が起こったのかを知ることができます。 e) 厳密な ref の再識別 - ref はポインタではありません。それはアイデンティティの証拠 (役割、パス、安定したテキスト、境界ハッシュ) です。すべてのアクションの前に、ライブ UI に対して再解決されます。 UI が変更された場合は、STALE_REF が返されます。 2 つの要素が一致すると、AMBIGUOUS_TARGET が返されます。それは決して推測ではありません。このすべてについて最も重要な部分は、Chromium アプリのアクセシビリティです。どうやってやったの？魔法の言葉はCDPです！現在、ほとんどのデスクトップ アプリは Chromium ベースです (Slack、VS Code、Obsidian、Discord など)。 1 つのコマンドは、エージェント デスクトップが実際に応答していることを確認する CDP エンドポイントを使用してアプリを起動します。

それを返す。そこから、Playwright、Puppeteer、エージェント ブラウザーなど、すでに使用しているあらゆるブラウザー自動化フレームワークを接続して Web コンテンツを駆動できます。 CDP 経由で Obsidian の Web コンテンツを読み取るのにかかる時間は 201 ミリ秒ですが、アクセシビリティ ツリーでは 2.3 秒かかります。 $ エージェントデスクトップ起動 "Obsidian" --cdp { "ok": true, "data": {
"レンダラー": "クロム",
"cdp": { "ポート": 57500,
"http_endpoint": "http://127.0.0.1:57500",
"websocket_url": "ws://127.0.0.1:57500/devtools/browser/..." },
"suggestion": "次: `agent-browser connect 57500` を実行します ..." } }
これにより、エージェント デスクトップがブラウザ自動化エコシステムと競合するのではなく、ブラウザ自動化エコシステム全体と相互運用できるようになります。エージェントデスクトップを試してみてください -> https://github.com/lahfir/agent-desktop

記事本文:
GitHub - lahfir/agent-desktop: AI エージェント用のネイティブ デスクトップ オートメーション CLI。構造化された JSON 出力と決定的な要素参照を備えた OS アクセシビリティ ツリーを通じて、あらゆるアプリケーションを制御します。 · GitHub
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
ラフィール
/
エージェントデスクトップ
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
179 コミット 179 コミット .githooks .githooks

.github .github アセット アセット ベンチマーク/ ロケーター解像度ベンチマーク/ ロケーター解像度 クレート クレート ドキュメント ドキュメント npm npm スクリプト スクリプト スキル スキル src src テスト テスト .gitignore .gitignore .gitleaks.toml .gitleaks.toml .release-please-manifest.json .release-please-manifest.json CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONCEPTS.md CONCEPTS.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md Clippy.toml Clippy.toml Deny.toml Deny.toml release-please-config.json release-please-config.json Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Agent-desktop は、Rust で構築された AI エージェント用に設計されたネイティブのデスクトップ オートメーション CLI です。 OS アクセシビリティ ツリーを通じて、あらゆるアプリケーションへの構造化されたアクセスが提供されます。スクリーンショット、ピクセル マッチング、ブラウザは必要ありません。
ネイティブ Rust CLI : 高速、単一バイナリ、ランタイム依存関係なし
C-ABI cdylib ( libagent_desktop_ffi ): 呼び出しごとに CLI をフォークする代わりに、Python / Swift / Go / Ruby / Node / C から 1 回ロードします。
58 個のコマンド名、54 個の操作コマンド: 観察、対話、キーボード、マウス、通知、クリップボード、ウィンドウ管理、セッション ライフサイクル、トレース読み取り/エクスポート、およびバンドルされたスキル ドキュメント ローダー。 4 つの保持入力名はステートフル デーモン用に予約されており、ステートレス CLI でフェールクローズされます。
プログレッシブ スケルトン トラバーサル: 浅い概要 + ターゲットを絞ったドリルダウンにより、高密度アプリで 78 ～ 96% のトークン削減
スナップショットと参照: コンパクトなスナップショット ID と修飾された要素参照を使用した AI に最適化されたワークフロー ( @s8f3k2p9:e1 、 @s8f3k2p9:e2 )
デフォルトでヘッドレスのインタラクション: Ref アクションはアクセシビリティ API を使用し、サイレント フォーカス、curs をブロックします。

または、キーボードまたはペーストボードの副作用
構造化された JSON 出力: エラー コードと回復ヒントを含む機械可読応答
あらゆるアプリで動作します: Finder、Safari、システム設定、Xcode、Slack — アクセシビリティ ツリーのあるものなら何でも
CDP を介した Chromium アプリの相互運用: launch --cdp は検証済みの DevTools ポートを開くため、CDP を使用するあらゆるフレームワーク (Playwright、Puppeteer、chrome-remote-interface 、agent-browser) が Web コンテンツを駆動できる一方で、ネイティブ メニュー、ダイアログ、ウィンドウはアクセシビリティ パス上に留まります。
npm install -g Agent-desktop # 事前に構築されたバイナリを自動的にダウンロードします
またはインストールせずに:
npx エージェント-デスクトップ スナップショット --app Finder -i
ソースから
git clone https://github.com/lahfir/agent-desktop
cd エージェント-デスクトップ
カーゴビルド --release
cp target/release/agent-desktop /usr/local/bin/
Rust 1.89 以降および macOS 13.0 以降が必要です。
macOS にはアクセシビリティ権限が必要です。スクリーンショットには画面録画権限も必要で、通知センター オープナーにはシステム イベントの自動化権限が必要です。単純な権限チェックではプロンプトが表示されません。以下を使用して、境界付き分離ヘルパーで欠落しているアクセス許可を要求します。
Agent-desktop Permissions --request # 分離されたヘルパーで不足している権限をリクエストします
権限フィールドは明示的なオブジェクトです。次に例を示します。
{
"アクセシビリティ" : { "状態" : " 付与 " },
"screen_recording" : { "state" : "拒否" , "suggestion" : "画面録画の許可を与える" },
"自動化" : { "状態" : " 不明 " }
}
自動化レポートが許可、拒否、または不明。不明 は、macOS でプロンプトが必要であるか、プロンプトなしではシステム イベントを調査できないことを意味します。
すべての GitHub リリースには、CLI tarball とともに、macOS、Linux、および Windows 用の事前構築済み C-ABI cdylib ( libagent_desktop_ffi ) が同梱されています。それをdl開いて、agent_desktop.hで宣言された関数を呼び出します。

n-process は、コマンドごとに fork-exec の代わりに呼び出します。
ctypesをインポートする
lib = ctypes 。 CDLL ( "./lib/libagent_desktop_ffi.dylib" )
ライブラリ 。 ad_init ( 3 ) # 呼び出しの前に ABI メジャー (AD_ABI_VERSION_MAJOR) を確認します
アダプター = lib 。 ad_adapter_create()
# 観察 -> 動作: ad_snapshot -> 修飾された ref を解析 -> ad_execute_by_ref ...
ライブラリ 。 ad_adapter_destroy (アダプター)
完全な消費者ガイド — エントリポイント、所有権、スレッド化、エラー処理、ビルド/リンク、リリース アーカイブ、および検証: skill/agent-desktop-ffi/ 。
高密度アプリ (Slack、VS Code、Notion) の場合は、プログレッシブ スケルトン トラバーサルを使用してトークンの使用量を最小限に抑えます。
# 1. 浅い概要 — 深さ 3 のマップ、切り詰められたコンテナーに Children_count が表示される
エージェントデスクトップスナップショット --skeleton --app Slack -i --compact
# snapshot_id を保持します (例: s8f3k2p9)
# 2. 対象領域にドリルします (名前付きコンテナはドリル ターゲットとして参照を取得します)
エージェントデスクトップスナップショット --root @e3 --snapshot s8f3k2p9 -i --compact
# 3. ドリルダウンで見つかった要素に基づいて動作する
エージェントデスクトップクリック @e12 --snapshot s8f3k2p9
# 4. 同じ領域を再ドリルして状態の変化を確認する
エージェントデスクトップスナップショット --root @e3 --snapshot s8f3k2p9 -i --compact
単純なアプリの場合は、完全なスナップショットで問題ありません。
Agent-desktop snapshot --app Finder -i # refs と snapshot_id を使用してインタラクティブな要素を取得します
Agent-desktop click @e3 --snapshot s8f3k2p9 # ref でボタンをクリックします
Agent-desktop type @e5 --snapshot s8f3k2p9 " 四半期レポート " # フィールドにテキストを挿入
エージェントデスクトップを押す cmd+s # キーボードショートカット
Agent-desktop snapshot -i # UI 変更後に再観察
エージェント ループ: スナップショット → 決定 → 行動 → スナップショット → 決定 → 行動 → ...
トレース ビューア (セッションを読み戻す)
session_id= $( エージェントデスクトップセッション開始 --スクリーンショット | jq -r ' .data.session_id ' )
エクスポート AGENT_DESKTOP_SESSION= " $session_id "
年齢

t-desktop snapshot --app Finder -i # 明示的なセッション スコープ内で動作します
エージェントデスクトップクリック @s8f3k2p9:e5
エージェントデスクトップトレースショー --limit 500 # エージェントの制限された JSON タイムライン
エージェントデスクトップトレースエクスポート --out run.html # 単一ファイルの HTML ビューア (file:// から動作)
トレース ショーは、すべてのセグメント ファイルを決定論的にマージし、権限を必要としません。トレース エクスポートは、タイムラインとスクリーンショットを 1 つの静的 HTML ファイルに Base64 として埋め込みます。 --out を指定しないと、HTML は現在のディレクトリではなくセッション ディレクトリ ( ~/.agent-desktop/sessions/<id>/trace-<id>.html ) に書き込まれます。 --out はパスをオーバーライドします。 artifacts: full が有効になっている場合、エクスポートされた HTML をスクリーンショットのように扱います。
マルチエージェントワークフローの共有セッション
エージェントの実行ごとにセッション開始を 1 回実行して、トレースが有効なセッション (マニフェスト トレース: デフォルトでオン) を作成し、返された ID を global --session <id> または AGENT_DESKTOP_SESSION=<id> で渡します。その明示的なスコープ内のコマンドは、~/.agent-desktop/sessions/<id>/trace/ の下にある自動 JSONL セグメントを取得し、セッションの最新スナップショット名前空間を共有します (呼び出しごとに --trace はありません)。
独立したエージェントを同時に実行するには、プロセスごとに AGENT_DESKTOP_SESSION=<id> を設定します。複数のエージェントが 1 つのセッション ID を共有する場合、各エージェントは、名前空間の最新のスナップショットが変更されていないと想定するのではなく、独自のスナップショット呼び出しからの修飾された参照に基づいて動作する必要があります。
マニフェストなし ( session start なし) の裸の --session <id> は、依然としてスナップショット名前空間のみをスコープし、トレース ファイルは書き込みません。スナップショット ID は、選択したセッション名前空間内でのみ解決されます。セッション間の検索は決してトリガーされません。
エージェントデスクトップセッション開始 --name release-fix # note data.session_id
エクスポート AGENT_DESKTOP_SESSION= < session_id >
エージェントデスクトップスナップショット --app Xcode -i --compact # 選択されたセッション + トレースを使用します
エージェントデスクトップは

it --element @s8f3k2p9:e9 --predicate actionable --timeout 5000
エージェントデスクトップクリック @s8f3k2p9:e9
エージェントデスクトップクリック @e9 --snapshot s2 # 従来のベア参照、明示的に固定
エージェント デスクトップ セッション終了 " $AGENT_DESKTOP_SESSION "
エージェントデスクトップセッションGC
Chromium アプリの駆動 (CDP)
Chromium ベースのアプリ (Slack、VS Code、Discord、Obsidian、Notion) の場合、 --cdp を起動すると、Web コンテンツ上で検証済みの Chrome DevTools プロトコル エンドポイントが開きます。 CDP を使用するフレームワークであれば、Playwright、Puppeteer、chrome-remote-interface、agent-browser を接続できます。ref ベースのエージェント ワークフローとバンドルされた Electron スキルで推奨されるエージェント ブラウザと接続できます。ネイティブ サーフェス (メニュー、ダイアログ、ウィンドウ、スクリーンショット) は、どちらの場合でもアクセシビリティ パス上に残ります。
エージェントデスクトップ起動「Obsidian」 --cdp
{ "アプリ" : "黒曜石" 、 "pid" : 4821 、 "cdp" : {
「ポート」: 9229 、
"http_endpoint" : " http://127.0.0.1:9229 " 、
"websocket_url" : " ws://127.0.0.1:9229/devtools/browser/<id> " ,
"製品" : " Chrome/142.0.7444.265 "
}、
"suggestion" : " 次に: `agent-browser connect 9229` を実行するか (推奨)、Playwright や Puppeteer などの CDP クライアントに接続します。 ... " }
--cdp には新たな起動が必要です。すでに実行中のターゲットは ACTION_FAILED を返します。最初に close-app を実行し、プロセスが終了したことを確認してから、再度 --cdp を起動します。
エンドポイントは 127.0.0.1 に固定されます。ユーザーとして実行されているローカル プロセスは、開いたままでもアクセスできます。 close-app はアプリとともに露出を終了します。
エージェントデスクトップスナップショット --app Safari -i # 参照を含むアクセシビリティツリー
エージェントデスクトップスナップショット --surface メニュー # オープンメニューをキャプチャ
エージェントデスクトップのスクリーンショット --app Finder # PNG スクリーンショット
Agent-desktop find --role button --app TextEdit # ロール、名前、値、テキストで検索
Agent-desktop get @e3 --snapshot s8f3k2p9 --property value # 要素プロパティの読み取り
ああ

ent-desktop は @e7 --snapshot s8f3k2p9 --property selected # ブール状態を確認します
Agent-desktop list-surfaces --app Notes # リスト メニュー、シート、ポップオーバー、アラート
get および is は参照を一度解決し、利用可能な場合はライブ プラットフォーム読み取りを優先し、そのライブ読み取りがアダプターでサポートされていない場合にのみフォールバックします。
エージェントデスクトップクリック @s8f3k2p9:e3 # 厳密なヘッドレス AX クリック
Agent-desktop --headed click @s8f3k2p9:e3 # 物理的なクリック、フォーカス/カーソルは許可されます
Agent-desktop --headed double-click @s8f3k2p9:e3 # 物理的なダブルクリック
Agent-desktop --headed トリプルクリック @s8f3k2p9:e3 # 物理的なトリプルクリック
エージェントデスクトップを右クリック @s8f3k2p9:e3 # コンテキスト メニューを開きます。再試行する前に効果を検査する
Agent-desktop type @s8f3k2p9:e5 " hello world " # 要素にテキストを挿入
Agent-desktop set-value @s8f3k2p9:e5 " new value " # AX 経由で値を直接設定
Agent-desktop clear @s8f3k2p9:e5 # 要素の値をクリア
Agent-desktop focus @s8f3k2p9:e5 # キーボード フォーカスを設定
Agent-desktop select @s8f3k2p9:e9 " Option B " # 検証済みのドロップダウン/リスト オプションを選択
Agent-desktop toggle @s8f3k2p9:e12 # チェックボックスまたはスイッチを反転します
エージェントデスクトップチェック @s8f3k2p9:e12 # 冪等チェック
エージェントデスクトップのチェックを外します @s8f3k2p9:e12 # 冪等のチェックを外します
エージェントデスクトップ展開 @s8f3k2p9:e15 # 展開

[切り捨てられた]

## Original Extract

Native desktop automation CLI for AI agents. Control any application through OS accessibility trees with structured JSON output and deterministic element refs. - lahfir/agent-desktop

That's a bold claim. But I genuinely feel like I might have actually solved computer use (demo: https://x.com/mdlahfir/status/2088109763783700827?s=20 ) For context, I've been building agent-desktop (Inspired by agent-browser by Vercel Labs), an automation CLI for desktop apps. It's like Playwright but for desktops, not just native, but for Chromium apps as well. Trust me, yes, Chromium apps whose accessibility tree is dense. MacOS is GA; I'm almost close to launching for Windows and Linux! So, how did I solve it? Basically interoperability. The biggest issue with computer use is that we have reliable frameworks for browser use, like Playwright, agent-browser, and many more, but not the same with desktops. We have really good solutions emerging, like tryCua, which I'm a big fan of. My vision with agent-desktop is to build the most reliable framework that agents can use for long-horizon tasks. agent-desktop is lightweight, built on Rust, fast, and not token-hungry (It can go for hours without exceeding the context window) Here's the approach I used to make it possible: a) skeleton snapshots - when you want to snapshot a window/app, it only snapshots the parent containers and gives back a ref id, not the entire accessibility tree. b) skeleton drilling - once the agent has that tree, it can then decide to drill into a specific region. All the subcommands like --find, --click, --wait... all work on that specific ref aware region. Meaning if you want to find an element in the entire app, it doesn't take forever searching for the entire app for that element; rather, the agent will have an exact clue on where that element might be for a fraction of the token costs. c) chained interaction fallback - a single click isn't one API call but it's an ordered chain of mechanisms (AXPress -> AXOpen -> activate through the inner cell -> write selection -> AXConfirm). Each step only runs if the element advertises it, and success is judged by watching the app's state change, not by the return code, because apps lie in both directions: Finder returns an error for an action that worked and success for one that did nothing. First observed effect wins. The response reports every step tried, so the agent knows exactly which mechanism landed. d) after-action feedback - every action reports its disposition (delivered and verified, delivered but unverified, not delivered) plus any surface that opened (dialog, menu, sheet), so the agent knows what happened without re-scanning the whole app. e) strict ref re-identification - a ref isn't a pointer; it's identity evidence (role, path, stable text, bounds hash). Before every action, it's re-resolved against the live UI. If the UI changed, you get STALE_REF; if two elements now match, you get AMBIGUOUS_TARGET. It never guesses. The most important part about all this is Chromium app accessibility. How did I do it? The magic word is CDP! Most desktop apps today are Chromium-based (Slack, VS Code, Obsidian, Discord...). One command launches the app with a CDP endpoint that agent-desktop verifies is actually answering before returning it. From there, any browser automation framework can connect and drive the web contents: Playwright, Puppeteer, agent-browser, whatever you already use. Reading Obsidian's web content over CDP takes 201ms vs 2.3s through the accessibility tree. $ agent-desktop launch "Obsidian" --cdp { "ok": true, "data": {
"renderer": "chromium",
"cdp": { "port": 57500,
"http_endpoint": "http://127.0.0.1:57500",
"websocket_url": "ws://127.0.0.1:57500/devtools/browser/..." },
"suggestion": "Next: run `agent-browser connect 57500` ..." } }
This is what makes agent-desktop interoperable with the entire browser automation ecosystem instead of competing with it. Go try agent-desktop -> https://github.com/lahfir/agent-desktop

GitHub - lahfir/agent-desktop: Native desktop automation CLI for AI agents. Control any application through OS accessibility trees with structured JSON output and deterministic element refs. · GitHub
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
lahfir
/
agent-desktop
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
179 Commits 179 Commits .githooks .githooks .github .github assets assets benchmarks/ locator-resolution benchmarks/ locator-resolution crates crates docs docs npm npm scripts scripts skills skills src src tests tests .gitignore .gitignore .gitleaks.toml .gitleaks.toml .release-please-manifest.json .release-please-manifest.json CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONCEPTS.md CONCEPTS.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md clippy.toml clippy.toml deny.toml deny.toml release-please-config.json release-please-config.json rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
agent-desktop is a native desktop automation CLI designed for AI agents, built with Rust. It gives structured access to any application through OS accessibility trees — no screenshots, no pixel matching, no browser required.
Native Rust CLI : Fast, single binary, no runtime dependencies
C-ABI cdylib ( libagent_desktop_ffi ): Load once from Python / Swift / Go / Ruby / Node / C instead of forking the CLI per call
58 command names, 54 operational commands : Observation, interaction, keyboard, mouse, notifications, clipboard, window management, session lifecycle, trace read/export, plus a bundled skills doc loader. The four held-input names are reserved for a stateful daemon and fail closed in the stateless CLI.
Progressive skeleton traversal : 78–96% token reduction on dense apps via shallow overview + targeted drill-down
Snapshot & refs : AI-optimized workflow using compact snapshot IDs and qualified element references ( @s8f3k2p9:e1 , @s8f3k2p9:e2 )
Headless-by-default interactions : Ref actions use accessibility APIs and block silent focus, cursor, keyboard, or pasteboard side effects
Structured JSON output : Machine-readable responses with error codes and recovery hints
Works with any app : Finder, Safari, System Settings, Xcode, Slack — anything with an accessibility tree
Chromium-app interop via CDP : launch --cdp opens a verified DevTools port so any framework that speaks CDP — Playwright, Puppeteer, chrome-remote-interface , agent-browser — can drive the web contents, while native menus, dialogs, and windows stay on the accessibility path
npm install -g agent-desktop # downloads prebuilt binary automatically
Or without installing:
npx agent-desktop snapshot --app Finder -i
From source
git clone https://github.com/lahfir/agent-desktop
cd agent-desktop
cargo build --release
cp target/release/agent-desktop /usr/local/bin/
Requires Rust 1.89+ and macOS 13.0+.
macOS requires Accessibility permission. Screenshots also require Screen Recording permission, and the Notification Center opener requires Automation permission for System Events. Plain permission checks never prompt. Request missing permissions in a bounded isolated helper with:
agent-desktop permissions --request # request missing permissions in an isolated helper
Permission fields are explicit objects, for example:
{
"accessibility" : { "state" : " granted " },
"screen_recording" : { "state" : " denied " , "suggestion" : " Grant Screen Recording permission " },
"automation" : { "state" : " unknown " }
}
Automation reports granted , denied , or unknown ; unknown means macOS would need to prompt or System Events could not be probed without prompting.
Every GitHub Release ships a prebuilt C-ABI cdylib ( libagent_desktop_ffi ) for macOS, Linux, and Windows alongside the CLI tarballs. dlopen it and call the functions declared in agent_desktop.h for in-process calls instead of fork-exec per command.
import ctypes
lib = ctypes . CDLL ( "./lib/libagent_desktop_ffi.dylib" )
lib . ad_init ( 3 ) # verify ABI major (AD_ABI_VERSION_MAJOR) before any call
adapter = lib . ad_adapter_create ()
# observe -> act: ad_snapshot -> parse a qualified ref -> ad_execute_by_ref ...
lib . ad_adapter_destroy ( adapter )
Full consumer guide — entrypoints, ownership, threading, error-handling, build/link, release archives, and verification: skills/agent-desktop-ffi/ .
For dense apps (Slack, VS Code, Notion), use progressive skeleton traversal to minimize token usage:
# 1. Shallow overview — depth-3 map, truncated containers show children_count
agent-desktop snapshot --skeleton --app Slack -i --compact
# Keep snapshot_id, for example s8f3k2p9
# 2. Drill into a region of interest (named containers get refs as drill targets)
agent-desktop snapshot --root @e3 --snapshot s8f3k2p9 -i --compact
# 3. Act on an element found in the drill-down
agent-desktop click @e12 --snapshot s8f3k2p9
# 4. Re-drill the same region to verify the state change
agent-desktop snapshot --root @e3 --snapshot s8f3k2p9 -i --compact
For simple apps, a full snapshot is fine:
agent-desktop snapshot --app Finder -i # get interactive elements with refs and snapshot_id
agent-desktop click @e3 --snapshot s8f3k2p9 # click a button by ref
agent-desktop type @e5 --snapshot s8f3k2p9 " quarterly report " # insert text into a field
agent-desktop press cmd+s # keyboard shortcut
agent-desktop snapshot -i # re-observe after UI changes
Agent loop: snapshot → decide → act → snapshot → decide → act → ...
Trace viewer (read back a session)
session_id= $( agent-desktop session start --screenshots | jq -r ' .data.session_id ' )
export AGENT_DESKTOP_SESSION= " $session_id "
agent-desktop snapshot --app Finder -i # work inside the explicit session scope
agent-desktop click @s8f3k2p9:e5
agent-desktop trace show --limit 500 # bounded JSON timeline for agents
agent-desktop trace export --out run.html # single-file HTML viewer (works from file://)
trace show merges all segment files deterministically and requires no permissions. trace export embeds the timeline plus screenshots as base64 in one static HTML file. Without --out , the HTML is written to the session directory ( ~/.agent-desktop/sessions/<id>/trace-<id>.html ), not the current directory; --out overrides the path. Treat exported HTML like a screenshot when artifacts: full was enabled.
Shared sessions for multi-agent workflows
Run session start once per agent run to create a trace-enabled session (manifest trace: on by default), then pass the returned ID with global --session <id> or AGENT_DESKTOP_SESSION=<id> . Commands in that explicit scope get automatic JSONL segments under ~/.agent-desktop/sessions/<id>/trace/ and share the session's latest-snapshot namespace — no --trace on every call.
For concurrent independent agents, set AGENT_DESKTOP_SESSION=<id> per process. When multiple agents share one session ID, each agent should act on the qualified refs from its own snapshot call rather than assuming the namespace's latest snapshot is unchanged.
Bare --session <id> without a manifest (no session start ) still scopes the snapshot namespace only and writes no trace files. Snapshot IDs resolve only inside the selected session namespace; they never trigger a cross-session search.
agent-desktop session start --name release-fix # note data.session_id
export AGENT_DESKTOP_SESSION= < session_id >
agent-desktop snapshot --app Xcode -i --compact # uses selected session + tracing
agent-desktop wait --element @s8f3k2p9:e9 --predicate actionable --timeout 5000
agent-desktop click @s8f3k2p9:e9
agent-desktop click @e9 --snapshot s2 # legacy bare ref, explicitly pinned
agent-desktop session end " $AGENT_DESKTOP_SESSION "
agent-desktop session gc
Driving Chromium apps (CDP)
For a Chromium-based app (Slack, VS Code, Discord, Obsidian, Notion), launch --cdp opens a verified Chrome DevTools Protocol endpoint on the web contents. Any framework that speaks CDP can connect — Playwright, Puppeteer, chrome-remote-interface , agent-browser — with agent-browser preferred for its ref-based agent workflow and bundled electron skill. Native surfaces (menus, dialogs, windows, screenshots) stay on the accessibility path either way.
agent-desktop launch " Obsidian " --cdp
{ "app" : " Obsidian " , "pid" : 4821 , "cdp" : {
"port" : 9229 ,
"http_endpoint" : " http://127.0.0.1:9229 " ,
"websocket_url" : " ws://127.0.0.1:9229/devtools/browser/<id> " ,
"product" : " Chrome/142.0.7444.265 "
},
"suggestion" : " Next: run `agent-browser connect 9229` (preferred) or connect any CDP client such as Playwright or Puppeteer. ... " }
--cdp needs a fresh launch — an already-running target returns ACTION_FAILED . Run close-app first, confirm the process exited, then launch --cdp again.
The endpoint is pinned to 127.0.0.1 . Any local process running as your user can still reach it while it stays open; close-app ends the exposure along with the app.
agent-desktop snapshot --app Safari -i # accessibility tree with refs
agent-desktop snapshot --surface menu # capture open menu
agent-desktop screenshot --app Finder # PNG screenshot
agent-desktop find --role button --app TextEdit # search by role, name, value, text
agent-desktop get @e3 --snapshot s8f3k2p9 --property value # read element property
agent-desktop is @e7 --snapshot s8f3k2p9 --property checked # check boolean state
agent-desktop list-surfaces --app Notes # list menus, sheets, popovers, alerts
get and is resolve the ref once, prefer live platform reads when available, and fall back only when that live read is unsupported by the adapter.
agent-desktop click @s8f3k2p9:e3 # strict headless AX click
agent-desktop --headed click @s8f3k2p9:e3 # physical click, focus/cursor allowed
agent-desktop --headed double-click @s8f3k2p9:e3 # physical double-click
agent-desktop --headed triple-click @s8f3k2p9:e3 # physical triple-click
agent-desktop right-click @s8f3k2p9:e3 # open context menu; inspect effect before retrying
agent-desktop type @s8f3k2p9:e5 " hello world " # insert text into element
agent-desktop set-value @s8f3k2p9:e5 " new value " # set value directly via AX
agent-desktop clear @s8f3k2p9:e5 # clear element value
agent-desktop focus @s8f3k2p9:e5 # set keyboard focus
agent-desktop select @s8f3k2p9:e9 " Option B " # select verified dropdown/list option
agent-desktop toggle @s8f3k2p9:e12 # flip checkbox or switch
agent-desktop check @s8f3k2p9:e12 # idempotent check
agent-desktop uncheck @s8f3k2p9:e12 # idempotent uncheck
agent-desktop expand @s8f3k2p9:e15 # expand di

[truncated]
