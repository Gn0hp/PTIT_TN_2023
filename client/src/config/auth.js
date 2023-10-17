import axios from 'axios'

let authToken = ''

export function getAuthToken () {
  return authToken
}

export function setAuthToken (token) {
  authToken = token
}

export default {
  methods: {
    setAuthHeader (token) {
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`
    }
  }
}

export const AxiosInstance = axios.create({
  contentType: 'application/json'
})
