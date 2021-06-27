# Spotify CLI

<img src="https://storage.googleapis.com/pr-newsroom-wp/1/2018/11/Spotify_Logo_RGB_Green.png" width="400">

## Description
Control an existing Spotify session without leaving the terminal.
- Support for Linux, MacOS, and Windows
- Download once, and keep up to date with `spotify update`
- Control playback for music and podcasts
- Play and queue songs by name
- Command aliases and autocompletion

## Installation
Get the latest version <a href="https://github.com/brianstrauch/spotify-cli/releases/latest">here</a>.

## Usage
```
$ spotify login
Logged in as Brian Strauch.
$ spotify play back pocket
🎵 Back Pocket
🎤 Vulfpeck
▶️ 0:00 [                ] 3:01
$ spotify pause
🎵 Back Pocket
🎤 Vulfpeck
⏸ 1:30 [========        ] 3:01
$ spotify queue blinding lights
Queued!
$ spotify next
🎵 Blinding Lights
🎤 The Weeknd
▶️ 0:00 [                ] 3:20
$ spotify status
🎵 Blinding Lights
🎤 The Weeknd
▶️ 0:20 [==              ] 3:20
$ spotify back
🎵 Back Pocket
🎤 Vulfpeck
▶️ 0:00 [                ] 3:01
$ spotify shuffle on
🔀 Shuffle on
$ spotify repeat off
🔁 Repeat off
```

## Aliases
<table>
  <tr>
    <td><code>spotify back</code></td>
    <td><code>spotify b</code></td>
  </tr>
  <tr>
    <td><code>spotify next</code></td>
    <td><code>spotify n</code></td>
  </tr>
  <tr>
    <td><code>spotify play</code></td>
    <td><code>spotify p</code></td>
  </tr>
  <tr>
    <td><code>spotify pause</code></td>
    <td><code>spotify p</code></td>
  </tr>
  <tr>
    <td><code>spotify queue</code></td>
    <td><code>spotify q</code></td>
  </tr>
  <tr>
    <td><code>spotify status</code></td>
    <td><code>spotify s</code></td>
  </tr>
</table>
