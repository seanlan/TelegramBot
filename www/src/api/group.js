import request from '@/utils/request'

// getGroupList
export function getGroupList(data) {
  return request({
    url: `/api/v1/group/list`,
    method: 'post',
    data
  })
}

// saveGroup
export function saveGroup(data) {
  return request({
    url: `/api/v1/group/save`,
    method: 'post',
    data
  })
}

// deleteGroup
export function deleteGroup(data) {
  return request({
    url: `/api/v1/group/delete`,
    method: 'post',
    data
  })
}

export function getMessagePushList(data) {
  return request({
    url: '/api/v1/message/list',
    method: 'post',
    data
  })
}

export function saveMessagePush(data) {
  return request({
    url: '/api/v1/message/save',
    method: 'post',
    data
  })
}

export function deleteMessagePush(data) {
  return request({
    url: '/api/v1/message/delete',
    method: 'post',
    data
  })
}
