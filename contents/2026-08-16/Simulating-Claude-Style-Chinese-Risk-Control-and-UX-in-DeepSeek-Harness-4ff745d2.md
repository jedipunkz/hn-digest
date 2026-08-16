---
source: "https://github.com/eri64/dsh-claude-ux"
hn_url: "https://news.ycombinator.com/item?id=49322653"
title: "Simulating Claude-Style Chinese Risk Control and UX in DeepSeek Harness"
article_title: "GitHub - eri64/dsh-claude-ux: DSH plugin: Claude-style Chinese risk control & conversation autonomy for DeepSeek Harness web · GitHub"
author: "Sha1rholder"
captured_at: "2026-08-16T19:14:01Z"
capture_tool: "hn-digest"
hn_id: 49322653
score: 1
comments: 0
posted_at: "2026-08-16T18:59:05Z"
tags:
  - hacker-news
  - translated
---

# Simulating Claude-Style Chinese Risk Control and UX in DeepSeek Harness

- HN: [49322653](https://news.ycombinator.com/item?id=49322653)
- Source: [github.com](https://github.com/eri64/dsh-claude-ux)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T18:59:05Z

## Translation

タイトル: DeepSeek ハーネスにおけるクロード流の中国のリスク管理と UX のシミュレーション
記事タイトル: GitHub - eri64/dsh-claude-ux: DSH プラグイン: DeepSeek Harness Web 用のクロード式中国語リスク制御と会話自律 · GitHub
説明: DSH プラグイン: DeepSeek Harness Web 用のクロード スタイルの中国語リスク コントロールと会話の自律性 - eri64/dsh-claude-ux

記事本文:
GitHub - eri64/dsh-claude-ux: DSH プラグイン: DeepSeek Harness Web 用のクロード スタイルの中国語リスク コントロールと会話の自律性 · GitHub
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
えり64
/
dsh-クロード-ux
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット docs docs 例 例 lib lib test test .gitignore .gitignore ライセンス ライセンス README.md README.mdcordis.patch.ymlcordis.patch.yml package-lock.json

package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード スタイルの「地域リスク管理 + 自律的な会話終了」プラグイン - DeepSeek Harness の Web プロファイルに適しています。
Anthropic/Claude の 2 種類の動作 (デフォルトでオフになっている 2 つのオプションの外部呼び出しを除くすべてローカル決定) を複製します (詳細については docs/PRIVACY.md を参照)。
地域リスク管理 (可逆的): ターゲット ユーザー (タイム ゾーン、システム/ブラウザ言語、中国語フォント、プロキシ、プロキシ/トランジット ドメイン名のブラックリスト、パブリック ネットワーク IP 帰属、WebRTC IP 一貫性) を検出します。 [regionTarget] では、cn = 中国人ユーザーのリスク管理 (クロードの本来の動作) を選択し、non-cn = 逆リスク管理 (中国人でない場合のリスク管理) を選択します。ヒット後、ペナルティはラダーに従って処理されます: コピーを拒否 (試行回数付き) → 拒否回数に達した後セッションを終了EndsAfter 回数 (チャット終了パネル + サーバーは拒否し続け、再起動後も有効) → システム プロンプト ワードがモデル レベルの地域命令に挿入されます。
自律性: ユーザーが悪用を続けたり、深刻な有害なコンテンツを繰り返し要求したりした場合、ユーザーはまず警告し、その後積極的に会話を終了します。自傷行為や他害行為のリスクメッセージが終わりを引き起こすことはありません（クロードの公的制限に沿って）。虐待的なエンディングと深刻な有害なエンディングでは、独立したコピーライティングが使用されています (どちらも設定ページでカスタマイズでき、空の値は組み込みのデフォルトとして灰色のテキストで表示されます)。侮辱判定は、単語リストの二次判定 + LLM 文脈分析を使用します。強い単語 (バカ/クソなど) は直接判定されます。弱い単語 (ジャンク/黙れなど) と見逃したメッセージは、コンテキスト判断のために個別に LLM に要求されます (目的タグ。会話ログやモデル コンテキストには入りません。メッセージは非同期の事前分類のためにキューに入れられ、追加の遅延はほぼゼロです)。分類モデルは、DSH 構成モデル カタログのドロップダウンから直接選択することも、空白のままにしてセッションのメイン モデルに従うこともできます。
ステガノグラフィック チャネル: システム プロンプト ワード日付形式 (2026-06-30 ↔ 2026/06/30) エンコード領域の決定、Unicode アポストロフィ バリアント (U+2019 ↔ U+02BC) エンコード ブラックリスト ヒット。
「クロードリスクコントロール」タブページ：右

メイン スイッチ、検出ステータス カード (リスク コントロール ターゲット/ホスト判定/ヒット/信号スコア/パブリック ネットワーク IP)、地域リスク コントロール パネル (リバース ターゲット ドロップダウン、ポリシー、ステガノグラフィック タグ、WebRTC)、分類 LLM パネル (判定モード、モデル ドロップダウン、メッセージ分類セルフテスト)、自律パネル (警告/終了しきい値)、およびコピーライティング パネル (組み込みのデフォルトを表示するには空白のままにします)。
ターゲットがヒットすると、各メッセージには拒否のコピーが返信され、「N/M 回目の試行 - 続行するとこの会話は終了します」が添付されます。 RefusalEndsAfter 回数に達すると、セッションはリージョン エンド コピーで終了され、サーバーは後続のメッセージを拒否し続けます。
侮辱を続ける場合は、まず警告が表示され、その後、独立したエンディング コピーで積極的に会話を終了します (「前のメッセージ」を引用せずに、しきい値を 1 に設定できます)。終了後、同じ会話内のすべてのメッセージはサーバーによって拒否されます。
未成年の性的コンテンツ、テロリズム、大規模な暴力などの重大な有害なコンテンツの場合は、警告なしに独立したテキストで会話を直接終了します。自傷行為やその他の危害のリスクに関するメッセージが終わりを引き起こすことはありません。
npx -y @deepseek-ai/dsh プラグイン --profile web add github:eri64/dsh-claude-ux
パッケージには登録エントリ (dsh.bundle) が付属しています。 dsh プラグインはインストールされると自動的に登録されるため、構成ファイルを手動で変更する必要はありません。
アップデート（最新バージョンへのアップグレード）も同じコマンドです。
Web プロファイル (dsh Web) を再起動し、ブラウザをハードリフレッシュします。
「Claude Risk Control」の独立タブが、設定ページの左側のタブ バーに表示されます (ビジュアル ツール プラグインと同じ)。
マスタースイッチ、リスク管理対象（中国人/非中国人ユーザー向けリバーススイッチ）、地域戦略、リアルタイム検知状況
(/_dsh/claude/status)、しきい値とコピーライティング。
デフォルトではオフになっています (ロックアウトを避けるために有効: false)。メインスイッチを入れて「保存して適用」するとすぐに有効になります。
インストールが成功したかどうかを確認します。ブラウザはローカル dsh Web の /_dsh/claude/status にアクセスします。
(デフォルト http://127.0.0.1:3080/_dsh/claude/status)、ホスト プラグインがロードされたことを意味する {"ok":true,...} を返します。
組み込みのデフォルト

値はそのまま使用できます。パッチレベルのオプション(ブラックリスト/cnTimezones/語彙など)をオーバーライドする必要がある場合
パッケージを変更するには、examples/cordis.patch.yml を参照してください。
node_modules/dsh-claude-ux/cordis.patch.yml を削除して再インストールしてください (設定ページで変更できる項目は、設定ページで優先されます)。
キー
デフォルト
説明
有効
偽
メインスイッチ（設定ページで変更可能）
地域.ターゲット
CN
誰がリスクをコントロールするか: cn = 中国人ユーザー | non-cn = 非中国人ユーザー (逆)
地域.ポリシー
ブロック
ブロック = 返信を拒否 |観察 = ログのみ + ステガノグラフィック マーク |オフ = 閉じる
地域.minSignals
2
信号閾値（強=2 / 中=1 / 弱=0.5）
地域.拒否終了後
3
N回拒否した後にセッションを終了する(「アカウント禁止」ラダーエリア)
地域.showAttempts
本当の
コピーに「N/M 回目の試行」を追加することを拒否する
地域.プロンプト施行
本当の
ターゲットのヒット時にモデルレベルの領域コマンドを挿入します
地域.ipCheck
偽
パブリック ネットワークの IP 所有権のクエリ (ipinfo.io / ip-api.com に問い合わせ、デフォルトでは閉じられています)
地域.webRtcCheck
偽
WebRTC IP 検出 (Google STUN にお問い合わせください、デフォルトではオフになっています)
地域.blockOnUnknown
偽
信号がまったくない場合 (フェールクローズ)、ターゲット ヒットとみなされるかどうか
地域.cnTimezones / 地域.ブラックリスト
内蔵
中国タイムゾーンテーブル/プロキシトランジットドメイン名ブラックリスト(丸ごと置き換え可能)
虐待。*
内蔵
暴言/重大な有害/自傷行為の言葉のリスト (通常、全体として置き換え可能)。 StrongPatterns /weakPatterns の分類
虐待.llm.モード
すべて
侮辱判定モード: all = すべてのメッセージ LLM コンテキスト判定 (最も正確) |ファジー = 弱いトーンのみ LLM |オフ = 純粋な単語リスト
虐待.llm.有効
本当の
LLM分類メインスイッチ
Abuse.llm.プロバイダー/モデル
会話に従ってください
リクエストを分類するためのモデル (デフォルトの再利用セッションのメイン モデル/エージェント デフォルト モデル)
ステガノグラフィー
本当の
ステガノグラフィー
設定ページは変更できます:enabled/regionPolicy/regionTarget/abuseEnabled/warnThreshold/endThreshold/refusalEndsAfter/warnEveryOffense/severeEndsImmediately/steganography/webRtcCheck/

llmMode / llmProvider / llmModel / llmTimeoutMs / 4 つのコピー (保存後すぐに有効になり、再起動する必要はありません)。その中で、分類モデルのドロップダウン リストには、DSH によって構成されたモデルが直接リストされます (設定のプロバイダーの名前空間をスキャンします)。空白のままにします = セッションのメイン モデルに従います。 「カスタマイズ...」では、ディレクトリの外にあるモデルを手動で入力できます。
blacklist/cnTimezones/vocabulary/minSignals/showAttempts/promptEnforcement/ipCheck/blockOnUnknown/abuse.llm.enabled にパッチを適用してから再起動する必要があります (プロバイダー/モデル/タイムアウトをページ変更に設定できます)。
デフォルト設定 (ipCheck: false、webRtcCheck: false) では、プラグインは外部サービスと通信しません。すべての検出 (タイム ゾーン、言語、プロキシ、ブラックリスト、ブラウザ言語/フォント) はローカルで完了し、テレメトリは報告されません。完全なデータ フローと露出ポイントの分析については、docs/PRIVACY.md を参照してください。
npm test # dsh インスタンスを使用せずにホスト (115 項目) およびクライアント (17 項目) の単体テストを実行します
テスト範囲：分類子（語彙分類+LLM文脈判断/曖昧さ回避/失敗ダウングレード/非同期事前分類）、領域判定（逆ターゲット含む）、
ステート マシン、llm/stream 合成ストリーム置換、終了リジェクト、日付ステガノグラフィー書き換え、プロジェクション フォールディング、ステート ルーティング
(GET/POST/競合/オリジナル検証/モデル カタログ/分類セルフテスト)、クライアントが契約を選択します。
DSH プラグイン: DeepSeek Harness Web 向けのクロード スタイルの中国語リスク コントロールと会話の自律性
github.com/eri64/dsh-claude-web トピック
Readme MIT ライセンス アクティビティ スター
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

DSH plugin: Claude-style Chinese risk control & conversation autonomy for DeepSeek Harness web - eri64/dsh-claude-ux

GitHub - eri64/dsh-claude-ux: DSH plugin: Claude-style Chinese risk control & conversation autonomy for DeepSeek Harness web · GitHub
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
eri64
/
dsh-claude-ux
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits docs docs examples examples lib lib test test .gitignore .gitignore LICENSE LICENSE README.md README.md cordis.patch.yml cordis.patch.yml package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Claude 式「区域风控 + 自主结束对话」插件 —— 适用于 DeepSeek Harness 的 web profile。
复刻 Anthropic/Claude 的两类行为， 除两个默认关闭的可选外部调用外全部本地判定 （详见 docs/PRIVACY.md ）：
区域风控（可反向） ：检测目标用户（时区、系统/浏览器语言、中文字体、代理、代理/中转域名黑名单、公网 IP 归属、WebRTC IP 一致性）。 regionTarget 选 cn = 风控中国用户（Claude 原版行为），选 non-cn = 反向风控 （检测到不是中国人就风控）。命中后按惩罚阶梯处置：拒绝文案（带尝试计数）→ 达到 refusalEndsAfter 次数后结束会话（Chat ended 面板 + 服务端持续拒绝，重启后依然生效）→ 系统提示词注入模型级区域指令。
自主性 ：用户持续辱骂或反复要求严重有害内容时，先警告、再主动结束对话；自伤/他伤风险消息永不触发结束（对齐 Claude 的公开限制）。辱骂结束与严重有害结束使用 独立文案 （均可在设置页自定义，空值灰字显示内置默认）。辱骂判定采用 词表秒判 + LLM 语境兜底 ：强词（傻逼/fuck 等）直接判；弱词（垃圾/闭嘴等）与未命中消息由独立 LLM 请求做语境裁决（ purpose 标记， 不进会话日志与模型上下文 ；消息入队即异步预分类，近零额外延迟）。分类模型可直接从 DSH 已配置的模型目录 下拉选择，或留空跟随会话主模型。
隐写通道 ：系统提示词日期格式（ 2026-06-30 ↔ 2026/06/30 ）编码区域判定，Unicode 撇号变体（U+2019 ↔ U+02BC）编码黑名单命中。
「Claude 风控」标签页：右上总开关、检测状态卡（风控目标 / 主机判定 / 命中 / 信号分 / 公网 IP）、区域风控面板（反向目标下拉、策略、隐写标记、WebRTC）、分类 LLM 面板（判定模式、模型下拉、消息分类自测）、自主性面板（警告 / 结束阈值）与文案面板（留空灰字显示内置默认）。
目标命中时每条消息回复拒绝文案并附「第 N/M 次尝试——继续将结束本次对话」；达到 refusalEndsAfter 次数后以区域结束文案终止会话，服务端对后续消息持续拒绝。
持续辱骂先警告、再以独立结束文案主动结束对话（不引用"上一条消息"，阈值可设为 1）；结束后同会话消息全部被服务端拒绝。
未成年性内容 / 恐怖主义 / 大规模暴力等严重有害内容不警告、直接以独立文案结束对话；自伤 / 他伤风险消息则永不触发结束。
npx -y @deepseek-ai/dsh plugin --profile web add github:eri64/dsh-claude-ux
包内自带注册条目（ dsh.bundle ）， dsh plugin 安装后 自动注册 ，无需手动改任何配置文件。
更新（升级到最新版）也是同一条命令。
重启 web profile（ dsh web ），浏览器硬刷新。
设置页左侧标签栏出现 「Claude 风控」 独立标签（与视觉工具插件同款）：
总开关、风控目标（中国/非中国用户反向开关）、区域策略、实时检测状态
（ /_dsh/claude/status ）、阈值与文案。
默认关闭 （ enabled: false ，避免把自己锁在外面）；打开总开关并「保存并应用」后即时生效。
校验安装是否成功：浏览器访问本机 dsh web 的 /_dsh/claude/status
（默认 http://127.0.0.1:3080/_dsh/claude/status ），返回 {"ok":true,...} 即主机插件已加载。
内置默认值开箱即用；需要覆盖 patch 级选项（ blacklist / cnTimezones / 词表等）时，
参考 examples/cordis.patch.yml 修改包内
node_modules/dsh-claude-ux/cordis.patch.yml 后重新安装即可（设置页可改的项优先用设置页）。
键
默认
说明
enabled
false
总开关（设置页可改）
region.target
cn
风控谁： cn =中国用户 | non-cn =非中国用户（反向）
region.policy
block
block =拒绝回复 | observe =只记录+隐写标记 | off =关闭
region.minSignals
2
信号分阈值（strong=2 / medium=1 / weak=0.5）
region.refusalEndsAfter
3
拒绝 N 次后结束会话（区域「封号」阶梯）
region.showAttempts
true
拒绝文案追加「第 N/M 次尝试」
region.promptEnforcement
true
目标命中时注入模型级区域指令
region.ipCheck
false
公网 IP 归属查询（联系 ipinfo.io / ip-api.com， 默认关闭 ）
region.webRtcCheck
false
WebRTC IP 探测（联系 Google STUN， 默认关闭 ）
region.blockOnUnknown
false
完全无信号时是否视为目标命中（fail-closed）
region.cnTimezones / region.blacklist
内置
中国时区表 / 代理中转域名黑名单（可整体替换）
abuse.*
内置
辱骂/严重有害/自伤词表（正则，可整体替换）； strongPatterns / weakPatterns 分级
abuse.llm.mode
all
辱骂判定模式： all =全部消息 LLM 语境判定（最准）| fuzzy =仅弱词调 LLM | off =纯词表
abuse.llm.enabled
true
LLM 分类总开关
abuse.llm.provider/model
跟随会话
分类请求的模型（缺省复用会话主模型 / agent-default-model）
steganography
true
隐写标记
设置页可改： enabled / regionPolicy / regionTarget / abuseEnabled / warnThreshold / endThreshold / refusalEndsAfter / warnEveryOffense / severeEndsImmediately / steganography / webRtcCheck / llmMode / llmProvider / llmModel / llmTimeoutMs / 四条文案 （保存即生效，无需重启）。其中分类模型 下拉直接列出 DSH 已配置的模型 （扫描 settings 的 providers 命名空间），留空=跟随会话主模型，「自定义…」可手填目录外的模型。
blacklist / cnTimezones / 词表 / minSignals / showAttempts / promptEnforcement / ipCheck / blockOnUnknown / abuse.llm.enabled 需改 patch 后重启（provider/model/timeout 已可设置页改）。
默认配置下（ ipCheck: false 、 webRtcCheck: false ） 插件不与任何外部服务通信 ：所有检测（时区、语言、代理、黑名单、浏览器语言/字体）都在本机完成，不上报任何遥测。完整的数据流向与暴露点分析见 docs/PRIVACY.md 。
npm test # 运行主机（115 项）与客户端（17 项）单元测试，无需 dsh 实例
测试覆盖：分类器（词表分级 + LLM 语境裁决/消歧/失败降级/异步预分类）、区域判定（含反向目标）、
状态机、llm/stream 合成流替换、ended reject、日期隐写改写、投影折叠、状态路由
（GET/POST/冲突/同源校验/模型目录/分类自测）、客户端 select 契约。
DSH plugin: Claude-style Chinese risk control & conversation autonomy for DeepSeek Harness web
github.com/eri64/dsh-claude-web Topics
Readme MIT license Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
