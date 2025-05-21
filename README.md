# Preliminary Test For LFX Mentorship@Antrea(Renovate)

## Simple Chat Application using Websockets

- This is a simple app that creates a websocket server through  which multiple websocket clients can connect and communicate with one another.
- The Screenshot shows the working of the websockets.
  ![Screenshot from 2025-05-21 13-54-50](https://github.com/user-attachments/assets/ec75c7f1-aa1b-4190-ba3d-c438326b78c9)


## Vulerability

- The Go module `gorilla/websockets@v1.4.0` had a `High` type vulnerability according to the CVSS:3.1 with a score of 7.5.
- I first downloaded the `Mind Renovate` app on my repo and then setup the `renovate.json` with required configurations.
- I also had to do additional configurations, toggling the `Silent Mode` Off and the `Automated PRs` On as shown in the below image -
![Screenshot from 2025-05-21 13-48-10](https://github.com/user-attachments/assets/5592a4b1-2666-428f-b760-9c4be1d79f22)
- Then, the `Renovate Bot` automatically creates a PR for the vulnerable go module.

## Links I Used For Reference

- https://nvd.nist.gov/vuln/detail/cve-2020-27813
- https://docs.renovatebot.com/configuration-options/#vulnerabilityalerts
