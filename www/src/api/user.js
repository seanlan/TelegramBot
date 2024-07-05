import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/api/v1/login/login',
    method: 'post',
    data
  })
}

export function logout() {
  return request({
    url: '/api/v1/login/logout',
    method: 'post'
  })
}

export function rePassword(data) {
  return request({
    url: '/api/v1/login/change_password',
    method: 'post',
    data
  })
}
