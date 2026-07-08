---
source: "https://github.com/Clear-Sights/Makoto"
hn_url: "https://news.ycombinator.com/item?id=48835278"
title: "Show HN: Makoto – blocks AI agent claims that contradict its logged actions"
article_title: "GitHub - Clear-Sights/Makoto: When Claude says “tests pass”, makoto checks the record — every claim held against the agent's own logged deeds; fakes blocked, not warned. Each check ships at measured zero false positives. 誠 · GitHub"
author: "Present_Flow"
captured_at: "2026-07-08T19:08:19Z"
capture_tool: "hn-digest"
hn_id: 48835278
score: 1
comments: 0
posted_at: "2026-07-08T18:11:17Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Makoto – blocks AI agent claims that contradict its logged actions

- HN: [48835278](https://news.ycombinator.com/item?id=48835278)
- Source: [github.com](https://github.com/Clear-Sights/Makoto)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T18:11:17Z

## Translation

タイトル: HN を表示: 誠 – ログに記録されたアクションと矛盾する AI エージェントの主張をブロックします
記事のタイトル: GitHub - Clear-Sights/Makoto: Claude が「テストに合格」と言ったとき、Makoto は記録をチェックします。エージェント自身の記録された行為に対して保持されているすべての請求です。偽物はブロックされますが、警告はされません。各小切手は誤検知ゼロの測定値で出荷されます。 誠 · GitHub
説明: クロードが「テストは合格」と言ったとき、マコトは記録をチェックします。エージェント自身の記録された行為に対して保持されているすべての請求です。偽物はブロックされますが、警告はされません。各小切手は誤検知ゼロの測定値で出荷されます。 誠 - Clear-Sights/Makoto

記事本文:
GitHub - Clear-Sights/Makoto: クロードが「テストに合格した」と言うと、マコトは記録をチェックします。エージェント自身の記録された行為に対して保持されているすべての請求です。偽物はブロックされますが、警告はされません。各小切手は誤検知ゼロの測定値で出荷されます。 誠 · GitHub
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

クリアサイト
/
マコト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
21 コミット 21 コミット .claude-plugin .claude-plugin .github .github canon canon checks checks コマンド コマンド データ データ ドキュメント docs フック フック lib lib prechecks prechecks stopchecks stopchecks テスト テスト ツール ツール .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md ライセンス ライセンス README.md README.md __init__.py __init__.py __main__.py __main__.py _dispatch.py _dispatch.py _dispatch_configchange.py _dispatch_configchange.py _dispatch_shim.sh _dispatch_shim.sh ackblock.py ackblock.py、audit.py、audit.py、引用、引用、py、コミットメント.py、コミットメント.py、configchange_verdict.py、configchange_verdict.py、db.py、db.py、install.py、install.py、ledger.py、ledger.py、lexicons.py、lexicons.py、plan.py、plan.py、posture.py、posture.py、pyproject.toml、pyproject.toml、receipt.py、retraction.py retraction.py schema.py schema.py state.py state.py Wire.py Wire.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェント自身のツール呼び出しを監視し、その呼び出しをブロックするクロード コードの整合性フック
それは小切手を偽造します。クロードが何かをした（テストを実行した、論文を引用した、コミットした）と言うとき、
修正し、証明書を確認しました）、マコトはその言葉を記録に保管します。権利書が無い場合は、
または検証が静かに無効になった場合、マコトはツール呼び出し (またはターン終了) をブロックし、手を差し出します。
エージェントに対して再試行する 1 行の修正。
それは決して世界の真実ではなく、エージェント自身の発言と記録に基づいてエージェントを判断します。いいえ
事実（「フランスには存在しない」）。主張された言葉が保持され、完全で、尊重されていることを確認するだけです
行為において。あ

通過する言葉は使いやすくなります。レビュー担当者や他のエージェントが信頼できる入札を行うことができます。
再導出せずに受け入れます。
makoto は、機械的フック イベント (すべての PreToolUse 、 PostToolUse 、 Stop 、および (勧告のみ)) で起動します。
以下の ConfigChange ウォッチを参照してください) ConfigChange 。
5 つの事前チェックで 22 回の事前チェックでブロックします (出口 2、クロード コードはブロックと再試行として扱います)。
ファミリと 14 のターン終了ゲート (さらに、gate.stale_establisher と Gate.undeclared_falsifiable 、
この 14 ゲート契約の範囲外にある 2 つの永久勧告カタログ完全性チェック: 16 がライブ
ストップ層チェックの合計)。すべての事前チェックと、2 ブロックを除く 14 個のターン終了ゲートのすべて。
これらに対してサイレントの「警告」層はありません (「火災レベル」を参照)。二人は文書化した
例外は、gate.self_wired およびgate.canon_fingerprints_advisory (下記)、アドバイザリーのみのチェックです。
設計上、決してブロックされません。
検証者の弱体化: チェックはサイレントに無効化されます
content.verifier_predicate_weakened ルーズコンパレータ検証器 (startswith / endwith / re.match ここで == を意味します)
content.verifier_exit_masking 終了コード マスキング ( || true 、 ; true 、テスト/ビルド/lint で +e を設定)
content.verifier_body_hollowed 中空検証本体 (True を返す / 構成チェックに合格する)
content.env_gated_audit 環境変数の背後でゲートされた監査/検証コード · content.integrity_suppression_flag 整合性名の抑制フラグ ( *_skip = true )
捏造された証拠 : 裏付けとなる成果物のない主張
content.phantom_cycling ファントム引用 (著者と年が docs/CITATIONS.md に含まれていない)
content.unsourced_webfetch 以前のツールでは見られなかった URL の WebFetch 結果がこのセッションで発生します
content.fabricated_commit_sha 捏造されたコミット SHA/タグがコミットの証拠として提示されました
content.deferred_checkbox_theater DEFERRED - 開いている To Do アイテムのスタイルのチェックボックス シアター
詐欺

tent.illusory_authorship_trailer エージェントが実際には書いていない架空の AI 著者コミット トレーラー
Security-verifier の無効化: それぞれに名前付き CWE
content.cert_verify_disabled / content.cert_none_mode / content.cert_reqs_none TLS / ピア証明書検証が無効 ( verify=False 、 CERT_NONE )
content.jwt_signature_disabled / content.jwt_none_alg JWT 署名検証オフ / アルゴリズムはホワイトリストに登録されていません
content.paramiko_host_key_weakened paramiko SSH ホストキーチェックが AutoAddPolicy に弱体化 · content.timing_unsafe_compare シークレット/HMAC/ダイジェストのtiming-unsafe ==
content.self_mute_guard マコトのセルフミュート ( settings.json 経由でマコトの無効化または配線解除)
範囲と契約規律: 幻想的な進捗状況と契約外のアクション (SPEC-5、Assay からの形状移植)
event.thrash_revert ファイル全体の書き込み。別の書き込みが介在した後、ファイルを以前のバイト同一のコンテンツに戻します (A→B→A、最終的な進行なし)。
event.forbidden_location ターゲットが作業ディレクトリの外、または保護されたシステム/資格情報の場所に解決される書き込み/編集
Gate.contract_order そのステップに対する宣言されたプランの依存関係がまだ解消されていない間に発行された結果生成呼び出し (停止時間の兄弟ゲートがターン終了時に残りを保護します)
ターン終了ゲート : 記録された台帳と照合して、エージェントの終了要求に発砲します。
Gate.completion : "done / created X " ですが、アーティファクトはディスク上にありません
Gate.advance : 前提条件が満たされたと記録されていないフェーズを進める
Gate.green_claim : 記録されたテスト失敗に対する「スイート グリーン」
Gate.dropped : 識別情報 (「デフォルト X を Y に追加します」、「Z で 3 つのテスト」) を伴う前方約束がターンエンドで解除されずに残され、言ったが実行されず、エージェント自身のタッチキーと照合されます。
Gate.fabricated_action : "X を実行しました

/ ツール呼び出しがまったくないターンで Y " を実行しました (コマンドテキストの一致ではなく、作業の存在)
Gate.named_test : その正確な名前付きテストの記録された FAILED に対して " test_foo が合格" (相互参照で固定されている、green_claim デルタ)
Gate.stale_pass : pytest 自身の lastfailed レコードに対して「すべてのテストが合格」し、まだライブで失敗したテストに名前を付けています (ランナーの台帳、古いために存在がフィルターされています)
Gate.liveness : クレームではなくコード自体で起動します。クローズされた関数内でライブ効果のないステートメントです (現在のクロージャ モデルが不活性であることが証明できる死んだ/幻想的な作業)。ターンのタッチされた .py AST を歩くため、ターンごとに 1 つではなく、錯覚ステートメントごとに結果が得られます。
Gate.self_wired : 勧告のみで、決してブロックしません (「火災レベル」を参照)。 makoto 自身の .claude/settings.json フック配線の部分ストリップ検出: PreToolUse / PostToolUse / Stop に makoto-dispatching エントリがないのに、他の少なくとも 1 つにまだ makoto-dispatching エントリがある場合に発生します。これには文書化された盲点があります。3 つすべてを同時に削除する 1 回の編集により、このチェックが実行されるのと同じ瞬間にこのチェックが無効になります (クロード コードはフック設定をセッション開始時に一度ではなく、ライブでリロードします)。そのため、フルストリップのケースに対してはまったくカバーされません。以下の ConfigChange ウォッチのセクションを参照してください。これにより、配線時にこのギャップが正確に解消されます。
Gate.hollow_test : クレームではなくコード自体で起動します。HOLLOWED テスト (SPIRIT.md §4)、内容は破壊されても名前だけは残るため、実際に失敗することはありません。 4 つのサブパターン: テスト本体にいかなる種類のアサーションもありません。アサートされたトートロジー (assert True 、または式をそれ自体と比較する)。 / テスト対象の唯一の呼び出しの失敗を黙って飲み込むことを除いて、広範な no-op try。独立して起動できないテスト形式の関数。別の関数内にネストされています (pytest の c

ollector はそれを決して発見しません)、またはおそらく常に true である Skipif / SkipIf 条件の背後でゲートされます。
Gate.canon : ブロック中 (以下の勧告例外とは異なり、 level="error" )。ターン自身の呼び出しストリーム上の 2 つの言語に依存しない停止プリミティブ。クローズド プロトコル フィールド (tool_name/tool_input ID、tool_response.interrupted / .error) のみを読み取り、テスト ランナー正規表現や言語トークンは読みません。 canon.timeout は、ターンの最後のツール呼び出しが直接エラー状態 (中断または自己生成エラー コード) で終了し、再表示または解決せずに終了したときに発生します。その後の成功した呼び出しによって解決されたエラーはサイレントのままです。 canon.recur は、同じツール呼び出し (同一の名前 + バイト同一の入力) が連続して再発行され、その連続実行内のすべての呼び出しが同じ直接エラー状態で終了した場合に起動されます。つまり、試行間で何も変更されないスタック再試行ループです。同じ実行で後で成功すると、それは沈黙します。保持された敵対的 RED フィクスチャ (発火する必要がある植え付けられたケース) と、ほぼ空のコーパス FP チェックによって認証されます。正直なコーパスは、トリガーとなる前提条件 (中断 / self_error_code ) をまったく持たないことがほとんどないため、クリーンなコーパスの再生だけでは決定的ではなく、認証とはなりません (修正されたフレーミングについては、ANCESTRY.md §2 パート B を参照してください)。
Gate.canon_fingerprints : ブロッキング。移植されたセッションレベルの「canon」フィンガープリント 17 個のうち 4 個 (SPEC-5 タスク 9、 REF-lever-graded-primitives/signalminer/grade_planted.py の 27-fingerprint THE_CANON から) は堅牢なコアであり、構造的にブロッキング可能です: gold-oracle 認定に従って、plant-clean および real-Claude-gold ネガティブ セットの両方で 0-FP ( nogreen_checkdisabled 、 nosrc_destruct 、 nosrc_green_timeout 、 notestedit_destruct )。
Gate.canon_fingerprints_advisory : 勧告のみであり、ブロックすることはありません。同じ1のうちの残りの13

7 つのフィンガープリント。それぞれが、ゴールドオラクルの調査結果が堅牢であると認定していないソフト/クレーム アトム (claimed_pa​​ss_no_run 、tool_timeout 、assertation_weakened ) に基づいているか、その調査結果で明示的に名前が指定されている最悪の失格フィンガープリントの 1 つ上にあります。 SPEC-5 の合計保持ルールに従って監査ログに記録され、ブロック決定として出力されることはありません。
Gate.contract_order : ブロッキング。停止時間の残りは、宣言された計画を保護します (SPEC-5、Assay の ContractOrder パターンから形状を移植)。プランの依存関係の残りが空でない状態でターンが終了したときに発生します。その PreToolUse 兄弟ガードは事前チェックであり、停止ゲートではありません。 makoto/checks/contractOrder.py を参照してください。
Gate.stale_establisher : 勧告のみ、ブロックはしない、オプトイン。宣言された Plan ノードは DONE を記録しましたが、その名前付きアーティファクトはディスク上に存在しません (このチェック ファミリが作成する 1 つのファイルシステム読み取り、Assay の stale_establisher からのシェイプによる移植)。これをブロックにエスカレートする製品の決定は、呼び出し元に委ねられます。ここでは作られていません。
Gate.undeclared_falsifiable : 勧告のみであり、ブロックすることはありません。チェック/カタログ自体に対する完全性監査: ローダーが検出できない孤立したモジュール、または対応するライブ モジュールのない宣言された ID。メンテナンス信号

[切り捨てられた]

## Original Extract

When Claude says “tests pass”, makoto checks the record — every claim held against the agent's own logged deeds; fakes blocked, not warned. Each check ships at measured zero false positives. 誠 - Clear-Sights/Makoto

GitHub - Clear-Sights/Makoto: When Claude says “tests pass”, makoto checks the record — every claim held against the agent's own logged deeds; fakes blocked, not warned. Each check ships at measured zero false positives. 誠 · GitHub
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
Clear-Sights
/
Makoto
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
21 Commits 21 Commits .claude-plugin .claude-plugin .github .github canon canon checks checks commands commands data data docs docs hooks hooks lib lib prechecks prechecks stopchecks stopchecks tests tests tools tools .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md LICENSE LICENSE README.md README.md __init__.py __init__.py __main__.py __main__.py _dispatch.py _dispatch.py _dispatch_configchange.py _dispatch_configchange.py _dispatch_shim.sh _dispatch_shim.sh ackblock.py ackblock.py audit.py audit.py citations.py citations.py commitments.py commitments.py configchange_verdict.py configchange_verdict.py db.py db.py install.py install.py ledger.py ledger.py lexicons.py lexicons.py plan.py plan.py posture.py posture.py pyproject.toml pyproject.toml receipt.py receipt.py retraction.py retraction.py schema.py schema.py state.py state.py wire.py wire.py View all files Repository files navigation
An integrity hook for Claude Code that watches the agent's own tool calls and blocks the ones
that fake a check. When Claude says it did something (ran the tests, cited a paper, committed the
fix, verified the certificate), makoto holds that word against its record. If the deed isn't there,
or the verification was quietly disabled, makoto blocks the tool call (or the end-of-turn) and hands
the agent a one-line correction to retry against.
It judges the agent against its own utterances and record, never the world's truth. It holds no
facts ("France doesn't exist to it"); it only checks that a claimed word is kept, whole, and honored
in deed. A word it lets through becomes spendable: trustworthy tender a reviewer or another agent can
accept without re-deriving it.
makoto fires on mechanical hook events: every PreToolUse , PostToolUse , Stop , and (advisory-only;
see ConfigChange watch below) ConfigChange .
It blocks (exit 2, which Claude Code treats as block-and-retry) on 22 pre-checks across five
families and 14 end-of-turn gates (plus gate.stale_establisher and gate.undeclared_falsifiable ,
two permanently-advisory catalog-completeness checks outside this 14-gate contract entirely: 16 live
Stop-tier checks in total). Every pre-check and every one of the 14 end-of-turn gates but two blocks;
there is no silent "warning" tier for those (see Fire level ). The two documented
exceptions are gate.self_wired and gate.canon_fingerprints_advisory (below), advisory-only checks
that by design never block.
Verifier weakening : a check silently neutered
content.verifier_predicate_weakened loose-comparator verifier ( startswith / endswith / re.match where == is meant)
content.verifier_exit_masking exit-code masking ( || true , ; true , set +e on a test/build/lint)
content.verifier_body_hollowed hollowed verifier body ( return True / pass in a constitution check)
content.env_gated_audit audit/verification code gated behind an env var · content.integrity_suppression_flag integrity-named suppression flag ( *_skip = true )
Fabricated evidence : a claim with no backing artifact
content.phantom_citation phantom citation (Author-Year not in docs/CITATIONS.md )
content.unsourced_webfetch WebFetch of a URL never seen in any prior tool result this session
content.fabricated_commit_sha fabricated commit SHA/tag presented as proof of a commit
content.deferred_checkbox_theater DEFERRED -style checkbox theater on an open to-do item
content.illusory_authorship_trailer an illusory AI-authorship commit trailer the agent never actually wrote
Security-verifier disabling : each a named CWE
content.cert_verify_disabled / content.cert_none_mode / content.cert_reqs_none TLS / peer-certificate verification disabled ( verify=False , CERT_NONE )
content.jwt_signature_disabled / content.jwt_none_alg JWT signature verification off / none algorithm whitelisted
content.paramiko_host_key_weakened paramiko SSH host-key check weakened to AutoAddPolicy · content.timing_unsafe_compare timing-unsafe == on a secret/HMAC/digest
content.self_mute_guard makoto self-mute (disabling or un-wiring makoto via settings.json )
Scope & contract discipline : illusory progress and out-of-contract action (SPEC-5, ported by shape from Assay)
event.thrash_revert a whole-file Write that reverts a file to an earlier byte-identical content after an intervening different Write (A→B→A, no net progress)
event.forbidden_location a Write/Edit whose target resolves outside the working directory or into a protected system/credential location
gate.contract_order a result-producing call issued while a declared Plan's dependency for that step is still undischarged (its Stop-time sibling gate guards the remainder at turn end)
End-of-turn gates : fire on the agent's closing claims, checked against the recorded ledger:
gate.completion : "done / created X " but the artifact isn't on disk
gate.advance : advancing a phase whose precondition isn't recorded as met
gate.green_claim : "suite green" against a recorded test failure
gate.dropped : a forward promise carrying identifying info ("I'll add def X to Y ", "3 tests in Z ") left undischarged at turn-end, said-but-not-done, checked against the agent's own touched-keys
gate.fabricated_action : "I ran X / executed Y " in a turn with no tool call at all (presence-of-work, not command-text match)
gate.named_test : " test_foo passes" against a recorded FAILED of that exact named test (coreference-pinned; the green_claim delta)
gate.stale_pass : "all tests pass" against pytest's own lastfailed record still naming a live failing test (the runner's ledger, existence-filtered for staleness)
gate.liveness : fires on the code itself, not a claim: a statement with no live effect inside a closed function (dead/illusory work the present-closure model can prove inert). It walks the turn's touched .py ASTs, so it yields a finding per illusory statement rather than one per turn.
gate.self_wired : advisory only, never blocks (see Fire level ). Partial-strip detection of makoto's own .claude/settings.json hook wiring: fires if PreToolUse / PostToolUse / Stop is missing a makoto-dispatching entry while at least one other still has one. It has a documented blind spot: a single edit that strips all three simultaneously disables this check in the same instant it would have fired (Claude Code reloads hook config live, not once at session start), so it provides zero coverage against that full-strip case. See the ConfigChange watch section below, which closes exactly this gap when wired.
gate.hollow_test : fires on the code itself, not a claim: a HOLLOWED test (SPIRIT.md §4), one that survives in name while its content is gutted, so it can never actually fail. Four sub-patterns: no assertion of any kind in the test body; an asserted tautology ( assert True , or comparing an expression to itself); a broad, no-op try / except that silently swallows the only call-under-test's failure; and a test-shaped function that can never fire independently, either nested inside another function (pytest's collector never discovers it) or gated behind a skipif / skipIf condition that is provably always true.
gate.canon : blocking ( level="error" , unlike the advisory exceptions below). Two language-agnostic Stop primitives over the turn's own call stream, reading only closed protocol fields (tool_name/tool_input identity, tool_response.interrupted / .error ), no test-runner regex, no language token. canon.timeout fires when the turn's LAST tool call ended in a direct error state (interrupted, or a self-emitted error code) and closed without resurfacing or resolving it; an error resolved by a later successful call stays silent. canon.recur fires when the SAME tool call (identical name + byte-identical input) was re-issued back-to-back and every call in that consecutive run ended in the same direct error state: a stuck retry loop with nothing changed between attempts; a later success in the same run silences it. Certified via held-out adversarial RED fixtures (planted cases that must fire) plus a near-vacuous corpus-FP check: the honest corpus almost never carries the triggering precondition ( interrupted / self_error_code ) at all, so a clean corpus replay alone is inconclusive, not a certification (see ANCESTRY.md §2 Part B for the corrected framing).
gate.canon_fingerprints : blocking . 4 of 17 ported session-level "canon" fingerprints (SPEC-5 Task 9, from REF-lever-graded-primitives/signalminer/grade_planted.py 's 27-fingerprint THE_CANON ) that are robust-core, blocking-capable by construction: 0-FP on both the planted-clean and real-Claude-gold negative sets per the gold-oracle certification ( nogreen_checkdisabled , nosrc_destruct , nosrc_green_timeout , notestedit_destruct ).
gate.canon_fingerprints_advisory : advisory only, never blocks . The other 13 of the same 17 fingerprints, each either resting on a soft/claim atom ( claimed_pass_no_run , tool_timeout , assertion_weakened ) the gold-oracle finding doesn't certify as robust, or one of that finding's explicitly-named worst-disqualified fingerprints. Recorded to the audit log per SPEC-5's total-retention rule, never emitted as a block decision.
gate.contract_order : blocking . The Stop-time remainder guard over a declared Plan (SPEC-5, ported by shape from Assay's ContractOrder pattern); fires when the turn ends with the plan's dependency remainder still non-empty. Its PreToolUse sibling guard is a pre-check, not a Stop gate; see makoto/checks/contractOrder.py .
gate.stale_establisher : advisory only, never blocks , opt-in. A declared Plan node recorded DONE whose named artifact no longer exists on disk (the one filesystem read this check family makes, ported by shape from Assay's stale_establisher ). A product decision to escalate this to blocking is left to the caller; it is not made here.
gate.undeclared_falsifiable : advisory only, never blocks . A completeness audit over the checks/ catalog itself: an orphan module the loader can't discover, or a declared id with no corresponding live module. A maintenance signal a

[truncated]
