import request from '@/utils/request'

// redpacketList
export function redpacketList(data) {
  return request({
    url: `/api/v1/redpacket/list`,
    method: 'post',
    data
  })
}

// redpacketAdd
export function redpacketAdd(data) {
  return request({
    url: `/api/v1/redpacket/add`,
    method: 'post',
    data
  })
}
