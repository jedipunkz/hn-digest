---
source: "https://github.com/SHIN620265/claude-pet"
hn_url: "https://news.ycombinator.com/item?id=48832902"
title: "Show HN: A Windows Desktop Pet for Claude Code, in PowerShell/Win32"
article_title: "GitHub - SHIN620265/claude-pet: Desktop companion pet for Claude Code (Windows) · GitHub"
author: "shin620265"
captured_at: "2026-07-08T15:10:05Z"
capture_tool: "hn-digest"
hn_id: 48832902
score: 1
comments: 0
posted_at: "2026-07-08T15:03:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A Windows Desktop Pet for Claude Code, in PowerShell/Win32

- HN: [48832902](https://news.ycombinator.com/item?id=48832902)
- Source: [github.com](https://github.com/SHIN620265/claude-pet)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T15:03:06Z

## Translation

タイトル: Show HN: PowerShell/Win32 のクロード コード用の Windows デスクトップ ペット
記事タイトル: GitHub - SHIN620265/claude-pet: クロード コードのデスクトップ コンパニオン ペット (Windows) · GitHub
説明: クロード コードのデスクトップ コンパニオン ペット (Windows)。 GitHub でアカウントを作成して、SHIN620265/claude-pet の開発に貢献してください。
HN テキスト: Codex からアイデアを得ました。終了時に思い出させる小さなペットが付いています。気に入りました。 Windows 上の Claude Code にはこれがなかったので、Claude Code 用に同じ種類のデスクトップ ペットを作成しました。その際、自分で設定したルールが 1 つあります。それは、Electron なし、バンドルされたランタイムなしというものです。私はターミナルでクロードコードを毎日使用しています。終了する瞬間や、停止してコマンドの承認を待つ瞬間を見逃してしまうことがよくあります。端末は通知しません。ペットはフックを通して私のクロード コード セッションを監視し、セッションごとに 1 枚のカード (「考え中」、「あなたを待っている」、または「完了」) を示します。セッションがあなたを必要とするとき、柔らかい音を奏でます。クロード コードが許可を求めると、ペットは組み込みの通知ではなく PermissionRequest フックをリッスンするため、カードには約 1 秒で許可が表示されます。約 200KB がインストールされており、ネットワーク要求は行われません。今のところ Windows のみです。裏話と、解決できなかった 1 つの問題を以下のコメントに記載しました。

記事本文:
GitHub - SHIN620265/claude-pet: クロード コードのデスクトップ コンパニオン ペット (Windows) · GitHub
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
SHIN620265
/
クロードペット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブランチタグ

s ファイルコードに移動 その他のアクションメニューを開く フォルダーとファイル
42 コミット 42 コミット .claude-plugin .claude-plugin docs docs pet pet scripts scripts tools tools vscode-ext/ claude-pet-jump vscode-ext/ claude-pet-jump .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md ライセンス ライセンス README.md README.md REVIEW-PROMPT.md REVIEW-PROMPT.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード デスクトップ ペット · クロード デスクトップ ペット · クロード デスクトップ ペット
🌸 デスクトップ上に住んでいて、クロード コードの進行状況を追跡するのに役立つ小さなペットです。
クロード コードの進捗を監視する小さなデスクトップ マスコットです。
クロード・コードの人の限界を见愿る、デスクトップの小さなマスコット。
↑実際の画面録画 — モックアップではありません・中国語と日本語のインターフェースは右クリックで瞬時に切り替え可能
言語/言語/言語：中国語・英語・日本語
⚠️ 非公式コミュニティ プロジェクト、Anthropic とは提携していません / 非公式コミュニティ プロジェクト、Anthropic とは提携していません / 抜きのコミュニティシステム。人間とは関係がありません
コードチェンジャー向け / 開発者 / 開発者向け → ARCHITECTURE.md
Windows デスクトップに住む小さなペット。ユーザーが実行しているクロード コードの会話を見つめ、カードと光のプロンプトを使用して、クロードが考えているのか、ユーザーの選択を待っているのか、それとも選択を終えたのかを表示します。そのため、コマンド ライン ウィンドウを見つめ続ける必要はありません。権限を確認する必要がある場合は、約 1 秒以内に通知されます。Claude Code の組み込み通知の 6 秒の遅延を待つ必要はありません。また、非常に軽量です。ローカルにインストールされるプラグインは約 200 KB (純粋な PowerShell スクリプト + サウンド エフェクト レンダリング)、Electron やバンドルされたランタイムはなく、Windows に付属のインターフェイス コンポーネントを駆動する PowerShell だけです。
必要なもの: Windows ・ Claude Code (最新バージョンを維持してください - インスタント許可リマインダーは新しいフッキング機能に依存しています) ・ PowerShell 7 (コマンド pwsh; そうでない場合は、winget install Microsoft.Po)

werShell、または Microsoft Store で「PowerShell」を検索してインストールします)
クロード コードに 2 つのコマンドを入力します。クロードは GitHub から直接コードを取得します。手動でダウンロードする必要はありません。
/プラグインマーケットプレイス追加 shin620265/claude-pet
/プラグインのインストール claude-pet@shin620265
次に、/reload-plugins を実行して、現在のセッションで有効にします (この手順は初回インストールにも必要です。再起動したり、ターミナル ラベルを失ったりする必要はありません。または、新しいクロード セッションを開くと、自動的に有効になります)。
終わり！ /my-pet を使用してペットを開きます (後で記憶され、新しいセッションが自動的に表示されます)。
アンインストール: /plugin uninstall claude-pet@shin620265 (完全にクリーンアップしたい場合は、~/.claude/pet-data フォルダーを削除してください。このフォルダーには設定とカード メモリが保存されています)
ウェアハウスをローカルにクローン/ダウンロードしましたか?最初のものを /plugin Marketplace add <ローカル フォルダー パス> に置き換え、残りは変更しないでください。
会話が実行されている場合、ステータス カード (最大 3 枚) がその横に積み重ねられ、各カードには次の内容が表示されます。
タイトル = その会話で最初に言うこと。
ステータス行: 考え中/確認が必要な場合 = 現在のステータス + 最新の文章入力。完了後 = モデル名 · 完了 · 応答の最初の文の要約 (モデル名はこの応答を生成したモデルを指し、/モデルは切り替え後の次の応答で自動的に更新されます)、色とアイコン:
🔵 考え中 (丸で囲む)・ 🟡 確認/選択が必要 (!、自動的に上部に固定されます)・ 🟢 完了 (✓)・ ⚪ アイドル・ ⚫ 中断 (Esc で現在のラウンドを中断)
オン/オフ: クロード コードに /my-pet と入力します (これは記憶されます。次回クロードを開いたときに自動的に表示され、すべてオフにすると消えます)。
プロンプトを見てください: 通常は心配する必要はありません。クロードが完了するか選択する必要がある場合、音が鳴り、それに応じてカードが点灯します。
ジャンプオーバー: カードをクリック = セッションが配置されているウィンドウを最前面に移動します (ホバーするとサンゴの輪が表示されます。配置できない場合、カードは首を左右に振りますが、間違ったウィンドウにジャンプすることはありません)。オプションの VS Code コンパニオン拡張機能がインストールされている場合、同じクリックでセッションのターミナル タブに直接移動できます。
操作: ドラッグ=移動(メモリ位置);ダブル

クリック = カードを折りたたむ/展開する;各行 ✎ = タイトルを変更、× = このカードを閉じます (新しいアクティビティがあると自動的に戻ります)。右クリック = メニュー (ペットを閉じる/デフォルトの位置に戻る/プロンプト音/プライバシー モード/言語)。
言語: 右クリック→「言語」→中国語/英語/日本語/自動 (自動 = システムに従います)、クリックすると即座に切り替わります。
プライバシーモード：右クリック→「プライバシーモード」のオン/オフ（デフォルトはオフ）。開くと、カードにはプロジェクト フォルダー名のみが表示され、入力テキストと返信テキストは非表示になります。画面/録画/ライブ ブロードキャストを共有する前に、ワンクリックで開きます。完全に非表示にしたい場合→ダブルクリックで折りたたんですべて非表示にします。
自動的に静かになる: 全画面でゲームやプレゼンテーションをプレイしているとき、または「フォーカス アシスタント/おやすみモード」をオンにしているときは音は鳴りません (カードは通常どおり更新されます)。
アニメーションを減らす：「設定」→「アクセシビリティ」→「視覚効果」→「アニメーション効果」で「アニメーション効果」をオフにすると、ペットは浮いたり点滅したり、円を描いて静止したりしなくなります。
Claude Code で、次の 4 つの項目を順番に入力します (インストール時にユーザー スコープを選択します)。
/プラグイン マーケットプレイスの更新 shin620265
/プラグインのアンインストール claude-pet@shin620265
/プラグインのインストール claude-pet@shin620265
/reload-プラグイン
このようにして、現在のセッションでは新しいバージョンが使用されます。
または、ターミナル タブを閉じて、claude --resume <セッション ID> で開きます (セッションは保持され、新しいプロセスは最新バージョンをロードする必要があります)。
デフォルトでは、カードをクリックするとウィンドウにジャンプします。 VS Code 内では、ターミナル ラベルを正確に指定できます。カード行をクリックすると、セッションのホスト ウィンドウ (Windows ターミナル / VS Code / PowerShell) が最前面に表示されます。最小化されている場合は、最初に復元されます。 VS Code ユーザーは、サポートする拡張機能 (vscode-ext/claude-pet-jump、オプション) をインストールし、カードをクリックしてウィンドウ内のセッションの端末ラベルに直接フォーカスできます。インストールされていない場合は、ウィンドウ レベルのままになります。 Windows ターミナルには複数のタブに対する信頼できる外部 API がないため、WT のウィンドウにジャンプした後、目的のタブをクリックする必要があります (これは仕様であり、バグではありません)。タブの精度が必要な場合は、VS Code (拡張機能をサポート) を使用してください。場所に到達できない場合 (セッションが閉じられている場合など)、カードは首を振り、ジャンプしません。
Windows のみ (powershell.exe + WinForms を使用)

。
ネイティブ GUI チャット パネルはスタックしていない可能性があります。ターミナルでクロード コードを実行すると、確実に利用できるようになります。拡張グラフィカル チャット パネルがフックをトリガーしない場合、または claude.exe を生成しない場合、セッションはスタックしない可能性があります (他のターミナル セッションには影響しません)。
カード名と終端タグ名は自動的には同期されません。カード タイトルはデフォルトで最初の文になり、終端タグ名とは独立しています。それらを同じにしたい場合は、✎ カード行の名前を手動で変更します (変更後はロックされ、自動的に上書きされません) をクリックします。
同じ画面上に最大 3 枚のカード: 残りのセッションはサイレントに追跡され、「要確認」または最近アクティブになるまで表示されません。これを超えると、3 番目のカードの右下に灰色の「+N」が表示され、バックグラウンドで N 個のセッションが追跡されていることを示します。
混合 DPI マルチスクリーン: インターフェースのスケーリングはメイン画面の DPI に従って計算されます。異なる拡大縮小率でペットを二次画面にドラッグすると、フォント サイズがわずかに大きく/小さくなる場合があります。
「非コマンド」操作を承認した後、カードは一時的に残る場合があります。コマンドを承認すると、コマンド自体の実行 (ビルド、テストなど) に時間がかかる場合でも、ペットは約 1 ～ 2 秒で実行が開始されたことを認識し、自動的に「思考」を再開します。ただし、承認された操作が他のタイプの操作 (MCP ツールへの長い呼び出しなど) である場合、クロード コードは「承認」と「実行完了」の間にイベントを生成せず、カードは処理を行わずに自動的に回復するまで操作の終了まで留まります。
ペットが見えませんか? /my-pet と入力して開きます。 Claude Code がまだ実行されていない場合は、実行していることを確認してください。
カード表示異常？ (個々のマシン/グラフィックス カード) ~/.claude/pet-data に空のファイル Card-region.flag を作成し (または環境変数 PET_CARD_REGION=1 を設定し)、ペットを再起動してシンプル カードにフォールバックします。
どの端末が使えるのでしょうか？はい: VS Code ターミナル、Windows ターミナル、および PowerShell ウィンドウはすべて同じです。
Windows デスクトップに常駐する小さなマスコット。実行中のクロード コードの会話を監視し、カードと柔らかいチャイムで、クロードが考えているか、あなたの選択を待っているか、または完了したかを通知します。そのため、端末を見つめ続ける必要はありません。クロードがあなたの許可を必要とするとき

オンにすると、約 1 秒以内にアラートが送信されます。Claude Code に組み込まれている 6 秒の通知遅延を待つ必要はありません。また、軽量です。インストールされるプラグインは約 200 KB (純粋な PowerShell スクリプトとサウンドとスプライト) です。Electron もバンドルされたランタイムもなく、Windows の組み込み UI スタックを駆動する PowerShell だけです。
必要なもの: Windows ・ クロード コード (最新の状態に保ちます。インスタント パーミッション アラートは新しいフック機能に依存します) ・ PowerShell 7 (pwsh コマンド。お持ちでない場合: winget install Microsoft.PowerShell 、または Microsoft Store から「PowerShell」をインストールします)
Claude Code で 2 つのコマンドを実行します。Claude は手動でダウンロードすることなく、GitHub から直接コードを取得します。
/プラグインマーケットプレイス追加 shin620265/claude-pet
/プラグインのインストール claude-pet@shin620265
次に、/reload-plugins を実行して、現在のセッションでアクティブ化します (最初のインストールでも必要です。再起動せず、ターミナルを保持します。または、新しいクロード セッションを開始するだけで、自動的にアクティブ化されます)。
完了 — /my-pet を使用してペットを開きます (記憶されており、新しいセッションに自動的に表示されます)。
アンインストール: /plugin uninstall claude-pet@shin620265 (完全なクリーンアップの場合は、~/.claude/pet-data フォルダーも削除します。このフォルダーには設定とカード メモリが保存されています)
すでにリポジトリをクローン/ダウンロードしましたか?最初のコマンドを /plugin Marketplace add <ローカルフォルダーへのパス> に置き換えます。他のコマンドはすべて同じです。
デスクトップ上の小さなキャラクター (クロード スパーク)。
会話の実行中、ステータス カードはその隣にスタックされます (最大 3 枚)。各カードには次のことが示されています。
タイトル = その会話で最初に入力した内容。
ステータス行: 思考中 / 入力を待機中 = 状態 + 最新のメッセージ。 Once Done = モデル · 完了 · 応答の最初の行 (バッジはその応答を生成したモデルの名前を示します。/model の後は次のモデルで更新されます)、色と i が付きます

短所:
🔵 思考中 (スピナー) ・ 🟡 入力が必要です ( ! 、自動的にトップに表示されます) ・ 🟢 完了 ( ✓ ) ・ ⚪ アイドル状態 ・ ⚫ 中断されました (Esc キーを押してターンを停止しました)
開く/閉じる : クロード コードに「/my-pet」と入力します (記憶されています。次にクロード コードを開くと再び表示され、すべて閉じると消えます)。
見てください。クロードが終了するか選択が必要になると、チャイムが鳴り、カードが一致する状態で点灯します。
Jump : カードをシングルクリック = そのセッションのウィンドウを最前面に表示します (ホバーするとサンゴのリングが表示されます。ウィンドウが見つからない場合、カードは首を振ります。間違ったウィンドウにジャンプすることはありません)。オプションの VS Code コンパニオン拡張機能を使用すると、同じクリックでそのセッションの正確なターミナル タブにも到達します。
アクション: ドラッグ = 移動 (記憶);ダブルクリック = カードを折りたたむ/展開します。行ごと ✎ = 名前変更、× = 却下 (新しいアクティビティで戻る)。右クリック = メニュー (ペットを閉じる / 位置をリセット / サウンド / プライバシー モード / 言語)。
言語: 右クリック → 「言語」 → 中文/英語/日本語/自動 (自動はシステムに従います);瞬時に切り替わります。
サウンド：右クリック→「サウンド」のオン/オフ。
プライバシーモード：右クリック→「プライバシーモード」のオン/オフ（デフォルトではオフ）。オンにすると、カードにはプロジェクト フォルダー名のみが表示され、プロンプト/応答テキストが非表示になります。画面共有、録画、またはストリーミングの前にオンに切り替えてください。すべてを非表示にするには、ダブルクリックして折りたたみます。
自動静音 : 全画面ゲーム、プレゼンテーション、または Windows フォーカス アシスト中は静音状態を保ちます (カードは引き続き更新されます)。
動きを減らす：設定→アクセシビリティ→視覚効果→「アニメーション効果」をオフにすると、ペットの揺れや瞬きが止まります。スピナーは静的な「…」になります。
クロード コードで、これらを順番に実行します ( install のユーザー スコープを選択します)。
/プラグイン マーケットプレイスの更新 shin620265
/プラグインのアンインストール claude-pet@shin620265
/

プラグインのインストール claude-pet@shin620265
/reload-プラグイン
これにより、現在のセッションが更新されます。
⚠️ 他の開いているセッションは古いバージョンのままです。更新するには、次のいずれかを実行します。
そのセッションでも /reload-plugins を実行します。
または、ターミナル タブを閉じて、claude --resume <session-id> で再度開きます (会話は維持されます。新しいプロセスは常に最新のものを読み込みます)。
新しいセッションは自動的に最新情報を取得します。特に何もする必要はありません。
カードをクリックするとウィンドウにジャンプし、VS Code 内では正確なターミナル タブにジャンプします。クリックすると、そのセッションのホスト ウィンドウ (Windows ターミナル / VS Code / PowerShell) が最前面に表示され、最小化されている場合は元に戻ります。 VS Code ユーザーはオプションでコンパニオン拡張機能 ( vscode-ext/claude-pet-jump ) をインストールできるため、同じクリックでそのセッションのターミナル タブにもフォーカスできます。これがないと、ジャンプはウィンドウレベルのままになります。特定の Windows ターミナル タブにフォーカスする場合、信頼できる外部 API がないため、Windows ターミナルでは設計上、ジャンプはウィンドウ レベルになります (タブを選択します)。タブの精度を高めるには、コンパニオン拡張機能を備えた VS Code を使用します。ウィンドウが見つからない場合 (セッションがなくなった場合など)、カードは推測する代わりに首を振ります。
Windows のみ (powershell.exe + WinForms を使用)。
ネイティブ GUI チャット パネルにはカードが表示されない場合があります。ターミナルでホストされるクロード コードは常に機能します。拡張機能の GUI パネルは機能しない可能性があります (フックを起動しない場合や claude.exe を生成しない場合)。他の端末セッションは影響を受けません。
カード名とターミナル タブ名は自動同期しません。カード タイトルはデフォルトで最初のプロンプトに設定され、タブ名とは独立しています。それらを一致させるには、行に ✎ が付いたカードの名前を変更します (新しい名前はロックされ、上書きされません)。
一度に表示されるカードは最大 3 枚です。他のセッションはサイレントに追跡され、必要なとき、または最新になったときに表示されます。さらにアクティブな場合は、小さな灰色の「+N」が表示されます。

3 番目のカードの右下には、バックグラウンドで追跡されている数が表示されます。
混合 DPI マルチモニター : UI スケールはプライマリ モニターの DPI から計算されるため、異なるスケール係数を持つモニターにペットをドラッグすると、テキストが若干大きすぎたり小さすぎたりする可能性があります。
コマンド以外のアクションを承認した後、カードが一時的に残る場合があります。コマンドを承認すると、コマンド自体が長時間実行される場合でも (ビルド、テスト スイート)、ペットはそのコマンドが実行され始めていることに気づき、1 ～ 2 秒以内に「思考」に戻ります。他の承認タイプ (長い MCP ツール呼び出しなど) の場合、クロード コードは「承認」と「完了」の間にイベントを発行しないため、その呼び出しが完了するまでカードはその状態を保持します。自然に回復します。
ペットが見えませんか？ /my-pet と入力します。それでも何も起こらない場合は、Claude Code が実行されていることを確認してください。
音が出ませんか？右クリック→「サウンド」がオンになっており、全画面表示/おやすみモードになっていないことを確認してください。
カードが見えなくなっていませんか？ (まれに、一部では

[切り捨てられた]

## Original Extract

Desktop companion pet for Claude Code (Windows). Contribute to SHIN620265/claude-pet development by creating an account on GitHub.

I got the idea from Codex. It has a little pet that reminds you when it finishes. I liked it. Claude Code on Windows did not have one, so I made the same kind of desktop pet for Claude Code, with one rule I set myself: no Electron, and no bundled runtime. I use Claude Code in the terminal every day. I often miss the moment it finishes, or the moment it stops and waits for me to approve a command. The terminal does not remind you. The pet watches my Claude Code sessions through the hooks, and shows one card for each session: thinking, waiting for you, or done. When a session needs you, it plays a soft sound. When Claude Code asks for permission, the card shows it in about one second for me, because the pet listens to the PermissionRequest hook instead of the built-in notification. It is about 200KB installed, and it does not make any network request. It is Windows only for now. I put the backstory, and the one problem I could not solve, in a comment below.

GitHub - SHIN620265/claude-pet: Desktop companion pet for Claude Code (Windows) · GitHub
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
SHIN620265
/
claude-pet
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
42 Commits 42 Commits .claude-plugin .claude-plugin docs docs pet pet scripts scripts tools tools vscode-ext/ claude-pet-jump vscode-ext/ claude-pet-jump .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md LICENSE LICENSE README.md README.md REVIEW-PROMPT.md REVIEW-PROMPT.md View all files Repository files navigation
Claude 桌面宠物 · Claude Desktop Pet · Claude デスクトップペット
🌸 一只住在桌面、帮你盯 Claude Code 进度的小宠物。
A little desktop mascot that watches your Claude Code progress for you.
Claude Code の進捗を見守る、デスクトップの小さなマスコット。
↑ real screen recording — not a mockup · 中文・日本語界面可右键即时切换
语言 / Language / 言語: 中文 ・ English ・ 日本語
⚠️ 非官方社区项目,与 Anthropic 无关 / Unofficial community project, not affiliated with Anthropic / 非公式のコミュニティ製。Anthropic とは無関係です
给改代码的人 / Developers / 開発者 → ARCHITECTURE.md
一只待在你 Windows 桌面上的小宠物。它盯着你正在跑的 Claude Code 对话,用 卡片 和 轻提示音 告诉你:Claude 是在 思考 、在 等你做选择 、还是 已经做完了 ——这样你不用一直盯着命令行窗口。需要你确认权限时,约 1 秒 内就会提醒——不必等 Claude Code 自带通知的 6 秒延迟。它也很轻:装到本地的插件约 200KB (纯 PowerShell 脚本 + 音效立绘),没有 Electron、没有捆绑运行时,只是 PowerShell 驱动着 Windows 自带的界面组件。
需要: Windows ・ Claude Code (请保持最新版——即时权限提醒依赖较新的钩子能力)・ PowerShell 7 (命令 pwsh ;没有就 winget install Microsoft.PowerShell ,或到 Microsoft Store 搜 "PowerShell" 安装)
在 Claude Code 里输入两条命令——Claude 会 直接从 GitHub 拉取 ,不用手动下载:
/plugin marketplace add shin620265/claude-pet
/plugin install claude-pet@shin620265
然后运行 /reload-plugins 让它在当前会话生效( 首次安装也要这步 ;无需重启、不丢终端标签。或者开一个新的 Claude 会话,它也会自动生效)。
完成!用 /my-pet 打开宠物(之后会记住,新会话自动出现)。
卸载: /plugin uninstall claude-pet@shin620265 (想彻底清理,再删除 ~/.claude/pet-data 文件夹——存的是设置与卡片记忆)
已经把仓库 clone / 下载到本地了?把第一条换成 /plugin marketplace add <本地文件夹路径> 即可,其余不变。
有对话在跑时,旁边叠出 状态卡 (最多 3 张),每张显示:
标题 = 你在那个对话里说的第一句话。
状态行 :思考/需确认时 = 当前状态 + 你最近一句输入;完成后 = 模型名 · 已完成 · 回复首句摘要 (模型名指产出这条回复的模型, /model 切换后随下一条回复自动更新),带颜色和图标:
🔵 正在思考(转圈)・ 🟡 需要你确认/选择( ! ,自动置顶)・ 🟢 已完成( ✓ )・ ⚪ 空闲・ ⚫ 已中断(Esc 打断了当前回合)
开/关 :在 Claude Code 里输入 /my-pet (会记住;下次开 Claude 自动出现,全部关掉后消失)。
看提示 :平时不用管它;Claude 做完或需要你选择时会响一声 + 卡片亮起对应状态。
跳过去 : 单击卡片 = 把那个会话所在的窗口拉到前台 (悬停出珊瑚色光环;定位不到时卡片会左右摇头,绝不跳错窗口)。装上可选的 VS Code 伴生扩展 后,同一次点击还会 直达该会话的终端标签 。
操作 :拖动=移动(记忆位置);双击=收起/展开卡片;每行 ✎ =改标题、 × =关掉这张卡(有新动静会自己回来); 右键 =菜单(关闭宠物/回默认位置/提示音/隐私模式/语言)。
语言 :右键 →「语言」→ 中文/English/日本語/自动(自动=跟随系统),点一下即时切换。
隐私模式 :右键 →「隐私模式」开/关(默认关)。开启后卡片 只显项目文件夹名、藏起你的输入和回复正文 ——共享屏幕/录屏/直播前一键打开。想彻底不露 → 双击折叠藏全部。
自动安静 :全屏游戏、演示、或开了"专注助手/勿扰"时不出声(卡片照常更新)。
减少动画 :在 设置 → 辅助功能 → 视觉效果 →「动画效果」关掉后,宠物不再浮动/眨眼,转圈变静止「…」。
在 Claude Code 里 依次 输入这四条( install 时选 user scope):
/plugin marketplace update shin620265
/plugin uninstall claude-pet@shin620265
/plugin install claude-pet@shin620265
/reload-plugins
这样 当前会话 就用上新版了。
或 关掉它的终端标签,再 claude --resume <会话id> 打开 (保留对话,新进程一定加载最新版);
点卡片默认跳到窗口;VS Code 内可精确到终端标签 :单击卡片行会把该会话的宿主窗口(Windows Terminal / VS Code / PowerShell)拉到前台,最小化了会先还原。VS Code 用户可另装配套扩展( vscode-ext/claude-pet-jump ,可选),点卡后在窗口内直接聚焦该会话的终端标签;不装则保持窗口级。Windows Terminal 的多个 tab 没有可靠的外部 API,所以在 WT 里跳到窗口后、具体哪个 tab 需你自己点(设计上如此、不是 bug);想要 tab 精度请用 VS Code(配套扩展)。定位不到时(如会话已关闭)卡片会摇头示意,不会瞎跳。
仅 Windows (用 powershell.exe + WinForms)。
原生 GUI 聊天面板可能不出卡 :在终端里跑 Claude Code 一定可用;扩展的图形聊天面板若不触发钩子/不产生 claude.exe ,该会话可能不出卡(不影响其它终端会话)。
卡片名与终端标签名不自动同步 :卡片标题默认取你的第一句话,与终端标签名各自独立。想让它俩一致, 点卡片行上的 ✎ 手动改名 即可(改完会锁定,不被自动覆盖)。
同屏最多 3 张卡 :其余会话静默跟踪,等"需确认"或最近活跃时才浮现;超出时第 3 张卡右下会显示灰色 "+N",表示还有 N 个会话在后台跟踪。
混合 DPI 多屏 :界面缩放按主屏 DPI 计算;把宠物拖到缩放比例不同的副屏时,字号可能略偏大/偏小。
批准"非命令类"操作后卡片可能短暂滞留 :批准一条命令后,宠物会在约 1-2 秒内认出它已经开跑,自动恢复"正在思考"——哪怕命令本身要跑很久(构建、测试等)。但批准的若是其它类型的操作(如 MCP 工具的长调用),Claude Code 在"已批准"和"跑完"之间不产生任何事件,卡片会停留到该操作结束才自动恢复,无需处理。
没看到宠物? 输入 /my-pet 打开;还没有就确认 Claude Code 在运行。
卡片显示异常? (个别机器/显卡)在 ~/.claude/pet-data 建个空文件 card-region.flag (或设环境变量 PET_CARD_REGION=1 ),重启宠物即回退到简约卡片。
哪个终端都能用? 能:VS Code 终端、Windows Terminal、PowerShell 窗口都一样。
A little mascot that lives on your Windows desktop. It watches your running Claude Code conversations and tells you — with cards and a soft chime — whether Claude is thinking , waiting for your choice , or done . So you don't have to keep staring at the terminal. When Claude needs your permission, you're alerted within about 1 second — no waiting for Claude Code's built-in 6-second notification delay. It's light, too: the installed plugin is about 200KB (pure PowerShell scripts plus sounds and sprites) — no Electron, no bundled runtime, just PowerShell driving Windows' built-in UI stack.
Requires: Windows ・ Claude Code (keep it up to date — the instant permission alerts rely on newer hook capabilities) ・ PowerShell 7 (the pwsh command; if you don't have it: winget install Microsoft.PowerShell , or install "PowerShell" from the Microsoft Store)
In Claude Code, run two commands — Claude fetches it straight from GitHub , no manual download:
/plugin marketplace add shin620265/claude-pet
/plugin install claude-pet@shin620265
Then run /reload-plugins to activate it in your current session ( needed on first install too ; no restart, keeps your terminals — or just start a new Claude session and it activates automatically).
Done — open the pet with /my-pet (it remembers, and appears automatically in new sessions).
Uninstall: /plugin uninstall claude-pet@shin620265 (for a full cleanup, also delete the ~/.claude/pet-data folder — it holds settings and card memory)
Already cloned/downloaded the repo? Replace the first command with /plugin marketplace add <path-to-the-local-folder> — everything else is the same.
A small character on your desktop (the Claude spark).
When conversations are running, status cards stack next to it (up to 3). Each card shows:
Title = the first thing you typed in that conversation.
Status line : while thinking / awaiting your input = state + your latest message; once done = model · Done · first line of the reply (the badge names the model that produced that reply; after /model it updates with the next one), with color and icon:
🔵 Thinking (spinner) ・ 🟡 Needs your input ( ! , auto-floats to top) ・ 🟢 Done ( ✓ ) ・ ⚪ Idle ・ ⚫ Interrupted (you pressed Esc to stop the turn)
Open/close : type /my-pet in Claude Code (it remembers; reappears next time you open Claude Code, disappears once all are closed).
Just watch : when Claude finishes or needs a choice, it chimes and the card lights up with the matching state.
Jump : single-click a card = bring that session's window to the foreground (a coral ring appears on hover; the card shakes its head when the window can't be located — it never jumps to the wrong one). With the optional VS Code companion extension , the same click also lands on that session's exact terminal tab .
Actions : drag = move (remembered); double-click = collapse/expand cards; per row ✎ = rename, × = dismiss (comes back on new activity); right-click = menu (close pet / reset position / sound / privacy mode / language).
Language : right-click → "Language" → 中文/English/日本語/Auto (Auto follows your system); switches instantly.
Sound : right-click → "Sound" on/off.
Privacy mode : right-click → "Privacy mode" on/off (off by default). When on, cards show only the project folder name and hide your prompt/reply text — flip it on before screen-sharing, recording, or streaming. To hide everything, double-click to collapse.
Auto-quiet : stays silent during fullscreen games, presentations, or Windows Focus Assist (cards still update).
Reduced motion : turn off Settings → Accessibility → Visual effects → "Animation effects" and the pet stops bobbing/blinking; the spinner becomes a static "…".
In Claude Code, run these in order (choose user scope for install ):
/plugin marketplace update shin620265
/plugin uninstall claude-pet@shin620265
/plugin install claude-pet@shin620265
/reload-plugins
That updates the current session .
⚠️ Other open sessions stay on the old version. To update them, either:
run /reload-plugins in that session too;
or close its terminal tab and reopen with claude --resume <session-id> (keeps your conversation; a fresh process always loads the latest);
new sessions get the latest automatically — nothing to do.
Clicking a card jumps to the window — and, inside VS Code, to the exact terminal tab : a click brings that session's host window (Windows Terminal / VS Code / PowerShell) to the foreground, restoring it if minimized. VS Code users can optionally install the companion extension ( vscode-ext/claude-pet-jump ) so the same click also focuses that session's terminal tab; without it the jump stays window-level. Focusing a specific Windows Terminal tab has no reliable external API, so in Windows Terminal the jump is window-level by design (you pick the tab); for tab precision use VS Code with the companion extension. When the window can't be located (e.g. the session is gone), the card shakes its head instead of guessing.
Windows only (uses powershell.exe + WinForms).
The native GUI chat panel may not show a card : terminal-hosted Claude Code always works; the extension's GUI panel may not (if it doesn't fire hooks / spawn claude.exe ). Other terminal sessions are unaffected.
Card name and terminal tab name don't auto-sync : the card title defaults to your first prompt and is independent of the tab name. To make them match, rename the card with the ✎ on its row (the new name is locked and won't be overwritten).
At most 3 cards show at once : other sessions are tracked silently and surface when they need you or become most-recent; when more are active, a small gray "+N" at the bottom-right of the third card shows how many are tracked in the background.
Mixed-DPI multi-monitor : UI scale is computed from the primary monitor's DPI, so text may look slightly too large/small if you drag the pet to a monitor with a different scale factor.
The card may briefly linger after approving a non-command action : when you approve a command, the pet spots it starting to run and flips back to "Thinking" within ~1-2 seconds — even if the command itself runs long (builds, test suites). For other approval types (e.g. long MCP tool calls), Claude Code emits no event between "approved" and "finished", so the card keeps its state until that call completes. It recovers on its own.
Don't see the pet? Type /my-pet . If still nothing, make sure Claude Code is running.
No sound? Check right-click → "Sound" is on, and you're not in fullscreen / Do-Not-Disturb.
Cards look off? (rare, on some

[truncated]
