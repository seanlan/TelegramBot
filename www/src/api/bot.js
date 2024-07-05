import request from '@/utils/request'

// getBotList
export function getBotList(data) {
  return request({
    url: `/api/v1/bot/list`,
    method: 'post',
    data
  })
}

// saveBot
export function saveBot(data) {
  return request({
    url: `/api/v1/bot/save`,
    method: 'post',
    data
  })
}

// switchBotStatus
export function switchBotStatus(data) {
  return request({
    url: `/api/v1/bot/switch_status`,
    method: 'post',
    data
  })
}

// getActionList
export function getActionList(data) {
  return request({
    url: `/api/v1/action/list`,
    method: 'post',
    data
  })
}

// saveAction
export function saveAction(data) {
  return request({
    url: `/api/v1/action/save`,
    method: 'post',
    data
  })
}
