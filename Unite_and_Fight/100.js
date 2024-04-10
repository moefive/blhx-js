// ==UserScript==
// @name         GBF剪切板跳转
// @version      1.0.3
// @description  acgp沙盒测试
// @icon         http://game.granbluefantasy.jp/favicon.ico
// @match        *://game.granbluefantasy.jp/*
// @match        *://gbf.game.mbga.jp/*
// @grant        GM_notification
// ==/UserScript==
(function () {
  'use strict';
  const AUDIO_URL = 'https://prd-game-a5-granbluefantasy.akamaized.net/assets/sound/voice/3040372000_ability_them2.mp3'
  const audio = new Audio(AUDIO_URL)
  const stickers = ['./mon_201903/16/fkQ5-4virK7T8S3c-3c.png']

   const send = (title) => {
     let sticker = stickers[Math.floor(Math.random() * stickers.length)]
   let image = `https://img.nga.178.com/attachments/${sticker.slice(2)}`
    GM_notification({
     title: title,
    text: '看一下',
    image: image,
      timeout: 10000
    })
     audio.volume = 0 //声音大小
     audio.play()
  }
  let eventOn = false
  window.addEventListener('hashchange', () => {
    let hash = location.hash
    if (/^#result(_multi)?\/\d/.test(hash)) {
      if (!eventOn) {
        eventOn = true
        $(document).ajaxSuccess(function(event, xhr, settings, data) {
          if (/\/result(multi)?\/content\/index\/\d+/.test(settings.url)) {           
              if (data.option.result_data.replicard?.has_occurred_event) {                 
                  window.location.replace(`https://game.granbluefantasy.jp/#${data.option.result_data.url}`)
              }else {
//                 跳转到url
                     location.hash = 'quest/supporter/917881/1/0/10116'
              }
          }
        })
      }
    }
  })
}())