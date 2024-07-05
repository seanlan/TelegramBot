import request from '@/utils/request'

// taskList
export function taskList(data) {
  return request({
    url: `/api/v1/task/list`,
    method: 'post',
    data
  })
}

// taskSave
export function taskSave(data) {
  return request({
    url: `/api/v1/task/save`,
    method: 'post',
    data
  })
}
