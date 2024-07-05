import request from '@/utils/request'

export function paymentList(data) {
  return request({
    url: '/api/v1/payment/list',
    method: 'post',
    data
  })
}
