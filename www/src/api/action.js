import request from '@/utils/request'

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

// deleteAction
export function deleteAction(data) {
  return request({
    url: `/api/v1/action/delete`,
    method: 'post',
    data
  })
}
