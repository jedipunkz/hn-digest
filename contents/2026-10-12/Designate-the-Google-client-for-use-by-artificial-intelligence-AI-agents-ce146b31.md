---
source: "https://support.google.com/cloud/answer/15549257?hl=en"
hn_url: "https://news.ycombinator.com/item?id=50061359"
title: "Designate the Google client for use by artificial intelligence (AI) agents"
article_title: "Manage OAuth Clients - Google Cloud Platform Console Help"
image: ""
author: "mooreds"
captured_at: "2026-10-12T20:08:46Z"
capture_tool: "hn-digest"
hn_id: 50061359
score: 1
comments: 0
posted_at: "2026-10-12T19:12:34Z"
tags:
  - hacker-news
---

# Designate the Google client for use by artificial intelligence (AI) agents

- HN: [50061359](https://news.ycombinator.com/item?id=50061359)
- Source: [support.google.com](https://support.google.com/cloud/answer/15549257?hl=en)
- Score: 1
- Comments: 0
- Posted: 2026-10-12T19:12:34Z

## Translation

Title: Designate the Google client for use by artificial intelligence (AI) agents
Article title: Manage OAuth Clients - Google Cloud Platform Console Help
Description: Your OAuth client is the credential which your application uses when making calls to Google OAuth 2.0 endpoint to receive an access token or ID token. After creating your OAuth client, you will receiv

Article text:
Skip to main content Google Cloud Platform Console Help Sign in Google Help Help Center
Your OAuth client is the credential which your application uses when making calls to Google OAuth 2.0 endpoint to receive an access token or ID token. After creating your OAuth client, you will receive a client ID and sometimes, a client secret.
Think of your client ID like your app's unique username when it needs to request an access token or ID token from Google's OAuth 2.0 endpoint. This ID helps Google identify your app and ensure that only authorized applications can access user data.
Similar to how you would use a username and password to log to online services, many applications use a client ID paired with a client secret. The client secret adds an extra layer of security, acting like your app's password.
Applications are categorized as either public or private clients:
Private Clients: These apps, like web server applications, can securely store the client secret because they run on servers you control.
Public Clients: Native apps or JavaScript-based apps fall under this category. They cannot securely store secrets, as they reside on user devices and as such do not use client secrets.
To create an OAuth 2.0 client ID in the console:
Navigate to the Google Auth Platform Clients page.
You will be prompted to create a project if you do not have one selected.
You will be prompted to register your application to use Google Auth if you are yet to do so. This is required before creating a client.
Select the appropriate application type for your application and enter any additional information required. Application types are described in more detail in the following sections.
Fill out the required information for the select client type and click the CREATE button to create the client.
Note : Your application's client secret will only be shown after you create the client. Store this information in a secure place such as Google Cloud Secret Manager because it will not be visible or accessible again. Learn more .
A web application is accessed by web browsers over a network.
Applications that use client-side JavaScript to access Google APIs must specify authorized JavaScript origins. The origins identify the domains from which your application can send API requests.
Specified origins must adhere to the following rules :
JavaScript origins must use the HTTPS scheme, not plain HTTP. Localhost URIs (including localhost IP address URIs) are exempt from this rule.
Hosts cannot be raw IP addresses. Localhost IP addresses are exempted from this rule.
If you use a port other than 80, you must specify it. For example: https://example.com:8080
Host TLDs (Top Level Domains) must belong to the public suffix list .
Host domains cannot be “googleusercontent.com”.
JavaScript origins cannot contain URL shortener domains (e.g. goo.gl) unless the app owns the domain.
JavaScript origins cannot contain the userinfo subcomponent.
JavaScript origins cannot contain the path component.
JavaScript origins cannot contain the query component.
JavaScript origins cannot contain the fragment component.
JavaScript origins cannot contain certain characters including:
Wildcard characters ('*')
Non-printable ASCII characters
Invalid percent encodings (any percent encoding that does not follow URL-encoding form of a percent sign followed by two hexadecimal digits)
Null characters (an encoded NULL character, e.g., %00, %C0%80 )
If you send a request to a Google OAuth 2.0 endpoint from an unregistered JavaScript origin, you will receive an origin_mismatch error.
Applications that access Google APIs from a server (often using languages and frameworks like Node.js, Java, .NET, and Python) must specify authorized redirect URIs. The redirect URIs are the endpoints of your application server to which the OAuth 2.0 server can send responses. Users are redirected to this path after they have authenticated with Google.
Redirect URIs must adhere to the following rules :
Redirect URIs must use the HTTPS scheme, not plain HTTP. Localhost URIs (including localhost IP address URIs) are exempt from this rule.
Hosts cannot be raw IP addresses. Localhost IP addresses are exempted from this rule.
Host TLDs (Top Level Domains) must belong to the public suffix list .
Redirect URIs cannot contain URL shortener domains (e.g. goo.gl ) unless the app owns the domain. Furthermore, if an app that owns a shortener domain chooses to redirect to that domain, that redirect URI must either contain “ /google-callback/ ” in its path or end with “ /google-callback ”.
Redirect URIs cannot contain the userinfo subcomponent.
Redirect URIs cannot contain a path traversal (also called directory backtracking), which is represented by an “/..” or “\..” or their URL encoding.
Redirect URIs cannot contain open redirects .
Redirect URIs cannot contain the fragment component.
Redirect URIs cannot contain certain characters including:
Wildcard characters ('*')
Non-printable ASCII characters
Invalid percent encodings (any percent encoding that does not follow URL-encoding form of a percent sign followed by two hexadecimal digits)
Null characters (an encoded NULL character, e.g., %00, %C0%80 )
If the redirect_uri passed in the authorization request does not match an authorized redirect URI for the OAuth client ID, you will receive a redirect_uri_mismatch error.
Note : It may take 5 minutes to a few hours for changes made to these settings to take effect
If your application is going to be installed on a device or computer (such as a system running Android, iOS, Universal Windows Platform, Chrome, or any desktop OS), you can use Google's OAuth 2.0 Mobile and desktop apps flow. If your application runs on devices with limited input capabilities, such as smart TVs, you can use Google’s OAuth 2.0 TV and limited-input device flow.
Note : Currently, obtaining OAuth 2.0 access tokens via AccountManager works for Android Ice Cream Sandwich (4.0) and newer versions.
You need to specify your Android app's package name and SHA1 fingerprint.
In the Package name field, enter your Android app's package name .
In a terminal, run the keytool utility to get the SHA1 fingerprint for your digitally signed .apk file's public certificate.
keytool -list -v -keystore path-to-debug-or-production-keystore -alias androiddebugkey
Note: For the debug.keystore , the password is android. For Android Studio, the debug keystore is typically located at ~/.android/debug.keystore .
The Keytool prints the fingerprint to the shell. For example:
$ keytool -list -v -keystore ~/.android/debug.keystore
Enter keystore password: Type "android" if using debug.keystore
Keystore type: JKS
Keystore provider: SUN
Your keystore contains 1 entry
Alias name: androiddebugkey
Creation date: Mar 13, 2020
Entry type: PrivateKeyEntry
Certificate chain length: 1
Certificate[1]:
Owner: C=US, O=Android, CN=Android Debug
Issuer: C=US, O=Android, CN=Android Debug
Serial number: 1
Valid from: Fri Mar 13 09:59:25 PDT 2020 until: Sun Mar 06 08:59:25 PST 2050
Certificate fingerprints:
SHA1: D9:E9:59:FA:7A:46:72:4E:69:1F:96:18:8C:F9:AE:82:3A:5D:2F:03
SHA256: 92:59:1E:F4:C9:BC:72:43:1C:59:57:24:AD:78:CA:A2:DB:C7:C5:AC:B1:A3:E8:52:04:B2:00:37:53:04:0B:8E
Signature algorithm name: SHA1withRSA
Subject Public Key Algorithm: 2048-bit RSA key
Version: 1
Copy the SHA1 fingerprint from the results that appear in your terminal.
Important: When you prepare to release your app to your users, follow these steps again in a production project and create a new OAuth 2.0 client ID for your production app. For production apps, use your own private key to sign the production app's .apk file. For more information, see Signing your applications .
Paste the SHA1 fingerprint into the form where requested.
(Optional) Verify ownership of your Android application.
You can verify ownership of your Android application to reduce the risk of app impersonation. Learn more about verifying ownership of your Android application .
If your application accesses APIs directly from iOS, you will need the application's Bundle ID and, optionally, its Apple App Store ID and Team ID:
The application's Bundle ID is the bundle identifier as listed in the app's .plist file. For example: com.example.myapp .
The application's App Store ID is in the app's App Store URL, if the app was published in the Apple App Store. For example, in the app URL https://apps.apple.com/us/app/google/id284815942 , the App Store ID is 284815942 .
The application's Team ID is a 10-character string that Apple assigns to your team. For information about your Team ID, see Locating your Team ID in the Apple App Distribution Guide.
After creating your iOS credentials and obtaining a Client ID, you use the Installed Application OAuth 2.0 flow to communicate with Google APIs.
When you enable App Check, Apple's App Attest service is used to verify that OAuth 2.0 requests originating from your OAuth client are genuine and come from your app. This helps to reduce the risk of app impersonation. Learn more about enabling App Check for your iOS app .
To enable App Check, turn on the Protect your OAuth client from abuse with Firebase App Check toggle button in the create/edit view of your iOS client.
The following requirements must be met to successfully enable App Check for your iOS client:
You must specify a team ID for your iOS client.
You must not use a wildcard in your bundle ID since it can resolve to more than one app. This means that the bundle ID must not include the asterisk (*) symbol.
Google Chrome apps and extensions are a special case of installed applications. Chrome exposes JavaScript APIs to allow your Chrome apps and extensions to perform various operations. Some of these APIs rely on knowing the identity of the user who is signed in to Chrome. If you're writing a Chrome app or extension that calls APIs that need to know the user's identity, and you want your app or extension to get user authorization for these requests using OAuth 2.0, then choose Chrome as the platform when you create your credentials. You will need to enter your Chrome app or extension's Application ID . The Item ID is the last part of your Chrome Extension's Chrome Web Store URL. For more information about these APIs, see the User Authentication documentation .
You can verify ownership of your Chrome application to reduce the risk of app impersonation. Learn more about verifying ownership of your Chrome application .
To delete a client ID, go to the Clients page , check the box next to the ID you want to delete, and then click the DELETE button.
Before deleting a Client ID, ensure to check the ID is not in use by monitoring your traffic in the overview page .
You can restore deleted clients within 30 days of the deletion. To restore a recently deleted client, navigate to the Deleted credentials page to find a list of clients you recently deleted and click the RESTORE button for the client you want to restore.
Any client deleted over 30 days ago cannot be restored and is permanently deleted.
Note : Clients can also be automatically deleted if they become inactive. Learn more .
Client secrets or credentials should be treated with extreme care as described in the OAuth 2.0 policies , because they allow anyone who has them to use your app's identity to gain access to user information. With the client secret rotation feature, you can add a new secret to your OAuth client configuration, migrate to the new secret while the old secret is still usable, and disable the old secret afterwards. This is useful when the client secret has been inadvertently disclosed or leaked. This also ensures good security practices by occasionally rotating your secrets without causing downtime of your app. In addition, Google started to issue more secure client secrets recommended by RFC 6749 in 2021. While apps that were created earlier are able to continue using the old secrets, we recommend that you migrate to the new secret with this rotation feature.
To

[truncated]

## Original Extract

Your OAuth client is the credential which your application uses when making calls to Google OAuth 2.0 endpoint to receive an access token or ID token. After creating your OAuth client, you will receiv

Skip to main content Google Cloud Platform Console Help Sign in Google Help Help Center
Your OAuth client is the credential which your application uses when making calls to Google OAuth 2.0 endpoint to receive an access token or ID token. After creating your OAuth client, you will receive a client ID and sometimes, a client secret.
Think of your client ID like your app's unique username when it needs to request an access token or ID token from Google's OAuth 2.0 endpoint. This ID helps Google identify your app and ensure that only authorized applications can access user data.
Similar to how you would use a username and password to log to online services, many applications use a client ID paired with a client secret. The client secret adds an extra layer of security, acting like your app's password.
Applications are categorized as either public or private clients:
Private Clients: These apps, like web server applications, can securely store the client secret because they run on servers you control.
Public Clients: Native apps or JavaScript-based apps fall under this category. They cannot securely store secrets, as they reside on user devices and as such do not use client secrets.
To create an OAuth 2.0 client ID in the console:
Navigate to the Google Auth Platform Clients page.
You will be prompted to create a project if you do not have one selected.
You will be prompted to register your application to use Google Auth if you are yet to do so. This is required before creating a client.
Select the appropriate application type for your application and enter any additional information required. Application types are described in more detail in the following sections.
Fill out the required information for the select client type and click the CREATE button to create the client.
Note : Your application's client secret will only be shown after you create the client. Store this information in a secure place such as Google Cloud Secret Manager because it will not be visible or accessible again. Learn more .
A web application is accessed by web browsers over a network.
Applications that use client-side JavaScript to access Google APIs must specify authorized JavaScript origins. The origins identify the domains from which your application can send API requests.
Specified origins must adhere to the following rules :
JavaScript origins must use the HTTPS scheme, not plain HTTP. Localhost URIs (including localhost IP address URIs) are exempt from this rule.
Hosts cannot be raw IP addresses. Localhost IP addresses are exempted from this rule.
If you use a port other than 80, you must specify it. For example: https://example.com:8080
Host TLDs (Top Level Domains) must belong to the public suffix list .
Host domains cannot be “googleusercontent.com”.
JavaScript origins cannot contain URL shortener domains (e.g. goo.gl) unless the app owns the domain.
JavaScript origins cannot contain the userinfo subcomponent.
JavaScript origins cannot contain the path component.
JavaScript origins cannot contain the query component.
JavaScript origins cannot contain the fragment component.
JavaScript origins cannot contain certain characters including:
Wildcard characters ('*')
Non-printable ASCII characters
Invalid percent encodings (any percent encoding that does not follow URL-encoding form of a percent sign followed by two hexadecimal digits)
Null characters (an encoded NULL character, e.g., %00, %C0%80 )
If you send a request to a Google OAuth 2.0 endpoint from an unregistered JavaScript origin, you will receive an origin_mismatch error.
Applications that access Google APIs from a server (often using languages and frameworks like Node.js, Java, .NET, and Python) must specify authorized redirect URIs. The redirect URIs are the endpoints of your application server to which the OAuth 2.0 server can send responses. Users are redirected to this path after they have authenticated with Google.
Redirect URIs must adhere to the following rules :
Redirect URIs must use the HTTPS scheme, not plain HTTP. Localhost URIs (including localhost IP address URIs) are exempt from this rule.
Hosts cannot be raw IP addresses. Localhost IP addresses are exempted from this rule.
Host TLDs (Top Level Domains) must belong to the public suffix list .
Redirect URIs cannot contain URL shortener domains (e.g. goo.gl ) unless the app owns the domain. Furthermore, if an app that owns a shortener domain chooses to redirect to that domain, that redirect URI must either contain “ /google-callback/ ” in its path or end with “ /google-callback ”.
Redirect URIs cannot contain the userinfo subcomponent.
Redirect URIs cannot contain a path traversal (also called directory backtracking), which is represented by an “/..” or “\..” or their URL encoding.
Redirect URIs cannot contain open redirects .
Redirect URIs cannot contain the fragment component.
Redirect URIs cannot contain certain characters including:
Wildcard characters ('*')
Non-printable ASCII characters
Invalid percent encodings (any percent encoding that does not follow URL-encoding form of a percent sign followed by two hexadecimal digits)
Null characters (an encoded NULL character, e.g., %00, %C0%80 )
If the redirect_uri passed in the authorization request does not match an authorized redirect URI for the OAuth client ID, you will receive a redirect_uri_mismatch error.
Note : It may take 5 minutes to a few hours for changes made to these settings to take effect
If your application is going to be installed on a device or computer (such as a system running Android, iOS, Universal Windows Platform, Chrome, or any desktop OS), you can use Google's OAuth 2.0 Mobile and desktop apps flow. If your application runs on devices with limited input capabilities, such as smart TVs, you can use Google’s OAuth 2.0 TV and limited-input device flow.
Note : Currently, obtaining OAuth 2.0 access tokens via AccountManager works for Android Ice Cream Sandwich (4.0) and newer versions.
You need to specify your Android app's package name and SHA1 fingerprint.
In the Package name field, enter your Android app's package name .
In a terminal, run the keytool utility to get the SHA1 fingerprint for your digitally signed .apk file's public certificate.
keytool -list -v -keystore path-to-debug-or-production-keystore -alias androiddebugkey
Note: For the debug.keystore , the password is android. For Android Studio, the debug keystore is typically located at ~/.android/debug.keystore .
The Keytool prints the fingerprint to the shell. For example:
$ keytool -list -v -keystore ~/.android/debug.keystore
Enter keystore password: Type "android" if using debug.keystore
Keystore type: JKS
Keystore provider: SUN
Your keystore contains 1 entry
Alias name: androiddebugkey
Creation date: Mar 13, 2020
Entry type: PrivateKeyEntry
Certificate chain length: 1
Certificate[1]:
Owner: C=US, O=Android, CN=Android Debug
Issuer: C=US, O=Android, CN=Android Debug
Serial number: 1
Valid from: Fri Mar 13 09:59:25 PDT 2020 until: Sun Mar 06 08:59:25 PST 2050
Certificate fingerprints:
SHA1: D9:E9:59:FA:7A:46:72:4E:69:1F:96:18:8C:F9:AE:82:3A:5D:2F:03
SHA256: 92:59:1E:F4:C9:BC:72:43:1C:59:57:24:AD:78:CA:A2:DB:C7:C5:AC:B1:A3:E8:52:04:B2:00:37:53:04:0B:8E
Signature algorithm name: SHA1withRSA
Subject Public Key Algorithm: 2048-bit RSA key
Version: 1
Copy the SHA1 fingerprint from the results that appear in your terminal.
Important: When you prepare to release your app to your users, follow these steps again in a production project and create a new OAuth 2.0 client ID for your production app. For production apps, use your own private key to sign the production app's .apk file. For more information, see Signing your applications .
Paste the SHA1 fingerprint into the form where requested.
(Optional) Verify ownership of your Android application.
You can verify ownership of your Android application to reduce the risk of app impersonation. Learn more about verifying ownership of your Android application .
If your application accesses APIs directly from iOS, you will need the application's Bundle ID and, optionally, its Apple App Store ID and Team ID:
The application's Bundle ID is the bundle identifier as listed in the app's .plist file. For example: com.example.myapp .
The application's App Store ID is in the app's App Store URL, if the app was published in the Apple App Store. For example, in the app URL https://apps.apple.com/us/app/google/id284815942 , the App Store ID is 284815942 .
The application's Team ID is a 10-character string that Apple assigns to your team. For information about your Team ID, see Locating your Team ID in the Apple App Distribution Guide.
After creating your iOS credentials and obtaining a Client ID, you use the Installed Application OAuth 2.0 flow to communicate with Google APIs.
When you enable App Check, Apple's App Attest service is used to verify that OAuth 2.0 requests originating from your OAuth client are genuine and come from your app. This helps to reduce the risk of app impersonation. Learn more about enabling App Check for your iOS app .
To enable App Check, turn on the Protect your OAuth client from abuse with Firebase App Check toggle button in the create/edit view of your iOS client.
The following requirements must be met to successfully enable App Check for your iOS client:
You must specify a team ID for your iOS client.
You must not use a wildcard in your bundle ID since it can resolve to more than one app. This means that the bundle ID must not include the asterisk (*) symbol.
Google Chrome apps and extensions are a special case of installed applications. Chrome exposes JavaScript APIs to allow your Chrome apps and extensions to perform various operations. Some of these APIs rely on knowing the identity of the user who is signed in to Chrome. If you're writing a Chrome app or extension that calls APIs that need to know the user's identity, and you want your app or extension to get user authorization for these requests using OAuth 2.0, then choose Chrome as the platform when you create your credentials. You will need to enter your Chrome app or extension's Application ID . The Item ID is the last part of your Chrome Extension's Chrome Web Store URL. For more information about these APIs, see the User Authentication documentation .
You can verify ownership of your Chrome application to reduce the risk of app impersonation. Learn more about verifying ownership of your Chrome application .
To delete a client ID, go to the Clients page , check the box next to the ID you want to delete, and then click the DELETE button.
Before deleting a Client ID, ensure to check the ID is not in use by monitoring your traffic in the overview page .
You can restore deleted clients within 30 days of the deletion. To restore a recently deleted client, navigate to the Deleted credentials page to find a list of clients you recently deleted and click the RESTORE button for the client you want to restore.
Any client deleted over 30 days ago cannot be restored and is permanently deleted.
Note : Clients can also be automatically deleted if they become inactive. Learn more .
Client secrets or credentials should be treated with extreme care as described in the OAuth 2.0 policies , because they allow anyone who has them to use your app's identity to gain access to user information. With the client secret rotation feature, you can add a new secret to your OAuth client configuration, migrate to the new secret while the old secret is still usable, and disable the old secret afterwards. This is useful when the client secret has been inadvertently disclosed or leaked. This also ensures good security practices by occasionally rotating your secrets without causing downtime of your app. In addition, Google started to issue more secure client secrets recommended by RFC 6749 in 2021. While apps that were created earlier are able to continue using the old secrets, we recommend that you migrate to the new secret with this rotation feature.
To

[truncated]
