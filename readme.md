# CUHWC Signupper

This repo contains a simple static web page that allows the easy collection of emails
for signups at freshers fairs.

## What it does:
- Shows a simple, offline (as in once it's loaded you don't need internet) page with
  an email input.
- That lets people enter their CRS ID...
- ... or their email if they are non-cambridge
- and saves the resulting data to the browser's [local storage](https://developer.mozilla.org/en-US/docs/Web/API/Web_Storage_API)
- with a download button for later collection of the data

## Why not...
- ... A piece of paper? Because peoples' handwriting sucks and then someone has to type
  it all up into a spreadsheet.
- ... A google sheet? Because the spreadsheet interface is awkward for people to
  navigate at a stall and also they can see/maybe mess up previous entries. Also you
  might not have reliable internet.
- ... A google form? Because it looks a bit shit and requires multiple clicks. Also the
  internet again.


## How to use
- Go to https://cuhwc.github.io/signupper
- Point your laptop at passing freshers
- Once you're done, click the download button in the bottom right to download a JSON file with all the collected emails.

## ❗️ Important Notes
The signups collected are stored on your laptop, in your browser's data store. This data should be safe across closing the tab/browser, restarting the laptop, etc. But there are some things to be aware of:
  - If you use incognito/private browsing mode, **this data will be lost when you close the tab!** This is because these modes clear all persistent data from during their use when you leave them. I recommend you don't use incognito.
  - If you use a different browser or different laptop, you will not see the signups - again, the data is only stored on your laptop - not in some cloud service.
