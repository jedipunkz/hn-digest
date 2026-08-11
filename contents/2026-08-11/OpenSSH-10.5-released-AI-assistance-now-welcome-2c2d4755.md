---
source: "https://www.openssh.org/releasenotes.html#10.5"
hn_url: "https://news.ycombinator.com/item?id=49261895"
title: "OpenSSH 10.5 released, AI assistance now welcome"
article_title: "OpenSSH: Release Notes"
author: "voxadam"
captured_at: "2026-08-11T18:48:03Z"
capture_tool: "hn-digest"
hn_id: 49261895
score: 24
comments: 8
posted_at: "2026-08-11T17:49:37Z"
tags:
  - hacker-news
  - translated
---

# OpenSSH 10.5 released, AI assistance now welcome

- HN: [49261895](https://news.ycombinator.com/item?id=49261895)
- Source: [www.openssh.org](https://www.openssh.org/releasenotes.html#10.5)
- Score: 24
- Comments: 8
- Posted: 2026-08-11T17:49:37Z

## Translation

タイトル: OpenSSH 10.5 がリリース、AI 支援を歓迎
記事のタイトル: OpenSSH: リリース ノート
説明: OpenSSH リリース ノート

記事本文:
OpenSSH 10.5 / 10.5p1 (2026-08-11)
OpenSSH 10.5 は 2026 年 8 月 11 日にリリースされました。から入手可能です。
ミラーは https://www.openssh.com/ にリストされています。
OpenSSH は 100% 完全な SSH プロトコル 2.0 実装であり、
SFTP クライアントとサーバーのサポートが含まれます。
最近、OpenSSH チームは大量のセキュリティを受け取りました。
バグレポート。その多くは AI モデルからの発見か、AI モデルを使用して作成されたものです。
AIによる支援。多くの AI レポートはそうではないと判断されていますが、
現実的な観点から考慮した場合のセキュリティへの影響
脅威モデルに関しては、特に次のような場合には、これらの報告を非常に歓迎します。
人間のトリアージ、分析、テストケース、特に
提案された修正を伴う場合。
私たちは、セキュリティ バグが特定されたケースを数多く見てきました。
AI ツールはその後、別のツールによって独立して発見されます。
研究者。これは、バグを報告しない攻撃者が、
OSS プロジェクトでもこれらのバグを発見できる可能性があります。
これを考慮して、OpenSSH チームは当面、より頻繁に
バグ修正をより迅速にユーザーの手に届けるためのリリース
次の計画されたリリースまでバッチ処理します。
改めて、OpenSSH コミュニティの皆様に感謝いたします。
プロジェクトへの継続的なサポート、特に貢献してくれた人たち
コードまたはパッチ、報告されたバグ、テストされたスナップショット、またはへの寄付
プロジェクト。寄付の詳細については、以下をご覧ください。
https://www.openssh.com/donations.html
互換性がない可能性のある変更
--------------------------------
* ポータブル OpenSSH には ECC (楕円曲線暗号) が必要になりました
libcrypto でのサポート (NISTP521 曲線のサポートを含む)。
ECC はすべてのデフォルトのビルド構成に含まれています。
現在サポートされているすべての libcrypto 実装のバージョン
OpenSSH (LibreSSL、OpenSSL、BoringSS を含む)

LとAWS LC。
--without-openssl ビルド構成は影響を受けません。
OpenSSH 10.4 以降の変更点
==========================
このリリースには、多数のセキュリティ修正と小さなバグ修正が含まれています。
セキュリティ
========
* ssh-agent(1) : インタラクションを修正
[切り捨てられた]
OpenSSH 10.4 は 2026 年 7 月 6 日にリリースされました。から入手可能です。
ミラーは https://www.openssh.com/ にリストされています。
OpenSSH は 100% 完全な SSH プロトコル 2.0 実装であり、
SFTP クライアントとサーバーのサポートが含まれます。
改めて、OpenSSH コミュニティの皆様に感謝いたします。
プロジェクトへの継続的なサポート、特に貢献してくれた人たち
コードまたはパッチ、報告されたバグ、テストされたスナップショット、またはへの寄付
プロジェクト。寄付の詳細については、以下をご覧ください。
https://www.openssh.com/donations.html
互換性がない可能性のある変更
--------------------------------
* sshd(8) : 設定ダンプ モード (「sshd -G」) がディレクティブを書き込むようになりました
大文字と小文字が混合された場合 (例: "PubkeyAuthentication")、以前は
小文字の名前のみを出力します。
* sshd(8) : seccomp サンドボックスが有効になっている Linux システム上では、
SECCOMP または NO_NEW_PRIVS を有効にしないと致命的になるようになりました。
以前は、sshd(8) はエラーをログに記録していましたが、操作は続行されました。
これらの機能が欠けているシステムをサポートするために。現在のシステムは、
これらが不足している場合は、構成時にサンドボックスを無効にする必要があります。
* ssh(1) 、 sshd(8) : トランスポート プロトコルをより厳密にします。
ピアがポスト中に非 KEX メッセージを送信した場合は切断します。
認証キーの再交換。以前は、悪意のあるピアによって、
ペナルティなしで非鍵交換メッセージを送信し続けます。これら
バッファリングされるため、メモリが浪費されることになります。
接続が終了したか、サーバー/クライアントがメモリ制限に達しました。
キー中に送信されるメッセージを制限しない実装
RFC4253 セクション 7.1 に従って交換することができます。

切断される。
マルコ・イェヴティッチ氏が報告した。
OpenSSH 10.3 以降の変更点
==========================
このリリースには、一般的な修正だけでなく、多数のセキュリティ修正も含まれています。
バグ修正といくつかの新機能。
セキュリティ
========
* sftp(1) : コマンドラインでファイルをダウンロードする場合
「sftp host:/path .」、悪意のあるサーバーによりファイルが
予期しない場所にダウンロードされる可能性があります。この問題が特定されました
Swivalセキュリティスキャナーによる。
* scp(1) : いつ
[切り捨てられた]
OpenSSH 10.3 は 2026 年 4 月 2 日にリリースされました。から入手可能です。
ミラーは https://www.openssh.com/ にリストされています。
OpenSSH は 100% 完全な SSH プロトコル 2.0 実装であり、
SFTP クライアントとサーバーのサポートが含まれます。
改めて、OpenSSH コミュニティの皆様に感謝いたします。
プロジェクトへの継続的なサポート、特に貢献してくれた人たち
コードまたはパッチ、報告されたバグ、テストされたスナップショット、またはへの寄付
プロジェクト。寄付の詳細については、以下をご覧ください。
https://www.openssh.com/donations.html
互換性がない可能性のある変更
--------------------------------
* ssh(1) 、 sshd(8) : 実装のバグ互換性を削除
それはキーの再生成をサポートしていません。そのような実装を試みると、
OpenSSH と相互運用するため、最終的に失敗するようになります。
トランスポートではキーの再作成が必要です。
* sshd(8) : このリリースより前は、空の証明書が含まれていました。
プリンシパル セクションは、任意のプリンシパルと一致するものとして扱われます。
(つまり、ワイルドカードとして)authorized_keys プリンシパル = "" 経由で使用される場合
オプション。これは意図的なものでしたが、驚くべき結果を生み出しました。
CA が誤って発行した場合、潜在的に危険な状況になります。
空のプリンシパルセクションを持つ証明書:
ご想像のとおり、役に立ちませんが、次のように認証するために使用できます。
authorized_keys を介して CA を信頼したユーザー。 [これに注意してください
条件が当てはまらなかった

sshd_config(5) 経由で信頼された CA へ
TrustedUserCAKeys オプション。]
このリリースでは、空のプリンシパル セクションは一致しないものとして扱われます。
任意のプリンシパル、およびワイルドカードの解釈も修正
証明書プリンシパル内の文字。今では彼らは一貫して
ホスト証明書に対して実装されており、ユーザーに対してはサポートされていません
証明書。
* ssh(1) : -J および同等の -oProxyJump="..." オプションが追加されました
渡された ProxyJump/-J オプションのユーザー名とホスト名を検証します
コマンドライン経由 (このような検証は実行されません)
設定ファイルのオプション)。これにより、シェルの注入が防止されます。
敵対者に直接さらされた状況
入力、それはテリーだったでしょう
[切り捨てられた]
OpenSSH 10.2 は 2025 年 10 月 10 日にリリースされました。から入手可能です。
ミラーは https://www.openssh.com/ にリストされています。
OpenSSH は 100% 完全な SSH プロトコル 2.0 実装であり、
SFTP クライアントとサーバーのサポートが含まれます。
改めて、OpenSSH コミュニティの皆様に感謝いたします。
プロジェクトへの継続的なサポート、特に貢献してくれた人たち
コードまたはパッチ、報告されたバグ、テストされたスナップショット、またはへの寄付
プロジェクト。寄付の詳細については、以下をご覧ください。
https://www.openssh.com/donations.html
将来の非推奨の警告
------------------------
* OpenSSH の将来のリリースでは、SHA1 SSHFP のサポートが廃止されます。
SHA1 ハッシュ関数の弱点によりレコードが失われる可能性があります。 SHA1 SSHFP
DNS レコードは無視され、ssh-keygen -r のみが生成されます。
SHA256 SSHFP レコード。
SHA256 ハッシュ アルゴリズムには既知の弱点はありませんが、
にリリースされた OpenSSH 6.1 以降、SSHFP レコードがサポートされています。
2012年。
OpenSSH 10.1 以降の変更点
==========================
これはバグ修正リリースであり、主にレンダリングされた問題を修正することを目的としています。
ControlPersist が有効な場合、ssh(1) は使用できません。
バグ修正
--------
*ssh(1)

: 時の端末接続の誤処理を修正
ControlPersist がアクティブになったため、セッションが使用できなくなりました。
bz3872
* ssh-keygen(1) : PKCS#11 トークンからのキーのダウンロードを修正しました。
* ssh-keygen(1) : CA キーが保持されている場合の CA 署名操作を修正
ssh-agent(1) 内。 bz3877
携帯性
-----------
* すべて: mmap(2) のないプラットフォームをサポートします。次のような WASM ビルド
https://hterm.org
* すべて: fnctl.h インクルードが欠落しているため、FreeBSD 上のビルドを修正。
* すべて: MacOS 10.12 Sierra 未満のビルドを修正します。
クロック_ゲットタイム(3)
* sshd(8) : リモートホストが「UNKNOWN」の場合は PAM_RHOST を行わないでください
プレースホルダー名。一部の PAM モジュールで発生する可能性のあるハングを回避します。
彼らはそれを解決しようとします。 sshd(8) は「UNKNOWN」のみを使用することに注意してください。
接続が IPv4 または IPv6 ソケット上にない場合の名前。
チェックサム:
==========
SHA1 ( openssh-10.2.tar.gz ) = 6fcda8004bad0fb0eaee60e8308f91b605ad0dce
SHA256 ( openssh-10.2.tar.gz ) = y0
[切り捨てられた]
OpenSSH 10.1 は 2025 年 10 月 6 日にリリースされました。から入手可能です。
ミラーは https://www.openssh.com/ にリストされています。
OpenSSH は 100% 完全な SSH プロトコル 2.0 実装であり、
SFTP クライアントとサーバーのサポートが含まれます。
改めて、OpenSSH コミュニティの皆様に感謝いたします。
プロジェクトへの継続的なサポート、特に貢献してくれた人たち
コードまたはパッチ、報告されたバグ、テストされたスナップショット、またはへの寄付
プロジェクト。寄付の詳細については、以下をご覧ください。
https://www.openssh.com/donations.html
将来の非推奨の警告
------------------------
* OpenSSH の将来のリリースでは、SHA1 SSHFP のサポートが廃止されます。
SHA1 ハッシュ関数の弱点によりレコードが失われる可能性があります。 SHA1 SSHFP
DNS レコードは無視され、ssh-keygen -r のみが生成されます。
SHA256 SSHFP レコード。
SHA256 ハッシュ アルゴリズムには既知の弱点はありませんが、
OpenSSH 6.1 以降、SSHFP レコードがサポートされています。

解放された
2012年。
互換性がない可能性のある変更
--------------------------------
* ssh(1) : 接続が非ポストをネゴシエートする場合に警告を追加します
量子鍵合意アルゴリズム。
この警告は、「今すぐ保存して復号化してください」というリスクがあるために追加されました。
後で」攻撃。詳細については、https://openssh.com/pq.html をご覧ください。
この警告は、新しい WarnWeakCrypto ssh_config によって制御される可能性があります。
オプション。デフォルトはオンです。このオプションはおそらく制御します
将来的には、弱い暗号に関する警告が追加される可能性があります。
* ssh(1) 、 sshd(8) : DSCP マーキング/IPQoS の処理に対する大幅な変更
クライアントとサーバーの両方でデフォルトの DSCP (別名 IPQoS) 値
が改訂され、実行時にこれらの値が使用される方法が変更されました。
変わりました。
インタラクティブ トラフィックが EF (優先) に割り当てられるようになりました。
Forwarding) クラスがデフォルトで設定されます。これにより、より適切な
中間ネットワークのパケット優先順位情報、
無線メディアなど (RFC 8325 を参照)。非インタラクティブなトラフィック
オペレーティング システムのデフォルトの DSCP マーキングが使用されるようになります。どちらも
インタラクティブおよび非インタラクティブ DSCP 値は、次の方法でオーバーライドできます。
IPQoS キーワード（「」で説明）
[切り捨てられた]
OpenSSH 10.0 は 2025 年 4 月 9 日にリリースされました。から入手可能です。
ミラーは https://www.openssh.com/ にリストされています。
OpenSSH は 100% 完全な SSH プロトコル 2.0 実装であり、
SFTP クライアントとサーバーのサポートが含まれます。
改めて、OpenSSH コミュニティの皆様に感謝いたします。
プロジェクトへの継続的なサポート、特に貢献してくれた人たち
コードまたはパッチ、報告されたバグ、テストされたスナップショット、またはへの寄付
プロジェクト。寄付の詳細については、以下をご覧ください。
https://www.openssh.com/donations.html
互換性がない可能性のある変更
--------------------------------
* このリリースでは、弱い DSA 署名のサポートが削除されます。
アルゴリズムを使用して、非推奨プロセスを完了します。

