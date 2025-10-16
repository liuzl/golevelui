import http from './http'

// Use dynamic axios types to avoid conflicts with existing usage
export function dbs() {
  return http.get('/golevelui/api/dbs')
}

interface KeysOption {
  db: string
  searchText: string
  prefix: string
}

export function keys(option: KeysOption) {
  return http.get('/golevelui/api/db/keys', {
    params: option,
  })
}

interface KeyInfoOption {
  db: string
  key: string
}

export function keyInfo(option: KeyInfoOption) {
  return http.get('/golevelui/api/db/key/info', {
    params: option,
  })
}

interface KeyDeleteOption {
  db: string
  key: string
}

export function keyDelete(option: KeyDeleteOption) {
  return http.post('/golevelui/api/db/key/delete', option)
}

interface KeyUpdateOption {
  db: string
  key: string
  value: string
}

export function keyUpdate(option: KeyUpdateOption) {
  return http.post('/golevelui/api/db/key/update', option)
}

interface KeysCountOption {
  db: string
  searchText: string
  prefix: string
}

export function keysCount(option: KeysCountOption) {
  return http.get('/golevelui/api/db/keys/count', {
    params: option,
  })
}
