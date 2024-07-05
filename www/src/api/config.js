import request from '@/utils/request'

// getConfigList
export function getConfigList(data) {
  return request({
    url: `/api/v1/config/list`,
    method: 'post',
    data
  })
}

// saveConfig
export function saveConfig(data) {
  return request({
    url: `/api/v1/config/save`,
    method: 'post',
    data
  })
}

// refreshConfig
export function refreshConfig(data) {
  return request({
    url: `/api/v1/config/refresh`,
    method: 'post',
    data
  })
}