に始まった
2015 (DSA がデフォルトで無効になっていたとき) と繰り返し警告されました
過去 12 か月にわたって。
* scp(1) 、 sftp(1) : によって呼び出されたときに「ControlMaster no」を ssh に渡します。
scp と sftp。これにより、これらによる暗黙的なセッションの作成が無効になります。
ControlMaster が設定で Yes/auto に設定されている場合のツール、
一部のユーザーはこれを意外だと感じました。この変更によって妨げられることはありません
scp/sftp が既存の多重化セッションを使用しないようにする
すでに作成されています。 GHPR557
* このリリースにはバージョン番号 10.0 があり、それ自体がアナウンスされます
「SSH-2.0-OpenSSH_10.0」として。単純に一致するソフトウェア
「OpenSSH_1*」のようなパターンを使用するバージョンは、混乱を招く可能性があります。
これ。
* sshd(8) : このリリースでは、
プロトコルごとのユーザー認証フェーズ
sshd-session バイナリを新しい sshd-auth バイナリに接続します。
このコードを別のバイナリに分割すると、
重要な認証前の攻撃対象領域には、完全に
残りの部分で使用されるコードから切り離されたアドレス空間
接続。また、実行時メモリもわずかに節約されます。
認証コードは認証後にアンロードされます
フェーズが完了します。この変化はほとんど目に見えないはずです
ただし、一部のログ メッセージは表示されない場合があります。

