---
source: "https://blog.fidelramos.net/software/emacs-straight-ai-review"
hn_url: "https://news.ycombinator.com/item?id=49220254"
title: "Supply-chain hygiene for Emacs: LLM review of package upgrades"
article_title: "blog.fidelramos.net\n– Supply-chain hygiene for Emacs: LLM review of package upgrades"
author: "fidelramos"
captured_at: "2026-08-08T10:21:34Z"
capture_tool: "hn-digest"
hn_id: 49220254
score: 1
comments: 0
posted_at: "2026-08-08T09:48:55Z"
tags:
  - hacker-news
  - translated
---

# Supply-chain hygiene for Emacs: LLM review of package upgrades

- HN: [49220254](https://news.ycombinator.com/item?id=49220254)
- Source: [blog.fidelramos.net](https://blog.fidelramos.net/software/emacs-straight-ai-review)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T09:48:55Z

## Translation

タイトル: Emacs のサプライチェーン衛生: パッケージアップグレードの LLM レビュー
記事のタイトル: blog.fidelramos.net
– Emacs のサプライチェーン衛生: パッケージアップグレードの LLM レビュー
説明: Emacs にはサンドボックスがありません。インストールするすべてのパッケージは、ファイル、資格情報、ネットワークなどの完全な権限で任意の Lisp を実行します。そのため、新しいパッケージをインストールしたり、既存のパッケージをアップグレードしたりするのが本当に怖くなります。私たちのほとんどが使用するパッケージは精査されたアーティファクトではありません。MELPA はいくつかのソースからビルドします。
[切り捨てられた]

記事本文:
blog.fidelramos.net
– Emacs のサプライチェーン衛生: パッケージアップグレードの LLM レビュー
フィデル・ラモス
Emacs のサプライチェーンの衛生管理: パッケージのアップグレードに関する LLM レビュー
投稿日: 2026-08-08 in ソフトウェア
Emacs にはサンドボックスがありません。インストールするすべてのパッケージは、完全な状態で任意の Lisp を実行します。
権限: ファイル、資格情報、ネットワーク。これは本当に怖いです
新しいパッケージをインストールするか、既存のパッケージをアップグレードします。
私たちのほとんどが使用するパッケージは精査されたアーティファクトではありません: MELPA は何らかのリポジトリからビルドします
(通常は署名されていない) コミット、多くのパッケージは 1 人によって維持され、誰も維持されません。
特定の更新に何が含まれるかを監査します。メンテナのアカウント乗っ取り、またはリポジトリ
新しい所有者に売却されると、M-x package-upgrade-all がリモート コード実行に変わります。
機械。このエコシステムにはさらに微妙なベクトルがあり、それはレシピです。
リダイレクト。 MELPA レシピを 1 行変更するだけで、既存のパッケージを暗黙的にポイントできる
別のリポジトリで名前を変更します。
答えは、アップグレードをやめることではなく、アップグレードを計画的に行うことです。つまり、すべてのパッケージを固定することです。
正確にコミットして、何かをマージする前に差分を確認してください。
手動で差分を確認するのは面倒なだけでなく、高度な専門知識が必要です。
Emacs Lisp の場合、私の解決策は、LLM に Emacs Lisp の初回のセキュリティ監査を行わせることです。
実行前に変更されます。
この投稿では、私が使用しているセットアップを紹介します。
ストレートエル、
Magit と gptel 。
ステップ1：straight.elブートストラップを確認する
ステップ 3: マージ前に差分を確認する
ステップ 4: LLM を利用した差分レビュー
Straight.el 独自のブートストラップ スクリプトは、固定された SHA-256 チェックサムに対して検証されます
評価される前に。
すべてのパッケージ (およびすべてのレシピ リポジトリ) は、次のロックファイル内のコミットに固定されます。
バージョン管理。
アップグレードはフェッチファーストです。 git fetch で新しく受信した変更を取得します。

git マージと
実際にはレビューに合格したもののみをアップグレードします。
カスタム Magit ビューには、フェッチされたもののマージされていない受信差分がすべて表示されます。
ストレートなリポジトリ。
カスタム gptel コマンドは、LLM を使用して完全な diff をレビューします。
変更がクリーンな場合は、straight.el で変更をマージし、
ロックファイル。
gptelがstraight.elパッケージのアップグレードをレビュー
ステップ1：straight.elブートストラップを確認する
Straight.el がインストールします
ダウンロードすることでそれ自体
そして起動時に install.el を評価します。それは額面通りのリモート コードなので、
early-init.el は SHA-256 を固定し、不一致で停止します。
;; Straight.el の install.el の SHA-256。での変更を確認した後、更新します
;; https://github.com/radian-software/straight.el/commits/develop/install.el
( defconst my-straight-install-el-sha256
"e29e07d52d16d4136971f0a822cb6a1a6e1e764a1cb9fe67cccbc7c048aba553")
( defvar ブートストラップ バージョン )
( let (( ブートストラップファイル
( 展開ファイル名
「straight/repos/straight.el/bootstrap.el」
(または (bound-and-true-p ストレートベースディレクトリ)
ユーザー emacs ディレクトリ )))
(ブートストラップバージョン 7 ))
(ただし ( file-exists-p bootstrap-file )
( 現在のバッファ付き
( URL 取得同期
「https://raw.githubusercontent.com/radian-software/straight.el/develop/install.el」
「サイレント」Cookie を禁止します)
;;ブートストラップ スクリプトを評価前にピン留めされたチェックサムと照合して検証する
( 'url-http が必要です)
( url-http-end-of-headers の場合
( delete-region ( point-min ) url-http-end-of-headers )
( 削除領域 ( ポイント分 )
( progn ( Skip-chars-forward " \t\n" ) ( point ))))
( let (( checksum ( secure-hash 'sha256 ( current-buffer ))))
(ただし ( string= checksum my-straight-install-el-sha256 )
(エラー「straight.el ブートストラップ チェックサムが一致しません!」)))
( goto-char ( point-max ))
( eval-print-last-sex )))
(ブートストラップファイルをロード nil 'nomessage ))
コピー
ステップ 2: すべてを固定する
ストレート.e

l はパッケージを git クローンとしてインストールするため、レビューが自然になります。ロックファイルのピン
それらすべて:
M-x Straight-freeze-versions は、すべてのパッケージの正確なコミットを次の場所に書き込みます。
Straight/versions/default.el (パッケージとレシピ リポジトリ、MELPA)
含まれています）。
ロックファイルをコミットします。実行後に git diff Straight/versions/default.el を確認する
アップグレードすると、何が移動したかが正確にわかります。
M-x ストレート解凍バージョンは、ロールバックする必要がある場合にそれらのリビジョンを復元します。
ステップ 3: マージ前に差分を確認する
安全なアップグレードの鍵は、フェッチとマージが別のステップであることです。 M-x
Straight-fetch-all は、チェックアウトを変更せずに、新しいアップストリーム コミットをダウンロードします。それから私は
M-x Straight-merge-all が何かを適用する前に、リポジトリごとに何が変わるかを確認してください。
次に、カスタム Magit 関数を使用して、すべてのリポジトリの単一の結合された差分ビューを表示します。
受信した上流コミットの場合:
;; https://magit.vc/
( パッケージ magit を使用
:bind (( "C-c v U" . my-straight-incoming-diffs ))
:序文
( defun my-straight-repo-behind-p ( repo )
「REPO の現在のブランチが上流の後ろにある場合は、非 nil を返します。」
( let* (( デフォルトディレクトリリポジトリ )
( 行 ( magit-git-lines "ステータス" "--porcelain=2" "--branch" )))
( seq-some ( lambda ( line )
( string-match-p "^# ブランチ\\.ab [+][0-9]+ -[1-9]" 行 ))
行)))
( defun my-straight-incoming-diffs ()
「受信したすべてのアップストリーム変更の完全なパッチを連結します。
すべての Straight.el チェックアウトに対して、「diff-mode」バッファー リストを作成します。
それは上流の背後にあり、作成者との受信コミットです。
日付の後に、マージが適用されるネット パッチが続きます。の
バッファは書き込み可能のままなので、フィードする前にハンクをトリミングできます
審査エージェントに送信します。」
(インタラクティブ)
( let* (( リポス ( magit-list-repos-1
(展開ファイル名 "straight/repos" ユーザー emacs ディレクトリ )
1))
( ( seq-filter #' my-straight-repo の後ろ)

-behind-p リポジトリ )))
( if ( 後ろに null )
(メッセージ「受信アップストリーム変更はありません」)
( with-current-buffer ( get-buffer-create "*straight-incoming-diffs*" )
(消去バッファ)
( dolist (リポジトリの後ろ)
( let* (( デフォルトディレクトリリポジトリ )
( ブランチ ( magit-git-string "rev-parse" "--abbrev-ref"
「頭」 ))
(上流 ( magit-git-string "rev-parse" "--abbrev-ref"
"@{アップストリーム}" ))
( コミット ( magit-git-lines
"ログ" "--format=%h %ad %an <%ae> %s"
"--date=short" "HEAD..@{upstream}" ))
( n ( コミットの長さ )))
(挿入 (make-string 80 ?= ) "\n" )
(挿入
( フォーマット "%s (%s <- %s, %d commit%s)\n"
( ファイル名非ディレクトリ ( ディレクトリファイル名 repo ))
(または分岐 "HEAD" ) (または上流 "?" )
n ( if ( = n 1 ) "" "s" )))
(挿入 (make-string 80 ?-) "\n" )
(挿入 (文字列結合コミット "\n" ) "\n\n" )
( magit-git-insert "diff" "HEAD...@{upstream}" )
(「\n」を挿入)))
( 差分モード )
( goto-char ( point-min ))
( Pop-to-buffer ( current-buffer ))))))
( defun my-straight-incoming-diffs-after-fetch (&rest _)
「インタラクション後に受信した完全な差分を表示する」
[切り捨てられた]
diff が本当に単純なものでない限り、コードレビューのために LLM に渡します。私は使用します
秘密保管用の KWallet と Venice を備えた gptel (紹介
リンク ) AI 推論の場合:
;; KWallet からの API キー (freedesktop Secret Service API 経由)。または置く
;;マシン api.venice.ai ログイン apikey パスワード <your-key>
;; ~/.authinfo.gpg 内で Secret: エントリを削除します。
( setq auth-sources ' ( "secrets:kdewallet" "~/.authinfo.gpg" "~/.authinfo" ))
( パッケージ gptel を使用
:bind (( "C-c a R" . my-gptel-review-malicious-code ))
:序文
( defvar my-gptel-review-system-prompt
「あなたはサードパーティのコードを監査する細心の注意を払ったセキュリティレビュー担当者です
マージまたは使用される前に。ユーザーは 1 つ以上の git を表示します
diff、それぞれの前に受信コミットのリストが続きます

著者と
日付。
マージまたは呼び出されると悪意を持って動作する可能性のあるものをスキャンします。
- バックドア、リモートコード実行、永続化メカニズム
- 資格情報、トークン、キー、またはデータの漏洩 (秘密チャネルも)
- 予期しないネットワークアクティビティ、ダウンロード、または新しいホストへの接続
- 動作を隠すように設計された難読化 (重いエンコーディング、デッドストア)
- ビルド、インストール、またはパッケージ化スクリプトの改ざん
- 異常なコミットの作成者: 新しいメンテナ、変更された電子メール アドレス、
プロジェクトに関係のない人によるコミット
すべての検出結果について、重大度 (高/中/低)、ファイル、および
ハンクと 1 段落の根拠。全体的な判断で締めくくります。
変更をマージしても安全であるかどうか。何も疑わしいところがなければ、
はっきりと簡潔に言ってください。発見結果をでっち上げないでください。」
「「my-gptel-review-malicious-code」のシステム プロンプト。」 )
( defun my-gptel-review-malicious-code ()
「領域またはバッファ全体に悪意のあるコードがないか確認してください。
セキュリティ監査システム プロンプトを使用してテキストを LLM に送信し、
*gptel-review* バッファー内の応答を表示します。」
(インタラクティブ)
( 'gptel が必要です)
( let* (( text ( バッファー部分文字列のプロパティ
( if ( use-region-p ) ( area-beginning ) ( point-min ))
( if ( use-region-p ) ( area-end ) ( point-max ))))
(バッファ ( get-buffer-create "*gptel-review*" ))
( マーカー ( with-current-buffer バッファ
(消去バッファ)
(組織モード)
( goto-char ( point-min ))
( ポイントマーカー ))))
( バッファへのポップ
[切り捨てられた]
M-x Straight-fetch-all : リモートをフェッチし、何も変更しません。レビューバッファがポップアップします
それ自体で。
C-c バッファー全体、または選択した領域に対する R。
満足したら M-x ストレートマージオール。
Emacs を再起動し、すべてが動作することを確認します (straight.el によって新しいパッケージがビルドされます)。
M-x ストレートフリーズバージョンを実行し、ロックファイルをコミットします。
LLM 監査はセコ

2 番目の目、保証ではありません。モデルがミス
そして、賢い敵は無害であると読み取れるコードを書きます。クリティカルな場合
パッケージがある場合でも、差分は自分で読んでください。監査をトリアージとして扱います: 「高」
見つけるということは、モデルが正しいということではなく、よく見るということです。
差分はマシンから離れます。ヴェネツィアはデータ保持ゼロを宣伝していますが、もし
あなたの脅威モデルではそれが許可されていません。gptel をローカルの llama.cpp サーバーに向けてください
代わりに。コードは同じですが、gptel プロバイダーが異なるだけです。
この設定が完璧ではないことは承知していますが、次のような場合に非常に安心感が得られます。
Emacs が私のコンピュータセットアップの中心部分であるため、パッケージをインストールまたはアップグレードしています
そして私はそれを完全に信頼する必要があります
盲点、欠点、またはより良い設定について何かアイデアはありますか?以下にコメントを書き込んでください。
emacs フリーソフトウェア ハウツー プログラミング セキュリティ
© 2026 - この作品はクリエイティブ コモンズ 表示 - 継承に基づいてライセンスされています。
Reflex テーマを使用して Pelican で構築
|
暗闇に切り替えます |ライト |ブラウザのテーマ

## Original Extract

Emacs has no sandbox. Every package you install runs arbitrary Lisp with your full privileges: your files, your credentials, your network. This makes me really scared of installing new packages, or upgrading existing ones. The packages most of us use are not vetted artifacts: MELPA builds from some
[truncated]

blog.fidelramos.net
– Supply-chain hygiene for Emacs: LLM review of package upgrades
Fidel Ramos
Supply-chain hygiene for Emacs: LLM review of package upgrades
Posted on 2026-08-08 in Software
Emacs has no sandbox. Every package you install runs arbitrary Lisp with your full
privileges: your files, your credentials, your network. This makes me really scared of
installing new packages, or upgrading existing ones.
The packages most of us use are not vetted artifacts: MELPA builds from some repository
(usually unsigned) commit, many packages are maintained by a single person, and nobody
audits what lands in a given update. A maintainer's account takeover, or a repository
sold to a new owner, turns M-x package-upgrade-all into remote code execution on your
machine. There is even a subtler vector unique to this ecosystem: recipe
redirection . A one-line change to a MELPA recipe can silently point an existing package
name at a different repository.
The answer isn't to stop upgrading, but to make upgrades deliberate : pin every package
to an exact commit and review the diffs before merging anything.
Manually reviewing diffs is not only cumbersome but needs a high level of expertise in
Emacs Lisp, so my solution is to have an LLM do a first-pass security audit of the
changes before they get to execute.
This post showcases the setup I use, built on
straight.el ,
Magit and gptel .
Step 1: verify the straight.el bootstrap
Step 3: review diffs before merging
Step 4: LLM-powered review of the diff
straight.el 's own bootstrap script is verified against a pinned SHA-256 checksum
before it is evaluated.
Every package (and every recipe repository) is pinned to a commit in a lockfile under
version control.
Upgrades are fetch-first: git fetch to pull new incoming changes, git merge and
actually upgrade only what passes review.
A custom Magit view displays all incoming fetched-but-unmerged diffs across all
straight repos.
A custom gptel command reviews the full diff using an LLM.
If the changes are clean, have straight.el merge the changes and update the
lockfile.
gptel reviewing a straight.el packages upgrade
Step 1: verify the straight.el bootstrap
straight.el installs
itself by downloading
and evaluating install.el at startup. That's remote code at face value, so my
early-init.el pins its SHA-256 and stops on mismatch:
;; SHA-256 of straight.el's install.el; update after reviewing changes at
;; https://github.com/radian-software/straight.el/commits/develop/install.el
( defconst my-straight-install-el-sha256
"e29e07d52d16d4136971f0a822cb6a1a6e1e764a1cb9fe67cccbc7c048aba553" )
( defvar bootstrap-version )
( let (( bootstrap-file
( expand-file-name
"straight/repos/straight.el/bootstrap.el"
( or ( bound-and-true-p straight-base-dir )
user-emacs-directory )))
( bootstrap-version 7 ))
( unless ( file-exists-p bootstrap-file )
( with-current-buffer
( url-retrieve-synchronously
"https://raw.githubusercontent.com/radian-software/straight.el/develop/install.el"
'silent 'inhibit-cookies )
;; verify bootstrap script against its pinned checksum before eval
( require 'url-http )
( when url-http-end-of-headers
( delete-region ( point-min ) url-http-end-of-headers )
( delete-region ( point-min )
( progn ( skip-chars-forward " \t\n" ) ( point ))))
( let (( checksum ( secure-hash 'sha256 ( current-buffer ))))
( unless ( string= checksum my-straight-install-el-sha256 )
( error "straight.el bootstrap checksum mismatch!" )))
( goto-char ( point-max ))
( eval-print-last-sexp )))
( load bootstrap-file nil 'nomessage ))
Copy
Step 2: pin everything
straight.el installs packages as git clones, which makes review natural. The lockfile pins
them all:
M-x straight-freeze-versions writes every package's exact commit to
straight/versions/default.el (packages and recipe repositories, MELPA
included).
Commit the lockfile. Reviewing git diff straight/versions/default.el after an
upgrade tells you exactly what moved.
M-x straight-thaw-versions restores those revisions if you ever need to roll back.
Step 3: review diffs before merging
The key to a safe upgrade is that fetching and merging are separate steps. M-x
straight-fetch-all downloads new upstream commits without changing any checkout. Then I
review what would change, per repo, before M-x straight-merge-all applies anything.
I then use a custom Magit function to display a single combined diff view of every repo
with incoming upstream commits:
;; https://magit.vc/
( use-package magit
:bind (( "C-c v U" . my-straight-incoming-diffs ))
:preface
( defun my-straight-repo-behind-p ( repo )
"Return non-nil if REPO's current branch is behind its upstream."
( let* (( default-directory repo )
( lines ( magit-git-lines "status" "--porcelain=2" "--branch" )))
( seq-some ( lambda ( line )
( string-match-p "^# branch\\.ab [+][0-9]+ -[1-9]" line ))
lines )))
( defun my-straight-incoming-diffs ()
"Concatenate the full patches of all incoming upstream changes.
Create a `diff-mode' buffer listing, for every straight.el checkout
that is behind its upstream, the incoming commits with author and
date, followed by the net patch that merging would apply. The
buffer is left writable so hunks can be trimmed before feeding it
to a reviewing agent."
( interactive )
( let* (( repos ( magit-list-repos-1
( expand-file-name "straight/repos" user-emacs-directory )
1 ))
( behind ( seq-filter #' my-straight-repo-behind-p repos )))
( if ( null behind )
( message "No incoming upstream changes" )
( with-current-buffer ( get-buffer-create "*straight-incoming-diffs*" )
( erase-buffer )
( dolist ( repo behind )
( let* (( default-directory repo )
( branch ( magit-git-string "rev-parse" "--abbrev-ref"
"HEAD" ))
( upstream ( magit-git-string "rev-parse" "--abbrev-ref"
"@{upstream}" ))
( commits ( magit-git-lines
"log" "--format=%h %ad %an <%ae> %s"
"--date=short" "HEAD..@{upstream}" ))
( n ( length commits )))
( insert ( make-string 80 ?= ) "\n" )
( insert
( format "%s (%s <- %s, %d commit%s)\n"
( file-name-nondirectory ( directory-file-name repo ))
( or branch "HEAD" ) ( or upstream "?" )
n ( if ( = n 1 ) "" "s" )))
( insert ( make-string 80 ?- ) "\n" )
( insert ( string-join commits "\n" ) "\n\n" )
( magit-git-insert "diff" "HEAD...@{upstream}" )
( insert "\n" )))
( diff-mode )
( goto-char ( point-min ))
( pop-to-buffer ( current-buffer ))))))
( defun my-straight-incoming-diffs-after-fetch ( &rest _ )
"Show incoming full diffs after an interact
[truncated]
Unless the diff is really straightforward I pass it to an LLM for a code review. I use
gptel with KWallet for secret storing and Venice ( referral
link ) for AI inference:
;; API key from KWallet via the freedesktop Secret Service API; or put
;; machine api.venice.ai login apikey password <your-key>
;; in ~/.authinfo.gpg and drop the secrets: entry.
( setq auth-sources ' ( "secrets:kdewallet" "~/.authinfo.gpg" "~/.authinfo" ))
( use-package gptel
:bind (( "C-c a R" . my-gptel-review-malicious-code ))
:preface
( defvar my-gptel-review-system-prompt
"You are a meticulous security reviewer auditing third-party code
before it is merged or used. The user will show you one or more git
diffs, each preceded by the list of incoming commits with author and
date.
Scan for anything that could act maliciously once merged or called:
- backdoors, remote code execution, persistence mechanisms
- credential, token, key or data exfiltration (also covert channels)
- unexpected network activity, downloads or connections to new hosts
- obfuscation designed to hide behavior (heavy encoding, dead stores)
- tampering with build, installation or packaging scripts
- anomalous commit authorship: new maintainer, changed email address,
commits by someone unrelated to the project
For every finding, state: severity (high/medium/low), the file and
hunk, and a one-paragraph rationale. Close with an overall verdict:
whether the changes look safe to merge. If nothing is suspicious,
say so plainly and briefly; do not invent findings."
"System prompt for `my-gptel-review-malicious-code' ." )
( defun my-gptel-review-malicious-code ()
"Review the region, or the whole buffer, for malicious code.
Send the text to the LLM with a security-audit system prompt and
show the response in the *gptel-review* buffer."
( interactive )
( require 'gptel )
( let* (( text ( buffer-substring-no-properties
( if ( use-region-p ) ( region-beginning ) ( point-min ))
( if ( use-region-p ) ( region-end ) ( point-max ))))
( buffer ( get-buffer-create "*gptel-review*" ))
( marker ( with-current-buffer buffer
( erase-buffer )
( org-mode )
( goto-char ( point-min ))
( point-marker ))))
( pop-to-buffer
[truncated]
M-x straight-fetch-all : fetches remotes, changes nothing; the review buffer pops up
on its own.
C-c a R on the whole buffer, or on a selected region.
M-x straight-merge-all when satisfied.
Restart Emacs, verify everything works ( straight.el would have built the new packages).
M-x straight-freeze-versions and commit the lockfile.
The LLM audit is a second pair of eyes, not a guarantee. Models miss
things, and a clever adversary writes code that reads as benign. For critical
packages, still read the diff yourself. Treat the audit as triage: a "high"
finding means look closer, not that the model is right.
Diffs leave your machine. Venice advertises zero data retention, but if
your threat model doesn't allow it, point gptel at a local llama.cpp server
instead. Code is the same, just a different gptel provider.
I know this setup won't be perfect, but it gives me much more peace of mind when
installing or upgrading packages, because Emacs is the center piece of my computer setup
and I need to trust it fully
Do you have any ideas on blind spots, shortcomings or a better setup? Drop a comment below!
emacs free-software howto programming security
© 2026 - This work is licensed under a Creative Commons Attribution-ShareAlike
Built with Pelican using Reflex theme
|
Switch to the dark | light | browser theme
