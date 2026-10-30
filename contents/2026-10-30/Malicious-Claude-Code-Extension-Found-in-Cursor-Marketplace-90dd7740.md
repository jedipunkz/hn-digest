---
source: "https://cursor.com/blog/claude-code-extension-security-notice"
hn_url: "https://news.ycombinator.com/item?id=50301004"
title: "Malicious Claude Code Extension Found in Cursor Marketplace"
article_title: "Security notice — malicious version of Anthropic.claude-code (2.1.313) identified in the Cursor Extension Marketplace · Cursor"
image: "https://ptht05hbb1ssoooe.public.blob.vercel-storage.com/assets/og/opengraph-default.png"
author: "meetpateltech"
captured_at: "2026-10-30T08:09:11Z"
capture_tool: "hn-digest"
hn_id: 50301004
score: 1
comments: 0
posted_at: "2026-10-30T08:04:26Z"
tags:
  - hacker-news
---

# Malicious Claude Code Extension Found in Cursor Marketplace

- HN: [50301004](https://news.ycombinator.com/item?id=50301004)
- Source: [cursor.com](https://cursor.com/blog/claude-code-extension-security-notice)
- Score: 1
- Comments: 0
- Posted: 2026-10-30T08:04:26Z

## Translation

Title: Malicious Claude Code Extension Found in Cursor Marketplace
Article title: Security notice — malicious version of Anthropic.claude-code (2.1.313) identified in the Cursor Extension Marketplace · Cursor
Description: Today we identified a security incident involving a specific extension package in the Cursor VS Code extension Marketplace: Anthropic.claude-code version 2.1.313.

Article text:
Security notice — malicious version of Anthropic.claude-code (2.1.313) identified in the Cursor Extension Marketplace · Cursor Skip to content Cursor Models ↓ Grok
Security notice — malicious version of Anthropic.claude-code (2.1.313) identified in the Cursor Extension Marketplace
Check whether you are impacted
Today we identified a security incident involving a specific extension package in the Cursor VS Code extension Marketplace: Anthropic.claude-code version 2.1.313 . This malicious version of the extension was not developed by Anthropic. At most 5,000 users were affected, and we will reach out directly.
Between about 2:27am PT and 1:35pm PT today, a maliciously modified version of Anthropic’s Claude Code extension was made available in the Cursor VS Code extension Marketplace. The extension runs malicious code that attempts to install a persistent SSH backdoor. Specifically, it:
Collects host and user information: hostname, username, OS, architecture, home directory path, environment variable names, working directories, configured Git email, and public IP.
Attempts to add an attacker-controlled SSH key to authorized_keys , enable the local SSH service, and open a reverse SSH tunnel. Outside of highly unusual circumstances where Cursor was started with administrator privileges, due to errors in the malicious code, this step would fail.
Attempts to make the reverse SSH tunnel persistent via a LaunchAgent (macOS) or a scheduled task that attempts to run as SYSTEM with a Startup-script fallback (Windows). Background scripts retry the tunnel roughly every 60 seconds.
Reports host details and backdoor status via HTTP POST to /beacon on port 8099.
Note: The attacker server no longer functions.
We worked with the cloud provider that was hosting it to take it offline, so the reverse shell and host detail POST are not functional.
These findings concern the identified package. We have no evidence that other versions of the extension or the official distribution are affected.
Additionally, turning on the local SSH service requires admin rights, which the malicious extension does not have on most macOS and Windows hosts. On those hosts the step fails and the attacker could not have logged in through the tunnel. The extension still adds the attacker’s SSH key, installs persistence, and sends host details to the attacker. Follow the remediation steps below regardless.
Removed the malicious version from the marketplace immediately upon detection and blocked further distribution. We followed with a full shut down of the VSCode Marketplace out of an abundance of caution.
We audited access and confirmed no other extensions were tampered with.
We rotated our keys, rebuilt the marketplace, and brought it back online
We worked with the hosting provider to take the malicious server offline, disrupting the attack by blocking the connection to that server.
If you installed or updated to Anthropic.claude-code 2.1.313 from Cursor VS Code extension marketplace between 2:27am PT and 1:35pm PT on October 29, 2026:
Run the check below to validate if you are impacted.
Search across endpoint, network, and extension-installation telemetry for the indicators below.
Isolate hosts showing evidence of payload execution or backdoor installation.
Remove the unauthorized SSH key, persistence mechanisms, and running tunnel processes using the remediation scripts below, and assess whether the host should be rebuilt. Uninstalling the extension may not be sufficient — the payload drops standalone persistence scripts that keep running.
Rotate credentials and sessions accessible from affected hosts, from a trusted system. The limited env-var collection does not rule out later credential access over the backdoor.
# Check whether you are impacted
The malicious extension appends the attacker's public key, labeled svc-telemetry , to your SSH authorized keys. To check for it:
grep -qs svc-telemetry ~/.ssh/authorized_keys && echo "svc-telemetry key found" || echo "No key found; machine not impacted"
On Windows, run in PowerShell:
$files = " $HOME \.ssh\authorized_keys" , "C:\ProgramData\ssh\administrators_authorized_keys" | Where-Object { Test-Path $_ }
if ($files -and ( Select-String svc - telemetry $files)) { "svc-telemetry key found" } else { "No key found; machine not impacted" }
If the key is found, your machine is impacted. Follow steps 3 to 5 above, and run the remediation script for your platform below.
These scripts remove the malicious extension, the persistence mechanisms, the running tunnel processes, and the attacker's SSH key. They exit with a non-zero status if anything remains.
On macOS, download remediate-macos.sh and run it as the affected user, not as root, from the folder you saved it to:
bash remediate-macos.sh
On Windows, download remediate-windows.ps1 and run it from PowerShell opened with Run as Administrator , in the folder you saved it to:
powershell - ExecutionPolicy Bypass - File .\remediate - windows.ps1
# Indicators of compromise
Modified VSIX SHA-256 (darwin-arm64): 432ccdee3801f99e52b8fe31f69ea9870118e96ea7ef9ecabed1479621b5ab62 (MD5 da5469c4221775e9fe2e413e59ca7fd9)
Modified VSIX SHA-256 (win32-x64): 3de1722901165769893bef57c72594e9a54459faa23f3091b136b79966edd081 (MD5 bd3b71ddf3d0cc3af76249afd2c0eda1)
Modified code: extension.js in the anthropic.claude-code-2.1.313-darwin-arm64 (macOS) or anthropic.claude-code-2.1.313-win32-x64 (Windows) extension folder contains 134[.]209[.]82[.]225 and svc-telemetry
Attacker endpoint: 134[.]209[.]82[.]225
Reverse SSH: tunnel@134[.]209[.]82[.]225, TCP 22
External beacon: hxxp://134[.]209[.]82[.]225:8099/beacon (HTTP POST)
Local beacon attempt: 127[.]0[.]0[.]1:8099/beacon (HTTP POST, loopback)
Public IP lookups to ifconfig[.]me / api[.]ipify[.]org — legitimate services; access alone is not an indicator
Unauthorized public key: ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDdvylCEJ/k2TpezcTfedrkHyFSK9IZ3vmP6vhY4Ss0P svc-telemetry
Unauthorized key fingerprint: SHA256:0Z6R+DF64Do2nwQic8a7kkfDLNNk4IiK+2U2Rci4htk
Reverse-tunnel key fingerprint: SHA256:KOAU57+ZG6u3MG4WwzPzPX2O5YEe9qdOC1Bl+6MVyDQ
macOS: unauthorized key in ~/.ssh/authorized_keys ; ~/.config/.smt , ~/.config/.smtk ; $TMPDIR/.stk_<10 hex> ; ~/Library/LaunchAgents/com.apple.sysmond.plist (label com.apple.sysmond )
Windows: %USERPROFILE%\.ssh\authorized_keys , %USERPROFILE%\.config\smt.cmd , %USERPROFILE%\.config\smtk , %TEMP%\.stk_<10 hex> ; scheduled task SysMon (attempts to run as SYSTEM); fallback %APPDATA%\Microsoft\Windows\Start Menu\Programs\Startup\sysmon.cmd
ssh -N -R '*:0:localhost:22' , autossh -M 0
SSH options: StrictHostKeyChecking=no , UserKnownHostsFile=/dev/null , ServerAliveInterval=30 , ServerAliveCountMax=3 , ExitOnForwardFailure=yes , IdentitiesOnly=yes
systemsetup -setremotelogin on
Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0
If you have further questions, reach out to security@cursor.com .
Cursor earns AIUC-1 certification for agent security and reliability
CFOs and the new economics of AI
Security notice — malicious version of Anthropic.claude-code (2.1.313) identified in the Cursor Extension Marketplace
Check whether you are impacted
Today we identified a security incident involving a specific extension package in the Cursor VS Code extension Marketplace: Anthropic.claude-code version 2.1.313 . This malicious version of the extension was not developed by Anthropic. At most 5,000 users were affected, and we will reach out directly.
Between about 2:27am PT and 1:35pm PT today, a maliciously modified version of Anthropic’s Claude Code extension was made available in the Cursor VS Code extension Marketplace. The extension runs malicious code that attempts to install a persistent SSH backdoor. Specifically, it:
Collects host and user information: hostname, username, OS, architecture, home directory path, environment variable names, working directories, configured Git email, and public IP.
Attempts to add an attacker-controlled SSH key to authorized_keys , enable the local SSH service, and open a reverse SSH tunnel. Outside of highly unusual circumstances where Cursor was started with administrator privileges, due to errors in the malicious code, this step would fail.
Attempts to make the reverse SSH tunnel persistent via a LaunchAgent (macOS) or a scheduled task that attempts to run as SYSTEM with a Startup-script fallback (Windows). Background scripts retry the tunnel roughly every 60 seconds.
Reports host details and backdoor status via HTTP POST to /beacon on port 8099.
Note: The attacker server no longer functions.
We worked with the cloud provider that was hosting it to take it offline, so the reverse shell and host detail POST are not functional.
These findings concern the identified package. We have no evidence that other versions of the extension or the official distribution are affected.
Additionally, turning on the local SSH service requires admin rights, which the malicious extension does not have on most macOS and Windows hosts. On those hosts the step fails and the attacker could not have logged in through the tunnel. The extension still adds the attacker’s SSH key, installs persistence, and sends host details to the attacker. Follow the remediation steps below regardless.
Removed the malicious version from the marketplace immediately upon detection and blocked further distribution. We followed with a full shut down of the VSCode Marketplace out of an abundance of caution.
We audited access and confirmed no other extensions were tampered with.
We rotated our keys, rebuilt the marketplace, and brought it back online
We worked with the hosting provider to take the malicious server offline, disrupting the attack by blocking the connection to that server.
If you installed or updated to Anthropic.claude-code 2.1.313 from Cursor VS Code extension marketplace between 2:27am PT and 1:35pm PT on October 29, 2026:
Run the check below to validate if you are impacted.
Search across endpoint, network, and extension-installation telemetry for the indicators below.
Isolate hosts showing evidence of payload execution or backdoor installation.
Remove the unauthorized SSH key, persistence mechanisms, and running tunnel processes using the remediation scripts below, and assess whether the host should be rebuilt. Uninstalling the extension may not be sufficient — the payload drops standalone persistence scripts that keep running.
Rotate credentials and sessions accessible from affected hosts, from a trusted system. The limited env-var collection does not rule out later credential access over the backdoor.
# Check whether you are impacted
The malicious extension appends the attacker's public key, labeled svc-telemetry , to your SSH authorized keys. To check for it:
grep -qs svc-telemetry ~/.ssh/authorized_keys && echo "svc-telemetry key found" || echo "No key found; machine not impacted"
On Windows, run in PowerShell:
$files = " $HOME \.ssh\authorized_keys" , "C:\ProgramData\ssh\administrators_authorized_keys" | Where-Object { Test-Path $_ }
if ($files -and ( Select-String svc - telemetry $files)) { "svc-telemetry key found" } else { "No key found; machine not impacted" }
If the key is found, your machine is impacted. Follow steps 3 to 5 above, and run the remediation script for your platform below.
These scripts remove the malicious extension, the persistence mechanisms, the running tunnel processes, and the attacker's SSH key. They exit with a non-zero status if anything remains.
On macOS, download remediate-macos.sh and run it as the affected user, not as root, from the folder you saved it to:
bash remediate-macos.sh
On Windows, download remediate-windows.ps1 and run it from PowerShell opened with Run as Administrator , in the folder you saved it to:
powershell - ExecutionPolicy Bypass - File .\remediate - windows.ps1
# Indicators of compromise
Modified VSIX SHA-256 (darwin-arm64): 432ccdee3801f99e52b8fe31f69ea9870118e96ea7ef9ecabed1479621b5ab62 (MD5 da5469c4221775e9fe2e4

[truncated]

## Original Extract

Today we identified a security incident involving a specific extension package in the Cursor VS Code extension Marketplace: Anthropic.claude-code version 2.1.313.

Security notice — malicious version of Anthropic.claude-code (2.1.313) identified in the Cursor Extension Marketplace · Cursor Skip to content Cursor Models ↓ Grok
Security notice — malicious version of Anthropic.claude-code (2.1.313) identified in the Cursor Extension Marketplace
Check whether you are impacted
Today we identified a security incident involving a specific extension package in the Cursor VS Code extension Marketplace: Anthropic.claude-code version 2.1.313 . This malicious version of the extension was not developed by Anthropic. At most 5,000 users were affected, and we will reach out directly.
Between about 2:27am PT and 1:35pm PT today, a maliciously modified version of Anthropic’s Claude Code extension was made available in the Cursor VS Code extension Marketplace. The extension runs malicious code that attempts to install a persistent SSH backdoor. Specifically, it:
Collects host and user information: hostname, username, OS, architecture, home directory path, environment variable names, working directories, configured Git email, and public IP.
Attempts to add an attacker-controlled SSH key to authorized_keys , enable the local SSH service, and open a reverse SSH tunnel. Outside of highly unusual circumstances where Cursor was started with administrator privileges, due to errors in the malicious code, this step would fail.
Attempts to make the reverse SSH tunnel persistent via a LaunchAgent (macOS) or a scheduled task that attempts to run as SYSTEM with a Startup-script fallback (Windows). Background scripts retry the tunnel roughly every 60 seconds.
Reports host details and backdoor status via HTTP POST to /beacon on port 8099.
Note: The attacker server no longer functions.
We worked with the cloud provider that was hosting it to take it offline, so the reverse shell and host detail POST are not functional.
These findings concern the identified package. We have no evidence that other versions of the extension or the official distribution are affected.
Additionally, turning on the local SSH service requires admin rights, which the malicious extension does not have on most macOS and Windows hosts. On those hosts the step fails and the attacker could not have logged in through the tunnel. The extension still adds the attacker’s SSH key, installs persistence, and sends host details to the attacker. Follow the remediation steps below regardless.
Removed the malicious version from the marketplace immediately upon detection and blocked further distribution. We followed with a full shut down of the VSCode Marketplace out of an abundance of caution.
We audited access and confirmed no other extensions were tampered with.
We rotated our keys, rebuilt the marketplace, and brought it back online
We worked with the hosting provider to take the malicious server offline, disrupting the attack by blocking the connection to that server.
If you installed or updated to Anthropic.claude-code 2.1.313 from Cursor VS Code extension marketplace between 2:27am PT and 1:35pm PT on October 29, 2026:
Run the check below to validate if you are impacted.
Search across endpoint, network, and extension-installation telemetry for the indicators below.
Isolate hosts showing evidence of payload execution or backdoor installation.
Remove the unauthorized SSH key, persistence mechanisms, and running tunnel processes using the remediation scripts below, and assess whether the host should be rebuilt. Uninstalling the extension may not be sufficient — the payload drops standalone persistence scripts that keep running.
Rotate credentials and sessions accessible from affected hosts, from a trusted system. The limited env-var collection does not rule out later credential access over the backdoor.
# Check whether you are impacted
The malicious extension appends the attacker's public key, labeled svc-telemetry , to your SSH authorized keys. To check for it:
grep -qs svc-telemetry ~/.ssh/authorized_keys && echo "svc-telemetry key found" || echo "No key found; machine not impacted"
On Windows, run in PowerShell:
$files = " $HOME \.ssh\authorized_keys" , "C:\ProgramData\ssh\administrators_authorized_keys" | Where-Object { Test-Path $_ }
if ($files -and ( Select-String svc - telemetry $files)) { "svc-telemetry key found" } else { "No key found; machine not impacted" }
If the key is found, your machine is impacted. Follow steps 3 to 5 above, and run the remediation script for your platform below.
These scripts remove the malicious extension, the persistence mechanisms, the running tunnel processes, and the attacker's SSH key. They exit with a non-zero status if anything remains.
On macOS, download remediate-macos.sh and run it as the affected user, not as root, from the folder you saved it to:
bash remediate-macos.sh
On Windows, download remediate-windows.ps1 and run it from PowerShell opened with Run as Administrator , in the folder you saved it to:
powershell - ExecutionPolicy Bypass - File .\remediate - windows.ps1
# Indicators of compromise
Modified VSIX SHA-256 (darwin-arm64): 432ccdee3801f99e52b8fe31f69ea9870118e96ea7ef9ecabed1479621b5ab62 (MD5 da5469c4221775e9fe2e413e59ca7fd9)
Modified VSIX SHA-256 (win32-x64): 3de1722901165769893bef57c72594e9a54459faa23f3091b136b79966edd081 (MD5 bd3b71ddf3d0cc3af76249afd2c0eda1)
Modified code: extension.js in the anthropic.claude-code-2.1.313-darwin-arm64 (macOS) or anthropic.claude-code-2.1.313-win32-x64 (Windows) extension folder contains 134[.]209[.]82[.]225 and svc-telemetry
Attacker endpoint: 134[.]209[.]82[.]225
Reverse SSH: tunnel@134[.]209[.]82[.]225, TCP 22
External beacon: hxxp://134[.]209[.]82[.]225:8099/beacon (HTTP POST)
Local beacon attempt: 127[.]0[.]0[.]1:8099/beacon (HTTP POST, loopback)
Public IP lookups to ifconfig[.]me / api[.]ipify[.]org — legitimate services; access alone is not an indicator
Unauthorized public key: ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDdvylCEJ/k2TpezcTfedrkHyFSK9IZ3vmP6vhY4Ss0P svc-telemetry
Unauthorized key fingerprint: SHA256:0Z6R+DF64Do2nwQic8a7kkfDLNNk4IiK+2U2Rci4htk
Reverse-tunnel key fingerprint: SHA256:KOAU57+ZG6u3MG4WwzPzPX2O5YEe9qdOC1Bl+6MVyDQ
macOS: unauthorized key in ~/.ssh/authorized_keys ; ~/.config/.smt , ~/.config/.smtk ; $TMPDIR/.stk_<10 hex> ; ~/Library/LaunchAgents/com.apple.sysmond.plist (label com.apple.sysmond )
Windows: %USERPROFILE%\.ssh\authorized_keys , %USERPROFILE%\.config\smt.cmd , %USERPROFILE%\.config\smtk , %TEMP%\.stk_<10 hex> ; scheduled task SysMon (attempts to run as SYSTEM); fallback %APPDATA%\Microsoft\Windows\Start Menu\Programs\Startup\sysmon.cmd
ssh -N -R '*:0:localhost:22' , autossh -M 0
SSH options: StrictHostKeyChecking=no , UserKnownHostsFile=/dev/null , ServerAliveInterval=30 , ServerAliveCountMax=3 , ExitOnForwardFailure=yes , IdentitiesOnly=yes
systemsetup -setremotelogin on
Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0
If you have further questions, reach out to security@cursor.com .
Cursor earns AIUC-1 certification for agent security and reliability
CFOs and the new economics of AI
Security notice — malicious version of Anthropic.claude-code (2.1.313) identified in the Cursor Extension Marketplace
Check whether you are impacted
Today we identified a security incident involving a specific extension package in the Cursor VS Code extension Marketplace: Anthropic.claude-code version 2.1.313 . This malicious version of the extension was not developed by Anthropic. At most 5,000 users were affected, and we will reach out directly.
Between about 2:27am PT and 1:35pm PT today, a maliciously modified version of Anthropic’s Claude Code extension was made available in the Cursor VS Code extension Marketplace. The extension runs malicious code that attempts to install a persistent SSH backdoor. Specifically, it:
Collects host and user information: hostname, username, OS, architecture, home directory path, environment variable names, working directories, configured Git email, and public IP.
Attempts to add an attacker-controlled SSH key to authorized_keys , enable the local SSH service, and open a reverse SSH tunnel. Outside of highly unusual circumstances where Cursor was started with administrator privileges, due to errors in the malicious code, this step would fail.
Attempts to make the reverse SSH tunnel persistent via a LaunchAgent (macOS) or a scheduled task that attempts to run as SYSTEM with a Startup-script fallback (Windows). Background scripts retry the tunnel roughly every 60 seconds.
Reports host details and backdoor status via HTTP POST to /beacon on port 8099.
Note: The attacker server no longer functions.
We worked with the cloud provider that was hosting it to take it offline, so the reverse shell and host detail POST are not functional.
These findings concern the identified package. We have no evidence that other versions of the extension or the official distribution are affected.
Additionally, turning on the local SSH service requires admin rights, which the malicious extension does not have on most macOS and Windows hosts. On those hosts the step fails and the attacker could not have logged in through the tunnel. The extension still adds the attacker’s SSH key, installs persistence, and sends host details to the attacker. Follow the remediation steps below regardless.
Removed the malicious version from the marketplace immediately upon detection and blocked further distribution. We followed with a full shut down of the VSCode Marketplace out of an abundance of caution.
We audited access and confirmed no other extensions were tampered with.
We rotated our keys, rebuilt the marketplace, and brought it back online
We worked with the hosting provider to take the malicious server offline, disrupting the attack by blocking the connection to that server.
If you installed or updated to Anthropic.claude-code 2.1.313 from Cursor VS Code extension marketplace between 2:27am PT and 1:35pm PT on October 29, 2026:
Run the check below to validate if you are impacted.
Search across endpoint, network, and extension-installation telemetry for the indicators below.
Isolate hosts showing evidence of payload execution or backdoor installation.
Remove the unauthorized SSH key, persistence mechanisms, and running tunnel processes using the remediation scripts below, and assess whether the host should be rebuilt. Uninstalling the extension may not be sufficient — the payload drops standalone persistence scripts that keep running.
Rotate credentials and sessions accessible from affected hosts, from a trusted system. The limited env-var collection does not rule out later credential access over the backdoor.
# Check whether you are impacted
The malicious extension appends the attacker's public key, labeled svc-telemetry , to your SSH authorized keys. To check for it:
grep -qs svc-telemetry ~/.ssh/authorized_keys && echo "svc-telemetry key found" || echo "No key found; machine not impacted"
On Windows, run in PowerShell:
$files = " $HOME \.ssh\authorized_keys" , "C:\ProgramData\ssh\administrators_authorized_keys" | Where-Object { Test-Path $_ }
if ($files -and ( Select-String svc - telemetry $files)) { "svc-telemetry key found" } else { "No key found; machine not impacted" }
If the key is found, your machine is impacted. Follow steps 3 to 5 above, and run the remediation script for your platform below.
These scripts remove the malicious extension, the persistence mechanisms, the running tunnel processes, and the attacker's SSH key. They exit with a non-zero status if anything remains.
On macOS, download remediate-macos.sh and run it as the affected user, not as root, from the folder you saved it to:
bash remediate-macos.sh
On Windows, download remediate-windows.ps1 and run it from PowerShell opened with Run as Administrator , in the folder you saved it to:
powershell - ExecutionPolicy Bypass - File .\remediate - windows.ps1
# Indicators of compromise
Modified VSIX SHA-256 (darwin-arm64): 432ccdee3801f99e52b8fe31f69ea9870118e96ea7ef9ecabed1479621b5ab62 (MD5 da5469c4221775e9fe2e4

[truncated]