[切り捨てられた]

## Original Extract

OpenSSH release notes

OpenSSH 10.5 / 10.5p1 (2026-08-11)
OpenSSH 10.5 was released on 2026-08-11. It is available from the
mirrors listed at https://www.openssh.com/ .
OpenSSH is a 100% complete SSH protocol 2.0 implementation and
includes sftp client and server support.
Recently the OpenSSH team have received a large number of security
bug reports, many of which are findings from AI models or made with
AI assistance. While many AI reports are determined not to have
security impact when considered in the context of a realistic
threat model, we very much welcome these reports, especially when
combined with human triage, analysis, test-cases and particularly
when accompanied by proposed fixes.
We have seen a number of cases where a security bug identified by
AI tools is subsequently independently discovered by a different
researcher. This suggests that adversaries who do not report bugs
to OSS projects are likely to be able to discover these bugs too.
Given this, the OpenSSH team will, for now, be making more frequent
releases to get bugfixes into users' hands more quickly rather than
batching them until the next planned release.
Once again, we would like to thank the OpenSSH community for their
continued support of the project, especially those who contributed
code or patches, reported bugs, tested snapshots or donated to the
project. More information on donations may be found at:
https://www.openssh.com/donations.html
Potentially-incompatible changes
--------------------------------
* Portable OpenSSH now requires ECC (Elliptic Curve Cryptography)
support in libcrypto, including support for the NISTP521 curve.
ECC is included in the default build configurations of all
versions of all libcrypto implementations currently supported by
OpenSSH, including LibreSSL, OpenSSL, BoringSSL and AWS LC.
The --without-openssl build configuration is not affected.
Changes since OpenSSH 10.4
==========================
This release contains a number of security fixes and small bugfixes.
Security
========
* ssh-agent(1) : fix an interaction
[truncated]
OpenSSH 10.4 was released on 2026-07-06. It is available from the
mirrors listed at https://www.openssh.com/ .
OpenSSH is a 100% complete SSH protocol 2.0 implementation and
includes sftp client and server support.
Once again, we would like to thank the OpenSSH community for their
continued support of the project, especially those who contributed
code or patches, reported bugs, tested snapshots or donated to the
project. More information on donations may be found at:
https://www.openssh.com/donations.html
Potentially-incompatible changes
--------------------------------
* sshd(8) : configuration dump mode ("sshd -G") now writes directives
in mixed case (e.g. "PubkeyAuthentication") whereas previously it
emitted only lower-case names.
* sshd(8) : on Linux systems with the seccomp sandbox enabled,
failures to enable SECCOMP or NO_NEW_PRIVS are now fatal.
Previously sshd(8) would log the error but continue operation,
to support systems that lacked these features. Now systems that
lack these should instead disable the sandbox at configure time.
* ssh(1) , sshd(8) : make the transport protocol stricter by
disconnecting if the peer sends non-KEX messages during a post-
authentication key re-exchange. Previously a malicious peer could
continue sending non-key exchange messages without penalty. These
would be buffered, causing memory to be wasted up until the
connection terminated or the server/client hit a memory limit.
Implementations that do not restrict messages sent during key
exchange as per RFC4253 section 7.1 may be disconnected.
Reported by Marko Jevtic.
Changes since OpenSSH 10.3
==========================
This release contains a number of security fixes as well as general
bugfixes and a couple of new features.
Security
========
* sftp(1) : when downloading files on the command-line using
"sftp host:/path .", a malicious server could cause the file to
be downloaded to an unexpected location. This issue was identified
by the Swival Security Scanner.
* scp(1) : when
[truncated]
OpenSSH 10.3 was released on 2026-04-02. It is available from the
mirrors listed at https://www.openssh.com/ .
OpenSSH is a 100% complete SSH protocol 2.0 implementation and
includes sftp client and server support.
Once again, we would like to thank the OpenSSH community for their
continued support of the project, especially those who contributed
code or patches, reported bugs, tested snapshots or donated to the
project. More information on donations may be found at:
https://www.openssh.com/donations.html
Potentially-incompatible changes
--------------------------------
* ssh(1) , sshd(8) : remove bug compatibility for implementations
that don't support rekeying. If such an implementation tries to
interoperate with OpenSSH, it will now eventually fail when the
transport needs rekeying.
* sshd(8) : prior to this release, a certificate that had an empty
principals section would be treated as matching any principal
(i.e. as a wildcard) when used via authorized_keys principals=""
option. This was intentional, but created a surprising and
potentially risky situation if a CA accidentally issued a
certificate with an empty principals section: instead of being
useless as one might expect, it could be used to authenticate as
any user who trusted the CA via authorized_keys. [Note that this
condition did not apply to CAs trusted via the sshd_config(5)
TrustedUserCAKeys option.]
This release treats an empty principals section as never matching
any principal, and also fixes interpretation of wildcard
characters in certificate principals. Now they are consistently
implemented for host certificates and not supported for user
certificates.
* ssh(1) : the -J and equivalent -oProxyJump="..." options now
validate user and host names for ProxyJump/-J options passed
via the command-line (no such validation is performed for this
option in configuration files). This prevents shell injection in
situations where these were directly exposed to adversarial
input, which would have been a terri
[truncated]
OpenSSH 10.2 was released on 2025-10-10. It is available from the
mirrors listed at https://www.openssh.com/ .
OpenSSH is a 100% complete SSH protocol 2.0 implementation and
includes sftp client and server support.
Once again, we would like to thank the OpenSSH community for their
continued support of the project, especially those who contributed
code or patches, reported bugs, tested snapshots or donated to the
project. More information on donations may be found at:
https://www.openssh.com/donations.html
Future deprecation warning
--------------------------
* A future release of OpenSSH will deprecate support for SHA1 SSHFP
records due to weaknesses in the SHA1 hash function. SHA1 SSHFP
DNS records will be ignored and ssh-keygen -r will generate only
SHA256 SSHFP records.
The SHA256 hash algorithm, which has no known weaknesses, has
been supported for SSHFP records since OpenSSH 6.1, released in
2012.
Changes since OpenSSH 10.1
==========================
This is a bugfix release, primarily to fix a problem that rendered
ssh(1) unusable when ControlPersist was enabled.
Bugfixes
--------
* ssh(1) : fix mishandling of terminal connections when
ControlPersist was active that rendered the session unusable.
bz3872
* ssh-keygen(1) : fix download of keys from PKCS#11 tokens.
* ssh-keygen(1) : fix CA signing operations when the CA key is held
in a ssh-agent(1) . bz3877
Portability
-----------
* All: support platforms without mmap(2), e.g. WASM builds such as
https://hterm.org
* All: fix builds on FreeBSD for missing fnctl.h include.
* All: fix builds on MacOS <10.12 Sierra, which lacks
clock_gettime(3)
* sshd(8) : don't PAM_RHOST if the remote host is the "UNKNOWN"
placeholder name. Avoids potential hangs in some PAM modules as
they try to resolve it. Note, sshd(8) only uses the "UNKNOWN"
name when the connection is not on an IPv4 or IPv6 socket.
Checksums:
==========
SHA1 ( openssh-10.2.tar.gz ) = 6fcda8004bad0fb0eaee60e8308f91b605ad0dce
SHA256 ( openssh-10.2.tar.gz ) = y0
[truncated]
OpenSSH 10.1 was released on 2025-10-06. It is available from the
mirrors listed at https://www.openssh.com/ .
OpenSSH is a 100% complete SSH protocol 2.0 implementation and
includes sftp client and server support.
Once again, we would like to thank the OpenSSH community for their
continued support of the project, especially those who contributed
code or patches, reported bugs, tested snapshots or donated to the
project. More information on donations may be found at:
https://www.openssh.com/donations.html
Future deprecation warning
--------------------------
* A future release of OpenSSH will deprecate support for SHA1 SSHFP
records due to weaknesses in the SHA1 hash function. SHA1 SSHFP
DNS records will be ignored and ssh-keygen -r will generate only
SHA256 SSHFP records.
The SHA256 hash algorithm, which has no known weaknesses, has
been supported for SSHFP records since OpenSSH 6.1, released in
2012.
Potentially-incompatible changes
--------------------------------
* ssh(1) : add a warning when the connection negotiates a non-post
quantum key agreement algorithm.
This warning has been added due to the risk of "store now, decrypt
later" attacks. More details at https://openssh.com/pq.html
This warning may be controlled via a new WarnWeakCrypto ssh_config
option, defaulting to on. This option is likely to control
additional weak crypto warnings in the future.
* ssh(1) , sshd(8) : major changes to handling of DSCP marking/IPQoS
In both client and server the default DSCP (a.k.a IPQoS) values
were revised and the way these values are used during runtime has
changed.
Interactive traffic is now assigned to the EF (Expedited
Forwarding) class by default. This provides more appropriate
packet prioritisation information for the intermediate network,
such as wireless media (cf. RFC 8325 ). Non-interactive traffic
will now use the operating system default DSCP marking. Both the
interactive and non-interactive DSCP values may be overridden via
the IPQoS keyword, described in s
[truncated]
OpenSSH 10.0 was released on 2025-04-09. It is available from the
mirrors listed at https://www.openssh.com/ .
OpenSSH is a 100% complete SSH protocol 2.0 implementation and
includes sftp client and server support.
Once again, we would like to thank the OpenSSH community for their
continued support of the project, especially those who contributed
code or patches, reported bugs, tested snapshots or donated to the
project. More information on donations may be found at:
https://www.openssh.com/donations.html
Potentially-incompatible changes
--------------------------------
* This release removes support for the weak DSA signature
algorithm, completing the deprecation process that began in
2015 (when DSA was disabled by default) and repeatedly warned
over the last 12 months.
* scp(1) , sftp(1) : pass "ControlMaster no" to ssh when invoked by
scp & sftp. This disables implicit session creation by these
tools when ControlMaster was set to yes/auto by configuration,
which some users found surprising. This change will not prevent
scp/sftp from using an existing multiplexing session if one had
already been created. GHPR557
* This release has the version number 10.0 and announces itself
as "SSH-2.0-OpenSSH_10.0". Software that naively matches
versions using patterns like "OpenSSH_1*" may be confused by
this.
* sshd(8) : this release removes the code responsible for the
user authentication phase of the protocol from the per-
connection sshd-session binary to a new sshd-auth binary.
Splitting this code into a separate binary ensures that the
crucial pre-authentication attack surface has an entirely
disjoint address space from the code used for the rest of the
connection. It also yields a small runtime memory saving as the
authentication code will be unloaded after the authentication
phase completes. This change should be largely invisible to
users, though some log messages may no

[truncated]
