import request from '@/utils/request'

export function getWithdrawList(data) {
  return request({
    url: '/api/v1/withdraw/list',
    method: 'post',
    data
  })
}
