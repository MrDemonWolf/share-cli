---
layout: default
title: ShareX
nav_order: 1
parent: Supported Uploaders
---


# Getting started with ShareX
{: .fs-6 .fw-300 }

#### Params
{: .no_toc }

| Field       | Type   | Description                                            |
| :---------- | :----- | :----------------------------------------------------- |
| YOUR_DOMAIN | string | The APPs running domain.                               |
| VERSION     | string | API Version you want to use                            |
| YOUR_TOKEN  | string | Token created from your account.  On the /tokens page. |

```json
{
  "Name": "Share",
  "DestinationType": "ImageUploader, FileUploader,TextUploader",
  "RequestMethod": "POST",
  "RequestURL": "https://[YOUR_DOMAIN]/api/[VERSION]upload/",
  "Headers": {
    "Authorization": "Bearer [YOUR_TOKEN]"
  },
  "Body": "MultipartFormData",
  "FileFormName": "file",
  "URL": "$json:file.url$",
  "DeletionURL": "$json:file.delete$"
}
```
