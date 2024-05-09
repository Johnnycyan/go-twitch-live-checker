<p align="center">
  <img src="https://raw.githubusercontent.com/Johnnycyan/Twitch-APIs/main/OneMoreDayIcon.svg" width="100" alt="project-logo">
</p>
<p align="center">
    <h1 align="center">GO-TWITCH-LIVE-CHECKER</h1>
</p>
<p align="center">
    <em>Check if a Twitch streamer is still alive... I mean live!</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/last-commit/Johnnycyan/go-twitch-live-checker?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/Johnnycyan/go-twitch-live-checker?style=default&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/Johnnycyan/go-twitch-live-checker?style=default&color=0080ff" alt="repo-language-count">
<p>
<p align="center">
	<!-- default option, no dependency badges. -->
</p>

<hr>

##  Overview

The go-twitch-live-checker project is designed to provide a streamlined solution for checking the live status of Twitch channels efficiently. By utilizing a caching mechanism, the software minimizes redundant API calls, ensuring fast retrieval of channel statuses. The core functionality involves handling HTTP requests, verifying channel availability, and maintaining cached results to optimize performance. The projects value lies in offering a convenient and resource-friendly method for monitoring Twitch channel live statuses, enhancing user experience and reducing load on the Twitch API.

---

##  Getting Started

**System Requirements:**

* **Internet**

###  Installation

<h4>From <code>releases</code></h4>

> 1. Download the latest release:
>     1. [Latest Releases](https://github.com/Johnnycyan/go-twitch-live-checker/releases) 

###  Usage

<h4>From <code>releases</code></h4>

> Run go-twitch-live-checker using the command below:
> ```console
> $ ./go-twitch-live-checker <port>
> ```
> or
> ```console
> $ go-twitch-live-checker.exe <port>
> ```
Endpoint      |     URL
------------- | -------------
Format  | <code>localhost:\<port>/?channel=\<channel-to-check></code>

---
