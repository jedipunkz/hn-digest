---
source: "https://www.mellonbot.app"
hn_url: "https://news.ycombinator.com/item?id=50408834"
title: "Initial submission of Mellonbot – an AI agent for Linux server diagnostics"
article_title: "Mellonbot — AI agent that fixes website problems: downtime, hacks, slow pages"
image: ""
author: "narekkaravarich"
captured_at: "2026-11-07T23:28:12Z"
capture_tool: "hn-digest"
hn_id: 50408834
score: 1
comments: 0
posted_at: "2026-11-07T22:57:12Z"
tags:
  - hacker-news
---

# Initial submission of Mellonbot – an AI agent for Linux server diagnostics

- HN: [50408834](https://news.ycombinator.com/item?id=50408834)
- Source: [www.mellonbot.app](https://www.mellonbot.app)
- Score: 1
- Comments: 0
- Posted: 2026-11-07T22:57:12Z

## Translation

Title: Initial submission of Mellonbot – an AI agent for Linux server diagnostics
Article title: Mellonbot — AI agent that fixes website problems: downtime, hacks, slow pages
Description: Website down, slow, hacked or “Not secure”? Mellonbot’s AI investigates your site and the server behind it, explains the real cause in plain words and fixes it with your OK. Free plan.

Article text:
Mellonbot — AI agent that fixes website problems: downtime, hacks, slow pages mellonbot How it works What it fixes Security Pricing Sign in Get started AI website agent · free plan
Website down, slow or hacked? Mellonbot finds out why and fixes it.
Database errors, expired certificates, failing backups, attacks, slow pages. Mellonbot’s AI digs through your site and the server it runs on, tells you the real cause in plain words and fixes it once you say OK.
1 of 6 : Your site says “Error establishing a database connection”
At the busiest moment your site’s PHP used more memory than the server has, so the database was pushed out and crashed.
Mellonbot can restart the database and lower PHP’s limit so both fit.
rosebakery.com · 06:12 Root cause found
The site shows a database error. Is the database running?
● mysql.service — failed (Result: oom-kill)
It was shut down for using too much memory. What else needs memory?
$ ps -C php-fpm --no-headers | wc -l
14 PHP workers × ~160 MB on a 2 GB server
millerplumbing.co · 09:41 Root cause found
Visitors get a security warning. Is the certificate still valid?
$ curl -I https://millerplumbing.co
● SSL certificate problem: certificate has expired
It should renew itself. Why didn’t it?
$ systemctl status certbot.timer
certbot.timer — inactive (dead) since Jun 3
lakesideyoga.com · 03:00 Root cause found
No new backup since September 1. What does the backup job say?
$ tail -n 1 /var/log/backup.log
● /opt/backup.sh: line 12: mysqldump: command not found
The tool it needs is gone. Did an update rename it?
$ grep mariadb /var/log/apt/history.log | tail -1
Start-Date: 2026-09-01 · Upgrade: mariadb-client
greenleafdental.com · 14:27 Root cause found
The site got slow this afternoon. Is it real visitors?
$ grep -c "POST /wp-login.php" access.log
● 48213 login attempts in the last 24 hours
Bots guessing passwords. Which door are they using most?
$ grep -c "POST /xmlrpc.php" access.log
31870 — an old side door the site never uses
oakandthread.shop · 11:05 Root cause found
Photo uploads are failing. Is the disk full?
Completely full. What is taking up the space?
$ du -sh wp-content/* | sort -h | tail -1
21G wp-content/updraft (1,412 backups)
harborinn.com · 16:48 Root cause found
Routine look-over. How current is the software?
● PHP 7.4.33 — security support ended Nov 2022
Can the site move to a newer PHP without breaking?
$ wp plugin list --update=available
9 plugins out of date · 2 won’t run on PHP 8
rosebakery.com · 06:12 AI investigating
The site shows a database error. Is the database running?
$ s y s t e m c t l s t a t u s m y s q l
● mysql.service — failed (Result: oom-kill)
It was shut down for using too much memory. What else needs memory?
$ p s - C p h p - f p m - - n o - h e a d e r s | w c - l
14 PHP workers × ~160 MB on a 2 GB server
Your site says “Error establishing a database connection”
At the busiest moment your site’s PHP used more memory than the server has, so the database was pushed out and crashed.
Mellonbot can restart the database and lower PHP’s limit so both fit.
Visitors see “Your connection is not private”
Your security certificate expired 12 days ago. Automatic renewal was switched off in June, so nothing noticed.
Mellonbot can renew it now and turn automatic renewal back on.
Your nightly backup hasn’t worked for 23 days
A server update on September 1 renamed a tool the backup script relies on, so it has been failing quietly every night.
Mellonbot can point the script at the new tool and take a fresh backup now.
Bots are trying to break into your site
Tens of thousands of password guesses a day, mostly through a side door your site doesn’t need, are slowing it down for real visitors.
Mellonbot can close the side door and slow down repeat guessers.
Old backups have filled up your whole server
A backup plugin has been saving copies for years and never deleting them. Now you can’t upload photos or take new orders.
Mellonbot can keep the five newest backups and clear out the rest.
Your site runs on software that stopped getting security fixes in 2022
Outdated PHP is one of the most common ways small sites get hacked — and two of your plugins would break if it were simply upgraded.
Mellonbot can update the plugins first, then walk you through upgrading PHP safely.
Your site says “Error establishing a database connection”
At the busiest moment your site’s PHP used more memory than the server has, so the database was pushed out and crashed.
Mellonbot can restart the database and lower PHP’s limit so both fit.
Websites break at the worst possible time
Sites go down, slow down, get attacked or quietly stop sending email. Finding out why usually means logs, forums, plugins and guesswork.
“ Black Friday, the busiest night of the year, and checkout just shows a white screen. ”
“ My site is down. The host says “check nginx”. I don’t even know what nginx is. ”
“ A customer texted me a screenshot: “Your website says it’s not secure.” I had no idea. ”
“ The site says 502 and I’m half asleep, digging through logs to find out which service died. ”
“ Our backups had been failing for a month. I only found out when I actually needed one. ”
“ Patients can’t book online, and the person who set up our server moved abroad years ago. ”
“ Twenty client servers. One full disk nobody noticed, and a certificate that expired over the weekend. ”
“ My site got so slow that visitors leave before the page even finishes loading. ”
“ Black Friday, the busiest night of the year, and checkout just shows a white screen. ”
“ My site is down. The host says “check nginx”. I don’t even know what nginx is. ”
“ A customer texted me a screenshot: “Your website says it’s not secure.” I had no idea. ”
“ The site says 502 and I’m half asleep, digging through logs to find out which service died. ”
“ Our backups had been failing for a month. I only found out when I actually needed one. ”
“ Patients can’t book online, and the person who set up our server moved abroad years ago. ”
“ Twenty client servers. One full disk nobody noticed, and a certificate that expired over the weekend. ”
“ My site got so slow that visitors leave before the page even finishes loading. ”
“ The contact form stopped sending emails weeks ago. I only noticed when a customer phoned to complain. ”
“ A client’s WordPress got hacked through a plugin nobody had updated in two years. Goodbye, weekend. ”
“ Google emailed me about “server errors” on my site. Is that bad? I honestly can’t tell. ”
“ Logs never rotated, the disk hit 100% at 3 am and the database quietly stopped saving orders. ”
“ My host says my server was sending spam. Now my booking emails land in people’s junk folders. ”
“ Someone spent the whole night trying to guess the root password. Thousands of attempts. ”
“ Once a month I lose a whole day logging into client sites just to check they’re still OK. ”
“ The site takes eight seconds to load, and three speed plugins later it still does. ”
“ The contact form stopped sending emails weeks ago. I only noticed when a customer phoned to complain. ”
“ A client’s WordPress got hacked through a plugin nobody had updated in two years. Goodbye, weekend. ”
“ Google emailed me about “server errors” on my site. Is that bad? I honestly can’t tell. ”
“ Logs never rotated, the disk hit 100% at 3 am and the database quietly stopped saving orders. ”
“ My host says my server was sending spam. Now my booking emails land in people’s junk folders. ”
“ Someone spent the whole night trying to guess the root password. Thousands of attempts. ”
“ Once a month I lose a whole day logging into client sites just to check they’re still OK. ”
“ The site takes eight seconds to load, and three speed plugins later it still does. ”
Other tools tell you what broke. Mellonbot works out why.
An alert leaves the hard part to you. Mellonbot’s AI asks follow-up questions, checks the answers and follows the clues until it finds the real cause — and the fix that stops it coming back.
Typical monitoring: CRITICAL · Disk usage 100% on /
What’s using the space? 21 GB in one folder — old backups from a plugin.
Why does it keep growing? The plugin saves a copy every night and never deletes any.
What has it broken? Photo uploads and new orders have failed since 11:05.
What happened: A backup plugin has saved a copy every night for three years. The disk filled up this morning, so your shop can’t save new orders.
Fix: Keep the five newest backups, delete the rest and make the plugin tidy up after itself.
CRITICAL · Disk usage 100% on /
That’s where it stops. Why it happened and what to do is up to you.
Mellonbot Root cause found What’s using the space? 21 GB in one folder — old backups from a plugin.
Why does it keep growing? The plugin saves a copy every night and never deletes any.
What has it broken? Photo uploads and new orders have failed since 11:05.
A backup plugin has saved a copy every night for three years. The disk filled up this morning, so your shop can’t save new orders.
Keep the five newest backups, delete the rest and make the plugin tidy up after itself.
DOWN · rosebakery.com returned 500
That’s where it stops. Why it happened and what to do is up to you.
Mellonbot Root cause found Is the website itself running? Yes, but it can’t reach its database.
Why is the database down? It was shut down at 06:12 for using too much memory.
What used up the memory? PHP is allowed 14 workers — more than a 2 GB server can hold.
When traffic peaked, your site’s PHP used more memory than the server has and pushed the database out.
Restart the database and lower PHP’s limit so the two fit side by side.
WARNING · CPU above 95% for 30 min
That’s where it stops. Why it happened and what to do is up to you.
Mellonbot Root cause found What’s keeping the server busy? PHP, answering login requests.
Are these real visitors? No — 48,213 login attempts from bots in a day.
How are they getting in? Mostly through an old side door the site never uses.
Bots are guessing passwords all day, and your server is busy answering them instead of your customers.
Close the side door and slow down anyone who keeps guessing.
Every check is green, so nobody looks.
Mellonbot Root cause found When was the last good backup? September 1 — 23 days ago.
What does the backup job say? It can’t find the tool it uses to copy the database.
What changed on September 1? A server update renamed that tool.
Your backups have failed quietly every night since an update on September 1. Nothing looked wrong, so nobody noticed.
Point the backup at the renamed tool and take a fresh backup now.
CRITICAL · Disk usage 100% on /
That’s where it stops. Why it happened and what to do is up to you.
Mellonbot Investigating What’s using the space? 21 GB in one folder — old backups from a plugin.
Why does it keep growing? The plugin saves a copy every night and never deletes any.
What has it broken? Photo uploads and new orders have failed since 11:05.
A backup plugin has saved a copy every night for three years. The disk filled up this morning, so your shop can’t save new orders.
Keep the five newest backups, delete the rest and make the plugin tidy up after itself.
From “something's wrong” to fixed in four steps
Run one command on your server — it installs a small agent that only connects out.
Ask, and it takes a read-only look at your site, SSL, services, disk, updates, security, databases and more.
The AI follows the clues — logs, recent changes, what depends on what — to the real cause, then explains it in everyday words.
You approve, and Mellonbot applies the fix, checks it worked and keeps a backup so you can undo it.
The problems that actually take sites down
Each check is where the thinking starts: when one fails, Mellonbot works out why and what will fix it for good.
You decide how much Mellonbot does on its own
Pick how much Mellonbot may do on its own — and watch what happens to the same four fixes. You can change it any time.
Ask Default You approve every fix before it runs. Auto Low-

[truncated]

## Original Extract

Website down, slow, hacked or “Not secure”? Mellonbot’s AI investigates your site and the server behind it, explains the real cause in plain words and fixes it with your OK. Free plan.

Mellonbot — AI agent that fixes website problems: downtime, hacks, slow pages mellonbot How it works What it fixes Security Pricing Sign in Get started AI website agent · free plan
Website down, slow or hacked? Mellonbot finds out why and fixes it.
Database errors, expired certificates, failing backups, attacks, slow pages. Mellonbot’s AI digs through your site and the server it runs on, tells you the real cause in plain words and fixes it once you say OK.
1 of 6 : Your site says “Error establishing a database connection”
At the busiest moment your site’s PHP used more memory than the server has, so the database was pushed out and crashed.
Mellonbot can restart the database and lower PHP’s limit so both fit.
rosebakery.com · 06:12 Root cause found
The site shows a database error. Is the database running?
● mysql.service — failed (Result: oom-kill)
It was shut down for using too much memory. What else needs memory?
$ ps -C php-fpm --no-headers | wc -l
14 PHP workers × ~160 MB on a 2 GB server
millerplumbing.co · 09:41 Root cause found
Visitors get a security warning. Is the certificate still valid?
$ curl -I https://millerplumbing.co
● SSL certificate problem: certificate has expired
It should renew itself. Why didn’t it?
$ systemctl status certbot.timer
certbot.timer — inactive (dead) since Jun 3
lakesideyoga.com · 03:00 Root cause found
No new backup since September 1. What does the backup job say?
$ tail -n 1 /var/log/backup.log
● /opt/backup.sh: line 12: mysqldump: command not found
The tool it needs is gone. Did an update rename it?
$ grep mariadb /var/log/apt/history.log | tail -1
Start-Date: 2026-09-01 · Upgrade: mariadb-client
greenleafdental.com · 14:27 Root cause found
The site got slow this afternoon. Is it real visitors?
$ grep -c "POST /wp-login.php" access.log
● 48213 login attempts in the last 24 hours
Bots guessing passwords. Which door are they using most?
$ grep -c "POST /xmlrpc.php" access.log
31870 — an old side door the site never uses
oakandthread.shop · 11:05 Root cause found
Photo uploads are failing. Is the disk full?
Completely full. What is taking up the space?
$ du -sh wp-content/* | sort -h | tail -1
21G wp-content/updraft (1,412 backups)
harborinn.com · 16:48 Root cause found
Routine look-over. How current is the software?
● PHP 7.4.33 — security support ended Nov 2022
Can the site move to a newer PHP without breaking?
$ wp plugin list --update=available
9 plugins out of date · 2 won’t run on PHP 8
rosebakery.com · 06:12 AI investigating
The site shows a database error. Is the database running?
$ s y s t e m c t l s t a t u s m y s q l
● mysql.service — failed (Result: oom-kill)
It was shut down for using too much memory. What else needs memory?
$ p s - C p h p - f p m - - n o - h e a d e r s | w c - l
14 PHP workers × ~160 MB on a 2 GB server
Your site says “Error establishing a database connection”
At the busiest moment your site’s PHP used more memory than the server has, so the database was pushed out and crashed.
Mellonbot can restart the database and lower PHP’s limit so both fit.
Visitors see “Your connection is not private”
Your security certificate expired 12 days ago. Automatic renewal was switched off in June, so nothing noticed.
Mellonbot can renew it now and turn automatic renewal back on.
Your nightly backup hasn’t worked for 23 days
A server update on September 1 renamed a tool the backup script relies on, so it has been failing quietly every night.
Mellonbot can point the script at the new tool and take a fresh backup now.
Bots are trying to break into your site
Tens of thousands of password guesses a day, mostly through a side door your site doesn’t need, are slowing it down for real visitors.
Mellonbot can close the side door and slow down repeat guessers.
Old backups have filled up your whole server
A backup plugin has been saving copies for years and never deleting them. Now you can’t upload photos or take new orders.
Mellonbot can keep the five newest backups and clear out the rest.
Your site runs on software that stopped getting security fixes in 2022
Outdated PHP is one of the most common ways small sites get hacked — and two of your plugins would break if it were simply upgraded.
Mellonbot can update the plugins first, then walk you through upgrading PHP safely.
Your site says “Error establishing a database connection”
At the busiest moment your site’s PHP used more memory than the server has, so the database was pushed out and crashed.
Mellonbot can restart the database and lower PHP’s limit so both fit.
Websites break at the worst possible time
Sites go down, slow down, get attacked or quietly stop sending email. Finding out why usually means logs, forums, plugins and guesswork.
“ Black Friday, the busiest night of the year, and checkout just shows a white screen. ”
“ My site is down. The host says “check nginx”. I don’t even know what nginx is. ”
“ A customer texted me a screenshot: “Your website says it’s not secure.” I had no idea. ”
“ The site says 502 and I’m half asleep, digging through logs to find out which service died. ”
“ Our backups had been failing for a month. I only found out when I actually needed one. ”
“ Patients can’t book online, and the person who set up our server moved abroad years ago. ”
“ Twenty client servers. One full disk nobody noticed, and a certificate that expired over the weekend. ”
“ My site got so slow that visitors leave before the page even finishes loading. ”
“ Black Friday, the busiest night of the year, and checkout just shows a white screen. ”
“ My site is down. The host says “check nginx”. I don’t even know what nginx is. ”
“ A customer texted me a screenshot: “Your website says it’s not secure.” I had no idea. ”
“ The site says 502 and I’m half asleep, digging through logs to find out which service died. ”
“ Our backups had been failing for a month. I only found out when I actually needed one. ”
“ Patients can’t book online, and the person who set up our server moved abroad years ago. ”
“ Twenty client servers. One full disk nobody noticed, and a certificate that expired over the weekend. ”
“ My site got so slow that visitors leave before the page even finishes loading. ”
“ The contact form stopped sending emails weeks ago. I only noticed when a customer phoned to complain. ”
“ A client’s WordPress got hacked through a plugin nobody had updated in two years. Goodbye, weekend. ”
“ Google emailed me about “server errors” on my site. Is that bad? I honestly can’t tell. ”
“ Logs never rotated, the disk hit 100% at 3 am and the database quietly stopped saving orders. ”
“ My host says my server was sending spam. Now my booking emails land in people’s junk folders. ”
“ Someone spent the whole night trying to guess the root password. Thousands of attempts. ”
“ Once a month I lose a whole day logging into client sites just to check they’re still OK. ”
“ The site takes eight seconds to load, and three speed plugins later it still does. ”
“ The contact form stopped sending emails weeks ago. I only noticed when a customer phoned to complain. ”
“ A client’s WordPress got hacked through a plugin nobody had updated in two years. Goodbye, weekend. ”
“ Google emailed me about “server errors” on my site. Is that bad? I honestly can’t tell. ”
“ Logs never rotated, the disk hit 100% at 3 am and the database quietly stopped saving orders. ”
“ My host says my server was sending spam. Now my booking emails land in people’s junk folders. ”
“ Someone spent the whole night trying to guess the root password. Thousands of attempts. ”
“ Once a month I lose a whole day logging into client sites just to check they’re still OK. ”
“ The site takes eight seconds to load, and three speed plugins later it still does. ”
Other tools tell you what broke. Mellonbot works out why.
An alert leaves the hard part to you. Mellonbot’s AI asks follow-up questions, checks the answers and follows the clues until it finds the real cause — and the fix that stops it coming back.
Typical monitoring: CRITICAL · Disk usage 100% on /
What’s using the space? 21 GB in one folder — old backups from a plugin.
Why does it keep growing? The plugin saves a copy every night and never deletes any.
What has it broken? Photo uploads and new orders have failed since 11:05.
What happened: A backup plugin has saved a copy every night for three years. The disk filled up this morning, so your shop can’t save new orders.
Fix: Keep the five newest backups, delete the rest and make the plugin tidy up after itself.
CRITICAL · Disk usage 100% on /
That’s where it stops. Why it happened and what to do is up to you.
Mellonbot Root cause found What’s using the space? 21 GB in one folder — old backups from a plugin.
Why does it keep growing? The plugin saves a copy every night and never deletes any.
What has it broken? Photo uploads and new orders have failed since 11:05.
A backup plugin has saved a copy every night for three years. The disk filled up this morning, so your shop can’t save new orders.
Keep the five newest backups, delete the rest and make the plugin tidy up after itself.
DOWN · rosebakery.com returned 500
That’s where it stops. Why it happened and what to do is up to you.
Mellonbot Root cause found Is the website itself running? Yes, but it can’t reach its database.
Why is the database down? It was shut down at 06:12 for using too much memory.
What used up the memory? PHP is allowed 14 workers — more than a 2 GB server can hold.
When traffic peaked, your site’s PHP used more memory than the server has and pushed the database out.
Restart the database and lower PHP’s limit so the two fit side by side.
WARNING · CPU above 95% for 30 min
That’s where it stops. Why it happened and what to do is up to you.
Mellonbot Root cause found What’s keeping the server busy? PHP, answering login requests.
Are these real visitors? No — 48,213 login attempts from bots in a day.
How are they getting in? Mostly through an old side door the site never uses.
Bots are guessing passwords all day, and your server is busy answering them instead of your customers.
Close the side door and slow down anyone who keeps guessing.
Every check is green, so nobody looks.
Mellonbot Root cause found When was the last good backup? September 1 — 23 days ago.
What does the backup job say? It can’t find the tool it uses to copy the database.
What changed on September 1? A server update renamed that tool.
Your backups have failed quietly every night since an update on September 1. Nothing looked wrong, so nobody noticed.
Point the backup at the renamed tool and take a fresh backup now.
CRITICAL · Disk usage 100% on /
That’s where it stops. Why it happened and what to do is up to you.
Mellonbot Investigating What’s using the space? 21 GB in one folder — old backups from a plugin.
Why does it keep growing? The plugin saves a copy every night and never deletes any.
What has it broken? Photo uploads and new orders have failed since 11:05.
A backup plugin has saved a copy every night for three years. The disk filled up this morning, so your shop can’t save new orders.
Keep the five newest backups, delete the rest and make the plugin tidy up after itself.
From “something's wrong” to fixed in four steps
Run one command on your server — it installs a small agent that only connects out.
Ask, and it takes a read-only look at your site, SSL, services, disk, updates, security, databases and more.
The AI follows the clues — logs, recent changes, what depends on what — to the real cause, then explains it in everyday words.
You approve, and Mellonbot applies the fix, checks it worked and keeps a backup so you can undo it.
The problems that actually take sites down
Each check is where the thinking starts: when one fails, Mellonbot works out why and what will fix it for good.
You decide how much Mellonbot does on its own
Pick how much Mellonbot may do on its own — and watch what happens to the same four fixes. You can change it any time.
Ask Default You approve every fix before it runs. Auto Low-

[truncated]
