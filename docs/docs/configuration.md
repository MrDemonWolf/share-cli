---
layout: default
title: Configuration
nav_order: 2
---

# Configuration

{: .no_toc }

All of Share configuration is saved in the .env and the database as well.
{: .fs-6 .fw-300 }

## Table of contents

{: .no_toc .text-delta }

1. TOC
   {:toc}

---

View [.env.example](https://github.com/MrDemonWolf/share/blob/master/.env.example) file as an example.

## URL

```yaml
# Domain or ip of the server running share
server:
  url: "http://localhost:8080"
```

## apikey

```yaml
# This is the API key you got from the share hosted app.
creds:
  apikey: >-
    eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1OTE3NjI4NTcsImV4cCI6NDc0NTM2Mjg1NywiaXNzIjoiU2hhcmUiLCJzdWIiOiI1ZWM1NTFlNzUwMGM4NDBmNzc0ZjVjNDQifQ.soQUFbTYXTQkqHhyA6eNwsq91R7BeIKNlJZfRxnzxJs
```
