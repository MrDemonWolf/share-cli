---
layout: default
title: Changelog
nav_order: 11
---

# Changelog
4.2.0
- Added tags to uploads which you can add when uploading and or after the fact.
- Added tags to the API v1 in the body.
- Fixed bug where you search for something that is not found it would spam the server.
- Added link shorter (with features such as link click tracking and limit of clicks.) also can get the link data from the API, delete the link , and create a link.
- Added Upload file name set via editing it after and also being able to add it while uploading via the API v2.
- v2 API will be taking over API v1 in future updates.  Would be good to move over as soon as you can.
- View page is back for support with currently only being images.
- Fixed some of the validation.
- Added support for exporting of your data such as links,uploads, account and other data.
- Added links to what get removed when you delete your account.
- Added focus to the modals.
- Removed old userlist code.
- Added terms page editer in owner plus a page for it to be displayed if it's enabled.
- Added a logger for logging the made on the app.
- Added new env variable for LOGGER so you can enable or disable that feature.
- Added better stableness due to try catching everything,

4.1.0
- Fixed bug in pm2 echosystem file.
- Fixed missing UPLOAD_LIMIT env.
- Fixed bug where it shows dev even in prod in footer.
- Fixed signup disabled middleware.
- Removed the display of signup links on all pages if they are disabled.
- You can now disable the /owner route to make it return a 404.
- Fixed front-end bug where it wont display the right token created date.
- Fixed the tests bug due to not removing mfa on the test account.
- Fixed bug where if you edited yourself in the admin panel it will make you a user.
- Fixed bug where the last login date wont show the right one.
- Added account space used.  WIth rate limited requests.
- Added admin dashboard space used.  And removed the users count.
- When logged in you can now make config files for supported uploads by pasting your token in and then picking a uploader.
- Added tests for the config.
- Fixed bug on check for username or email when signing up.

4.0.1
* Fixed a bug where users can still create accounts even when signups are disabled (Hotfix)

4.0.0

__This is a fair big of a update but this is a list "all" changes that have been made__

* Made a lot of the code cleaner and easier to mange for developers
* Added last login date and location
* Added logo and favicon support
* Added Service worker and PWA support (this is in very beta).
* Switched to bower for frameworks (Must npm i again).
* Added transfer ownership.
* Added MFA.
* Added a verify checkmark.
* Redesigned and redo of the upload,token and user lising pages.
* Added Tests.
* Removed avatars.
* Added new docs.
* Switched to sendgrid official mailer
* Changed from csurf to lusca.
* Changed the routes.
* Changed from passport-local-mongoose to passport-local.
* Integrated Docker both for development and production modes.  Thanks to @exia for adding this.
* Adding a better way to handle the emails.  As there is so many templates that are just reused.  For easier reuse.
* Added Streamer mode.  It will stop leaks for both the user and if they have admin others as well.
* Now you can make yourself owner  by going to /owner.  If the email matches the one in the env then it will change the account to owner (This is safe as the email has to be verified anyways which makes sure its yours.).
* API route has been changed file and image route to just be as one.
* All upload lists have been reworked to be faster.
* You can now limit the size of each upload by default its 100M.
* Fixed dashboard so it will add a 's' when there is more the one user or upload.
* Added suspend users 24 hours , a week, and even a month. With unsuspend.
* Added ban and unban users. Confirm the ban of a user has been added.
* Changed text from 'Create user' to 'Create new user' in admin users.
* Fixed the link in the admin nav gallery.  Now it links to the gallery.
* Fixed the navbar in admin gallery as it was linking the wrong one.

3.1.2
* Hotfix for createdAt date in the upload.js. Now when  you upload it should show the right date.

3.1.1
* Fixed bug with middleware.

3.1.0
* Removed File and Text views.
* Removed Removed text support due to others being better. (You should use github gists if anything.).

3.0.1
* Added Dec to image and files.

3.0.0
* Better Gallery
* Added support for proxy more.
* Now the upload link will add https based on the proxy and even the domain of the server.
* Fixed login bug with tokens via the proxy trust.
* Updated env template.
* Added psd, doc, docx, xls, xlsx file support.
* Added better ZIP upload support.
* Changed to npm start from npm run web.
* Added templates for sharex imput (All you have to do is replace the name and add your API key).
* Added robots.txt.

2.0.0
* Many changes
* Better upload URLs
* View pages for files,images, and text

1.0.0
* First public build.
