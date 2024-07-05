import request from '@/utils/request'

// redeemList
export function redeemList(data) {
  return request({
    url: `/api/v1/redeem/list`,
    method: 'post',
    data
  })
}

// redeemAdd
export function redeemAdd(data) {
  return request({
    url: `/api/v1/redeem/add`,
    method: 'post',
    data
  })
}
